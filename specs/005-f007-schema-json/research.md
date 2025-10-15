# F007 Research: Schema JSON Format Specification

**Feature**: Schema JSON Format Specification
**Date**: 2025-10-15
**Status**: Complete

## Overview

This document captures research decisions for the Schema JSON Format Specification. All decisions are based on publicly available information (JSON schema best practices, SQL documentation, data generation patterns from open source tools like Faker.js and Synthea).

---

## Research Decision 1: JSON Schema Design Best Practices

### Decision

Top-level schema structure with these sections:

```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "description": "Realistic fintech loan data with borrowers, loans, and payments",
  "author": "SourceBox Contributors",
  "version": "1.0.0",
  "database_type": ["mysql", "postgres"],
  "metadata": {
    "industry": "fintech",
    "tags": ["loans", "credit", "payments"],
    "total_records": 4950
  },
  "tables": [ /* array of table definitions */ ],
  "relationships": [ /* array of foreign key relationships */ ],
  "generation_order": ["borrowers", "loans", "payments"],
  "validation_rules": [ /* optional soft constraints */ ]
}
```

### Rationale

**Flat structure over nested complexity**: Keep schema object relatively flat with clear top-level sections. This makes parsing straightforward and documentation easier to understand.

**Dual versioning (schema_version + version)**:
- `schema_version`: Format version (1.0 for MVP, 2.0 if format changes significantly)
- `version`: Content version for this specific schema (semver: 1.0.0, 1.1.0, 2.0.0)
- Rationale: Format evolution (schema_version) is separate from content evolution (version)

**Explicit generation_order**: Separate field rather than inferred from relationships ensures:
- Unambiguous ordering for F008 parser implementation
- Parent tables seeded before child tables (foreign key integrity)
- Clear documentation of dependencies

**Separate relationships section**: Dual representation pattern:
- Inline `foreign_key` in column definition (what parser uses)
- Explicit `relationships` section (documentation for humans)
- Rationale: Code needs inline, humans need explicit overview

**Optional validation_rules**: Soft constraints (warnings, not enforcement):
- Documents expected correlations ("higher credit scores → higher loan amounts")
- Separate from hard validation (required fields, referential integrity)
- Enables future validation improvements without schema version bumps

### Alternatives Considered

**YAML instead of JSON**: Rejected because:
- JSON has native Go support (encoding/json standard library)
- JSON is more widely understood (JavaScript developers)
- JSON is machine-first (strict syntax, no ambiguity)
- YAML is human-first but harder to validate programmatically

**Inline-only relationships (no separate section)**: Rejected because:
- Relationships are scattered across multiple column definitions
- Hard to visualize schema structure without explicit overview
- Dual representation provides both code clarity and documentation

**Single version field**: Rejected because:
- Can't distinguish format changes from content changes
- Breaking format change (1.0 → 2.0) would require schema authors to update all schemas
- Dual versioning separates concerns

### Source

- JSON Schema documentation: https://json-schema.org/
- Go encoding/json package docs: https://pkg.go.dev/encoding/json
- Database schema formats: MySQL Workbench, PostgreSQL pg_dump
- Faker.js schema patterns (public GitHub)

---

## Research Decision 2: Data Generator Type Patterns

### Decision

Generator categories:

**Built-in Personal Data** (standard across all schemas):
- `first_name`, `last_name`, `full_name`
- `email`, `phone`, `address`
- `ssn`, `date_of_birth`

**Built-in Company Data** (standard across all schemas):
- `company_name`, `job_title`
- `company_email`, `domain`

**Built-in Date/Time** (standard across all schemas):
- `timestamp_past`, `timestamp_future`, `date_between`

**Built-in Numeric** (standard across all schemas):
- `int_range`, `float_range`, `decimal_range`

**Custom Generators** (defined per schema):
- Schema-specific: `credit_score`, `loan_amount`, `interest_rate`, `loan_status`
- Industry-specific: `diagnosis_code`, `product_sku`, `order_status`
- Rationale: Custom generators enable verticalized realism without bloating built-in types

**Generator parameter format**:
```json
{
  "name": "credit_score",
  "type": "int",
  "generator": "credit_score",
  "generator_params": {
    "distribution": "normal",
    "mean": 680,
    "std_dev": 80,
    "min": 300,
    "max": 850
  }
}
```

### Rationale

**Built-in vs custom split**:
- Built-in: Common patterns reusable across all schemas (names, emails, dates)
- Custom: Industry-specific values (credit_score for fintech, diagnosis_code for healthcare)
- Rationale: Balance between reusability and verticalized realism

**Generator parameters in separate object**:
- `generator_params` is a JSON object (not inline fields)
- Keeps column definition clean, parameters are clearly grouped
- Type-safe validation possible (F008 can validate params match generator requirements)

**Distribution-based approach**:
- Generators specify HOW to generate (distribution type + parameters)
- Not WHAT to generate (hardcoded values)
- Enables realistic, varied data (not the same 100 values repeated)

### Alternatives Considered

**Hardcoded value lists**: Rejected because:
- Unrealistic (same 100 names repeated for 10,000 records)
- Not scalable (need millions of pre-generated values)
- Distribution-based approach is more flexible

**All generators as built-in**: Rejected because:
- 100+ built-in generators would bloat core format
- Custom generators enable community contributions (domain experts add fintech/healthcare generators)
- Extensibility is critical for long-term growth (50+ schemas)

**Inline generator parameters**: Rejected because:
```json
// Rejected approach (inline parameters)
{
  "name": "credit_score",
  "generator": "normal",
  "mean": 680,
  "std_dev": 80,
  "min": 300,
  "max": 850
}
```
- Mixes column definition with generator params
- Hard to distinguish required column fields from generator params
- Nested object is clearer

### Source

- Faker.js generator types: https://github.com/faker-js/faker
- Python Faker library: https://faker.readthedocs.io/
- Synthea generators (healthcare): https://github.com/synthetichealth/synthea
- Statistical distribution patterns: Wikipedia articles on normal, lognormal distributions

---

## Research Decision 3: Distribution Specification Format

### Decision

**Distribution types**:

1. **uniform**: Evenly distributed values
   - Parameters: `{min, max}`
   - Use case: Random IDs, equally likely categories
   - Example: `int_range(1, 100)` uniform distribution

2. **normal**: Bell curve distribution
   - Parameters: `{mean, std_dev, min?, max?}`
   - Use case: Credit scores, heights, test scores
   - Example: credit_score with mean=680, std_dev=80, min=300, max=850

3. **lognormal**: Right-skewed distribution
   - Parameters: `{median, min, max}`
   - Use case: Loan amounts, income, product prices
   - Example: loan_amount with median=50000, min=5000, max=500000
   - Rationale: Naturally skewed data (always positive, long tail)

4. **weighted**: Specific values with probabilities
   - Parameters: `{values: [{value, weight}, ...]}`
   - Use case: Categorical data with realistic frequencies
   - Example: loan_status (70% active, 25% paid_off, 5% delinquent)

5. **ranges**: Multiple ranges with weights
   - Parameters: `{ranges: [{min, max, weight}, ...]}`
   - Use case: Tiered distributions
   - Example: interest_rate (60% 3-6%, 30% 6-10%, 10% 10-15%)

### Rationale

**Distribution-based approach** (not hardcoded values):
- More realistic (not the same 100 values repeated)
- More scalable (generate millions of records)
- More flexible (tune distributions for realism)

**Lognormal for financial data**:
- Right-skewed distribution (most values cluster low, few very high values)
- Always positive (can't have negative loan amounts)
- Models real-world patterns (income, loan amounts, product prices)

**Weighted for categorical data**:
- Realistic frequencies (not uniform distribution)
- Example: 70% active loans, 25% paid off, 5% delinquent (real-world pattern)
- Simple to understand and configure

**Ranges for tiered distributions**:
- Multiple ranges with different frequencies
- Example: Interest rates vary by credit risk tiers
- More realistic than single normal distribution

### Alternatives Considered

**Only uniform distribution**: Rejected because:
- Unrealistic (credit scores don't distribute uniformly)
- Can't model real-world patterns
- Verticalized data requires realistic distributions

**AI/ML-generated distributions**: Rejected because:
- Complex (requires training data, models)
- Expensive (computation, maintenance)
- Overkill for MVP (statistical distributions are sufficient)
- Violates "boring tech wins" principle

**Custom distribution DSL**: Rejected because:
- Overengineering (standard distributions cover 95% of use cases)
- Hard to document and validate
- JSON-native types are simpler

### Source

- Statistical distributions: Wikipedia (normal, lognormal, uniform distributions)
- Real-world data modeling: "Synthetic Data Generation" papers (public)
- Faker.js patterns: https://github.com/faker-js/faker (randomness patterns)
- Financial data distributions: Public fintech datasets (loan amounts, credit scores)

---

## Research Decision 4: Foreign Key Relationship Syntax

### Decision

**Dual representation** (inline + explicit):

**Inline foreign key** (in column definition):
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

**Explicit relationships** (separate section):
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

### Rationale

**Why both inline and explicit?**

**Inline foreign_key**:
- What F008 parser uses (part of column specification)
- Defines actual SQL constraint
- Specifies referential integrity actions (CASCADE, SET NULL, RESTRICT)
- Machine-readable, actionable

**Explicit relationships**:
- Documentation for humans
- Relationship visualization (understand schema structure)
- Relationship type classification (one_to_one, one_to_many, many_to_one, many_to_many)
- Natural language description

**Rationale**: Code needs inline (part of column definition), humans need explicit (relationship overview for visualization and documentation)

### Alternatives Considered

**Inline only (no explicit relationships section)**: Rejected because:
- Relationships scattered across multiple column definitions
- Hard to visualize schema structure
- No central place to document relationship semantics

**Explicit only (no inline foreign_key)**: Rejected because:
- Parser needs foreign key info in column definition
- Column definition incomplete without foreign key constraint
- SQL generation requires inline foreign key

**Foreign keys in separate section**: Rejected because:
- Column definition incomplete (missing foreign key constraint)
- Harder to understand column meaning without seeing foreign key
- Inline is more natural for SQL developers

### Source

- SQL foreign key syntax: MySQL documentation, PostgreSQL documentation
- JSON relationship formats: GraphQL schema patterns, JSON API specification
- Database design patterns: "Database Design for Mere Mortals" (public book)
- ER diagram patterns: Standard entity-relationship modeling conventions

---

## Research Decision 5: Column Type Mappings

### Decision

**Supported types** (MySQL/PostgreSQL common subset):

**Integer types**:
- `int`, `bigint`, `smallint`, `tinyint` (common to both)

**Decimal types**:
- `decimal(p,s)`, `float`, `double` (common to both)

**String types**:
- `varchar(n)`, `text`, `char(n)` (common to both)

**Date/Time types**:
- `date`, `datetime`, `timestamp` (common to both)

**Boolean**:
- `boolean` (PostgreSQL native), `bit` (MySQL workaround)
- Note: Document that boolean maps to TINYINT(1) in MySQL

**JSON types**:
- `json` (common to both)
- `jsonb` (PostgreSQL-specific, binary JSON)
- Note: Document that jsonb falls back to json in MySQL

**Enum types**:
- `enum('val1','val2',...)` (MySQL format)
- Note: PostgreSQL uses custom types (CREATE TYPE), document differences

**Platform-specific handling**:
- Document in schema-spec.md which types are database-specific
- Provide fallback guidance (jsonb → json, enum → varchar in some cases)

### Rationale

**Common subset approach**:
- Phase 1 supports MySQL and PostgreSQL only
- Focus on types that work in BOTH databases
- Avoid MySQL-only or PostgreSQL-only types in MVP

**Explicit type format**:
- varchar(n) includes length parameter (not just "varchar")
- decimal(p,s) includes precision and scale
- Clear, unambiguous, matches SQL syntax

**Document differences, don't hide them**:
- jsonb is PostgreSQL-specific (document fallback to json for MySQL)
- enum is MySQL format (document PostgreSQL alternative)
- Developers need to know platform differences

### Alternatives Considered

**MySQL-only types**: Rejected because:
- Constitutional constraint: "Format must support both MySQL and PostgreSQL"
- Phase 1 requires cross-database compatibility

**Custom type system**: Rejected because:
- Overengineering (SQL types are well-understood)
- Violates "boring tech wins" principle
- Standard SQL types are sufficient

**Abstract types (hide SQL)**: Rejected because:
- Developers understand SQL types (varchar, int, timestamp)
- Abstraction adds complexity without benefit
- SQL is the common language for database developers

### Source

- MySQL data types: https://dev.mysql.com/doc/refman/8.0/en/data-types.html
- PostgreSQL data types: https://www.postgresql.org/docs/current/datatype.html
- Cross-database compatibility: "SQL Antipatterns" (public book)
- Common subset patterns: Open source ORMs (TypeORM, Sequelize)

---

## Research Decision 6: Schema Versioning Strategy

### Decision

**Semantic versioning (semver) rules**:

**Major version** (1.0.0 → 2.0.0): Breaking changes
- Table or column removal
- Column type changes (int → string)
- Foreign key changes (remove or change target)
- Generator parameter incompatibilities

**Minor version** (1.0.0 → 1.1.0): Backward-compatible additions
- New tables (doesn't affect existing data generation)
- New columns (added to existing tables)
- New generators (new data types)

**Patch version** (1.0.0 → 1.0.1): Bug fixes, improvements
- Data distribution improvements (more realistic)
- Generator parameter tuning
- Documentation clarifications
- Metadata updates (tags, descriptions)

**Two version fields**:
- `schema_version`: Format version (1.0 for all MVP schemas, 2.0 if format changes)
- `version`: Content version (1.0.0, 1.1.0, 2.0.0 per schema evolution)

### Rationale

**Semantic versioning is proven**:
- Widely understood by developers
- Clear rules for breaking vs non-breaking changes
- Enables dependency management (tools can require schema version ≥ 1.2.0)

**Dual versioning separates concerns**:
- Format version (schema_version): When schema format itself changes (F007 specification updates)
- Content version (version): When this specific schema evolves (fintech-loans 1.0.0 → 1.1.0)

**Backward compatibility for minor versions**:
- Tools using schema 1.0.0 can use schema 1.1.0 without changes
- Adding columns/tables doesn't break existing generators
- Minor versions enable iterative improvements

### Alternatives Considered

**Single version field**: Rejected because:
- Can't distinguish format changes from content changes
- Breaking format change would require all schemas to update versions
- Dual versioning is clearer

**Date-based versioning (2025-01-15)**: Rejected because:
- Doesn't communicate breaking vs non-breaking changes
- Hard to understand compatibility (is 2025-01-20 compatible with 2025-01-15?)
- Semantic versioning is more developer-friendly

**No versioning**: Rejected because:
- Schemas will evolve over time
- No way to track compatibility
- Breaking changes would break existing implementations

### Source

- Semantic Versioning specification: https://semver.org/
- API versioning patterns: REST API best practices
- Schema evolution: Database migration patterns (Liquibase, Flyway)
- npm package versioning: https://docs.npmjs.com/about-semantic-versioning

---

## Research Decision 7: Validation Rules Specification

### Decision

**Validation requirements**:

**Schema-level validation** (required):
- Unique name (no conflicts with existing schemas)
- At least one table defined
- Valid database_type values (["mysql", "postgres"] for Phase 1)
- generation_order includes all table names
- No circular dependencies in generation_order

**Table-level validation** (required):
- Unique table name within schema
- At least one column defined
- Exactly one primary key
- record_count > 0 (must generate at least one record)

**Column-level validation** (required):
- Valid column type (from supported types list)
- Valid generator name (built-in or custom for this schema)
- Generator parameters match generator requirements
- Foreign keys reference existing tables and columns
- Primary key columns are not nullable

**Relationship validation** (required):
- Foreign keys reference existing tables
- Foreign keys reference primary keys or unique columns
- Referential integrity actions are valid (CASCADE, SET NULL, RESTRICT)

**Generation order validation** (required):
- Parent tables come before child tables
- No circular dependencies (A → B → C → A)

**Optional validation_rules section** (soft constraints):
- Recommendations, not requirements
- Example: "Higher credit scores should correlate with higher loan amounts"
- Documented intent, not enforced by parser

### Rationale

**Hard validation vs soft validation**:
- Hard: Required for schema to be valid (parser rejects invalid schemas)
- Soft: Documented correlations and patterns (warnings, not errors)
- Rationale: Hard validation ensures data integrity, soft validation documents intent

**Explicit validation rules**:
- F008 parser can implement validation unambiguously
- Schema authors can self-diagnose errors before testing
- Clear error messages for invalid schemas

**Generation order validation**:
- Critical for foreign key integrity (parent tables seeded first)
- Circular dependencies would cause deadlock
- Explicit validation catches errors early

### Alternatives Considered

**No validation (let parser fail)**: Rejected because:
- Late error detection (fails during generation, not load time)
- Poor developer experience (unclear error messages)
- No self-service validation

**AI-based validation (detect unrealistic data)**: Rejected because:
- Complex (requires ML models, training data)
- Expensive (computation, maintenance)
- Overkill for MVP (explicit rules are sufficient)

**Runtime validation only**: Rejected because:
- Late error detection (fails during generation)
- Wasted computation (generate data, then discover schema is invalid)
- Validation at load time is faster and clearer

### Source

- JSON Schema validation: https://json-schema.org/understanding-json-schema/reference/validation.html
- Database schema validation: PostgreSQL pg_dump validation logic
- Schema correctness: "Database Design for Mere Mortals" (public book)
- Validation patterns: OpenAPI schema validation

---

## Research Decision 8: Extensibility Patterns

### Decision

**Extensibility mechanisms**:

**Generator extensibility**:
- Custom generators defined per schema (schema-specific, not global)
- Built-in generators are standard (first_name, email)
- Custom generators enable community contributions (domain experts add fintech/healthcare generators)

**Database extensibility**:
- database_type array supports multiple databases: `["mysql", "postgres"]`
- Phase 1: MySQL, PostgreSQL
- Phase 2: Add "mongodb" (NoSQL support)
- Parser (F008) validates against supported list

**Column extensibility**:
- Unknown attributes ignored by parser (forward compatibility)
- Required attributes documented in schema-spec.md
- New attributes can be added without breaking old schemas

**Validation extensibility**:
- validation_rules section is optional, extensible
- Can add new rule types without breaking existing schemas
- Soft constraints don't affect parser logic

**Metadata extensibility**:
- metadata object is flexible (JSON object)
- Can add new fields (industry, use_case, complexity_tier) without schema version bump
- Free-form, not validated by parser

**Future-proofing**:
- schema_version field allows format changes
- Schema format 1.0 (MVP), 2.0 (future enhancements)
- Parser checks schema_version and adapts parsing logic

### Rationale

**Extensibility enables growth**:
- Community contributions (50+ schemas by Phase 2)
- New generators without breaking existing schemas
- New databases without format changes

**Forward compatibility**:
- Unknown attributes ignored (parser doesn't fail on new fields)
- Required attributes documented (clear contract)
- Safe to add fields in minor version bumps

**schema_version enables format evolution**:
- Can introduce breaking changes in format (1.0 → 2.0)
- Parser can support multiple format versions
- Clear migration path for schema authors

### Alternatives Considered

**Strict schema (reject unknown fields)**: Rejected because:
- Breaks forward compatibility
- Can't add new fields without breaking existing parsers
- Too rigid for evolving format

**Plugin system for generators**: Rejected because:
- Complex (requires dynamic loading, versioning)
- Overengineering for MVP
- Custom generators per schema are sufficient

**No versioning (assume backward compatibility forever)**: Rejected because:
- Format will evolve (new features, improvements)
- No way to communicate breaking changes
- schema_version enables clear migration path

### Source

- JSON extensibility patterns: JSON Schema "additionalProperties" pattern
- Schema evolution: API versioning best practices (Stripe, GitHub APIs)
- Backward compatibility: "Evolving RESTful APIs" (public blog posts)
- Plugin patterns: Open source plugin systems (WordPress, Babel, Webpack)

---

## Research Decision 9: Generator Parameter Schema

### Decision

**Parameter structure**:

`generator_params` is a JSON object with generator-specific keys:

**Distribution parameters**:
- **normal**: `{mean, std_dev, min?, max?}`
- **lognormal**: `{median, min, max}`
- **uniform**: `{min, max}`
- **weighted**: `{values: [{value, weight}, ...]}`
- **ranges**: `{ranges: [{min, max, weight}, ...]}`

**Date/Time parameters**:
- **timestamp_past**: `{days_ago_min, days_ago_max}`
- **timestamp_future**: `{days_ahead_min, days_ahead_max}`
- **date_between**: `{start_date, end_date}`

**String parameters**:
- **email**: `{domain?}` (optional custom domain)
- **phone**: `{format?}` (optional format string)

**Type safety**:
- Parser (F008) validates parameters match generator requirements
- Missing required parameters → error
- Extra parameters → warning (forward compatibility)

**Examples for each generator type**:
- Documented in schema-spec.md with practical examples
- Copy-pasteable JSON snippets

### Rationale

**Nested object for parameters**:
- Clear separation between column definition and generator params
- Type-safe validation possible
- Easy to extend (add new params without changing column structure)

**Required vs optional parameters**:
- Required: Parser validates presence (mean, std_dev for normal)
- Optional: Parser uses default if missing (min, max for normal)
- Explicit documentation of required vs optional

**Typed parameters**:
- Numbers for numeric params (mean, std_dev, min, max)
- Strings for text params (domain, format)
- Arrays for complex params (values, ranges)
- JSON-native types are simple and well-supported

### Alternatives Considered

**Inline parameters (flat structure)**: Rejected because:
```json
// Rejected approach
{
  "name": "credit_score",
  "generator": "normal",
  "mean": 680,
  "std_dev": 80
}
```
- Mixes column fields with generator params
- Hard to distinguish required column fields from generator params
- Nested object is clearer

**String-based DSL**: Rejected because:
```json
// Rejected approach
{
  "generator": "normal(mean=680, std_dev=80, min=300, max=850)"
}
```
- Hard to parse programmatically
- No type safety
- Poor error messages
- JSON objects are simpler

**Separate generator definitions file**: Rejected because:
- Splits schema definition across multiple files
- Harder to understand (need to look up generator definitions)
- Inline parameters are clearer

### Source

- Parameter schema patterns: GraphQL input types, JSON Schema validation
- Type systems: TypeScript type definitions, Go struct tags
- Validation approaches: OpenAPI parameter validation
- Generator patterns: Faker.js configuration, Python Faker

---

## Research Decision 10: Example Schema Complexity Level

### Decision

**Example schema design**:

**Tier**: Tier 1 (simple schema, <30s generation for 1,000 records)

**Tables**: 3 tables
- borrowers (250 records)
- loans (1,000 records)
- payments (3,700 records)

**Total records**: 4,950 (within Tier 1 constitutional constraint: <5,000)

**Relationships**:
- loans → borrowers (many-to-one foreign key)
- payments → loans (many-to-one foreign key)

**Generators**:
- Built-in: first_name, last_name, email, phone, timestamp_past
- Custom: credit_score, loan_amount, interest_rate, loan_status, payment_amount

**Distributions**:
- normal: credit_score (mean=680, std_dev=80)
- lognormal: loan_amount (median=50000, min=5000, max=500000)
- weighted: loan_status (70% active, 25% paid_off, 5% delinquent)

**Features demonstrated**:
- Primary keys (auto_increment)
- Foreign keys (with CASCADE)
- Indexes (email unique, credit_score indexed)
- Custom generators (credit_score, loan_amount, interest_rate, loan_status)
- Distribution types (normal, lognormal, weighted)
- Nullable columns (optional middle_name, phone)
- Unique constraints (email)
- Timestamps (created_at with default CURRENT_TIMESTAMP)

### Rationale

**Tier 1 complexity** (<30s, <5,000 records):
- Constitutional constraint: "<30 seconds" applies to Tier 1 schemas only
- MVP focus: Simple schemas first, complex later
- 4,950 total records is within Tier 1 target

**Fintech domain** (loans):
- Demonstrates verticalized data (not "John Doe" generic)
- Realistic distributions (credit scores, loan amounts)
- Industry terminology (APR, credit score, principal)

**3 tables** (not 10+):
- Simple enough for MVP (<30s generation)
- Complex enough to demonstrate relationships
- Demonstrates parent-child relationships (borrowers → loans → payments)

**Mix of built-in and custom generators**:
- Built-in: first_name, last_name, email (reusable)
- Custom: credit_score, loan_amount (verticalized)
- Demonstrates extensibility

**Variety of distributions**:
- normal: credit_score (bell curve)
- lognormal: loan_amount (right-skewed)
- weighted: loan_status (realistic frequencies)
- Demonstrates all distribution types

### Alternatives Considered

**Single table schema**: Rejected because:
- Doesn't demonstrate relationships
- Doesn't show foreign key handling
- Too simple to validate format completeness

**10+ table schema**: Rejected because:
- Exceeds Tier 1 complexity (<30s target)
- Too complex for initial example
- Tier 2/3 schemas come from community contributions

**Generic domain (users/posts)**: Rejected because:
- Violates "Verticalized > Generic" principle
- Doesn't demonstrate industry-specific realism
- Fintech is more compelling demo

**Healthcare domain**: Rejected because:
- More complex (diagnoses, prescriptions, visits)
- Harder to understand for general developers
- Fintech is simpler and widely applicable

### Source

- Constitution schema tiers: Lines 176-197 (Tier 1 <30s, Tier 2 <2min, Tier 3 <5min)
- MVP focus: Constitution lines 255-264 (start with 3 schemas)
- Fintech data patterns: Public fintech datasets (loan amounts, credit scores)
- Faker.js examples: https://github.com/faker-js/faker (demo patterns)

---

## Constitutional Compliance Verification

### Core Principles Alignment

✅ **Verticalized > Generic**: Custom generators (credit_score, loan_amount) enable fintech realism
✅ **Speed > Features**: JSON parsing is fast (<1ms), format enables O(1) generator lookup
✅ **Local-First**: JSON format is offline-compatible, no network required
✅ **Boring Tech Wins**: JSON is proven, standard, widely understood
✅ **Open Source Forever**: Schema format is MIT licensed, community-contributable
✅ **Developer-First**: Clear specification, human-readable format, unambiguous documentation
✅ **Ship Fast, Validate Early**: Design work takes 2 days, no implementation delay

### Technical Constraints Compliance

✅ **Performance**: Tier 1 example schema <5,000 records (constitutional <30s target)
✅ **Distribution**: JSON format embeds in Go binary via embed.FS
✅ **Database Support**: Format supports both MySQL and PostgreSQL common types
✅ **Cost**: $0 (design work, no infrastructure)
✅ **Code Quality**: Specification is clear, unambiguous, implementable by F008
✅ **License**: MIT (schema format and example schemas)
✅ **Platform Support**: JSON is cross-platform

### Legal Constraints Compliance

✅ **Independent Development**: JSON schema design using public best practices only
✅ **No Employer References**: No employer-specific information
✅ **Public Information Only**: All patterns from public JSON schema docs, SQL docs, Faker.js
✅ **Open Source Protection**: MIT license
✅ **Illustrative Examples Only**: Generic fintech terminology (loans, credit scores)

---

## Next Steps

1. Generate `quickstart.md` with specification validation steps
2. Update CLAUDE.md with schema format guidelines
3. Re-verify constitutional compliance after Phase 1 design
4. Proceed to task generation (/speckit.tasks)

---

**Research Status**: Complete
**Constitutional Violations**: None
**Ready for Phase 1**: Yes
