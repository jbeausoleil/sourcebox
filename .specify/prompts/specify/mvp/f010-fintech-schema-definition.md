# Feature Specification Prompt: F010 - Fintech Schema Definition (JSON)

## Feature Metadata
- **Feature ID**: F010
- **Name**: Fintech Schema Definition (JSON)
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F007 (Schema specification), F008 (Schema parser)

## Constitutional Alignment

### Core Principles
- ✅ **Verticalized > Generic (NON-NEGOTIABLE)**: Fintech data MUST look like fintech
- ✅ **Realism > Convenience**: Real-world distributions, proper terminology, accurate relationships
- ✅ **Developer-First**: Schema demonstrates realistic loan management system

### Technical Constraints
- ✅ **Performance**: 4,950 records should seed in < 30 seconds
- ✅ **Data Quality**: Realistic distributions, proper relationships

### Development Practices
- ✅ **Quality Bar**: If data isn't 10x more realistic than Faker, don't ship it

## User Story
**US-MVP-005**: "As a backend developer building a fintech demo, I want realistic loan data with proper credit scores, interest rates, and payment histories so my demos look professional to investors and customers."

## Problem Statement
Generic fake data from Faker.js makes fintech demos look unprofessional. "John Doe" with employee ID "123" borrowing "$1,000.00" at "5%" interest doesn't reflect real-world lending patterns. A realistic fintech schema must include:
- **Realistic credit scores**: Bell curve distribution (300-850, mean 680)
- **Realistic loan amounts**: Log-normal distribution ($5K-$500K, median $50K)
- **Realistic interest rates**: Weighted by creditworthiness (3-15%)
- **Realistic loan statuses**: Weighted frequency (active 70%, paid 25%, delinquent 5%)
- **Proper relationships**: Borrowers → Loans → Payments
- **Industry terminology**: APR, not "interest percentage"

Without realistic fintech data, developers waste hours creating manual test data or use generic data that doesn't reflect production scenarios.

## Solution Overview
Create a comprehensive fintech loans schema (`schemas/fintech-loans.json`) with 4 tables: borrowers, loans, payments, credit_scores. Define custom generators for credit scores (normal distribution), loan amounts (log-normal), interest rates (weighted ranges), and loan statuses (weighted values). Ensure proper foreign key relationships and realistic data distributions. Total: 4,950 records across all tables.

## Detailed Requirements

### Acceptance Criteria
1. **Schema File Created**: `schemas/fintech-loans.json` following F007 specification
2. **4 Tables Defined**:
   - `borrowers` (250 records)
   - `loans` (1,000 records)
   - `payments` (3,500 records)
   - `credit_scores` (250 records)
3. **Foreign Key Relationships**:
   - `borrowers.id` → `loans.borrower_id`
   - `loans.id` → `payments.loan_id`
   - `borrowers.id` → `credit_scores.borrower_id`
4. **Realistic Columns**: id, name, email, credit_score, amount, interest_rate, status, payment_date, created_at
5. **Custom Generators Defined**:
   - `credit_score` (normal distribution: mean=680, std=80, range 300-850)
   - `loan_amount` (log-normal: median=$50K, range $5K-$500K)
   - `interest_rate` (weighted: 3-6%=60%, 6-10%=30%, 10-15%=10%)
   - `loan_status` (weighted: active=70%, paid_off=25%, delinquent=5%)
6. **Schema Validates**: Passes F008 schema parser validation
7. **Total Expected Records**: 4,950 (250+1000+3500+250)

### Technical Specifications

#### Schema Structure

**File**: `schemas/fintech-loans.json`

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
    "tags": ["loans", "credit", "payments", "borrowers", "fintech"]
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
          "name": "phone",
          "type": "varchar(20)",
          "generator": "phone",
          "nullable": true
        },
        {
          "name": "date_of_birth",
          "type": "date",
          "generator": "date_of_birth",
          "generator_params": {
            "min_age": 21,
            "max_age": 75
          },
          "nullable": false
        },
        {
          "name": "ssn_last_4",
          "type": "varchar(4)",
          "generator": "ssn_last_4",
          "nullable": true
        },
        {
          "name": "annual_income",
          "type": "decimal(12,2)",
          "generator": "annual_income",
          "generator_params": {
            "distribution": "lognormal",
            "median": 65000,
            "min": 25000,
            "max": 500000
          },
          "nullable": false
        },
        {
          "name": "employment_status",
          "type": "varchar(50)",
          "generator": "employment_status",
          "generator_params": {
            "values": [
              {"value": "employed", "weight": 75},
              {"value": "self_employed", "weight": 15},
              {"value": "unemployed", "weight": 5},
              {"value": "retired", "weight": 5}
            ]
          },
          "nullable": false
        },
        {
          "name": "created_at",
          "type": "timestamp",
          "generator": "timestamp_past",
          "generator_params": {
            "days_ago_min": 730,
            "days_ago_max": 30
          },
          "default": "CURRENT_TIMESTAMP",
          "nullable": false
        }
      ],
      "indexes": [
        {
          "name": "idx_email",
          "columns": ["email"],
          "unique": true
        },
        {
          "name": "idx_last_name",
          "columns": ["last_name"]
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
          "auto_increment": true,
          "nullable": false
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
          "name": "loan_number",
          "type": "varchar(50)",
          "generator": "loan_number",
          "unique": true,
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
          "name": "term_months",
          "type": "int",
          "generator": "loan_term",
          "generator_params": {
            "values": [
              {"value": 12, "weight": 15},
              {"value": 24, "weight": 20},
              {"value": 36, "weight": 30},
              {"value": 48, "weight": 20},
              {"value": 60, "weight": 15}
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
          "name": "funded_date",
          "type": "date",
          "generator": "date_past",
          "generator_params": {
            "days_ago_min": 730,
            "days_ago_max": 1
          },
          "nullable": false
        },
        {
          "name": "maturity_date",
          "type": "date",
          "generator": "maturity_date",
          "generator_params": {
            "based_on": "funded_date",
            "months_forward": "term_months"
          },
          "nullable": false
        },
        {
          "name": "created_at",
          "type": "timestamp",
          "default": "CURRENT_TIMESTAMP",
          "nullable": false
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
        },
        {
          "name": "idx_funded_date",
          "columns": ["funded_date"]
        }
      ]
    },
    {
      "name": "payments",
      "description": "Loan payment history",
      "record_count": 3500,
      "columns": [
        {
          "name": "id",
          "type": "int",
          "primary_key": true,
          "auto_increment": true,
          "nullable": false
        },
        {
          "name": "loan_id",
          "type": "int",
          "foreign_key": {
            "table": "loans",
            "column": "id",
            "on_delete": "CASCADE"
          },
          "nullable": false
        },
        {
          "name": "payment_number",
          "type": "int",
          "generator": "payment_number",
          "nullable": false
        },
        {
          "name": "payment_date",
          "type": "date",
          "generator": "payment_date",
          "generator_params": {
            "based_on_loan": "funded_date",
            "monthly_interval": true
          },
          "nullable": false
        },
        {
          "name": "amount",
          "type": "decimal(10,2)",
          "generator": "payment_amount",
          "generator_params": {
            "calculated_from": "loan.amount",
            "term": "loan.term_months"
          },
          "nullable": false
        },
        {
          "name": "principal",
          "type": "decimal(10,2)",
          "generator": "payment_principal",
          "nullable": false
        },
        {
          "name": "interest",
          "type": "decimal(10,2)",
          "generator": "payment_interest",
          "nullable": false
        },
        {
          "name": "status",
          "type": "varchar(20)",
          "generator": "payment_status",
          "generator_params": {
            "values": [
              {"value": "paid", "weight": 85},
              {"value": "pending", "weight": 10},
              {"value": "late", "weight": 4},
              {"value": "missed", "weight": 1}
            ]
          },
          "nullable": false
        },
        {
          "name": "created_at",
          "type": "timestamp",
          "default": "CURRENT_TIMESTAMP",
          "nullable": false
        }
      ],
      "indexes": [
        {
          "name": "idx_loan",
          "columns": ["loan_id"]
        },
        {
          "name": "idx_payment_date",
          "columns": ["payment_date"]
        },
        {
          "name": "idx_status",
          "columns": ["status"]
        }
      ]
    },
    {
      "name": "credit_scores",
      "description": "Credit score history for borrowers",
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
          "name": "score",
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
          "name": "bureau",
          "type": "varchar(50)",
          "generator": "credit_bureau",
          "generator_params": {
            "values": [
              {"value": "Experian", "weight": 33},
              {"value": "Equifax", "weight": 33},
              {"value": "TransUnion", "weight": 34}
            ]
          },
          "nullable": false
        },
        {
          "name": "checked_date",
          "type": "date",
          "generator": "date_past",
          "generator_params": {
            "days_ago_min": 365,
            "days_ago_max": 1
          },
          "nullable": false
        },
        {
          "name": "created_at",
          "type": "timestamp",
          "default": "CURRENT_TIMESTAMP",
          "nullable": false
        }
      ],
      "indexes": [
        {
          "name": "idx_borrower",
          "columns": ["borrower_id"]
        },
        {
          "name": "idx_score",
          "columns": ["score"]
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
      "description": "Each loan belongs to one borrower (borrowers can have multiple loans)"
    },
    {
      "from_table": "payments",
      "from_column": "loan_id",
      "to_table": "loans",
      "to_column": "id",
      "relationship_type": "many_to_one",
      "description": "Each payment belongs to one loan (loans have multiple payments)"
    },
    {
      "from_table": "credit_scores",
      "from_column": "borrower_id",
      "to_table": "borrowers",
      "to_column": "id",
      "relationship_type": "one_to_one",
      "description": "Each borrower has one current credit score"
    }
  ],
  "generation_order": ["borrowers", "loans", "payments", "credit_scores"],
  "validation_rules": [
    {
      "rule": "loan_amount_credit_correlation",
      "description": "Higher credit scores should correlate with lower interest rates",
      "enforcement": "soft"
    },
    {
      "rule": "payment_amounts_match_loan",
      "description": "Sum of payment principal should approximately equal loan amount",
      "enforcement": "soft"
    },
    {
      "rule": "delinquent_loans_have_missed_payments",
      "description": "Loans with status 'delinquent' should have some missed/late payments",
      "enforcement": "soft"
    }
  ]
}
```

### Data Realism Requirements

**Credit Scores** (Normal Distribution):
- Mean: 680 (US average)
- Standard deviation: 80
- Range: 300-850 (FICO scale)
- Distribution should show bell curve when visualized

**Loan Amounts** (Log-Normal Distribution):
- Median: $50,000
- Range: $5,000 - $500,000
- More common: $20K-$100K
- Less common: $200K+
- Reflects real lending patterns (many small loans, few large loans)

**Interest Rates** (Weighted):
- 3-6%: 60% (good credit, prime rates)
- 6-10%: 30% (average credit)
- 10-15%: 10% (subprime)
- Correlate with credit scores (higher score = lower rate)

**Loan Status** (Weighted):
- Active: 70% (most loans are current)
- Paid off: 25% (some loans complete)
- Delinquent: 5% (realistic default rate)

### Performance Considerations
- 4,950 total records
- Target: < 30 seconds seeding time
- Foreign key constraints maintained
- Indexes for common queries

### Testing Strategy

**Schema Validation**:
```bash
# Validate schema parses correctly
go test ./pkg/schema/... -v

# Specifically test fintech schema
go test ./pkg/schema/ -run TestLoadFintechSchema
```

**Manual Review**:
1. Read schema JSON - verify it's readable and well-documented
2. Check all tables have descriptions
3. Verify foreign keys are correct
4. Verify generators have appropriate parameters
5. Verify record counts make sense (borrowers < loans < payments)

**Future Integration Testing** (F020):
- Generate data from schema
- Verify distributions match specifications
- Verify foreign key integrity
- Query sample data to verify realism

## Dependencies
- **Upstream**:
  - F007 (Schema specification defined)
  - F008 (Schema parser exists to validate)
- **Downstream**:
  - F012 (Custom Fintech Generators) implements generators defined here
  - F013 (Data Generation Engine) uses this schema
  - F020 (Integration Tests) validates this schema

## Deliverables
1. Fintech schema JSON file (`schemas/fintech-loans.json`)
2. Schema validates successfully with F008 parser
3. Documentation of data distributions and realism
4. Example queries demonstrating realistic use cases

## Success Criteria
- ✅ Schema file exists and is valid JSON
- ✅ Passes F008 schema parser validation
- ✅ All 4 tables defined with proper relationships
- ✅ Custom generators defined with realistic parameters
- ✅ Total record count = 4,950
- ✅ Schema demonstrates verticalized fintech data (NOT generic Faker data)

## Anti-Patterns to Avoid
- ❌ Generic data (e.g., "John Doe", "Acme Corp", employee ID "123")
- ❌ Unrealistic distributions (uniform random for loan amounts)
- ❌ Missing foreign key relationships
- ❌ Vague generator parameters (must be specific)
- ❌ Missing industry terminology (use "APR", not "percentage")
- ❌ Too few records (won't demonstrate realistic patterns)
- ❌ Too many records (will violate <30s seeding time)

## Implementation Notes
- Schema format follows F007 specification exactly
- Custom generators will be implemented in F012
- Data generation engine (F013) will use this schema
- This is the first of 3 schemas (healthcare F016, retail F018)
- Quality bar: 10x more realistic than Faker or don't ship

## TDD Requirements
**Not applicable for schema definition** - This is declarative data. However, verify schema quality by:
1. Validating with F008 parser
2. Reviewing distributions for realism
3. Checking foreign key relationships
4. Future: Generate sample data and verify realism (F020)

## Related Constitution Sections
- **Verticalized > Generic (Principle I, NON-NEGOTIABLE)**: Fintech must look like fintech
- **Realism Standard**: Real-world distributions, industry terminology, proper relationships
- **Quality Bar**: If not 10x more realistic than Faker, don't ship it
- **Performance (Technical Constraint 1)**: <30s seeding for 4,950 records
