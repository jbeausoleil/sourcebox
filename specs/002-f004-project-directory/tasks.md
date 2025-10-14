# Tasks: F004 - Project Directory Structure & Build System

**Input**: Design documents from `/specs/002-f004-project-directory/`
**Prerequisites**: plan.md ✅, spec.md ✅, research.md ✅, quickstart.md ✅

**Tests**: Tests are NOT requested in this feature specification - no test tasks generated.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions
- **Single project**: Go CLI at repository root
- Structure: `cmd/sourcebox/`, `pkg/`, `schemas/`, `docker/`, `docs/`, `examples/`, `dist/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Create foundational directory structure and build configuration

- [X] T001 [P] Create `/cmd/sourcebox/` directory for CLI entry point
- [X] T002 [P] Create `/pkg/` directory with subdirectories (`generators/`, `schema/`, `database/`)
- [X] T003 [P] Create `/schemas/` directory for schema JSON definitions (future use)
- [X] T004 [P] Create `/docker/` directory for container files (future use)
- [X] T005 [P] Create `/docs/` directory for documentation (future use)
- [X] T006 [P] Create `/examples/` directory for usage examples (future use)
- [X] T007 [P] Create `/dist/` directory with `.gitkeep` file
- [X] T008 Update `.gitignore` to exclude `/dist/` and `coverage.txt`

**Checkpoint**: All directories exist and are properly tracked/ignored in git ✓

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core build system that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T009 Create `Makefile` in repository root with all required variables (BINARY_NAME, VERSION, BUILD_DIR, LDFLAGS, PLATFORMS)
- [X] T010 [P] Implement `make help` target (default) in Makefile showing all available commands
- [X] T011 [P] Declare all targets as `.PHONY` in Makefile

**Checkpoint**: Build system foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Developer Navigates Codebase (Priority: P1) 🎯 MVP

**Goal**: Enable developers to quickly locate code using standard Go project layout

**Independent Test**: Clone the repository, review the directory structure, and successfully locate the appropriate package directory for schema, database, or generator logic within 2 minutes.

### Implementation for User Story 1

- [ ] T012 [US1] Create `cmd/sourcebox/main.go` with minimal bootstrap code (version flag, placeholder output)
- [ ] T013 [US1] Create `cmd/sourcebox/main_test.go` with placeholder test (TestPlaceholder verifies build system)
- [ ] T014 [US1] Add version variable in `main.go` that gets overridden at build time via ldflags

**Checkpoint**: Directory structure is navigable and main.go exists with basic functionality

---

## Phase 4: User Story 2 - Developer Builds Application (Priority: P1) 🎯 MVP

**Goal**: Enable developers to build the application for their local platform in under 30 seconds

**Independent Test**: Run `make build` from the project root and verify a working binary is created in the /dist directory that executes without errors.

### Implementation for User Story 2

- [ ] T015 [US2] Implement `make build` target in Makefile (builds for current platform → dist/sourcebox)
- [ ] T016 [US2] Configure VERSION variable in Makefile using `git describe --tags --always --dirty`
- [ ] T017 [US2] Configure LDFLAGS in Makefile with `-s -w` (strip symbols) and `-X main.version=$(VERSION)`
- [ ] T018 [US2] Verify binary is executable and includes version information

**Checkpoint**: `make build` produces working binary in < 30 seconds

---

## Phase 5: User Story 3 - Developer Runs Test Suite (Priority: P1) 🎯 MVP

**Goal**: Enable developers to verify their changes haven't broken existing functionality

**Independent Test**: Run `make test` and verify tests execute successfully with coverage reporting, even if only placeholder tests exist initially.

### Implementation for User Story 3

- [ ] T019 [US3] Implement `make test` target in Makefile (runs tests with race detection and coverage)
- [ ] T020 [US3] Configure test target to run `go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...`
- [ ] T021 [US3] Verify placeholder test in `main_test.go` passes and generates coverage report

**Checkpoint**: `make test` runs successfully with coverage reporting

---

## Phase 6: User Story 4 - Developer Installs CLI Locally (Priority: P2)

**Goal**: Enable developers to install and use the SourceBox CLI from anywhere on their system

**Independent Test**: Run `make install`, then open a new terminal and execute `sourcebox --version` from any directory to verify installation.

### Implementation for User Story 4

- [ ] T022 [US4] Implement `make install` target in Makefile (installs binary to $GOPATH/bin)
- [ ] T023 [US4] Configure install target to build and copy binary to $GOPATH/bin/sourcebox
- [ ] T024 [US4] Verify installed binary is globally accessible and shows correct version

**Checkpoint**: `make install` enables global `sourcebox` command access

---

## Phase 7: User Story 5 - Maintainer Creates Release Binaries (Priority: P2)

**Goal**: Enable project maintainers to create distribution binaries for all supported platforms

**Independent Test**: Run `make build-all`, verify 5 compressed binaries (.gz files) are created in /dist, and confirm each is under 50MB.

### Implementation for User Story 5

- [ ] T025 [US5] Define PLATFORMS variable in Makefile (darwin/amd64, darwin/arm64, linux/amd64, linux/arm64, windows/amd64)
- [ ] T026 [P] [US5] Implement build target for macOS Intel (darwin/amd64) with gzip compression
- [ ] T027 [P] [US5] Implement build target for macOS Apple Silicon (darwin/arm64) with gzip compression
- [ ] T028 [P] [US5] Implement build target for Linux x86_64 (linux/amd64) with gzip compression
- [ ] T029 [P] [US5] Implement build target for Linux ARM64 (linux/arm64) with gzip compression
- [ ] T030 [P] [US5] Implement build target for Windows x86_64 (windows/amd64 with .exe extension) with gzip compression
- [ ] T031 [US5] Implement `make build-all` target that orchestrates all 5 platform builds
- [ ] T032 [US5] Verify all compressed binaries are under 50MB (constitutional requirement)

**Checkpoint**: `make build-all` creates 5 platform binaries in under 2 minutes (with parallel builds)

---

## Phase 8: User Story 6 - Developer Cleans Build Artifacts (Priority: P3)

**Goal**: Enable developers to clean up build artifacts before committing or to free disk space

**Independent Test**: Run `make build` to create artifacts, then run `make clean` and verify /dist directory and coverage files are removed.

### Implementation for User Story 6

- [ ] T033 [US6] Implement `make clean` target in Makefile (removes dist/ and coverage.txt)
- [ ] T034 [US6] Verify clean target removes all build artifacts without affecting source files
- [ ] T035 [US6] Verify `git status` shows no leftover build artifacts after `make clean`

**Checkpoint**: `make clean` successfully removes all build artifacts

---

## Phase 9: Polish & Cross-Cutting Concerns

**Purpose**: Final integration, documentation, and validation

- [ ] T036 Update `CLAUDE.md` with build system information (project structure, Makefile targets, build conventions)
- [ ] T037 Verify all tasks from `quickstart.md` pass successfully
- [ ] T038 Run full build workflow: `make clean && make test && make build && make build-all`
- [ ] T039 Verify version injection works correctly with `./dist/sourcebox --version`
- [ ] T040 Confirm constitutional compliance: build time < 2 min, binaries < 50MB compressed

**Checkpoint**: Feature fully integrated, documented, and validated

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3-8)**: All depend on Foundational phase completion
  - US1 (Navigate Codebase): Can start after Foundational - No dependencies on other stories
  - US2 (Build Application): Depends on US1 (needs main.go)
  - US3 (Run Tests): Depends on US1 (needs main_test.go)
  - US4 (Install CLI): Depends on US2 (needs build target)
  - US5 (Release Binaries): Depends on US2 (needs build infrastructure)
  - US6 (Clean Artifacts): Independent - can implement anytime after Foundational
- **Polish (Phase 9)**: Depends on all user stories being complete

### User Story Dependencies

```
Setup (Phase 1) → Foundational (Phase 2)
                        ↓
                     US1 (Phase 3) ← Foundation for code structure
                        ↓
        ┌───────────────┼───────────────┐
        ↓               ↓               ↓
     US2 (Phase 4)   US3 (Phase 5)   US6 (Phase 8)
        ↓
        ├───────────────┐
        ↓               ↓
     US4 (Phase 6)   US5 (Phase 7)
```

### Within Each User Story

- US1: main.go before main_test.go (test needs main package)
- US2: Makefile variables before build targets
- US3: Test infrastructure depends on US1 test file
- US4: Install depends on US2 build target
- US5: All platform targets can run in parallel (marked [P])
- US6: Independent implementation

### Parallel Opportunities

**Phase 1 (Setup)**: All tasks (T001-T008) can run in parallel - different directories

**Phase 2 (Foundational)**: T010 and T011 can run in parallel with T009

**Phase 5 (User Story 5)**: Tasks T026-T030 can run in parallel - different platform builds

```bash
# Example: Launch all platform builds in parallel
Task: "Build macOS Intel binary with gzip compression"
Task: "Build macOS ARM binary with gzip compression"
Task: "Build Linux x86_64 binary with gzip compression"
Task: "Build Linux ARM64 binary with gzip compression"
Task: "Build Windows x86_64 binary with gzip compression"
```

---

## Parallel Example: Setup Phase

```bash
# Launch all directory creation tasks together:
Task: "Create /cmd/sourcebox/ directory"
Task: "Create /pkg/ directory with subdirectories"
Task: "Create /schemas/ directory"
Task: "Create /docker/ directory"
Task: "Create /docs/ directory"
Task: "Create /examples/ directory"
Task: "Create /dist/ directory"
Task: "Update .gitignore"
```

---

## Implementation Strategy

### MVP First (User Stories 1-3 Only)

1. Complete Phase 1: Setup → All directories exist
2. Complete Phase 2: Foundational → Makefile foundation ready
3. Complete Phase 3: US1 → Code structure navigable
4. Complete Phase 4: US2 → Local builds work
5. Complete Phase 5: US3 → Tests run successfully
6. **STOP and VALIDATE**: Test build workflow independently
7. This is the **minimum viable build system** (MVP!)

### Incremental Delivery

1. **Foundation** (Phases 1-2) → Project structure + Makefile ready
2. **MVP** (Add US1-US3) → Navigable code + local builds + tests
3. **Developer Convenience** (Add US4) → Global CLI installation
4. **Release Capability** (Add US5) → Cross-platform binaries
5. **Cleanup** (Add US6) → Build artifact management
6. Each phase adds value without breaking previous functionality

### Sequential Recommended (Single Developer)

Since most tasks depend on previous work:

1. Complete Setup (Phase 1) first
2. Complete Foundational (Phase 2) second
3. Then proceed with user stories in priority order:
   - US1 (P1) → US2 (P1) → US3 (P1) → US4 (P2) → US5 (P2) → US6 (P3)
4. Only US5 platform builds (T026-T030) benefit from parallelization

---

## Task Summary

- **Total Tasks**: 40 tasks
- **Setup Phase**: 8 tasks (all parallelizable)
- **Foundational Phase**: 3 tasks (1 parallel opportunity)
- **User Story 1**: 3 tasks (sequential)
- **User Story 2**: 4 tasks (sequential)
- **User Story 3**: 3 tasks (sequential)
- **User Story 4**: 3 tasks (sequential)
- **User Story 5**: 8 tasks (5 platform builds parallelizable)
- **User Story 6**: 3 tasks (sequential)
- **Polish**: 5 tasks (sequential)

**Parallelization Opportunities**:
- Phase 1: 8 tasks can run in parallel
- Phase 2: 2 tasks can run in parallel
- US5: 5 platform build tasks can run in parallel

**Estimated Effort**: 1 day (small effort, foundational feature)

**MVP Scope**: Phases 1-5 (Setup + Foundational + US1-US3) = **18 tasks**

---

## Notes

- [P] tasks = different files/directories, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story builds on previous stories (US2 needs US1, US4 needs US2, etc.)
- No tests generated - not requested in feature specification
- Stop at US3 completion for minimal viable build system
- Commit after each user story phase for clear checkpoints
- Verify quickstart.md validation passes before marking feature complete
