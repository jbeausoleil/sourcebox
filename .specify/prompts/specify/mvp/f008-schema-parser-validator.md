# Feature Specification Prompt: F008 - Schema Parser & Validator

## Feature Metadata
- **Feature ID**: F008
- **Name**: Schema Parser & Validator
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F007 (Schema specification must be defined)

## Constitutional Alignment

### Core Principles
- ✅ **Developer-First Design**: Clear error messages for invalid schemas
- ✅ **Quality Standards**: Validation prevents runtime errors
- ✅ **Fail Gracefully**: Actionable validation errors

### Technical Constraints
- ✅ **Code Quality**: > 80% test coverage (TDD required)
- ✅ **Performance**: Fast schema loading (<100ms)

### Development Practices
- ✅ **TDD Required for Core Functionality**: Schema parser is core logic
- ✅ **Test-first workflow**: Write test → implement → refactor

## User Story
**US-MVP-004**: "As a developer, I want schema files to be validated when loaded so I get clear error messages if something is wrong, rather than cryptic runtime failures."

## Problem Statement
SourceBox needs to load and validate schema JSON files to ensure they conform to the specification (F007). Invalid schemas cause runtime errors that are hard to debug. A robust parser must:
- Parse JSON into Go structs
- Validate all required fields exist
- Check data types are valid for target database
- Verify foreign key relationships are correct
- Provide clear, actionable error messages
- Support both file-based and embedded schemas

Without proper validation, developers waste time debugging malformed schemas instead of building features.

## Solution Overview
Build a schema parser in `/pkg/schema/parser.go` that unmarshals JSON into Go structs, validates all constraints, and returns clear errors. Use Go's standard `encoding/json` for parsing. Implement comprehensive validation rules: required fields, valid types, foreign key integrity, table ordering. Provide 100% test coverage using TDD methodology.

## Detailed Requirements

### Acceptance Criteria
1. **Parser Implemented**: `pkg/schema/parser.go` with `LoadSchema` function
2. **Go Structs Defined**: Schema data structures in `pkg/schema/types.go`
3. **JSON Unmarshaling**: Properly decodes JSON into structs
4. **Validation Rules Implemented**:
   - Required fields present (name, tables, columns)
   - Valid data types for target database
   - Foreign key integrity (referenced tables/columns exist)
   - Primary key exists in each table
   - No duplicate table/column names
   - `generation_order` includes all tables
   - Generator names are valid
5. **Clear Error Messages**: Errors indicate exactly what's wrong and where
6. **File and Embedded Support**: Load from file path or `embed.FS`
7. **Unit Tests**: `go test ./pkg/schema/... -v` passes with 100% coverage

### Technical Specifications

#### Go Data Structures: `pkg/schema/types.go`

```go
package schema

import "time"

// Schema represents a complete data schema
type Schema struct {
	SchemaVersion string            `json:"schema_version"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Author        string            `json:"author"`
	Version       string            `json:"version"`
	DatabaseType  []string          `json:"database_type"` // ["mysql", "postgres"]
	Metadata      SchemaMetadata    `json:"metadata"`
	Tables        []Table           `json:"tables"`
	Relationships []Relationship    `json:"relationships"`
	GenerationOrder []string        `json:"generation_order"`
	ValidationRules []ValidationRule `json:"validation_rules,omitempty"`
}

// SchemaMetadata contains schema metadata
type SchemaMetadata struct {
	TotalRecords int      `json:"total_records"`
	Industry     string   `json:"industry"`
	Tags         []string `json:"tags"`
}

// Table represents a database table
type Table struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	RecordCount int       `json:"record_count"`
	Columns     []Column  `json:"columns"`
	Indexes     []Index   `json:"indexes,omitempty"`
}

// Column represents a table column
type Column struct {
	Name            string                 `json:"name"`
	Type            string                 `json:"type"`
	PrimaryKey      bool                   `json:"primary_key,omitempty"`
	AutoIncrement   bool                   `json:"auto_increment,omitempty"`
	Nullable        bool                   `json:"nullable,omitempty"`
	Unique          bool                   `json:"unique,omitempty"`
	Default         string                 `json:"default,omitempty"`
	Generator       string                 `json:"generator,omitempty"`
	GeneratorParams map[string]interface{} `json:"generator_params,omitempty"`
	ForeignKey      *ForeignKey            `json:"foreign_key,omitempty"`
}

// ForeignKey represents a foreign key constraint
type ForeignKey struct {
	Table    string `json:"table"`
	Column   string `json:"column"`
	OnDelete string `json:"on_delete,omitempty"` // CASCADE, SET NULL, RESTRICT
	OnUpdate string `json:"on_update,omitempty"` // CASCADE, SET NULL, RESTRICT
}

// Index represents a table index
type Index struct {
	Name    string   `json:"name"`
	Columns []string `json:"columns"`
	Unique  bool     `json:"unique,omitempty"`
}

// Relationship documents table relationships
type Relationship struct {
	FromTable        string `json:"from_table"`
	FromColumn       string `json:"from_column"`
	ToTable          string `json:"to_table"`
	ToColumn         string `json:"to_column"`
	RelationshipType string `json:"relationship_type"` // one_to_one, one_to_many, many_to_one, many_to_many
	Description      string `json:"description"`
}

// ValidationRule represents a data validation rule
type ValidationRule struct {
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Enforcement string `json:"enforcement"` // soft, hard
}
```

#### Parser Implementation: `pkg/schema/parser.go`

```go
package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// LoadSchema loads and validates a schema from a file
func LoadSchema(path string) (*Schema, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open schema file %s: %w", path, err)
	}
	defer file.Close()

	return ParseSchema(file)
}

// ParseSchema parses and validates a schema from a reader
func ParseSchema(r io.Reader) (*Schema, error) {
	var schema Schema

	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields() // Strict parsing

	if err := decoder.Decode(&schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
	}

	// Validate schema
	if err := ValidateSchema(&schema); err != nil {
		return nil, fmt.Errorf("schema validation failed: %w", err)
	}

	return &schema, nil
}

// ValidateSchema validates a schema
func ValidateSchema(s *Schema) error {
	// Required fields
	if s.Name == "" {
		return fmt.Errorf("schema name is required")
	}
	if len(s.Tables) == 0 {
		return fmt.Errorf("schema must have at least one table")
	}
	if len(s.DatabaseType) == 0 {
		return fmt.Errorf("database_type is required")
	}

	// Validate database types
	validDBTypes := map[string]bool{"mysql": true, "postgres": true}
	for _, dbType := range s.DatabaseType {
		if !validDBTypes[dbType] {
			return fmt.Errorf("invalid database_type: %s (must be 'mysql' or 'postgres')", dbType)
		}
	}

	// Validate tables
	tableNames := make(map[string]bool)
	for i, table := range s.Tables {
		if err := ValidateTable(&table, tableNames); err != nil {
			return fmt.Errorf("table %d (%s): %w", i, table.Name, err)
		}
		tableNames[table.Name] = true
	}

	// Validate foreign keys reference existing tables
	for _, table := range s.Tables {
		for _, col := range table.Columns {
			if col.ForeignKey != nil {
				if !tableNames[col.ForeignKey.Table] {
					return fmt.Errorf("table %s: column %s references non-existent table %s",
						table.Name, col.Name, col.ForeignKey.Table)
				}
			}
		}
	}

	// Validate generation order
	if err := ValidateGenerationOrder(s, tableNames); err != nil {
		return err
	}

	return nil
}

// ValidateTable validates a single table
func ValidateTable(t *Table, existingTables map[string]bool) error {
	if t.Name == "" {
		return fmt.Errorf("table name is required")
	}
	if existingTables[t.Name] {
		return fmt.Errorf("duplicate table name: %s", t.Name)
	}
	if len(t.Columns) == 0 {
		return fmt.Errorf("table must have at least one column")
	}
	if t.RecordCount <= 0 {
		return fmt.Errorf("record_count must be positive, got %d", t.RecordCount)
	}

	// Validate at least one primary key
	hasPrimaryKey := false
	columnNames := make(map[string]bool)

	for i, col := range t.Columns {
		if err := ValidateColumn(&col, columnNames); err != nil {
			return fmt.Errorf("column %d (%s): %w", i, col.Name, err)
		}
		columnNames[col.Name] = true

		if col.PrimaryKey {
			hasPrimaryKey = true
		}
	}

	if !hasPrimaryKey {
		return fmt.Errorf("table must have a primary key")
	}

	return nil
}

// ValidateColumn validates a single column
func ValidateColumn(c *Column, existingColumns map[string]bool) error {
	if c.Name == "" {
		return fmt.Errorf("column name is required")
	}
	if existingColumns[c.Name] {
		return fmt.Errorf("duplicate column name: %s", c.Name)
	}
	if c.Type == "" {
		return fmt.Errorf("column type is required")
	}

	// Validate data type
	if err := ValidateDataType(c.Type); err != nil {
		return err
	}

	// Validate foreign key actions
	if c.ForeignKey != nil {
		validActions := map[string]bool{"CASCADE": true, "SET NULL": true, "RESTRICT": true, "": true}
		if !validActions[c.ForeignKey.OnDelete] {
			return fmt.Errorf("invalid on_delete action: %s", c.ForeignKey.OnDelete)
		}
		if !validActions[c.ForeignKey.OnUpdate] {
			return fmt.Errorf("invalid on_update action: %s", c.ForeignKey.OnUpdate)
		}
	}

	return nil
}

// ValidateDataType checks if a data type is valid
func ValidateDataType(dataType string) error {
	// Common types
	validTypes := []string{
		"int", "bigint", "smallint", "tinyint",
		"decimal", "float", "double",
		"varchar", "text", "char",
		"date", "datetime", "timestamp",
		"boolean", "bit", "json", "jsonb", "enum",
	}

	typeLower := strings.ToLower(dataType)
	for _, valid := range validTypes {
		if strings.HasPrefix(typeLower, valid) {
			return nil
		}
	}

	return fmt.Errorf("unsupported data type: %s", dataType)
}

// ValidateGenerationOrder ensures all tables are in generation order
func ValidateGenerationOrder(s *Schema, tableNames map[string]bool) error {
	if len(s.GenerationOrder) == 0 {
		return fmt.Errorf("generation_order is required")
	}

	orderSet := make(map[string]bool)
	for _, tableName := range s.GenerationOrder {
		if !tableNames[tableName] {
			return fmt.Errorf("generation_order references non-existent table: %s", tableName)
		}
		if orderSet[tableName] {
			return fmt.Errorf("generation_order contains duplicate table: %s", tableName)
		}
		orderSet[tableName] = true
	}

	// Ensure all tables are in generation order
	for tableName := range tableNames {
		if !orderSet[tableName] {
			return fmt.Errorf("table %s is not in generation_order", tableName)
		}
	}

	return nil
}
```

### Validation Rules Summary
1. **Schema level**: name, tables, database_type required
2. **Table level**: name, columns, record_count > 0, has primary key
3. **Column level**: name, type required; valid data type
4. **Foreign keys**: reference existing tables and columns
5. **Generation order**: includes all tables, no duplicates
6. **No duplicates**: table names, column names, index names

### Error Message Quality
❌ **Bad**: "validation error"
✅ **Good**: "table 2 (loans): column 5 (borrower_id): foreign key references non-existent table 'users'"

### Performance Considerations
- JSON parsing: O(n) where n = schema size
- Validation: O(n) where n = number of tables + columns
- Target: < 100ms for typical schemas (< 10 tables)

### Testing Strategy (TDD Required)

#### Test File: `pkg/schema/parser_test.go`

```go
package schema

import (
	"strings"
	"testing"
)

// Test valid schema parses successfully
func TestParseValidSchema(t *testing.T) {
	schemaJSON := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"metadata": {"total_records": 100, "industry": "test", "tags": ["test"]},
		"tables": [{
			"name": "users",
			"description": "Users table",
			"record_count": 100,
			"columns": [{
				"name": "id",
				"type": "int",
				"primary_key": true,
				"auto_increment": true
			}]
		}],
		"generation_order": ["users"]
	}`

	schema, err := ParseSchema(strings.NewReader(schemaJSON))
	if err != nil {
		t.Fatalf("expected valid schema to parse, got error: %v", err)
	}

	if schema.Name != "test-schema" {
		t.Errorf("expected name 'test-schema', got '%s'", schema.Name)
	}

	if len(schema.Tables) != 1 {
		t.Errorf("expected 1 table, got %d", len(schema.Tables))
	}
}

// Test missing required field
func TestParseMissingName(t *testing.T) {
	schemaJSON := `{
		"schema_version": "1.0",
		"database_type": ["mysql"],
		"tables": []
	}`

	_, err := ParseSchema(strings.NewReader(schemaJSON))
	if err == nil {
		t.Fatal("expected error for missing name, got nil")
	}

	if !strings.Contains(err.Error(), "name is required") {
		t.Errorf("expected error about missing name, got: %v", err)
	}
}

// Test invalid foreign key
func TestParseInvalidForeignKey(t *testing.T) {
	schemaJSON := `{
		"schema_version": "1.0",
		"name": "test",
		"database_type": ["mysql"],
		"tables": [{
			"name": "orders",
			"record_count": 100,
			"columns": [{
				"name": "id",
				"type": "int",
				"primary_key": true
			}, {
				"name": "user_id",
				"type": "int",
				"foreign_key": {
					"table": "users",
					"column": "id"
				}
			}]
		}],
		"generation_order": ["orders"]
	}`

	_, err := ParseSchema(strings.NewReader(schemaJSON))
	if err == nil {
		t.Fatal("expected error for non-existent foreign key table")
	}

	if !strings.Contains(err.Error(), "non-existent table") {
		t.Errorf("expected error about non-existent table, got: %v", err)
	}
}

// Additional tests: duplicate tables, missing primary key, invalid data type, etc.
```

## Dependencies
- **Upstream**: F007 (Schema specification must be defined)
- **Downstream**:
  - F010 (Fintech Schema) will be validated by this parser
  - F013 (Data Generation Engine) will use parsed schemas
  - F016, F018 (Healthcare, Retail Schemas) will be validated

## Deliverables
1. Go structs for schema representation (`pkg/schema/types.go`)
2. Schema parser implementation (`pkg/schema/parser.go`)
3. Comprehensive validation logic
4. Unit tests with 100% coverage (`pkg/schema/parser_test.go`)
5. Test coverage report

## Success Criteria
- ✅ Valid schemas parse successfully
- ✅ Invalid schemas return clear errors
- ✅ All validation rules implemented
- ✅ Tests pass: `go test ./pkg/schema/... -v`
- ✅ 100% test coverage: `go test -coverprofile=coverage.txt ./pkg/schema/...`
- ✅ Parser handles edge cases (empty files, malformed JSON, etc.)

## Anti-Patterns to Avoid
- ❌ Vague error messages ("schema error")
- ❌ Allowing invalid schemas to parse (strict validation required)
- ❌ Poor test coverage (< 100% for core parser)
- ❌ Ignoring edge cases (empty arrays, null values, etc.)
- ❌ Panic on errors (return errors, don't crash)

## Implementation Notes
- Use Go's `encoding/json` (standard library, well-tested)
- `DisallowUnknownFields()` catches typos in schema files
- Validation happens during parsing, not later
- Clear error messages save hours of debugging

## TDD Requirements
**REQUIRED for this feature** - Follow strict TDD workflow:

1. **Write test first**: Test for valid schema parsing
2. **Run test**: Should fail (parser not implemented)
3. **Implement**: Basic parser structure
4. **Run test**: Should pass
5. **Refactor**: Clean up code
6. **Repeat**: For each validation rule (missing name, invalid FK, etc.)

**Coverage target**: 100% for all parser code

## Related Constitution Sections
- **TDD Required for Core Functionality (Development Practice 1)**: Schema parser is core logic
- **Code Quality Standards (Technical Constraint 5)**: >80% coverage required
- **Fail Gracefully (UX Principle 4)**: Clear, actionable error messages
- **Developer-First Design (Principle VI)**: Good errors enable self-service debugging
