# Feature Specification Prompt: F041 - Incremental Data Management

## Feature Metadata
- **Feature ID**: F041
- **Name**: Incremental Data Management (Modify Command)
- **Category**: CLI Tool
- **Phase**: Phase 2 (Post-MVP Enhancement)
- **Priority**: P1 (High Value)
- **Estimated Effort**: Large (5-7 days)
- **Dependencies**: F013 (Data Generation Engine), F021 (Seed Command), F023 (MySQL Connector), F024 (PostgreSQL Connector)

## Constitutional Alignment

### Core Principles
- ✅ **Developer-First Design**: Enables iterative test data workflows without full re-seeding
- ✅ **Speed > Features**: Incremental operations faster than full re-seed
- ✅ **Local-First**: No cloud dependencies, works offline

### Technical Constraints
- ✅ **Performance**: Modify operations complete in <10s for 1,000 records
- ✅ **TDD Required**: Core modification logic requires test coverage
- ✅ **Code Quality**: >80% test coverage for modification engine

### Development Practices
- ✅ **TDD Required for Core Functionality**: Modification engine is core logic
- ✅ **Fail Gracefully**: Clear errors for invalid operations

## User Story

**US-P2-001**: "As a developer, I want to add, delete, or update records in already-seeded tables without re-seeding the entire database, so I can iterate quickly on test scenarios and simulate data lifecycle events."

## Problem Statement

SourceBox currently only supports one-shot database seeding. Once a database is seeded, developers have limited options:
1. **Drop and re-seed** - Slow (30s+ for large schemas), loses custom test data
2. **Manual SQL** - Time-consuming, error-prone, defeats purpose of SourceBox
3. **Live with what you have** - Limits test scenario coverage

**Real-world use cases SourceBox cannot handle today:**

1. **QA Engineer**: "I seeded 1,000 users for load testing, but now I need 100 more - don't want to wait 30s to re-seed everything"
2. **Backend Developer**: "I want to simulate 6 months of user churn - delete 30% of users and cascade to their transactions"
3. **Data Engineer**: "Testing my ETL pipeline - need to add new records daily for a week to simulate incremental loads"
4. **Full-stack Developer**: "I accidentally corrupted the `transactions` table during testing - need to regenerate just that table"
5. **Performance Tester**: "Need to grow database from 1K to 10K records gradually to test scaling behavior"

Without incremental data management, SourceBox is a **one-shot tool**, not a **test data platform**.

## Solution Overview

Implement `sourcebox modify` command to perform incremental database operations (append, delete, update, replace, truncate) on already-seeded tables. Operations respect foreign key constraints, maintain referential integrity, and provide clear error messages. Supports both MySQL and PostgreSQL.

### Command Syntax

```bash
sourcebox modify <schema> <table> [operation] [flags]
```

### Operations

1. **Append**: Add new records to existing table
2. **Delete**: Remove records (random or by criteria)
3. **Update**: Regenerate specific fields in existing records
4. **Replace**: Drop table and regenerate with new data
5. **Truncate**: Remove all records, optionally reseed

## Detailed Requirements

### Acceptance Criteria

1. **Command Structure**: `sourcebox modify` with subcommands for each operation type
2. **Append Operation**:
   - Add N new records to table
   - Maintains foreign key constraints
   - Auto-increments primary keys correctly
   - Validates against schema definition
3. **Delete Operation**:
   - Remove N random records OR records matching criteria
   - Respects foreign key constraints (CASCADE, SET NULL, RESTRICT)
   - Reports number of records deleted
   - Supports `--where` clause for targeted deletion
4. **Update Operation**:
   - Regenerate specified fields for N records or X% of table
   - Supports field list: `--fields email,phone,address`
   - Maintains data consistency (e.g., don't orphan foreign keys)
   - Reports number of records updated
5. **Replace Operation**:
   - Drops table and regenerates with fresh data
   - Faster than delete + append for full table refresh
   - Validates no dependent foreign keys exist (or CASCADE them)
6. **Truncate Operation**:
   - Removes all records from table
   - Optional `--reseed` flag to repopulate after truncate
   - Handles foreign key constraints gracefully
7. **Multi-table Support**: Operate on multiple tables in one command
8. **Dry-run Mode**: Preview operation without executing
9. **Transaction Safety**: All operations wrapped in transactions (rollback on error)
10. **Clear Errors**: Actionable error messages for constraint violations

### Command Examples

```bash
# Append 100 new users
sourcebox modify fintech users --append --count 100

# Delete 50 random transactions
sourcebox modify fintech transactions --delete --count 50

# Delete users created before a date (respects CASCADE foreign keys)
sourcebox modify fintech users --delete --where "created_at < '2024-01-01'"

# Update email and phone for 20 random users
sourcebox modify fintech users --update email,phone --count 20

# Update email for 15% of all users (simulate churn)
sourcebox modify fintech users --update email --percentage 15

# Replace entire transactions table with new data (1000 records)
sourcebox modify fintech transactions --replace --count 1000

# Truncate loans table and reseed with 500 records
sourcebox modify fintech loans --truncate --reseed --count 500

# Multi-table: truncate and reseed users (500) and loans (2000)
sourcebox modify fintech users,loans --truncate --reseed --count 500,2000

# Dry-run: preview delete operation
sourcebox modify fintech users --delete --count 100 --dry-run

# Verbose: show detailed operation logs
sourcebox modify fintech transactions --append --count 200 --verbose
```

### Technical Specifications

#### Package Structure

```
pkg/modify/
├── engine.go           # Core modification engine
├── append.go           # Append operation logic
├── delete.go           # Delete operation logic
├── update.go           # Update operation logic
├── replace.go          # Replace operation logic
├── truncate.go         # Truncate operation logic
├── constraints.go      # Foreign key constraint handling
└── modify_test.go      # TDD test suite
```

#### Operation Types (Go Enum)

```go
type OperationType string

const (
    OperationAppend   OperationType = "append"
    OperationDelete   OperationType = "delete"
    OperationUpdate   OperationType = "update"
    OperationReplace  OperationType = "replace"
    OperationTruncate OperationType = "truncate"
)
```

#### Modify Command Flags

```
Global Flags (apply to all operations):
  --database string       Database connection string (required)
  --dry-run              Preview operation without executing
  --verbose              Show detailed operation logs
  --quiet                Suppress non-error output

Append Flags:
  --count int            Number of records to append (required)

Delete Flags:
  --count int            Number of records to delete (random)
  --where string         SQL WHERE clause for targeted deletion
  --percentage int       Delete X% of records (alternative to --count)

Update Flags:
  --fields string        Comma-separated list of fields to regenerate (required)
  --count int            Number of records to update (random)
  --percentage int       Update X% of records (alternative to --count)
  --where string         SQL WHERE clause for targeted update

Replace Flags:
  --count int            Number of records in regenerated table (required)

Truncate Flags:
  --reseed               Repopulate table after truncation
  --count int            Number of records to reseed (if --reseed specified)
```

#### Constraint Handling

**Foreign Key Validation**:
- Before DELETE: Check for dependent records (unless CASCADE specified in schema)
- Before TRUNCATE: Validate no orphaned foreign keys created
- Before REPLACE: Handle cascading deletes for dependent tables

**Primary Key Handling**:
- APPEND: Auto-increment from MAX(id) + 1
- UPDATE: Never modify primary key values
- REPLACE: Regenerate all primary keys

#### Transaction Management

All operations wrapped in database transactions:

```go
func (e *ModifyEngine) ExecuteOperation(op Operation) error {
    tx, err := e.db.Begin()
    if err != nil {
        return fmt.Errorf("failed to start transaction: %w", err)
    }
    defer tx.Rollback() // Rollback if not committed

    // Execute operation
    if err := e.execute(tx, op); err != nil {
        return fmt.Errorf("operation failed: %w", err)
    }

    // Commit transaction
    if err := tx.Commit(); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}
```

#### Error Messages

❌ **Bad**: "modification error"
✅ **Good**: "cannot delete 50 records from 'users' table: would orphan 120 records in 'loans' table (foreign key constraint). Use CASCADE or delete dependent records first."

❌ **Bad**: "update failed"
✅ **Good**: "cannot update 'email' field: column not found in schema definition. Available fields: id, first_name, last_name, email_address"

### Validation Rules

1. **Schema Validation**: Table and columns must exist in schema definition
2. **Constraint Validation**: Operations must respect foreign key constraints
3. **Data Type Validation**: Generated data must match column data types
4. **Count Validation**: --count must be positive and <= table row count (for delete/update)
5. **Percentage Validation**: --percentage must be 1-100
6. **Field Validation**: --fields must reference existing columns in schema
7. **Operation Compatibility**: Can't combine conflicting flags (e.g., --append + --delete)

### Performance Targets

- **Append 1,000 records**: <10 seconds
- **Delete 1,000 records**: <5 seconds
- **Update 1,000 records**: <15 seconds (depends on field complexity)
- **Replace table (1,000 records)**: <20 seconds
- **Truncate + reseed (1,000 records)**: <15 seconds

*(Measured on 2020 MacBook Pro - constitutional hardware baseline)*

## Dependencies

### Upstream Dependencies
- **F013**: Data Generation Engine (reuse generators for append/update)
- **F021**: Seed Command (command structure patterns)
- **F023**: MySQL Connector (database operations)
- **F024**: PostgreSQL Connector (database operations)
- **F008**: Schema Parser (validate operations against schema)

### Downstream Impact
- **F025**: Integration tests will cover modify operations
- **F037-F038**: Documentation will include modify command examples
- **F039**: Alpha testers will provide feedback on modify workflows

## Deliverables

1. Modify command implementation (`cmd/sourcebox/cmd/modify.go`)
2. Modification engine package (`pkg/modify/`)
3. Unit tests with >80% coverage (`pkg/modify/modify_test.go`)
4. Integration tests with TestContainers (MySQL + PostgreSQL)
5. CLI help text and examples
6. Documentation updates (README.md, docs site)

## Success Criteria

- ✅ All five operations (append, delete, update, replace, truncate) work correctly
- ✅ Foreign key constraints respected for all operations
- ✅ Transactions rollback on error (no partial modifications)
- ✅ Clear error messages for constraint violations
- ✅ Performance targets met (<10s for 1,000 record operations)
- ✅ Unit tests pass: `go test ./pkg/modify/... -v`
- ✅ Integration tests pass with MySQL and PostgreSQL
- ✅ Dry-run mode accurately previews operations
- ✅ Multi-table operations work correctly
- ✅ Help text comprehensive: `sourcebox modify --help`

## Anti-Patterns to Avoid

- ❌ Ignoring foreign key constraints (causes orphaned records)
- ❌ Not wrapping operations in transactions (partial failures corrupt data)
- ❌ Poor error messages (users don't understand why operation failed)
- ❌ Modifying primary keys during UPDATE (breaks referential integrity)
- ❌ Allowing percentage > 100 or count > table size (confusing UX)
- ❌ Generating inconsistent data (e.g., future dates for created_at)
- ❌ Slow operations (>30s defeats purpose of incremental approach)

## Use Case Examples

### Use Case 1: Gradual Database Growth (Performance Testing)

```bash
# Day 1: Seed initial dataset
sourcebox seed fintech --count 1000

# Day 2: Add 500 more users
sourcebox modify fintech users --append --count 500

# Day 3: Add 1000 more transactions
sourcebox modify fintech transactions --append --count 1000

# Day 4: Add 2000 more transactions
sourcebox modify fintech transactions --append --count 2000

# Result: Simulated 4 days of organic growth without re-seeding
```

### Use Case 2: User Churn Simulation (Backend Testing)

```bash
# Simulate 6 months of user churn (30% attrition)
sourcebox modify fintech users --delete --percentage 30

# Result: Removes 30% of users, CASCADE deletes their loans/transactions
```

### Use Case 3: Data Corruption Recovery (Development Workflow)

```bash
# Developer accidentally corrupts transactions table
# Replace just that table without touching users/loans
sourcebox modify fintech transactions --replace --count 5000

# Result: Fresh transactions data in <20s, users/loans untouched
```

### Use Case 4: ETL Pipeline Testing (Data Engineering)

```bash
# Week 1: Initial load
sourcebox seed fintech --count 1000

# Week 2: Incremental load
sourcebox modify fintech transactions --append --count 200

# Week 3: Incremental load
sourcebox modify fintech transactions --append --count 300

# Result: Test incremental ETL without full refresh
```

### Use Case 5: GDPR Testing (Compliance)

```bash
# Test user data deletion (GDPR "right to be forgotten")
sourcebox modify fintech users --delete --where "email = 'test@example.com'"

# Result: Deletes specific user, CASCADE removes their personal data
```

## TDD Requirements

**REQUIRED for this feature** - Follow strict TDD workflow:

1. **Write test first**: Test for append operation
2. **Run test**: Should fail (modify engine not implemented)
3. **Implement**: Basic modify engine structure
4. **Run test**: Should pass
5. **Refactor**: Clean up code
6. **Repeat**: For each operation type (delete, update, replace, truncate)

**Coverage target**: >80% for all modification code

### Test Strategy

```go
// Test append operation
func TestModifyAppend(t *testing.T) {
    // Setup: Seed initial data
    // Execute: Append 100 records
    // Assert: Total records = initial + 100
    // Assert: New records have correct schema
    // Assert: Foreign keys valid
}

// Test delete with CASCADE
func TestModifyDeleteCascade(t *testing.T) {
    // Setup: Seed users + loans with FK
    // Execute: Delete 10 users
    // Assert: Users deleted
    // Assert: Dependent loans also deleted (CASCADE)
}

// Test update operation
func TestModifyUpdate(t *testing.T) {
    // Setup: Seed initial data
    // Execute: Update email field for 20 records
    // Assert: Exactly 20 records updated
    // Assert: Email format valid
    // Assert: Other fields unchanged
}

// Test foreign key constraint violation
func TestModifyDeleteConstraintViolation(t *testing.T) {
    // Setup: Seed users + loans (FK without CASCADE)
    // Execute: Try to delete users with loans
    // Assert: Operation fails with clear error
    // Assert: No records deleted (transaction rollback)
}
```

## Related Constitution Sections

- **Developer-First Design (Principle VI)**: Incremental workflows reduce iteration time
- **Speed > Features (Principle II)**: Modify operations faster than full re-seed
- **TDD Required for Core Functionality (Development Practice 1)**: Modify engine is core logic
- **Fail Gracefully (UX Principle 4)**: Clear errors for constraint violations
- **Code Quality Standards (Technical Constraint 5)**: >80% coverage required
- **Performance Gates (Technical Constraint 1)**: <10s for 1,000 record operations

## Phase 2 Positioning

This feature is a **high-value Phase 2 enhancement** that transforms SourceBox from:
- **One-shot seeder** → **Test data platform**
- **Static datasets** → **Dynamic data lifecycle management**
- **All-or-nothing** → **Surgical, iterative workflows**

**Market differentiation**: Most test data tools only support initial seeding. SourceBox would be unique in supporting full data lifecycle management.
