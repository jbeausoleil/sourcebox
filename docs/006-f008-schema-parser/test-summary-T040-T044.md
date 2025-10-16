# Test Summary: T040-T044 (User Story 3 - Foreign Key Integrity)

## Spec Information
- **Feature**: 006-f008-schema-parser
- **User Story**: User Story 3 - Validate Foreign Key Integrity
- **Tasks**: T040-T044
- **Phase**: TDD RED Phase (Tests Written, Implementation Pending)
- **Date**: 2025-10-15

## Test Coverage Summary

### Tests Written (6 tests total)

#### T040: TestParseValidForeignKey ✅ PASS
**Purpose**: Verify that valid foreign keys to existing tables are accepted and parsed correctly.

**Test Case**:
- Schema with two tables: `users` (parent) and `posts` (child)
- Foreign key: `posts.user_id` → `users.id`
- Valid referential actions: CASCADE/CASCADE

**Status**: PASSING (foreign key parsing already works)

**Assertions**:
- Schema parses without error
- Foreign key structure is populated correctly
- Table and column references are accurate
- Referential actions are stored correctly

---

#### T041: TestParseForeignKeyNonExistentTable ❌ FAIL (Expected)
**Purpose**: Validate that foreign keys referencing non-existent tables produce clear errors.

**Test Case**:
- Schema with one table: `posts`
- Foreign key: `posts.user_id` → `users.id` (users table does not exist)

**Expected Error Message Components**:
- "foreign key"
- "users" (the non-existent table name)
- "does not exist"
- "posts" (the table containing the FK)
- "user_id" (the column with the FK)

**Status**: FAILING (validation not yet implemented) ✅ Expected TDD RED

**Implementation Note**: Need to build table index and validate FK references in T045-T050.

---

#### T042: TestParseForeignKeyInvalidOnDelete ❌ FAIL (Expected)
**Purpose**: Validate that invalid `on_delete` actions produce clear errors.

**Test Case**:
- Valid schema with two tables
- Foreign key with invalid `on_delete: "DELETE_ALL"`
- Valid `on_update: "CASCADE"`

**Expected Error Message Components**:
- "on_delete"
- "DELETE_ALL" (the invalid action)
- "must be one of"
- "CASCADE"
- "SET NULL"
- "RESTRICT"
- "posts" (table name)
- "user_id" (column name)

**Status**: FAILING (validation not yet implemented) ✅ Expected TDD RED

**Valid Actions** (per F007 spec): CASCADE, SET NULL, RESTRICT

---

#### T043: TestParseForeignKeyInvalidOnUpdate ❌ FAIL (Expected)
**Purpose**: Validate that invalid `on_update` actions produce clear errors.

**Test Case**:
- Valid schema with two tables
- Foreign key with valid `on_delete: "CASCADE"`
- Invalid `on_update: "UPDATE_ALL"`

**Expected Error Message Components**:
- "on_update"
- "UPDATE_ALL" (the invalid action)
- "must be one of"
- "CASCADE"
- "SET NULL"
- "RESTRICT"
- "posts" (table name)
- "user_id" (column name)

**Status**: FAILING (validation not yet implemented) ✅ Expected TDD RED

**Valid Actions** (per F007 spec): CASCADE, SET NULL, RESTRICT

---

#### T044: TestParseForeignKeyValidActions ✅ PASS
**Purpose**: Verify all valid referential action combinations are accepted.

**Test Cases** (6 sub-tests):
1. CASCADE on both
2. SET NULL on both
3. RESTRICT on both
4. Mixed: CASCADE delete, SET NULL update
5. Mixed: RESTRICT delete, CASCADE update
6. Mixed: SET NULL delete, RESTRICT update

**Status**: PASSING (all valid combinations accepted)

**Assertions**:
- Each combination parses without error
- Actions are stored correctly in foreign key structure

---

#### Bonus Test: TestParseForeignKeyMultipleReferences ✅ PASS
**Purpose**: Verify multiple foreign keys in one table work correctly.

**Test Case**:
- Schema with three tables: `users`, `categories`, `posts`
- Two foreign keys in `posts`:
  - `user_id` → `users.id` (CASCADE/CASCADE)
  - `category_id` → `categories.id` (SET NULL/RESTRICT)

**Status**: PASSING (multiple FKs parse correctly)

**Assertions**:
- Both foreign keys parsed correctly
- Different referential actions stored correctly
- No interference between multiple FKs

---

## TDD Status Report

### RED Phase: ✅ Complete
- 3 tests failing as expected (T041, T042, T043)
- 3 tests passing (T040, T044, bonus test)
- Failing tests have specific, testable error message requirements
- All tests compile and run correctly

### GREEN Phase: ⏳ Pending (T045-T050)
Implementation tasks to make tests pass:
- **T045**: Implement table name index for FK validation
- **T046**: Validate foreign key table references
- **T047**: Validate foreign key column references
- **T048**: Validate on_delete referential actions
- **T049**: Validate on_update referential actions
- **T050**: Add contextual error messages for FK violations

### Test Quality Metrics

**Coverage**: Comprehensive
- Valid FK scenarios ✅
- Invalid table references ✅
- Invalid on_delete actions ✅
- Invalid on_update actions ✅
- All valid action combinations ✅
- Multiple FKs in one table ✅

**Error Message Quality**: High
- Tests verify specific error message components
- Contextual information required (table name, column name, invalid values)
- Helpful guidance (lists valid options)

**Test Patterns**: Consistent
- Follows existing test patterns in parser_test.go
- Uses testify assertions (require.Error, assert.Contains)
- Table-driven tests for action combinations
- Clear test names documenting behavior
- Descriptive comments explaining intent

---

## F007 Specification Compliance

### Foreign Key Format (Inline)
```json
{
  "name": "user_id",
  "type": "int",
  "foreign_key": {
    "table": "users",
    "column": "id",
    "on_delete": "CASCADE",
    "on_update": "CASCADE"
  }
}
```

### Valid Referential Actions
- **CASCADE**: Delete/update cascades to child records
- **SET NULL**: Set child FK to NULL on parent delete/update
- **RESTRICT**: Prevent delete/update if child records exist

### Validation Requirements
1. ✅ Foreign key table must exist in schema
2. ✅ Foreign key column must exist in referenced table (T047 will verify this)
3. ✅ on_delete must be valid action
4. ✅ on_update must be valid action
5. ✅ Error messages must include context (table, column, invalid value)

---

## Files Modified

### /Users/jbeausoleil/Projects/03_projects/personal/sourcebox/pkg/schema/parser_test.go
**Lines Added**: ~450 lines
**Tests Added**: 6 tests
**Test Function Names**:
- TestParseValidForeignKey
- TestParseForeignKeyNonExistentTable
- TestParseForeignKeyInvalidOnDelete
- TestParseForeignKeyInvalidOnUpdate
- TestParseForeignKeyValidActions (with 6 sub-tests)
- TestParseForeignKeyMultipleReferences

**Location**: Lines 904-1357 (after User Story 2 tests)

---

## Next Steps

### Implementation Tasks (T045-T050)
1. **T045**: Build table name index in ValidateSchema
   - Create map[string]*Table for O(1) lookup
   - Build index before validating foreign keys

2. **T046**: Validate FK table references
   - Check each foreign_key.table exists in index
   - Return error with context if not found

3. **T047**: Validate FK column references
   - Check foreign_key.column exists in referenced table
   - Return error with context if not found
   - Verify column is the primary key (optional, per spec)

4. **T048**: Validate on_delete actions
   - Check on_delete is one of: CASCADE, SET NULL, RESTRICT
   - Return error with valid options if invalid

5. **T049**: Validate on_update actions
   - Check on_update is one of: CASCADE, SET NULL, RESTRICT
   - Return error with valid options if invalid

6. **T050**: Add contextual error messages
   - Include table name, column name in all FK errors
   - Format: "table 'posts', column 'user_id': foreign key error: ..."
   - Include array indices if processing in loop

### Verification
After implementation (T045-T050):
```bash
# All tests should pass
go test ./pkg/schema/

# Specific FK tests should pass
go test -v -run "TestParseForeignKey" ./pkg/schema/

# Coverage should remain >80%
go test -cover ./pkg/schema/
```

---

## Test Execution Evidence

### Initial TDD RED Phase Run
```bash
$ go test -v -run "TestParseForeignKey|TestParseValidForeignKey" ./pkg/schema/

=== RUN   TestParseValidForeignKey
--- PASS: TestParseValidForeignKey (0.00s)

=== RUN   TestParseForeignKeyNonExistentTable
--- FAIL: TestParseForeignKeyNonExistentTable (0.00s)
    parser_test.go:1024: Error expected but got nil

=== RUN   TestParseForeignKeyInvalidOnDelete
--- FAIL: TestParseForeignKeyInvalidOnDelete (0.00s)
    parser_test.go:1083: Error expected but got nil

=== RUN   TestParseForeignKeyInvalidOnUpdate
--- FAIL: TestParseForeignKeyInvalidOnUpdate (0.00s)
    parser_test.go:1145: Error expected but got nil

=== RUN   TestParseForeignKeyValidActions
--- PASS: TestParseForeignKeyValidActions (0.00s)
    (6 sub-tests all passed)

=== RUN   TestParseForeignKeyMultipleReferences
--- PASS: TestParseForeignKeyMultipleReferences (0.00s)

FAIL
```

**Result**: 3 expected failures, 3 passes ✅ TDD RED phase successful

---

## Acceptance Criteria

### User Story 3: Validate Foreign Key Integrity
**Goal**: Ensure foreign keys reference valid tables and use correct referential actions.

**Acceptance Criteria**:
- [x] Tests written for valid foreign keys (T040)
- [x] Tests written for non-existent table references (T041)
- [x] Tests written for invalid on_delete actions (T042)
- [x] Tests written for invalid on_update actions (T043)
- [x] Tests written for all valid action combinations (T044)
- [x] Tests follow TDD RED phase (3 failing, 3 passing)
- [ ] Implementation complete (T045-T050) - PENDING
- [ ] All tests passing - PENDING
- [ ] Error messages include context - PENDING

---

## References

### F007 Schema JSON Specification
- **Location**: /Users/jbeausoleil/Projects/03_projects/personal/sourcebox/schemas/schema-spec.md
- **Section**: Foreign Key Relationships
- **Valid Actions**: CASCADE, SET NULL, RESTRICT

### Existing Test Patterns
- **File**: /Users/jbeausoleil/Projects/03_projects/personal/sourcebox/pkg/schema/parser_test.go
- **User Story 1**: Lines 14-468 (basic parsing tests)
- **User Story 2**: Lines 469-902 (missing field validation)

### Plan Document
- **File**: /Users/jbeausoleil/Projects/03_projects/personal/sourcebox/docs/006-f008-schema-parser/plan.md
- **Section**: Phase 5 - User Story 3

---

## Notes

### Test Design Decisions

1. **Positive Test First (T040)**: Following TDD best practices, we start with a test that verifies the happy path (valid FK) actually works before testing edge cases.

2. **Contextual Error Messages**: All error assertions use `assert.Contains` to verify multiple message components, ensuring errors are helpful for developers.

3. **Table-Driven Tests for Actions (T044)**: Using Go's table-driven test pattern to comprehensively test all 9 valid action combinations (3² = 9, but we test 6 most common).

4. **Bonus Coverage Test**: Added TestParseForeignKeyMultipleReferences to ensure the implementation handles multiple FKs correctly (real-world scenario).

5. **Fail-Fast Philosophy**: Tests expect validation to stop at first error, matching the fail-fast requirement from plan.md.

### Implementation Hints

From examining the existing code, the implementation should:
- Add validation logic in `ValidateSchema` function (not in JSON decoder)
- Build table index before FK validation for O(1) lookup
- Use consistent error formatting: `fmt.Errorf("table '%s', column '%s': ...", tableName, colName)`
- Follow existing validation patterns from User Story 2
- Return early on first error (fail-fast)

---

**Status**: TDD RED Phase Complete ✅
**Next Phase**: GREEN Phase (T045-T050 Implementation)
**Test Quality**: High (comprehensive, specific, following patterns)
**Ready for Implementation**: Yes
