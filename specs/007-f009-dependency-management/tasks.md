# Tasks: Dependency Management Setup (F009)

**Input**: Design documents from `/specs/007-f009-dependency-management/`
**Prerequisites**: plan.md, spec.md, research.md, quickstart.md

**Tests**: No automated tests required for this infrastructure feature. Validation is through manual verification commands.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3, US4)
- Include exact file paths in descriptions

## Path Conventions
- Repository root: `/Users/jbeausoleil/Projects/03_projects/personal/sourcebox/`
- Dependencies defined in: `go.mod`, `go.sum`
- Documentation: `README.md`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: No additional setup required - F003 (go.mod initialization) is complete

**Current State**:
- ✅ go.mod exists with module path `github.com/jbeausoleil/sourcebox`
- ✅ Go 1.21 specified
- ✅ Cobra v1.10.1 already installed (newer than F009 target of v1.8.0)
- ⚠️ Missing 5 dependencies: gofakeit, progressbar, mysql driver, pq driver, color

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core dependency setup that MUST be complete before ANY user story validation

**⚠️ CRITICAL**: No user story work can begin until these dependencies are installed

- [X] T001 Add gofakeit dependency: `go get github.com/brianvoe/gofakeit/v6@v6.27.0`
- [X] T002 Add progressbar dependency: `go get github.com/schollz/progressbar/v3@v3.14.1`
- [X] T003 Add MySQL driver dependency: `go get github.com/go-sql-driver/mysql@v1.7.1`
- [X] T004 Add PostgreSQL driver dependency: `go get github.com/lib/pq@v1.10.9`
- [X] T005 Add color dependency: `go get github.com/fatih/color@v1.16.0`
- [X] T006 Run `go mod tidy` to organize dependencies and generate go.sum checksums

**Checkpoint**: Foundation ready - all 6 core dependencies installed, go.sum generated

---

## Phase 3: User Story 1 - Initial Project Setup (Priority: P1) 🎯 MVP

**Goal**: Developers can clone the repository and build/test without manual dependency installation

**Independent Test**:
```bash
# Simulate fresh clone (optional: test in new directory)
go mod download  # Download dependencies
go build ./cmd/sourcebox  # Build should succeed
go test ./...  # Tests should pass
```

### Implementation for User Story 1

- [X] T007 [US1] Verify dependency integrity with `go mod verify` - expect "all modules verified"
- [X] T008 [US1] Test clean build with timing: `time go build ./cmd/sourcebox` - expect < 30 seconds
- [X] T009 [US1] Run test suite: `go test ./...` - verify no import errors
- [X] T010 [US1] Verify clean state: `go mod tidy && git diff go.mod go.sum` - expect no changes

**Acceptance Validation**:
- ✅ Build succeeds without manual steps
- ✅ Tests execute successfully
- ✅ Build time < 30 seconds
- ✅ No uncommitted go.mod/go.sum changes

**Checkpoint**: User Story 1 complete - automatic dependency download works

---

## Phase 4: User Story 2 - Dependency Verification (Priority: P2)

**Goal**: Developers can verify dependency integrity and ensure exact versions are used

**Independent Test**:
```bash
go mod verify  # Should pass with all checksums matching
go mod tidy  # Should produce no changes
```

### Implementation for User Story 2

- [X] T011 [US2] Test checksum verification: Run `go mod verify` and confirm "all modules verified" output
- [X] T012 [US2] Test corruption detection: Simulate tampered dependency (optional manual test)
- [X] T013 [US2] Verify go.sum is committed: `git status go.sum` - should show as tracked file
- [X] T014 [US2] Test clean state verification: `go mod tidy && git status` - should show no changes

**Acceptance Validation**:
- ✅ `go mod verify` passes
- ✅ go.sum file committed to version control
- ✅ `go mod tidy` produces no changes
- ✅ Checksums match for all dependencies

**Checkpoint**: User Story 2 complete - dependency verification works

---

## Phase 5: User Story 3 - Dependency Documentation (Priority: P2)

**Goal**: Developers understand dependency choices, licensing, and update procedures

**Independent Test**: Read README.md and verify all required information is present and accurate

### Implementation for User Story 3

- [X] T015 [P] [US3] Add "Dependencies" section to README.md after "Installation", before "Usage"
- [X] T016 [P] [US3] Document dependencies by category in README.md:
  - CLI & UX: Cobra (v1.10.1), progressbar (v3.14.1), color (v1.16.0)
  - Data Generation: gofakeit (v6.27.0)
  - Database Drivers: go-sql-driver/mysql (v1.7.1), lib/pq (v1.10.9)
- [X] T017 [US3] Add license compatibility statement to README.md: "All dependencies are MIT or similarly permissive licenses compatible with our MIT license"
- [X] T018 [US3] Add each dependency with: name (GitHub link), version (backticks), one-line purpose
- [X] T019 [US3] Verify README.md versions match go.mod exactly

**Acceptance Validation**:
- ✅ Each dependency listed with purpose and version
- ✅ Dependencies grouped by category
- ✅ All licenses noted as MIT-compatible
- ✅ Update instructions available
- ✅ README versions match go.mod

**Checkpoint**: User Story 3 complete - dependency documentation is complete

---

## Phase 6: User Story 4 - Clean Dependency State (Priority: P3)

**Goal**: Project maintains clean dependency state with no unused dependencies

**Independent Test**:
```bash
go mod tidy  # Should produce no changes
go list -m all  # Verify all listed dependencies are used
```

### Implementation for User Story 4

- [X] T020 [US4] Verify no unused direct dependencies: Run `go mod tidy` and check for removals
- [X] T021 [US4] Verify indirect dependencies are properly marked: Check go.mod for `// indirect` comments
- [X] T022 [US4] Test dependency cleanup: Temporarily add unused dependency, run `go mod tidy`, verify removal
- [X] T023 [US4] Document cleanup process: Add to CONTRIBUTING.md dependency update section (if needed)

**Acceptance Validation**:
- ✅ `go mod tidy` produces no changes on clean repository
- ✅ All direct dependencies are imported in code
- ✅ Indirect dependencies properly marked
- ✅ Unused dependencies removed automatically

**Checkpoint**: User Story 4 complete - clean dependency state maintained

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Final validation and documentation updates

- [X] T024 [P] Run full quickstart.md validation checklist
- [X] T025 [P] Verify all Success Criteria (SC-001 through SC-008) from spec.md
- [X] T026 Commit all changes: `git add go.mod go.sum README.md`
- [X] T027 Create commit with clear message documenting all 6 dependencies and their versions
- [X] T028 Update CLAUDE.md with dependency management context (if not already done)
- [X] T029 Verify constitutional compliance: <10 dependencies, MIT-compatible licenses, <30s builds

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: SKIPPED - F003 already complete
- **Foundational (Phase 2)**: T001-T006 - Install 5 missing dependencies, run go mod tidy - BLOCKS all user stories
- **User Stories (Phase 3-6)**: All depend on Foundational phase completion
  - User Story 1 (Phase 3): Tests automatic dependency download
  - User Story 2 (Phase 4): Tests dependency verification
  - User Story 3 (Phase 5): Documents dependencies
  - User Story 4 (Phase 6): Validates clean state
- **Polish (Phase 7)**: Depends on all user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start immediately after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - Independent of US1 (can run in parallel)
- **User Story 3 (P2)**: Can start after Foundational (Phase 2) - Independent of US1/US2 (can run in parallel)
- **User Story 4 (P3)**: Should run after US1/US2/US3 to verify final clean state

### Within Each User Story

- Foundational tasks (T001-T006) must complete sequentially (each `go get` modifies go.mod)
- T006 (`go mod tidy`) must run after all `go get` commands
- User Story tasks can run in parallel after Foundational phase
- Documentation tasks (US3) can run in parallel with verification tasks (US1, US2, US4)

### Parallel Opportunities

- After Foundational (T001-T006): All user stories can be validated in parallel
- Within US3: All documentation tasks marked [P] can be written simultaneously
- Verification tasks (US1, US2, US4) can run concurrently if desired

---

## Parallel Example: After Foundational Phase

```bash
# After T001-T006 complete, these can run in parallel:

# Terminal 1: User Story 1 - Build verification
Task: "T007-T010: Verify build, test, and clean state"

# Terminal 2: User Story 2 - Checksum verification
Task: "T011-T014: Verify checksums and clean state"

# Terminal 3: User Story 3 - Documentation
Task: "T015-T019: Document dependencies in README.md"

# Then run User Story 4 after others complete
Task: "T020-T023: Verify clean dependency state"
```

---

## Implementation Strategy

### MVP First (User Stories 1 & 2 Only)

1. Complete Phase 2: Foundational (T001-T006) - Install all dependencies
2. Complete Phase 3: User Story 1 (T007-T010) - Verify builds work
3. Complete Phase 4: User Story 2 (T011-T014) - Verify checksums work
4. **STOP and VALIDATE**: Dependencies installed, builds work, verification passes
5. Dependencies are now usable by downstream features (F013, F021, F022)

### Full Delivery (All User Stories)

1. Complete Foundational → Dependencies installed
2. Complete User Story 1 → Builds work automatically
3. Complete User Story 2 → Verification works
4. Complete User Story 3 → Documentation complete
5. Complete User Story 4 → Clean state validated
6. Complete Polish → All validated, committed, documented

### Sequential Strategy (Recommended for Single Developer)

1. **Phase 2 (Foundational)**: T001-T006 (15 minutes)
2. **Phase 3 (US1)**: T007-T010 (10 minutes)
3. **Phase 4 (US2)**: T011-T014 (10 minutes)
4. **Phase 5 (US3)**: T015-T019 (20 minutes) - Documentation writing
5. **Phase 6 (US4)**: T020-T023 (10 minutes)
6. **Phase 7 (Polish)**: T024-T029 (15 minutes)

**Total Time**: ~90 minutes for complete implementation with validation

---

## Notes

- [P] tasks = different files/independent operations, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story validates a different aspect of dependency management
- Verification is through manual command execution, not automated tests
- This feature is infrastructure/configuration - no data entities or code changes
- All Success Criteria (SC-001 through SC-008) should be verified during polish phase
- Avoid: skipping go mod tidy, forgetting to commit go.sum, documenting wrong versions

---

## Verification Checklist (from quickstart.md)

### go.mod Verification
- [ ] go.mod exists at repo root
- [ ] Module path is `github.com/jbeausoleil/sourcebox`
- [ ] Go version is 1.21
- [ ] Require block has 6 direct dependencies (Cobra, gofakeit, progressbar, mysql driver, pq driver, color)
- [ ] Require block has indirect dependencies marked with `// indirect`
- [ ] Dependencies are alphabetically sorted
- [ ] All versions are exact (e.g., v1.8.0, not ranges)

### go.sum Verification
- [ ] go.sum exists at repo root
- [ ] go.sum has entries for all direct dependencies
- [ ] go.sum has entries for all indirect dependencies
- [ ] `go mod verify` passes with "all modules verified"
- [ ] go.sum is committed to git

### Build Verification
- [ ] `go build ./cmd/sourcebox` succeeds without errors
- [ ] Build time is < 30 seconds
- [ ] Binary is created (`sourcebox`)
- [ ] Binary can be executed (`./sourcebox --help` works)

### Test Verification
- [ ] `go test ./...` runs without errors
- [ ] No dependency-related import errors

### Dependency State Verification
- [ ] `go mod tidy` produces no changes to go.mod or go.sum
- [ ] `git diff go.mod go.sum` shows no changes after tidy
- [ ] No unused dependencies

### Documentation Verification
- [ ] README.md has "Dependencies" section (after "Installation", before "Usage")
- [ ] Each dependency listed with name, version, purpose
- [ ] Dependencies grouped by category (CLI & UX, Data Generation, Database Drivers)
- [ ] License compatibility statement included
- [ ] Dependency versions in README match go.mod exactly

---

## Success Criteria Mapping

| Task | Success Criteria |
|------|------------------|
| T001-T006 | SC-001, SC-007 (automatic dependency download, <10 deps) |
| T007 | SC-002 (dependency verification) |
| T008 | SC-003 (build time <30s) |
| T009-T010 | SC-005 (go mod tidy produces no changes) |
| T015-T019 | SC-006 (complete documentation) |
| T015-T019 | SC-004, SC-008 (license documentation) |
| T024-T029 | All SC validation |

---

## Edge Cases Handled

- **Missing go.sum**: T007 (`go mod verify`) will detect and provide clear error
- **Security vulnerabilities**: Out of scope for MVP - addressed in quarterly review
- **Corporate proxy**: Works automatically via standard Go environment variables
- **Conflicting transitive dependencies**: Go MVS handles automatically
- **Dependency unavailability**: Go module proxy provides cached versions

---

## Next Steps After F009

Once F009 is complete, these features can proceed:

1. **F013: Data Generation Engine** - Uses gofakeit for realistic fake data
2. **F021: Seed command implementation** - Uses Cobra, progressbar, database drivers
3. **F022: List-schemas command** - Uses Cobra framework
4. **F005: GitHub Actions CI/CD** - Includes `go mod verify` in pipeline
