# Implementation Planning Prompt: F007 - Schema JSON Format Specification

## Feature Metadata
- **Feature ID**: F007
- **Name**: Schema JSON Format Specification
- **Feature Branch**: `005-f007-schema-json`
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (2 days)
- **Dependencies**: None (design work)
- **Spec Location**: `.specify/prompts/specify/mvp/f007-schema-json-format-specification.md`

## Constitutional Alignment

### Core Principles Verification
- ✅ **Verticalized > Generic**: Schema format MUST support industry-specific data generators and realistic distributions (custom generators for fintech, healthcare, retail)
- ✅ **Speed > Features**: N/A (design work) - However, format design enables fast parsing in F008 (JSON parsing is fast in Go)
- ✅ **Local-First, Cloud Optional**: N/A (specification document)
- ✅ **Boring Tech Wins**: JSON format is proven, standard, boring - maximum compatibility, extensive tooling, well-understood
- ✅ **Open Source Forever**: Schema format will be MIT licensed, enables community schema contributions
- ✅ **Developer-First Design**: Clear specification document enables developers to contribute schemas, format is human-readable and machine-parseable
- ✅ **Ship Fast, Validate Early**: Design-only feature (2 days), no implementation delay, spec validated by creating example schema

### Technical Constraints Verification
- ✅ **Performance**: JSON parsing is fast in Go standard library, schema loading happens once at startup
- ✅ **Distribution**: Schemas will be embedded in binary (no runtime file reads), format supports this
- ✅ **Database Support**: Format must support both MySQL and PostgreSQL column types, relationships, constraints
- ✅ **Cost**: $0 (design work, no infrastructure)
- ✅ **Code Quality**: N/A (documentation feature) - However, specification must be clear, unambiguous, implementable
- ✅ **License**: Schema format is MIT, schemas themselves are MIT (community can contribute)
- ✅ **Platform Support**: JSON is cross-platform, format design is platform-agnostic

### Legal Constraints Verification (CRITICAL)
- ✅ **Independent Development**: JSON schema design using public best practices only
- ✅ **No Employer References**: N/A (technical specification)
- ✅ **Public Information Only**: All patterns from public JSON schema documentation, Go JSON parsing docs, database documentation
- ✅ **Open Source Protection**: Format will be MIT licensed
- ✅ **Illustrative Examples Only**: Example schemas use generic fintech terminology (loans, credit scores) from public domain

## Planning Context

### Feature Summary
Define a comprehensive JSON-based schema format for describing database schemas, data generators, relationships, and validation rules. The format must support verticalized data generation (industry-specific realism), multiple tables with foreign keys, custom generators with distribution parameters, and schema versioning. Deliverables are specification document (`schemas/schema-spec.md`), example schema (`schemas/example-schema.json`), and optional JSON Schema validation file. This specification will be implemented by F008 (Schema Parser).

### Key Technical Decisions Required

**Phase 0 Research Topics**:
1. **JSON Schema Design Best Practices**: What makes a good schema format? How to balance simplicity with extensibility? How to structure nested objects (tables, columns, relationships)?
2. **Data Generator Type Patterns**: What generator types are needed? How to specify parameters (distribution, ranges, weights)? Built-in vs custom generators?
3. **Distribution Specification Format**: How to specify distributions (uniform, normal, lognormal, weighted)? What parameters does each need? How to make it clear and unambiguous?
4. **Foreign Key Relationship Syntax**: How to express foreign keys in JSON? What referential integrity rules (CASCADE, SET NULL, RESTRICT)? How to ensure clarity?
5. **Column Type Mappings**: What SQL column types to support? How to handle MySQL vs PostgreSQL differences? What's the common subset? What's platform-specific?
6. **Schema Versioning Strategy**: How do schemas evolve? Semantic versioning (major.minor.patch)? What constitutes breaking changes? How to handle backward compatibility?
7. **Validation Rules Specification**: What makes a schema valid? Required fields? Referential integrity checks? How to specify validation rules explicitly?
8. **Extensibility Patterns**: How to add new generator types without breaking existing schemas? How to support future database types? How to enable community contributions?
9. **Generator Parameter Schema**: How to structure generator parameters? Typed parameters? Validation rules? Clear examples for each generator type?
10. **Example Schema Complexity Level**: What complexity for example schema? Tier 1 (simple, <30s generation)? How many tables? What relationships? What demonstrates all features without overwhelming?

### Technical Context (Pre-filled)

**Language/Version**: N/A (JSON specification document, language-agnostic)
**Primary Dependencies**: None (design work, no code dependencies)
**Storage**: Schemas will be stored as JSON files, embedded in Go binary via embed.FS (F008 implementation)
**Testing**: Manual validation (create example schema, verify JSON syntax, verify completeness)
**Target Platform**: JSON format is cross-platform (macOS, Linux, Windows)
**Project Type**: Specification document (design work, not implementation)
**Performance Goals**:
  - JSON parsing: Fast in Go standard library (microseconds for typical schema)
  - Schema validation: Once at load time (startup overhead negligible)
  - Format must enable fast generator lookup (O(1) or O(log n))
**Constraints**:
  - Must be human-readable (developers will write and edit schemas)
  - Must be machine-parseable (Go code will load and validate)
  - Must support MySQL and PostgreSQL (Phase 1 databases)
  - Must support custom generators (verticalized data)
  - Must be extensible (new generators, new databases, new features)
  - Must include validation rules (what's valid, what's invalid)
  - Must follow semver for schema versioning
  - Example schema must be Tier 1 complexity (<30s generation, MVP focus)
**Scale/Scope**: Foundation for all schemas (F010 fintech, F016 healthcare, F018 retail), community contributions (50+ schemas by Phase 2)

## Planning Workflow

### Phase 0: Research & Technical Decisions

Generate `research.md` with documented decisions for:

#### 1. JSON Schema Design Best Practices
- **Decision Point**: What structure for top-level schema object? What sections (metadata, tables, relationships, generation_order)?
- **Research**: JSON schema best practices, database schema formats, Faker.js schema, Synthea schema (healthcare)
- **Output**: Schema structure with:
  - **schema_version**: Format version (for future schema format changes)
  - **name**: Unique identifier (e.g., "fintech-loans")
  - **description**: Human-readable description
  - **author**: Schema creator
  - **version**: Schema content version (semver)
  - **database_type**: Supported databases (["mysql", "postgres"])
  - **metadata**: Industry, tags, total_records (for documentation)
  - **tables**: Array of table definitions (core content)
  - **relationships**: Array of foreign key relationships (explicit referential integrity)
  - **generation_order**: Array of table names (ensures parent tables generated first)
  - **validation_rules**: Array of validation rules (soft vs hard enforcement)

#### 2. Data Generator Type Patterns
- **Decision Point**: What generator types are needed? How to categorize (built-in vs custom)? How to specify parameters?
- **Research**: Faker.js generators, Python Faker library, Synthea generators, common data patterns
- **Output**: Generator type categories:
  - **Built-in Personal Data**: first_name, last_name, full_name, email, phone, address, ssn, date_of_birth
  - **Built-in Company Data**: company_name, job_title, company_email, domain
  - **Built-in Date/Time**: timestamp_past, timestamp_future, date_between
  - **Built-in Numeric**: int_range, float_range, decimal_range
  - **Custom Generators**: Defined per schema (credit_score, loan_amount, diagnosis_code, product_name)
  - **Generator parameter format**:
    ```json
    "generator": "credit_score",
    "generator_params": {
      "distribution": "normal",
      "mean": 680,
      "std_dev": 80,
      "min": 300,
      "max": 850
    }
    ```

#### 3. Distribution Specification Format
- **Decision Point**: What distributions are needed? How to specify parameters clearly?
- **Research**: Statistical distributions, realistic data modeling, lognormal distributions for skewed data
- **Output**: Distribution types:
  - **uniform**: Evenly distributed values
    - Parameters: min, max
    - Use case: Random IDs, equally likely categories
  - **normal**: Bell curve distribution
    - Parameters: mean, std_dev, min (optional), max (optional)
    - Use case: Credit scores, heights, test scores
  - **lognormal**: Right-skewed distribution (naturally skewed data)
    - Parameters: median, min, max
    - Use case: Loan amounts, income, product prices
  - **weighted**: Specific values with weights (realistic frequency)
    - Parameters: values array with {value, weight}
    - Use case: Loan status (70% active, 25% paid_off, 5% delinquent)
  - **ranges**: Multiple ranges with weights (tiered distributions)
    - Parameters: ranges array with {min, max, weight}
    - Use case: Interest rates (60% 3-6%, 30% 6-10%, 10% 10-15%)

#### 4. Foreign Key Relationship Syntax
- **Decision Point**: How to express foreign keys in JSON? Inline in column definition or separate relationships section?
- **Research**: SQL foreign key syntax, JSON relationship formats, clarity vs verbosity tradeoffs
- **Output**: Dual representation (inline + explicit):
  - **Inline foreign key** (in column definition):
    ```json
    {
      "name": "borrower_id",
      "type": "int",
      "foreign_key": {
        "table": "borrowers",
        "column": "id",
        "on_delete": "CASCADE",
        "on_update": "CASCADE"
      },
      "nullable": false
    }
    ```
  - **Explicit relationships** (separate section for documentation):
    ```json
    {
      "from_table": "loans",
      "from_column": "borrower_id",
      "to_table": "borrowers",
      "to_column": "id",
      "relationship_type": "many_to_one",
      "description": "Each loan belongs to one borrower"
    }
    ```
  - **Rationale**: Inline for code (what F008 parser uses), explicit for documentation (human readability, relationship visualization)

#### 5. Column Type Mappings
- **Decision Point**: What SQL column types to support? How to handle MySQL vs PostgreSQL differences?
- **Research**: MySQL data types documentation, PostgreSQL data types documentation, common subset
- **Output**: Supported types:
  - **Integer types**: int, bigint, smallint, tinyint (MySQL/PostgreSQL common)
  - **Decimal types**: decimal(p,s), float, double (MySQL/PostgreSQL common)
  - **String types**: varchar(n), text, char(n) (MySQL/PostgreSQL common)
  - **Date/Time types**: date, datetime, timestamp (MySQL/PostgreSQL common)
  - **Boolean**: boolean, bit (MySQL/PostgreSQL common)
  - **JSON types**: json, jsonb (jsonb is PostgreSQL-specific, json is common)
  - **Enum types**: enum('val1','val2',...) (MySQL format, PostgreSQL uses custom types - document differences)
  - **Platform-specific handling**: Document in schema-spec.md which types are database-specific, provide fallback guidance

#### 6. Schema Versioning Strategy
- **Decision Point**: How do schemas evolve? What versioning scheme? What's a breaking change?
- **Research**: Semantic versioning (semver) documentation, API versioning patterns, schema evolution best practices
- **Output**: Semantic versioning rules:
  - **Major version** (1.0.0 → 2.0.0): Breaking changes
    - Table or column removal
    - Column type changes (int → string)
    - Foreign key changes
    - Generator parameter incompatibilities
  - **Minor version** (1.0.0 → 1.1.0): Backward-compatible additions
    - New tables (doesn't affect existing data generation)
    - New columns (added to existing tables)
    - New generators (new data types)
  - **Patch version** (1.0.0 → 1.0.1): Bug fixes, improvements
    - Data distribution improvements (more realistic)
    - Generator parameter tuning
    - Documentation clarifications
    - Metadata updates (tags, descriptions)
  - **Version field**: Two versions in schema:
    - `schema_version`: Format version (1.0 for all MVP schemas)
    - `version`: Content version (1.0.0, 1.1.0, 2.0.0 per schema evolution)

#### 7. Validation Rules Specification
- **Decision Point**: What makes a schema valid? Required fields? Referential integrity? How to specify explicitly?
- **Research**: JSON Schema validation, database schema validation, schema correctness rules
- **Output**: Validation requirements:
  - **Schema-level validation**:
    - Unique name (no conflicts with existing schemas)
    - At least one table defined
    - Valid database_type values
    - generation_order includes all table names
    - No circular dependencies in generation_order
  - **Table-level validation**:
    - Unique table name within schema
    - At least one column defined
    - Exactly one primary key
    - record_count > 0
  - **Column-level validation**:
    - Valid column type (from supported types)
    - Valid generator name (built-in or custom for this schema)
    - Generator parameters match generator requirements
    - Foreign keys reference existing tables and columns
    - Primary key columns are not nullable
  - **Relationship validation**:
    - Foreign keys reference existing tables
    - Foreign keys reference primary keys or unique columns
    - Referential integrity actions are valid (CASCADE, SET NULL, RESTRICT)
  - **Generation order validation**:
    - Parent tables come before child tables
    - No circular dependencies
  - **Optional validation_rules section**: Soft constraints (recommendations, not requirements)
    - Example: "Higher credit scores should correlate with higher loan amounts" (soft enforcement)

#### 8. Extensibility Patterns
- **Decision Point**: How to add new features without breaking existing schemas? How to enable community contributions?
- **Research**: JSON extensibility patterns, schema evolution, backward compatibility
- **Output**: Extensibility mechanisms:
  - **Generator extensibility**: Custom generators defined per schema
    - Schema can define new generators in custom_generators section (future enhancement)
    - Built-in generators are standard (first_name, email, etc.)
    - Custom generators are schema-specific (credit_score, loan_amount)
  - **Database extensibility**: database_type array supports multiple databases
    - Phase 1: ["mysql", "postgres"]
    - Phase 2: Add "mongodb" (NoSQL support)
    - Parser (F008) validates against supported list
  - **Column extensibility**: New column attributes can be added without breaking old schemas
    - Unknown attributes ignored by parser (forward compatibility)
    - Required attributes documented in schema-spec.md
  - **Validation extensibility**: validation_rules section is optional, extensible
    - Can add new rule types without breaking existing schemas
  - **Metadata extensibility**: metadata object is flexible (JSON object)
    - Can add new fields (industry, use_case, complexity_tier) without schema version bump
  - **Future-proofing**: schema_version field allows format changes
    - Schema format 1.0 (MVP)
    - Schema format 2.0 (future enhancements)
    - Parser checks schema_version and adapts parsing logic

#### 9. Generator Parameter Schema
- **Decision Point**: How to structure generator parameters? Typed? Validated? Clear examples?
- **Research**: Parameter schema patterns, type systems, validation approaches
- **Output**: Parameter structure:
  - **Structure**: `generator_params` is a JSON object with generator-specific keys
  - **Type safety**: Parser (F008) validates parameters match generator requirements
  - **Distribution parameters**:
    - **normal**: {mean, std_dev, min?, max?}
    - **lognormal**: {median, min, max}
    - **uniform**: {min, max}
    - **weighted**: {values: [{value, weight}, ...]}
    - **ranges**: {ranges: [{min, max, weight}, ...]}
  - **Date/Time parameters**:
    - **timestamp_past**: {days_ago_min, days_ago_max}
    - **timestamp_future**: {days_ahead_min, days_ahead_max}
    - **date_between**: {start_date, end_date}
  - **String parameters**:
    - **email**: {domain?} (optional custom domain)
    - **phone**: {format?} (optional format string)
  - **Examples for each generator type**: Documented in schema-spec.md with practical examples

#### 10. Example Schema Complexity Level
- **Decision Point**: What complexity for example schema? Tier 1 (simple)? What demonstrates all features?
- **Research**: Constitution schema tiers (lines 176-197), MVP focus on Tier 1 (<30s generation)
- **Output**: Example schema design:
  - **Tier**: Tier 1 (simple schema, <30s generation for 1,000 records)
  - **Tables**: 2-4 tables (borrowers, loans, payments) - demonstrates relationships
  - **Relationships**: 1-2 foreign keys (loans → borrowers, payments → loans)
  - **Generators**: Mix of built-in (first_name, email) and custom (credit_score, loan_amount)
  - **Distributions**: Demonstrate normal, lognormal, weighted (variety of patterns)
  - **Total records**: ~250 borrowers, ~1,000 loans, ~3,700 payments = 4,950 total (constitutional <30s target)
  - **Features demonstrated**:
    - Primary keys (auto_increment)
    - Foreign keys (with CASCADE)
    - Indexes (email unique, credit_score indexed)
    - Custom generators (credit_score, loan_amount, interest_rate, loan_status)
    - Distribution types (normal, lognormal, weighted)
    - Nullable columns (some optional fields)
    - Unique constraints (email)
    - Timestamps (created_at with default CURRENT_TIMESTAMP)
  - **Rationale**: Simple enough for MVP (<30s), complex enough to demonstrate all format features, realistic fintech example

**Deliverable**: `specs/005-f007-schema-json/research.md`

### Phase 1: Design & Contracts

#### 1. Data Model (SKIP for this feature)
**Rationale**: F007 IS the data model format definition. Generating a data-model.md would be recursive (defining the format for defining formats). Skip this deliverable.

#### 2. API Contracts (SKIP for this feature)
**Rationale**: F007 has no REST/GraphQL APIs. The JSON schema format itself IS the contract for F008 parser. No separate contracts/ directory needed. The schema-spec.md document serves as the contract.

#### 3. Quickstart Guide (REQUIRED - Adapted for Documentation)
Generate `quickstart.md` with focus on **specification validation** (not code execution):

```markdown
# F007 Quickstart: Schema Format Specification Validation

## Prerequisites
- Constitution read (lines 176-197 for schema complexity tiers)
- F007 spec read (understanding of requirements)
- JSON syntax validator available (jsonlint, IDE, or online tool)

## Validation Overview

The Schema JSON Format Specification is validated by:
1. Creating comprehensive specification document (schema-spec.md)
2. Creating example schema demonstrating all features (example-schema.json)
3. Verifying JSON syntax is valid
4. Verifying all sections are documented
5. Verifying specification is implementable by F008 parser

## Manual Validation Steps

### 1. Verify Specification Document Completeness
```bash
# Check schema-spec.md exists and contains all required sections
cat schemas/schema-spec.md

# Required sections checklist:
# - Overview (what is the format, why JSON)
# - Schema Structure (top-level fields)
# - Table Definition (table object structure)
# - Column Definition (column object structure)
# - Generator Types (built-in and custom)
# - Distribution Types (uniform, normal, lognormal, weighted, ranges)
# - Foreign Key Relationships (inline and explicit)
# - Schema Versioning (semver rules)
# - Validation Rules (what makes a schema valid)
# - Supported Data Types (MySQL and PostgreSQL common types)
# - Example Schema (reference to example-schema.json)
```

Expected: All sections present with clear explanations, code examples, and practical guidance

### 2. Verify Example Schema JSON Syntax
```bash
# Validate JSON syntax
cat schemas/example-schema.json | python -m json.tool

# Or use jq
cat schemas/example-schema.json | jq '.'

# Or use jsonlint online: https://jsonlint.com/
```

Expected: Valid JSON with no syntax errors, properly formatted

### 3. Verify Example Schema Demonstrates All Features
```bash
# Check example schema includes:
# - schema_version field (format version)
# - name, description, author, version (metadata)
# - database_type array (MySQL and PostgreSQL)
# - metadata object (industry, tags, total_records)
# - tables array with multiple tables
# - columns with various types (int, varchar, decimal, timestamp, etc.)
# - Primary keys (auto_increment)
# - Foreign keys (inline foreign_key object)
# - Indexes (unique and non-unique)
# - Generators (built-in: first_name, email; custom: credit_score, loan_amount)
# - Generator parameters (distribution, mean, std_dev, min, max, weights)
# - Relationships section (explicit documentation)
# - generation_order array (parent tables first)
# - validation_rules section (optional soft constraints)
```

Expected: Example schema demonstrates all documented features

### 4. Verify Tier 1 Complexity Compliance
```bash
# Check example schema record counts
cat schemas/example-schema.json | jq '.metadata.total_records'

# Verify Tier 1 target: <5,000 total records (constitutional <30s generation target)
# Example breakdown:
# - borrowers: 250
# - loans: 1,000
# - payments: 3,700
# Total: 4,950 (within Tier 1)
```

Expected: Total records <5,000 (Tier 1 complexity, constitutional constraint)

### 5. Verify Schema Versioning Documentation
```bash
# Check schema-spec.md has versioning section
grep -A 20 "Schema Versioning" schemas/schema-spec.md

# Verify documents:
# - schema_version (format version, currently 1.0)
# - version (content version, semver format)
# - Major version rules (breaking changes)
# - Minor version rules (backward-compatible additions)
# - Patch version rules (bug fixes, improvements)
```

Expected: Clear versioning strategy documented with examples

### 6. Verify Validation Rules Documentation
```bash
# Check schema-spec.md has validation section
grep -A 30 "Validation Rules" schemas/schema-spec.md

# Verify documents:
# - Schema-level validation (unique name, at least one table, etc.)
# - Table-level validation (unique name, at least one column, one primary key)
# - Column-level validation (valid types, valid generators, FK references exist)
# - Relationship validation (FKs reference existing tables/columns)
# - Generation order validation (parent tables first, no circular deps)
```

Expected: Comprehensive validation rules that F008 parser can implement

### 7. Verify Generator Documentation
```bash
# Check schema-spec.md has generator section
grep -A 50 "Generator Types" schemas/schema-spec.md

# Verify documents:
# - Built-in personal data generators (first_name, last_name, email, etc.)
# - Built-in company data generators (company_name, job_title, etc.)
# - Built-in date/time generators (timestamp_past, timestamp_future, etc.)
# - Built-in numeric generators (int_range, float_range, decimal_range)
# - Custom generator pattern (how schemas define custom generators)
# - Generator parameter structure (generator_params object)
# - Distribution types with parameters (normal, lognormal, weighted, etc.)
```

Expected: All generator types documented with parameter examples

### 8. Verify Data Type Documentation
```bash
# Check schema-spec.md has data types section
grep -A 30 "Supported Data Types" schemas/schema-spec.md

# Verify documents:
# - MySQL/PostgreSQL common types (int, varchar, decimal, timestamp, etc.)
# - Platform-specific types (jsonb for PostgreSQL)
# - Type format specifications (varchar(n), decimal(p,s))
# - Enum type format (enum('val1','val2'))
```

Expected: All supported SQL data types documented with examples

### 9. Verify Foreign Key Relationship Documentation
```bash
# Check schema-spec.md has foreign key section
grep -A 40 "Foreign Key Relationships" schemas/schema-spec.md

# Verify documents:
# - Inline foreign_key object structure
# - Referential integrity actions (CASCADE, SET NULL, RESTRICT)
# - Explicit relationships section format
# - Relationship types (one_to_one, one_to_many, many_to_one, many_to_many)
```

Expected: Foreign key format clear, both inline and explicit representations

### 10. Verify F008 Implementability
```bash
# Review specification for clarity:
# - Can F008 parser unambiguously parse this format?
# - Are all required fields clearly marked?
# - Are all optional fields clearly marked?
# - Are validation rules specific enough to implement?
# - Are error messages for invalid schemas described?
```

Expected: Specification is unambiguous, implementable by F008 parser

## Verification Checklist
- [ ] schema-spec.md exists with all required sections
- [ ] example-schema.json exists and is valid JSON
- [ ] Example schema demonstrates all documented features
- [ ] Example schema is Tier 1 complexity (<5,000 records, <30s generation)
- [ ] Schema versioning strategy documented (semver)
- [ ] Validation rules comprehensively documented
- [ ] All generator types documented with parameters
- [ ] All supported data types documented
- [ ] Foreign key relationships clearly documented
- [ ] Specification is unambiguous and implementable
- [ ] Example schema can be used as template for F010, F016, F018
- [ ] Format supports custom generators (verticalized data)
- [ ] Format supports multiple databases (MySQL, PostgreSQL)
- [ ] Format is extensible (future generators, databases, features)

## Common Issues and Fixes

### Issue: JSON syntax errors in example schema
- **Check**: Missing commas, trailing commas, unquoted keys, unescaped strings
- **Fix**: Use JSON validator (jsonlint, jq, IDE JSON validator)

### Issue: Example schema too complex (>30s generation)
- **Check**: Total records >10,000? Too many foreign key lookups?
- **Fix**: Reduce record counts, simplify relationships, target Tier 1 (<5,000 records)

### Issue: Specification ambiguous
- **Check**: Can you implement a parser from the spec? Are required fields clear?
- **Fix**: Add explicit "Required" vs "Optional" labels, add code examples

### Issue: Generator parameters unclear
- **Check**: Are parameter names clear? Are types documented? Are examples provided?
- **Fix**: Add generator parameter schema for each generator type with examples

### Issue: Validation rules missing edge cases
- **Check**: What happens if... (circular dependencies, missing foreign keys, invalid types)?
- **Fix**: Add edge case handling to validation rules section

### Issue: Foreign key syntax confusing
- **Check**: Inline vs explicit - are both needed? Is purpose of each clear?
- **Fix**: Document rationale (inline for code, explicit for documentation)

## Performance Validation (Not Applicable for F007)
F007 is specification work. Performance validation happens in F008 (parser) and F010/F016/F018 (schema implementations).

## Next Steps
- F008: Implement schema parser in Go (uses this specification)
- F010: Create fintech-loans schema (uses this format)
- F016: Create healthcare schema (uses this format)
- F018: Create retail schema (uses this format)
- Community contributions: Enable schema contributions using this format

## Design Quality Checklist

### Simplicity
- [ ] JSON format is simple, readable, human-editable
- [ ] No unnecessary nesting or complexity
- [ ] Field names are clear, self-explanatory
- [ ] Examples are easy to understand

### Extensibility
- [ ] New generators can be added without breaking existing schemas
- [ ] New databases can be added without format changes
- [ ] New column attributes can be added (forward compatibility)
- [ ] schema_version allows future format evolution

### Clarity
- [ ] Specification answers common questions (How do I...?)
- [ ] Required vs optional fields are explicit
- [ ] Validation rules are unambiguous
- [ ] Error conditions are documented

### Realism
- [ ] Format enables verticalized data (custom generators)
- [ ] Supports realistic distributions (normal, lognormal, weighted)
- [ ] Supports complex relationships (foreign keys, indexes)
- [ ] Example schema demonstrates real-world fintech data

### Implementability
- [ ] F008 parser can be implemented from this spec
- [ ] Validation rules are specific enough to implement
- [ ] Error messages for invalid schemas are described
- [ ] Edge cases are documented
```

**Deliverable**: `specs/005-f007-schema-json/quickstart.md`

#### 4. Update Agent Context
Run: `.specify/scripts/bash/update-agent-context.sh claude`

This updates the Claude-specific context file (CLAUDE.md) with:
- Schema JSON format overview
- Generator types and parameters
- Distribution types (normal, lognormal, weighted, ranges)
- Foreign key relationship format
- Schema versioning (semver)
- Validation rules
- Tier 1 complexity targets (<30s, <5,000 records)
- How to create new schemas using this format

**Deliverable**: Updated `CLAUDE.md` with schema format documentation

## Constitution Re-verification

After Phase 1 design, verify:
- [ ] JSON format is simple, readable (Simple > Complex principle)
- [ ] Format supports verticalized data via custom generators (Verticalized > Generic principle)
- [ ] Format enables fast parsing in Go (<1ms) (Speed > Features principle)
- [ ] Format is extensible for future needs (Ship Fast, Validate Early principle)
- [ ] Example schema is Tier 1 complexity (<5,000 records, <30s) (Technical Constraint 1)
- [ ] Format supports both MySQL and PostgreSQL (Technical Constraint 3)
- [ ] Specification is clear enough for community contributions (Developer-First Design principle)
- [ ] Format is boring, proven (JSON standard) (Boring Tech Wins principle)
- [ ] Design work takes 2 days (Ship Fast principle)
- [ ] No over-engineering (format is as simple as possible, but no simpler) (Anti-Pattern 5)
- [ ] Validation rules prevent invalid schemas (Code Quality constraint)
- [ ] Schema versioning enables evolution (Product Philosophy 1)

## Deliverables Summary

**Generated by /speckit.plan**:
1. ✅ `specs/005-f007-schema-json/plan.md` - This file
2. ✅ `specs/005-f007-schema-json/research.md` - Phase 0 output (10 schema design decisions)
3. ✅ `specs/005-f007-schema-json/quickstart.md` - Phase 1 output (spec validation guide)
4. ⏭️ `specs/005-f007-schema-json/data-model.md` - SKIP (F007 defines the model format itself)
5. ⏭️ `specs/005-f007-schema-json/contracts/` - SKIP (schema format IS the contract)
6. ✅ Updated CLAUDE.md with schema format guidelines

**NOT generated by /speckit.plan** (created later):
- `specs/005-f007-schema-json/tasks.md` - Phase 2, separate command (/speckit.tasks)
- `schemas/schema-spec.md` - Created during implementation
- `schemas/example-schema.json` - Created during implementation

## Success Criteria for Planning Phase

- ✅ All 10 research decisions documented with clear rationale
- ✅ JSON schema structure follows best practices (simplicity, clarity, extensibility)
- ✅ Generator types and parameters clearly defined
- ✅ Distribution types documented (uniform, normal, lognormal, weighted, ranges)
- ✅ Foreign key relationship format is clear (inline + explicit)
- ✅ Column type mappings cover MySQL and PostgreSQL common types
- ✅ Schema versioning strategy follows semver
- ✅ Validation rules are comprehensive and implementable
- ✅ Extensibility patterns enable future enhancements
- ✅ Example schema is Tier 1 complexity (constitutional constraint)
- ✅ Quickstart provides clear specification validation steps
- ✅ Constitution compliance verified (no violations)
- ✅ Agent context updated with schema format guidelines
- ✅ Planning artifacts reference constitution and spec correctly

## Anti-Patterns to Avoid

- ❌ Don't generate data-model.md (F007 IS the model format definition)
- ❌ Don't generate contracts/ (schema format itself is the contract)
- ❌ Don't implement code in F007 (design phase only, implementation in F008)
- ❌ Don't overcomplicate JSON format (simplicity enables clarity)
- ❌ Don't skip versioning strategy (schemas will evolve)
- ❌ Don't make format database-specific (must support MySQL AND PostgreSQL)
- ❌ Don't hardcode generator types (must be extensible for custom generators)
- ❌ Don't skip validation rules (F008 needs clear implementation guidance)
- ❌ Don't make example schema too complex (Tier 1 target: <5,000 records, <30s)
- ❌ Don't use obscure JSON patterns (keep it standard, widely understood)
- ❌ Don't forget extensibility (community will contribute 50+ schemas)
- ❌ Don't skip distribution parameter documentation (generators need clear params)
- ❌ Don't make foreign key syntax ambiguous (clarity is critical)
- ❌ Don't ignore edge cases in validation rules (circular dependencies, missing FKs)

## Implementation Notes

### For the AI Agent
When executing `/speckit.plan` with this prompt:

1. **Start with comprehensive research.md**: Document all 10 research decisions with:
   - Decision: What was chosen (JSON structure, generator types, distribution format, etc.)
   - Rationale: Why chosen (simplicity, clarity, extensibility, constitution alignment)
   - Alternatives considered: What else was evaluated (YAML vs JSON, inline vs separate relationships)
   - Source: Where information came from (JSON Schema docs, SQL docs, Faker patterns)

2. **Be explicit about skips**: Clearly state why data-model.md and contracts/ are not needed for F007:
   - F007 IS the data model format definition (recursive to define format for defining formats)
   - Schema format itself IS the contract for F008 parser implementation
   - No REST/GraphQL APIs (no need for API contracts)

3. **Focus on quickstart.md for documentation validation**: This is the primary deliverable beyond research. Include:
   - How to validate specification completeness (all sections present)
   - How to validate example schema JSON syntax (jsonlint, jq)
   - How to verify all features are demonstrated
   - How to check Tier 1 complexity compliance (<5,000 records)
   - How to verify F008 implementability (unambiguous, clear)
   - Troubleshooting common issues (JSON syntax, complexity, ambiguity)
   - Design quality checklist (simplicity, extensibility, clarity, realism, implementability)

4. **Verify constitutional compliance**:
   - Verticalized > Generic: Custom generators enable industry-specific realism
   - Simple > Complex: JSON format is simple, human-readable, machine-parseable
   - Developer-First: Clear specification enables community contributions
   - Boring Tech: JSON is proven, standard, widely understood
   - Speed: Format enables fast parsing in Go (JSON standard library)
   - Tier 1 complexity: Example schema <5,000 records (<30s generation target)

5. **Keep it standard**: This is intentionally simple specification work:
   - Use standard JSON format (no custom extensions)
   - Follow JSON naming conventions (snake_case for consistency)
   - Standard SQL column types (MySQL/PostgreSQL common subset)
   - Semantic versioning for schemas (semver is proven, understood)
   - Predictable and maintainable format

### Schema Format Design Principles

**Simplicity First**:
- JSON format over YAML (machine-parseable, Go standard library)
- Flat structure where possible (minimal nesting)
- Clear field names (self-documenting)
- No magic values or implicit behavior

**Human-Readable**:
- Developers will write and edit schemas
- Field names are descriptive (generator_params, foreign_key, generation_order)
- Examples demonstrate common patterns
- Comments in example schema (via "description" fields)

**Machine-Parseable**:
- Go JSON unmarshalling is straightforward
- Type-safe (parser can validate types)
- Validation rules are explicit (not inferred)
- Clear error messages for invalid schemas

**Extensible**:
- Custom generators per schema
- New generators don't break existing schemas
- New database types can be added
- schema_version enables format evolution

**Verticalized**:
- Custom generators enable industry-specific realism
- Distribution parameters allow realistic data patterns
- Example schema demonstrates fintech realism (not generic "John Doe")

### Generator Design Philosophy

**Built-in Generators** (standard across all schemas):
- Personal data (first_name, last_name, email, phone)
- Company data (company_name, job_title)
- Date/time (timestamp_past, timestamp_future, date_between)
- Numeric ranges (int_range, float_range, decimal_range)
- Rationale: Common patterns, reusable across schemas

**Custom Generators** (defined per schema):
- Industry-specific values (credit_score, loan_amount, diagnosis_code, product_name)
- Verticalized data (fintech, healthcare, retail)
- Realistic distributions (normal, lognormal, weighted)
- Rationale: Enables verticalized realism without bloating core format

### Distribution Type Selection Guide

**When to use uniform**:
- Equally likely values (dice roll, random ID)
- No natural central tendency
- Example: int_range(1, 100) uniform distribution

**When to use normal** (bell curve):
- Natural central tendency (heights, test scores, credit scores)
- Most values cluster around mean
- Example: credit_score with mean=680, std_dev=80

**When to use lognormal** (right-skewed):
- Naturally skewed data (income, loan amounts, product prices)
- Cannot be negative (always positive)
- Long tail (few very high values)
- Example: loan_amount with median=50000, min=5000, max=500000

**When to use weighted** (specific values):
- Categorical data with realistic frequencies
- Example: loan_status (70% active, 25% paid_off, 5% delinquent)

**When to use ranges** (tiered distributions):
- Multiple ranges with different frequencies
- Example: interest_rate (60% 3-6%, 30% 6-10%, 10% 10-15%)

### Foreign Key Relationship Dual Representation

**Why both inline and explicit?**
- **Inline foreign_key** (in column definition):
  - Used by F008 parser (what code reads)
  - Defines actual SQL constraint
  - On_delete and on_update rules
  - Part of column specification
- **Explicit relationships** (separate section):
  - Documentation for humans (what developers read)
  - Relationship visualization
  - Relationship type (one_to_one, one_to_many, many_to_one)
  - Natural language description
- **Rationale**: Code needs inline (part of column), humans need explicit (relationship overview)

### Schema Versioning Examples

**Major version bump** (1.0.0 → 2.0.0):
- Remove "payments" table → breaking change
- Change "credit_score" from int to varchar → breaking change
- Remove "borrower_id" foreign key → breaking change

**Minor version bump** (1.0.0 → 1.1.0):
- Add "transactions" table → backward compatible
- Add "middle_name" column to borrowers → backward compatible
- Add "payment_method" generator → backward compatible

**Patch version bump** (1.0.0 → 1.0.1):
- Improve credit_score distribution (mean 680→700) → bug fix
- Update schema description → documentation
- Add industry tag to metadata → metadata update

### Related Constitution Sections
- **Core Principle I**: Verticalized > Generic (custom generators enable industry-specific realism)
- **Core Principle IV**: Boring Tech Wins (JSON is proven, standard, widely understood)
- **Core Principle VI**: Developer-First Design (clear specification enables community contributions)
- **Core Principle VII**: Ship Fast, Validate Early (design work, 2 days, no implementation delay)
- **Technical Constraint 1**: Performance (Tier 1 target: <30s generation, <5,000 records)
- **Technical Constraint 3**: Database Support (MySQL and PostgreSQL, format supports both)
- **Technical Constraint 5**: Code Quality Standards (specification must be clear, implementable)
- **Product Philosophy 1**: MVP Mindset (start with simple format, expand via community)
- **Anti-Pattern 5**: Over-Engineering (format is as simple as possible, but no simpler)
- **Anti-Pattern 6**: Generic Data (custom generators enable verticalized realism)

## Drag-and-Drop Usage

**To use this prompt**:
1. Drag this file into Claude Code
2. Claude will execute the `/speckit.plan` workflow for F007
3. Expected outputs:
   - research.md with 10 documented schema design decisions
   - quickstart.md with specification validation guide
   - Updated CLAUDE.md with schema format guidelines
   - Constitution compliance verified

**Estimated time**: 20-30 minutes for complete planning phase

**Next command**: `/speckit.tasks` to generate tasks.md from this plan

**Success indicators**:
- Research decisions are clear and actionable
- JSON schema structure follows best practices (simplicity, extensibility, clarity)
- Generator types and distribution formats are well-documented
- Foreign key relationship format is unambiguous
- Schema versioning strategy follows semver
- Validation rules are comprehensive and implementable
- Example schema is Tier 1 complexity (<5,000 records, <30s target)
- Quickstart provides clear specification validation steps
- No constitutional violations identified
- Ready to proceed to task generation (/speckit.tasks)
- Specification is implementable by F008 parser
- Format enables community schema contributions
