# F008 Quickstart: Schema Parser & Validator Development

**Feature**: Schema Parser & Validator
**Branch**: `006-f008-schema-parser`
**Date**: 2025-10-15
**Approach**: Test-Driven Development (TDD)

## Overview

This quickstart guide walks through developing the Schema Parser & Validator using strict TDD discipline. The parser loads JSON schema files (F007 format), validates all constraints, and returns clear error messages.

**Constitutional Requirement**: TDD is NON-NEGOTIABLE for core functionality (Development Practice 1). Target: 100% coverage for `pkg/schema/parser.go`.

---

## Prerequisites

### Required

- ✅ Go 1.21+ installed (`go version`)
- ✅ F007 schema specification complete (`schemas/schema-spec.md`)
- ✅ F007 example schema available for testing (`schemas/example-schema.json`)
- ✅ testify library available (already in project from F006)
- ✅ Git repository initialized (existing from F003)

### Knowledge

- Go basics (structs, pointers, error handling)
- JSON parsing with `encoding/json`
- Table-driven tests
- Test-first development mindset

---

## Development Philosophy

### TDD Workflow (Strict Discipline)

**RED → GREEN → REFACTOR**:

1. **RED**: Write a failing test for one specific behavior
2. **GREEN**: Write minimal code to make that test pass
3. **REFACTOR**: Clean up code while keeping tests green
4. **REPEAT**: Move to next behavior

**Test-First Discipline**:
- ❌ Don't write implementation before test
- ❌ Don't write multiple tests before implementing
- ❌ Don't skip tests for "simple" code
- ✅ Write one test at a time
- ✅ Implement only what's needed to pass
- ✅ Refactor with test safety net

**Coverage Target**: 100% for parser.go (critical path validation)

---

## Phase 1: Setup Test Structure

### Step 1.1: Create Test File First

```bash
# Create package directory
mkdir -p pkg/schema

# Create test file FIRST (TDD discipline)
touch pkg/schema/parser_test.go
```

### Step 1.2: Write First Test (Valid Schema)

**File**: `pkg/schema/parser_test.go`

```go
package schema

import (
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestParseValidMinimalSchema(t *testing.T) {
    schemaJSON := `{
        "schema_version": "1.0",
        "name": "test-schema",
        "description": "Minimal valid schema for testing",
        "author": "Test Author",
        "version": "1.0.0",
        "database_type": ["mysql"],
        "metadata": {
            "total_records": 100,
            "industry": "test",
            "tags": ["test"]
        },
        "tables": [{
            "name": "users",
            "description": "User table",
            "record_count": 100,
            "columns": [{
                "name": "id",
                "type": "int",
                "primary_key": true
            }]
        }],
        "relationships": [],
        "generation_order": ["users"]
    }`

    schema, err := ParseSchema(strings.NewReader(schemaJSON))

    require.NoError(t, err, "Valid schema should parse without error")
    assert.NotNil(t, schema, "Schema should not be nil")
    assert.Equal(t, "test-schema", schema.Name)
    assert.Equal(t, "1.0", schema.SchemaVersion)
    assert.Len(t, schema.Tables, 1)
    assert.Equal(t, "users", schema.Tables[0].Name)
}
```

### Step 1.3: Run Test (Should Fail - RED)

```bash
cd pkg/schema
go test -v

# Expected output:
# parser_test.go:10:18: undefined: ParseSchema
# FAIL    github.com/jbeausoleil/sourcebox/pkg/schema [build failed]
```

**Status**: ❌ RED (test fails - function doesn't exist yet)

---

## Phase 2: Implement Basic Parser

### Step 2.1: Create Types File

**File**: `pkg/schema/types.go`

```go
package schema

// Schema represents a complete database schema definition
type Schema struct {
    SchemaVersion   string           `json:"schema_version"`
    Name            string           `json:"name"`
    Description     string           `json:"description"`
    Author          string           `json:"author"`
    Version         string           `json:"version"`
    DatabaseType    []string         `json:"database_type"`
    Metadata        SchemaMetadata   `json:"metadata"`
    Tables          []Table          `json:"tables"`
    Relationships   []Relationship   `json:"relationships"`
    GenerationOrder []string         `json:"generation_order"`
    ValidationRules []ValidationRule `json:"validation_rules,omitempty"`
}

// SchemaMetadata contains additional schema information
type SchemaMetadata struct {
    TotalRecords int      `json:"total_records"`
    Industry     string   `json:"industry"`
    Tags         []string `json:"tags"`
}

// Table represents a database table definition
type Table struct {
    Name        string   `json:"name"`
    Description string   `json:"description"`
    RecordCount int      `json:"record_count"`
    Columns     []Column `json:"columns"`
    Indexes     []Index  `json:"indexes,omitempty"`
}

// Column represents a table column definition
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
    OnDelete string `json:"on_delete,omitempty"`
    OnUpdate string `json:"on_update,omitempty"`
}

// Index represents a table index
type Index struct {
    Name    string   `json:"name"`
    Columns []string `json:"columns"`
    Unique  bool     `json:"unique,omitempty"`
}

// Relationship represents an explicit relationship definition
type Relationship struct {
    FromTable        string `json:"from_table"`
    FromColumn       string `json:"from_column"`
    ToTable          string `json:"to_table"`
    ToColumn         string `json:"to_column"`
    RelationshipType string `json:"relationship_type"`
    Description      string `json:"description"`
}

// ValidationRule represents a data validation rule
type ValidationRule struct {
    Rule        string `json:"rule"`
    Description string `json:"description"`
    Enforcement string `json:"enforcement"`
}
```

### Step 2.2: Create Parser File

**File**: `pkg/schema/parser.go`

```go
package schema

import (
    "encoding/json"
    "fmt"
    "io"
)

// ParseSchema parses a schema from an io.Reader
func ParseSchema(r io.Reader) (*Schema, error) {
    decoder := json.NewDecoder(r)
    decoder.DisallowUnknownFields() // Strict parsing

    var schema Schema
    if err := decoder.Decode(&schema); err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %w", err)
    }

    // Validation will be added in Phase 3
    return &schema, nil
}
```

### Step 2.3: Run Test (Should Pass - GREEN)

```bash
go test -v

# Expected output:
# === RUN   TestParseValidMinimalSchema
# --- PASS: TestParseValidMinimalSchema (0.00s)
# PASS
# ok      github.com/jbeausoleil/sourcebox/pkg/schema    0.003s
```

**Status**: ✅ GREEN (test passes)

### Step 2.4: Add LoadSchema Convenience Function

**File**: `pkg/schema/parser.go` (append)

```go
import (
    "os"
    // ... other imports
)

// LoadSchema loads and parses a schema from a file path
func LoadSchema(path string) (*Schema, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("failed to open schema file: %w", err)
    }
    defer file.Close()

    return ParseSchema(file)
}
```

---

## Phase 3: Add Validation (TDD Cycle)

### Validation Order

Implement validations in this order (each with test-first):

1. ✅ Schema name required
2. ✅ At least one table required
3. ✅ Database type required
4. ✅ Table name required
5. ✅ At least one column per table
6. ✅ Column name and type required
7. ✅ Exactly one primary key per table
8. ✅ Positive record count
9. ✅ Valid data types
10. ✅ Foreign keys reference existing tables
11. ✅ Valid referential integrity actions
12. ✅ Generation order includes all tables
13. ✅ No duplicate tables in generation order
14. ✅ No missing tables in generation order

### Example TDD Cycle: Schema Name Required

#### Step 3.1: Write Failing Test (RED)

**File**: `pkg/schema/parser_test.go` (append)

```go
func TestParseMissingSchemaName(t *testing.T) {
    schemaJSON := `{
        "schema_version": "1.0",
        "description": "Schema without name",
        "author": "Test",
        "version": "1.0.0",
        "database_type": ["mysql"],
        "metadata": {"total_records": 0, "industry": "test", "tags": []},
        "tables": [],
        "relationships": [],
        "generation_order": []
    }`

    schema, err := ParseSchema(strings.NewReader(schemaJSON))

    require.Error(t, err, "Should error on missing schema name")
    assert.Nil(t, schema, "Schema should be nil on error")
    assert.Contains(t, err.Error(), "name is required",
        "Error message should indicate missing name")
}
```

#### Step 3.2: Run Test (Should Fail - RED)

```bash
go test -v -run TestParseMissingSchemaName

# Expected output:
# --- FAIL: TestParseMissingSchemaName (0.00s)
#     parser_test.go:45: Should error on missing schema name
# FAIL
```

**Status**: ❌ RED (test fails - no validation yet)

#### Step 3.3: Implement Validation (GREEN)

**File**: `pkg/schema/parser.go` (update ParseSchema)

```go
func ParseSchema(r io.Reader) (*Schema, error) {
    decoder := json.NewDecoder(r)
    decoder.DisallowUnknownFields()

    var schema Schema
    if err := decoder.Decode(&schema); err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %w", err)
    }

    // Add validation
    if err := ValidateSchema(&schema); err != nil {
        return nil, fmt.Errorf("schema validation failed: %w", err)
    }

    return &schema, nil
}

// ValidateSchema validates a parsed schema
func ValidateSchema(s *Schema) error {
    // Validate schema name
    if s.Name == "" {
        return fmt.Errorf("schema name is required")
    }

    // More validations will be added here
    return nil
}
```

#### Step 3.4: Run Test (Should Pass - GREEN)

```bash
go test -v -run TestParseMissingSchemaName

# Expected output:
# === RUN   TestParseMissingSchemaName
# --- PASS: TestParseMissingSchemaName (0.00s)
# PASS
```

**Status**: ✅ GREEN (test passes)

#### Step 3.5: Run All Tests (Ensure No Regression)

```bash
go test -v

# Expected output:
# === RUN   TestParseValidMinimalSchema
# --- PASS: TestParseValidMinimalSchema (0.00s)
# === RUN   TestParseMissingSchemaName
# --- PASS: TestParseMissingSchemaName (0.00s)
# PASS
```

**Status**: ✅ GREEN (all tests pass - no regression)

### Repeat TDD Cycle for Each Validation

Continue this pattern for all validation rules:
1. Write failing test (RED)
2. Implement minimal validation (GREEN)
3. Refactor if needed (keep tests GREEN)
4. Move to next validation

---

## Phase 4: Edge Cases

### Test Edge Cases

Add tests for these scenarios:

```go
func TestParseEmptyFile(t *testing.T) {
    schema, err := ParseSchema(strings.NewReader(""))
    require.Error(t, err)
    assert.Nil(t, schema)
    assert.Contains(t, err.Error(), "parse")
}

func TestParseMalformedJSON(t *testing.T) {
    schema, err := ParseSchema(strings.NewReader("{invalid json"))
    require.Error(t, err)
    assert.Nil(t, schema)
}

func TestParseUnknownFields(t *testing.T) {
    schemaJSON := `{
        "schema_version": "1.0",
        "name": "test",
        "unknown_field": "value",
        "tables": []
    }`

    schema, err := ParseSchema(strings.NewReader(schemaJSON))
    require.Error(t, err)
    assert.Nil(t, schema)
    assert.Contains(t, err.Error(), "unknown")
}
```

---

## Running Tests

### Run All Tests

```bash
cd pkg/schema
go test -v
```

### Run Specific Test

```bash
go test -v -run TestParseValidMinimalSchema
```

### Run with Coverage

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Check Coverage Percentage

```bash
go test -cover

# Target output:
# PASS
# coverage: 100.0% of statements
# ok      github.com/jbeausoleil/sourcebox/pkg/schema    0.005s
```

### Run with Race Detection

```bash
go test -race
```

---

## Verification Checklist

### Implementation Complete

- [ ] `pkg/schema/types.go` defines all structs
- [ ] `pkg/schema/parser.go` implements `ParseSchema` and `LoadSchema`
- [ ] `pkg/schema/parser.go` implements `ValidateSchema` and validation functions
- [ ] JSON parsing uses `DisallowUnknownFields()`
- [ ] Error messages include context (table name, column name, index)

### Tests Complete

- [ ] Test for valid schema parsing passes
- [ ] Tests for all required fields pass
- [ ] Tests for invalid data types pass
- [ ] Tests for foreign key validation pass
- [ ] Tests for generation order validation pass
- [ ] Tests for duplicate names pass
- [ ] Tests for edge cases pass (empty file, malformed JSON, unknown fields)
- [ ] All tests use table-driven test pattern where appropriate
- [ ] Test coverage is 100% for `parser.go`

### Validation Complete

- [ ] F007 example schema parses successfully
- [ ] Invalid schemas return clear error messages with context
- [ ] Parser handles both file and `io.Reader` input
- [ ] Performance is <100ms for typical schemas
- [ ] `go vet ./pkg/schema/...` passes with zero warnings
- [ ] `go test ./pkg/schema/...` passes all tests

### Integration Ready

- [ ] Parser can be imported: `import "github.com/jbeausoleil/sourcebox/pkg/schema"`
- [ ] LoadSchema works with absolute file paths
- [ ] ParseSchema works with `strings.NewReader`, `os.File`, `embed.FS`
- [ ] Ready for F013 data generation engine integration

---

## Common Issues and Fixes

### Issue: Test fails with "unknown field" error

**Cause**: JSON has field not in Go struct
**Fix**: Add field to struct in `types.go` with correct JSON tag

```go
// Add missing field
type Schema struct {
    // ... existing fields
    MissingField string `json:"missing_field"`
}
```

### Issue: Test fails with "name is required" but name is present

**Cause**: JSON field name doesn't match struct tag
**Fix**: Ensure JSON uses snake_case to match F007 format

```go
// Correct JSON
{"schema_version": "1.0", "name": "test"}

// Incorrect JSON
{"schemaVersion": "1.0", "name": "test"} // Won't parse
```

### Issue: Foreign key validation not working

**Cause**: Table names map not populated before validation
**Fix**: Build `tableNames` map before validating foreign keys

```go
func ValidateSchema(s *Schema) error {
    // Build table names map first
    tableNames := make(map[string]bool, len(s.Tables))
    for _, table := range s.Tables {
        tableNames[table.Name] = true
    }

    // Then validate foreign keys
    for _, table := range s.Tables {
        for _, col := range table.Columns {
            if col.ForeignKey != nil {
                if !tableNames[col.ForeignKey.Table] {
                    return fmt.Errorf("foreign key references non-existent table '%s'",
                        col.ForeignKey.Table)
                }
            }
        }
    }

    return nil
}
```

### Issue: Coverage is not 100%

**Cause**: Missing test for validation path
**Fix**: Check coverage report, add test for uncovered error path

```bash
go test -coverprofile=coverage.out
go tool cover -func=coverage.out | grep parser.go

# Look for lines with 0% coverage
# Add tests for those paths
```

### Issue: Error messages are unclear

**Cause**: Missing context in error wrapping
**Fix**: Use `fmt.Errorf("context: %w", err)` pattern

```go
// Bad (no context)
return err

// Good (with context)
return fmt.Errorf("table %d (%s): column %d (%s): %w",
    tableIdx, table.Name, colIdx, col.Name, err)
```

### Issue: Performance is slow (>100ms)

**Cause**: Inefficient validation (nested loops, string concatenation)
**Fix**: Use maps for O(1) lookup, avoid unnecessary allocations

```go
// Slow: O(n²) lookup
for _, table := range schema.Tables {
    for _, refTable := range schema.Tables {
        if table.Name == refTable.Name { ... }
    }
}

// Fast: O(1) lookup with map
tableNames := make(map[string]bool, len(schema.Tables))
for _, table := range schema.Tables {
    tableNames[table.Name] = true
}
// Now: if tableNames[name] { ... }
```

---

## Manual Verification Steps

### 1. Verify Valid Schema Parsing

```bash
# Create test schema file
cat > /tmp/test-schema.json <<EOF
{
  "schema_version": "1.0",
  "name": "test",
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
      "primary_key": true
    }]
  }],
  "relationships": [],
  "generation_order": ["users"]
}
EOF

# Test with Go code
go run -C /tmp <<EOF
package main
import (
    "fmt"
    "log"
    "github.com/jbeausoleil/sourcebox/pkg/schema"
)
func main() {
    s, err := schema.LoadSchema("/tmp/test-schema.json")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Loaded schema: %s (v%s)\n", s.Name, s.Version)
}
EOF
```

Expected: `Loaded schema: test (v1.0.0)`

### 2. Verify Error Messages

```bash
# Create invalid schema (missing name)
cat > /tmp/invalid-schema.json <<EOF
{
  "schema_version": "1.0",
  "description": "Missing name field",
  "tables": []
}
EOF

# Test error message
go run -C /tmp <<EOF
package main
import (
    "fmt"
    "github.com/jbeausoleil/sourcebox/pkg/schema"
)
func main() {
    _, err := schema.LoadSchema("/tmp/invalid-schema.json")
    if err != nil {
        fmt.Printf("Error (expected): %v\n", err)
    }
}
EOF
```

Expected: Error message contains "name is required"

### 3. Verify Performance

```bash
# Run benchmark
go test -bench=BenchmarkParseSchema -benchmem

# Manual timing
time go run -C /tmp <<EOF
package main
import "github.com/jbeausoleil/sourcebox/pkg/schema"
func main() {
    for i := 0; i < 100; i++ {
        schema.LoadSchema("/tmp/test-schema.json")
    }
}
EOF
```

Expected: <100ms for 100 iterations (<1ms per parse)

---

## Testing Best Practices

### Test Organization

```go
// Group related tests using subtests
func TestParseSchema(t *testing.T) {
    t.Run("valid minimal schema", func(t *testing.T) { /* ... */ })
    t.Run("valid full schema", func(t *testing.T) { /* ... */ })
}

func TestValidation(t *testing.T) {
    t.Run("missing name", func(t *testing.T) { /* ... */ })
    t.Run("missing tables", func(t *testing.T) { /* ... */ })
}
```

### Table-Driven Tests

```go
func TestValidateDataType(t *testing.T) {
    tests := []struct {
        name        string
        dataType    string
        expectError bool
    }{
        {"valid int", "int", false},
        {"valid varchar", "varchar(255)", false},
        {"valid decimal", "decimal(10,2)", false},
        {"invalid type", "foobar", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateDataType(tt.dataType)
            if tt.expectError {
                require.Error(t, err)
            } else {
                require.NoError(t, err)
            }
        })
    }
}
```

### Error Assertions

```go
// Use testify for clear assertions
require.Error(t, err, "Should return error for invalid input")
assert.Contains(t, err.Error(), "expected substring",
    "Error message should contain specific text")
assert.Nil(t, schema, "Schema should be nil on error")
assert.NotNil(t, schema, "Schema should not be nil on success")
```

---

## Next Steps

After F008 parser is complete and tested:

1. **F013**: Data Generation Engine
   - Uses parser to load schemas
   - Implements generator logic
   - Generates realistic data

2. **F010**: Fintech Schema
   - Validated by parser
   - First verticalized schema

3. **F016**: Healthcare Schema
   - Validated by parser
   - Second verticalized schema

4. **F018**: Retail Schema
   - Validated by parser
   - Third verticalized schema

---

## Constitutional Compliance

✅ **TDD Required**: All core parser logic developed test-first
✅ **>80% Coverage**: Target 100% for parser.go (critical path)
✅ **Developer-First**: Clear error messages with context
✅ **Speed > Features**: <100ms parsing (linear-time validation)
✅ **Boring Tech Wins**: Standard library only (encoding/json)
✅ **Ship Fast**: 3-day implementation (TDD discipline prevents scope creep)

---

## Summary

**TDD Workflow**:
1. Write failing test (RED)
2. Implement minimal code (GREEN)
3. Refactor if needed (keep tests GREEN)
4. Repeat for each validation rule

**Coverage Target**: 100% for `parser.go`

**Success Criteria**:
- All tests pass (`go test ./pkg/schema/...`)
- Coverage is 100% (`go test -cover`)
- F007 example schema parses successfully
- Invalid schemas return clear, actionable errors
- Performance is <100ms for typical schemas

**Ready for Integration**: Parser can be imported by F013 data generation engine
