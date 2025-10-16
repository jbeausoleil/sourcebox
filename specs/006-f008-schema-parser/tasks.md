# Tasks: F008 - Schema Parser & Validator

**Input**: Design documents from `/specs/006-f008-schema-parser/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, quickstart.md

**Tests**: This feature uses **TDD (Test-Driven Development)** - tests are written BEFORE implementation for every validation rule. This is a constitutional requirement for core functionality.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story. However, due to the TDD nature and shared validation infrastructure, some dependencies exist across stories.

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions
- Go CLI project with package structure
- Core implementation: `pkg/schema/`
- Tests: `pkg/schema/*_test.go`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic Go package structure

- [X] T001 Create pkg/schema/ directory structure
- [X] T002 [P] Create pkg/schema/types.go stub file (Go structs for Schema, Table, Column, ForeignKey, Index, Relationship, ValidationRule)
- [X] T003 [P] Create pkg/schema/parser.go stub file (functions: ParseSchema, LoadSchema)
- [X] T004 [P] Create pkg/schema/parser_test.go file with package declaration and testify imports

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core structs and basic parser infrastructure that ALL user stories depend on

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T005 Implement all Go structs in pkg/schema/types.go with complete JSON tags (Schema, SchemaMetadata, Table, Column, ForeignKey, Index, Relationship, ValidationRule)
- [X] T006 Implement basic ParseSchema function in pkg/schema/parser.go (JSON decoder setup with DisallowUnknownFields)
- [X] T007 Implement LoadSchema convenience function in pkg/schema/parser.go (file opening and delegation to ParseSchema)
- [X] T008 Create ValidateSchema function stub in pkg/schema/parser.go (to be filled with validation logic)

**Checkpoint**: Foundation ready - user story implementation can now begin

---

## Phase 3: User Story 1 - Load Valid Schema (Priority: P1) 🎯 MVP

**Goal**: Enable developers to load valid schema JSON files and get usable data structures

**Independent Test**: Load a minimal valid schema with all required fields and verify all fields are correctly populated in the returned Schema object

### Tests for User Story 1

**NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [X] T009 [US1] Write TestParseValidMinimalSchema in pkg/schema/parser_test.go (tests parsing of minimal valid schema with one table, one column, primary key)
- [X] T010 [US1] Write TestParseValidFullSchema in pkg/schema/parser_test.go (tests parsing of complex schema with multiple tables, foreign keys, indexes, relationships)
- [X] T011 [US1] Write TestParseValidSchemaWithOptionalFields in pkg/schema/parser_test.go (tests schemas with metadata, validation_rules, indexes)

### Implementation for User Story 1

- [X] T012 [US1] Ensure ParseSchema correctly unmarshals JSON into Schema struct (verify JSON decoder works with all struct fields)
- [X] T013 [US1] Ensure ParseSchema correctly preserves all table definitions (verify Tables array is populated)
- [X] T014 [US1] Ensure ParseSchema correctly preserves column definitions and foreign keys (verify nested structures work)
- [X] T015 [US1] Ensure ParseSchema correctly handles optional fields (verify omitempty works, nil pointers for missing optional fields)
- [X] T016 [US1] Verify TestParseValidMinimalSchema passes
- [X] T017 [US1] Verify TestParseValidFullSchema passes
- [X] T018 [US1] Verify TestParseValidSchemaWithOptionalFields passes

**Checkpoint**: At this point, valid schemas should load successfully. All User Story 1 tests should pass.

---

## Phase 4: User Story 2 - Detect Missing Required Fields (Priority: P1)

**Goal**: Provide clear error messages when required fields are missing from schemas

**Independent Test**: Provide schemas with various missing required fields and verify specific error messages are returned

### Tests for User Story 2

- [X] T019 [P] [US2] Write TestParseMissingSchemaName in pkg/schema/parser_test.go
- [X] T020 [P] [US2] Write TestParseMissingTables in pkg/schema/parser_test.go
- [X] T021 [P] [US2] Write TestParseMissingDatabaseType in pkg/schema/parser_test.go
- [X] T022 [P] [US2] Write TestParseMissingGenerationOrder in pkg/schema/parser_test.go
- [X] T023 [P] [US2] Write TestParseTableMissingName in pkg/schema/parser_test.go
- [X] T024 [P] [US2] Write TestParseTableMissingColumns in pkg/schema/parser_test.go
- [X] T025 [P] [US2] Write TestParseTableMissingPrimaryKey in pkg/schema/parser_test.go
- [X] T026 [P] [US2] Write TestParseColumnMissingName in pkg/schema/parser_test.go
- [X] T027 [P] [US2] Write TestParseColumnMissingType in pkg/schema/parser_test.go
- [X] T028 [P] [US2] Write TestParseZeroRecordCount in pkg/schema/parser_test.go
- [X] T029 [P] [US2] Write TestParseNegativeRecordCount in pkg/schema/parser_test.go

### Implementation for User Story 2

- [X] T030 [US2] Implement ValidateSchema: check schema name is required (add to ValidateSchema function)
- [X] T031 [US2] Implement ValidateSchema: check at least one table exists (add to ValidateSchema function)
- [X] T032 [US2] Implement ValidateSchema: check database_type is required and valid (add to ValidateSchema function)
- [X] T033 [US2] Implement ValidateSchema: check generation_order is non-empty (add to ValidateSchema function)
- [X] T034 [US2] Implement ValidateTable function in pkg/schema/parser.go (checks table name, columns, record_count requirements)
- [X] T035 [US2] Implement ValidateColumn function in pkg/schema/parser.go (checks column name and type requirements)
- [X] T036 [US2] Implement primary key validation in ValidateTable (exactly one primary key per table)
- [X] T037 [US2] Implement record_count validation in ValidateTable (must be positive)
- [X] T038 [US2] Integrate table and column validation into ValidateSchema (call ValidateTable and ValidateColumn in loop)
- [X] T039 [US2] Verify all User Story 2 tests pass

**Checkpoint**: At this point, missing required fields should produce clear, actionable error messages

---

## Phase 5: User Story 3 - Validate Foreign Key Integrity (Priority: P2)

**Goal**: Ensure foreign keys reference existing tables and have valid integrity actions

**Independent Test**: Provide schemas with valid/invalid foreign keys and verify validation results

### Tests for User Story 3

- [ ] T040 [P] [US3] Write TestParseValidForeignKey in pkg/schema/parser_test.go (foreign key to existing table succeeds)
- [ ] T041 [P] [US3] Write TestParseForeignKeyNonExistentTable in pkg/schema/parser_test.go (error for non-existent table reference)
- [ ] T042 [P] [US3] Write TestParseForeignKeyInvalidOnDelete in pkg/schema/parser_test.go (error for invalid on_delete action)
- [ ] T043 [P] [US3] Write TestParseForeignKeyInvalidOnUpdate in pkg/schema/parser_test.go (error for invalid on_update action)
- [ ] T044 [P] [US3] Write TestParseForeignKeyValidActions in pkg/schema/parser_test.go (CASCADE, SET NULL, RESTRICT all valid)

### Implementation for User Story 3

- [ ] T045 [US3] Build table names map in ValidateSchema (map[string]bool of all table names for O(1) lookup)
- [ ] T046 [US3] Implement ValidateForeignKeys function in pkg/schema/parser.go (checks foreign key table existence)
- [ ] T047 [US3] Implement ValidateReferentialAction function in pkg/schema/parser.go (checks on_delete/on_update actions)
- [ ] T048 [US3] Integrate foreign key validation into ValidateSchema (call ValidateForeignKeys after table validation)
- [ ] T049 [US3] Add error context to foreign key errors (include table name, column name, array indices)
- [ ] T050 [US3] Verify all User Story 3 tests pass

**Checkpoint**: At this point, foreign key validation should catch invalid references and actions

---

## Phase 6: User Story 4 - Validate Data Types (Priority: P2)

**Goal**: Ensure column data types are from the supported set for MySQL/PostgreSQL

**Independent Test**: Provide schemas with valid/invalid data types and verify validation results

### Tests for User Story 4

- [ ] T051 [P] [US4] Write TestValidateDataTypeInt in pkg/schema/parser_test.go (int, bigint, smallint, tinyint valid)
- [ ] T052 [P] [US4] Write TestValidateDataTypeDecimal in pkg/schema/parser_test.go (decimal, float, double, decimal(10,2) valid)
- [ ] T053 [P] [US4] Write TestValidateDataTypeString in pkg/schema/parser_test.go (varchar(255), text, char(50) valid)
- [ ] T054 [P] [US4] Write TestValidateDataTypeDateTime in pkg/schema/parser_test.go (date, datetime, timestamp valid)
- [ ] T055 [P] [US4] Write TestValidateDataTypeBoolean in pkg/schema/parser_test.go (boolean, bit valid)
- [ ] T056 [P] [US4] Write TestValidateDataTypeJSON in pkg/schema/parser_test.go (json, jsonb valid)
- [ ] T057 [P] [US4] Write TestValidateDataTypeEnum in pkg/schema/parser_test.go (enum('a','b','c') valid)
- [ ] T058 [P] [US4] Write TestValidateDataTypeInvalid in pkg/schema/parser_test.go (unsupported type returns error)
- [ ] T059 [P] [US4] Write TestValidateDataTypeCaseInsensitive in pkg/schema/parser_test.go (INT, Int, int all valid)

### Implementation for User Story 4

- [ ] T060 [US4] Implement ValidateDataType function in pkg/schema/parser.go (case-insensitive prefix matching against supported types list)
- [ ] T061 [US4] Define validTypes slice in pkg/schema/parser.go (list of all supported data types from F007)
- [ ] T062 [US4] Integrate data type validation into ValidateColumn (call ValidateDataType for each column)
- [ ] T063 [US4] Verify all User Story 4 tests pass

**Checkpoint**: At this point, data type validation should reject unsupported types with clear error messages

---

## Phase 7: User Story 5 - Validate Generation Order (Priority: P2)

**Goal**: Ensure generation_order includes all tables exactly once with no duplicates or missing tables

**Independent Test**: Provide schemas with complete/incomplete/duplicate generation orders and verify validation results

### Tests for User Story 5

- [ ] T064 [P] [US5] Write TestValidateGenerationOrderComplete in pkg/schema/parser_test.go (all tables included succeeds)
- [ ] T065 [P] [US5] Write TestValidateGenerationOrderMissingTable in pkg/schema/parser_test.go (missing table produces error)
- [ ] T066 [P] [US5] Write TestValidateGenerationOrderDuplicate in pkg/schema/parser_test.go (duplicate table name produces error)
- [ ] T067 [P] [US5] Write TestValidateGenerationOrderNonExistentTable in pkg/schema/parser_test.go (table in order but not in schema produces error)
- [ ] T068 [P] [US5] Write TestValidateGenerationOrderEmpty in pkg/schema/parser_test.go (empty generation_order produces error)

### Implementation for User Story 5

- [ ] T069 [US5] Implement ValidateGenerationOrder function in pkg/schema/parser.go (checks all tables included, no duplicates, no missing)
- [ ] T070 [US5] Build generation order set in ValidateGenerationOrder (map[string]bool to detect duplicates)
- [ ] T071 [US5] Check all tables exist in generation_order (compare tableNames map with orderSet)
- [ ] T072 [US5] Check generation_order only references existing tables (validate each entry against tableNames map)
- [ ] T073 [US5] Integrate generation order validation into ValidateSchema (call ValidateGenerationOrder after table validation)
- [ ] T074 [US5] Verify all User Story 5 tests pass

**Checkpoint**: At this point, generation order validation should ensure correct table ordering for data generation

---

## Phase 8: User Story 6 - Detect Duplicate Names (Priority: P3)

**Goal**: Detect and reject duplicate table names or duplicate column names within tables

**Independent Test**: Provide schemas with duplicate table/column names and verify appropriate errors

### Tests for User Story 6

- [ ] T075 [P] [US6] Write TestParseDuplicateTableNames in pkg/schema/parser_test.go (two tables with same name produces error)
- [ ] T076 [P] [US6] Write TestParseDuplicateColumnNames in pkg/schema/parser_test.go (two columns in same table with same name produces error)

### Implementation for User Story 6

- [ ] T077 [US6] Implement duplicate table name detection in ValidateSchema (check tableNames map for duplicates)
- [ ] T078 [US6] Implement duplicate column name detection in ValidateTable (check for duplicate column names within each table)
- [ ] T079 [US6] Verify all User Story 6 tests pass

**Checkpoint**: At this point, duplicate names should be detected with clear error messages

---

## Phase 9: Edge Cases (Priority: P3)

**Goal**: Handle all edge cases gracefully with clear error messages

**Independent Test**: Provide edge case inputs and verify appropriate error handling

### Tests for Edge Cases

- [ ] T080 [P] [Edge] Write TestParseEmptyFile in pkg/schema/parser_test.go (empty file produces parse error)
- [ ] T081 [P] [Edge] Write TestParseMalformedJSON in pkg/schema/parser_test.go (invalid JSON syntax produces parse error)
- [ ] T082 [P] [Edge] Write TestParseUnknownFields in pkg/schema/parser_test.go (unknown fields in JSON produce error due to DisallowUnknownFields)
- [ ] T083 [P] [Edge] Write TestParseNullRequiredFields in pkg/schema/parser_test.go (null values for required fields produce validation error)
- [ ] T084 [P] [Edge] Write TestLoadSchemaNonExistentFile in pkg/schema/parser_test.go (file path doesn't exist produces clear error)

### Implementation for Edge Cases

- [ ] T085 [Edge] Ensure ParseSchema handles empty input gracefully (verify JSON decoder error handling)
- [ ] T086 [Edge] Ensure ParseSchema handles malformed JSON gracefully (verify JSON decode errors are wrapped with context)
- [ ] T087 [Edge] Verify DisallowUnknownFields is enabled (ensure strict parsing catches typos)
- [ ] T088 [Edge] Ensure LoadSchema handles file open errors gracefully (verify os.Open error wrapping)
- [ ] T089 [Edge] Verify all edge case tests pass

**Checkpoint**: All edge cases should be handled with clear, actionable error messages

---

## Phase 10: Polish & Cross-Cutting Concerns

**Purpose**: Final improvements and verification

- [ ] T090 [P] Run go fmt ./pkg/schema/... (format all schema package files)
- [ ] T091 [P] Run go vet ./pkg/schema/... (verify no lint issues)
- [ ] T092 Verify test coverage with go test -coverprofile=coverage.out ./pkg/schema/... (target: 100% for parser.go)
- [ ] T093 Run go test ./pkg/schema/... -v (ensure all tests pass)
- [ ] T094 Test parser with F007 example schema file (manual verification that example schema from F007 loads successfully)
- [ ] T095 [P] Add package documentation comment to pkg/schema/types.go
- [ ] T096 [P] Add package documentation comment to pkg/schema/parser.go
- [ ] T097 Run quickstart.md verification steps (follow quickstart guide to verify all steps work)
- [ ] T098 Measure parse performance with benchmark (verify <100ms for typical schemas)

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Story 1 (Phase 3)**: Depends on Foundational phase - validates basic parsing works
- **User Story 2 (Phase 4)**: Depends on Foundational phase - can start after US1 or in parallel
- **User Story 3 (Phase 5)**: Depends on Foundational + table names map - can start after US2
- **User Story 4 (Phase 6)**: Depends on Foundational + ValidateColumn - can start after US2
- **User Story 5 (Phase 7)**: Depends on Foundational + table names map - can start after US3
- **User Story 6 (Phase 8)**: Depends on Foundational + basic validation - can start after US2
- **Edge Cases (Phase 9)**: Depends on basic parser implementation - can start after US1
- **Polish (Phase 10)**: Depends on all user stories complete

### TDD Workflow (CRITICAL)

For EVERY validation rule:
1. **RED**: Write failing test first
2. **GREEN**: Implement minimal validation to pass test
3. **REFACTOR**: Clean up code while keeping tests green
4. **VERIFY**: Run all tests to ensure no regression

### Within Each User Story

- Tests MUST be written FIRST and MUST FAIL before implementation
- Tests marked [P] can be written in parallel
- Implementation tasks must be sequential (build on previous validation logic)
- Verify all tests pass before moving to next user story

### Parallel Opportunities

- Phase 1 (Setup): T002, T003, T004 can run in parallel (different files)
- Phase 2 (Foundational): Tasks are sequential (types → parser → validation stub)
- Phase 3 (US1): T009, T010, T011 tests can be written in parallel
- Phase 4 (US2): T019-T029 tests can be written in parallel
- Phase 5 (US3): T040-T044 tests can be written in parallel
- Phase 6 (US4): T051-T059 tests can be written in parallel
- Phase 7 (US5): T064-T068 tests can be written in parallel
- Phase 8 (US6): T075-T076 tests can be written in parallel
- Phase 9 (Edge): T080-T084 tests can be written in parallel
- Phase 10 (Polish): T090, T091, T095, T096 can run in parallel

---

## Parallel Example: User Story 2

```bash
# Write all tests for User Story 2 in parallel:
Task: "Write TestParseMissingSchemaName in pkg/schema/parser_test.go"
Task: "Write TestParseMissingTables in pkg/schema/parser_test.go"
Task: "Write TestParseMissingDatabaseType in pkg/schema/parser_test.go"
# ... all other US2 test tasks

# Then implement validation sequentially:
Task: "Implement ValidateSchema: check schema name is required"
Task: "Implement ValidateSchema: check at least one table exists"
# ... etc.
```

---

## Implementation Strategy

### MVP First (User Stories 1 + 2 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL)
3. Complete Phase 3: User Story 1 (load valid schemas)
4. Complete Phase 4: User Story 2 (detect missing fields)
5. **STOP and VALIDATE**: Parser can load valid schemas and reject invalid ones
6. This is the minimum viable parser for data generation

### Incremental Delivery

1. **Foundation** (Phases 1-2): Basic structs and parser infrastructure
2. **MVP** (Phases 3-4): Load valid schemas + detect missing fields → Can start using parser
3. **Enhanced Validation** (Phases 5-7): Foreign keys + data types + generation order → Catch more schema errors
4. **Complete** (Phases 8-9): Duplicate detection + edge cases → Production-ready
5. **Polish** (Phase 10): Documentation + performance verification

### Test Coverage Strategy

- Target: **100% coverage for pkg/schema/parser.go** (critical path)
- Every validation rule MUST have corresponding test(s)
- Every error path MUST be tested
- Use table-driven tests for similar test cases (data type validation, foreign key validation)
- Use `go test -coverprofile=coverage.out` and `go tool cover -html=coverage.out` to identify untested paths

---

## Notes

- [P] tasks = different files or independent test cases, no dependencies
- [Story] label maps task to specific user story for traceability
- **TDD is non-negotiable**: Constitutional requirement for core functionality (Development Practice 1)
- Error messages MUST include context (table name, column index, specific field)
- Use testify assertions: `require.Error(t, err)` and `assert.Contains(t, err.Error(), "expected text")`
- Validation is fail-fast: Return first error encountered (don't accumulate multiple errors)
- Run `go test -v` frequently during development to ensure tests pass
- Commit after each logical group of tests + implementation
- Stop at any checkpoint to validate independently before proceeding
- Avoid: vague error messages, skipping tests, implementing before testing
