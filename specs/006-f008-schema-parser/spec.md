# Feature Specification: Schema Parser & Validator

**Feature Branch**: `006-f008-schema-parser`
**Created**: 2025-10-15
**Status**: Draft
**Input**: User description: "F008 - Schema Parser & Validator"

## User Scenarios & Testing

### User Story 1 - Load Valid Schema (Priority: P1)

A developer wants to load a schema JSON file to use for data generation. They provide a valid schema file path, and the system successfully parses it into usable data structures.

**Why this priority**: This is the core happy path - without the ability to load valid schemas, no other functionality works. This delivers immediate value by enabling developers to use schema files.

**Independent Test**: Can be fully tested by providing a valid schema JSON file with all required fields (name, tables, columns, database_type, generation_order) and verifying it loads without errors and returns a populated schema object with correct field values.

**Acceptance Scenarios**:

1. **Given** a valid schema JSON file exists with all required fields, **When** a developer loads the schema using the parser, **Then** the system returns a schema object with all fields correctly populated
2. **Given** a schema file with multiple tables and foreign key relationships, **When** the parser loads the file, **Then** all table definitions and relationships are correctly preserved in the schema object
3. **Given** a schema file with optional fields (metadata, validation_rules), **When** loaded, **Then** both required and optional fields are correctly parsed

---

### User Story 2 - Detect Missing Required Fields (Priority: P1)

A developer provides a schema file that is missing required fields (e.g., no table name, no database_type). The system detects this during parsing and provides a clear error message indicating exactly what is missing and where.

**Why this priority**: Catching validation errors early prevents runtime failures and reduces debugging time. This is critical for developer experience - without clear errors, developers waste hours troubleshooting.

**Independent Test**: Can be tested by providing schema files with various missing fields (name, tables, database_type, generation_order) and verifying that each produces a specific error message identifying the missing field.

**Acceptance Scenarios**:

1. **Given** a schema file missing the "name" field, **When** parsed, **Then** system returns error: "schema name is required"
2. **Given** a schema file with no tables defined, **When** parsed, **Then** system returns error: "schema must have at least one table"
3. **Given** a table definition missing a primary key, **When** parsed, **Then** system returns error: "table [name] must have a primary key"
4. **Given** a schema missing database_type, **When** parsed, **Then** system returns error: "database_type is required"

---

### User Story 3 - Validate Foreign Key Integrity (Priority: P2)

A developer creates a schema where one table references another table via foreign key. The system validates that the referenced table exists and the referenced column is valid.

**Why this priority**: Foreign key validation prevents data generation errors that would only surface at runtime. This catches schema design errors early in development.

**Independent Test**: Can be tested by providing schema files with various foreign key scenarios (valid references, non-existent tables, non-existent columns) and verifying appropriate validation results.

**Acceptance Scenarios**:

1. **Given** a schema where table A has a foreign key to table B, and table B exists, **When** parsed, **Then** validation succeeds
2. **Given** a foreign key reference to a non-existent table, **When** parsed, **Then** system returns error: "table [X]: column [Y] references non-existent table [Z]"
3. **Given** a foreign key with invalid integrity actions (on_delete, on_update), **When** parsed, **Then** system returns error: "invalid on_delete action: [action]"

---

### User Story 4 - Validate Data Types (Priority: P2)

A developer specifies column data types in their schema. The system validates that each data type is supported by the target database (MySQL/PostgreSQL common subset).

**Why this priority**: Data type validation prevents schema generation errors when creating database tables. This ensures schemas are portable across supported databases.

**Independent Test**: Can be tested by providing schemas with various data types (valid: int, varchar, timestamp; invalid: unsupported types) and verifying validation results.

**Acceptance Scenarios**:

1. **Given** columns using standard types (int, varchar, datetime), **When** parsed, **Then** validation succeeds
2. **Given** a column with an unsupported data type, **When** parsed, **Then** system returns error: "unsupported data type: [type]"
3. **Given** columns using database-specific types (jsonb for PostgreSQL), **When** parsed, **Then** validation succeeds

---

### User Story 5 - Validate Generation Order (Priority: P2)

A developer defines the order in which tables should be generated. The system validates that all tables are included in the generation order and no duplicates exist.

**Why this priority**: Generation order is critical for foreign key integrity - parent tables must be generated before child tables. Invalid order would cause data generation to fail.

**Independent Test**: Can be tested by providing schemas with various generation_order configurations (complete, missing tables, duplicates) and verifying validation results.

**Acceptance Scenarios**:

1. **Given** generation_order includes all tables exactly once, **When** parsed, **Then** validation succeeds
2. **Given** generation_order missing a table, **When** parsed, **Then** system returns error: "table [name] is not in generation_order"
3. **Given** generation_order with duplicate table names, **When** parsed, **Then** system returns error: "generation_order contains duplicate table: [name]"
4. **Given** generation_order referencing a non-existent table, **When** parsed, **Then** system returns error: "generation_order references non-existent table: [name]"

---

### User Story 6 - Detect Duplicate Names (Priority: P3)

A developer accidentally uses the same name for multiple tables or columns. The system detects these duplicates and provides clear error messages indicating which names are duplicated.

**Why this priority**: Duplicate names would cause database errors. While less common than other validation errors, catching this early prevents subtle bugs.

**Independent Test**: Can be tested by providing schemas with duplicate table names or duplicate column names within a table and verifying appropriate errors are returned.

**Acceptance Scenarios**:

1. **Given** two tables with the same name, **When** parsed, **Then** system returns error: "duplicate table name: [name]"
2. **Given** a table with duplicate column names, **When** parsed, **Then** system returns error: "duplicate column name: [name]"

---

### Edge Cases

- What happens when a schema file is empty or contains only whitespace?
- What happens when a schema file contains malformed JSON (syntax errors)?
- What happens when a schema file contains unknown/extra fields not in the specification?
- What happens when numeric fields contain non-numeric values (e.g., record_count: "abc")?
- What happens when required boolean fields are missing (e.g., primary_key not specified)?
- What happens when a file path doesn't exist or cannot be read?
- What happens when a foreign key references a column that doesn't exist in the target table?
- What happens when a table has zero or negative record_count?

## Requirements

### Functional Requirements

- **FR-001**: System MUST load schema definitions from file paths on the filesystem
- **FR-002**: System MUST parse JSON content into structured data objects
- **FR-003**: System MUST validate that all required top-level fields are present (schema_version, name, description, author, version, database_type, metadata, tables, generation_order)
- **FR-004**: System MUST validate that each table has required fields (name, description, record_count, columns)
- **FR-005**: System MUST validate that each column has required fields (name, type)
- **FR-006**: System MUST verify database_type contains only supported values ("mysql" or "postgres")
- **FR-007**: System MUST verify each table has exactly one primary key column
- **FR-008**: System MUST validate that column data types are from the supported set (int, bigint, smallint, tinyint, decimal, float, double, varchar, text, char, date, datetime, timestamp, boolean, bit, json, jsonb, enum)
- **FR-009**: System MUST validate that foreign key references point to existing tables
- **FR-010**: System MUST validate that foreign key integrity actions (on_delete, on_update) are valid (CASCADE, SET NULL, RESTRICT, or empty)
- **FR-011**: System MUST validate that generation_order includes all tables exactly once with no duplicates
- **FR-012**: System MUST detect and reject duplicate table names
- **FR-013**: System MUST detect and reject duplicate column names within the same table
- **FR-014**: System MUST reject schemas where record_count is zero or negative
- **FR-015**: System MUST reject malformed JSON with clear error messages
- **FR-016**: System MUST reject JSON containing unknown fields (strict parsing)
- **FR-017**: System MUST provide error messages that include the specific location of the error (e.g., "table 2 (loans): column 5 (borrower_id): [error]")
- **FR-018**: System MUST support loading schemas from in-memory readers (not just files)

### Key Entities

- **Schema**: The top-level container representing a complete database schema definition, including metadata (name, version, author), target databases, tables, relationships, and generation order
- **Table**: A database table definition including its name, description, target record count, columns, and optional indexes
- **Column**: A table column definition including name, data type, constraints (primary key, nullable, unique), default values, data generator configuration, and optional foreign key relationships
- **ForeignKey**: A relationship constraint defining how a column references another table's column, including referential integrity actions (CASCADE, SET NULL, RESTRICT)
- **Relationship**: Documentation of table relationships including source/target tables and columns, relationship type (one-to-one, one-to-many, many-to-one, many-to-many), and description
- **ValidationRule**: Optional rules for data validation including the rule definition, description, and enforcement level (soft or hard)

## Success Criteria

### Measurable Outcomes

- **SC-001**: Developers can load a valid schema file and receive a parsed schema object in under 100 milliseconds for schemas with up to 10 tables
- **SC-002**: When provided with an invalid schema, developers receive an error message that includes the specific field or location causing the error within 1 second
- **SC-003**: 100% of validation rules specified in the schema specification (F007) are enforced by the parser
- **SC-004**: Developers can identify and fix schema errors without needing to consult documentation, based solely on error messages
- **SC-005**: All edge cases (empty files, malformed JSON, missing fields, invalid references) are handled gracefully with clear error messages
- **SC-006**: Parser successfully validates all example schemas from the schema specification (F007) without false positives or negatives

## Assumptions

- Schema files use UTF-8 encoding
- Schema files are reasonably sized (< 10MB) for typical use cases
- File system has read permissions for schema file paths
- JSON structure follows the schema specification defined in F007
- Error messages are displayed to developers in a terminal or log output
- No concurrent schema loading is required (single-threaded access is sufficient)
- Schema validation is synchronous (blocking operation)
