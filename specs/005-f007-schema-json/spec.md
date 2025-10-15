# Feature Specification: Schema JSON Format Specification

**Feature Branch**: `005-f007-schema-json`
**Created**: 2025-10-15
**Status**: Draft
**Input**: User description: "F007 - Schema JSON Format Specification"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Understanding Schema Format (Priority: P1)

As a developer new to SourceBox, I want to read clear documentation about the schema format so I can understand how to create schemas without needing to reverse-engineer existing examples.

**Why this priority**: Without schema format documentation, developers cannot contribute new schemas or modify existing ones. This is foundational to the entire schema ecosystem and enables all other schema-related work.

**Independent Test**: Can be fully tested by reviewing the specification document (`schemas/schema-spec.md`) and verifying that it comprehensively explains all schema format elements (tables, columns, generators, relationships) with clear examples. Delivers immediate value by enabling developers to understand the system.

**Acceptance Scenarios**:

1. **Given** a developer has no prior knowledge of SourceBox, **When** they read `schemas/schema-spec.md`, **Then** they can identify all required fields for creating a valid schema
2. **Given** the specification document exists, **When** a developer wants to add a new table to a schema, **Then** they can find clear instructions and examples in the specification
3. **Given** the specification document exists, **When** a developer needs to define a foreign key relationship, **Then** they can follow the documented format without trial and error
4. **Given** a developer reads the specification, **When** they want to use a custom generator (like `credit_score`), **Then** they understand how to specify generator parameters correctly

---

### User Story 2 - Validating Schema Correctness (Priority: P2)

As a schema author, I want to know the validation rules that determine whether my schema is valid so I can create schemas that will work correctly without runtime errors.

**Why this priority**: Clear validation rules prevent developers from creating invalid schemas that fail at runtime. This reduces frustration and enables self-service schema creation.

**Independent Test**: Can be fully tested by reviewing the validation rules section and attempting to create both valid and invalid schemas according to the rules. Delivers value by enabling schema authors to self-validate their work before submission.

**Acceptance Scenarios**:

1. **Given** validation rules are documented, **When** a developer creates a schema missing a primary key, **Then** they can identify this violation by reading the validation rules
2. **Given** validation rules are documented, **When** a developer references a non-existent table in a foreign key, **Then** they understand why this is invalid before testing
3. **Given** validation rules exist, **When** a developer uses an invalid generator name, **Then** the rules explain what constitutes a valid generator name
4. **Given** the validation rules section, **When** a developer wants to ensure referential integrity, **Then** they can verify their `generation_order` includes all tables

---

### User Story 3 - Using Built-in Generators (Priority: P2)

As a schema author creating realistic test data, I want to understand what built-in generators are available and how to configure them so I can generate appropriate data for my domain without writing custom code.

**Why this priority**: Built-in generators enable rapid schema creation for common data types. Understanding their capabilities and parameters allows developers to leverage existing functionality before creating custom generators.

**Independent Test**: Can be fully tested by reviewing the generator documentation and creating schemas that use various built-in generators with different parameter configurations. Delivers value by showing what's possible with zero custom code.

**Acceptance Scenarios**:

1. **Given** the built-in generators list exists, **When** a developer needs to generate personal names, **Then** they can identify `first_name`, `last_name`, and `full_name` generators
2. **Given** generator documentation exists, **When** a developer wants normally-distributed numeric values, **Then** they understand how to use the `normal` distribution with `mean` and `std_dev` parameters
3. **Given** timestamp generators are documented, **When** a developer needs historical dates, **Then** they can configure `timestamp_past` with appropriate `days_ago_min` and `days_ago_max` parameters
4. **Given** the generator parameters section, **When** a developer needs weighted values (like loan status), **Then** they understand how to specify values with probability weights

---

### User Story 4 - Creating Example Schemas (Priority: P1)

As a schema author, I want to see a complete, working example schema that demonstrates all major features so I can use it as a reference when creating my own schemas.

**Why this priority**: A comprehensive example schema serves as both documentation and validation of the format. It provides a concrete reference that developers can copy, modify, and learn from.

**Independent Test**: Can be fully tested by examining `schemas/example-schema.json` and verifying it contains examples of all documented features (multiple tables, various column types, generators with parameters, foreign keys, indexes, relationships). Delivers value by providing immediate, copyable patterns.

**Acceptance Scenarios**:

1. **Given** the example schema file exists, **When** a developer wants to create a multi-table schema, **Then** they can see how tables relate through foreign keys
2. **Given** `example-schema.json` is provided, **When** a developer needs to use custom generators, **Then** they can see examples with various distribution types (normal, lognormal, weighted)
3. **Given** the example demonstrates all features, **When** a developer wants to add indexes, **Then** they can copy the index syntax from the example
4. **Given** the example schema exists, **When** a developer needs to specify generation order, **Then** they understand the purpose and format by seeing it in context

---

### User Story 5 - Versioning Schemas Over Time (Priority: P3)

As a schema maintainer, I want to understand the versioning strategy so I can evolve schemas without breaking existing users or implementations.

**Why this priority**: Schema versioning enables long-term maintenance and evolution. While less critical than initial creation, it's essential for sustainable schema development.

**Independent Test**: Can be fully tested by reviewing the versioning section and understanding when to increment major, minor, or patch versions. Delivers value by establishing clear version change guidelines.

**Acceptance Scenarios**:

1. **Given** versioning strategy is documented, **When** a developer adds a new column to an existing table, **Then** they understand this requires a minor version bump (backward compatible)
2. **Given** the versioning section exists, **When** a developer removes a table from a schema, **Then** they know this requires a major version bump (breaking change)
3. **Given** versioning guidelines exist, **When** a developer improves data distribution quality, **Then** they understand this is a patch-level change
4. **Given** semantic versioning is explained, **When** a new schema version is released, **Then** downstream tools can determine compatibility based on version numbers

---

### Edge Cases

- What happens when a schema defines a `generation_order` that includes tables not defined in the schema?
- How does the system handle circular foreign key relationships between tables?
- What validation occurs when generator parameters conflict with column type constraints (e.g., `varchar(10)` with a generator that produces 50-character strings)?
- How are schema name conflicts resolved when multiple schemas share the same name but different versions?
- What happens when a foreign key references a column that isn't a primary key?
- How does the system validate that `on_delete` and `on_update` constraints are supported by both MySQL and Postgres?
- What validation ensures that custom generator names don't conflict with built-in generator names?
- How are inconsistent `total_records` metadata values (sum of table record counts doesn't match metadata) handled?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Specification document MUST define all required fields for a valid schema (schema_version, name, description, tables)
- **FR-002**: Specification document MUST define all optional fields for a schema (author, database_type, metadata, relationships, validation_rules)
- **FR-003**: Specification document MUST document all supported data types for MySQL and PostgreSQL common types
- **FR-004**: Specification document MUST list all built-in generator types organized by category (Personal Data, Company Data, Date/Time, Numeric)
- **FR-005**: Specification document MUST explain generator parameter formats for all distribution types (uniform, normal, lognormal, weighted)
- **FR-006**: Specification document MUST define the structure for foreign key relationships including on_delete and on_update constraints
- **FR-007**: Specification document MUST explain semantic versioning strategy for schemas (major, minor, patch)
- **FR-008**: Example schema file MUST demonstrate all major schema features (multiple tables, foreign keys, custom generators, indexes, relationships)
- **FR-009**: Specification document MUST include validation rules that define what constitutes a valid schema
- **FR-010**: Validation rules MUST require at least one table per schema
- **FR-011**: Validation rules MUST require at least one column per table
- **FR-012**: Validation rules MUST require each table to have a primary key
- **FR-013**: Validation rules MUST require foreign keys to reference existing tables and columns
- **FR-014**: Validation rules MUST require generator names to be either built-in or custom (defined in schema)
- **FR-015**: Validation rules MUST require generator parameters to match generator requirements
- **FR-016**: Validation rules MUST require `generation_order` to include all defined tables
- **FR-017**: Validation rules MUST require record counts to be positive integers
- **FR-018**: Example schema MUST use realistic fintech domain data (borrowers, loans) to demonstrate verticalized data generation
- **FR-019**: Specification document MUST explain how custom generators differ from built-in generators
- **FR-020**: Specification document MUST document index definition format including unique and composite indexes
- **FR-021**: Specification document MUST explain the purpose of the `generation_order` field
- **FR-022**: Example schema MUST be valid according to all documented validation rules

### Key Entities

- **Schema**: Represents a complete database schema definition with metadata, tables, and relationships
  - Includes name, version, description, author, database compatibility
  - Contains one or more table definitions
  - Optionally includes relationships and validation rules

- **Table**: Represents a database table within a schema
  - Has a name, description, and record count target
  - Contains one or more column definitions
  - May include index definitions
  - Must have exactly one primary key

- **Column**: Represents a column within a table
  - Has a name, data type, and nullability
  - May be a primary key or have a foreign key constraint
  - May specify a data generator with parameters
  - May have uniqueness constraint or default value

- **Generator**: Represents a data generation strategy for a column
  - Built-in generators (name, email, timestamp) or custom generators (credit_score, loan_amount)
  - Accepts parameters that control distribution and constraints
  - Parameters include distribution type (uniform, normal, lognormal, weighted) and specific values (mean, min, max, weights)

- **Relationship**: Represents a foreign key relationship between tables
  - Defines source table/column and target table/column
  - Specifies relationship type (many_to_one, one_to_many, many_to_many)
  - Documents business meaning of the relationship

- **Validation Rule**: Represents a constraint or quality check for generated data
  - Can be "soft" (warning) or "hard" (enforced)
  - Documents expected correlations or patterns in realistic data

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Developers with no prior SourceBox knowledge can create a valid single-table schema within 15 minutes of reading the specification
- **SC-002**: The example schema file passes validation according to all documented validation rules (100% compliance)
- **SC-003**: Documentation answers all common schema authoring questions without requiring external clarification (measured by zero [NEEDS CLARIFICATION] markers remaining in spec)
- **SC-004**: Specification document covers 100% of schema format elements shown in the example schema
- **SC-005**: Schema format supports both MySQL and PostgreSQL without database-specific features that prevent cross-compatibility
- **SC-006**: Developers can identify which built-in generators to use for common data types (names, emails, dates) within 2 minutes of reading the generator documentation
- **SC-007**: Schema versioning strategy allows schema evolution without breaking existing implementations (backward compatibility for minor versions)
- **SC-008**: Validation rules enable schema authors to self-diagnose 90% of schema errors before runtime testing

## Assumptions *(include if making informed guesses)*

1. **Generator Parameter Formats**: Assuming JSON-native types (numbers, strings, objects, arrays) are sufficient for generator parameters; no need for custom expression languages
2. **Database Compatibility**: Assuming MySQL and PostgreSQL are the only supported databases initially; other databases (SQLite, SQL Server) deferred to future versions
3. **Custom Generator Definition**: Assuming custom generators are defined externally (in Go code) rather than inline in JSON schemas; schema only references generator names
4. **Schema Validation Timing**: Assuming schema validation occurs at load time (when parsing JSON) rather than at generation time; invalid schemas fail fast
5. **Index Complexity**: Assuming basic index support (single/multi-column, unique/non-unique) is sufficient; advanced features (partial indexes, expression indexes) not required
6. **Relationship Documentation**: Assuming relationships section is documentation-only (describes foreign keys) rather than generative (doesn't create additional constraints beyond foreign_key field)
7. **Validation Rule Enforcement**: Assuming "soft" validation rules are warnings only; "hard" rules would require implementation changes beyond this specification
8. **Schema Embedding**: Assuming schemas are embedded in the binary at compile time (per project context); no runtime schema file loading needed for this specification

## Dependencies *(include if feature depends on other work)*

### Upstream Dependencies
- None (this is a design/specification feature with no code dependencies)

### Downstream Dependencies
- **F008 (Schema Parser)**: Will implement and validate this specification in Go code
- **F010 (Fintech Schema)**: Will use this format to define fintech loan data schema
- **F016 (Healthcare Schema)**: Will use this format to define healthcare patient data schema
- **F018 (Retail Schema)**: Will use this format to define retail product/order data schema

### Blocking Relationships
- This feature (F007) must be completed before F008 can begin implementation
- F008 must validate that this specification is implementable and complete
- Schema authors cannot create new schemas until this specification is documented and example is provided

## Scope *(include to clarify what's included/excluded)*

### In Scope
- JSON schema format specification for SourceBox data generation
- Comprehensive documentation of all schema elements (tables, columns, generators, relationships)
- Built-in generator types and parameter documentation
- Custom generator reference documentation (names and parameters, not implementation)
- Foreign key relationship definition format
- Schema versioning strategy (semantic versioning)
- Validation rules for schema correctness
- Example schema demonstrating all major features (fintech domain)
- Documentation answering common schema authoring questions

### Out of Scope
- Implementation of schema parser in Go (F008)
- Implementation of data generators (F009, F011-F020)
- Creation of production schemas beyond the example (F010, F016, F018)
- Schema validation tooling or linting (future enhancement)
- Visual schema designer or GUI tools
- Schema migration or upgrade tooling
- Performance optimization of schema parsing (specification only)
- Database-specific features beyond MySQL/PostgreSQL common subset
- Implementation of custom generators (specification documents interface only)
- Test suite for schema validation (implementation concern, not specification)

### Boundary Clarifications
- This feature produces documentation and an example file, not executable code
- The specification should be implementable by F008 without ambiguity
- Example schema should be realistic but doesn't need to match production data distributions exactly
- Generator documentation describes interface only, not implementation algorithm details
- Validation rules describe what makes a schema valid, not how to implement validation