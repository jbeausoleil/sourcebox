# F008 Schema Parser & Validator - Research Decisions

**Feature**: Schema Parser & Validator
**Branch**: `006-f008-schema-parser`
**Date**: 2025-10-15
**Phase**: 0 - Research & Technical Decisions

## Overview

This document captures all technical decisions made during the research phase for the Schema Parser & Validator (F008). Each decision includes the chosen approach, rationale, alternatives considered, and source references.

The schema parser is a foundational component that loads JSON schema files (defined by F007), validates all constraints, and provides clear error messages. It must be fast (<100ms), comprehensive (100% of F007 validation rules), and developer-friendly (actionable errors).

---

## Decision 1: Go Struct Design Patterns for Schema Representation

### Decision

Use Go structs with JSON tags to mirror the F007 JSON schema format exactly. Structs will be defined in `pkg/schema/types.go`:

**Core Structs**:
- `Schema` - Top-level container (required fields as values, optional as pointers)
- `SchemaMetadata` - Metadata container (all values)
- `Table` - Table definition (required fields as values, optional as pointers)
- `Column` - Column definition (required fields as values, optional as pointers)
- `ForeignKey` - Foreign key constraint (all values, used as pointer in Column)
- `Index` - Table index (all values, used as pointer in Table)
- `Relationship` - Explicit relationship documentation (all values)
- `ValidationRule` - Optional validation rule (all values)

**Pointer vs Value Types**:
- **Pointers for optional fields**: `*ForeignKey`, `[]Index` (omitempty), `[]ValidationRule` (omitempty)
  - Rationale: Nil distinguishes "not set" from "set to zero value"
- **Values for required fields**: `string` for Name/Type, `[]Table`, `[]Column`
  - Rationale: Simpler code, no nil checks needed

**JSON Tags**:
- Use snake_case to match F007 format: `json:"schema_version"`, `json:"foreign_key"`
- Add `omitempty` for optional fields: `json:"indexes,omitempty"`

**GeneratorParams as map[string]interface{}**:
- Rationale: Generator parameters vary by type (mean/std_dev for normal, weights for weighted)
- Trade-off: Runtime validation instead of compile-time type safety
- Acceptable: Flexibility outweighs type safety for extensibility

### Rationale

**Why this approach**:
1. **Exact format matching**: Go structs mirror F007 JSON structure, enabling standard unmarshaling
2. **Type safety**: Compile-time checks for required fields, runtime checks for optional fields
3. **Clarity**: Pointer vs value distinction makes required/optional explicit in code
4. **Standard library**: Uses encoding/json (constitutional: boring tech wins)
5. **Future-proof**: Easy to extend with new fields without breaking existing schemas

**Why not alternatives**:
- **Custom unmarshaling**: Adds complexity without benefit (standard unmarshaling works)
- **Interface-based design**: Over-engineering for MVP (KISS principle)
- **String maps for everything**: Loses type safety, requires extensive runtime validation

### Alternatives Considered

**Alternative 1: Interface-based polymorphic design**
```go
type Column interface {
    GetName() string
    GetType() string
}
type PrimaryKeyColumn struct { ... }
type ForeignKeyColumn struct { ... }
```
- **Rejected**: Over-engineering, adds complexity without benefit
- **Constitutional violation**: Anti-Pattern 5 (Over-Engineering)

**Alternative 2: All fields as map[string]interface{}**
```go
type Schema struct {
    Data map[string]interface{} `json:",inline"`
}
```
- **Rejected**: Loses all type safety, requires extensive runtime validation
- **Constitutional violation**: Code Quality (strict mode, no implicit any)

**Alternative 3: Code generation from JSON Schema**
- **Rejected**: Adds build complexity, harder to understand generated code
- **Constitutional violation**: Anti-Pattern 4 (Shiny Tech without proven necessity)

### Source References

- Go JSON struct tags: https://pkg.go.dev/encoding/json#Marshal
- Go pointer semantics: https://go.dev/tour/methods/1
- F007 schema specification: `/schemas/schema-spec.md` (lines 50-117 for structure example)
- Constitutional reference: Core Principle IV (Boring Tech Wins), Anti-Pattern 5 (Over-Engineering)

---

## Decision 2: JSON Unmarshaling Strategies

### Decision

Use **strict parsing** with `decoder.DisallowUnknownFields()` to catch schema typos and format errors early. Parsing flow:

1. Create `json.Decoder` from `io.Reader`
2. Call `decoder.DisallowUnknownFields()` (strict mode)
3. Decode into `Schema` struct
4. If decode error: wrap with "failed to parse JSON" context
5. Call `ValidateSchema()` for semantic validation
6. If validation error: wrap with "schema validation failed" context
7. Return validated schema

**Error wrapping pattern**:
```go
if err := decoder.Decode(&schema); err != nil {
    return nil, fmt.Errorf("failed to parse JSON: %w", err)
}

if err := ValidateSchema(&schema); err != nil {
    return nil, fmt.Errorf("schema validation failed: %w", err)
}
```

**Missing optional fields**: Acceptable (Go's zero values or nil pointers)
**Missing required fields**: Caught in validation phase, not parsing phase
**Custom UnmarshalJSON**: Not needed (standard unmarshaling sufficient)

### Rationale

**Why strict parsing**:
1. **Catch typos early**: "tabes" instead of "tables" errors immediately
2. **Version mismatch detection**: Old schema with deprecated fields errors
3. **Developer-friendly**: Explicit errors instead of silent failures
4. **Constitutional alignment**: UX Principle 4 (Fail Gracefully with actionable errors)

**Why fail on unknown fields**:
- Copy-paste errors from examples or other schemas
- Outdated schema versions with renamed fields
- Typos in field names that would otherwise be ignored

**Why validate after parsing**:
- Separation of concerns: parsing (syntax) vs validation (semantics)
- Better error messages: "foreign key references non-existent table" vs "JSON parse error"
- Testability: Can test parsing and validation separately

### Alternatives Considered

**Alternative 1: Lenient parsing (ignore unknown fields)**
```go
// No DisallowUnknownFields() call
decoder.Decode(&schema)
```
- **Rejected**: Silent failures on typos, harder to debug
- **Example**: `"tabes": [...]` would be silently ignored, causing "no tables" error later
- **Constitutional violation**: UX Principle 4 (errors should be actionable, not vague)

**Alternative 2: Parse and validate in one pass**
```go
func (s *Schema) UnmarshalJSON(data []byte) error {
    // Custom unmarshaling + validation
}
```
- **Rejected**: Mixes concerns, harder to test, less clear error messages
- **Constitutional violation**: Anti-Pattern 5 (Over-Engineering)

**Alternative 3: Accumulate all parsing errors before returning**
```go
var errors []error
// Collect all validation errors, return all at once
```
- **Rejected**: Overwhelming for developers (10 errors at once)
- **Constitutional violation**: UX Principle 4 (Fail Gracefully = one error at a time)

### Source References

- Go json.Decoder documentation: https://pkg.go.dev/encoding/json#Decoder
- DisallowUnknownFields: https://pkg.go.dev/encoding/json#Decoder.DisallowUnknownFields
- Error wrapping: https://go.dev/blog/go1.13-errors
- Constitutional reference: UX Principle 4 (Fail Gracefully), Anti-Pattern 5 (Simple > Complex)

---

## Decision 3: Validation Architecture

### Decision

Use **multi-phase fail-fast validation** with clear error context at each level:

**Validation Phases**:
1. **Schema-level**: `ValidateSchema(s *Schema) error`
   - Check required top-level fields (name, tables, database_type, generation_order)
   - Check at least one table exists
   - Build table names map for later validation
2. **Table-level**: For each table, call `ValidateTable(t *Table, tableNames map[string]bool) error`
   - Check required fields (name, columns, record_count)
   - Check at least one column exists
   - Check exactly one primary key
   - Build column names map for foreign key validation
3. **Column-level**: For each column, call `ValidateColumn(c *Column, columnNames map[string]bool) error`
   - Check required fields (name, type)
   - Validate data type with `ValidateDataType(dataType string) error`
   - Check for duplicate column names
4. **Foreign Key Validation**: After all tables parsed
   - Check foreign_key.Table exists in tableNames map
   - Validate referential integrity actions (CASCADE, SET NULL, RESTRICT)
5. **Generation Order Validation**: `ValidateGenerationOrder(s *Schema, tableNames map[string]bool) error`
   - Check all tables in generation_order
   - Check no duplicates
   - Check no missing tables

**Fail-Fast Strategy**:
- Return first error encountered (not all errors at once)
- Easier for developers to fix one issue at a time
- Reduces cognitive load

**Error Context Pattern**:
```go
return fmt.Errorf("table %d (%s): column %d (%s): %w",
    tableIdx, table.Name, colIdx, col.Name, err)
```

### Rationale

**Why multi-phase validation**:
1. **Logical progression**: Schema → Table → Column (top-down)
2. **Clear error context**: Each phase adds context (which table, which column)
3. **Efficient referential checks**: Build maps once, reuse for lookups
4. **Separation of concerns**: Each function has single responsibility

**Why fail-fast**:
1. **Developer-friendly**: Fix one error, re-run parser, repeat
2. **Reduces overwhelm**: 10 errors at once is harder to process than 1
3. **Constitutional alignment**: UX Principle 4 (Fail Gracefully)

**Why build maps**:
- O(1) lookup for foreign key validation (check if table exists)
- O(n) total validation complexity (linear time)
- Constitutional alignment: Performance Constraint 1 (<100ms parsing)

### Alternatives Considered

**Alternative 1: Accumulate all errors, return list**
```go
var errors []error
// Collect all validation errors
return errors.Join(errors...)
```
- **Rejected**: Overwhelming, harder to prioritize which error to fix first
- **Constitutional violation**: UX Principle 4 (errors should be actionable)

**Alternative 2: Single monolithic validation function**
```go
func ValidateEverything(s *Schema) error {
    // 500 lines of validation logic
}
```
- **Rejected**: Unmaintainable, hard to test, unclear error context
- **Constitutional violation**: Anti-Pattern 5 (Simple > Complex)

**Alternative 3: Validation during parsing (custom UnmarshalJSON)**
```go
func (s *Schema) UnmarshalJSON(data []byte) error {
    // Parse + validate in one pass
}
```
- **Rejected**: Mixes concerns, unclear where errors come from (parse vs validation)
- **Constitutional violation**: Code Quality (testable design via TDD)

### Source References

- Go error handling: https://go.dev/blog/go1.13-errors
- Validation patterns: https://github.com/go-ozzo/ozzo-validation (inspiration, not dependency)
- F007 validation rules: `/schemas/schema-spec.md` (Validation Rules section)
- Constitutional reference: UX Principle 4 (Fail Gracefully), Performance Constraint 1 (<100ms)

---

## Decision 4: Error Message Quality Patterns

### Decision

Use **structured error messages** with context, location, and exact error:

**Error Message Template**:
```
{context}: {what's wrong}
```

**Context Includes**:
- Array index for easy location: "table 0", "column 3"
- Entity name for clarity: "(borrowers)", "(borrower_id)"
- Hierarchy: "table 2 (loans): column 5 (borrower_id)"

**Examples**:
- ✅ Good: `"table 2 (loans): column 5 (borrower_id): foreign key references non-existent table 'users'"`
- ✅ Good: `"generation_order: table 'payments' is not in generation_order"`
- ✅ Good: `"schema validation failed: name is required"`
- ❌ Bad: `"validation error"` (too vague)
- ❌ Bad: `"error"` (no context)

**Implementation Pattern**:
```go
// At column validation level
if err := ValidateColumn(col); err != nil {
    return fmt.Errorf("table %d (%s): column %d (%s): %w",
        tableIdx, table.Name, colIdx, col.Name, err)
}

// At base validation level
if fk.Table == "" {
    return fmt.Errorf("foreign key references empty table name")
}
if !tableExists(fk.Table, tableNames) {
    return fmt.Errorf("foreign key references non-existent table '%s'", fk.Table)
}
```

**No suggestions in MVP**: Keep messages concise, add "Fix:" suggestions in future if needed

### Rationale

**Why structured errors**:
1. **Exact location**: Array index ("table 2") makes it easy to find in JSON
2. **Entity name**: Human-readable ("loans") adds context
3. **What's wrong**: Exact error ("references non-existent table 'users'")
4. **Constitutional alignment**: Core Principle VI (Developer-First Design), UX Principle 4 (Fail Gracefully)

**Why include array index**:
- JSON editors show line numbers, array indices help locate
- Easy to map "table 2" to third table in array (0-indexed)
- Consistent with developer tools (error on line N)

**Why no suggestions yet**:
- Keep MVP simple (messages already actionable with context)
- Add "Fix:" suggestions in Phase 2 based on user feedback
- Avoid verbose messages that obscure the actual error

### Alternatives Considered

**Alternative 1: Simple errors without context**
```go
return errors.New("foreign key error")
```
- **Rejected**: Developer has to debug which foreign key, which table
- **Constitutional violation**: UX Principle 4 (Fail Gracefully = actionable errors)

**Alternative 2: Include suggestions in every error**
```go
return fmt.Errorf("table missing: add tables array to your schema JSON")
```
- **Rejected**: Verbose, suggestions may be obvious, adds maintenance burden
- **Decision**: Add suggestions in Phase 2 if user feedback indicates need

**Alternative 3: Error codes instead of messages**
```go
return &ValidationError{Code: "ERR_FK_001", Table: "loans"}
```
- **Rejected**: Developer-unfriendly, requires documentation lookup
- **Constitutional violation**: UX Principle 5 (Boring CLI = clear messages, not codes)

### Source References

- Go error formatting: https://pkg.go.dev/fmt#Errorf
- Error wrapping best practices: https://go.dev/blog/go1.13-errors
- Constitutional reference: Core Principle VI (Developer-First Design), UX Principle 4 (Fail Gracefully)

---

## Decision 5: File vs Embedded Schema Loading

### Decision

Provide **two public functions** for flexible schema loading:

**Public API**:
```go
// LoadSchema loads and parses a schema from file path
func LoadSchema(path string) (*Schema, error)

// ParseSchema parses a schema from any io.Reader
func ParseSchema(r io.Reader) (*Schema, error)
```

**LoadSchema Implementation**:
```go
func LoadSchema(path string) (*Schema, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("failed to open schema file: %w", err)
    }
    defer file.Close()

    return ParseSchema(file)
}
```

**ParseSchema Implementation**:
```go
func ParseSchema(r io.Reader) (*Schema, error) {
    decoder := json.NewDecoder(r)
    decoder.DisallowUnknownFields()

    var schema Schema
    if err := decoder.Decode(&schema); err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %w", err)
    }

    if err := ValidateSchema(&schema); err != nil {
        return nil, fmt.Errorf("schema validation failed: %w", err)
    }

    return &schema, nil
}
```

**embed.FS Support** (caller code, not in parser):
```go
// In CLI package (not parser package)
//go:embed schemas/*.json
var schemasFS embed.FS

file, _ := schemasFS.Open("schemas/fintech-loans.json")
schema, _ := schema.ParseSchema(file)
```

### Rationale

**Why two functions**:
1. **LoadSchema for convenience**: Common case (load from file path)
2. **ParseSchema for flexibility**: Works with files, embed.FS, bytes.Reader, strings.Reader
3. **io.Reader abstraction**: Standard Go interface, maximum compatibility
4. **Constitutional alignment**: Core Principle VI (Developer-First = simple API)

**Why io.Reader**:
- Standard Go interface (works with anything that reads)
- Enables testing with strings.NewReader (no temp files needed)
- Supports embed.FS, os.File, bytes.Buffer, http.Response.Body
- Future-proof (works with any source: network, stdin, memory)

**Why defer close in LoadSchema**:
- Standard Go pattern for resource cleanup
- Ensures file handle released even if ParseSchema errors
- Constitutional alignment: Code Quality (go vet passing)

### Alternatives Considered

**Alternative 1: Only file-based loading**
```go
func LoadSchema(path string) (*Schema, error) {
    data, _ := os.ReadFile(path)
    json.Unmarshal(data, &schema)
}
```
- **Rejected**: Doesn't support embed.FS (needed for F013 data generation)
- **Constitutional violation**: Technical Constraint 2 (distribution via single binary with embedded schemas)

**Alternative 2: Three functions (file, bytes, reader)**
```go
func LoadSchemaFromFile(path string) (*Schema, error)
func LoadSchemaFromBytes(data []byte) (*Schema, error)
func LoadSchemaFromReader(r io.Reader) (*Schema, error)
```
- **Rejected**: Unnecessary complexity, io.Reader handles all cases
- **Constitutional violation**: Anti-Pattern 5 (Simple > Complex)

**Alternative 3: embed.FS as parser parameter**
```go
func LoadSchema(fs embed.FS, path string) (*Schema, error)
```
- **Rejected**: Couples parser to embed package, less flexible
- **Decision**: Let caller handle embed.FS, parser takes io.Reader

### Source References

- Go io.Reader interface: https://pkg.go.dev/io#Reader
- Go embed package: https://pkg.go.dev/embed
- Defer pattern: https://go.dev/tour/flowcontrol/12
- Constitutional reference: Technical Constraint 2 (Distribution), Anti-Pattern 5 (Simple > Complex)

---

## Decision 6: Data Type Validation Approach

### Decision

Use **case-insensitive prefix matching** against a list of supported types from F007:

**Supported Types** (from F007, lines 182-188):
```go
var validTypes = []string{
    // Integer types
    "int", "bigint", "smallint", "tinyint",
    // Decimal types
    "decimal", "float", "double",
    // String types
    "varchar", "text", "char",
    // Date/Time types
    "date", "datetime", "timestamp",
    // Boolean type
    "boolean", "bit",
    // JSON types
    "json", "jsonb",
    // Enum type
    "enum",
}
```

**Validation Algorithm**:
```go
func ValidateDataType(dataType string) error {
    // Normalize to lowercase for case-insensitive matching
    normalized := strings.ToLower(dataType)

    // Check if normalized type starts with any valid type
    for _, validType := range validTypes {
        if strings.HasPrefix(normalized, validType) {
            return nil // Valid type found
        }
    }

    return fmt.Errorf("unsupported data type: %s", dataType)
}
```

**Prefix Matching Examples**:
- `"varchar(255)"` matches `"varchar"` ✅
- `"DECIMAL(10,2)"` matches `"decimal"` ✅
- `"INT"` matches `"int"` ✅
- `"foobar"` matches nothing ❌

**MySQL vs PostgreSQL Handling**:
- Parser accepts both MySQL and PostgreSQL types
- `database_type` field indicates which databases supported
- `jsonb` valid even for MySQL schemas (parser doesn't validate compatibility)
- **Future enhancement**: Cross-check type against database_type (Phase 2)

### Rationale

**Why prefix matching**:
1. **Handles parameterized types**: `varchar(255)`, `decimal(10,2)`, `char(50)`
2. **Case-insensitive**: `INT` and `int` both valid
3. **Simple to implement**: Standard library only (strings.HasPrefix)
4. **Constitutional alignment**: Core Principle IV (Boring Tech Wins)

**Why not database-specific validation yet**:
- MVP focus: Validate syntax, not cross-database compatibility
- Future enhancement: Check `jsonb` only valid if "postgres" in database_type
- Decision: Ship basic type validation first, add compatibility checks in Phase 2

**Why lowercase normalization**:
- SQL data types are case-insensitive by convention
- Developers may use `INT`, `int`, or `Int`
- Normalization makes validation consistent

### Alternatives Considered

**Alternative 1: Regular expressions for type matching**
```go
var typeRegex = regexp.MustCompile(`^(int|varchar\(\d+\)|decimal\(\d+,\d+\))$`)
```
- **Rejected**: Over-engineering, harder to maintain, slower
- **Constitutional violation**: Anti-Pattern 5 (Simple > Complex)

**Alternative 2: Parse parameterized types (extract length, precision, scale)**
```go
type DataType struct {
    Base      string // "varchar"
    Length    int    // 255
    Precision int    // 10 for decimal(10,2)
    Scale     int    // 2 for decimal(10,2)
}
```
- **Rejected**: MVP doesn't need to extract parameters, just validate syntax
- **Future enhancement**: Extract parameters for data generation (F013)

**Alternative 3: Validate type compatibility with database_type**
```go
if dataType == "jsonb" && !contains(schema.DatabaseType, "postgres") {
    return fmt.Errorf("jsonb not supported by MySQL")
}
```
- **Rejected**: Adds complexity, not critical for MVP
- **Decision**: Add in Phase 2 if users request it

### Source References

- F007 supported types: `/schemas/schema-spec.md` (lines 182-188, "Supported Data Types")
- Go strings package: https://pkg.go.dev/strings
- Constitutional reference: Core Principle IV (Boring Tech Wins), Anti-Pattern 5 (Simple > Complex)

---

## Decision 7: Foreign Key Integrity Validation

### Decision

Validate **table-level foreign key references** in MVP. Column-level and primary key validation deferred to Phase 2.

**MVP Validation** (table exists):
```go
func ValidateForeignKeys(schema *Schema, tableNames map[string]bool) error {
    for tableIdx, table := range schema.Tables {
        for colIdx, col := range table.Columns {
            if col.ForeignKey != nil {
                fk := col.ForeignKey

                // Check referenced table exists
                if !tableNames[fk.Table] {
                    return fmt.Errorf("table %d (%s): column %d (%s): foreign key references non-existent table '%s'",
                        tableIdx, table.Name, colIdx, col.Name, fk.Table)
                }

                // Check referential integrity actions
                if err := ValidateReferentialAction(fk.OnDelete, "on_delete"); err != nil {
                    return fmt.Errorf("table %d (%s): column %d (%s): %w",
                        tableIdx, table.Name, colIdx, col.Name, err)
                }
                if err := ValidateReferentialAction(fk.OnUpdate, "on_update"); err != nil {
                    return fmt.Errorf("table %d (%s): column %d (%s): %w",
                        tableIdx, table.Name, colIdx, col.Name, err)
                }
            }
        }
    }
    return nil
}

func ValidateReferentialAction(action, actionType string) error {
    if action == "" {
        return nil // Empty is valid (database default)
    }

    validActions := []string{"CASCADE", "SET NULL", "RESTRICT"}
    for _, valid := range validActions {
        if strings.EqualFold(action, valid) {
            return nil
        }
    }

    return fmt.Errorf("invalid %s action: %s (valid: CASCADE, SET NULL, RESTRICT)", actionType, action)
}
```

**Future Enhancements** (Phase 2):
1. **Column existence**: Validate `foreign_key.Column` exists in referenced table
2. **Primary key reference**: Validate referenced column is primary key or unique
3. **Type compatibility**: Validate foreign key type matches referenced column type

### Rationale

**Why table-level only in MVP**:
1. **Covers 90% of errors**: Referencing wrong table name is most common mistake
2. **Simple implementation**: Single map lookup (O(1))
3. **Constitutional alignment**: Ship Fast (MVP in 12 weeks, not 6 months)

**Why defer column-level validation**:
- Requires building map of all columns per table (more complexity)
- Most schemas reference primary keys by convention (column name = "id")
- Add when user feedback indicates need

**Why validate referential actions**:
- Invalid actions (e.g., "FOOBAR") would cause database errors
- Simple string comparison (minimal overhead)
- Improves developer experience (catch errors early)

### Alternatives Considered

**Alternative 1: Full referential integrity validation in MVP**
```go
// Build map of columns per table
tableColumns := make(map[string]map[string]*Column)
// Validate foreign_key.Column exists
// Validate referenced column is primary key
// Validate type compatibility
```
- **Rejected**: Adds complexity, delays MVP launch
- **Constitutional violation**: Core Principle VII (Ship Fast), Anti-Pattern 5 (Over-Engineering)

**Alternative 2: No foreign key validation (defer to database)**
```go
// Skip foreign key validation entirely
// Let MySQL/Postgres error at table creation time
```
- **Rejected**: Developer-unfriendly, cryptic database errors
- **Constitutional violation**: Core Principle VI (Developer-First Design)

**Alternative 3: Graph-based dependency validation**
```go
// Build dependency graph from foreign keys
// Topologically sort to detect circular dependencies
```
- **Rejected**: Complex algorithm, not MVP requirement
- **Decision**: Add circular dependency detection in Phase 2 if needed

### Source References

- F007 foreign key format: `/schemas/schema-spec.md` (Foreign Key Relationships section)
- MySQL foreign key docs: https://dev.mysql.com/doc/refman/8.0/en/create-table-foreign-keys.html
- PostgreSQL foreign key docs: https://www.postgresql.org/docs/current/ddl-constraints.html#DDL-CONSTRAINTS-FK
- Constitutional reference: Core Principle VII (Ship Fast), Anti-Pattern 5 (Over-Engineering)

---

## Decision 8: Generation Order Dependency Checking

### Decision

Validate **generation_order completeness** in MVP (all tables included, no duplicates, no missing tables). Circular dependency detection deferred to Phase 2.

**MVP Validation**:
```go
func ValidateGenerationOrder(schema *Schema, tableNames map[string]bool) error {
    // Check generation_order is not empty
    if len(schema.GenerationOrder) == 0 {
        return fmt.Errorf("generation_order is empty (must include all tables)")
    }

    // Build set of tables in generation_order
    orderSet := make(map[string]bool)

    for _, tableName := range schema.GenerationOrder {
        // Check table exists in schema
        if !tableNames[tableName] {
            return fmt.Errorf("generation_order references non-existent table: %s", tableName)
        }

        // Check for duplicates
        if orderSet[tableName] {
            return fmt.Errorf("generation_order contains duplicate table: %s", tableName)
        }

        orderSet[tableName] = true
    }

    // Check all tables are in generation_order
    for tableName := range tableNames {
        if !orderSet[tableName] {
            return fmt.Errorf("table '%s' is not in generation_order", tableName)
        }
    }

    return nil
}
```

**Future Enhancements** (Phase 2):
1. **Circular dependency detection**: Build dependency graph from foreign keys, detect cycles
2. **Automatic ordering**: Generate correct generation_order from foreign key relationships
3. **Ordering suggestions**: If order invalid, suggest correct order

### Rationale

**Why completeness validation only**:
1. **Covers 90% of errors**: Missing table or duplicate table are most common
2. **Simple implementation**: Two map lookups (O(n) total)
3. **Constitutional alignment**: Ship Fast (MVP in 12 weeks)

**Why defer circular dependency detection**:
- Complex graph algorithm (DFS or topological sort)
- Rare in practice (developers usually notice circular refs)
- Add when user feedback indicates need

**Why no automatic ordering yet**:
- Requires graph algorithm to compute topological sort
- MVP assumes developers manually order tables (documented in F007)
- Add in Phase 2 as convenience feature

### Alternatives Considered

**Alternative 1: Circular dependency detection in MVP**
```go
// Build dependency graph from foreign keys
graph := buildDependencyGraph(schema)
// Detect cycles using DFS
if hasCycle(graph) {
    return fmt.Errorf("circular dependency detected")
}
```
- **Rejected**: Complex algorithm, delays MVP, rare edge case
- **Constitutional violation**: Core Principle VII (Ship Fast), Anti-Pattern 5 (Over-Engineering)

**Alternative 2: Automatic generation_order computation**
```go
// Topologically sort tables based on foreign keys
generationOrder := topologicalSort(schema)
schema.GenerationOrder = generationOrder
```
- **Rejected**: Changes schema (parser should validate, not modify)
- **Decision**: Add as separate tool in Phase 2 (`sourcebox validate --fix`)

**Alternative 3: No generation_order validation (trust developer)**
```go
// Skip generation_order validation
// Let data generation fail at runtime
```
- **Rejected**: Runtime errors are harder to debug than parse-time errors
- **Constitutional violation**: Core Principle VI (Developer-First Design)

### Source References

- F007 generation_order requirements: `/schemas/schema-spec.md` (Generation Order section)
- Graph algorithms: https://en.wikipedia.org/wiki/Topological_sorting
- Constitutional reference: Core Principle VII (Ship Fast), Anti-Pattern 5 (Simple > Complex)

---

## Decision 9: Test-Driven Development Patterns for Parsers

### Decision

Use **table-driven tests** with strict TDD discipline for 100% parser coverage:

**Test Structure**:
```
pkg/schema/
├── parser.go           # Implementation (written after tests)
├── parser_test.go      # Tests (written first)
├── types.go            # Go structs (written with parser)
└── types_test.go       # Type tests (if needed)
```

**Test Categories** (in order of implementation):
1. **Valid schema test**: Minimal valid schema parses successfully
2. **Missing required field tests**: Each required field tested
3. **Invalid type tests**: Unsupported data types
4. **Foreign key tests**: Non-existent table, invalid actions
5. **Generation order tests**: Missing table, duplicate, extra table
6. **Duplicate name tests**: Duplicate table/column names
7. **Edge case tests**: Empty file, malformed JSON, null values

**Table-Driven Test Pattern**:
```go
func TestParseSchema(t *testing.T) {
    tests := []struct {
        name          string
        schemaJSON    string
        expectError   bool
        errorContains string
    }{
        {
            name: "valid minimal schema",
            schemaJSON: `{
                "schema_version": "1.0",
                "name": "test",
                "description": "Test schema",
                "author": "Test",
                "version": "1.0.0",
                "database_type": ["mysql"],
                "tables": [{
                    "name": "users",
                    "record_count": 100,
                    "columns": [{
                        "name": "id",
                        "type": "int",
                        "primary_key": true
                    }]
                }],
                "generation_order": ["users"]
            }`,
            expectError: false,
        },
        {
            name: "missing schema name",
            schemaJSON: `{
                "schema_version": "1.0",
                "database_type": ["mysql"],
                "tables": []
            }`,
            expectError:   true,
            errorContains: "name is required",
        },
        // ... more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            schema, err := ParseSchema(strings.NewReader(tt.schemaJSON))
            if tt.expectError {
                require.Error(t, err)
                assert.Contains(t, err.Error(), tt.errorContains)
                assert.Nil(t, schema)
            } else {
                require.NoError(t, err)
                assert.NotNil(t, schema)
            }
        })
    }
}
```

**TDD Workflow**:
1. Write test for valid schema (fails - not implemented)
2. Implement basic parser (test passes)
3. Write test for missing name (fails - no validation)
4. Add name validation (test passes)
5. Refactor if needed (tests still pass)
6. Repeat for each validation rule

**Coverage Target**: 100% for `parser.go` (critical path)

### Rationale

**Why table-driven tests**:
1. **Scalability**: Easy to add new test cases (just add entry to table)
2. **Clarity**: Each test case is self-contained and readable
3. **Standard Go pattern**: Widely used in Go stdlib and community
4. **Constitutional alignment**: Code Quality Standard 5 (>80% coverage, targeting 100%)

**Why TDD workflow**:
1. **Testable design**: Writing tests first forces simple, testable API
2. **No missed validation**: Each validation rule has corresponding test
3. **Refactoring safety**: Tests pass after refactor = correct behavior preserved
4. **Constitutional alignment**: Development Practice 1 (TDD Required for Core Functionality)

**Why 100% coverage target**:
- Parser is critical path (all data generation depends on it)
- Validation errors are common (missing fields, typos, etc.)
- High coverage ensures no edge cases missed
- Constitutional alignment: Code Quality (>80% required, 100% ideal for core logic)

### Alternatives Considered

**Alternative 1: Integration tests only (no unit tests)**
```go
func TestParseRealSchemas(t *testing.T) {
    // Load fintech-loans.json, healthcare.json, retail.json
    // Assert they parse without errors
}
```
- **Rejected**: Doesn't test error paths, hard to test edge cases
- **Constitutional violation**: Development Practice 1 (TDD Required)

**Alternative 2: Test-after (write implementation first)**
```go
// Implement parser completely, then write tests
```
- **Rejected**: Often results in lower coverage, harder to test
- **Constitutional violation**: Development Practice 1 (TDD Required for Core Functionality)

**Alternative 3: Individual test functions (no table-driven)**
```go
func TestParseMissingName(t *testing.T) { ... }
func TestParseMissingTables(t *testing.T) { ... }
// ... 50 separate test functions
```
- **Rejected**: Verbose, harder to maintain, duplication
- **Decision**: Table-driven is more maintainable for parser tests

### Source References

- Table-driven tests: https://go.dev/wiki/TableDrivenTests
- Go testing package: https://pkg.go.dev/testing
- testify assertions: https://pkg.go.dev/github.com/stretchr/testify/assert
- Constitutional reference: Development Practice 1 (TDD Required), Code Quality Standard 5 (>80% coverage)

---

## Decision 10: Performance Optimization Strategies

### Decision

Use **simple linear-time algorithms** with pre-allocated maps for O(n) total validation complexity:

**Optimization Techniques**:
1. **Pre-allocate maps**: `tableNames := make(map[string]bool, len(schema.Tables))`
2. **Single pass validation**: Validate each entity once (no backtracking)
3. **Map lookups for references**: O(1) lookup (not O(n) search)
4. **Avoid string concatenation**: Use `fmt.Errorf` for error messages
5. **Reuse io.Reader**: Pass reader as parameter, don't re-open files

**Complexity Analysis**:
- JSON parsing: O(n) where n = JSON size (Go stdlib)
- Schema validation: O(1) - check required fields
- Table validation: O(t) where t = number of tables
- Column validation: O(c) where c = total columns across all tables
- Foreign key validation: O(c) with map lookup
- Generation order validation: O(t) with map lookup
- **Total complexity**: O(n + t + c) = O(n) where n = schema size

**Estimated Performance** (typical MVP schemas):
- Small (3 tables, 15 columns): <2ms total
- Medium (10 tables, 50 columns): <10ms total
- Large (50 tables, 250 columns): <50ms total
- **Target**: <100ms for typical schemas ✅

**Optimization Priority**: Correctness > Speed
- Get validation right first (all F007 rules enforced)
- Optimize only if benchmarks show >100ms
- Don't premature optimize

### Rationale

**Why simple algorithms**:
1. **Linear time is fast enough**: <100ms for typical schemas
2. **Standard library is optimized**: encoding/json is well-tuned
3. **Maps are efficient**: O(1) lookup for foreign key validation
4. **Constitutional alignment**: Anti-Pattern 5 (Simple > Complex, avoid premature optimization)

**Why pre-allocated maps**:
- Avoids reallocation during map growth
- Minimal overhead (just `make(map, capacity)`)
- Standard Go optimization pattern

**Why no custom JSON parser**:
- Go stdlib encoding/json is fast (native code)
- Custom parser would be complex, error-prone
- Constitutional alignment: Core Principle IV (Boring Tech Wins)

### Alternatives Considered

**Alternative 1: Custom JSON parser for speed**
```go
// Hand-written JSON parser optimized for schemas
func parseSchemaFast(data []byte) (*Schema, error) { ... }
```
- **Rejected**: Over-engineering, harder to maintain, minimal benefit
- **Constitutional violation**: Anti-Pattern 4 (Shiny Tech without proven necessity)

**Alternative 2: Parallel validation (goroutines)**
```go
// Validate tables in parallel
var wg sync.WaitGroup
for _, table := range schema.Tables {
    wg.Add(1)
    go func(t Table) {
        defer wg.Done()
        ValidateTable(&t)
    }(table)
}
wg.Wait()
```
- **Rejected**: Adds complexity, overhead of goroutines > benefit for small schemas
- **Constitutional violation**: Anti-Pattern 5 (Over-Engineering)

**Alternative 3: Lazy validation (validate on-demand)**
```go
// Skip validation during parsing
// Validate fields only when accessed
schema.Name() // validates name field
```
- **Rejected**: Fail-late instead of fail-fast, worse developer experience
- **Constitutional violation**: UX Principle 4 (Fail Gracefully = fail early with clear errors)

### Source References

- Go performance best practices: https://go.dev/doc/effective_go#performance
- Profiling guide: https://go.dev/blog/pprof
- Constitutional reference: Performance Constraint 1 (<100ms), Anti-Pattern 5 (Simple > Complex)

---

## Summary

All 10 technical decisions have been documented with clear rationale, alternatives considered, and source references. Key themes:

**Simplicity**: Use Go standard library, avoid over-engineering, simple algorithms
**Developer Experience**: Clear errors, fail-fast, structured messages with context
**Performance**: Linear-time validation, map lookups, pre-allocated structures (<100ms target)
**Constitutional Alignment**: Boring tech, ship fast, TDD required, developer-first design

**Next Steps**: Phase 1 (Design & Contracts)
- Generate `data-model.md` with complete Go struct design
- Generate `quickstart.md` with TDD workflow and verification steps
- Update agent context with parser implementation guidance
