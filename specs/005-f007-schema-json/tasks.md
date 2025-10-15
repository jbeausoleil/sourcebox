# Tasks: Schema JSON Format Specification (F007)

**Feature Branch**: `005-f007-schema-json`
**Input**: Design documents from `/specs/005-f007-schema-json/`
**Prerequisites**: plan.md ✅, spec.md ✅, research.md ✅, quickstart.md ✅

**Important**: This is a **design-only feature** (specification and documentation, no code implementation). No tests are required or included.

**Organization**: Tasks are grouped by user story to enable independent implementation and validation of each documentation section.

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3, US4, US5)
- File paths are absolute from repository root

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Create directory structure for schema specifications

- [X] T001 Create `schemas/` directory at repository root if it doesn't exist
- [X] T002 Verify CLAUDE.md exists and is ready for schema format updates

**Checkpoint**: Directory structure ready for specification files

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: None required - this is documentation work with no foundational code dependencies

**Note**: No foundational phase needed. Documentation can proceed directly to user stories.

---

## Phase 3: User Story 1 - Understanding Schema Format (Priority: P1) 🎯 MVP CORE

**Goal**: Create comprehensive schema format specification document that enables developers to understand and create schemas

**Independent Test**: Developer with no prior SourceBox knowledge can read `schemas/schema-spec.md` and identify all required fields for creating a valid schema

**Why P1**: Without schema format documentation, developers cannot contribute new schemas or modify existing ones. This is foundational to the entire schema ecosystem.

### Implementation for User Story 1

- [X] T003 [US1] Create `schemas/schema-spec.md` with overview section explaining JSON format rationale
- [X] T004 [US1] Document schema structure section (schema_version, name, description, author, version, database_type, metadata, tables, relationships, generation_order, validation_rules)
- [X] T005 [US1] Document table definition format (name, description, record_count, columns, indexes)
- [X] T006 [US1] Document column definition format (name, type, nullable, primary_key, unique, default, generator, generator_params, foreign_key)
- [X] T007 [US1] Add copyable JSON examples for each schema element in schema-spec.md
- [X] T008 [US1] Document dual versioning strategy (schema_version vs version fields)
- [X] T009 [US1] Document generation_order purpose and format (parent tables first)

**Checkpoint**: At this point, User Story 1 (Understanding Schema Format) should be complete - developers can understand the basic structure

---

## Phase 4: User Story 4 - Creating Example Schemas (Priority: P1) 🎯 MVP CORE

**Goal**: Create a complete, working example schema demonstrating all major features as a reference for schema authors

**Independent Test**: Examine `schemas/example-schema.json` and verify it contains examples of all documented features (multiple tables, various column types, generators with parameters, foreign keys, indexes, relationships)

**Why P1**: A comprehensive example schema serves as both documentation and validation of the format. It provides a concrete reference that developers can copy, modify, and learn from.

**Note**: This depends on US1 being substantially complete so we know what to demonstrate

### Implementation for User Story 4

- [X] T010 [US4] Create `schemas/example-schema.json` with fintech schema structure (borrowers, loans, payments tables)
- [X] T011 [US4] Define borrowers table (250 records) with personal data columns (first_name, last_name, email, phone, credit_score)
- [X] T012 [US4] Define loans table (1,000 records) with custom generators (loan_amount lognormal, interest_rate ranges, loan_status weighted)
- [X] T013 [US4] Define payments table (3,700 records) with foreign key to loans and payment amounts
- [X] T014 [US4] Add inline foreign_key definitions to loans.borrower_id and payments.loan_id
- [X] T015 [US4] Add explicit relationships section documenting all foreign key relationships
- [X] T016 [US4] Add indexes (email unique on borrowers, credit_score indexed, loan_id indexed on payments)
- [X] T017 [US4] Add generation_order array (borrowers, loans, payments)
- [X] T018 [US4] Add metadata section (industry: fintech, tags, total_records: 4950)
- [X] T019 [US4] Validate JSON syntax using `cat schemas/example-schema.json | jq '.'`
- [X] T020 [US4] Verify Tier 1 complexity compliance (<5,000 total records per constitution line 176-197)

**Checkpoint**: At this point, User Stories 1 and 4 are complete - developers have both documentation and a working example

---

## Phase 5: User Story 2 - Validating Schema Correctness (Priority: P2)

**Goal**: Document validation rules that determine whether a schema is valid so schema authors can create schemas that work correctly without runtime errors

**Independent Test**: Review the validation rules section and attempt to create both valid and invalid schemas according to the rules

**Why P2**: Clear validation rules prevent developers from creating invalid schemas that fail at runtime. Enables self-service schema creation.

**Note**: Can proceed in parallel with or after US1/US4, as validation rules formalize the structure already documented

### Implementation for User Story 2

- [X] T021 [P] [US2] Document schema-level validation rules in schema-spec.md (unique name, at least one table, valid database_type, generation_order completeness)
- [X] T022 [P] [US2] Document table-level validation rules in schema-spec.md (unique name, at least one column, exactly one primary key, record_count > 0)
- [X] T023 [P] [US2] Document column-level validation rules in schema-spec.md (valid types, valid generators, params match requirements, foreign keys reference existing tables)
- [X] T024 [P] [US2] Document relationship-level validation rules in schema-spec.md (foreign keys reference primary keys, valid integrity actions CASCADE/SET NULL/RESTRICT)
- [X] T025 [US2] Document generation_order validation rules (parent tables first, no circular dependencies)
- [X] T026 [US2] Add edge case documentation (circular dependencies, missing foreign key targets, invalid generator parameters)
- [X] T027 [US2] Add validation error message guidance for F008 implementers

**Checkpoint**: Validation rules are complete - schema authors can self-validate their work

---

## Phase 6: User Story 3 - Using Built-in Generators (Priority: P2)

**Goal**: Document all built-in generators and their configuration so schema authors can generate appropriate data without writing custom code

**Independent Test**: Review the generator documentation and create schemas that use various built-in generators with different parameter configurations

**Why P2**: Built-in generators enable rapid schema creation for common data types. Understanding their capabilities allows developers to leverage existing functionality before creating custom generators.

**Note**: Can proceed in parallel with US2, as generators are independent documentation sections

### Implementation for User Story 3

- [ ] T028 [P] [US3] Document built-in personal data generators in schema-spec.md (first_name, last_name, full_name, email, phone, address, ssn, date_of_birth)
- [ ] T029 [P] [US3] Document built-in company data generators in schema-spec.md (company_name, job_title, company_email, domain)
- [ ] T030 [P] [US3] Document built-in date/time generators in schema-spec.md (timestamp_past, timestamp_future, date_between with parameters)
- [ ] T031 [P] [US3] Document built-in numeric generators in schema-spec.md (int_range, float_range, decimal_range)
- [ ] T032 [US3] Document custom generator pattern (how schemas define custom generators like credit_score, loan_amount)
- [ ] T033 [US3] Document generator_params structure and format
- [ ] T034 [US3] Document distribution types: uniform (min, max), normal (mean, std_dev, min, max), lognormal (median, min, max), weighted (values array), ranges (ranges array)
- [ ] T035 [US3] Add copyable examples for each distribution type with realistic use cases
- [ ] T036 [US3] Document parameter types and validation (required vs optional parameters)

**Checkpoint**: Generator documentation is complete - developers understand all available generators and how to configure them

---

## Phase 7: User Story 5 - Versioning Schemas Over Time (Priority: P3)

**Goal**: Document versioning strategy so schema maintainers can evolve schemas without breaking existing users or implementations

**Independent Test**: Review the versioning section and understand when to increment major, minor, or patch versions

**Why P3**: Schema versioning enables long-term maintenance and evolution. Less critical than initial creation but essential for sustainable schema development.

**Note**: Can be done last as it documents schema evolution rather than initial creation

### Implementation for User Story 5

- [ ] T037 [US5] Document semantic versioning (semver) strategy in schema-spec.md
- [ ] T038 [US5] Define major version rules (breaking changes: table removal, type changes, foreign key changes)
- [ ] T039 [US5] Define minor version rules (backward-compatible: new tables, new columns, new generators)
- [ ] T040 [US5] Define patch version rules (bug fixes: distribution improvements, parameter tuning, docs updates)
- [ ] T041 [US5] Add versioning examples (when to bump 1.0.0 → 1.1.0 vs 1.0.0 → 2.0.0)
- [ ] T042 [US5] Document backward compatibility guidelines for minor versions

**Checkpoint**: Versioning strategy is documented - schema maintainers understand how to evolve schemas safely

---

## Phase 8: Supported Data Types Documentation (Cross-Cutting)

**Goal**: Document all supported SQL data types for MySQL and PostgreSQL

**Note**: Supports multiple user stories (US1, US2, US3) - developers need to know valid types

### Implementation

- [ ] T043 [P] [Cross] Document integer types in schema-spec.md (int, bigint, smallint, tinyint)
- [ ] T044 [P] [Cross] Document decimal types in schema-spec.md (decimal(p,s), float, double)
- [ ] T045 [P] [Cross] Document string types in schema-spec.md (varchar(n), text, char(n))
- [ ] T046 [P] [Cross] Document date/time types in schema-spec.md (date, datetime, timestamp)
- [ ] T047 [P] [Cross] Document boolean type in schema-spec.md (boolean for PostgreSQL, bit for MySQL)
- [ ] T048 [P] [Cross] Document JSON types in schema-spec.md (json for both, jsonb for PostgreSQL, document fallback behavior)
- [ ] T049 [P] [Cross] Document enum types in schema-spec.md (enum format, MySQL vs PostgreSQL differences)
- [ ] T050 [Cross] Add platform-specific guidance (type compatibility matrix for MySQL and PostgreSQL)

**Checkpoint**: Data types are fully documented - developers know what types are supported and how they map across databases

---

## Phase 9: Foreign Key Relationships Documentation (Cross-Cutting)

**Goal**: Document foreign key format for both inline and explicit representations

**Note**: Supports US1 and US2 - developers need clear foreign key syntax

### Implementation

- [ ] T051 [US1+US2] Document inline foreign_key object structure in schema-spec.md (table, column, on_delete, on_update)
- [ ] T052 [US1+US2] Document explicit relationships section format in schema-spec.md (from_table, from_column, to_table, to_column, relationship_type, description)
- [ ] T053 [US1+US2] Document rationale for dual representation (inline for code, explicit for documentation)
- [ ] T054 [US1+US2] Document relationship types (one_to_one, one_to_many, many_to_one, many_to_many)
- [ ] T055 [US1+US2] Document referential integrity actions (CASCADE, SET NULL, RESTRICT) with examples
- [ ] T056 [US1+US2] Add foreign key examples showing both inline and explicit representations

**Checkpoint**: Foreign key documentation is complete - developers understand how to define relationships

---

## Phase 10: Polish & Cross-Cutting Concerns

**Purpose**: Final improvements, CLAUDE.md update, and validation

- [ ] T057 [P] Update CLAUDE.md with schema format guidelines (overview, generator types, distribution types, foreign keys, versioning, validation rules, Tier 1 targets, creating new schemas, supported data types)
- [ ] T058 [P] Add table of contents to schema-spec.md for easy navigation
- [ ] T059 [P] Proofread schema-spec.md for clarity, consistency, and completeness
- [ ] T060 [P] Verify all [TODO] and [NEEDS CLARIFICATION] markers are resolved
- [ ] T061 Run quickstart.md validation steps (all 10 verification checks)
- [ ] T062 Verify constitutional compliance (Tier 1 complexity, verticalized data, boring tech, implementable by F008)
- [ ] T063 Final review of example schema for feature completeness and Tier 1 compliance

**Checkpoint**: F007 complete and ready for F008 implementation

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Skipped for this feature (no code dependencies)
- **User Story 1 (Phase 3)**: Depends on Setup - START HERE 🎯
- **User Story 4 (Phase 4)**: Depends on US1 being substantially complete (need to know what to demonstrate)
- **User Story 2 (Phase 5)**: Can start after US1 is complete, parallel with US3
- **User Story 3 (Phase 6)**: Can start after US1 is complete, parallel with US2
- **User Story 5 (Phase 7)**: Can start anytime after US1, preferably last
- **Data Types (Phase 8)**: Can proceed in parallel with other phases once US1 structure is defined
- **Foreign Keys (Phase 9)**: Can proceed in parallel with other phases once US1 structure is defined
- **Polish (Phase 10)**: Depends on all other phases being complete

### User Story Dependencies

- **User Story 1 (P1)**: No dependencies - foundation for all others 🎯 START HERE
- **User Story 4 (P1)**: Depends on US1 (need schema structure to create example) 🎯 SECOND
- **User Story 2 (P2)**: Depends on US1, can run parallel with US3
- **User Story 3 (P2)**: Depends on US1, can run parallel with US2
- **User Story 5 (P3)**: Minimal dependencies, can be done last

### Critical Path

1. **Phase 1: Setup** (T001-T002) → 2 tasks
2. **Phase 3: User Story 1** (T003-T009) → 7 tasks 🎯 CRITICAL
3. **Phase 4: User Story 4** (T010-T020) → 11 tasks 🎯 CRITICAL
4. **Phases 5-9: Documentation Completion** (T021-T056) → Can proceed in parallel
5. **Phase 10: Polish** (T057-T063) → Final review

### Parallel Opportunities

Within each phase, tasks marked [P] can run in parallel:

- **Phase 5 (US2)**: T021-T024 can all run in parallel (different validation rule categories)
- **Phase 6 (US3)**: T028-T031 can all run in parallel (different generator categories)
- **Phase 8 (Data Types)**: T043-T049 can all run in parallel (different type categories)
- **Phase 10 (Polish)**: T057-T060 can all run in parallel (different files)

User Stories can also proceed in parallel after US1 is complete:
- After US1 done: US2, US3, US5 can proceed in parallel
- After US4 done: Example can be referenced by other documentation

---

## Parallel Example: Documentation Sections

```bash
# After US1 core structure is complete, launch these in parallel:

# User Story 2: Validation Rules
Task: "Document schema-level validation rules"
Task: "Document table-level validation rules"
Task: "Document column-level validation rules"
Task: "Document relationship-level validation rules"

# User Story 3: Generator Documentation
Task: "Document built-in personal data generators"
Task: "Document built-in company data generators"
Task: "Document built-in date/time generators"
Task: "Document built-in numeric generators"

# Data Types Documentation
Task: "Document integer types"
Task: "Document decimal types"
Task: "Document string types"
Task: "Document date/time types"
```

---

## Implementation Strategy

### MVP First (User Stories 1 + 4 Only)

1. Complete Phase 1: Setup (2 tasks)
2. Complete Phase 3: User Story 1 - Understanding Schema Format (7 tasks) 🎯
3. Complete Phase 4: User Story 4 - Creating Example Schemas (11 tasks) 🎯
4. **STOP and VALIDATE**: Review schema-spec.md + example-schema.json
5. This delivers immediate value: developers can understand and copy the example

**Total MVP tasks**: 20 tasks

### Incremental Delivery

1. MVP (US1 + US4) → Schema format documented with example ✅
2. Add US2 (Validation Rules) → Schema authors can self-validate ✅
3. Add US3 (Built-in Generators) → Generator reference complete ✅
4. Add Data Types + Foreign Keys → Full reference documentation ✅
5. Add US5 (Versioning) → Long-term maintenance guide ✅
6. Polish → Production-ready specification ✅

### Validation Points

- **After US1**: Can developers understand the schema structure?
- **After US4**: Does example demonstrate all documented features?
- **After US2**: Can developers identify validation errors?
- **After US3**: Can developers configure generators correctly?
- **After Phase 8-9**: Are all SQL types and relationships documented?
- **After Phase 10**: Does specification pass all quickstart.md checks?

---

## Task Summary

- **Total Tasks**: 63
- **MVP Tasks (US1 + US4)**: 20 tasks
- **Setup**: 2 tasks
- **User Story 1 (P1)**: 7 tasks
- **User Story 4 (P1)**: 11 tasks
- **User Story 2 (P2)**: 7 tasks
- **User Story 3 (P2)**: 9 tasks
- **User Story 5 (P3)**: 6 tasks
- **Data Types (Cross-Cutting)**: 8 tasks
- **Foreign Keys (Cross-Cutting)**: 6 tasks
- **Polish**: 7 tasks

**Parallel Opportunities**: 23 tasks marked [P] can run in parallel

**Estimated Effort**:
- MVP (US1 + US4): 1-2 days (comprehensive documentation + example schema)
- Full Feature (all user stories): 2-3 days (per plan.md timeline)
- Each user story after MVP: ~0.5 day

---

## Notes

- [P] tasks = different sections/files, can write in parallel
- [Story] label maps task to specific user story for traceability
- [Cross] label indicates task supports multiple user stories
- This is documentation work only - no code implementation
- No tests required (design feature, not code feature)
- Tests in this context = "manual validation" via quickstart.md
- Each user story section can be validated independently
- Example schema (US4) serves as validation of specification (US1)
- Commit after each major section or logical group
- Final validation: Run all 10 quickstart.md verification steps

---

## Success Criteria

F007 is complete when:

✅ All 63 tasks are complete
✅ All user stories deliver their independent test criteria
✅ schemas/schema-spec.md exists with all required sections
✅ schemas/example-schema.json exists and demonstrates all features
✅ Example schema is valid JSON and Tier 1 compliant (<5,000 records)
✅ All validation rules are documented and implementable
✅ All generator types are documented with examples
✅ All data types are documented (MySQL/PostgreSQL common subset)
✅ Foreign key format is clearly documented (inline + explicit)
✅ Versioning strategy is documented (semver rules)
✅ CLAUDE.md is updated with schema format guidelines
✅ All 10 quickstart.md verification checks pass
✅ No [TODO] or [NEEDS CLARIFICATION] markers remain
✅ Specification is implementable by F008 without ambiguity
✅ Constitutional compliance verified (Tier 1, verticalized, boring tech)
