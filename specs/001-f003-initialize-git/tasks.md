# Tasks: F003 - Initialize Git Repository & Go Module

**Input**: Design documents from `/specs/001-f003-initialize-git/`
**Prerequisites**: plan.md (✅), spec.md (✅), research.md (✅), quickstart.md (✅)
**Data Model**: N/A (configuration feature)
**Contracts**: N/A (no APIs)

**Tests**: Not applicable - This is a configuration/documentation feature with manual verification only. No executable logic to test.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1, US2, US3, US4)
- Include exact file paths in descriptions

## Path Conventions
F003 creates foundational configuration files at repository root. No source code directories yet (those come in F004).

---

## Phase 1: Setup (Repository Initialization)

**Purpose**: Initialize Git repository with basic structure

- [X] T001 [Setup] Create GitHub repository at github.com/jbeausoleil/sourcebox (public, empty, no .gitignore or license yet)
- [X] T002 [Setup] Clone empty repository to local machine
- [X] T003 [Setup] Create initial branch `001-f003-initialize-git` from main/master

---

## Phase 2: Foundational (Core Configuration Files)

**Purpose**: Create the essential configuration files that enable all user stories

**⚠️ CRITICAL**: All user stories depend on these files existing

- [X] T004 [P] [Foundation] Create .gitignore file with Go patterns per research.md Decision 1
- [X] T005 [P] [Foundation] Create LICENSE file with MIT license text per research.md Decision 3
- [X] T006 [P] [Foundation] Initialize Go module: `go mod init github.com/jbeausoleil/sourcebox` per research.md Decision 2
- [X] T007 [Foundation] Verify go.mod contains `go 1.25` directive
- [X] T008 [Foundation] Create empty go.sum file (will populate when dependencies added in F009)

**Checkpoint**: Core configuration complete - documentation can now be added

---

## Phase 3: User Story 1 - Clone and Understand Project (Priority: P1) 🎯 MVP

**Goal**: Developer can discover SourceBox on GitHub and immediately understand what it is, how it's licensed, and how to contribute

**Independent Test**:
1. View repository on GitHub (should see README with project description and legal notice)
2. Read LICENSE file (should see MIT license)
3. Read CONTRIBUTING.md (should understand bug/feature/PR process)
4. Read CODE_OF_CONDUCT.md (should understand community standards)
5. Answer: "What is SourceBox?", "Can I use it commercially?", "How do I contribute?"

### Implementation for User Story 1

- [X] T009 [P] [US1] Create README.md with structure per research.md Decision 3
  - **File**: `README.md`
  - **Content**:
    - Project name and tagline: "SourceBox - Verticalized demo data for developers in 30 seconds"
    - Badges: Build Status (placeholder), License (MIT), Go Version (1.21+)
    - Legal independence notice per research.md Decision 4: "SourceBox is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information."
    - Problem statement from spec.md (why SourceBox exists)
    - Features: Brief bulleted list (placeholder for future features)
    - Installation: Placeholder section (to be filled in F006-F013)
    - Quick start: Placeholder section (to be filled in F006-F013)
    - Contributing: Link to CONTRIBUTING.md
    - License: MIT with link to LICENSE file

- [X] T010 [P] [US1] Create CONTRIBUTING.md per research.md Decision 5
  - **File**: `CONTRIBUTING.md`
  - **Content**:
    - Bug Reports section: Template for reporting bugs (description, steps to reproduce, environment, logs)
    - Feature Requests section: Focus on problem, not solution (problem statement, workaround, benefit)
    - Pull Requests section: Process (fork/branch, code standards, commit format, testing)
    - Code Style section: gofmt, go vet, TDD for core logic, comments for non-obvious logic
    - Testing Requirements section: TDD for core, >80% coverage, manual QA on Mac/Linux/Windows

- [X] T011 [P] [US1] Create CODE_OF_CONDUCT.md per research.md Decision 6
  - **File**: `CODE_OF_CONDUCT.md`
  - **Content**: Full Contributor Covenant v2.1 text from contributor-covenant.org
  - **Customizations**: Project name (SourceBox), contact email (to be determined), scope (all project spaces)

- [X] T012 [US1] Add .gitignore patterns for documentation tools
  - **File**: `.gitignore` (append to existing)
  - **Content**: Add patterns for Markdown editors (*.md.backup, .obsidian/, etc.)

**Checkpoint**: User Story 1 complete - Developer can understand project from GitHub repo page

**Verification** (Manual QA from quickstart.md):
- [ ] View repository on GitHub - README renders correctly with legal notice visible
- [ ] LICENSE file contains full MIT license text
- [ ] CONTRIBUTING.md has clear bug/feature/PR instructions
- [ ] CODE_OF_CONDUCT.md contains Contributor Covenant v2.1
- [ ] All documentation files render correctly on GitHub (no broken markdown)

---

## Phase 4: User Story 2 - Set Up Local Development Environment (Priority: P1)

**Goal**: Developer can clone repository and verify Go module is properly configured

**Independent Test**:
1. Clone repository: `git clone https://github.com/jbeausoleil/sourcebox.git`
2. Run `go mod download` (should complete without errors)
3. Run `go list -m` (should output: github.com/jbeausoleil/sourcebox)
4. Verify go.mod contains `go 1.21`

### Implementation for User Story 2

User Story 2 has no additional implementation tasks - it depends entirely on Phase 2 (Foundational) tasks being complete.

**Verification Prerequisites**:
- [x] T006: Go module initialized
- [x] T007: Go 1.21 directive present
- [x] T001-T003: Repository exists and is cloneable

**Checkpoint**: User Story 2 complete - Developer can set up local environment

**Verification** (Manual QA from quickstart.md):
- [ ] Clone repository completes in < 30 seconds
- [ ] `go mod download` runs without errors
- [ ] `go list -m` outputs correct module path
- [ ] Repository size is < 1MB (verify with `du -sh .git`)

---

## Phase 5: User Story 3 - Understand Contribution Guidelines (Priority: P2)

**Goal**: Developer understands the process for contributing bugs, features, and code

**Independent Test**:
1. Read CONTRIBUTING.md
2. Answer: "How do I report a bug?", "How do I request a feature?", "What's the PR process?"

### Implementation for User Story 3

User Story 3 has no additional implementation tasks - it depends entirely on T010 (CONTRIBUTING.md) being complete.

**Verification Prerequisites**:
- [x] T010: CONTRIBUTING.md created with clear instructions

**Checkpoint**: User Story 3 complete - Developer knows how to contribute

**Verification** (Manual QA from spec.md):
- [ ] Developer can identify where to report bugs (GitHub issues)
- [ ] Developer knows what information to include in bug reports
- [ ] Developer knows what information to include in feature requests
- [ ] Developer understands PR process (fork, branch, standards, review)

---

## Phase 6: User Story 4 - Understand Community Standards (Priority: P3)

**Goal**: Developer understands expected behavior and values of SourceBox community

**Independent Test**:
1. Read CODE_OF_CONDUCT.md
2. Answer: "What behaviors are expected?", "What happens if I witness inappropriate behavior?"

### Implementation for User Story 4

User Story 4 has no additional implementation tasks - it depends entirely on T011 (CODE_OF_CONDUCT.md) being complete.

**Verification Prerequisites**:
- [x] T011: CODE_OF_CONDUCT.md created with Contributor Covenant v2.1

**Checkpoint**: User Story 4 complete - Developer understands community standards

**Verification** (Manual QA from spec.md):
- [ ] Developer can identify acceptable behaviors
- [ ] Developer can identify unacceptable behaviors
- [ ] Developer knows how to report inappropriate behavior
- [ ] Developer understands enforcement mechanisms

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Final verification and quality checks across all user stories

- [ ] T013 [P] [Polish] Verify all documentation links work (README → CONTRIBUTING, LICENSE)
- [ ] T014 [P] [Polish] Verify all badges render correctly on GitHub (build status placeholder, license, Go version)
- [ ] T015 [Polish] Run full quickstart.md verification checklist
  - **Location**: specs/001-f003-initialize-git/quickstart.md
  - **Checklist**: All 24 items in "Verification Checklist" section
- [ ] T016 [Polish] Verify repository size < 1MB per success criteria SC-002
- [ ] T017 [Polish] Verify legal notice is prominently visible in README per success criteria SC-005
- [ ] T018 [P] [Polish] Search for any accidental employer references: `grep -r "employer-name" .`
- [ ] T019 [P] [Polish] Search for any accidentally committed secrets: `git log --all --full-history -- .env`
- [ ] T020 [Polish] Verify all success criteria from spec.md (SC-001 through SC-007)
- [ ] T021 [Polish] Commit all changes with message:
  ```
  feat: Initialize Git repository and Go module (F003)

  - Add MIT license
  - Add README with legal independence notice
  - Add CONTRIBUTING.md and CODE_OF_CONDUCT.md
  - Initialize Go module (github.com/jbeausoleil/sourcebox)
  - Configure .gitignore for Go projects

  Implements:
  - FR-001 through FR-014 from spec.md
  - All decisions from research.md
  - Enables all user stories (US1-US4)

  🤖 Generated with [Claude Code](https://claude.com/claude-code)

  Co-Authored-By: Claude <noreply@anthropic.com>
  ```
- [ ] T022 [Polish] Push branch to remote: `git push -u origin 001-f003-initialize-git`

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup (T001-T003) completion
- **User Story 1 (Phase 3)**: Depends on Foundational (T004-T008) completion
- **User Story 2 (Phase 4)**: Depends on Foundational (T004-T008) completion - No additional implementation
- **User Story 3 (Phase 5)**: Depends on User Story 1 (T010) completion - No additional implementation
- **User Story 4 (Phase 6)**: Depends on User Story 1 (T011) completion - No additional implementation
- **Polish (Phase 7)**: Depends on all user stories (T009-T012) being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - Creates all documentation files
- **User Story 2 (P1)**: Can start after Foundational (Phase 2) - Only verification, no new tasks
- **User Story 3 (P2)**: Depends on US1 (T010: CONTRIBUTING.md) - Only verification, no new tasks
- **User Story 4 (P3)**: Depends on US1 (T011: CODE_OF_CONDUCT.md) - Only verification, no new tasks

**Note**: US2, US3, and US4 have no implementation tasks - they are verification-only stories that depend on foundational files created in Phase 2 and Phase 3.

### Within Each Phase

**Phase 1 (Setup)**:
- T001 → T002 → T003 (sequential: must create repo before cloning before creating branch)

**Phase 2 (Foundational)**:
- T004, T005, T006 can run in parallel [P]
- T007 depends on T006 (must initialize module before verifying version)
- T008 depends on T007 (create go.sum after go.mod is correct)

**Phase 3 (User Story 1)**:
- T009, T010, T011 can run in parallel [P] (different files, no dependencies)
- T012 depends on T009-T011 (add documentation patterns after docs exist)

**Phase 7 (Polish)**:
- T013, T014, T016, T017, T018, T019 can run in parallel [P]
- T015, T020 must run after parallel tasks (comprehensive verification)
- T021 must run after all verification (commit only after quality checks)
- T022 must run after T021 (push only after commit)

### Parallel Opportunities

**Phase 2 Foundational** - Launch together:
```bash
Task: "Create .gitignore file" (T004)
Task: "Create LICENSE file" (T005)
Task: "Initialize Go module" (T006)
```

**Phase 3 User Story 1** - Launch together:
```bash
Task: "Create README.md" (T009)
Task: "Create CONTRIBUTING.md" (T010)
Task: "Create CODE_OF_CONDUCT.md" (T011)
```

**Phase 7 Polish** - Launch together:
```bash
Task: "Verify documentation links" (T013)
Task: "Verify badges render" (T014)
Task: "Verify repository size" (T016)
Task: "Verify legal notice visible" (T017)
Task: "Search for employer references" (T018)
Task: "Search for secrets" (T019)
```

---

## Implementation Strategy

### MVP First (User Stories 1 & 2 Only - Both P1)

1. **Complete Phase 1**: Setup (T001-T003)
   - Duration: ~5 minutes
   - Result: Empty repository cloned locally on branch

2. **Complete Phase 2**: Foundational (T004-T008)
   - Duration: ~10 minutes
   - Result: Core configuration files (.gitignore, LICENSE, go.mod) exist

3. **Complete Phase 3**: User Story 1 (T009-T012)
   - Duration: ~30 minutes
   - Result: All documentation files (README, CONTRIBUTING, CODE_OF_CONDUCT) exist

4. **Verify Phase 4**: User Story 2 (verification only)
   - Duration: ~5 minutes
   - Result: Confirm Go module works correctly

5. **STOP and VALIDATE**: Can developers clone and understand the project?
   - Clone repository
   - Read README (understand project, see legal notice)
   - Run `go mod download` (verify Go setup works)
   - Read CONTRIBUTING.md and CODE_OF_CONDUCT.md

6. **Complete Phase 7**: Polish (T013-T022)
   - Duration: ~15 minutes
   - Result: Quality verified, committed, pushed

**Total MVP Time**: ~65 minutes (1 hour)

**MVP Scope**: User Stories 1 & 2 (both P1) - Developer can discover, understand, and set up SourceBox

### Incremental Delivery (All User Stories)

1. Complete Setup + Foundational → Files exist
2. Add User Story 1 → Documentation complete → Verify independently
3. Verify User Story 2 → Go module works → Verify independently
4. Verify User Story 3 → Contribution guidelines understood → Verify independently
5. Verify User Story 4 → Community standards understood → Verify independently
6. Polish → All quality checks pass → Commit and push

**Total Time for All Stories**: ~75 minutes (1.25 hours)

### Solo Developer Strategy

For one developer working alone:

1. **Day 1, Hour 1** (Complete F003):
   - Complete all phases sequentially (T001-T022)
   - Run verification checklist from quickstart.md
   - Push to GitHub
   - Verify on GitHub web interface that everything renders correctly

2. **Immediate Next Steps**:
   - F004: Project Directory Structure
   - F005: GitHub Actions CI/CD (will make build badge functional)

---

## Implementation Notes

### TDD Not Applicable
This feature creates configuration and documentation files. There is no executable logic to test. Manual verification using quickstart.md checklist is the appropriate quality assurance approach.

### Manual QA Required
Per Constitution Development Practice 3, manually verify:
- Repository clones successfully
- Go module initializes correctly
- All documentation renders on GitHub
- Legal notice is visible
- No employer references present
- No secrets committed

### File Count & Size
- **Total files created**: 6 files (.gitignore, LICENSE, README.md, CONTRIBUTING.md, CODE_OF_CONDUCT.md, go.mod)
- **Expected repository size**: < 100KB (well under 1MB constraint)
- **Clone time**: < 10 seconds on typical broadband

### Success Criteria Mapping
- **SC-001**: T015 verification (< 2 minutes clone + setup)
- **SC-002**: T016 verification (< 1MB repository size)
- **SC-003**: T013, T014 verification (documentation renders correctly)
- **SC-004**: T015 verification (answer 3 key questions)
- **SC-005**: T017 verification (legal notice visible)
- **SC-006**: T018 verification (no proprietary info or employer references)
- **SC-007**: T015 verification (Go module path resolves correctly)

### Constitutional Compliance
- ✅ **Boring Tech**: Git + Go module (standard, proven)
- ✅ **Open Source Forever**: MIT license in T005
- ✅ **Speed**: All tasks complete in < 2 hours
- ✅ **Legal Independence**: Legal notice in T009, verification in T017, T018
- ✅ **Developer-First**: Standard Go conventions, clear documentation

---

## Task Summary

**Total Tasks**: 22 tasks (T001-T022)

### Tasks per Phase
- **Phase 1 (Setup)**: 3 tasks (T001-T003)
- **Phase 2 (Foundational)**: 5 tasks (T004-T008)
- **Phase 3 (User Story 1 - P1)**: 4 tasks (T009-T012) 🎯
- **Phase 4 (User Story 2 - P1)**: 0 tasks (verification only) 🎯
- **Phase 5 (User Story 3 - P2)**: 0 tasks (verification only)
- **Phase 6 (User Story 4 - P3)**: 0 tasks (verification only)
- **Phase 7 (Polish)**: 10 tasks (T013-T022)

### Tasks per User Story
- **US1 (P1) - Clone and Understand**: 4 implementation tasks (T009-T012)
- **US2 (P1) - Set Up Local Dev**: 0 implementation tasks (uses Phase 2 foundational tasks)
- **US3 (P2) - Understand Contribution**: 0 implementation tasks (uses T010 from US1)
- **US4 (P3) - Understand Community**: 0 implementation tasks (uses T011 from US1)

### Parallel Opportunities
- **3 parallel tasks** in Phase 2: T004, T005, T006
- **3 parallel tasks** in Phase 3: T009, T010, T011
- **6 parallel tasks** in Phase 7: T013, T014, T016, T017, T018, T019

**Total Parallelizable Tasks**: 12 tasks (55% of all tasks)

### MVP Scope
**MVP = User Stories 1 & 2 (both P1)**
- Phase 1: Setup (3 tasks)
- Phase 2: Foundational (5 tasks)
- Phase 3: User Story 1 (4 tasks)
- Phase 4: User Story 2 (verification)
- Phase 7: Polish (10 tasks)

**MVP Total**: 22 tasks (100% of feature - this is a small foundational feature)

---

## Next Steps

1. Execute tasks in order (T001 → T022)
2. Mark tasks complete as you go
3. Verify each user story independently using quickstart.md
4. Commit and push when all verification passes
5. Move to F004 (Project Directory Structure & Build System)

**Ready to implement!** 🚀
