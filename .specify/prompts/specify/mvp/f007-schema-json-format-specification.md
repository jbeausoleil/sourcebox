# Feature Specification Prompt: F007 - Schema JSON Format Specification

## Feature Metadata
- **Feature ID**: F007
- **Name**: Schema JSON Format Specification
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (2 days)
- **Dependencies**: None (design work)

## Constitutional Alignment

### Core Principles
- ✅ **Verticalized > Generic**: Schema format must support industry-specific data structures
- ✅ **Simple > Complex**: JSON format is human-readable and machine-parseable
- ✅ **Developer-First Design**: Clear specification, easy to understand and extend

### Technical Constraints
- ✅ **Extensibility**: Support for future schema types and generators
- ✅ **Validation**: Clear rules for schema correctness

### Development Practices
- ✅ **Ship Fast**: Start with simple format, iterate based on feedback

## User Story
**US-MVP-004**: "As a developer, I want to understand the schema format so I can create new schemas or modify existing ones for my specific use cases."

## Problem Statement
SourceBox needs a standardized format for defining data schemas that support multiple tables, relationships, custom data generators, and verticalized distributions. The format must be:
- **Human-readable**: Developers can read and edit schemas
- **Machine-parseable**: Go code can load and validate schemas
- **Extensible**: Support for future generator types and database features
- **Industry-specific**: Enable realistic, verticalized data generation

Without a clear schema specification, developers cannot contribute new schemas, and the codebase cannot reliably generate data.

## Solution Overview
Define a JSON-based schema format that describes tables, columns, data types, generators (for realistic data), and foreign key relationships. Document the format with a comprehensive specification file (`schemas/schema-spec.md`) and provide an example schema (`schemas/example-schema.json`). The format will support built-in generators (name, email, date) and custom generators for verticalized data (loan amounts, credit scores, diagnoses).

## Detailed Requirements

### Acceptance Criteria
1. **Schema Specification Document Created**: `schemas/schema-spec.md` with complete format documentation
2. **JSON Schema Format Defined**: Clear structure with all required and optional fields
3. **Generator Types Documented**: Built-in and custom generators with parameters
4. **Foreign Key Relationships Supported**: Table relationships can be defined
5. **Schema Versioning Strategy Defined**: How schemas evolve over time
6. **Example Schema Provided**: `schemas/example-schema.json` demonstrating all features
7. **Clear Validation Rules**: What makes a schema valid or invalid

### Technical Specifications

#### Schema JSON Structure

```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "description": "Realistic fintech loan data with borrowers, loans, payments, and credit scores",
  "author": "SourceBox",
  "version": "1.0.0",
  "database_type": ["mysql", "postgres"],
  "metadata": {
    "total_records": 4950,
    "industry": "fintech",
    "tags": ["loans", "credit", "payments"]
  },
  "tables": [
    {
      "name": "borrowers",
      "description": "Loan borrowers with personal information",
      "record_count": 250,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "primary_key": true,
          "auto_increment": true,
          "nullable": false
        },
        {
          "name": "first_name",
          "type": "varchar(100)",
          "generator": "first_name",
          "nullable": false
        },
        {
          "name": "last_name",
          "type": "varchar(100)",
          "generator": "last_name",
          "nullable": false
        },
        {
          "name": "email",
          "type": "varchar(255)",
          "generator": "email",
          "unique": true,
          "nullable": false
        },
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
          },
          "nullable": false
        },
        {
          "name": "created_at",
          "type": "timestamp",
          "generator": "timestamp_past",
          "generator_params": {
            "days_ago_min": 365,
            "days_ago_max": 30
          },
          "default": "CURRENT_TIMESTAMP"
        }
      ],
      "indexes": [
        {
          "name": "idx_email",
          "columns": ["email"],
          "unique": true
        },
        {
          "name": "idx_credit_score",
          "columns": ["credit_score"]
        }
      ]
    },
    {
      "name": "loans",
      "description": "Loan records with amounts, rates, and status",
      "record_count": 1000,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "primary_key": true,
          "auto_increment": true
        },
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {
            "table": "borrowers",
            "column": "id",
            "on_delete": "CASCADE"
          },
          "nullable": false
        },
        {
          "name": "amount",
          "type": "decimal(12,2)",
          "generator": "loan_amount",
          "generator_params": {
            "distribution": "lognormal",
            "median": 50000,
            "min": 5000,
            "max": 500000
          },
          "nullable": false
        },
        {
          "name": "interest_rate",
          "type": "decimal(5,2)",
          "generator": "interest_rate",
          "generator_params": {
            "distribution": "weighted",
            "ranges": [
              {"min": 3.0, "max": 6.0, "weight": 60},
              {"min": 6.0, "max": 10.0, "weight": 30},
              {"min": 10.0, "max": 15.0, "weight": 10}
            ]
          },
          "nullable": false
        },
        {
          "name": "status",
          "type": "varchar(20)",
          "generator": "loan_status",
          "generator_params": {
            "values": [
              {"value": "active", "weight": 70},
              {"value": "paid_off", "weight": 25},
              {"value": "delinquent", "weight": 5}
            ]
          },
          "nullable": false
        },
        {
          "name": "created_at",
          "type": "timestamp",
          "generator": "timestamp_past",
          "default": "CURRENT_TIMESTAMP"
        }
      ],
      "indexes": [
        {
          "name": "idx_borrower",
          "columns": ["borrower_id"]
        },
        {
          "name": "idx_status",
          "columns": ["status"]
        }
      ]
    }
  ],
  "relationships": [
    {
      "from_table": "loans",
      "from_column": "borrower_id",
      "to_table": "borrowers",
      "to_column": "id",
      "relationship_type": "many_to_one",
      "description": "Each loan belongs to one borrower"
    }
  ],
  "generation_order": ["borrowers", "loans", "payments", "credit_scores"],
  "validation_rules": [
    {
      "rule": "loan_amount_credit_correlation",
      "description": "Higher credit scores should correlate with higher loan amounts",
      "enforcement": "soft"
    }
  ]
}
```

#### Supported Data Types

**MySQL/PostgreSQL Common Types**:
- `int`, `bigint`, `smallint`, `tinyint`
- `decimal(p,s)`, `float`, `double`
- `varchar(n)`, `text`, `char(n)`
- `date`, `datetime`, `timestamp`
- `boolean`, `bit`
- `json`, `jsonb` (Postgres)
- `enum('val1','val2',...)`

#### Built-in Generator Types

**Personal Data**:
- `first_name`, `last_name`, `full_name`
- `email`, `phone`, `address`
- `ssn`, `date_of_birth`

**Company Data**:
- `company_name`, `job_title`
- `company_email`, `domain`

**Date/Time**:
- `timestamp_past` - Random timestamp in the past
- `timestamp_future` - Random timestamp in the future
- `date_between` - Date between two dates

**Numeric**:
- `int_range` - Integer within range
- `float_range` - Float within range
- `decimal_range` - Decimal within range

**Custom Generators** (defined per schema):
- Fintech: `credit_score`, `loan_amount`, `interest_rate`, `loan_status`
- Healthcare: `diagnosis_code`, `medication`, `visit_type`, `insurance_provider`
- Retail: `product_name`, `price`, `sku`, `category`

#### Generator Parameters

Generators support various distribution types:
- **uniform**: Evenly distributed values
- **normal**: Bell curve distribution (mean, std_dev)
- **lognormal**: For naturally skewed data (median, min, max)
- **weighted**: Specific values with weights (realistic frequency)

#### Foreign Key Relationships

```json
{
  "foreign_key": {
    "table": "parent_table",
    "column": "parent_column",
    "on_delete": "CASCADE|SET NULL|RESTRICT",
    "on_update": "CASCADE|SET NULL|RESTRICT"
  }
}
```

#### Schema Versioning

Schemas follow semantic versioning (semver):
- **Major (1.0.0 → 2.0.0)**: Breaking changes (table/column removal, type changes)
- **Minor (1.0.0 → 1.1.0)**: New tables or columns (backward compatible)
- **Patch (1.0.0 → 1.0.1)**: Bug fixes, data distribution improvements

### Validation Rules

A valid schema must:
1. Have a unique name (no conflicts with existing schemas)
2. Specify at least one table
3. Each table must have at least one column
4. Each table must have a primary key
5. Foreign keys must reference existing tables and columns
6. Generator names must be valid (built-in or custom)
7. Generator parameters must match generator requirements
8. `generation_order` must include all tables
9. Record counts must be positive integers

### Performance Considerations
- JSON parsing is fast in Go (standard library)
- Schema validation happens once at load time
- Schemas are embedded in binary (no runtime file reads needed)

### Testing Strategy

**Manual Validation**:
1. Create example schema following specification
2. Verify all fields are documented
3. Verify schema is valid JSON
4. Check specification against Go parser requirements (F008)

**Documentation Tests**:
- Example schema should be self-explanatory
- Specification should answer common questions:
  - How do I add a new table?
  - How do I define a foreign key?
  - How do I use custom generators?
  - How do I specify data distributions?

## Dependencies
- **Upstream**: None (this is a design/specification feature)
- **Downstream**:
  - F008 (Schema Parser) implements this specification
  - F010 (Fintech Schema) uses this format
  - F016 (Healthcare Schema) uses this format
  - F018 (Retail Schema) uses this format

## Deliverables
1. **Schema Specification Document**: `schemas/schema-spec.md` (comprehensive documentation)
2. **Example Schema**: `schemas/example-schema.json` (demonstrates all features)
3. **JSON Schema Validation File** (optional): `schemas/schema-schema.json` (JSON Schema for validation)
4. **Generator Documentation**: List of all built-in generators with parameters

## Success Criteria
- ✅ Specification is complete and unambiguous
- ✅ Example schema demonstrates all major features
- ✅ Format supports verticalized data generation
- ✅ Format is extensible for future needs
- ✅ Documentation is clear enough for contributors to add schemas
- ✅ Validation rules are explicitly defined

## Anti-Patterns to Avoid
- ❌ Overly complex format (keep it simple and readable)
- ❌ Database-specific features that lock in to one database (support both MySQL and Postgres)
- ❌ Vague generator definitions (be explicit about parameters)
- ❌ No versioning strategy (schemas will evolve)
- ❌ Poor documentation (developers need clear examples)
- ❌ No validation rules (leads to invalid schemas)

## Implementation Notes
- JSON was chosen over YAML for machine-parseability and Go standard library support
- Format is designed to be extended without breaking existing schemas
- Custom generators allow verticalized data without bloating the core format
- Foreign keys ensure referential integrity
- `generation_order` ensures parent tables are created before children

## TDD Requirements
**Not applicable for specification/design work** - This is documentation and design. However, the specification should be validated by:
1. Creating an example schema that uses all features
2. Validating JSON syntax
3. Ensuring F008 (parser) can implement this specification

## Related Constitution Sections
- **Verticalized > Generic (Principle I)**: Custom generators enable industry-specific realism
- **Simple > Complex**: JSON format is simple, standard, readable
- **Developer-First Design (Principle VI)**: Clear specification enables contributions
- **MVP Mindset (Product Philosophy 1)**: Start with 3 schemas, expand via community
