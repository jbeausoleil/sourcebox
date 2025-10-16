# Implementation Planning Prompt: F008 - Schema Parser & Validator

## Feature Metadata
- **Feature ID**: F008
- **Name**: Schema Parser & Validator
- **Feature Branch**: `006-f008-schema-parser`
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F007 (Schema JSON specification must be defined)
- **Spec Location**: `.specify/prompts/specify/mvp/f008-schema-parser-validator.md`

## Constitutional Alignment

### Core Principles Verification
- ✅ **Verticalized > Generic**: N/A (foundation infrastructure) - Parser validates verticalized schemas
- ✅ **Speed > Features**: Fast schema loading (<100ms) is critical for CLI startup time
- ✅ **Local-First, Cloud Optional**: Pure local operation, no network dependencies
- ✅ **Boring Tech Wins**: Go standard library (encoding/json), proven patterns, no exotic dependencies
- ✅ **Open Source Forever**: Parser is core CLI functionality, MIT licensed
- ✅ **Developer-First Design**: Clear error messages enable self-service debugging
- ✅ **Ship Fast, Validate Early**: TDD required, 3-day implementation, foundational for all data generation

### Technical Constraints Verification
- ✅ **Performance**: <100ms schema loading (constitutional requirement), <30s total data generation
- ✅ **Distribution**: Parser embedded in binary, schemas loaded from embed.FS
- ✅ **Database Support**: Validates schemas for both MySQL and PostgreSQL
- ✅ **Cost**: $0 (standard library, no external dependencies)
- ✅ **Code Quality**: TDD required, 100% coverage target for core parser logic
- ✅ **License**: MIT, uses only standard library (no license conflicts)
- ✅ **Platform Support**: Cross-platform (Go stdlib encoding/json works everywhere)

### Legal Constraints Verification (CRITICAL)
- ✅ **Independent Development**: Go standard library patterns, public JSON parsing documentation
- ✅ **No Employer References**: N/A (technical implementation)
- ✅ **Public Information Only**: All patterns from public Go documentation, JSON parsing best practices
- ✅ **Open Source Protection**: MIT license, Go standard library (BSD license, compatible)
- ✅ **Illustrative Examples Only**: Parser validates generic schema format (fintech, healthcare, retail)

## Planning Context

### Feature Summary
Build a robust schema parser in `/pkg/schema/parser.go` that loads JSON schema files (defined by F007), unmarshals into Go structs, validates all constraints (required fields, valid types, foreign key integrity, generation order), and returns clear, actionable error messages. Use Go's standard `encoding/json` for parsing. Implement comprehensive validation: required fields, data type validation, foreign key references, table ordering, no duplicates. TDD required with 100% test coverage. Foundation for all data generation features (F010 fintech, F016 healthcare, F018 retail).

### Key Technical Decisions Required

**Phase 0 Research Topics**:

1. **Go Struct Design Patterns for Schema Representation**: How to structure Go structs to match F007 JSON schema format? What field tags (`json:"name"`)? How to handle nested objects (tables, columns, foreign keys)? What pointer vs value types? How to handle optional fields?

2. **JSON Unmarshaling Strategies**: Strict vs lenient parsing? Use `DisallowUnknownFields()` to catch typos? How to handle missing optional fields vs missing required fields? Custom UnmarshalJSON methods needed? Error handling patterns?

3. **Validation Architecture**: When to validate (during parsing vs after)? How to structure validation functions (schema-level, table-level, column-level)? How to accumulate errors vs fail-fast? How to validate referential integrity (foreign keys exist)?

4. **Error Message Quality Patterns**: How to provide actionable error messages? What context to include (file path, line number if possible, exact field)? How to suggest fixes? Error message templates?

5. **File vs Embedded Schema Loading**: How to support both file-based (`LoadSchema(path)`) and embedded (`embed.FS`) loading? How to abstract reading mechanism? How to handle different file paths across platforms?

6. **Data Type Validation Approach**: How to validate SQL data types (int, varchar(n), decimal(p,s))? How to handle MySQL vs PostgreSQL differences? What types are supported (from F007)? How to validate type format (varchar needs length)?

7. **Foreign Key Integrity Validation**: How to validate foreign keys reference existing tables? How to validate foreign keys reference existing columns? How to validate foreign keys reference primary keys? How to validate referential integrity actions (CASCADE, SET NULL, RESTRICT)?

8. **Generation Order Dependency Checking**: How to validate generation_order includes all tables? How to validate no duplicates in generation_order? How to validate parent tables come before child tables? How to detect circular dependencies?

9. **Test-Driven Development Patterns for Parsers**: How to structure parser tests? Table-driven tests for valid/invalid schemas? How to test error messages? How to achieve 100% coverage? How to test edge cases (empty files, malformed JSON, missing fields)?

10. **Performance Optimization Strategies**: How to optimize JSON parsing? How to minimize allocations? How to make validation efficient (O(n) where n = schema size)? Target: <100ms for typical schemas (<10 tables, <50 columns total)

### Technical Context (Pre-filled)

**Language/Version**: Go 1.21+ (existing project configuration from F003/F004)
**Primary Dependencies**:
  - encoding/json (Go standard library for JSON parsing)
  - github.com/stretchr/testify (testing assertions, already in project from F006)
**Storage**: Schemas loaded from files (development) or embed.FS (production binary)
**Testing**: TDD required with 100% coverage target for pkg/schema/parser.go and types.go
**Target Platform**: Cross-platform (macOS Intel/ARM, Linux x86_64/ARM64, Windows x86_64)
**Project Type**: CLI tool library (pkg/schema package)
**Performance Goals**:
  - Schema parsing: <100ms for typical schemas (<10 tables, <50 columns)
  - JSON unmarshaling: O(n) where n = schema size
  - Validation: O(n) where n = number of tables + columns
  - Target measured on 2020 MacBook Pro (constitutional hardware baseline)
**Constraints**:
  - Must validate all F007 schema rules (required fields, valid types, foreign key integrity)
  - Must provide clear, actionable error messages (UX Principle 4: Fail Gracefully)
  - Must use Go standard library encoding/json (no external JSON libraries)
  - Must support both file and embedded schema loading
  - TDD required (100% coverage for core parser logic)
  - Must be cross-platform compatible (no platform-specific file operations)
  - Error messages must indicate exactly what's wrong and where
**Scale/Scope**: Foundation for all schema-based features (F010 fintech, F016 healthcare, F018 retail, F013 data generation engine)

## Planning Workflow

### Phase 0: Research & Technical Decisions

Generate `research.md` with documented decisions for:

#### 1. Go Struct Design Patterns for Schema Representation
- **Decision Point**: How to structure Go structs to match F007 JSON schema format?
- **Research**: Go JSON struct tags, pointer vs value types, optional field handling, nested struct patterns
- **Output**: Complete struct design for types.go:
  - **Schema struct**: Top-level schema representation
    - Fields: SchemaVersion, Name, Description, Author, Version, DatabaseType, Metadata, Tables, Relationships, GenerationOrder, ValidationRules
    - Tags: `json:"schema_version"` format (snake_case to match JSON)
  - **SchemaMetadata struct**: Metadata container
    - Fields: TotalRecords, Industry, Tags
  - **Table struct**: Table representation
    - Fields: Name, Description, RecordCount, Columns, Indexes
  - **Column struct**: Column representation
    - Fields: Name, Type, PrimaryKey, AutoIncrement, Nullable, Unique, Default, Generator, GeneratorParams, ForeignKey
    - GeneratorParams: map[string]interface{} for flexibility
  - **ForeignKey struct**: Foreign key constraint
    - Fields: Table, Column, OnDelete, OnUpdate
  - **Index struct**: Table index
    - Fields: Name, Columns, Unique
  - **Relationship struct**: Explicit relationship documentation
    - Fields: FromTable, FromColumn, ToTable, ToColumn, RelationshipType, Description
  - **ValidationRule struct**: Validation rule
    - Fields: Rule, Description, Enforcement
- **Pointer vs Value Types**:
  - Use pointers for optional fields (ForeignKey, Indexes, ValidationRules)
  - Use values for required fields (Name, Type, Tables)
  - Rationale: Pointers allow nil to distinguish "not set" from "set to zero value"

#### 2. JSON Unmarshaling Strategies
- **Decision Point**: Strict vs lenient parsing? How to handle unknown fields and missing required fields?
- **Research**: Go json.Decoder options, DisallowUnknownFields(), error handling patterns
- **Output**: Parsing strategy:
  - **Strict parsing**: Use `decoder.DisallowUnknownFields()` to catch typos in schema files
    - Rationale: Typos like "tabes" instead of "tables" should be caught early
  - **Error on unknown fields**: Fail if schema has fields not in struct definition
    - Helps developers catch copy-paste errors, outdated schema versions
  - **Missing optional fields**: Acceptable (use zero values or nil pointers)
  - **Missing required fields**: Caught in validation phase (ValidateSchema function)
  - **Custom UnmarshalJSON**: Not needed (standard unmarshaling is sufficient)
  - **Error wrapping**: Use fmt.Errorf("context: %w", err) for clear error chains
  - **Parsing flow**:
    1. Create json.Decoder from io.Reader
    2. Set DisallowUnknownFields()
    3. Decode into Schema struct
    4. If decode error: return with "failed to parse JSON" context
    5. Call ValidateSchema()
    6. If validation error: return with "schema validation failed" context
    7. Return validated schema

#### 3. Validation Architecture
- **Decision Point**: When to validate? How to structure validation functions? Fail-fast or accumulate errors?
- **Research**: Validation patterns, error accumulation strategies, referential integrity checking
- **Output**: Validation architecture:
  - **Validation phases**:
    1. **Schema-level validation**: ValidateSchema() checks top-level requirements
    2. **Table-level validation**: ValidateTable() checks each table
    3. **Column-level validation**: ValidateColumn() checks each column
    4. **Referential integrity validation**: After all tables parsed, check foreign keys reference existing tables/columns
    5. **Generation order validation**: ValidateGenerationOrder() checks all tables included, no duplicates
  - **Fail-fast strategy**: Return first error encountered
    - Rationale: Easier for developers to fix one error at a time
    - Multiple errors can be overwhelming
    - Constitutional principle: Clear, actionable errors
  - **Validation functions**:
    - ValidateSchema(s *Schema) error
    - ValidateTable(t *Table, existingTables map[string]bool) error
    - ValidateColumn(c *Column, existingColumns map[string]bool) error
    - ValidateDataType(dataType string) error
    - ValidateGenerationOrder(s *Schema, tableNames map[string]bool) error
  - **Error context**: Each error includes:
    - What's wrong (required field missing, invalid type, etc.)
    - Where (table name, column name, index in array)
    - How to fix (suggestion if possible)

#### 4. Error Message Quality Patterns
- **Decision Point**: How to provide clear, actionable error messages that save developer debugging time?
- **Research**: Go error formatting, constitutional UX Principle 4 (Fail Gracefully)
- **Output**: Error message standards:
  - **Bad error** ❌: "validation error"
  - **Good error** ✅: "table 2 (loans): column 5 (borrower_id): foreign key references non-existent table 'users'"
  - **Error message template**:
    ```
    {context}: {what's wrong}
    ```
  - **Context examples**:
    - "schema validation failed: name is required"
    - "table 0 (borrowers): must have at least one column"
    - "table 1 (loans): column 3 (borrower_id): foreign key references non-existent table 'users'"
    - "generation_order: table 'payments' is not in generation_order"
  - **Always include**:
    - Array index (table 0, column 3) for easy location
    - Entity name (table "loans", column "borrower_id") for clarity
    - Exact error (what's wrong)
    - No suggestions for now (keep messages concise, add in future if needed)
  - **Use fmt.Errorf** for error wrapping and context:
    ```go
    return fmt.Errorf("table %d (%s): column %d (%s): %w",
        tableIdx, table.Name, colIdx, col.Name, err)
    ```

#### 5. File vs Embedded Schema Loading
- **Decision Point**: How to support both file-based (development) and embedded (production) schema loading?
- **Research**: Go embed package, io.Reader abstraction, file path handling
- **Output**: Loading strategy:
  - **Public API**:
    - `LoadSchema(path string) (*Schema, error)` - Loads from file path
    - `ParseSchema(r io.Reader) (*Schema, error)` - Parses from any reader
  - **LoadSchema implementation**:
    1. Open file with os.Open(path)
    2. Defer close
    3. Call ParseSchema(file)
    4. Return result
  - **ParseSchema implementation**:
    1. Create json.Decoder from reader
    2. Set DisallowUnknownFields()
    3. Decode into Schema struct
    4. Validate schema
    5. Return validated schema
  - **embed.FS support** (for F013 data generation):
    ```go
    // Caller code (not in parser, in CLI)
    //go:embed schemas/*.json
    var schemasFS embed.FS

    file, _ := schemasFS.Open("schemas/fintech-loans.json")
    schema, _ := schema.ParseSchema(file)
    ```
  - **Rationale**: ParseSchema takes io.Reader, works with files, embed.FS, bytes.Reader, strings.Reader

#### 6. Data Type Validation Approach
- **Decision Point**: How to validate SQL data types? Handle MySQL vs PostgreSQL differences?
- **Research**: F007 supported types (lines 182-188), MySQL vs PostgreSQL type compatibility
- **Output**: Data type validation strategy:
  - **Supported types** (from F007):
    - Integer: int, bigint, smallint, tinyint
    - Decimal: decimal, decimal(p,s), float, double
    - String: varchar(n), text, char(n)
    - Date/Time: date, datetime, timestamp
    - Boolean: boolean, bit
    - JSON: json, jsonb
    - Enum: enum('val1','val2',...)
  - **Validation approach**:
    - Case-insensitive matching (convert to lowercase)
    - Prefix matching (varchar(100) matches "varchar")
    - Valid types list:
      ```go
      validTypes := []string{
          "int", "bigint", "smallint", "tinyint",
          "decimal", "float", "double",
          "varchar", "text", "char",
          "date", "datetime", "timestamp",
          "boolean", "bit", "json", "jsonb", "enum",
      }
      ```
    - Check if dataType starts with any valid type
    - Error if no match: "unsupported data type: {type}"
  - **MySQL vs PostgreSQL handling**:
    - Parser accepts both (jsonb valid even if MySQL doesn't support)
    - Rationale: database_type field indicates which databases supported
    - Future enhancement: Validate type compatibility with database_type

#### 7. Foreign Key Integrity Validation
- **Decision Point**: How to validate foreign keys reference existing tables and columns?
- **Research**: F007 foreign key format (inline foreign_key object), referential integrity rules
- **Output**: Foreign key validation strategy:
  - **Validation steps**:
    1. Build map of all table names: `tableNames map[string]bool`
    2. For each table, for each column:
       - If column has foreign_key:
         - Check foreign_key.Table exists in tableNames
         - Error if not: "table {t}: column {c}: foreign key references non-existent table {fk.Table}"
    3. Validate referential integrity actions:
       - Valid actions: CASCADE, SET NULL, RESTRICT, "" (empty = default)
       - Check OnDelete and OnUpdate are valid
       - Error if invalid: "invalid on_delete action: {action}"
  - **Column reference validation** (future enhancement):
    - For now: Assume foreign key references primary key (standard pattern)
    - Future: Build map of all columns per table, validate foreign_key.Column exists
  - **Primary key reference validation** (future enhancement):
    - For now: Assume foreign keys reference primary keys
    - Future: Validate referenced column is primary key or unique
  - **Rationale**: Start with table-level validation (MVP), add column-level in future

#### 8. Generation Order Dependency Checking
- **Decision Point**: How to validate generation_order includes all tables and parent tables come first?
- **Research**: F007 generation_order requirements, dependency validation patterns
- **Output**: Generation order validation strategy:
  - **Validation requirements**:
    1. generation_order array is not empty
    2. All tables in generation_order exist (reference valid table names)
    3. No duplicate tables in generation_order
    4. All tables are in generation_order (no missing tables)
  - **Validation algorithm**:
    1. Check len(s.GenerationOrder) > 0
    2. Build orderSet map[string]bool
    3. For each tableName in generation_order:
       - Check tableName exists in tableNames
       - Error if not: "generation_order references non-existent table: {name}"
       - Check tableName not in orderSet (no duplicates)
       - Error if duplicate: "generation_order contains duplicate table: {name}"
       - Add tableName to orderSet
    4. For each tableName in tableNames:
       - Check tableName exists in orderSet
       - Error if not: "table {name} is not in generation_order"
  - **Circular dependency detection** (future enhancement):
    - For now: Assume generation_order is correct (developer responsibility)
    - Future: Build dependency graph from foreign keys, detect cycles
    - Rationale: Complex graph algorithms, not MVP requirement

#### 9. Test-Driven Development Patterns for Parsers
- **Decision Point**: How to structure parser tests? How to achieve 100% coverage?
- **Research**: Go testing patterns, table-driven tests, testify assertions
- **Output**: TDD testing strategy:
  - **Test file structure**:
    - `pkg/schema/parser_test.go` - Main parser tests
    - `pkg/schema/types_test.go` - Type tests (if needed)
  - **Test categories**:
    1. **Valid schema tests**: Schema parses successfully, all fields populated
    2. **Missing required field tests**: Each required field tested (name, tables, etc.)
    3. **Invalid type tests**: Unsupported data types, invalid column types
    4. **Foreign key tests**: Non-existent table, invalid referential actions
    5. **Generation order tests**: Missing table, duplicate table, extra table
    6. **Duplicate name tests**: Duplicate table names, duplicate column names
    7. **Edge case tests**: Empty schema, malformed JSON, null values
  - **Table-driven test pattern**:
    ```go
    tests := []struct {
        name        string
        schemaJSON  string
        expectError bool
        errorContains string
    }{
        {"valid schema", validJSON, false, ""},
        {"missing name", missingNameJSON, true, "name is required"},
        // ... more test cases
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            schema, err := ParseSchema(strings.NewReader(tt.schemaJSON))
            if tt.expectError {
                require.Error(t, err)
                assert.Contains(t, err.Error(), tt.errorContains)
            } else {
                require.NoError(t, err)
                assert.NotNil(t, schema)
            }
        })
    }
    ```
  - **Coverage target**: 100% for parser.go (critical path)
  - **TDD workflow**:
    1. Write test for valid schema (fails - parser not implemented)
    2. Implement basic parser (test passes)
    3. Write test for missing required field (fails)
    4. Add validation (test passes)
    5. Repeat for each validation rule
    6. Refactor after tests pass

#### 10. Performance Optimization Strategies
- **Decision Point**: How to ensure <100ms schema loading? What optimizations?
- **Research**: Go JSON parsing performance, allocation reduction, validation efficiency
- **Output**: Performance optimization strategy:
  - **JSON parsing**: Go standard library encoding/json is fast (no custom JSON parser needed)
  - **Validation efficiency**:
    - Single pass through tables: O(n) where n = number of tables
    - Single pass through columns: O(m) where m = number of columns
    - Map lookups for table names: O(1) lookup
    - Total complexity: O(n + m) - linear time
  - **Allocation reduction**:
    - Pre-allocate maps: `tableNames := make(map[string]bool, len(s.Tables))`
    - Avoid unnecessary string concatenations (use fmt.Errorf)
    - Reuse readers (io.Reader passed as parameter)
  - **Typical schema performance** (estimated):
    - Small schema (3 tables, 15 columns): <1ms parsing + <1ms validation = <2ms total
    - Medium schema (10 tables, 50 columns): <5ms parsing + <5ms validation = <10ms total
    - Large schema (50 tables, 250 columns): <25ms parsing + <25ms validation = <50ms total
  - **Target**: <100ms for typical MVP schemas (3-10 tables, 15-50 columns)
  - **Measurement**: Benchmark tests if needed, not MVP requirement
  - **Optimization priority**: Correctness > Speed (get it right first, optimize if needed)

**Deliverable**: `specs/006-f008-schema-parser/research.md`

### Phase 1: Design & Contracts

#### 1. Data Model (REQUIRED - Go Struct Design)
Generate `data-model.md` documenting the Go struct architecture from types.go:

```markdown
# Data Model: Schema Parser Go Structs

## Overview
Go struct representation of F007 JSON schema format. These structs are used by the parser to unmarshal JSON schemas and provide type-safe access to schema data.

## Core Structs

### Schema
**Purpose**: Top-level schema representation
**File**: pkg/schema/types.go
**Usage**: Result of ParseSchema(), input to data generation engine

```go
type Schema struct {
    SchemaVersion    string            `json:"schema_version"`
    Name             string            `json:"name"`
    Description      string            `json:"description"`
    Author           string            `json:"author"`
    Version          string            `json:"version"`
    DatabaseType     []string          `json:"database_type"`
    Metadata         SchemaMetadata    `json:"metadata"`
    Tables           []Table           `json:"tables"`
    Relationships    []Relationship    `json:"relationships"`
    GenerationOrder  []string          `json:"generation_order"`
    ValidationRules  []ValidationRule  `json:"validation_rules,omitempty"`
}
```

**Fields**:
- SchemaVersion: Format version (e.g., "1.0")
- Name: Unique schema identifier (e.g., "fintech-loans")
- Description: Human-readable description
- Author: Schema creator
- Version: Content version (semver: "1.0.0")
- DatabaseType: Supported databases (["mysql", "postgres"])
- Metadata: Additional schema metadata
- Tables: Array of table definitions (core content)
- Relationships: Explicit foreign key documentation
- GenerationOrder: Table generation order (parent tables first)
- ValidationRules: Optional soft validation rules

### SchemaMetadata
**Purpose**: Schema metadata container
**File**: pkg/schema/types.go
**Usage**: Documentation, complexity tier classification

```go
type SchemaMetadata struct {
    TotalRecords int      `json:"total_records"`
    Industry     string   `json:"industry"`
    Tags         []string `json:"tags"`
}
```

**Fields**:
- TotalRecords: Total number of records across all tables
- Industry: Industry category (e.g., "fintech", "healthcare")
- Tags: Keywords for discovery (e.g., ["loans", "credit"])

### Table
**Purpose**: Database table representation
**File**: pkg/schema/types.go
**Usage**: Table structure, column definitions, indexes

```go
type Table struct {
    Name        string    `json:"name"`
    Description string    `json:"description"`
    RecordCount int       `json:"record_count"`
    Columns     []Column  `json:"columns"`
    Indexes     []Index   `json:"indexes,omitempty"`
}
```

**Fields**:
- Name: Table name (unique within schema)
- Description: Table purpose
- RecordCount: Number of records to generate
- Columns: Array of column definitions
- Indexes: Optional table indexes

### Column
**Purpose**: Table column representation
**File**: pkg/schema/types.go
**Usage**: Column structure, data generation parameters

```go
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
```

**Fields**:
- Name: Column name (unique within table)
- Type: SQL data type (e.g., "int", "varchar(255)", "decimal(10,2)")
- PrimaryKey: Is primary key column
- AutoIncrement: Auto-increment for integer primary keys
- Nullable: Allows NULL values
- Unique: Unique constraint
- Default: Default value expression
- Generator: Data generator name (e.g., "first_name", "credit_score")
- GeneratorParams: Generator parameters (distribution, ranges, etc.)
- ForeignKey: Foreign key constraint (pointer = optional)

### ForeignKey
**Purpose**: Foreign key constraint representation
**File**: pkg/schema/types.go
**Usage**: Referential integrity, data generation ordering

```go
type ForeignKey struct {
    Table    string `json:"table"`
    Column   string `json:"column"`
    OnDelete string `json:"on_delete,omitempty"`
    OnUpdate string `json:"on_update,omitempty"`
}
```

**Fields**:
- Table: Referenced table name
- Column: Referenced column name
- OnDelete: Referential action on delete (CASCADE, SET NULL, RESTRICT)
- OnUpdate: Referential action on update (CASCADE, SET NULL, RESTRICT)

### Index
**Purpose**: Table index representation
**File**: pkg/schema/types.go
**Usage**: Database performance optimization

```go
type Index struct {
    Name    string   `json:"name"`
    Columns []string `json:"columns"`
    Unique  bool     `json:"unique,omitempty"`
}
```

**Fields**:
- Name: Index name
- Columns: Indexed column names
- Unique: Unique index constraint

### Relationship
**Purpose**: Explicit relationship documentation (human-readable)
**File**: pkg/schema/types.go
**Usage**: Schema documentation, relationship visualization

```go
type Relationship struct {
    FromTable        string `json:"from_table"`
    FromColumn       string `json:"from_column"`
    ToTable          string `json:"to_table"`
    ToColumn         string `json:"to_column"`
    RelationshipType string `json:"relationship_type"`
    Description      string `json:"description"`
}
```

**Fields**:
- FromTable: Source table name
- FromColumn: Source column name
- ToTable: Target table name
- ToColumn: Target column name
- RelationshipType: Relationship type (one_to_one, one_to_many, many_to_one, many_to_many)
- Description: Human-readable relationship description

### ValidationRule
**Purpose**: Optional validation rule (soft constraint)
**File**: pkg/schema/types.go
**Usage**: Soft validation, documentation

```go
type ValidationRule struct {
    Rule        string `json:"rule"`
    Description string `json:"description"`
    Enforcement string `json:"enforcement"`
}
```

**Fields**:
- Rule: Rule identifier
- Description: Rule description
- Enforcement: Enforcement level (soft, hard)

## Struct Relationships

```
Schema
├── Metadata (SchemaMetadata)
├── Tables []Table
│   ├── Columns []Column
│   │   ├── GeneratorParams (map)
│   │   └── ForeignKey *ForeignKey
│   └── Indexes []Index
├── Relationships []Relationship
└── ValidationRules []ValidationRule
```

## Design Decisions

### Pointer vs Value Types
- **Pointers**: Used for optional fields (ForeignKey, Indexes, ValidationRules)
  - Rationale: Allows nil to distinguish "not set" from "set to zero value"
- **Values**: Used for required fields (Name, Type, Tables)
  - Rationale: Simpler, no nil checks needed

### map[string]interface{} for GeneratorParams
- Rationale: Generator parameters vary by generator type
- Flexibility: Supports arbitrary parameters (mean, std_dev, min, max, weights)
- Trade-off: Type safety (validated at runtime, not compile time)

### JSON Tags (snake_case)
- Rationale: Matches F007 JSON format (schema_version, foreign_key, etc.)
- Standard: Go convention for JSON unmarshaling

### Omitempty Tags
- Applied to: Optional fields (omitempty)
- Rationale: Clean JSON output (omit empty arrays, nil pointers)

## Validation Requirements

See research.md for detailed validation rules. Key validations:
- Schema: name, tables, database_type required
- Table: name, columns, record_count required; exactly one primary key
- Column: name, type required; valid data type; foreign keys reference existing tables
- ForeignKey: valid referential actions (CASCADE, SET NULL, RESTRICT)
- GenerationOrder: includes all tables, no duplicates, no missing tables
```

**Deliverable**: `specs/006-f008-schema-parser/data-model.md`

#### 2. API Contracts (SKIP for this feature)
**Rationale**: F008 is a parser library (internal package), not a REST/GraphQL API. No HTTP endpoints, no API contracts. The Go struct design (data-model.md) serves as the contract between parser and data generation engine.

#### 3. Quickstart Guide (REQUIRED - Adapted for TDD Verification)
Generate `quickstart.md` with focus on **TDD workflow and parser verification**:

```markdown
# F008 Quickstart: Schema Parser & Validator Development

## Prerequisites
- Go 1.21+ installed
- F007 schema specification complete (schemas/schema-spec.md)
- F007 example schema available (schemas/example-schema.json)
- TDD mindset (test-first development)

## Development Overview

The Schema Parser is developed using strict TDD:
1. Write test for valid schema parsing
2. Implement basic parser structure
3. Write test for validation rule (e.g., missing name)
4. Implement validation rule
5. Repeat for each validation requirement
6. Refactor after tests pass
7. Achieve 100% coverage

## TDD Workflow

### Phase 1: Setup Test Structure
```bash
# Create test file
mkdir -p pkg/schema
touch pkg/schema/parser_test.go

# Write first test (before implementation)
# Test: Parse valid minimal schema
```

### Phase 2: Implement Basic Parser
```bash
# Create implementation files
touch pkg/schema/types.go    # Go structs
touch pkg/schema/parser.go   # Parser functions

# Run test (should fail - not implemented yet)
go test ./pkg/schema/... -v

# Implement basic parser
# - Define Schema, Table, Column structs in types.go
# - Implement LoadSchema and ParseSchema in parser.go
# - Run test (should pass)
```

### Phase 3: Add Validation (TDD Cycle)
For each validation rule:
1. Write test expecting specific error
2. Run test (should fail)
3. Implement validation
4. Run test (should pass)
5. Refactor if needed

**Validation rules to test**:
- Schema name required
- At least one table required
- Database type required and valid
- Table name required
- At least one column per table
- Column name and type required
- Exactly one primary key per table
- Positive record count
- Valid data types
- Foreign keys reference existing tables
- Valid referential integrity actions
- Generation order includes all tables
- No duplicate tables in generation order
- No duplicate table names
- No duplicate column names within table

### Phase 4: Edge Cases
Test edge cases:
- Empty file
- Malformed JSON (syntax errors)
- Unknown fields (should error with DisallowUnknownFields)
- Null values for required fields
- Empty arrays for required arrays
- Circular dependencies (future enhancement)

## Running Tests

### Run all tests
```bash
go test ./pkg/schema/... -v
```

### Run with coverage
```bash
go test ./pkg/schema/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run specific test
```bash
go test ./pkg/schema/... -v -run TestParseValidSchema
```

### Check coverage percentage
```bash
go test ./pkg/schema/... -cover
# Target: 100% for parser.go
```

## Parser Verification Steps

### 1. Verify Valid Schema Parsing
```bash
# Test with F007 example schema
go test ./pkg/schema/... -v -run TestParseValidSchema

# Manual verification (optional)
cd pkg/schema
go run test_parse.go ../../schemas/example-schema.json
```

Expected: Schema parses successfully, all fields populated correctly

### 2. Verify Missing Required Fields
```bash
# Run tests for missing required fields
go test ./pkg/schema/... -v -run TestParseMissing
```

Expected: Clear error messages:
- "schema name is required"
- "schema must have at least one table"
- "database_type is required"

### 3. Verify Data Type Validation
```bash
# Run data type validation tests
go test ./pkg/schema/... -v -run TestValidateDataType
```

Expected:
- Valid types pass (int, varchar(255), decimal(10,2))
- Invalid types error ("unsupported data type: foobar")

### 4. Verify Foreign Key Validation
```bash
# Run foreign key tests
go test ./pkg/schema/... -v -run TestForeignKey
```

Expected:
- Valid foreign keys pass
- Non-existent table error ("foreign key references non-existent table 'users'")
- Invalid referential action error ("invalid on_delete action: FOOBAR")

### 5. Verify Generation Order Validation
```bash
# Run generation order tests
go test ./pkg/schema/... -v -run TestGenerationOrder
```

Expected:
- Valid generation order passes
- Missing table error ("table 'payments' is not in generation_order")
- Duplicate table error ("generation_order contains duplicate table: 'loans'")

### 6. Verify Error Message Quality
```bash
# Run all tests and review error messages
go test ./pkg/schema/... -v 2>&1 | grep "Error:"
```

Expected: All error messages are clear and actionable:
- ✅ "table 2 (loans): column 5 (borrower_id): foreign key references non-existent table 'users'"
- ❌ "validation error" (too vague)

### 7. Verify Performance
```bash
# Run benchmark (optional)
go test ./pkg/schema/... -bench=BenchmarkParseSchema

# Manual timing test
time go run test_parse.go ../../schemas/example-schema.json
```

Expected: <100ms for typical schemas (<10 tables, <50 columns)

### 8. Verify Test Coverage
```bash
# Generate coverage report
go test ./pkg/schema/... -coverprofile=coverage.out
go tool cover -func=coverage.out | grep parser.go

# View coverage in browser
go tool cover -html=coverage.out
```

Expected: 100% coverage for parser.go (all validation paths tested)

## Verification Checklist

### Implementation Complete
- [ ] types.go defines all structs (Schema, Table, Column, ForeignKey, etc.)
- [ ] parser.go implements LoadSchema(path) and ParseSchema(reader)
- [ ] parser.go implements ValidateSchema() and validation functions
- [ ] JSON parsing uses DisallowUnknownFields()
- [ ] Error messages are clear and actionable

### Tests Complete
- [ ] Test for valid schema parsing passes
- [ ] Tests for all required fields pass
- [ ] Tests for invalid data types pass
- [ ] Tests for foreign key validation pass
- [ ] Tests for generation order validation pass
- [ ] Tests for duplicate names pass
- [ ] Tests for edge cases pass (empty file, malformed JSON)
- [ ] All tests use table-driven test pattern
- [ ] Test coverage is 100% for parser.go

### Validation Complete
- [ ] F007 example schema parses successfully
- [ ] Invalid schemas return clear error messages
- [ ] Parser handles both file and io.Reader input
- [ ] Performance is <100ms for typical schemas
- [ ] go vet ./pkg/schema/... passes
- [ ] go test ./pkg/schema/... passes

### Integration Ready
- [ ] Parser can be imported by F013 data generation engine
- [ ] Parser validates F010 fintech schema (once created)
- [ ] Parser validates F016 healthcare schema (once created)
- [ ] Parser validates F018 retail schema (once created)

## Common Issues and Fixes

### Issue: Test fails with "unknown field" error
- **Cause**: JSON has field not in Go struct
- **Fix**: Add field to struct in types.go with correct JSON tag

### Issue: Test fails with "name is required" but name is present
- **Cause**: JSON field name doesn't match struct tag (e.g., "schemaName" vs "schema_name")
- **Fix**: Use snake_case in JSON to match F007 format

### Issue: Foreign key validation not working
- **Cause**: Table names map not populated before validation
- **Fix**: Build tableNames map before validating foreign keys

### Issue: Coverage is not 100%
- **Cause**: Missing test for validation path
- **Fix**: Add test for uncovered error path (check with go tool cover)

### Issue: Error messages are unclear
- **Cause**: Missing context in error wrapping
- **Fix**: Use fmt.Errorf("context: %w", err) pattern

### Issue: Performance is slow (>100ms)
- **Cause**: Inefficient validation (nested loops, string concatenation)
- **Fix**: Use maps for O(1) lookup, avoid unnecessary allocations

## Next Steps

After F008 parser is complete:
- F013: Data Generation Engine (uses parser to load schemas)
- F010: Fintech Schema (validated by parser)
- F016: Healthcare Schema (validated by parser)
- F018: Retail Schema (validated by parser)

## Testing Best Practices

### Test Organization
```go
// Good test structure
func TestParseValidSchema(t *testing.T) { /* minimal valid schema */ }
func TestParseMissingName(t *testing.T) { /* missing name field */ }
func TestParseMissingTables(t *testing.T) { /* empty tables array */ }
```

### Table-Driven Tests
```go
// Use table-driven tests for multiple cases
tests := []struct {
    name        string
    schemaJSON  string
    expectError bool
    errorContains string
}{
    {"valid schema", validJSON, false, ""},
    {"missing name", missingNameJSON, true, "name is required"},
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // test logic
    })
}
```

### Error Assertions
```go
// Use testify for clear assertions
require.Error(t, err) // Test fails if err is nil
assert.Contains(t, err.Error(), "expected substring")
```
```

**Deliverable**: `specs/006-f008-schema-parser/quickstart.md`

#### 4. Update Agent Context
Run: `.specify/scripts/bash/update-agent-context.sh claude`

This updates the Claude-specific context file (CLAUDE.md) with:
- Schema parser package overview (pkg/schema)
- Parser functions (LoadSchema, ParseSchema)
- Validation functions (ValidateSchema, ValidateTable, ValidateColumn, etc.)
- Go struct design (Schema, Table, Column, ForeignKey)
- Error message standards (clear, actionable errors)
- TDD requirements (100% coverage target)
- Usage examples (how to load and validate schemas)

**Deliverable**: Updated `CLAUDE.md` with parser implementation context

## Constitution Re-verification

After Phase 1 design, verify:
- [ ] Go structs match F007 JSON format exactly (Boring Tech, Simple > Complex)
- [ ] Validation is comprehensive (Code Quality Standard 5: >80% coverage, targeting 100%)
- [ ] Error messages are clear and actionable (UX Principle 4: Fail Gracefully, Developer-First Design)
- [ ] Parser uses Go standard library (Boring Tech Wins: no exotic dependencies)
- [ ] Performance target is achievable (<100ms for typical schemas)
- [ ] TDD workflow is defined (Development Practice 1: TDD Required for Core Functionality)
- [ ] Parser supports both file and embedded loading (Local-First: works offline)
- [ ] Design is simple, not over-engineered (Anti-Pattern 5: Simple > Complex)
- [ ] Parser is cross-platform compatible (Technical Constraint 7: Platform Support)
- [ ] Implementation is 3 days (Ship Fast: 12-week MVP on track)

## Deliverables Summary

**Generated by /speckit.plan**:
1. ✅ `specs/006-f008-schema-parser/plan.md` - This file
2. ✅ `specs/006-f008-schema-parser/research.md` - Phase 0 output (10 parser design decisions)
3. ✅ `specs/006-f008-schema-parser/data-model.md` - Phase 1 output (Go struct design)
4. ✅ `specs/006-f008-schema-parser/quickstart.md` - Phase 1 output (TDD verification guide)
5. ⏭️ `specs/006-f008-schema-parser/contracts/` - SKIP (no REST/GraphQL APIs)
6. ✅ Updated CLAUDE.md with parser implementation context

**NOT generated by /speckit.plan** (created later):
- `specs/006-f008-schema-parser/tasks.md` - Phase 2, separate command (/speckit.tasks)
- `pkg/schema/types.go` - Created during implementation (Go structs)
- `pkg/schema/parser.go` - Created during implementation (parser functions)
- `pkg/schema/parser_test.go` - Created during implementation (TDD tests)

## Success Criteria for Planning Phase

- ✅ All 10 research decisions documented with clear rationale
- ✅ Go struct design follows best practices (pointer vs value, JSON tags, optional fields)
- ✅ JSON unmarshaling strategy is strict (DisallowUnknownFields)
- ✅ Validation architecture is comprehensive (schema, table, column, foreign key, generation order)
- ✅ Error message standards are clear (actionable, contextual)
- ✅ File and embedded schema loading strategy defined
- ✅ Data type validation approach covers MySQL and PostgreSQL
- ✅ Foreign key integrity validation strategy defined
- ✅ Generation order validation strategy defined
- ✅ TDD patterns documented (table-driven tests, 100% coverage)
- ✅ Performance optimization strategy defined (<100ms target)
- ✅ Data model documented (Go struct architecture)
- ✅ Quickstart provides clear TDD workflow and verification steps
- ✅ Constitution compliance verified (no violations)
- ✅ Agent context updated with parser implementation guidance
- ✅ Planning artifacts reference constitution and F007 spec correctly

## Anti-Patterns to Avoid

- ❌ Don't skip TDD (100% coverage required for core parser)
- ❌ Don't use external JSON parsing libraries (Go stdlib is sufficient)
- ❌ Don't accumulate validation errors (fail-fast for clarity)
- ❌ Don't provide vague error messages ("validation error" not acceptable)
- ❌ Don't ignore edge cases (empty files, malformed JSON, null values)
- ❌ Don't skip performance measurement (benchmark if needed)
- ❌ Don't use custom JSON unmarshaling unless necessary (keep it simple)
- ❌ Don't validate column-level foreign key references in MVP (table-level is sufficient)
- ❌ Don't detect circular dependencies in MVP (complex graph algorithms)
- ❌ Don't over-optimize prematurely (get correctness first, optimize if needed)
- ❌ Don't skip DisallowUnknownFields (catch typos early)
- ❌ Don't hardcode validation rules (make them clear and testable)
- ❌ Don't forget to test error paths (every validation rule needs test)
- ❌ Don't use lenient parsing (strict parsing prevents subtle bugs)

## Implementation Notes

### For the AI Agent
When executing `/speckit.plan` with this prompt:

1. **Start with comprehensive research.md**: Document all 10 research decisions with:
   - Decision: What was chosen (Go struct design, validation approach, etc.)
   - Rationale: Why chosen (clarity, performance, simplicity, constitution alignment)
   - Alternatives considered: What else was evaluated (lenient vs strict parsing, fail-fast vs accumulate errors)
   - Source: Where information came from (Go docs, JSON parsing patterns, F007 spec)

2. **Generate detailed data-model.md**: Document Go struct architecture:
   - All structs with fields and JSON tags
   - Struct relationships (nested objects)
   - Design decisions (pointer vs value, map for GeneratorParams)
   - Validation requirements summary
   - Usage examples

3. **Focus on TDD in quickstart.md**: This is critical for F008:
   - TDD workflow (write test → implement → refactor)
   - Phase-by-phase development (setup, basic parser, validation, edge cases)
   - Running tests (go test commands)
   - Coverage measurement (go tool cover)
   - Verification checklist
   - Common issues and fixes

4. **Verify constitutional compliance**:
   - TDD Required: 100% coverage target for core parser (Development Practice 1)
   - Developer-First Design: Clear error messages enable self-service debugging (Core Principle VI)
   - Speed > Features: <100ms schema loading (Core Principle II)
   - Boring Tech Wins: Go standard library, no exotic dependencies (Core Principle IV)
   - Code Quality Standards: >80% coverage, targeting 100% for parser (Technical Constraint 5)
   - Fail Gracefully: Actionable error messages (UX Principle 4)

5. **Keep it standard**: This is intentionally straightforward Go development:
   - Use Go standard library encoding/json (no custom JSON parsers)
   - Follow Go conventions (snake_case JSON tags, CamelCase struct fields)
   - Standard validation patterns (required field checks, type validation, referential integrity)
   - Table-driven tests (Go testing best practice)
   - Fail-fast error handling (return first error, clear context)

### Parser Design Principles

**Simplicity First**:
- Go standard library over third-party dependencies
- Strict parsing over lenient (catch errors early)
- Fail-fast over error accumulation (one error at a time)
- Clear error messages over vague messages

**Validation Rigor**:
- Validate all F007 requirements (required fields, valid types, referential integrity)
- Catch typos early (DisallowUnknownFields)
- Validate at parse time (not later)
- Provide context in errors (table name, column name, array index)

**TDD Discipline**:
- Write test first, then implement
- 100% coverage for parser.go (every validation path tested)
- Table-driven tests for multiple cases
- Test edge cases (empty files, malformed JSON, null values)

**Performance Awareness**:
- Linear time validation (O(n) where n = schema size)
- Map lookups for table names (O(1) lookup)
- Minimize allocations (pre-allocate maps, reuse readers)
- Target <100ms for typical schemas

### Go Struct Design Philosophy

**Type Safety**:
- Structs match F007 JSON format exactly
- JSON tags for correct unmarshaling
- Pointer types for optional fields (nil = not set)
- Value types for required fields (simpler, no nil checks)

**Flexibility**:
- GeneratorParams as map[string]interface{} (arbitrary parameters)
- Validation rules as array (extensible)
- Metadata as struct (typed but extensible)

**Clarity**:
- Field names are descriptive (SchemaVersion, GenerationOrder, ForeignKey)
- JSON tags match F007 format (snake_case)
- Omitempty for optional fields (clean JSON output)

### Validation Strategy Philosophy

**Fail-Fast**:
- Return first error encountered
- Easier for developers (fix one error at a time)
- Avoids overwhelming error lists

**Clear Errors**:
- Always include context (table name, column name, index)
- What's wrong (required field missing, invalid type, etc.)
- Use fmt.Errorf for error wrapping

**Comprehensive**:
- Schema-level validation (name, tables, database_type)
- Table-level validation (name, columns, primary key, record_count)
- Column-level validation (name, type, data type, generator)
- Referential integrity (foreign keys reference existing tables)
- Generation order (includes all tables, no duplicates)

### Related Constitution Sections
- **Core Principle II**: Speed > Features (<100ms schema loading target)
- **Core Principle IV**: Boring Tech Wins (Go standard library, no exotic dependencies)
- **Core Principle VI**: Developer-First Design (clear error messages enable self-service debugging)
- **Technical Constraint 5**: Code Quality Standards (>80% coverage, targeting 100% for parser)
- **Development Practice 1**: TDD Required for Core Functionality (schema parser is core logic)
- **UX Principle 4**: Fail Gracefully (actionable error messages with context)
- **Anti-Pattern 5**: Over-Engineering (keep parser simple, no unnecessary complexity)

## Drag-and-Drop Usage

**To use this prompt**:
1. Drag this file into Claude Code
2. Claude will execute the `/speckit.plan` workflow for F008
3. Expected outputs:
   - research.md with 10 documented parser design decisions
   - data-model.md with Go struct architecture
   - quickstart.md with TDD verification guide
   - Updated CLAUDE.md with parser implementation context
   - Constitution compliance verified

**Estimated time**: 30-40 minutes for complete planning phase

**Next command**: `/speckit.tasks` to generate tasks.md from this plan

**Success indicators**:
- Research decisions are clear and actionable (Go struct design, validation strategy, error messages)
- Go struct design matches F007 JSON format exactly
- Validation architecture is comprehensive (all F007 rules covered)
- Error message standards are clear (actionable, contextual)
- TDD workflow is well-documented (write test → implement → refactor)
- Performance target is achievable (<100ms for typical schemas)
- Data model documentation is complete (all structs, fields, relationships)
- Quickstart provides clear TDD development and verification steps
- No constitutional violations identified
- Ready to proceed to task generation (/speckit.tasks)
- Parser implementation is straightforward Go development (no exotic patterns)
