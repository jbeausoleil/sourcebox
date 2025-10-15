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
