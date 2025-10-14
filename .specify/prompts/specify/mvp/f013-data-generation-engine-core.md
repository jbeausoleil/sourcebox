# Feature Specification Prompt: F013 - Data Generation Engine Core

## Feature Metadata
- **Feature ID**: F013
- **Name**: Data Generation Engine Core
- **Category**: Data Generation
- **Phase**: Week 6
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Large (5 days)
- **Dependencies**: F008 (Schema parser), F011 (Base generators), F012 (Custom generators)

## Constitutional Alignment

### Core Principles
- ✅ **Speed > Features**: < 30 seconds to generate 1,000 records
- ✅ **Verticalized > Generic**: Use schema-defined custom generators
- ✅ **Quality Standards**: Maintain referential integrity

### Technical Constraints
- ✅ **Performance**: 100+ records/second generation rate
- ✅ **Code Quality**: TDD required, >80% coverage

### Development Practices
- ✅ **TDD Required for Core Functionality**: Engine is critical core logic

## User Story
**US-MVP-008**: "As a CLI tool, I need a data generation engine that reads schemas, invokes correct generators, maintains referential integrity, and produces structured data ready for database insertion."

## Problem Statement
SourceBox needs a central engine that orchestrates data generation:
- Reads schema definitions (F008)
- Invokes correct generators (F011, F012) for each column
- Maintains referential integrity (foreign keys)
- Generates tables in correct order (parents before children)
- Produces structured data ready for SQL insertion
- Handles configurable record counts
- Generates data efficiently (>100 records/sec)

Without this engine, individual generators are useless - there's no way to coordinate them into complete, consistent datasets.

## Solution Overview
Implement a data generation engine in `pkg/generators/engine.go` that:
1. Takes a schema (from F008) as input
2. Validates generation order (parents before children)
3. For each table in order:
   - Creates specified number of records
   - For each column, invokes appropriate generator
   - Stores primary keys for foreign key references
   - Maintains referential integrity
4. Returns structured data (map of table → rows)
5. Ready for database insertion (F023, F024)

## Detailed Requirements

### Acceptance Criteria
1. **Engine Module**: `pkg/generators/engine.go` created
2. **Generate Function**: `Generate(schema *Schema, recordCounts map[string]int) (*Dataset, error)`
3. **Schema Reading**: Reads schema and validates generation order
4. **Generator Invocation**: Correctly invokes generators per column type
5. **Referential Integrity**: Maintains FK constraints (parent IDs exist)
6. **Generation Order**: Tables generated in correct order (parents first)
7. **Configurable Counts**: Override default record counts per table
8. **Structured Output**: Returns data ready for INSERT statements
9. **Unit Tests**: Test with fintech schema, verify integrity, 100% coverage

### Technical Specifications

#### Engine Core: `pkg/generators/engine.go`

```go
package generators

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yourusername/sourcebox/pkg/schema"
)

// Dataset represents generated data for all tables
type Dataset struct {
	Schema *schema.Schema
	Tables map[string]*TableData
}

// TableData represents generated data for a single table
type TableData struct {
	Name    string
	Columns []string
	Rows    []map[string]interface{}
}

// Engine orchestrates data generation
type Engine struct {
	schema    *schema.Schema
	dataset   *Dataset
	pkTracker map[string][]int // tracks generated primary keys for FK references
}

// NewEngine creates a new data generation engine
func NewEngine(s *schema.Schema) *Engine {
	return &Engine{
		schema:    s,
		pkTracker: make(map[string][]int),
		dataset: &Dataset{
			Schema: s,
			Tables: make(map[string]*TableData),
		},
	}
}

// Generate generates data for all tables in the schema
func (e *Engine) Generate() (*Dataset, error) {
	// Validate generation order
	if err := e.validateGenerationOrder(); err != nil {
		return nil, err
	}

	// Generate data for each table in order
	for _, tableName := range e.schema.GenerationOrder {
		table := e.findTable(tableName)
		if table == nil {
			return nil, fmt.Errorf("table not found in generation order: %s", tableName)
		}

		if err := e.generateTable(table); err != nil {
			return nil, fmt.Errorf("failed to generate table %s: %w", tableName, err)
		}
	}

	return e.dataset, nil
}

// validateGenerationOrder ensures parent tables come before children
func (e *Engine) validateGenerationOrder() error {
	generated := make(map[string]bool)

	for _, tableName := range e.schema.GenerationOrder {
		table := e.findTable(tableName)
		if table == nil {
			return fmt.Errorf("table in generation_order not found: %s", tableName)
		}

		// Check foreign keys reference already-generated tables
		for _, col := range table.Columns {
			if col.ForeignKey != nil {
				fkTable := col.ForeignKey.Table
				if !generated[fkTable] {
					return fmt.Errorf("table %s has FK to %s, but %s comes later in generation_order",
						tableName, fkTable, fkTable)
				}
			}
		}

		generated[tableName] = true
	}

	return nil
}

// generateTable generates data for a single table
func (e *Engine) generateTable(table *schema.Table) error {
	tableData := &TableData{
		Name:    table.Name,
		Columns: make([]string, len(table.Columns)),
		Rows:    make([]map[string]interface{}, 0, table.RecordCount),
	}

	// Extract column names
	for i, col := range table.Columns {
		tableData.Columns[i] = col.Name
	}

	// Generate records
	for i := 0; i < table.RecordCount; i++ {
		row, err := e.generateRow(table, i+1)
		if err != nil {
			return fmt.Errorf("failed to generate row %d: %w", i, err)
		}
		tableData.Rows = append(tableData.Rows, row)
	}

	// Store table data
	e.dataset.Tables[table.Name] = tableData

	return nil
}

// generateRow generates a single row of data
func (e *Engine) generateRow(table *schema.Table, rowID int) (map[string]interface{}, error) {
	row := make(map[string]interface{})

	for _, col := range table.Columns {
		value, err := e.generateColumnValue(table, col, rowID)
		if err != nil {
			return nil, fmt.Errorf("failed to generate column %s: %w", col.Name, err)
		}
		row[col.Name] = value
	}

	return row, nil
}

// generateColumnValue generates a value for a single column
func (e *Engine) generateColumnValue(table *schema.Table, col *schema.Column, rowID int) (interface{}, error) {
	// Handle primary key
	if col.PrimaryKey && col.AutoIncrement {
		// Track primary key for FK references
		e.pkTracker[table.Name] = append(e.pkTracker[table.Name], rowID)
		return rowID, nil
	}

	// Handle foreign key
	if col.ForeignKey != nil {
		return e.generateForeignKeyValue(col.ForeignKey)
	}

	// Handle generator-based values
	if col.Generator != "" {
		return e.invokeGenerator(col.Generator, col.GeneratorParams)
	}

	// Fallback: generate based on type
	return e.generateByType(col.Type)
}

// generateForeignKeyValue generates a foreign key value
func (e *Engine) generateForeignKeyValue(fk *schema.ForeignKey) (interface{}, error) {
	// Get available primary keys from referenced table
	pks, exists := e.pkTracker[fk.Table]
	if !exists || len(pks) == 0 {
		return nil, fmt.Errorf("no primary keys available for FK reference to %s", fk.Table)
	}

	// Randomly select one of the available primary keys
	idx := rand.Intn(len(pks))
	return pks[idx], nil
}

// invokeGenerator invokes the appropriate generator
func (e *Engine) invokeGenerator(generatorName string, params map[string]interface{}) (interface{}, error) {
	switch generatorName {
	// Base generators
	case "first_name":
		return FirstName(), nil
	case "last_name":
		return LastName(), nil
	case "full_name":
		return FullName(), nil
	case "email":
		return Email(), nil
	case "phone":
		return Phone(), nil
	case "company":
		return Company(), nil

	// Date/time generators
	case "timestamp_past":
		daysAgo := 365
		if val, ok := params["days_ago_max"].(float64); ok {
			daysAgo = int(val)
		}
		return TimestampPast(daysAgo), nil
	case "date_past":
		daysAgo := 365
		if val, ok := params["days_ago_max"].(float64); ok {
			daysAgo = int(val)
		}
		return TimestampPast(daysAgo).Format("2006-01-02"), nil

	// Fintech generators
	case "loan_amount":
		return LoanAmount(), nil
	case "credit_score":
		return CreditScore(), nil
	case "interest_rate":
		return InterestRate(), nil
	case "loan_status":
		return LoanStatus(), nil
	case "loan_number":
		return LoanNumber(), nil
	case "loan_term":
		return LoanTerm(), nil
	case "annual_income":
		return AnnualIncome(), nil
	case "employment_status":
		return EmploymentStatus(), nil
	case "payment_status":
		return PaymentStatus(), nil
	case "credit_bureau":
		return CreditBureau(), nil

	// Numeric generators
	case "int_range":
		min := 0
		max := 1000
		if val, ok := params["min"].(float64); ok {
			min = int(val)
		}
		if val, ok := params["max"].(float64); ok {
			max = int(val)
		}
		return IntRange(min, max), nil

	default:
		return nil, fmt.Errorf("unknown generator: %s", generatorName)
	}
}

// generateByType generates a value based on column type (fallback)
func (e *Engine) generateByType(colType string) (interface{}, error) {
	switch {
	case contains(colType, "int"):
		return IntRange(1, 1000), nil
	case contains(colType, "varchar") || contains(colType, "text"):
		return FullName(), nil
	case contains(colType, "decimal") || contains(colType, "float"):
		return FloatRange(0, 1000), nil
	case contains(colType, "date"):
		return TimestampPast(365).Format("2006-01-02"), nil
	case contains(colType, "timestamp"):
		return TimestampPast(365), nil
	default:
		return "DEFAULT", nil
	}
}

// findTable finds a table by name in the schema
func (e *Engine) findTable(name string) *schema.Table {
	for _, table := range e.schema.Tables {
		if table.Name == name {
			return &table
		}
	}
	return nil
}

// contains checks if string contains substring (case-insensitive)
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		   (s == substr ||
		    len(s) > len(substr) &&
			(s[:len(substr)] == substr ||
			 s[len(s)-len(substr):] == substr))
}
```

#### Unit Tests: `pkg/generators/engine_test.go`

```go
package generators

import (
	"testing"

	"github.com/yourusername/sourcebox/pkg/schema"
)

func TestEngineWithFintechSchema(t *testing.T) {
	// Load fintech schema
	s, err := schema.LoadSchema("../../schemas/fintech-loans.json")
	if err != nil {
		t.Fatalf("failed to load schema: %v", err)
	}

	// Create engine
	engine := NewEngine(s)

	// Generate data
	dataset, err := engine.Generate()
	if err != nil {
		t.Fatalf("generation failed: %v", err)
	}

	// Verify all tables generated
	if len(dataset.Tables) != 4 {
		t.Errorf("expected 4 tables, got %d", len(dataset.Tables))
	}

	// Verify record counts
	expectedCounts := map[string]int{
		"borrowers":     250,
		"loans":         1000,
		"payments":      3500,
		"credit_scores": 250,
	}

	for tableName, expectedCount := range expectedCounts {
		tableData, exists := dataset.Tables[tableName]
		if !exists {
			t.Errorf("table %s not generated", tableName)
			continue
		}

		if len(tableData.Rows) != expectedCount {
			t.Errorf("table %s: expected %d rows, got %d",
				tableName, expectedCount, len(tableData.Rows))
		}
	}
}

func TestReferentialIntegrity(t *testing.T) {
	s, err := schema.LoadSchema("../../schemas/fintech-loans.json")
	if err != nil {
		t.Fatalf("failed to load schema: %v", err)
	}

	engine := NewEngine(s)
	dataset, err := engine.Generate()
	if err != nil {
		t.Fatalf("generation failed: %v", err)
	}

	// Get borrower IDs
	borrowers := dataset.Tables["borrowers"]
	borrowerIDs := make(map[int]bool)
	for _, row := range borrowers.Rows {
		if id, ok := row["id"].(int); ok {
			borrowerIDs[id] = true
		}
	}

	// Verify all loans have valid borrower_id
	loans := dataset.Tables["loans"]
	for i, row := range loans.Rows {
		borrowerID, ok := row["borrower_id"].(int)
		if !ok {
			t.Errorf("loan row %d: borrower_id is not an int", i)
			continue
		}

		if !borrowerIDs[borrowerID] {
			t.Errorf("loan row %d: borrower_id %d does not exist in borrowers table",
				i, borrowerID)
		}
	}

	// Get loan IDs
	loanIDs := make(map[int]bool)
	for _, row := range loans.Rows {
		if id, ok := row["id"].(int); ok {
			loanIDs[id] = true
		}
	}

	// Verify all payments have valid loan_id
	payments := dataset.Tables["payments"]
	for i, row := range payments.Rows {
		loanID, ok := row["loan_id"].(int)
		if !ok {
			t.Errorf("payment row %d: loan_id is not an int", i)
			continue
		}

		if !loanIDs[loanID] {
			t.Errorf("payment row %d: loan_id %d does not exist in loans table",
				i, loanID)
		}
	}
}

func TestGenerationPerformance(t *testing.T) {
	s, err := schema.LoadSchema("../../schemas/fintech-loans.json")
	if err != nil {
		t.Fatalf("failed to load schema: %v", err)
	}

	engine := NewEngine(s)

	// Time the generation
	start := time.Now()
	_, err = engine.Generate()
	duration := time.Since(start)

	if err != nil {
		t.Fatalf("generation failed: %v", err)
	}

	// Should generate 4,950 records in < 5 seconds (conservative)
	// Target: < 30 seconds for database seeding, but generation alone should be faster
	if duration > 5*time.Second {
		t.Errorf("generation took too long: %v (expected < 5s)", duration)
	}

	// Calculate records/second
	recordsPerSec := 4950 / duration.Seconds()
	if recordsPerSec < 100 {
		t.Errorf("generation rate too slow: %.0f records/sec (expected > 100)", recordsPerSec)
	}
}

func BenchmarkDataGeneration(b *testing.B) {
	s, _ := schema.LoadSchema("../../schemas/fintech-loans.json")
	engine := NewEngine(s)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.Generate()
	}
}
```

### Performance Considerations
- Target: 100+ records/second generation
- Memory: Hold entire dataset in memory (acceptable for MVP)
- CPU-bound: No I/O during generation
- Benchmarks verify performance targets

### Testing Strategy
- Unit tests with fintech schema (F010)
- Verify all tables generated
- Verify record counts match schema
- Verify referential integrity (FK values exist)
- Performance tests (generation time, records/sec)
- Benchmark tests

## Dependencies
- **Upstream**:
  - F008 (Schema parser to load schemas)
  - F011 (Base generators)
  - F012 (Custom fintech generators)
- **Downstream**: F020 (Integration tests use engine)

## Deliverables
1. Data generation engine module (`pkg/generators/engine.go`)
2. Dataset structures (Dataset, TableData)
3. Comprehensive unit tests (`pkg/generators/engine_test.go`)
4. Performance benchmarks
5. Test coverage report (>90%)

## Success Criteria
- ✅ Engine generates complete datasets from schemas
- ✅ Referential integrity maintained (FK tests pass)
- ✅ Tests pass with fintech schema
- ✅ Performance: >100 records/sec
- ✅ Test coverage >90%

## Anti-Patterns to Avoid
- ❌ Generating children before parents (FK violations)
- ❌ Ignoring foreign key constraints
- ❌ Slow generation (< 100 records/sec)
- ❌ Poor error messages
- ❌ Insufficient tests

## Implementation Notes
- Engine is stateful (tracks PKs for FK references)
- Generator invocation uses switch/case for now (can be extended)
- PK tracking is simple (store all generated PKs)
- FK selection is random from available PKs

## TDD Requirements
**REQUIRED** - Engine is core functionality:

1. Write test: Generate fintech schema, verify tables exist
2. Run test - should fail
3. Implement: Basic engine structure
4. Run test - should pass
5. Write test: Verify referential integrity
6. Run test - should fail
7. Implement: FK tracking and generation
8. Run test - should pass
9. Repeat for performance, etc.

## Related Constitution Sections
- **Speed > Features (Principle II, NON-NEGOTIABLE)**: < 30s seeding time
- **Performance (Technical Constraint 1)**: 100+ records/sec
- **TDD Required (Development Practice 1)**: Core generation logic
- **Code Quality (Technical Constraint 5)**: >80% coverage
