---
description: "Implementation tasks for F006 - Cobra CLI Framework Integration"
---

# Tasks: Cobra CLI Framework Integration

**Feature**: F006 - Cobra CLI Framework Integration
**Branch**: `004-f006-cobra-cli`
**Input**: Design documents from `/specs/004-f006-cobra-cli/`

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

**Tests**: Tests are OPTIONAL in this feature. Per the specification, no explicit TDD requirement was requested. Test tasks are included for framework validation but not for business logic (as commands are placeholders).

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1-US5)
- Exact file paths included in descriptions

## Path Conventions
Project uses standard Go CLI structure:
- `cmd/sourcebox/` - CLI entry point
- `cmd/sourcebox/cmd/` - Command definitions
- Tests colocated with implementation (`*_test.go` pattern)

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Install tools and dependencies needed for all user stories

- [X] T001 [P] Install cobra-cli tool: `go install github.com/spf13/cobra-cli@latest`
- [X] T002 [P] Add Cobra dependency: `go get github.com/spf13/cobra@latest`
- [X] T003 [P] Add testify dependency: `go get github.com/stretchr/testify@latest`
- [X] T004 Run `go mod tidy` to clean up dependencies

**Checkpoint**: Dependencies installed - ready for Cobra initialization

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core Cobra structure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T005 Initialize Cobra structure: `cobra-cli init --config .cobra.yaml` (creates `cmd/sourcebox/main.go` and `cmd/sourcebox/cmd/root.go`)
- [X] T006 Update `cmd/sourcebox/main.go` to add version variable for build-time injection
- [X] T007 Export `SetVersion()` function in `cmd/sourcebox/cmd/root.go`
- [X] T008 Update `main.go` to call `cmd.SetVersion(version)` before `cmd.Execute()`

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Get Help and Documentation (Priority: P1) 🎯 MVP

**Goal**: Enable users to discover SourceBox capabilities and learn proper usage through comprehensive command-line help

**Independent Test**: Run help commands (`sourcebox --help`, `sourcebox seed --help`) and verify output contains clear descriptions, usage examples, and available commands

### Implementation for User Story 1

- [X] T009 [US1] Customize root command metadata in `cmd/sourcebox/cmd/root.go`:
  - Set Use: "sourcebox"
  - Set Short: "Generate realistic, verticalized demo data instantly"
  - Set Long: Multi-paragraph description from quickstart.md
- [X] T010 [US1] Add Example field to root command showing common usage patterns
- [X] T011 [US1] Verify help output contains all required fields: `go run ./cmd/sourcebox --help`

**Checkpoint**: At this point, User Story 1 should be fully functional - users can access comprehensive help documentation

---

## Phase 4: User Story 2 - Check Version Information (Priority: P1) 🎯 MVP

**Goal**: Enable users to verify installed version for troubleshooting and support

**Independent Test**: Run `sourcebox --version` and verify accurate version string is displayed

### Implementation for User Story 2

- [X] T012 [US2] Update Makefile to add VERSION variable: `VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")`
- [X] T013 [US2] Update Makefile build target to add LDFLAGS: `LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"`
- [X] T014 [US2] Update Makefile build command to use LDFLAGS: `go build $(LDFLAGS) -o dist/sourcebox ./cmd/sourcebox`
- [X] T015 [US2] Build and verify version injection: `make build && ./dist/sourcebox --version`

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Control Output Verbosity (Priority: P2)

**Goal**: Enable users to control output detail level for different use cases (debugging vs production)

**Independent Test**: Run commands with `--verbose` and `--quiet` flags and verify output volume changes appropriately

### Implementation for User Story 3

- [X] T016 [US3] Add global flag variables in `cmd/sourcebox/cmd/root.go`:
  - `var verbose bool`
  - `var quiet bool`
- [X] T017 [US3] Add PersistentFlags in root command init():
  - `rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")`
  - `rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress non-error output")`
- [X] T018 [US3] Verify flags parse correctly: `go run ./cmd/sourcebox --verbose --help`
- [X] T019 [US3] Verify flags appear in help output under "Global Flags" section

**Checkpoint**: All P1 and P2 user stories should now be independently functional

---

## Phase 6: User Story 4 - Use Custom Configuration (Priority: P3)

**Goal**: Enable users to specify custom configuration files for multi-environment workflows

**Independent Test**: Run command with `--config=/path/to/file` and verify flag is accepted without error

### Implementation for User Story 4

- [X] T020 [US4] Add config flag variable in `cmd/sourcebox/cmd/root.go`: `var cfgFile string`
- [X] T021 [US4] Add PersistentFlag in root command init():
  - `rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path (default: ~/.sourcebox.yaml)")`
- [X] T022 [US4] Verify config flag parses correctly: `go run ./cmd/sourcebox --config=/tmp/test.yaml --help`

**Checkpoint**: All P1, P2, and P3 user stories should now be independently functional

---

## Phase 7: User Story 5 - Access Core Commands (Priority: P1) 🎯 MVP

**Goal**: Establish command structure for core SourceBox functionality (seed, list-schemas) with clear help text

**Independent Test**: Verify commands are registered, show appropriate help text, and acknowledge future implementation

### Implementation for User Story 5

#### Seed Command

- [X] T023 [US5] Scaffold seed command: `cobra-cli add seed` (creates `cmd/sourcebox/cmd/seed.go`)
- [X] T024 [US5] Customize seed command in `cmd/sourcebox/cmd/seed.go`:
  - Set Use: "seed <database>"
  - Set Short: "Seed a database with realistic demo data" (<60 chars)
  - Set Long: Multi-paragraph description from quickstart.md
  - Set Example: 3 practical usage examples from quickstart.md
  - Set Args: `cobra.ExactArgs(1)` for database argument validation
- [X] T025 [US5] Add local flags to seed command init():
  - `--schema` (required, string, short: -s)
  - `--records` (optional, int, short: -n, default: 1000)
  - `--host` (optional, string, default: "localhost")
  - `--port` (optional, int, default: 0 for auto-detect)
  - `--user` (optional, string, default: "root")
  - `--password` (optional, string)
  - `--db-name` (optional, string, default: "demo")
  - `--output` (optional, string, for SQL file export)
  - `--dry-run` (optional, bool)
- [X] T026 [US5] Mark schema flag as required: `seedCmd.MarkFlagRequired("schema")`
- [X] T027 [US5] Implement placeholder Run function that prints:
  - "Seed command - implementation coming in F021"
  - Database name from args[0]
  - Schema from flag
  - Records from flag

#### List-Schemas Command

- [X] T028 [P] [US5] Scaffold list-schemas command: `cobra-cli add list-schemas` (creates `cmd/sourcebox/cmd/list_schemas.go`)
- [X] T029 [P] [US5] Customize list-schemas command in `cmd/sourcebox/cmd/list_schemas.go`:
  - Set Use: "list-schemas"
  - Set Aliases: []string{"ls"}
  - Set Short: "List all available data schemas" (<60 chars)
  - Set Long: Multi-paragraph description from quickstart.md
  - Set Example: 2 usage examples (full name and alias) from quickstart.md
- [X] T030 [P] [US5] Implement placeholder Run function that prints:
  - "List-schemas command - implementation coming in F022"
  - List of sample schemas (fintech-loans, healthcare-patients, retail-orders)

#### Verification

- [X] T031 [US5] Build binary: `make build`
- [X] T032 [US5] Verify seed command registration: `./dist/sourcebox --help` shows "seed" in available commands
- [X] T033 [US5] Verify list-schemas command registration: `./dist/sourcebox --help` shows "list-schemas" in available commands
- [X] T034 [US5] Verify seed help output: `./dist/sourcebox seed --help` shows comprehensive help
- [X] T035 [US5] Verify list-schemas help output: `./dist/sourcebox list-schemas --help` shows comprehensive help
- [X] T036 [US5] Verify alias works: `./dist/sourcebox ls --help` shows same output as list-schemas
- [X] T037 [US5] Test seed placeholder execution: `./dist/sourcebox seed mysql --schema=fintech-loans`
- [X] T038 [US5] Test list-schemas placeholder execution: `./dist/sourcebox list-schemas`

**Checkpoint**: All user stories should now be independently functional

---

## Phase 8: Framework Validation Tests (Optional Quality Assurance)

**Purpose**: Validate Cobra framework integration with unit tests (not TDD, but quality verification)

**Note**: These tests validate framework behavior, not business logic (commands are placeholders for F021/F022)

- [X] T039 [P] Create `cmd/sourcebox/cmd/root_test.go`:
  - Test root command help output contains expected text
  - Test version flag displays version correctly
  - Test commands are registered (seed, list-schemas)
- [X] T040 [P] Create `cmd/sourcebox/cmd/seed_test.go`:
  - Test seed help output contains comprehensive information
  - Test required schema flag validation
  - Test flag parsing (valid and invalid combinations)
- [X] T041 [P] Create `cmd/sourcebox/cmd/list_schemas_test.go`:
  - Test list-schemas help output
  - Test alias (ls) works correctly
- [X] T042 Run all tests: `go test ./cmd/sourcebox/cmd/... -v`
- [X] T043 Check test coverage: `go test ./cmd/sourcebox/cmd/... -coverprofile=coverage.txt && go tool cover -func=coverage.txt`
- [X] T044 Verify coverage >80% for cmd/ package (excluding init() and static help text)

**Checkpoint**: All tests passing, coverage target met

---

## Phase 9: Polish & Cross-Cutting Concerns

**Purpose**: Final improvements affecting multiple user stories

- [X] T045 [P] Update CLAUDE.md with Cobra framework context (copy from plan.md "Cobra CLI Framework" section)
- [X] T046 [P] Verify performance: `time ./dist/sourcebox --help` <1 second
- [X] T047 [P] Verify performance: `time ./dist/sourcebox --version` <100ms
- [X] T048 Run quickstart.md verification workflow end-to-end
- [X] T049 Test on macOS (current platform)
- [X] T050 Test on Linux (via CI or VM)
- [X] T051 Test on Windows (via CI or VM)
- [X] T052 Final build: `make build` produces working binary
- [X] T053 Commit all changes with message from quickstart.md

**Checkpoint**: Feature complete, validated on all platforms

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup (T001-T004) completion - BLOCKS all user stories
- **User Stories (Phase 3-7)**: All depend on Foundational (T005-T008) completion
  - User stories can then proceed in parallel (if staffed)
  - Or sequentially in priority order (P1 → P2 → P3)
- **Framework Validation (Phase 8)**: Depends on all user stories being complete
- **Polish (Phase 9)**: Depends on all desired user stories and tests being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (T008) - No dependencies on other stories
- **User Story 2 (P1)**: Can start after Foundational (T008) - No dependencies on other stories
- **User Story 3 (P2)**: Can start after Foundational (T008) - No dependencies on other stories
- **User Story 4 (P3)**: Can start after Foundational (T008) - No dependencies on other stories
- **User Story 5 (P1)**: Can start after Foundational (T008) - Integrates with US1, US2, US3 flags

### Within Each User Story

- US1: Sequential (T009 → T010 → T011)
- US2: Sequential (T012 → T013 → T014 → T015) - Makefile edits must be ordered
- US3: Sequential (T016 → T017 → T018 → T019)
- US4: Sequential (T020 → T021 → T022)
- US5: Seed and List-schemas scaffolding can be parallel, verification sequential

### Parallel Opportunities

- Phase 1 Setup: All tasks (T001-T003) marked [P] can run in parallel
- Phase 2 Foundational: Sequential (building on each other)
- Phase 3-7 User Stories: Once Phase 2 completes, all user stories can start in parallel
  - US1 (T009-T011) can run in parallel with US2 (T012-T015)
  - US3 (T016-T019) can run in parallel with US4 (T020-T022)
  - US5 seed and list-schemas scaffolding (T023-T027 and T028-T030) can run in parallel
- Phase 8 Tests: Test file creation (T039-T041) marked [P] can run in parallel

---

## Parallel Example: User Story 5 (Seed + List-Schemas)

```bash
# Launch seed and list-schemas scaffolding in parallel:
Task: "Scaffold seed command: cobra-cli add seed"
Task: "Scaffold list-schemas command: cobra-cli add list-schemas"

# Customize seed command (T024-T027)
# Customize list-schemas command (T029-T030) [can run parallel to seed]

# Then verification (T031-T038) runs sequentially
```

---

## Implementation Strategy

### MVP First (P1 User Stories Only)

1. Complete Phase 1: Setup (T001-T004)
2. Complete Phase 2: Foundational (T005-T008) - CRITICAL
3. Complete Phase 3: User Story 1 (T009-T011) - Help system
4. Complete Phase 4: User Story 2 (T012-T015) - Version display
5. Complete Phase 7: User Story 5 (T023-T038) - Core commands
6. **STOP and VALIDATE**: Test all P1 stories independently
7. Deploy/demo if ready

**Defer to later**: US3 (verbose/quiet flags - P2), US4 (config flag - P3)

### Incremental Delivery

1. Setup + Foundational → Foundation ready
2. Add US1 (Help) → Test independently → Deploy/Demo
3. Add US2 (Version) → Test independently → Deploy/Demo (MVP core!)
4. Add US5 (Commands) → Test independently → Deploy/Demo (MVP complete!)
5. Add US3 (Verbosity) → Test independently → Deploy/Demo
6. Add US4 (Config) → Test independently → Deploy/Demo
7. Each story adds value without breaking previous stories

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together (T001-T008)
2. Once Foundational is done:
   - Developer A: User Story 1 (Help) + User Story 2 (Version)
   - Developer B: User Story 3 (Verbosity) + User Story 4 (Config)
   - Developer C: User Story 5 (Commands)
3. Stories complete and integrate independently

---

## Success Criteria Mapping

Tasks mapped to success criteria from spec.md:

- **SC-001**: Help in <3 seconds → T009-T011, T034-T036 (US1)
- **SC-002**: 100% commands have help → T024, T029 (US5)
- **SC-003**: Version in <1 second → T012-T015, T047 (US2)
- **SC-004**: Discover all commands → T032-T033 (US5)
- **SC-005**: Clear error messages → T037 (US5 - placeholder acknowledgment)
- **SC-006**: <100ms response → T046-T047 (Phase 9)
- **SC-007**: Verbose mode detail → T016-T019 (US3 - framework only, behavior in F021/F022)
- **SC-008**: Quiet mode reduction → T016-T019 (US3 - framework only, behavior in F021/F022)
- **SC-009**: Config path recognition → T020-T022 (US4 - framework only, behavior in F021/F022)
- **SC-010**: Alias accessibility → T028-T030, T036 (US5)

---

## Notes

- [P] tasks = different files, no dependencies, can run in parallel
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Framework validation tests (Phase 8) are optional but recommended
- Business logic tests deferred to F021 (Seed) and F022 (List-schemas)
- Commit after each user story phase completion
- Stop at any checkpoint to validate story independently
- Performance verification (Phase 9) validates constitutional requirements (<100ms help/version)

---

## Task Count Summary

- **Phase 1 (Setup)**: 4 tasks
- **Phase 2 (Foundational)**: 4 tasks
- **Phase 3 (US1 - Help)**: 3 tasks
- **Phase 4 (US2 - Version)**: 4 tasks
- **Phase 5 (US3 - Verbosity)**: 4 tasks
- **Phase 6 (US4 - Config)**: 3 tasks
- **Phase 7 (US5 - Commands)**: 16 tasks
- **Phase 8 (Tests)**: 6 tasks
- **Phase 9 (Polish)**: 9 tasks

**Total**: 53 tasks

**MVP Scope** (P1 stories only): T001-T015, T023-T038, T045-T053 = 42 tasks

**Parallel Opportunities**:
- Setup: 3 tasks can run in parallel
- User Stories: 5 stories can run in parallel after Foundational
- Tests: 3 test files can be created in parallel
- Estimated 40-50% time savings with parallel execution
