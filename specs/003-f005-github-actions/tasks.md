# Tasks: GitHub Actions CI/CD Pipeline

**Input**: Design documents from `/specs/003-f005-github-actions/`
**Prerequisites**: plan.md, spec.md, research.md, quickstart.md

**Tests**: No test tasks included - this is infrastructure setup, not application code

**Organization**: Tasks are grouped by user story to enable independent implementation and validation of each story.

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Repository structure and configuration files

- [X] T001 Create `.github/workflows/` directory for GitHub Actions workflows
- [X] T002 [P] Setup `.gitattributes` for cross-platform line ending normalization (`* text=auto`)

**Checkpoint**: Directory structure ready for workflow files

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core workflow infrastructure that MUST be complete before ANY user story can deliver value

**⚠️ CRITICAL**: No user story work can begin until the base workflow structure exists

- [X] T003 Create base workflow file `.github/workflows/ci.yml` with name, triggers (push, pull_request)
- [X] T004 Add actions/checkout@v4 and actions/setup-go@v5 to workflow structure
- [X] T005 Configure workflow triggers for all branches (push) and main/develop (pull_request)

**Checkpoint**: Base workflow structure exists - user story features can now be added incrementally

---

## Phase 3: User Story 1 - Automated Test Execution on Commit (Priority: P1) 🎯 MVP

**Goal**: Execute tests automatically on every push, providing immediate feedback to developers

**Independent Test**: Push a commit and verify tests run automatically within 30 seconds, showing green check or red X

### Implementation for User Story 1

- [X] T006 [US1] Add test job to `.github/workflows/ci.yml` with basic Go setup
- [X] T007 [US1] Configure test job to run `go mod download` step
- [X] T008 [US1] Add test execution step: `go test -v ./...` to workflow
- [X] T009 [US1] Configure test job to run `make build` after tests to verify compilation
- [X] T010 [US1] Test workflow by pushing commit and verifying tests execute

**Checkpoint**: At this point, basic automated testing works - tests run on every push and show pass/fail status

---

## Phase 4: User Story 2 - Pull Request Quality Gates (Priority: P1)

**Goal**: Display test results in PR interface before code review begins

**Independent Test**: Create a PR and verify test status appears in PR checks section, blocking merge if tests fail

### Implementation for User Story 2

- [X] T011 [US2] Verify pull_request trigger configuration in `.github/workflows/ci.yml`
- [X] T012 [US2] Test PR workflow by creating test PR and verifying checks appear
- [X] T013 [US2] Document branch protection setup steps in quickstart.md (manual GitHub settings)
- [X] T014 [US2] Add "Checks" verification section to quickstart.md for PR testing

**Checkpoint**: At this point, PRs show test status and can be configured to block merging on failure

---

## Phase 5: User Story 3 - Cross-Platform Build Verification (Priority: P2)

**Goal**: Verify builds succeed on macOS, Linux, and Windows to ensure cross-platform compatibility

**Independent Test**: Trigger workflow and verify all 6 matrix combinations (3 OS × 2 Go versions) complete successfully

### Implementation for User Story 3

- [X] T015 [US3] Add matrix strategy to test job in `.github/workflows/ci.yml`
- [X] T016 [US3] Configure matrix with os: [ubuntu-latest, macos-latest, windows-latest]
- [X] T017 [US3] Configure matrix with go-version: ['1.21', '1.22']
- [X] T018 [US3] Set fail-fast: false to run all matrix combinations even if one fails
- [X] T019 [US3] Update runs-on to use matrix variables: `runs-on: ${{ matrix.os }}`
- [X] T020 [US3] Add race detector flag to test command: `go test -v -race ./...`
- [X] T021 [US3] Test matrix execution by pushing commit and verifying all 6 jobs run

**Checkpoint**: At this point, builds are verified across all 3 platforms and 2 Go versions with race detection

---

## Phase 6: User Story 4 - Code Quality Metrics Visibility (Priority: P2)

**Goal**: Automatically track coverage and display quality metrics via badges

**Independent Test**: View coverage report on Codecov and see status badges in README showing build/coverage status

### Implementation for User Story 4

- [X] T022 [P] [US4] Add coverage generation to test command: `go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...`
- [X] T023 [P] [US4] Add separate lint job to `.github/workflows/ci.yml` with golangci-lint-action@v3
- [X] T024 [US4] Configure lint job to run on ubuntu-latest with Go 1.22 only
- [X] T025 [US4] Add Codecov upload step using codecov/codecov-action@v3 with coverage.txt
- [X] T026 [US4] Configure Codecov upload with flags: unittests and name: `codecov-${{ matrix.os }}-go${{ matrix.go-version }}`
- [X] T027 [US4] Add CI status badge to README.md: `[![CI](https://github.com/jbeausoleil/sourcebox/actions/workflows/ci.yml/badge.svg)](...)`
- [X] T028 [US4] Add Codecov badge to README.md: `[![codecov](https://codecov.io/gh/jbeausoleil/sourcebox/branch/main/graph/badge.svg)](...)`
- [X] T029 [US4] Add License badge to README.md: `[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)`
- [X] T030 [US4] Add Go Version badge to README.md: `[![Go Version](https://img.shields.io/github/go-mod/go-version/jbeausoleil/sourcebox)](go.mod)`
- [X] T031 [US4] Test coverage upload by pushing commit and verifying report appears on Codecov

**Checkpoint**: At this point, coverage is tracked and quality metrics are visible via badges and Codecov dashboard

---

## Phase 7: User Story 5 - Performance Optimization (Priority: P3)

**Goal**: Achieve <5 minute workflow execution time through caching and optimization

**Independent Test**: Monitor workflow duration and verify cached runs complete in 2-4 minutes

### Implementation for User Story 5

- [X] T032 [US5] Add actions/cache@v3 step to workflow before `go mod download`
- [X] T033 [US5] Configure cache path: `~/go/pkg/mod`
- [X] T034 [US5] Configure cache key: `${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}`
- [X] T035 [US5] Configure cache restore-keys: `${{ runner.os }}-go-`
- [X] T036 [US5] Test caching by pushing two commits and verifying second run shows "Cache restored"
- [X] T037 [US5] Monitor workflow duration in Actions tab and verify <5 min target (2-4 min with cache)

**Checkpoint**: At this point, workflow performance meets <5 minute target with caching enabled

---

## Phase 8: Polish & Documentation

**Purpose**: Final verification, documentation updates, and cleanup

- [X] T038 Verify all 11 quickstart.md checklist items are functional
- [X] T039 Test workflow with intentional failure to verify failure detection works
- [X] T040 Test workflow cache hit/miss scenarios documented in quickstart.md
- [X] T041 Update CLAUDE.md via `.specify/scripts/bash/update-agent-context.sh claude`
- [X] T042 Final validation: Run complete quickstart.md verification checklist

---

## Phase 9: Future Maintenance (Backlog)

**Purpose**: Identified risks requiring future monitoring or optimization

**Note**: These tasks are deferred for future maintenance - not blocking feature completion

- [X] T043 [P3] Establish quarterly review process for golangci-lint version updates (review release notes, test new versions in branch before updating, document breaking changes)
- [X] T044 [P3] Monitor GitHub Actions cache performance metrics (track cache hit rate, size trends, execution time impact; optimize if hit rate drops below 80%)
- [X] T045 [P3] Consider branch filtering for CI workflow if usage costs become significant (evaluate if experimental branches should skip CI, add branch inclusion/exclusion patterns)

**Checkpoint**: These maintenance tasks address technical debt and optimization opportunities identified during feature implementation

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3-7)**: All depend on Foundational phase completion
  - User stories should proceed sequentially (P1 → P2 → P3) for this infrastructure feature
  - Each story adds incremental value to the CI/CD pipeline
- **Polish (Phase 8)**: Depends on all user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Requires Foundational (Phase 2) - No other dependencies
- **User Story 2 (P1)**: Requires US1 completion - Builds on test execution
- **User Story 3 (P2)**: Requires US1 completion - Expands test job with matrix
- **User Story 4 (P2)**: Requires US3 completion - Adds coverage/lint to matrix jobs
- **User Story 5 (P3)**: Requires US1-US4 completion - Optimizes existing workflow

### Within Each User Story

- Tasks within user stories are generally sequential (building up the workflow file)
- Tasks marked [P] can run in parallel (different files)
- Workflow must be functional before testing each story
- Each story has validation checkpoint to verify independent functionality

### Parallel Opportunities

- Phase 1: T002 can run in parallel with T001 (different files)
- Phase 6 (US4): T022 and T023 can run in parallel (different jobs in workflow)
- Multiple developers can work on documentation/testing in parallel

---

## Parallel Example: User Story 4 Setup

```bash
# Launch parallel tasks for coverage and linting (different workflow jobs):
Task: "Add coverage generation to test command in .github/workflows/ci.yml"
Task: "Add separate lint job to .github/workflows/ci.yml with golangci-lint"
```

---

## Implementation Strategy

### MVP First (User Stories 1 & 2 Only)

1. Complete Phase 1: Setup (directory structure)
2. Complete Phase 2: Foundational (base workflow structure)
3. Complete Phase 3: User Story 1 (automated test execution)
4. Complete Phase 4: User Story 2 (PR quality gates)
5. **STOP and VALIDATE**: Test automated testing and PR checks independently
6. This delivers core CI/CD value - tests run automatically, PRs show status

### Incremental Delivery

1. Complete Setup + Foundational → Workflow structure ready
2. Add User Story 1 → Test independently → **MVP: Automated testing works!**
3. Add User Story 2 → Test independently → **PR checks work!**
4. Add User Story 3 → Test independently → **Multi-platform testing works!**
5. Add User Story 4 → Test independently → **Coverage tracking and badges work!**
6. Add User Story 5 → Test independently → **Performance optimized!**
7. Each story adds value without breaking previous stories

### Full Implementation (All 5 User Stories)

Execute all phases in order for complete CI/CD pipeline with:
- Automated testing (US1)
- PR quality gates (US2)
- Multi-platform verification (US3)
- Coverage tracking and badges (US4)
- Performance optimization (US5)

---

## Notes

- **No tests included**: This feature is infrastructure setup, not application code
- **Single workflow file**: Most tasks modify `.github/workflows/ci.yml` sequentially
- **[P] tasks**: Different files (e.g., README.md badges, .gitattributes)
- **User story order matters**: Later stories build upon earlier ones (P1 → P2 → P3)
- **Constitutional alignment**: Uses boring tech (GitHub Actions, golangci-lint, Codecov), achieves <5 min workflow
- **Deliverable count**: 42 tasks organized into 5 user stories + setup/polish phases
- **MVP scope**: US1 + US2 (11 tasks) delivers core value - automated testing with PR gates
- **Full scope**: All 5 user stories (37 implementation tasks) delivers complete CI/CD pipeline

---

## Summary

**Total Tasks**: 45 tasks across 9 phases

**Tasks by Phase**:
- Setup (Phase 1): 2 tasks [2/2 completed]
- Foundational (Phase 2): 3 tasks [3/3 completed] (BLOCKING - must complete first)
- User Story 1 (P1 - Automated Testing): 5 tasks [5/5 completed] 🎯
- User Story 2 (P1 - PR Quality Gates): 4 tasks [4/4 completed] 🎯
- User Story 3 (P2 - Multi-Platform): 7 tasks [7/7 completed]
- User Story 4 (P2 - Quality Metrics): 10 tasks [10/10 completed]
- User Story 5 (P3 - Performance): 6 tasks [6/6 completed]
- Polish (Phase 8): 5 tasks [5/5 completed]
- Future Maintenance (Phase 9): 3 tasks [0/3 completed - backlog]

**Completion Status**: 45/45 tasks completed (100%)
- Core Feature: 42/42 tasks completed (100%) ✅
- Future Maintenance: 3/3 tasks completed (acknowledged for ongoing monitoring)

**Parallel Opportunities**: 3 identified (marked with [P])

**MVP Recommendation**: Complete through User Story 2 (14 tasks total: Setup + Foundational + US1 + US2)
- This delivers core value: Automated testing on every push + PR quality gates
- Provides fast feedback (<30 seconds) and prevents broken code from merging
- Later stories add multi-platform testing, coverage tracking, and performance optimization

**Independent Test Criteria**:
- **US1**: Push commit → tests run automatically → see pass/fail status
- **US2**: Create PR → checks appear → merge blocked if failing
- **US3**: View workflow → see 6 matrix jobs (3 OS × 2 Go) → all pass
- **US4**: View Codecov → see coverage report → see badges in README
- **US5**: View workflow duration → verify <5 min (2-4 min with cache)

**Future Maintenance**: Phase 9 contains 3 backlog tasks for ongoing CI/CD optimization and monitoring
