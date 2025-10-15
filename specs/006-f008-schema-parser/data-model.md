# F008 Data Model: Schema Parser Go Structs

**Feature**: Schema Parser & Validator
**Branch**: `006-f008-schema-parser`
**Date**: 2025-10-15
**File**: `pkg/schema/types.go`

## Overview

This document defines the Go struct architecture for the Schema Parser & Validator. These structs represent the in-memory format of F007 JSON schema definitions, enabling type-safe access to schema data for validation and data generation.

The structs mirror the F007 JSON format exactly, using standard `encoding/json` unmarshaling with strict field validation.

---

## Design Principles

### 1. Exact JSON Format Mirroring

Go structs match F007 JSON structure 1:1:
- Field names use snake_case JSON tags (`json:"schema_version"`)
- Nested objects become nested structs (Table.Columns, Column.ForeignKey)
- Arrays become Go slices ([]Table, []Column)

### 2. Pointer vs Value Types

**Pointers for Optional Fields**:
- `*ForeignKey` - Not all columns have foreign keys
- `[]Index` with omitempty - Not all tables have indexes
- `[]ValidationRule` with omitempty - Validation rules are optional

**Values for Required Fields**:
- `string` for Name, Type, Description (required, never nil)
- `[]Table`, `[]Column` (required arrays, empty = validation error)
- `int` for RecordCount (required, zero = validation error)

**Rationale**: Pointers allow `nil` to distinguish "not set" from "set to zero value". Required fields use values for simpler code (no nil checks).

### 3. Flexible Generator Parameters

`GeneratorParams` as `map[string]interface{}`:
- Supports arbitrary parameters: `mean`, `std_dev`, `min`, `max`, `weights`, `ranges`
- Trade-off: Runtime validation instead of compile-time type safety
- Justification: Extensibility for custom generators outweighs type safety

### 4. Constitutional Alignment

- **Boring Tech Wins**: Standard library `encoding/json`, no custom unmarshaling
- **Simple > Complex**: Flat struct hierarchy, no interfaces or polymorphism
- **Developer-First**: Clear field names, explicit required/optional distinction

---

## Core Structs

### Schema

**Purpose**: Top-level container for complete schema definition
**File**: `pkg/schema/types.go`
**Usage**: Returned by `ParseSchema()`, input to data generation engine (F013)

```go
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
```

**Fields**:

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `SchemaVersion` | string | Yes | Format version | `"1.0"` |
| `Name` | string | Yes | Unique schema identifier | `"fintech-loans"` |
| `Description` | string | Yes | Human-readable description | `"Realistic fintech loan data"` |
| `Author` | string | Yes | Schema creator | `"SourceBox Contributors"` |
| `Version` | string | Yes | Content version (semver) | `"1.0.0"` |
| `DatabaseType` | []string | Yes | Supported databases | `["mysql", "postgres"]` |
| `Metadata` | SchemaMetadata | Yes | Additional metadata | See SchemaMetadata |
| `Tables` | []Table | Yes | Table definitions | See Table |
| `Relationships` | []Relationship | Yes | Explicit FK documentation | See Relationship |
| `GenerationOrder` | []string | Yes | Table generation order | `["borrowers", "loans"]` |
| `ValidationRules` | []ValidationRule | No | Optional validation rules | See ValidationRule |

**Validation Requirements** (enforced by parser):
- `SchemaVersion` must be "1.0" (MVP format version)
- `Name` must be non-empty and unique (lowercase-hyphenated by convention)
- `DatabaseType` must contain at least one of: "mysql", "postgres"
- `Tables` must have at least one table
- `GenerationOrder` must include all table names exactly once

---

### SchemaMetadata

**Purpose**: Additional schema metadata for documentation and classification
**File**: `pkg/schema/types.go`
**Usage**: Complexity tier classification, industry tagging, search/discovery

```go
type SchemaMetadata struct {
    TotalRecords int      `json:"total_records"`
    Industry     string   `json:"industry"`
    Tags         []string `json:"tags"`
}
```

**Fields**:

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `TotalRecords` | int | Yes | Sum of all table record_count | `4950` |
| `Industry` | string | Yes | Industry category | `"fintech"`, `"healthcare"` |
| `Tags` | []string | Yes | Keywords for discovery | `["loans", "credit", "borrowers"]` |

**Validation Requirements**:
- `TotalRecords` must be positive (sum of Table.RecordCount)
- `Industry` must be non-empty string
- `Tags` must have at least one tag

**Tier Classification** (based on TotalRecords and complexity):
- **Tier 1**: <5,000 records, <30s generation time (MVP focus)
- **Tier 2**: 5,000-50,000 records, <2min generation time (Phase 2)
- **Tier 3**: >50,000 records, <5min generation time (Community-driven)

---

### Table

**Purpose**: Database table definition with columns and constraints
**File**: `pkg/schema/types.go`
**Usage**: Schema validation, SQL table creation, data generation

```go
type Table struct {
    Name        string   `json:"name"`
    Description string   `json:"description"`
    RecordCount int      `json:"record_count"`
    Columns     []Column `json:"columns"`
    Indexes     []Index  `json:"indexes,omitempty"`
}
```

**Fields**:

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `Name` | string | Yes | Table name (unique within schema) | `"borrowers"` |
| `Description` | string | Yes | Table purpose | `"Loan borrower profiles"` |
| `RecordCount` | int | Yes | Number of records to generate | `1000` |
| `Columns` | []Column | Yes | Column definitions | See Column |
| `Indexes` | []Index | No | Table indexes | See Index |

**Validation Requirements**:
- `Name` must be unique within schema (case-insensitive check)
- `RecordCount` must be positive (>0)
- `Columns` must have at least one column
- Exactly one column must have `PrimaryKey: true`
- No duplicate column names within table

**Naming Conventions** (best practices):
- Use snake_case: `loan_applications`, not `LoanApplications`
- Use plurals for entity tables: `borrowers`, not `borrower`
- Use descriptive names: `loan_payments`, not `lp`

---

### Column

**Purpose**: Table column definition with type, constraints, and generator
**File**: `pkg/schema/types.go`
**Usage**: SQL column creation, data generation, validation

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

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `Name` | string | Yes | Column name (unique within table) | `"borrower_id"` |
| `Type` | string | Yes | SQL data type | `"int"`, `"varchar(255)"`, `"decimal(10,2)"` |
| `PrimaryKey` | bool | No | Is primary key column | `true` (exactly one per table) |
| `AutoIncrement` | bool | No | Auto-increment for integer PKs | `true` |
| `Nullable` | bool | No | Allows NULL values | `true` |
| `Unique` | bool | No | Unique constraint | `true` |
| `Default` | string | No | Default value expression | `"0"`, `"CURRENT_TIMESTAMP"` |
| `Generator` | string | No | Data generator name | `"first_name"`, `"credit_score"` |
| `GeneratorParams` | map[string]interface{} | No | Generator parameters | `{"mean": 680, "std_dev": 80}` |
| `ForeignKey` | *ForeignKey | No | Foreign key constraint | See ForeignKey |

**Validation Requirements**:
- `Name` must be unique within table (case-sensitive)
- `Type` must be from supported types list (see Supported Types section)
- If `PrimaryKey: true`, exactly one per table
- If `AutoIncrement: true`, must be integer type (int, bigint, smallint, tinyint)
- If `ForeignKey` is set, must reference existing table

**Generator Parameters Examples**:
```go
// Normal distribution (credit scores)
GeneratorParams: map[string]interface{}{
    "mean":    680,
    "std_dev": 80,
    "min":     300,
    "max":     850,
}

// Lognormal distribution (loan amounts)
GeneratorParams: map[string]interface{}{
    "median": 15000,
    "min":    1000,
    "max":    50000,
}

// Weighted distribution (loan status)
GeneratorParams: map[string]interface{}{
    "values": []map[string]interface{}{
        {"value": "active", "weight": 0.7},
        {"value": "paid", "weight": 0.25},
        {"value": "delinquent", "weight": 0.05},
    },
}
```

---

### ForeignKey

**Purpose**: Foreign key constraint for referential integrity
**File**: `pkg/schema/types.go`
**Usage**: Schema validation, SQL constraint creation, data generation ordering

```go
type ForeignKey struct {
    Table    string `json:"table"`
    Column   string `json:"column"`
    OnDelete string `json:"on_delete,omitempty"`
    OnUpdate string `json:"on_update,omitempty"`
}
```

**Fields**:

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `Table` | string | Yes | Referenced table name | `"borrowers"` |
| `Column` | string | Yes | Referenced column name | `"id"` |
| `OnDelete` | string | No | Action on delete | `"CASCADE"`, `"SET NULL"`, `"RESTRICT"` |
| `OnUpdate` | string | No | Action on update | `"CASCADE"`, `"SET NULL"`, `"RESTRICT"` |

**Validation Requirements**:
- `Table` must reference existing table in schema (MVP validation)
- `Column` must be non-empty (column existence validation in Phase 2)
- `OnDelete` if set, must be: "CASCADE", "SET NULL", "RESTRICT", or empty (database default)
- `OnUpdate` if set, must be: "CASCADE", "SET NULL", "RESTRICT", or empty (database default)

**Usage in Column**:
```go
Column{
    Name: "borrower_id",
    Type: "int",
    ForeignKey: &ForeignKey{
        Table:    "borrowers",
        Column:   "id",
        OnDelete: "CASCADE",
        OnUpdate: "CASCADE",
    },
}
```

**Referential Integrity Actions**:
- **CASCADE**: Delete/update referenced row → delete/update this row
- **SET NULL**: Delete/update referenced row → set this column to NULL (requires Nullable: true)
- **RESTRICT**: Prevent delete/update of referenced row if this row exists
- **Empty**: Use database default (typically RESTRICT)

---

### Index

**Purpose**: Table index for query performance optimization
**File**: `pkg/schema/types.go`
**Usage**: SQL index creation (optional performance optimization)

```go
type Index struct {
    Name    string   `json:"name"`
    Columns []string `json:"columns"`
    Unique  bool     `json:"unique,omitempty"`
}
```

**Fields**:

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `Name` | string | Yes | Index name | `"idx_borrower_email"` |
| `Columns` | []string | Yes | Indexed column names | `["email"]`, `["last_name", "first_name"]` |
| `Unique` | bool | No | Unique index constraint | `true` |

**Validation Requirements** (MVP: minimal validation):
- `Name` must be non-empty string
- `Columns` must have at least one column name
- Columns must reference existing columns in table (Phase 2 validation)

**Index Types**:
- **Single-column**: `["email"]` - Index on one column
- **Composite**: `["last_name", "first_name"]` - Multi-column index (order matters)
- **Unique**: `Unique: true` - Enforces uniqueness constraint

**Usage in Table**:
```go
Table{
    Name: "borrowers",
    Indexes: []Index{
        {
            Name:    "idx_borrower_email",
            Columns: []string{"email"},
            Unique:  true,
        },
        {
            Name:    "idx_borrower_name",
            Columns: []string{"last_name", "first_name"},
        },
    },
}
```

---

### Relationship

**Purpose**: Explicit relationship documentation for human readers
**File**: `pkg/schema/types.go`
**Usage**: Schema documentation, relationship visualization, future tooling

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

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `FromTable` | string | Yes | Source table name | `"loans"` |
| `FromColumn` | string | Yes | Source column name | `"borrower_id"` |
| `ToTable` | string | Yes | Target table name | `"borrowers"` |
| `ToColumn` | string | Yes | Target column name | `"id"` |
| `RelationshipType` | string | Yes | Relationship type | `"many_to_one"`, `"one_to_many"` |
| `Description` | string | Yes | Human-readable description | `"Each loan belongs to one borrower"` |

**Validation Requirements** (MVP: minimal validation):
- All fields must be non-empty strings
- `FromTable` and `ToTable` must reference existing tables (Phase 2)
- `RelationshipType` must be: "one_to_one", "one_to_many", "many_to_one", "many_to_many" (Phase 2)

**Relationship Types**:
- **one_to_one**: User → Profile (each user has one profile)
- **one_to_many**: Borrower → Loans (one borrower, many loans)
- **many_to_one**: Loans → Borrower (many loans, one borrower) - inverse of one_to_many
- **many_to_many**: Students → Courses (via junction table)

**Dual Representation** (inline + explicit):
- **Inline**: `Column.ForeignKey` object (what parser uses for validation)
- **Explicit**: `Schema.Relationships` array (documentation for humans and tools)
- **Rationale**: Code uses inline, documentation uses explicit (see F007 spec)

---

### ValidationRule

**Purpose**: Optional soft validation rules for data quality
**File**: `pkg/schema/types.go`
**Usage**: Documentation, future data quality checks (Phase 2)

```go
type ValidationRule struct {
    Rule        string `json:"rule"`
    Description string `json:"description"`
    Enforcement string `json:"enforcement"`
}
```

**Fields**:

| Field | Type | Required | Description | Example |
|-------|------|----------|-------------|---------|
| `Rule` | string | Yes | Rule identifier | `"credit_score_range"` |
| `Description` | string | Yes | Rule description | `"Credit scores must be between 300 and 850"` |
| `Enforcement` | string | Yes | Enforcement level | `"soft"` (warn), `"hard"` (error) |

**Validation Requirements** (MVP: minimal validation):
- All fields must be non-empty strings
- `Enforcement` must be "soft" or "hard" (Phase 2 validation)

**Enforcement Levels**:
- **soft**: Warning only (data generation proceeds)
- **hard**: Error (data generation fails)

**Usage** (future enhancement):
```go
Schema{
    ValidationRules: []ValidationRule{
        {
            Rule:        "credit_score_range",
            Description: "Credit scores must be between 300 and 850",
            Enforcement: "hard",
        },
    },
}
```

---

## Supported Data Types

**Source**: F007 Schema Specification (Supported Data Types section)

### Type Categories

Parser validates column types against this list (case-insensitive, prefix matching):

**Integer Types**:
- `int` - Standard 32-bit integer
- `bigint` - 64-bit integer
- `smallint` - 16-bit integer
- `tinyint` - 8-bit integer (MySQL: 0-255, Postgres: -128 to 127)

**Decimal Types**:
- `decimal` or `decimal(p,s)` - Fixed-point decimal (e.g., `decimal(10,2)` for money)
- `float` - 32-bit floating point
- `double` - 64-bit floating point

**String Types**:
- `varchar(n)` - Variable-length string (e.g., `varchar(255)`)
- `text` - Unlimited length text
- `char(n)` - Fixed-length string (e.g., `char(50)`)

**Date/Time Types**:
- `date` - Date only (YYYY-MM-DD)
- `datetime` - Date and time (YYYY-MM-DD HH:MM:SS)
- `timestamp` - Unix timestamp or datetime with timezone

**Boolean Type**:
- `boolean` - PostgreSQL native boolean
- `bit` - MySQL boolean (BIT(1))

**JSON Types**:
- `json` - JSON data (both MySQL and PostgreSQL)
- `jsonb` - Binary JSON (PostgreSQL only, parser accepts for MySQL schemas but doesn't validate compatibility)

**Enum Types**:
- `enum('val1','val2',...)` - Enumerated values (e.g., `enum('active','paid','delinquent')`)

### Validation Algorithm

```go
func ValidateDataType(dataType string) error {
    normalized := strings.ToLower(dataType)

    validTypes := []string{
        "int", "bigint", "smallint", "tinyint",
        "decimal", "float", "double",
        "varchar", "text", "char",
        "date", "datetime", "timestamp",
        "boolean", "bit",
        "json", "jsonb",
        "enum",
    }

    for _, validType := range validTypes {
        if strings.HasPrefix(normalized, validType) {
            return nil // Valid type found
        }
    }

    return fmt.Errorf("unsupported data type: %s", dataType)
}
```

**Examples**:
- `"VARCHAR(255)"` → matches `"varchar"` ✅
- `"DECIMAL(10,2)"` → matches `"decimal"` ✅
- `"int"` → matches `"int"` ✅
- `"foobar"` → no match ❌ → error: "unsupported data type: foobar"

---

## Struct Relationships

Visual hierarchy of nested structs:

```
Schema
├── Metadata: SchemaMetadata
│   ├── TotalRecords: int
│   ├── Industry: string
│   └── Tags: []string
├── Tables: []Table
│   ├── Name: string
│   ├── Description: string
│   ├── RecordCount: int
│   ├── Columns: []Column
│   │   ├── Name: string
│   │   ├── Type: string
│   │   ├── PrimaryKey: bool
│   │   ├── Generator: string
│   │   ├── GeneratorParams: map[string]interface{}
│   │   └── ForeignKey: *ForeignKey
│   │       ├── Table: string
│   │       ├── Column: string
│   │       ├── OnDelete: string
│   │       └── OnUpdate: string
│   └── Indexes: []Index
│       ├── Name: string
│       ├── Columns: []string
│       └── Unique: bool
├── Relationships: []Relationship
│   ├── FromTable: string
│   ├── FromColumn: string
│   ├── ToTable: string
│   ├── ToColumn: string
│   ├── RelationshipType: string
│   └── Description: string
├── GenerationOrder: []string
└── ValidationRules: []ValidationRule
    ├── Rule: string
    ├── Description: string
    └── Enforcement: string
```

---

## Usage Examples

### Parsing a Schema

```go
import (
    "github.com/jbeausoleil/sourcebox/pkg/schema"
)

// Load from file path
schema, err := schema.LoadSchema("/path/to/fintech-loans.json")
if err != nil {
    log.Fatalf("Failed to load schema: %v", err)
}

// Or parse from io.Reader
file, _ := os.Open("schema.json")
defer file.Close()
schema, err := schema.ParseSchema(file)
```

### Accessing Schema Data

```go
// Top-level metadata
fmt.Printf("Schema: %s v%s\n", schema.Name, schema.Version)
fmt.Printf("Industry: %s\n", schema.Metadata.Industry)
fmt.Printf("Total records: %d\n", schema.Metadata.TotalRecords)

// Iterate tables
for _, table := range schema.Tables {
    fmt.Printf("Table: %s (%d records)\n", table.Name, table.RecordCount)

    // Iterate columns
    for _, col := range table.Columns {
        fmt.Printf("  Column: %s (%s)\n", col.Name, col.Type)

        // Check for foreign key
        if col.ForeignKey != nil {
            fmt.Printf("    Foreign key: %s.%s\n",
                col.ForeignKey.Table, col.ForeignKey.Column)
        }
    }
}

// Check generation order
fmt.Printf("Generation order: %v\n", schema.GenerationOrder)
```

### Type Checking

```go
// Check if column is primary key
if col.PrimaryKey {
    fmt.Println("This is a primary key column")
}

// Check if column has foreign key
if col.ForeignKey != nil {
    fmt.Printf("References: %s.%s\n",
        col.ForeignKey.Table, col.ForeignKey.Column)
}

// Check generator parameters
if params, ok := col.GeneratorParams["mean"]; ok {
    fmt.Printf("Mean: %v\n", params)
}
```

---

## Validation Summary

### Schema-Level Validation

✅ Required fields: `name`, `tables`, `database_type`, `generation_order`
✅ At least one table exists
✅ `database_type` contains valid values ("mysql", "postgres")
✅ No duplicate table names

### Table-Level Validation

✅ Required fields: `name`, `record_count`, `columns`
✅ At least one column exists
✅ Exactly one primary key column
✅ `record_count` > 0
✅ No duplicate column names within table

### Column-Level Validation

✅ Required fields: `name`, `type`
✅ `type` is from supported types list
✅ If `auto_increment: true`, type must be integer
✅ If `foreign_key` set, referenced table must exist

### Foreign Key Validation

✅ `foreign_key.table` references existing table (MVP)
✅ `on_delete` and `on_update` are valid actions or empty
⏸️ Column existence validation (Phase 2)
⏸️ Primary key reference validation (Phase 2)
⏸️ Type compatibility validation (Phase 2)

### Generation Order Validation

✅ `generation_order` is non-empty
✅ All tables included exactly once
✅ No duplicate table names
✅ No missing tables
⏸️ Circular dependency detection (Phase 2)

---

## Future Enhancements (Phase 2)

### Column-Level Foreign Key Validation
- Validate `foreign_key.column` exists in referenced table
- Validate referenced column is primary key or unique
- Validate type compatibility (foreign key type = referenced column type)

### Circular Dependency Detection
- Build dependency graph from foreign keys
- Detect cycles using depth-first search
- Suggest correct generation_order

### Automatic Generation Order
- Compute topological sort from foreign key relationships
- Generate correct generation_order automatically
- Tool: `sourcebox validate --fix` to auto-correct order

### Enhanced Error Messages
- Add "Fix:" suggestions to validation errors
- Show example of correct format
- Provide links to documentation

### Type Compatibility Validation
- Check `jsonb` only used with PostgreSQL schemas
- Validate `tinyint` usage (MySQL vs PostgreSQL semantics)
- Warn about database-specific types

---

## Constitutional Alignment

✅ **Boring Tech Wins**: Standard library `encoding/json`, no custom unmarshaling
✅ **Simple > Complex**: Flat struct hierarchy, clear field names
✅ **Developer-First**: Explicit required/optional via pointer/value distinction
✅ **Speed > Features**: Simple structs enable fast parsing (<100ms)
✅ **TDD Required**: All structs testable via JSON unmarshaling tests

---

## Related Documentation

- **F007 Schema Specification**: `/schemas/schema-spec.md` - JSON format definition
- **Research Decisions**: `research.md` - Technical decision rationale
- **Quickstart Guide**: `quickstart.md` - TDD workflow and verification
- **Parser Implementation**: `pkg/schema/parser.go` - Parsing and validation logic
- **Parser Tests**: `pkg/schema/parser_test.go` - TDD test suite
