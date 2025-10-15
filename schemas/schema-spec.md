# Schema JSON Format Specification

**Version**: 1.0
**Last Updated**: 2025-10-15
**Status**: Design Specification

## Overview

SourceBox uses JSON-based schema definitions to describe database schemas, data generators, and relationships. This specification defines the format for creating realistic, verticalized test data across different industries and use cases.

### Why JSON?

JSON was chosen as the schema format for several pragmatic reasons:

1. **Universal Understanding**: Every developer knows JSON. No learning curve, no special tooling required.
2. **Human-Readable**: Schemas can be written and edited in any text editor without specialized IDE support.
3. **Machine-Parsable**: Standard libraries exist in every language, making future integrations trivial.
4. **Validation-Friendly**: JSON Schema exists for structural validation, and custom validators can easily inspect the parsed structure.
5. **Version Control**: Plain text format works seamlessly with Git, enabling clear diffs and collaborative schema development.
6. **Boring Technology**: JSON is proven, stable, and will be supported forever. No risk of format obsolescence.

### Purpose

Schema JSON files serve three critical functions:

**1. Database Schema Definition**
Define tables, columns, data types, constraints, and relationships in a database-agnostic format that can target both MySQL and PostgreSQL (with future expansion to other databases).

**2. Data Generation Rules**
Specify how to generate realistic test data using built-in generators (names, emails, dates) and custom generators (industry-specific patterns like credit scores, diagnosis codes, product SKUs).

**3. Verticalized Realism**
Enable industry-specific data generation through custom generators, distributions, and relationships. A fintech schema generates realistic loan amounts with lognormal distributions. A healthcare schema generates valid diagnosis codes with proper statistical distributions. Generic test data tools cannot achieve this level of domain accuracy.

### How Schemas Enable Verticalized Data

Schemas bridge the gap between generic test data generation and domain-specific realism:

- **Custom Generators**: Define industry-specific data patterns (e.g., `credit_score`, `diagnosis_code`, `sku_format`)
- **Statistical Distributions**: Model real-world data patterns (lognormal for income, weighted for categorical status)
- **Relationship Modeling**: Represent complex business logic (loans belong to borrowers, payments reference loans)
- **Validation Rules**: Enforce domain constraints (credit scores 300-850, valid ICD-10 codes)

A schema is not just a database blueprint—it is a recipe for generating data that looks and behaves like production data.

### Quick Example

Here's what a minimal SourceBox schema looks like:

```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "description": "Realistic fintech loan data with borrowers and loans",
  "author": "SourceBox Contributors",
  "version": "1.0.0",
  "database_type": ["mysql", "postgres"],
  "tables": [
    {
      "name": "borrowers",
      "record_count": 1000,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"]
        },
        {
          "name": "email",
          "type": "varchar(255)",
          "generator": "email",
          "constraints": ["UNIQUE", "NOT NULL"]
        },
        {
          "name": "credit_score",
          "type": "int",
          "generator": "int_range",
          "distribution": {
            "type": "normal",
            "params": {"mean": 680, "std_dev": 80, "min": 300, "max": 850}
          }
        }
      ]
    },
    {
      "name": "loans",
      "record_count": 2500,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"]
        },
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {
            "table": "borrowers",
            "column": "id",
            "on_delete": "CASCADE"
          }
        },
        {
          "name": "loan_amount",
          "type": "decimal(10,2)",
          "generator": "decimal_range",
          "distribution": {
            "type": "lognormal",
            "params": {"median": 15000, "min": 1000, "max": 50000}
          }
        }
      ]
    }
  ],
  "generation_order": ["borrowers", "loans"]
}
```

This schema defines two tables (borrowers and loans) with realistic data generation using statistical distributions. The rest of this specification explains each field in detail.

### Document Structure

This specification is organized into the following sections:

1. **Overview** (this section): Format rationale and purpose
2. **Schema Structure**: Top-level JSON fields and metadata
3. **Table Definitions**: How to define tables, columns, and constraints
4. **Generator Types**: Built-in and custom data generators
5. **Distribution Types**: Statistical distributions for realistic data
6. **Relationships**: Foreign keys and referential integrity
7. **Validation Rules**: Schema validation requirements
8. **Versioning Strategy**: Semantic versioning for schemas
9. **Examples**: Complete schema examples by industry

## Schema Structure

Every SourceBox schema is a JSON object with a specific set of top-level fields. Some fields are required (the schema parser will reject files without them), while others are optional but recommended for discoverability and maintainability.

### Required Fields

#### schema_version (string)

The format version of this specification. Used by the parser to determine which validation rules to apply.

- **Current version**: `"1.0"`
- **Purpose**: Enable backward compatibility if the format evolves
- **Format**: Semantic versioning (`major.minor`)
- **When to change**: Only when SourceBox releases a new format specification

```json
{
  "schema_version": "1.0"
}
```

#### name (string)

Unique identifier for this schema. Used in CLI commands and file naming.

- **Format**: lowercase-kebab-case (e.g., `fintech-loans`, `healthcare-patient-records`)
- **Constraints**: Must be unique across all schemas, no spaces, alphanumeric + hyphens only
- **Purpose**: Human-readable identifier for `sourcebox seed` and `sourcebox list-schemas`

```json
{
  "name": "fintech-loans"
}
```

#### description (string)

Human-readable explanation of what this schema represents.

- **Length**: 1-200 characters recommended
- **Purpose**: Shown in schema listings, helps users discover relevant schemas
- **Style**: Clear, concise, actionable (e.g., "Realistic fintech loan data with borrowers, loans, and payments")

```json
{
  "description": "Realistic fintech loan data with borrowers, loans, and payments"
}
```

#### author (string)

Creator or maintainer of this schema.

- **Format**: Name, organization, or GitHub username
- **Purpose**: Attribution and contact for questions/contributions
- **Examples**: `"SourceBox Contributors"`, `"Jane Doe (@janedoe)"`, `"Acme Corp Data Team"`

```json
{
  "author": "SourceBox Contributors"
}
```

#### version (string)

Content version of this schema (not the format version).

- **Format**: Semantic versioning (`major.minor.patch`)
- **Purpose**: Track changes to table structure, generators, relationships
- **Versioning rules**:
  - **Major** (1.0.0 → 2.0.0): Breaking changes (table removal, type changes)
  - **Minor** (1.0.0 → 1.1.0): Backward-compatible additions (new tables, columns)
  - **Patch** (1.0.0 → 1.0.1): Bug fixes (distribution tweaks, docs)

```json
{
  "version": "1.0.0"
}
```

#### database_type (array of strings)

List of database systems this schema supports.

- **Valid values**: `["mysql", "postgres"]` (MVP), future: `"sqlite"`, `"mssql"`, `"oracle"`
- **Purpose**: Enable database-specific type mapping and validation
- **Constraints**: Must include at least one database type

```json
{
  "database_type": ["mysql", "postgres"]
}
```

### Optional Fields

#### metadata (object)

Free-form metadata for categorization and discoverability.

- **Common fields**:
  - `industry` (string): Industry vertical (e.g., `"fintech"`, `"healthcare"`, `"retail"`)
  - `tags` (array of strings): Searchable keywords (e.g., `["loans", "credit", "payments"]`)
  - `total_records` (integer): Sum of record_count across all tables (helpful for complexity estimation)
- **Purpose**: Enable schema search, filtering, and complexity assessment
- **Extensible**: Add custom fields as needed (e.g., `complexity_tier`, `license`, `created_date`)

```json
{
  "metadata": {
    "industry": "fintech",
    "tags": ["loans", "credit", "borrowers", "payments"],
    "total_records": 4950,
    "complexity_tier": 1
  }
}
```

#### tables (array of objects)

Definitions of database tables, columns, and generators. Required for data generation but optional in the structure (allows for metadata-only schema files).

- **Purpose**: Define database schema and data generation rules
- **Details**: See section 3 (Table Definitions)

```json
{
  "tables": [
    {
      "name": "borrowers",
      "description": "Individual loan borrowers",
      "record_count": 1000,
      "columns": [...]
    }
  ]
}
```

#### relationships (array of objects)

Explicit documentation of foreign key relationships between tables.

- **Purpose**: Human-readable relationship documentation (parser uses inline `foreign_key` in columns)
- **Details**: See section 6 (Relationships)
- **Note**: This is complementary to inline foreign key definitions, not a replacement

```json
{
  "relationships": [
    {
      "from_table": "loans",
      "from_column": "borrower_id",
      "to_table": "borrowers",
      "to_column": "id",
      "relationship_type": "many_to_one",
      "description": "Each loan belongs to one borrower"
    }
  ]
}
```

#### generation_order (array of strings)

Explicit order for table population (parent tables before child tables).

- **Purpose**: Ensure foreign key constraints are satisfied during data generation
- **Constraints**: Must include all table names, parent tables must appear before children
- **Example**: `["borrowers", "loans", "payments"]` (borrowers first, then loans, then payments)

```json
{
  "generation_order": ["borrowers", "loans", "payments"]
}
```

#### validation_rules (array of objects)

Schema-level validation rules for cross-table constraints.

- **Purpose**: Define business logic constraints that span multiple tables
- **Details**: See section 7 (Validation Rules)
- **Examples**: "Total payments cannot exceed loan amount", "Credit score must align with loan approval rate"

```json
{
  "validation_rules": [
    {
      "rule": "payments.sum(amount) <= loans.amount",
      "description": "Total payments cannot exceed original loan amount",
      "severity": "error"
    }
  ]
}
```

### Complete Structure Example

Here is a minimal but complete schema showing all top-level fields:

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
    "tags": ["loans", "credit", "borrowers", "payments"],
    "total_records": 4950,
    "complexity_tier": 1
  },
  "tables": [
    {
      "name": "borrowers",
      "description": "Individual loan borrowers",
      "record_count": 1000,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"]
        },
        {
          "name": "first_name",
          "type": "varchar(100)",
          "generator": "first_name"
        },
        {
          "name": "email",
          "type": "varchar(255)",
          "generator": "email",
          "constraints": ["UNIQUE", "NOT NULL"]
        }
      ]
    },
    {
      "name": "loans",
      "description": "Loan records",
      "record_count": 2500,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"]
        },
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {
            "table": "borrowers",
            "column": "id",
            "on_delete": "CASCADE",
            "on_update": "CASCADE"
          }
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
  "generation_order": ["borrowers", "loans"],
  "validation_rules": [
    {
      "rule": "loans.borrower_id REFERENCES borrowers.id",
      "description": "All loans must reference valid borrowers",
      "severity": "error"
    }
  ]
}
```

### Field Summary

| Field | Required | Type | Purpose |
|-------|----------|------|---------|
| `schema_version` | Yes | string | Format version (currently `"1.0"`) |
| `name` | Yes | string | Unique schema identifier (lowercase-kebab-case) |
| `description` | Yes | string | Human-readable schema purpose |
| `author` | Yes | string | Creator/maintainer attribution |
| `version` | Yes | string | Content version (semver: `major.minor.patch`) |
| `database_type` | Yes | array | Supported databases (`["mysql", "postgres"]`) |
| `metadata` | No | object | Free-form categorization (industry, tags, total_records) |
| `tables` | No | array | Table definitions (see section 3) |
| `relationships` | No | array | Explicit relationship documentation (see section 6) |
| `generation_order` | No | array | Table population order (parent tables first) |
| `validation_rules` | No | array | Schema-level validation rules (see section 7) |

---

## Table Definitions

Tables are the core building blocks of a schema. Each table defines a database table structure, the number of records to generate, and how to populate each column with realistic data.

### Structure

A table definition is a JSON object within the `tables` array. Tables are defined independently but can reference each other through foreign key relationships (see section 6).

```json
{
  "tables": [
    {
      "name": "borrowers",
      "description": "Individual loan borrowers with credit profiles",
      "record_count": 1000,
      "columns": [...],
      "indexes": [...]
    }
  ]
}
```

### Required Fields

#### name (string)

The table name as it will appear in the database.

- **Format**: lowercase with underscores for multi-word names (e.g., `borrowers`, `loan_payments`, `user_sessions`)
- **Constraints**: Must be unique within the schema, must be a valid SQL identifier
- **Purpose**: Used for CREATE TABLE statements, foreign key references, and generation order
- **Examples**: `borrowers`, `loans`, `payments`, `transaction_history`

```json
{
  "name": "borrowers"
}
```

#### record_count (integer)

The number of records (rows) to generate for this table.

- **Constraints**: Must be > 0 (positive integer)
- **Purpose**: Controls dataset size and influences generation time
- **Tier 1 target**: Individual tables should stay under 5,000 records for MVP schemas
- **Considerations**:
  - Child tables typically have higher counts than parent tables (e.g., 1,000 borrowers, 2,500 loans, 5,000 payments)
  - Record count affects foreign key distribution (more child records = more realistic relationships)

```json
{
  "record_count": 1000
}
```

#### columns (array of objects)

Array of column definitions describing the table structure and data generation rules.

- **Constraints**: Must contain at least one column, must include exactly one PRIMARY KEY
- **Purpose**: Defines both database schema (types, constraints) and data generation (generators, distributions)
- **Details**: See section 4 (Column Definitions) for complete column structure
- **Minimum requirement**: Every table must have a primary key column

```json
{
  "columns": [
    {
      "name": "id",
      "type": "int",
      "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"]
    },
    {
      "name": "email",
      "type": "varchar(255)",
      "generator": "email",
      "constraints": ["UNIQUE", "NOT NULL"]
    }
  ]
}
```

### Optional Fields

#### description (string)

Human-readable explanation of the table's purpose and contents.

- **Length**: 1-150 characters recommended
- **Purpose**: Documentation for schema authors and users
- **Style**: Clear, specific, describes what the table represents (not just the table name)
- **Examples**:
  - `"Individual loan borrowers with credit profiles"`
  - `"Historical payment records for all loans"`
  - `"User authentication sessions with IP tracking"`

```json
{
  "description": "Individual loan borrowers with credit profiles"
}
```

#### indexes (array of objects)

Database indexes for query optimization (beyond the primary key).

- **Purpose**: Improve query performance on frequently accessed columns
- **Common use cases**: Foreign key columns, timestamp columns, search fields
- **Note**: MVP implementation may create indexes for foreign keys automatically
- **Structure**: Each index specifies name, columns, and type

```json
{
  "indexes": [
    {
      "name": "idx_borrower_email",
      "columns": ["email"],
      "type": "BTREE",
      "unique": true
    },
    {
      "name": "idx_created_at",
      "columns": ["created_at"],
      "type": "BTREE"
    }
  ]
}
```

### Complete Table Example

Here is a complete table definition showing all fields and realistic column definitions:

```json
{
  "name": "borrowers",
  "description": "Individual loan borrowers with credit profiles and contact information",
  "record_count": 1000,
  "columns": [
    {
      "name": "id",
      "type": "int",
      "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"],
      "description": "Unique borrower identifier"
    },
    {
      "name": "first_name",
      "type": "varchar(100)",
      "generator": "first_name",
      "constraints": ["NOT NULL"],
      "description": "Borrower's legal first name"
    },
    {
      "name": "last_name",
      "type": "varchar(100)",
      "generator": "last_name",
      "constraints": ["NOT NULL"],
      "description": "Borrower's legal last name"
    },
    {
      "name": "email",
      "type": "varchar(255)",
      "generator": "email",
      "constraints": ["UNIQUE", "NOT NULL"],
      "description": "Primary contact email (unique across all borrowers)"
    },
    {
      "name": "phone",
      "type": "varchar(20)",
      "generator": "phone",
      "description": "Primary contact phone number"
    },
    {
      "name": "date_of_birth",
      "type": "date",
      "generator": "date_between",
      "params": {
        "start_date": "1950-01-01",
        "end_date": "2005-12-31"
      },
      "constraints": ["NOT NULL"],
      "description": "Borrower's date of birth (ages 18-75)"
    },
    {
      "name": "credit_score",
      "type": "int",
      "generator": "int_range",
      "distribution": {
        "type": "normal",
        "params": {
          "mean": 680,
          "std_dev": 80,
          "min": 300,
          "max": 850
        }
      },
      "constraints": ["NOT NULL"],
      "description": "FICO credit score (300-850 scale, normally distributed around 680)"
    },
    {
      "name": "created_at",
      "type": "timestamp",
      "generator": "timestamp_past",
      "params": {
        "years_ago": 5
      },
      "constraints": ["NOT NULL", "DEFAULT CURRENT_TIMESTAMP"],
      "description": "Record creation timestamp"
    }
  ],
  "indexes": [
    {
      "name": "idx_borrower_email",
      "columns": ["email"],
      "type": "BTREE",
      "unique": true
    },
    {
      "name": "idx_credit_score",
      "columns": ["credit_score"],
      "type": "BTREE"
    },
    {
      "name": "idx_created_at",
      "columns": ["created_at"],
      "type": "BTREE"
    }
  ]
}
```

### Multiple Tables Example

Here is how multiple tables are defined in a schema, showing parent-child relationships:

```json
{
  "tables": [
    {
      "name": "borrowers",
      "description": "Individual loan borrowers",
      "record_count": 1000,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"]
        },
        {
          "name": "email",
          "type": "varchar(255)",
          "generator": "email",
          "constraints": ["UNIQUE", "NOT NULL"]
        }
      ]
    },
    {
      "name": "loans",
      "description": "Loan records linked to borrowers",
      "record_count": 2500,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"]
        },
        {
          "name": "borrower_id",
          "type": "int",
          "constraints": ["NOT NULL"],
          "foreign_key": {
            "table": "borrowers",
            "column": "id",
            "on_delete": "CASCADE",
            "on_update": "CASCADE"
          },
          "description": "Reference to borrower who took out this loan"
        },
        {
          "name": "loan_amount",
          "type": "decimal(10,2)",
          "generator": "decimal_range",
          "distribution": {
            "type": "lognormal",
            "params": {
              "median": 15000,
              "min": 1000,
              "max": 50000
            }
          },
          "constraints": ["NOT NULL"]
        }
      ],
      "indexes": [
        {
          "name": "idx_borrower_id",
          "columns": ["borrower_id"],
          "type": "BTREE"
        }
      ]
    },
    {
      "name": "payments",
      "description": "Payment records for loans",
      "record_count": 7500,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "constraints": ["PRIMARY KEY", "AUTO_INCREMENT"]
        },
        {
          "name": "loan_id",
          "type": "int",
          "constraints": ["NOT NULL"],
          "foreign_key": {
            "table": "loans",
            "column": "id",
            "on_delete": "CASCADE",
            "on_update": "CASCADE"
          },
          "description": "Reference to loan this payment applies to"
        },
        {
          "name": "payment_amount",
          "type": "decimal(10,2)",
          "generator": "decimal_range",
          "params": {
            "min": 50,
            "max": 2000
          },
          "constraints": ["NOT NULL"]
        }
      ],
      "indexes": [
        {
          "name": "idx_loan_id",
          "columns": ["loan_id"],
          "type": "BTREE"
        }
      ]
    }
  ]
}
```

### Table Naming Conventions

**Lowercase with underscores**:
- Good: `borrowers`, `loan_payments`, `user_sessions`, `transaction_history`
- Bad: `Borrowers`, `loanPayments`, `UserSessions`, `TransactionHistory`

**Plural for data tables**:
- Good: `borrowers`, `loans`, `payments`
- Bad: `borrower`, `loan`, `payment`

**Singular for lookup/reference tables** (optional convention):
- Acceptable: `status`, `country`, `currency`
- Also acceptable: `statuses`, `countries`, `currencies`

**Descriptive but concise**:
- Good: `loan_payments`, `user_sessions`, `auth_tokens`
- Bad: `lp`, `us`, `tokens` (too cryptic)
- Bad: `loan_payment_transaction_records` (too verbose)

### Table Relationships and Generation Order

When defining multiple tables with foreign key relationships:

1. **Define parent tables first** in the `tables` array (though this is not required—the `generation_order` field controls population order)
2. **Specify foreign keys** in the child table's column definition (inline `foreign_key` object)
3. **Set generation_order** to ensure parent tables are populated before child tables
4. **Document relationships** in the `relationships` array for human readability

**Example generation flow**:
```json
{
  "generation_order": ["borrowers", "loans", "payments"]
}
```

This ensures:
1. 1,000 borrowers are created first (with IDs 1-1000)
2. 2,500 loans are created next, each referencing a valid borrower_id (1-1000)
3. 7,500 payments are created last, each referencing a valid loan_id

### Field Summary

| Field | Required | Type | Purpose |
|-------|----------|------|---------|
| `name` | Yes | string | Table name (lowercase_with_underscores) |
| `record_count` | Yes | integer | Number of records to generate (must be > 0) |
| `columns` | Yes | array | Column definitions (at least 1, exactly 1 PRIMARY KEY) |
| `description` | No | string | Human-readable table purpose |
| `indexes` | No | array | Database indexes for query optimization |

---

## Column Definitions

Columns define the structure of a database table and the rules for generating data to populate it. Each column specifies a SQL data type, optional constraints, and how to generate realistic values.

### Structure

A column definition is a JSON object within the `columns` array of a table. Columns combine database schema metadata (type, constraints) with data generation instructions (generator, distribution).

```json
{
  "columns": [
    {
      "name": "email",
      "type": "varchar(255)",
      "nullable": false,
      "primary_key": false,
      "unique": true,
      "default": null,
      "generator": "email",
      "generator_params": {},
      "foreign_key": null,
      "description": "Primary contact email"
    }
  ]
}
```

### Required Fields

#### name (string)

The column name as it will appear in the database.

- **Format**: lowercase with underscores for multi-word names (e.g., `email`, `first_name`, `created_at`)
- **Constraints**: Must be unique within the table, must be a valid SQL identifier
- **Purpose**: Used in CREATE TABLE statements, foreign key references, and data generation
- **Naming conventions**: Use descriptive names (avoid abbreviations like `fn`, `em`, `dt`)

```json
{
  "name": "email"
}
```

#### type (string)

SQL data type for this column.

- **Format**: Standard SQL type with optional size/precision (e.g., `int`, `varchar(255)`, `decimal(10,2)`)
- **Constraints**: Must be from the supported data type list (see Data Types section below)
- **Purpose**: Determines database column type and influences data generation constraints
- **Portability**: Types are validated against MySQL/PostgreSQL common subset

```json
{
  "type": "varchar(255)"
}
```

### Optional Fields

#### nullable (boolean)

Whether this column allows NULL values.

- **Default**: `false` (columns are NOT NULL by default)
- **Purpose**: Controls NULL constraint in database schema
- **Best practice**: Set to `true` only when NULL is semantically valid (e.g., optional fields like `middle_name`, `secondary_phone`)
- **Generator interaction**: If `nullable: true`, generator may randomly produce NULL values based on a configured probability

```json
{
  "nullable": true
}
```

**Examples**:

```json
{
  "name": "middle_name",
  "type": "varchar(100)",
  "nullable": true,
  "generator": "first_name",
  "description": "Optional middle name (30% will be NULL)"
}
```

```json
{
  "name": "email",
  "type": "varchar(255)",
  "nullable": false,
  "generator": "email",
  "description": "Required email address"
}
```

#### primary_key (boolean)

Whether this column is the primary key for the table.

- **Default**: `false`
- **Constraints**: Exactly one column per table must have `primary_key: true`
- **Purpose**: Marks the column as the primary key in CREATE TABLE statements
- **Auto-increment**: Primary keys are typically auto-incrementing integers (though UUIDs are also valid)
- **Generator interaction**: Primary key columns typically use auto-increment (no generator needed)

```json
{
  "primary_key": true
}
```

**Examples**:

```json
{
  "name": "id",
  "type": "int",
  "primary_key": true,
  "nullable": false,
  "description": "Auto-incrementing primary key"
}
```

```json
{
  "name": "uuid",
  "type": "varchar(36)",
  "primary_key": true,
  "nullable": false,
  "generator": "uuid",
  "description": "UUID primary key"
}
```

#### unique (boolean)

Whether this column must have unique values across all records.

- **Default**: `false`
- **Purpose**: Adds a UNIQUE constraint in the database schema
- **Use cases**: Email addresses, usernames, external IDs, natural keys
- **Generator interaction**: Generator must ensure uniqueness (e.g., email generator appends incremental numbers if needed)

```json
{
  "unique": true
}
```

**Examples**:

```json
{
  "name": "email",
  "type": "varchar(255)",
  "unique": true,
  "nullable": false,
  "generator": "email",
  "description": "Unique email address across all users"
}
```

```json
{
  "name": "username",
  "type": "varchar(50)",
  "unique": true,
  "nullable": false,
  "generator": "username",
  "description": "Unique username for login"
}
```

#### default (any)

Default value for this column if no value is provided during insertion.

- **Default**: `null` (no default)
- **Type**: Must match the column's data type (string for varchar, number for int/decimal, boolean for boolean)
- **Purpose**: Sets the DEFAULT constraint in the database schema
- **Common use cases**: Status flags, timestamps, boolean flags
- **Special values**: `CURRENT_TIMESTAMP` for timestamp columns

```json
{
  "default": "active"
}
```

**Examples**:

```json
{
  "name": "status",
  "type": "varchar(20)",
  "default": "active",
  "nullable": false,
  "description": "Loan status (defaults to active)"
}
```

```json
{
  "name": "is_verified",
  "type": "boolean",
  "default": false,
  "nullable": false,
  "description": "Email verification status (defaults to false)"
}
```

```json
{
  "name": "created_at",
  "type": "timestamp",
  "default": "CURRENT_TIMESTAMP",
  "nullable": false,
  "description": "Record creation timestamp"
}
```

#### generator (string)

Name of the data generator to use for populating this column.

- **Default**: `null` (no generation—use default value or auto-increment)
- **Purpose**: Specifies which built-in or custom generator creates values for this column
- **Generator types**:
  - **Built-in generators**: `first_name`, `last_name`, `email`, `phone`, `address`, `date_between`, `int_range`, `decimal_range`, etc. (see section 5)
  - **Custom generators**: Schema-specific generators like `credit_score`, `loan_amount`, `diagnosis_code` (defined in schema metadata)
- **When to omit**: Primary keys (auto-increment), columns with static defaults, foreign keys (populated from parent table)

```json
{
  "generator": "email"
}
```

**Examples**:

```json
{
  "name": "first_name",
  "type": "varchar(100)",
  "generator": "first_name",
  "description": "Generated using built-in first_name generator"
}
```

```json
{
  "name": "credit_score",
  "type": "int",
  "generator": "int_range",
  "generator_params": {
    "min": 300,
    "max": 850
  },
  "description": "Generated using int_range with custom bounds"
}
```

```json
{
  "name": "id",
  "type": "int",
  "primary_key": true,
  "description": "No generator needed—auto-increments"
}
```

#### generator_params (object)

Parameters for configuring the generator's behavior.

- **Default**: `{}` (empty object—use generator defaults)
- **Purpose**: Pass configuration to the generator (e.g., min/max for ranges, date ranges, distribution parameters)
- **Structure**: Free-form object—schema depends on the generator
- **Common parameters**:
  - `min`, `max` for numeric/date ranges
  - `start_date`, `end_date` for date ranges
  - `years_ago` for relative dates
  - `distribution` for statistical distributions (see section 6)
- **Validation**: Parser validates params against generator requirements

```json
{
  "generator_params": {
    "min": 1000,
    "max": 50000
  }
}
```

**Examples**:

```json
{
  "name": "loan_amount",
  "type": "decimal(10,2)",
  "generator": "decimal_range",
  "generator_params": {
    "min": 1000,
    "max": 50000,
    "distribution": {
      "type": "lognormal",
      "params": {
        "median": 15000,
        "min": 1000,
        "max": 50000
      }
    }
  },
  "description": "Loan amount with lognormal distribution"
}
```

```json
{
  "name": "date_of_birth",
  "type": "date",
  "generator": "date_between",
  "generator_params": {
    "start_date": "1950-01-01",
    "end_date": "2005-12-31"
  },
  "description": "Birth date (ages 18-75)"
}
```

```json
{
  "name": "created_at",
  "type": "timestamp",
  "generator": "timestamp_past",
  "generator_params": {
    "years_ago": 5
  },
  "description": "Record created within last 5 years"
}
```

#### foreign_key (object)

Foreign key relationship definition (inline representation).

- **Default**: `null` (not a foreign key)
- **Purpose**: Defines referential integrity constraints and enables relational data generation
- **Structure**: Object with `table`, `column`, `on_delete`, `on_update` fields
- **Parser usage**: Parser uses this field to generate foreign key constraints and populate values from parent table
- **Dual representation**: This inline definition is used by the parser; the `relationships` array (section 6) is for human documentation

**Foreign key object structure**:

| Field | Required | Type | Purpose |
|-------|----------|------|---------|
| `table` | Yes | string | Parent table name (must exist in schema) |
| `column` | Yes | string | Parent column name (must be primary key or unique) |
| `on_delete` | No | string | Action on parent deletion (`CASCADE`, `SET NULL`, `RESTRICT`) |
| `on_update` | No | string | Action on parent update (`CASCADE`, `SET NULL`, `RESTRICT`) |

```json
{
  "foreign_key": {
    "table": "borrowers",
    "column": "id",
    "on_delete": "CASCADE",
    "on_update": "CASCADE"
  }
}
```

**Examples**:

```json
{
  "name": "borrower_id",
  "type": "int",
  "nullable": false,
  "foreign_key": {
    "table": "borrowers",
    "column": "id",
    "on_delete": "CASCADE",
    "on_update": "CASCADE"
  },
  "description": "Reference to parent borrower (cascading deletes)"
}
```

```json
{
  "name": "user_id",
  "type": "int",
  "nullable": true,
  "foreign_key": {
    "table": "users",
    "column": "id",
    "on_delete": "SET NULL",
    "on_update": "CASCADE"
  },
  "description": "Optional user reference (set to NULL if user deleted)"
}
```

```json
{
  "name": "country_code",
  "type": "varchar(2)",
  "nullable": false,
  "foreign_key": {
    "table": "countries",
    "column": "code",
    "on_delete": "RESTRICT",
    "on_update": "CASCADE"
  },
  "description": "ISO country code (prevent country deletion if referenced)"
}
```

**Referential integrity actions**:

- **CASCADE**: Propagate parent changes/deletions to child records (e.g., delete all loans when borrower deleted)
- **SET NULL**: Set foreign key to NULL when parent deleted (requires `nullable: true`)
- **RESTRICT**: Prevent parent deletion if child records exist (enforce strict referential integrity)
- **NO ACTION** / **SET DEFAULT**: Supported by some databases (not part of MVP)

#### description (string)

Human-readable description of the column's purpose and contents.

- **Default**: `null` (no description)
- **Length**: 1-150 characters recommended
- **Purpose**: Documentation for schema authors and users
- **Style**: Clear, specific, mentions constraints and semantics
- **Include**: Nullability notes, value ranges, units, relationships

```json
{
  "description": "Borrower's credit score (300-850 scale, normally distributed around 680)"
}
```

**Examples**:

```json
{
  "name": "email",
  "type": "varchar(255)",
  "description": "Primary contact email (unique across all borrowers)"
}
```

```json
{
  "name": "loan_amount",
  "type": "decimal(10,2)",
  "description": "Loan principal amount in USD (1000-50000 range, lognormal distribution)"
}
```

```json
{
  "name": "is_active",
  "type": "boolean",
  "description": "Active status flag (defaults to true, set to false when loan is closed)"
}
```

### Data Types

SourceBox supports the MySQL/PostgreSQL common subset of data types to ensure cross-database portability. Database-specific types (e.g., PostgreSQL's `jsonb`, MySQL's `enum`) are annotated with compatibility notes.

#### Integer Types

| Type | MySQL | PostgreSQL | Range | Use Case |
|------|-------|------------|-------|----------|
| `tinyint` | Yes | Yes (smallint) | -128 to 127 | Small numeric values, flags |
| `smallint` | Yes | Yes | -32,768 to 32,767 | Small IDs, counts |
| `int` | Yes | Yes | -2.1B to 2.1B | Primary keys, standard IDs |
| `bigint` | Yes | Yes | -9.2E18 to 9.2E18 | Large IDs, timestamps |

**Examples**:

```json
{
  "name": "id",
  "type": "int",
  "primary_key": true
}
```

```json
{
  "name": "user_count",
  "type": "smallint",
  "default": 0
}
```

#### Decimal Types

| Type | MySQL | PostgreSQL | Precision | Use Case |
|------|-------|------------|-----------|----------|
| `decimal(p,s)` | Yes | Yes | Exact | Financial data, currency |
| `float` | Yes | Yes | Approximate | Measurements, percentages |
| `double` | Yes | Yes (double precision) | Approximate | Scientific data, coordinates |

**Examples**:

```json
{
  "name": "loan_amount",
  "type": "decimal(10,2)",
  "generator": "decimal_range",
  "description": "Loan amount in USD (exact precision)"
}
```

```json
{
  "name": "interest_rate",
  "type": "float",
  "generator": "float_range",
  "description": "Annual interest rate (percentage)"
}
```

#### String Types

| Type | MySQL | PostgreSQL | Max Length | Use Case |
|------|-------|------------|------------|----------|
| `char(n)` | Yes | Yes | Fixed-length | Codes, fixed-width IDs |
| `varchar(n)` | Yes | Yes | Variable, up to n | Names, emails, short text |
| `text` | Yes | Yes | Unlimited | Long text, descriptions, JSON |

**Examples**:

```json
{
  "name": "email",
  "type": "varchar(255)",
  "generator": "email",
  "unique": true
}
```

```json
{
  "name": "country_code",
  "type": "char(2)",
  "generator": "country_code",
  "description": "ISO 3166-1 alpha-2 code"
}
```

```json
{
  "name": "description",
  "type": "text",
  "generator": "lorem_paragraph",
  "description": "Long-form text description"
}
```

#### Date/Time Types

| Type | MySQL | PostgreSQL | Format | Use Case |
|------|-------|------------|--------|----------|
| `date` | Yes | Yes | YYYY-MM-DD | Birth dates, due dates |
| `datetime` | Yes | Yes (timestamp) | YYYY-MM-DD HH:MM:SS | Event timestamps |
| `timestamp` | Yes | Yes | YYYY-MM-DD HH:MM:SS | Record creation/update |

**Examples**:

```json
{
  "name": "date_of_birth",
  "type": "date",
  "generator": "date_between",
  "generator_params": {
    "start_date": "1950-01-01",
    "end_date": "2005-12-31"
  }
}
```

```json
{
  "name": "created_at",
  "type": "timestamp",
  "default": "CURRENT_TIMESTAMP",
  "nullable": false
}
```

```json
{
  "name": "updated_at",
  "type": "datetime",
  "generator": "timestamp_past",
  "generator_params": {
    "years_ago": 1
  }
}
```

#### Boolean Type

| Type | MySQL | PostgreSQL | Values | Use Case |
|------|-------|------------|--------|----------|
| `boolean` | Yes (tinyint(1)) | Yes | TRUE/FALSE | Flags, status indicators |

**Note**: MySQL stores boolean as `tinyint(1)` internally but accepts `TRUE`/`FALSE` syntax.

**Examples**:

```json
{
  "name": "is_active",
  "type": "boolean",
  "default": true,
  "description": "Active status flag"
}
```

```json
{
  "name": "is_verified",
  "type": "boolean",
  "default": false,
  "generator": "weighted_boolean",
  "generator_params": {
    "true_weight": 0.8
  },
  "description": "Email verification status (80% verified)"
}
```

#### JSON Type

| Type | MySQL | PostgreSQL | Notes |
|------|-------|------------|-------|
| `json` | Yes (5.7+) | Yes | Standard JSON storage |
| `jsonb` | No | Yes (binary) | PostgreSQL-only (binary JSON, faster queries) |

**Examples**:

```json
{
  "name": "metadata",
  "type": "json",
  "generator": "json_object",
  "generator_params": {
    "schema": {
      "user_agent": "string",
      "ip_address": "string",
      "referrer": "string"
    }
  },
  "description": "Session metadata as JSON object"
}
```

**Compatibility note**: Use `json` for cross-database schemas. Use `jsonb` only if targeting PostgreSQL exclusively and need binary JSON performance.

#### Enum Type

| Type | MySQL | PostgreSQL | Notes |
|------|-------|------------|-------|
| `enum('val1','val2')` | Yes | Yes (CREATE TYPE) | Predefined value list |

**Examples**:

```json
{
  "name": "status",
  "type": "enum('active','paid','delinquent','defaulted')",
  "generator": "enum",
  "generator_params": {
    "values": [
      {"value": "active", "weight": 0.70},
      {"value": "paid", "weight": 0.25},
      {"value": "delinquent", "weight": 0.04},
      {"value": "defaulted", "weight": 0.01}
    ]
  },
  "description": "Loan status (weighted distribution)"
}
```

**Compatibility note**: MySQL uses `ENUM('val1','val2')` syntax. PostgreSQL requires `CREATE TYPE` statements (parser handles conversion).

### Complete Column Examples

#### Primary Key Column

```json
{
  "name": "id",
  "type": "int",
  "primary_key": true,
  "nullable": false,
  "description": "Auto-incrementing primary key"
}
```

#### Generated Text Column

```json
{
  "name": "first_name",
  "type": "varchar(100)",
  "nullable": false,
  "generator": "first_name",
  "description": "Borrower's legal first name"
}
```

#### Unique Email Column

```json
{
  "name": "email",
  "type": "varchar(255)",
  "nullable": false,
  "unique": true,
  "generator": "email",
  "description": "Primary contact email (unique across all borrowers)"
}
```

#### Foreign Key Column

```json
{
  "name": "borrower_id",
  "type": "int",
  "nullable": false,
  "foreign_key": {
    "table": "borrowers",
    "column": "id",
    "on_delete": "CASCADE",
    "on_update": "CASCADE"
  },
  "description": "Reference to parent borrower (cascading deletes)"
}
```

#### Generated Numeric Column with Distribution

```json
{
  "name": "credit_score",
  "type": "int",
  "nullable": false,
  "generator": "int_range",
  "generator_params": {
    "min": 300,
    "max": 850,
    "distribution": {
      "type": "normal",
      "params": {
        "mean": 680,
        "std_dev": 80,
        "min": 300,
        "max": 850
      }
    }
  },
  "description": "FICO credit score (300-850 scale, normally distributed around 680)"
}
```

#### Nullable Optional Column

```json
{
  "name": "middle_name",
  "type": "varchar(100)",
  "nullable": true,
  "generator": "first_name",
  "generator_params": {
    "null_probability": 0.3
  },
  "description": "Optional middle name (30% will be NULL)"
}
```

#### Timestamp Column with Default

```json
{
  "name": "created_at",
  "type": "timestamp",
  "nullable": false,
  "default": "CURRENT_TIMESTAMP",
  "description": "Record creation timestamp (defaults to current time)"
}
```

#### Generated Date Column with Range

```json
{
  "name": "date_of_birth",
  "type": "date",
  "nullable": false,
  "generator": "date_between",
  "generator_params": {
    "start_date": "1950-01-01",
    "end_date": "2005-12-31"
  },
  "description": "Borrower's date of birth (ages 18-75)"
}
```

#### Decimal Column with Lognormal Distribution

```json
{
  "name": "loan_amount",
  "type": "decimal(10,2)",
  "nullable": false,
  "generator": "decimal_range",
  "generator_params": {
    "min": 1000,
    "max": 50000,
    "distribution": {
      "type": "lognormal",
      "params": {
        "median": 15000,
        "min": 1000,
        "max": 50000
      }
    }
  },
  "description": "Loan principal amount in USD (lognormal distribution, realistic skew)"
}
```

#### Enum Column with Weighted Distribution

```json
{
  "name": "loan_status",
  "type": "enum('active','paid','delinquent','defaulted')",
  "nullable": false,
  "default": "active",
  "generator": "enum",
  "generator_params": {
    "values": [
      {"value": "active", "weight": 0.70},
      {"value": "paid", "weight": 0.25},
      {"value": "delinquent", "weight": 0.04},
      {"value": "defaulted", "weight": 0.01}
    ]
  },
  "description": "Current loan status (70% active, 25% paid, 4% delinquent, 1% defaulted)"
}
```

#### Boolean Column with Default

```json
{
  "name": "is_verified",
  "type": "boolean",
  "nullable": false,
  "default": false,
  "description": "Email verification status (defaults to false)"
}
```

### Field Summary

| Field | Required | Type | Default | Purpose |
|-------|----------|------|---------|---------|
| `name` | Yes | string | N/A | Column name (lowercase_with_underscores) |
| `type` | Yes | string | N/A | SQL data type from supported list |
| `nullable` | No | boolean | `false` | Allow NULL values |
| `primary_key` | No | boolean | `false` | Mark as primary key (exactly one per table) |
| `unique` | No | boolean | `false` | Enforce unique constraint |
| `default` | No | any | `null` | Default value for column |
| `generator` | No | string | `null` | Data generator name (built-in or custom) |
| `generator_params` | No | object | `{}` | Generator configuration parameters |
| `foreign_key` | No | object | `null` | Foreign key relationship definition |
| `description` | No | string | `null` | Human-readable column description |

### Column Definition Best Practices

**1. Be explicit about nullability**:
- Set `nullable: false` for required fields
- Set `nullable: true` only when NULL is semantically valid
- Document what NULL means (e.g., "NULL if middle name not provided")

**2. Use meaningful names**:
- Good: `email`, `first_name`, `created_at`, `borrower_id`
- Bad: `em`, `fn`, `dt`, `bid`

**3. Match types to data semantics**:
- Use `decimal(p,s)` for currency (exact precision)
- Use `float` for measurements (approximate precision)
- Use `varchar(n)` with appropriate length (not always 255)
- Use `text` for unlimited-length content

**4. Document generator parameters**:
- Explain distribution choices in `description`
- Use realistic min/max bounds
- Leverage distributions for domain accuracy (normal for credit scores, lognormal for loan amounts)

**5. Define foreign keys inline**:
- Always specify `on_delete` and `on_update` actions
- Use `CASCADE` for tight parent-child relationships
- Use `SET NULL` for optional references (requires `nullable: true`)
- Use `RESTRICT` to prevent orphaned references

**6. Provide descriptions**:
- Document value ranges, units, constraints
- Explain relationships to other columns
- Note any special generation logic

---

## Validation Rules

Schema validation ensures that schema JSON files are structurally correct, semantically valid, and will generate data that satisfies database constraints. This section documents all validation rules enforced by the SourceBox schema parser (F008), enabling schema authors to self-validate their schemas and F008 implementers to build comprehensive validators.

### Validation Philosophy

SourceBox validation follows a **fail-fast, fail-loud** philosophy:

- **Fail-fast**: Invalid schemas are rejected immediately during parsing, before any data generation begins
- **Fail-loud**: Error messages are clear, actionable, and guide schema authors toward fixes
- **Comprehensive**: All structural, semantic, and referential integrity constraints are validated
- **Implementable**: Every rule is testable, unambiguous, and can be mechanically verified

Validation occurs in multiple phases:

1. **Structural validation**: JSON syntax, required fields, data types
2. **Schema-level validation**: Top-level field constraints and completeness
3. **Table-level validation**: Table structure and primary key requirements
4. **Column-level validation**: Column definitions, data types, and generators
5. **Relationship validation**: Foreign key integrity and referential actions
6. **Generation order validation**: Dependency resolution and circular dependency detection
7. **Cross-validation**: Consistency across tables, relationships, and generation order

### Validation Categories

This section is organized into the following categories:

1. **Schema-Level Validation**: Top-level field validation (name, version, database_type, etc.)
2. **Table-Level Validation**: Table structure and completeness
3. **Column-Level Validation**: Column definitions and data types
4. **Relationship-Level Validation**: Foreign key integrity
5. **Generation Order Validation**: Dependency ordering and circular dependency detection
6. **Edge Cases**: Common failure scenarios and how to detect them
7. **Error Message Guidance**: Standards for helpful error messages

---

### Schema-Level Validation

Schema-level validation ensures that the top-level structure of the schema JSON is complete and correct.

#### V-S001: Required Fields Must Be Present

**Rule**: All required top-level fields must exist in the schema JSON.

**Required fields**:
- `schema_version` (string)
- `name` (string)
- `description` (string)
- `author` (string)
- `version` (string)
- `database_type` (array of strings)

**Validation logic**:
```
FOR EACH required_field IN [schema_version, name, description, author, version, database_type]:
  IF required_field NOT IN schema:
    RAISE ERROR "Missing required field: {required_field}"
```

**Examples**:

**Valid**:
```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "description": "Realistic fintech loan data",
  "author": "SourceBox Contributors",
  "version": "1.0.0",
  "database_type": ["mysql", "postgres"]
}
```

**Invalid** (missing `author`):
```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "description": "Realistic fintech loan data",
  "version": "1.0.0",
  "database_type": ["mysql", "postgres"]
}
```

**Error message**:
```
Missing required field: author
```

**F008 Implementation Note**: Check for field existence before validating field values.

---

#### V-S002: Schema Name Must Be Valid

**Rule**: The `name` field must:
- Be non-empty
- Use lowercase-kebab-case format
- Contain only alphanumeric characters and hyphens
- Not start or end with a hyphen
- Be unique across all schemas in the repository (enforced at repository level)

**Validation logic**:
```
IF schema.name IS EMPTY:
  RAISE ERROR "Schema name cannot be empty"

IF NOT MATCHES(schema.name, "^[a-z0-9]+(-[a-z0-9]+)*$"):
  RAISE ERROR "Schema name must use lowercase-kebab-case format (e.g., 'fintech-loans')"
```

**Examples**:

**Valid**:
- `"fintech-loans"`
- `"healthcare-patients"`
- `"retail-orders"`
- `"ecommerce-catalog-v2"`

**Invalid**:
- `""` (empty)
- `"Fintech-Loans"` (uppercase)
- `"fintech_loans"` (underscore)
- `"-fintech-loans"` (starts with hyphen)
- `"fintech-loans-"` (ends with hyphen)
- `"fintech--loans"` (double hyphen)

**Error messages**:
```
Schema name cannot be empty
Schema name must use lowercase-kebab-case format (e.g., 'fintech-loans')
```

---

#### V-S003: Schema Version Must Be Valid Semantic Version

**Rule**: The `version` field must follow semantic versioning format: `major.minor.patch` where each component is a non-negative integer.

**Validation logic**:
```
IF NOT MATCHES(schema.version, "^[0-9]+\\.[0-9]+\\.[0-9]+$"):
  RAISE ERROR "Schema version must follow semantic versioning (e.g., '1.0.0')"
```

**Examples**:

**Valid**:
- `"1.0.0"`
- `"2.1.3"`
- `"0.0.1"`
- `"10.20.30"`

**Invalid**:
- `"1.0"` (missing patch)
- `"1"` (missing minor and patch)
- `"v1.0.0"` (prefix not allowed)
- `"1.0.0-beta"` (pre-release suffix not supported in MVP)
- `"1.0.0.0"` (too many components)

**Error message**:
```
Schema version must follow semantic versioning (e.g., '1.0.0')
```

---

#### V-S004: Schema Format Version Must Be Valid

**Rule**: The `schema_version` field must:
- Be a string
- Match the current supported format version (`"1.0"` for MVP)
- Follow `major.minor` format

**Validation logic**:
```
IF schema.schema_version != "1.0":
  RAISE ERROR "Unsupported schema_version: {schema.schema_version}. Parser supports: 1.0"
```

**Examples**:

**Valid**:
- `"1.0"`

**Invalid** (for MVP parser):
- `"2.0"` (future format version)
- `"1.1"` (future format version)
- `"1"` (missing minor version)
- `1.0` (number instead of string)

**Error message**:
```
Unsupported schema_version: 2.0. Parser supports: 1.0
```

**F008 Implementation Note**: This enables forward compatibility—new parsers can support multiple schema versions.

---

#### V-S005: Database Type Must Be Non-Empty and Valid

**Rule**: The `database_type` field must:
- Be a non-empty array
- Contain at least one database type
- Contain only valid database type strings: `"mysql"`, `"postgres"` (MVP)
- Not contain duplicates

**Validation logic**:
```
IF database_type IS EMPTY:
  RAISE ERROR "database_type must contain at least one database type"

FOR EACH db IN database_type:
  IF db NOT IN ["mysql", "postgres"]:
    RAISE ERROR "Invalid database_type: {db}. Supported: mysql, postgres"

IF HAS_DUPLICATES(database_type):
  RAISE ERROR "database_type contains duplicates"
```

**Examples**:

**Valid**:
- `["mysql"]`
- `["postgres"]`
- `["mysql", "postgres"]`

**Invalid**:
- `[]` (empty array)
- `["sqlite"]` (unsupported database)
- `["mysql", "mysql"]` (duplicate)
- `["MySQL"]` (incorrect case)

**Error messages**:
```
database_type must contain at least one database type
Invalid database_type: sqlite. Supported: mysql, postgres
database_type contains duplicates
```

---

#### V-S006: Schema Must Have At Least One Table

**Rule**: If the `tables` field is present, it must contain at least one table definition.

**Validation logic**:
```
IF "tables" IN schema:
  IF schema.tables IS EMPTY:
    RAISE ERROR "Schema must define at least one table"
```

**Examples**:

**Valid**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "record_count": 1000,
      "columns": [...]
    }
  ]
}
```

**Invalid**:
```json
{
  "tables": []
}
```

**Error message**:
```
Schema must define at least one table
```

**F008 Implementation Note**: Empty `tables` array is invalid if `tables` is present. Omitting `tables` entirely is valid (metadata-only schema).

---

#### V-S007: Generation Order Must Include All Tables

**Rule**: If the `generation_order` field is present, it must include all table names defined in `tables`, with no duplicates or extra entries.

**Validation logic**:
```
IF "generation_order" IN schema:
  defined_tables = {table.name FOR table IN schema.tables}
  ordered_tables = SET(schema.generation_order)

  IF ordered_tables != defined_tables:
    missing = defined_tables - ordered_tables
    extra = ordered_tables - defined_tables

    IF missing IS NOT EMPTY:
      RAISE ERROR "Tables missing from generation_order: {missing}"

    IF extra IS NOT EMPTY:
      RAISE ERROR "Unknown tables in generation_order: {extra}"
```

**Examples**:

**Valid**:
```json
{
  "tables": [
    {"name": "borrowers", ...},
    {"name": "loans", ...}
  ],
  "generation_order": ["borrowers", "loans"]
}
```

**Invalid** (missing table):
```json
{
  "tables": [
    {"name": "borrowers", ...},
    {"name": "loans", ...}
  ],
  "generation_order": ["borrowers"]
}
```

**Invalid** (extra table):
```json
{
  "tables": [
    {"name": "borrowers", ...}
  ],
  "generation_order": ["borrowers", "payments"]
}
```

**Error messages**:
```
Tables missing from generation_order: ['loans']
Unknown tables in generation_order: ['payments']
```

---

### Table-Level Validation

Table-level validation ensures that each table definition is structurally correct and contains all required elements.

#### V-T001: Table Name Must Be Valid

**Rule**: Each table's `name` field must:
- Be non-empty
- Use lowercase with underscores for multi-word names
- Contain only lowercase letters, numbers, and underscores
- Be a valid SQL identifier
- Be unique across all tables in the schema

**Validation logic**:
```
FOR EACH table IN schema.tables:
  IF table.name IS EMPTY:
    RAISE ERROR "Table name cannot be empty"

  IF NOT MATCHES(table.name, "^[a-z][a-z0-9_]*$"):
    RAISE ERROR "Table '{table.name}' uses invalid format. Use lowercase_with_underscores"

  IF table.name IN seen_table_names:
    RAISE ERROR "Duplicate table name: {table.name}"

  seen_table_names.ADD(table.name)
```

**Examples**:

**Valid**:
- `"borrowers"`
- `"loan_payments"`
- `"user_sessions"`
- `"transaction_history_2024"`

**Invalid**:
- `""` (empty)
- `"Borrowers"` (uppercase)
- `"loan-payments"` (hyphen)
- `"1_table"` (starts with number)
- `"user sessions"` (space)

**Error messages**:
```
Table name cannot be empty
Table 'Borrowers' uses invalid format. Use lowercase_with_underscores
Duplicate table name: borrowers
```

---

#### V-T002: Record Count Must Be Positive

**Rule**: Each table's `record_count` field must be a positive integer (> 0).

**Validation logic**:
```
FOR EACH table IN schema.tables:
  IF table.record_count <= 0:
    RAISE ERROR "Table '{table.name}' has invalid record_count: {table.record_count}. Must be > 0"

  IF NOT IS_INTEGER(table.record_count):
    RAISE ERROR "Table '{table.name}' has non-integer record_count: {table.record_count}"
```

**Examples**:

**Valid**:
- `1`
- `1000`
- `999999`

**Invalid**:
- `0`
- `-100`
- `1.5` (float)
- `"1000"` (string)

**Error messages**:
```
Table 'borrowers' has invalid record_count: 0. Must be > 0
Table 'loans' has non-integer record_count: 1.5
```

---

#### V-T003: Table Must Have At Least One Column

**Rule**: Each table must define at least one column in its `columns` array.

**Validation logic**:
```
FOR EACH table IN schema.tables:
  IF table.columns IS EMPTY:
    RAISE ERROR "Table '{table.name}' must define at least one column"
```

**Examples**:

**Valid**:
```json
{
  "name": "borrowers",
  "columns": [
    {"name": "id", "type": "int", "primary_key": true}
  ]
}
```

**Invalid**:
```json
{
  "name": "borrowers",
  "columns": []
}
```

**Error message**:
```
Table 'borrowers' must define at least one column
```

---

#### V-T004: Table Must Have Exactly One Primary Key

**Rule**: Each table must have exactly one column marked with `"primary_key": true`.

**Validation logic**:
```
FOR EACH table IN schema.tables:
  primary_keys = [col FOR col IN table.columns IF col.primary_key == true]

  IF LENGTH(primary_keys) == 0:
    RAISE ERROR "Table '{table.name}' has no primary key. Exactly one column must have primary_key: true"

  IF LENGTH(primary_keys) > 1:
    pk_names = [pk.name FOR pk IN primary_keys]
    RAISE ERROR "Table '{table.name}' has multiple primary keys: {pk_names}. Only one column can be primary key"
```

**Examples**:

**Valid**:
```json
{
  "name": "borrowers",
  "columns": [
    {"name": "id", "type": "int", "primary_key": true},
    {"name": "email", "type": "varchar(255)"}
  ]
}
```

**Invalid** (no primary key):
```json
{
  "name": "borrowers",
  "columns": [
    {"name": "email", "type": "varchar(255)"}
  ]
}
```

**Invalid** (multiple primary keys):
```json
{
  "name": "borrowers",
  "columns": [
    {"name": "id", "type": "int", "primary_key": true},
    {"name": "uuid", "type": "varchar(36)", "primary_key": true}
  ]
}
```

**Error messages**:
```
Table 'borrowers' has no primary key. Exactly one column must have primary_key: true
Table 'borrowers' has multiple primary keys: ['id', 'uuid']. Only one column can be primary key
```

**F008 Implementation Note**: Composite primary keys are not supported in MVP.

---

### Column-Level Validation

Column-level validation ensures that column definitions are correct, generators are valid, and parameters match requirements.

#### V-C001: Column Name Must Be Valid

**Rule**: Each column's `name` field must:
- Be non-empty
- Use lowercase with underscores for multi-word names
- Contain only lowercase letters, numbers, and underscores
- Be a valid SQL identifier
- Be unique within the table

**Validation logic**:
```
FOR EACH table IN schema.tables:
  seen_column_names = SET()

  FOR EACH column IN table.columns:
    IF column.name IS EMPTY:
      RAISE ERROR "Table '{table.name}': Column name cannot be empty"

    IF NOT MATCHES(column.name, "^[a-z][a-z0-9_]*$"):
      RAISE ERROR "Table '{table.name}': Column '{column.name}' uses invalid format. Use lowercase_with_underscores"

    IF column.name IN seen_column_names:
      RAISE ERROR "Table '{table.name}': Duplicate column name: {column.name}"

    seen_column_names.ADD(column.name)
```

**Examples**:

**Valid**:
- `"id"`
- `"first_name"`
- `"created_at"`
- `"user_id_2"`

**Invalid**:
- `""` (empty)
- `"firstName"` (camelCase)
- `"First-Name"` (hyphen, uppercase)
- `"1st_name"` (starts with number)

**Error messages**:
```
Table 'borrowers': Column name cannot be empty
Table 'borrowers': Column 'firstName' uses invalid format. Use lowercase_with_underscores
Table 'borrowers': Duplicate column name: email
```

---

#### V-C002: Column Data Type Must Be Valid

**Rule**: Each column's `type` field must be a supported SQL data type from the MySQL/PostgreSQL common subset.

**Supported types**:
- Integer: `int`, `bigint`, `smallint`, `tinyint`
- Decimal: `decimal(p,s)`, `float`, `double`
- String: `varchar(n)`, `text`, `char(n)`
- Date/Time: `date`, `datetime`, `timestamp`
- Boolean: `boolean`
- JSON: `json`, `jsonb` (PostgreSQL only, warn if used with MySQL)
- Enum: `enum('val1','val2',...)`

**Validation logic**:
```
valid_types_regex = "^(int|bigint|smallint|tinyint|float|double|boolean|text|json|jsonb|date|datetime|timestamp|decimal\\([0-9]+,[0-9]+\\)|varchar\\([0-9]+\\)|char\\([0-9]+\\)|enum\\(.+\\))$"

FOR EACH column IN table.columns:
  IF NOT MATCHES(column.type, valid_types_regex):
    RAISE ERROR "Table '{table.name}', Column '{column.name}': Invalid type '{column.type}'"

  # Warn about jsonb with MySQL
  IF column.type == "jsonb" AND "mysql" IN schema.database_type:
    WARN "Table '{table.name}', Column '{column.name}': jsonb is PostgreSQL-only. MySQL schemas should use json"
```

**Examples**:

**Valid**:
- `"int"`
- `"varchar(255)"`
- `"decimal(10,2)"`
- `"enum('active','paid','defaulted')"`
- `"timestamp"`

**Invalid**:
- `"string"` (use `varchar` or `text`)
- `"integer"` (use `int`)
- `"VARCHAR(255)"` (uppercase not allowed)
- `"decimal"` (missing precision/scale)
- `"varchar"` (missing length)

**Error messages**:
```
Table 'borrowers', Column 'email': Invalid type 'string'
Table 'borrowers', Column 'name': Invalid type 'VARCHAR(255)'
```

**F008 Implementation Note**: Use regex validation for complex types (decimal, varchar, enum).

---

#### V-C003: Generator Must Be Valid

**Rule**: If a column specifies a `generator`, it must be a valid built-in or custom generator name.

**Built-in generators** (MVP):
- Personal data: `first_name`, `last_name`, `full_name`, `email`, `phone`, `address`, `ssn`, `date_of_birth`
- Company data: `company_name`, `job_title`, `company_email`, `domain`
- Date/time: `timestamp_past`, `timestamp_future`, `date_between`
- Numeric: `int_range`, `float_range`, `decimal_range`
- Boolean: `weighted_boolean`
- Enum: `enum`
- UUID: `uuid`

**Validation logic**:
```
built_in_generators = ["first_name", "last_name", "full_name", "email", "phone", "address", "ssn", "date_of_birth", "company_name", "job_title", "company_email", "domain", "timestamp_past", "timestamp_future", "date_between", "int_range", "float_range", "decimal_range", "weighted_boolean", "enum", "uuid"]

FOR EACH column IN table.columns:
  IF "generator" IN column:
    IF column.generator NOT IN built_in_generators:
      RAISE ERROR "Table '{table.name}', Column '{column.name}': Unknown generator '{column.generator}'"
```

**Examples**:

**Valid**:
- `"email"`
- `"int_range"`
- `"date_between"`

**Invalid**:
- `"random_email"` (not a built-in generator)
- `"Email"` (incorrect case)
- `"emailGenerator"` (incorrect naming)

**Error message**:
```
Table 'borrowers', Column 'email': Unknown generator 'random_email'
```

**F008 Implementation Note**: Custom generators (defined in schema metadata) will be supported in post-MVP.

---

#### V-C004: Generator Parameters Must Match Requirements

**Rule**: If a column specifies `generator_params`, they must match the generator's requirements.

**Parameter requirements by generator**:

| Generator | Required Params | Optional Params |
|-----------|----------------|-----------------|
| `int_range` | `min`, `max` OR `distribution` | `distribution` |
| `float_range` | `min`, `max` OR `distribution` | `distribution` |
| `decimal_range` | `min`, `max` OR `distribution` | `distribution` |
| `date_between` | `start_date`, `end_date` | - |
| `timestamp_past` | `years_ago` | - |
| `timestamp_future` | `years_ahead` | - |
| `weighted_boolean` | `true_weight` | - |
| `enum` | `values` (array of {value, weight}) | - |

**Validation logic**:
```
FOR EACH column IN table.columns:
  IF "generator" IN column AND "generator_params" IN column:
    params = column.generator_params

    CASE column.generator:
      WHEN "int_range", "float_range", "decimal_range":
        IF "distribution" NOT IN params:
          IF "min" NOT IN params OR "max" NOT IN params:
            RAISE ERROR "Table '{table.name}', Column '{column.name}': {column.generator} requires 'min' and 'max' parameters OR 'distribution'"

      WHEN "date_between":
        IF "start_date" NOT IN params OR "end_date" NOT IN params:
          RAISE ERROR "Table '{table.name}', Column '{column.name}': date_between requires 'start_date' and 'end_date'"

      WHEN "timestamp_past":
        IF "years_ago" NOT IN params:
          RAISE ERROR "Table '{table.name}', Column '{column.name}': timestamp_past requires 'years_ago'"

      WHEN "enum":
        IF "values" NOT IN params OR NOT IS_ARRAY(params.values):
          RAISE ERROR "Table '{table.name}', Column '{column.name}': enum requires 'values' array"
```

**Examples**:

**Valid**:
```json
{
  "generator": "int_range",
  "generator_params": {
    "min": 300,
    "max": 850
  }
}
```

**Invalid** (missing params):
```json
{
  "generator": "int_range",
  "generator_params": {
    "max": 850
  }
}
```

**Error message**:
```
Table 'borrowers', Column 'credit_score': int_range requires 'min' and 'max' parameters OR 'distribution'
```

---

#### V-C005: Distribution Parameters Must Be Valid

**Rule**: If a column uses a `distribution`, the distribution type and parameters must be valid.

**Distribution types**:
- `uniform`: Requires `min`, `max`
- `normal`: Requires `mean`, `std_dev`, optional `min`, `max`
- `lognormal`: Requires `median`, `min`, `max`
- `weighted`: Requires `values` array with `{value, weight}` objects
- `ranges`: Requires `ranges` array with `{min, max, weight}` objects

**Validation logic**:
```
FOR EACH column IN table.columns:
  IF "distribution" IN column.generator_params:
    dist = column.generator_params.distribution

    IF dist.type NOT IN ["uniform", "normal", "lognormal", "weighted", "ranges"]:
      RAISE ERROR "Table '{table.name}', Column '{column.name}': Unknown distribution type '{dist.type}'"

    CASE dist.type:
      WHEN "uniform":
        IF "min" NOT IN dist.params OR "max" NOT IN dist.params:
          RAISE ERROR "Table '{table.name}', Column '{column.name}': uniform distribution requires 'min' and 'max'"

      WHEN "normal":
        IF "mean" NOT IN dist.params OR "std_dev" NOT IN dist.params:
          RAISE ERROR "Table '{table.name}', Column '{column.name}': normal distribution requires 'mean' and 'std_dev'"

      WHEN "lognormal":
        IF "median" NOT IN dist.params OR "min" NOT IN dist.params OR "max" NOT IN dist.params:
          RAISE ERROR "Table '{table.name}', Column '{column.name}': lognormal distribution requires 'median', 'min', and 'max'"

      WHEN "weighted":
        IF "values" NOT IN dist.params OR NOT IS_ARRAY(dist.params.values):
          RAISE ERROR "Table '{table.name}', Column '{column.name}': weighted distribution requires 'values' array"

        FOR EACH value_obj IN dist.params.values:
          IF "value" NOT IN value_obj OR "weight" NOT IN value_obj:
            RAISE ERROR "Table '{table.name}', Column '{column.name}': weighted values must have 'value' and 'weight'"

      WHEN "ranges":
        IF "ranges" NOT IN dist.params OR NOT IS_ARRAY(dist.params.ranges):
          RAISE ERROR "Table '{table.name}', Column '{column.name}': ranges distribution requires 'ranges' array"

        FOR EACH range_obj IN dist.params.ranges:
          IF "min" NOT IN range_obj OR "max" NOT IN range_obj OR "weight" NOT IN range_obj:
            RAISE ERROR "Table '{table.name}', Column '{column.name}': range objects must have 'min', 'max', and 'weight'"
```

**Examples**:

**Valid** (normal distribution):
```json
{
  "distribution": {
    "type": "normal",
    "params": {
      "mean": 680,
      "std_dev": 80,
      "min": 300,
      "max": 850
    }
  }
}
```

**Invalid** (missing params):
```json
{
  "distribution": {
    "type": "normal",
    "params": {
      "mean": 680
    }
  }
}
```

**Error message**:
```
Table 'borrowers', Column 'credit_score': normal distribution requires 'mean' and 'std_dev'
```

---

#### V-C006: Foreign Key Must Reference Existing Table and Column

**Rule**: If a column defines a `foreign_key`, the referenced table and column must exist in the schema.

**Validation logic**:
```
FOR EACH table IN schema.tables:
  FOR EACH column IN table.columns:
    IF "foreign_key" IN column:
      fk = column.foreign_key

      # Check referenced table exists
      ref_table = FIND_TABLE(schema.tables, fk.table)
      IF ref_table IS NULL:
        RAISE ERROR "Table '{table.name}', Column '{column.name}': Foreign key references non-existent table '{fk.table}'"

      # Check referenced column exists
      ref_column = FIND_COLUMN(ref_table.columns, fk.column)
      IF ref_column IS NULL:
        RAISE ERROR "Table '{table.name}', Column '{column.name}': Foreign key references non-existent column '{fk.table}.{fk.column}'"

      # Check referenced column is primary key or unique
      IF NOT (ref_column.primary_key == true OR ref_column.unique == true):
        RAISE ERROR "Table '{table.name}', Column '{column.name}': Foreign key must reference a primary key or unique column. '{fk.table}.{fk.column}' is neither"
```

**Examples**:

**Valid**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "loans",
      "columns": [
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {
            "table": "borrowers",
            "column": "id"
          }
        }
      ]
    }
  ]
}
```

**Invalid** (non-existent table):
```json
{
  "name": "borrower_id",
  "type": "int",
  "foreign_key": {
    "table": "users",
    "column": "id"
  }
}
```

**Invalid** (non-existent column):
```json
{
  "name": "borrower_id",
  "type": "int",
  "foreign_key": {
    "table": "borrowers",
    "column": "user_id"
  }
}
```

**Error messages**:
```
Table 'loans', Column 'borrower_id': Foreign key references non-existent table 'users'
Table 'loans', Column 'borrower_id': Foreign key references non-existent column 'borrowers.user_id'
Table 'loans', Column 'borrower_id': Foreign key must reference a primary key or unique column. 'borrowers.email' is neither
```

---

### Relationship-Level Validation

Relationship-level validation ensures that foreign key relationships are consistent and referential integrity actions are valid.

#### V-R001: Foreign Keys Must Reference Primary Keys or Unique Columns

**Rule**: Foreign key columns must reference columns that are either primary keys or have unique constraints.

**Validation logic**: (Covered in V-C006 above)

**Rationale**: Referencing non-unique columns would create ambiguous relationships (which parent record does the foreign key reference?).

---

#### V-R002: Foreign Key Data Types Must Match

**Rule**: Foreign key columns must have the same data type as the referenced column.

**Validation logic**:
```
FOR EACH table IN schema.tables:
  FOR EACH column IN table.columns:
    IF "foreign_key" IN column:
      fk = column.foreign_key
      ref_table = FIND_TABLE(schema.tables, fk.table)
      ref_column = FIND_COLUMN(ref_table.columns, fk.column)

      IF column.type != ref_column.type:
        RAISE ERROR "Table '{table.name}', Column '{column.name}': Foreign key type '{column.type}' does not match referenced column type '{ref_column.type}' in '{fk.table}.{fk.column}'"
```

**Examples**:

**Valid**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "loans",
      "columns": [
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {"table": "borrowers", "column": "id"}
        }
      ]
    }
  ]
}
```

**Invalid** (type mismatch):
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "loans",
      "columns": [
        {
          "name": "borrower_id",
          "type": "varchar(36)",
          "foreign_key": {"table": "borrowers", "column": "id"}
        }
      ]
    }
  ]
}
```

**Error message**:
```
Table 'loans', Column 'borrower_id': Foreign key type 'varchar(36)' does not match referenced column type 'int' in 'borrowers.id'
```

---

#### V-R003: Referential Integrity Actions Must Be Valid

**Rule**: Foreign key `on_delete` and `on_update` actions must be one of: `CASCADE`, `SET NULL`, `RESTRICT`.

**Validation logic**:
```
valid_actions = ["CASCADE", "SET NULL", "RESTRICT"]

FOR EACH column IN table.columns:
  IF "foreign_key" IN column:
    fk = column.foreign_key

    IF "on_delete" IN fk AND fk.on_delete NOT IN valid_actions:
      RAISE ERROR "Table '{table.name}', Column '{column.name}': Invalid on_delete action '{fk.on_delete}'. Valid: CASCADE, SET NULL, RESTRICT"

    IF "on_update" IN fk AND fk.on_update NOT IN valid_actions:
      RAISE ERROR "Table '{table.name}', Column '{column.name}': Invalid on_update action '{fk.on_update}'. Valid: CASCADE, SET NULL, RESTRICT"
```

**Examples**:

**Valid**:
- `"CASCADE"`
- `"SET NULL"`
- `"RESTRICT"`

**Invalid**:
- `"DELETE CASCADE"` (incorrect syntax)
- `"cascade"` (lowercase not allowed)
- `"NO ACTION"` (not supported in MVP)
- `"SET DEFAULT"` (not supported in MVP)

**Error message**:
```
Table 'loans', Column 'borrower_id': Invalid on_delete action 'NO ACTION'. Valid: CASCADE, SET NULL, RESTRICT
```

---

#### V-R004: SET NULL Requires Nullable Column

**Rule**: If a foreign key uses `on_delete: "SET NULL"` or `on_update: "SET NULL"`, the foreign key column must have `nullable: true`.

**Validation logic**:
```
FOR EACH column IN table.columns:
  IF "foreign_key" IN column:
    fk = column.foreign_key

    IF (fk.on_delete == "SET NULL" OR fk.on_update == "SET NULL") AND column.nullable != true:
      RAISE ERROR "Table '{table.name}', Column '{column.name}': Foreign key uses 'SET NULL' but column is not nullable. Set nullable: true"
```

**Examples**:

**Valid**:
```json
{
  "name": "user_id",
  "type": "int",
  "nullable": true,
  "foreign_key": {
    "table": "users",
    "column": "id",
    "on_delete": "SET NULL"
  }
}
```

**Invalid**:
```json
{
  "name": "user_id",
  "type": "int",
  "nullable": false,
  "foreign_key": {
    "table": "users",
    "column": "id",
    "on_delete": "SET NULL"
  }
}
```

**Error message**:
```
Table 'sessions', Column 'user_id': Foreign key uses 'SET NULL' but column is not nullable. Set nullable: true
```

---

### Generation Order Validation

Generation order validation ensures that tables can be populated in a valid sequence without violating foreign key constraints.

#### V-G001: Parent Tables Must Appear Before Child Tables

**Rule**: In `generation_order`, any table referenced by a foreign key (parent) must appear before the table containing the foreign key (child).

**Validation logic**:
```
# Build dependency map: child -> [parents]
dependencies = {}
FOR EACH table IN schema.tables:
  dependencies[table.name] = []
  FOR EACH column IN table.columns:
    IF "foreign_key" IN column:
      dependencies[table.name].APPEND(column.foreign_key.table)

# Validate generation order
table_positions = {name: index FOR index, name IN ENUMERATE(schema.generation_order)}

FOR child_table, parent_tables IN dependencies:
  FOR parent_table IN parent_tables:
    IF table_positions[child_table] < table_positions[parent_table]:
      RAISE ERROR "Invalid generation_order: '{child_table}' has foreign key to '{parent_table}', but '{parent_table}' appears later in generation_order (position {table_positions[parent_table]} vs {table_positions[child_table]})"
```

**Examples**:

**Valid**:
```json
{
  "generation_order": ["borrowers", "loans", "payments"]
}
```

**Invalid** (child before parent):
```json
{
  "generation_order": ["loans", "borrowers", "payments"]
}
```

**Error message**:
```
Invalid generation_order: 'loans' has foreign key to 'borrowers', but 'borrowers' appears later in generation_order (position 1 vs 0)
```

---

#### V-G002: No Circular Dependencies

**Rule**: The schema must not contain circular foreign key dependencies (e.g., A → B → A).

**Validation logic**:
```
# Build dependency graph
graph = {}
FOR EACH table IN schema.tables:
  graph[table.name] = []
  FOR EACH column IN table.columns:
    IF "foreign_key" IN column:
      graph[table.name].APPEND(column.foreign_key.table)

# Detect cycles using DFS
FOR EACH table IN schema.tables:
  visited = SET()
  stack = [table.name]

  WHILE stack IS NOT EMPTY:
    current = stack.POP()

    IF current IN visited:
      cycle_path = RECONSTRUCT_CYCLE(visited, current)
      RAISE ERROR "Circular dependency detected: {cycle_path}"

    visited.ADD(current)

    FOR EACH dependency IN graph[current]:
      IF dependency NOT IN visited:
        stack.PUSH(dependency)
```

**Examples**:

**Valid** (no cycles):
```
borrowers → (none)
loans → borrowers
payments → loans
```

**Invalid** (circular dependency):
```
users → addresses
addresses → users
```

**Error message**:
```
Circular dependency detected: users -> addresses -> users
```

**F008 Implementation Note**: Use topological sort or DFS cycle detection.

---

### Edge Cases and Common Errors

This section documents common schema errors and how to detect them.

#### Edge Case 1: Missing Foreign Key Target

**Scenario**: Foreign key references a table that doesn't exist.

**Detection**: V-C006

**Example**:
```json
{
  "tables": [
    {
      "name": "loans",
      "columns": [
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {"table": "users", "column": "id"}
        }
      ]
    }
  ]
}
```

**Error**:
```
Table 'loans', Column 'borrower_id': Foreign key references non-existent table 'users'
```

**Fix**: Add the `users` table or change the foreign key to reference an existing table.

---

#### Edge Case 2: Circular Dependency

**Scenario**: Two tables reference each other.

**Detection**: V-G002

**Example**:
```json
{
  "tables": [
    {
      "name": "users",
      "columns": [
        {
          "name": "primary_address_id",
          "type": "int",
          "foreign_key": {"table": "addresses", "column": "id"}
        }
      ]
    },
    {
      "name": "addresses",
      "columns": [
        {
          "name": "user_id",
          "type": "int",
          "foreign_key": {"table": "users", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["users", "addresses"]
}
```

**Error**:
```
Circular dependency detected: users -> addresses -> users
```

**Fix**: Remove one foreign key relationship or make it nullable and populate in a second pass (post-MVP feature).

---

#### Edge Case 3: Invalid Generator Parameters

**Scenario**: Generator parameters don't match requirements.

**Detection**: V-C004

**Example**:
```json
{
  "name": "credit_score",
  "type": "int",
  "generator": "int_range",
  "generator_params": {
    "max": 850
  }
}
```

**Error**:
```
Table 'borrowers', Column 'credit_score': int_range requires 'min' and 'max' parameters OR 'distribution'
```

**Fix**: Add `"min": 300` to generator_params.

---

#### Edge Case 4: Type Mismatch in Foreign Key

**Scenario**: Foreign key column type doesn't match referenced column type.

**Detection**: V-R002

**Example**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "bigint", "primary_key": true}
      ]
    },
    {
      "name": "loans",
      "columns": [
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {"table": "borrowers", "column": "id"}
        }
      ]
    }
  ]
}
```

**Error**:
```
Table 'loans', Column 'borrower_id': Foreign key type 'int' does not match referenced column type 'bigint' in 'borrowers.id'
```

**Fix**: Change `borrower_id` type to `bigint`.

---

#### Edge Case 5: SET NULL Without Nullable

**Scenario**: Foreign key uses `SET NULL` but column is not nullable.

**Detection**: V-R004

**Example**:
```json
{
  "name": "user_id",
  "type": "int",
  "nullable": false,
  "foreign_key": {
    "table": "users",
    "column": "id",
    "on_delete": "SET NULL"
  }
}
```

**Error**:
```
Table 'sessions', Column 'user_id': Foreign key uses 'SET NULL' but column is not nullable. Set nullable: true
```

**Fix**: Change to `"nullable": true`.

---

#### Edge Case 6: Multiple Primary Keys

**Scenario**: Table has more than one column marked as primary key.

**Detection**: V-T004

**Example**:
```json
{
  "name": "borrowers",
  "columns": [
    {"name": "id", "type": "int", "primary_key": true},
    {"name": "uuid", "type": "varchar(36)", "primary_key": true}
  ]
}
```

**Error**:
```
Table 'borrowers' has multiple primary keys: ['id', 'uuid']. Only one column can be primary key
```

**Fix**: Remove `"primary_key": true` from one column.

---

#### Edge Case 7: Empty Generation Order

**Scenario**: `generation_order` is defined but empty.

**Detection**: V-S007

**Example**:
```json
{
  "tables": [
    {"name": "borrowers", ...}
  ],
  "generation_order": []
}
```

**Error**:
```
Tables missing from generation_order: ['borrowers']
```

**Fix**: Add all tables to `generation_order`.

---

### Error Message Guidance for F008 Implementers

This section provides standards for error messages to ensure they are clear, actionable, and helpful.

#### Error Message Structure

**Format**:
```
[SEVERITY] {Context}: {Problem}. {Suggestion}
```

**Components**:
- **SEVERITY**: `ERROR` (blocks schema usage) or `WARNING` (non-blocking)
- **Context**: Where the error occurred (table, column, field)
- **Problem**: What is wrong
- **Suggestion**: How to fix (when possible)

**Examples**:

**Good**:
```
ERROR Table 'loans', Column 'borrower_id': Foreign key references non-existent table 'users'. Did you mean 'borrowers'?
```

**Bad**:
```
Invalid foreign key
```

---

#### Error Message Principles

**1. Be Specific**:
- Always include context (table name, column name, field name)
- Quote values to distinguish them from prose
- Use exact error locations (line numbers if available)

**Good**:
```
Table 'borrowers', Column 'credit_score': int_range requires 'min' and 'max' parameters OR 'distribution'
```

**Bad**:
```
Missing parameters
```

---

**2. Be Actionable**:
- Explain what needs to be fixed
- Suggest corrections when possible
- Link to documentation for complex rules

**Good**:
```
Schema name 'Fintech-Loans' must use lowercase-kebab-case format (e.g., 'fintech-loans')
```

**Bad**:
```
Invalid schema name
```

---

**3. Be Consistent**:
- Use consistent terminology (e.g., always "foreign key" not "FK")
- Use consistent formatting (e.g., always quote table/column names)
- Group related errors (e.g., all missing fields together)

---

**4. Provide Examples**:
- Show correct format when rejecting incorrect format
- Reference existing schema examples
- Use realistic examples (not "foo", "bar")

**Good**:
```
Schema version must follow semantic versioning (e.g., '1.0.0')
```

**Bad**:
```
Invalid version format
```

---

**5. Prioritize Errors**:
- Show structural errors first (missing required fields)
- Show semantic errors second (invalid values)
- Show cross-validation errors last (foreign key references)

---

**6. Suggest Corrections**:
- Use "Did you mean?" for likely typos
- Suggest valid values when rejecting invalid values
- Reference similar valid schemas

**Good**:
```
Unknown table in generation_order: 'borrower'. Did you mean 'borrowers'?
```

**Bad**:
```
Unknown table
```

---

#### Error Message Examples by Category

**Structural Errors**:
```
Missing required field: author
Invalid JSON syntax at line 42: unexpected token '}'
Field 'database_type' must be an array, got string
```

**Schema-Level Errors**:
```
Schema name 'Fintech_Loans' must use lowercase-kebab-case format (e.g., 'fintech-loans')
Schema version '1.0' must follow semantic versioning (e.g., '1.0.0')
database_type must contain at least one database type. Valid: mysql, postgres
```

**Table-Level Errors**:
```
Table 'borrowers' has no primary key. Exactly one column must have primary_key: true
Table 'loans' has invalid record_count: 0. Must be > 0
Duplicate table name: borrowers
```

**Column-Level Errors**:
```
Table 'borrowers', Column 'email': Invalid type 'string'. Use 'varchar(n)' or 'text'
Table 'borrowers', Column 'credit_score': int_range requires 'min' and 'max' parameters OR 'distribution'
Table 'borrowers': Duplicate column name: email
```

**Relationship Errors**:
```
Table 'loans', Column 'borrower_id': Foreign key references non-existent table 'users'. Did you mean 'borrowers'?
Table 'loans', Column 'borrower_id': Foreign key type 'int' does not match referenced column type 'bigint' in 'borrowers.id'
Table 'sessions', Column 'user_id': Foreign key uses 'SET NULL' but column is not nullable. Set nullable: true
```

**Generation Order Errors**:
```
Tables missing from generation_order: ['loans', 'payments']
Invalid generation_order: 'loans' has foreign key to 'borrowers', but 'borrowers' appears later in generation_order
Circular dependency detected: users -> addresses -> users
```

---

#### Severity Levels

**ERROR** (blocks schema usage):
- Missing required fields
- Invalid syntax
- Structural violations (no primary key, circular dependencies)
- Type mismatches
- Invalid references (non-existent tables/columns)

**WARNING** (non-blocking, but should be reviewed):
- Using `jsonb` with MySQL (database-specific type)
- Very large record counts (may exceed Tier 1 limits)
- Complex distributions (may impact generation time)
- Missing descriptions (optional but recommended)

---

#### Multi-Error Reporting

When multiple errors exist, report them in priority order:

1. **Structural errors** (JSON syntax, missing required fields)
2. **Schema-level errors** (invalid name, version, database_type)
3. **Table-level errors** (invalid tables, missing primary keys)
4. **Column-level errors** (invalid columns, types, generators)
5. **Relationship errors** (invalid foreign keys)
6. **Generation order errors** (invalid ordering, circular dependencies)

**Example multi-error output**:
```
Schema validation failed with 3 errors:

ERROR: Missing required field: author
ERROR: Table 'borrowers' has no primary key. Exactly one column must have primary_key: true
ERROR: Invalid generation_order: 'loans' has foreign key to 'borrowers', but 'borrowers' appears later in generation_order
```

---

### Validation Summary

**Schema-Level Validation** (7 rules):
- V-S001: Required fields present
- V-S002: Valid schema name
- V-S003: Valid semantic version
- V-S004: Valid format version
- V-S005: Valid database types
- V-S006: At least one table
- V-S007: Generation order includes all tables

**Table-Level Validation** (4 rules):
- V-T001: Valid table names
- V-T002: Positive record counts
- V-T003: At least one column
- V-T004: Exactly one primary key

**Column-Level Validation** (6 rules):
- V-C001: Valid column names
- V-C002: Valid data types
- V-C003: Valid generators
- V-C004: Valid generator parameters
- V-C005: Valid distribution parameters
- V-C006: Foreign keys reference existing tables/columns

**Relationship-Level Validation** (4 rules):
- V-R001: Foreign keys reference primary keys or unique columns
- V-R002: Foreign key type matches referenced type
- V-R003: Valid referential integrity actions
- V-R004: SET NULL requires nullable column

**Generation Order Validation** (2 rules):
- V-G001: Parent tables before child tables
- V-G002: No circular dependencies

**Total**: 23 validation rules covering all aspects of schema correctness.

---

## Versioning Strategy

SourceBox schemas use a dual versioning system to separately track format evolution and content evolution. Understanding the distinction between these two version types is critical for both schema authors and schema consumers.

### Why Two Version Fields?

Schema files contain two independent version numbers that serve different purposes:

1. **schema_version**: The format specification version (controlled by SourceBox maintainers)
2. **version**: The schema content version (controlled by schema authors)

This separation exists because **format changes** and **content changes** happen at different rates and are controlled by different parties:

- **Format changes** are rare, breaking, and require parser updates (e.g., adding new required top-level fields, changing validation rules)
- **Content changes** are frequent, incremental, and require only schema file updates (e.g., adding tables, tweaking distributions)

Without this separation, every schema author would need to coordinate version bumps with SourceBox maintainers, creating unnecessary coupling and slowing down schema development.

### schema_version (Format Version)

The `schema_version` field tracks the version of this format specification itself. It tells the parser which validation rules to apply and which features are available.

#### Purpose

- **Parser compatibility**: The parser checks this field to determine how to validate and process the schema
- **Format evolution**: Enables backward-compatible format changes over time
- **Breaking change detection**: Prevents old parsers from incorrectly processing new schema formats

#### Format

- **Type**: String (semantic versioning: `major.minor`)
- **Current value**: `"1.0"` (MVP format specification)
- **Control**: SourceBox maintainers (not schema authors)

#### When It Changes

`schema_version` is incremented only when the schema JSON format itself changes:

**Major version bump** (1.0 → 2.0):
- New required top-level fields added (e.g., `license`, `dependencies`)
- Existing field semantics changed (e.g., `generation_order` becomes required)
- Validation rules significantly changed (e.g., new constraints on generator parameters)
- Parser behavior fundamentally altered

**Minor version bump** (1.0 → 1.1):
- New optional top-level fields added (e.g., `custom_generators`, `macros`)
- New generator types introduced (e.g., `ai_generator` for LLM-based generation)
- Backward-compatible enhancements (e.g., additional distribution types)

#### Examples

**Current MVP schema**:
```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "version": "1.0.0"
}
```

**Hypothetical future format** (if SourceBox adds required license field):
```json
{
  "schema_version": "2.0",
  "name": "fintech-loans",
  "version": "1.0.0",
  "license": "CC0-1.0"
}
```

**Hypothetical future format** (if SourceBox adds optional macros):
```json
{
  "schema_version": "1.1",
  "name": "fintech-loans",
  "version": "1.0.0",
  "macros": {
    "common_timestamps": ["created_at", "updated_at"]
  }
}
```

#### Schema Author Guidance

**Do NOT change `schema_version` when**:
- Adding/removing tables from your schema
- Changing generators or distributions
- Updating record counts or relationships
- Fixing bugs in data generation parameters
- Improving documentation

**Schema authors should always use the current format version** (`"1.0"` for MVP). SourceBox maintainers will announce format version changes through release notes and migration guides.

### version (Content Version)

The `version` field tracks the evolution of the schema's content—the tables, columns, generators, and relationships defined in the schema file. This is the version that schema authors control and increment as they improve their schemas.

#### Purpose

- **Change tracking**: Enable users to understand what changed between schema versions
- **Compatibility signaling**: Major version bumps signal breaking changes that may affect downstream users
- **Rollback support**: Allow users to reference specific schema versions (e.g., `fintech-loans@1.2.0`)
- **Collaboration**: Enable multiple contributors to coordinate schema updates

#### Format

- **Type**: String (semantic versioning: `major.minor.patch`)
- **Current value**: `"1.0.0"` for new schemas
- **Control**: Schema authors (you)

#### Semantic Versioning Rules

Schema content versions follow standard semantic versioning (semver) conventions:

**MAJOR version** (1.0.0 → 2.0.0) - Breaking changes:
- Table removed
- Column removed
- Column data type changed (e.g., `int` → `varchar`)
- Column renamed (breaks references)
- Foreign key relationships removed
- Record counts significantly changed (e.g., 1000 → 10 breaks downstream tests)
- Generation order changed in a way that breaks relationships

**MINOR version** (1.0.0 → 1.1.0) - Backward-compatible additions:
- New table added
- New column added to existing table
- New foreign key relationship added
- Generator changed (but generates compatible data)
- Distribution tweaked (but still produces similar data patterns)
- Record counts increased moderately
- Documentation improved

**PATCH version** (1.0.0 → 1.0.1) - Bug fixes and minor tweaks:
- Distribution parameters fixed (e.g., credit score mean was 500, should be 680)
- Generator parameters corrected (e.g., date range was wrong)
- Typos in descriptions fixed
- Documentation clarified
- Metadata updated (tags, industry classification)
- Cosmetic changes (formatting, comments)

#### Examples

**Initial schema release**:
```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "version": "1.0.0",
  "tables": [
    {
      "name": "borrowers",
      "record_count": 1000,
      "columns": [...]
    }
  ]
}
```

**Minor version bump** (added new `payments` table):
```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "version": "1.1.0",
  "tables": [
    {
      "name": "borrowers",
      "record_count": 1000,
      "columns": [...]
    },
    {
      "name": "payments",
      "record_count": 3000,
      "columns": [...]
    }
  ]
}
```

**Patch version bump** (fixed credit score distribution):
```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "version": "1.1.1",
  "tables": [
    {
      "name": "borrowers",
      "record_count": 1000,
      "columns": [
        {
          "name": "credit_score",
          "type": "int",
          "generator": "int_range",
          "distribution": {
            "type": "normal",
            "params": {
              "mean": 680,  // was 500 in 1.1.0
              "std_dev": 80,
              "min": 300,
              "max": 850
            }
          }
        }
      ]
    }
  ]
}
```

**Major version bump** (removed `borrowers` table, replaced with `customers`):
```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "version": "2.0.0",
  "tables": [
    {
      "name": "customers",  // breaking change: renamed table
      "record_count": 1000,
      "columns": [...]
    }
  ]
}
```

#### When to Increment Which Component

Use this decision tree when updating a schema:

```
Did you remove or rename tables/columns?
└─ YES → MAJOR (2.0.0)
└─ NO → Continue...

Did you change column data types or remove relationships?
└─ YES → MAJOR (2.0.0)
└─ NO → Continue...

Did you add new tables or columns?
└─ YES → MINOR (1.1.0)
└─ NO → Continue...

Did you change generators or distributions (non-breaking)?
└─ YES → MINOR (1.1.0)
└─ NO → Continue...

Did you fix bugs, improve docs, or tweak metadata?
└─ YES → PATCH (1.0.1)
└─ NO → No version change needed
```

### Backward Compatibility Considerations

#### For Schema Authors

When incrementing schema versions, consider the impact on downstream users:

**Breaking changes should be rare**:
- Removing tables/columns forces users to update their code
- Renaming breaks existing references and scripts
- Type changes may break data import pipelines

**Prefer additions over removals**:
- Add new columns instead of changing existing ones
- Deprecate old tables/columns in MINOR versions, remove in MAJOR versions
- Provide migration guides for major version bumps

**Document changes clearly**:
- Maintain a CHANGELOG.md alongside your schema
- Explain why breaking changes were necessary
- Provide migration scripts or examples

**Test before releasing**:
- Generate data with both old and new schema versions
- Verify foreign key relationships still work
- Check that distributions produce expected data patterns

#### For Schema Consumers

When using schemas, be mindful of version dependencies:

**Pin to specific versions**:
- Use exact versions in production (e.g., `fintech-loans@1.2.0`)
- Test against specific versions to ensure reproducibility
- Update dependencies deliberately, not automatically

**Monitor for updates**:
- Watch for MAJOR version bumps (breaking changes)
- Review MINOR versions for new features
- Apply PATCH versions for bug fixes

**Handle version conflicts**:
- If a MAJOR version breaks your workflow, stay on the old version until ready to migrate
- Report issues to schema authors if changes are unexpectedly breaking

### Version Tracking Best Practices

**For new schemas**:
1. Start with `schema_version: "1.0"` (current format)
2. Start with `version: "1.0.0"` (initial content)
3. Increment `version` as you make changes (never change `schema_version`)

**For schema updates**:
1. Determine the scope of your changes (breaking, addition, fix)
2. Increment the appropriate version component (major, minor, patch)
3. Update documentation and CHANGELOG
4. Tag the schema file in version control (if using Git)

**For schema repositories**:
1. Use Git tags to track schema versions (e.g., `git tag v1.2.0`)
2. Maintain a CHANGELOG.md documenting all changes
3. Provide migration guides for major version bumps
4. Consider publishing schemas to a schema registry

### Version Summary Table

| Field | Type | Controlled By | Purpose | Example |
|-------|------|---------------|---------|---------|
| `schema_version` | string (major.minor) | SourceBox maintainers | Format specification version | `"1.0"` |
| `version` | string (major.minor.patch) | Schema authors | Schema content version | `"1.2.0"` |

### Quick Reference

**schema_version changes when**:
- SourceBox adds required fields to the format
- Parser validation rules change significantly
- New top-level schema features are added
- **You (schema author) never change this**

**version changes when**:
- You add/remove tables or columns (MAJOR if removal, MINOR if addition)
- You change data types or relationships (MAJOR)
- You improve generators or distributions (MINOR)
- You fix bugs or update docs (PATCH)
- **You (schema author) always change this**

---

## Generation Order

The `generation_order` field is a critical component of schema definitions that ensures data is generated in the correct sequence to satisfy foreign key constraints. When tables have relationships (parent-child hierarchies), child tables cannot be populated until their parent tables exist—otherwise, foreign key values would reference non-existent records.

### Purpose

The primary purpose of `generation_order` is to **ensure referential integrity during data generation**. By explicitly specifying the order in which tables should be populated, schema authors guarantee that:

1. **Parent tables are created before child tables**: A child table referencing a parent via foreign key must wait until the parent's records exist
2. **Foreign key values are valid**: When generating child records, the generator can select from the pool of existing parent record IDs
3. **No circular dependencies**: The schema parser can detect impossible dependency cycles (e.g., Table A depends on Table B, which depends on Table A)
4. **Deterministic generation**: The same schema always generates data in the same order, ensuring reproducibility

Without `generation_order`, the parser would need to infer the correct order by analyzing foreign key relationships—a complex problem prone to errors, especially with multi-level hierarchies and independent tables.

### Format

`generation_order` is an array of strings, where each string is the name of a table defined in the schema.

```json
{
  "generation_order": ["borrowers", "loans", "payments"]
}
```

**Requirements**:
- **Array of strings**: Each element is a table name (must match a table's `name` field exactly)
- **All tables included**: Every table in the `tables` array must appear exactly once in `generation_order`
- **No duplicates**: Each table name can only appear once
- **Parent tables first**: Tables referenced by foreign keys must appear before tables that reference them
- **Case-sensitive**: Table names must match exactly (case-sensitive)

### How the Parser Uses generation_order

When the SourceBox generator processes a schema, it follows this sequence:

1. **Validate generation_order**: Check that all tables are present, no duplicates exist, and parent-child ordering is correct
2. **Populate tables sequentially**: Iterate through `generation_order` and generate records for each table in order
3. **Track generated IDs**: As each table is populated, store the range of generated primary key IDs (e.g., borrowers have IDs 1-1000)
4. **Generate foreign key values**: When generating child records, randomly select from the pool of existing parent IDs
5. **Enforce referential integrity**: Because parent records exist before children are created, all foreign keys reference valid records

**Example**: For `generation_order: ["borrowers", "loans", "payments"]`:

```
Step 1: Generate 1,000 borrowers (IDs: 1-1000)
Step 2: Generate 2,500 loans (each loan.borrower_id randomly selected from 1-1000)
Step 3: Generate 7,500 payments (each payment.loan_id randomly selected from 1-2500)
```

This ensures every loan has a valid borrower, and every payment has a valid loan.

### Validation Rules

The schema parser enforces the following rules on `generation_order`:

#### Rule 1: All Tables Must Appear

Every table defined in the `tables` array must appear exactly once in `generation_order`.

**Valid**:
```json
{
  "tables": [
    {"name": "borrowers", ...},
    {"name": "loans", ...}
  ],
  "generation_order": ["borrowers", "loans"]
}
```

**Invalid** (missing `loans`):
```json
{
  "tables": [
    {"name": "borrowers", ...},
    {"name": "loans", ...}
  ],
  "generation_order": ["borrowers"]
}
```

**Parser error**: `"Table 'loans' is defined but not included in generation_order"`

#### Rule 2: No Duplicates

Each table name can only appear once in `generation_order`.

**Valid**:
```json
{
  "generation_order": ["borrowers", "loans", "payments"]
}
```

**Invalid**:
```json
{
  "generation_order": ["borrowers", "loans", "borrowers"]
}
```

**Parser error**: `"Table 'borrowers' appears multiple times in generation_order"`

#### Rule 3: Parent Tables Before Child Tables

Any table referenced by a foreign key (parent) must appear before the table containing the foreign key (child).

**Valid**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [{"name": "id", "type": "int", "primary_key": true}]
    },
    {
      "name": "loans",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {"table": "borrowers", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["borrowers", "loans"]
}
```

**Invalid** (child before parent):
```json
{
  "generation_order": ["loans", "borrowers"]
}
```

**Parser error**: `"Table 'loans' has foreign key to 'borrowers', but 'borrowers' appears later in generation_order"`

#### Rule 4: No Circular Dependencies

Circular foreign key relationships are not allowed (e.g., Table A references Table B, Table B references Table A).

**Invalid schema**:
```json
{
  "tables": [
    {
      "name": "users",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "primary_address_id",
          "type": "int",
          "foreign_key": {"table": "addresses", "column": "id"}
        }
      ]
    },
    {
      "name": "addresses",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "user_id",
          "type": "int",
          "foreign_key": {"table": "users", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["users", "addresses"]
}
```

**Parser error**: `"Circular dependency detected: users -> addresses -> users"`

**Solution**: Remove one foreign key or make it nullable and populate it in a second pass (advanced pattern, not MVP).

#### Rule 5: Exact Name Matching

Table names in `generation_order` must exactly match the `name` field in table definitions (case-sensitive).

**Valid**:
```json
{
  "tables": [
    {"name": "borrowers", ...}
  ],
  "generation_order": ["borrowers"]
}
```

**Invalid** (case mismatch):
```json
{
  "tables": [
    {"name": "borrowers", ...}
  ],
  "generation_order": ["Borrowers"]
}
```

**Parser error**: `"Table 'Borrowers' in generation_order does not match any defined table (did you mean 'borrowers'?)"`

### Relationship Patterns and Ordering

#### One-to-Many Relationships (Single Level)

**Pattern**: Parent table has many children (e.g., one borrower has many loans).

**Ordering rule**: Parent first, then child.

**Example**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "loans",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {"table": "borrowers", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["borrowers", "loans"]
}
```

**Relationship diagram**:
```
borrowers (1) ----< loans (many)
```

**Generation flow**:
1. Create 1,000 borrowers (IDs: 1-1000)
2. Create 2,500 loans (each references a borrower_id from 1-1000)

#### Multi-Level Hierarchies (Grandparent → Parent → Child)

**Pattern**: Three or more levels of relationships (e.g., borrowers → loans → payments).

**Ordering rule**: Grandparent first, parent second, child last.

**Example**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "loans",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {"table": "borrowers", "column": "id"}
        }
      ]
    },
    {
      "name": "payments",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "loan_id",
          "type": "int",
          "foreign_key": {"table": "loans", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["borrowers", "loans", "payments"]
}
```

**Relationship diagram**:
```
borrowers (1) ----< loans (many) ----< payments (many)
```

**Generation flow**:
1. Create 1,000 borrowers (IDs: 1-1000)
2. Create 2,500 loans (each references a borrower_id from 1-1000)
3. Create 7,500 payments (each references a loan_id from 1-2500)

#### Independent Tables (No Relationships)

**Pattern**: Tables with no foreign keys or relationships to other tables.

**Ordering rule**: Any order is valid (independent tables can be generated in parallel, though sequential generation is simpler for MVP).

**Example**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "products",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    }
  ],
  "generation_order": ["borrowers", "products"]
}
```

**Alternative valid ordering**:
```json
{
  "generation_order": ["products", "borrowers"]
}
```

Both are correct because `borrowers` and `products` have no relationship to each other.

#### Multiple Parents (Child References Multiple Parents)

**Pattern**: A child table has foreign keys to multiple parent tables (e.g., loan_applications references both borrowers and loan_officers).

**Ordering rule**: Both parents before the child (parent order relative to each other doesn't matter if they're independent).

**Example**:
```json
{
  "tables": [
    {
      "name": "borrowers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "loan_officers",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "loan_applications",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {"table": "borrowers", "column": "id"}
        },
        {
          "name": "loan_officer_id",
          "type": "int",
          "foreign_key": {"table": "loan_officers", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["borrowers", "loan_officers", "loan_applications"]
}
```

**Alternative valid ordering** (parents swapped):
```json
{
  "generation_order": ["loan_officers", "borrowers", "loan_applications"]
}
```

Both are correct because `borrowers` and `loan_officers` are independent—only `loan_applications` depends on them.

**Relationship diagram**:
```
borrowers (1) ----< loan_applications (many) >---- loan_officers (1)
```

#### Many-to-Many Relationships (Junction Tables)

**Pattern**: Two tables related through a junction/join table (e.g., students and courses related through enrollments).

**Ordering rule**: Both entity tables first, junction table last.

**Example**:
```json
{
  "tables": [
    {
      "name": "students",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "courses",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "enrollments",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "student_id",
          "type": "int",
          "foreign_key": {"table": "students", "column": "id"}
        },
        {
          "name": "course_id",
          "type": "int",
          "foreign_key": {"table": "courses", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["students", "courses", "enrollments"]
}
```

**Alternative valid ordering** (entity tables swapped):
```json
{
  "generation_order": ["courses", "students", "enrollments"]
}
```

**Relationship diagram**:
```
students (many) ----< enrollments (junction) >---- courses (many)
```

### Common Mistakes and How to Fix Them

#### Mistake 1: Child Before Parent

**Problem**: Child table appears before the parent it references.

**Invalid**:
```json
{
  "generation_order": ["loans", "borrowers"]
}
```

**Error**: `"Table 'loans' has foreign key to 'borrowers', but 'borrowers' appears later in generation_order"`

**Fix**: Move parent table before child.

**Correct**:
```json
{
  "generation_order": ["borrowers", "loans"]
}
```

#### Mistake 2: Missing Table

**Problem**: A table is defined but not included in `generation_order`.

**Invalid**:
```json
{
  "tables": [
    {"name": "borrowers", ...},
    {"name": "loans", ...},
    {"name": "payments", ...}
  ],
  "generation_order": ["borrowers", "loans"]
}
```

**Error**: `"Table 'payments' is defined but not included in generation_order"`

**Fix**: Add all tables to `generation_order`.

**Correct**:
```json
{
  "generation_order": ["borrowers", "loans", "payments"]
}
```

#### Mistake 3: Typo in Table Name

**Problem**: Table name in `generation_order` doesn't match defined table (typo, case mismatch).

**Invalid**:
```json
{
  "tables": [
    {"name": "borrowers", ...}
  ],
  "generation_order": ["borrower"]
}
```

**Error**: `"Table 'borrower' in generation_order does not match any defined table (did you mean 'borrowers'?)"`

**Fix**: Correct the table name to match exactly.

**Correct**:
```json
{
  "generation_order": ["borrowers"]
}
```

#### Mistake 4: Circular Dependency

**Problem**: Two tables reference each other (impossible to determine ordering).

**Invalid**:
```json
{
  "tables": [
    {
      "name": "users",
      "columns": [
        {
          "name": "primary_address_id",
          "type": "int",
          "foreign_key": {"table": "addresses", "column": "id"}
        }
      ]
    },
    {
      "name": "addresses",
      "columns": [
        {
          "name": "user_id",
          "type": "int",
          "foreign_key": {"table": "users", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["users", "addresses"]
}
```

**Error**: `"Circular dependency detected: users -> addresses -> users"`

**Fix**: Remove one foreign key or restructure the schema (e.g., make one relationship optional).

**Correct** (remove circular dependency):
```json
{
  "tables": [
    {
      "name": "users",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true}
      ]
    },
    {
      "name": "addresses",
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "user_id",
          "type": "int",
          "foreign_key": {"table": "users", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": ["users", "addresses"]
}
```

### Determining Correct Ordering for Complex Schemas

For schemas with many tables and complex relationships, follow this process to determine the correct `generation_order`:

#### Step 1: List All Tables

Write down all table names from the `tables` array.

**Example**:
```
- borrowers
- loans
- payments
- loan_officers
- loan_statuses
```

#### Step 2: Identify All Foreign Keys

For each table, note which other tables it references via foreign keys.

**Example**:
```
- borrowers: (no foreign keys)
- loans: → borrowers, → loan_officers, → loan_statuses
- payments: → loans
- loan_officers: (no foreign keys)
- loan_statuses: (no foreign keys)
```

#### Step 3: Draw a Dependency Graph

Visualize the relationships. An arrow from A → B means "A depends on B" (A has a foreign key to B).

**Example**:
```
borrowers <---- loans <---- payments
                ↑   ↑
loan_officers --+   |
                    |
loan_statuses ------+
```

#### Step 4: Perform Topological Sort

Arrange tables so that every parent appears before its children.

**Process**:
1. Find tables with no dependencies (no outgoing arrows): `borrowers`, `loan_officers`, `loan_statuses`
2. Place them first (any order among themselves)
3. Find tables whose dependencies are now satisfied: `loans` (all its parents are placed)
4. Place `loans` next
5. Find tables whose dependencies are now satisfied: `payments`
6. Place `payments` last

**Result**:
```json
{
  "generation_order": [
    "borrowers",
    "loan_officers",
    "loan_statuses",
    "loans",
    "payments"
  ]
}
```

**Alternative valid orderings** (independent tables can be reordered):
```json
{
  "generation_order": [
    "loan_statuses",
    "borrowers",
    "loan_officers",
    "loans",
    "payments"
  ]
}
```

Both are correct because `borrowers`, `loan_officers`, and `loan_statuses` are independent.

#### Step 5: Validate with Parser

Once you've determined the order, validate it by running the schema through the SourceBox parser. The parser will catch any mistakes (missing tables, incorrect ordering, circular dependencies).

### Complete Example

Here is a complete schema with a realistic `generation_order`:

```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "description": "Realistic fintech loan data with multi-level relationships",
  "author": "SourceBox Contributors",
  "version": "1.0.0",
  "database_type": ["mysql", "postgres"],
  "tables": [
    {
      "name": "borrowers",
      "record_count": 1000,
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {"name": "email", "type": "varchar(255)", "generator": "email"}
      ]
    },
    {
      "name": "loan_officers",
      "record_count": 50,
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {"name": "name", "type": "varchar(200)", "generator": "full_name"}
      ]
    },
    {
      "name": "loan_statuses",
      "record_count": 4,
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {"name": "status_name", "type": "varchar(50)"}
      ]
    },
    {
      "name": "loans",
      "record_count": 2500,
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "borrower_id",
          "type": "int",
          "foreign_key": {"table": "borrowers", "column": "id"}
        },
        {
          "name": "loan_officer_id",
          "type": "int",
          "foreign_key": {"table": "loan_officers", "column": "id"}
        },
        {
          "name": "status_id",
          "type": "int",
          "foreign_key": {"table": "loan_statuses", "column": "id"}
        }
      ]
    },
    {
      "name": "payments",
      "record_count": 7500,
      "columns": [
        {"name": "id", "type": "int", "primary_key": true},
        {
          "name": "loan_id",
          "type": "int",
          "foreign_key": {"table": "loans", "column": "id"}
        }
      ]
    }
  ],
  "generation_order": [
    "borrowers",
    "loan_officers",
    "loan_statuses",
    "loans",
    "payments"
  ]
}
```

**Explanation**:
1. **borrowers**, **loan_officers**, **loan_statuses**: Independent tables (no foreign keys), generated first in any order
2. **loans**: Depends on all three independent tables, generated after they exist
3. **payments**: Depends on loans, generated last

**Generation flow**:
```
1. Generate 1,000 borrowers (IDs: 1-1000)
2. Generate 50 loan_officers (IDs: 1-50)
3. Generate 4 loan_statuses (IDs: 1-4)
4. Generate 2,500 loans (each references valid borrower_id, loan_officer_id, status_id)
5. Generate 7,500 payments (each references valid loan_id)
```

### Summary

**Key points**:
- `generation_order` ensures parent tables are populated before child tables
- It is an array of table names, specifying the order for sequential table population
- All tables must appear exactly once
- Parent tables must appear before child tables
- The parser validates ordering and detects circular dependencies
- Independent tables can be ordered arbitrarily
- Multi-level hierarchies require grandparent → parent → child ordering
- Use topological sorting for complex schemas with many relationships

**Best practices**:
- Draw dependency diagrams for complex schemas
- Validate with the parser after defining `generation_order`
- Document complex ordering logic in schema descriptions
- Group independent tables together for clarity (e.g., all lookup tables first)
- Use descriptive comments if the ordering is non-obvious

**What to avoid**:
- Listing child tables before parent tables
- Omitting tables from `generation_order`
- Creating circular dependencies (A → B → A)
- Typos or case mismatches in table names
- Guessing the order without analyzing foreign keys

---

_Subsequent sections will be added in tasks T010-T012._
