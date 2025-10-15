# Implementation Plan: F004 - Project Directory Structure & Build System

**Branch**: `002-f004-project-directory` | **Date**: 2025-10-14 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/002-f004-project-directory/spec.md`

## Summary

Establish a well-organized Go project structure following golang-standards/project-layout best practices with directories for CLI commands (`/cmd`), packages (`/pkg`), schemas (`/schemas`), Docker files (`/docker`), documentation (`/docs`), and examples (`/examples`). Implement a Makefile with targets for building, testing, installing, and cross-compiling. Support 5 platforms (macOS Intel/ARM, Linux x86/ARM, Windows x86) with optimized, compressed binaries under 50MB.

**Technical Approach** (from research):
- Use standard Go project layout (boring, proven, widely documented)
- Build automation via Makefile (simple, predictable, universally available)
- Cross-compilation for 5 platforms using Go's built-in toolchain
- Binary optimization with `-ldflags="-s -w"` and gzip compression
- Version injection via `git describe --tags --always --dirty`
- Build artifacts in `/dist` directory (gitignored, easy cleanup)

## Technical Context

**Language/Version**: Go 1.21+ (specified in spec and constitution)
**Primary Dependencies**: None (standard library only for MVP foundation)
**Storage**: File system (directory structure, build artifacts in /dist)
**Testing**: go test (built-in Go testing framework, placeholder test for build verification)
**Target Platform**: Cross-platform (macOS Intel + ARM, Linux x86 + ARM, Windows x86)
**Project Type**: Single project (CLI tool with standard Go layout)
**Performance Goals**:
  - Build time < 30 seconds per platform (typical: 5-10s on 2020 MacBook Pro)
  - Cross-compilation < 2 minutes for all 5 platforms (with `make -j4 build-all`)
  - Binary size < 50MB compressed (constitutional requirement)
**Constraints**:
  - Must use standard Go project layout (golang-standards/project-layout)
  - Must support TDD workflow (directory structure for tests)
  - Must be boring and predictable (no exotic build systems)
  - Build artifacts in /dist, not root directory
  - Version injection via ldflags (no hardcoded versions)
**Scale/Scope**: Foundation for all subsequent development (weeks 4-12)

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

### Core Principles ✅

| Principle | Status | Verification |
|-----------|--------|-------------|
| **Verticalized > Generic** | ✅ PASS (N/A) | Infrastructure setup, no data generation |
| **Speed > Features** | ✅ PASS | Build time < 2 min, optimized flags, parallel compilation |
| **Local-First, Cloud Optional** | ✅ PASS | Build system works 100% offline, no network calls |
| **Boring Tech Wins** | ✅ PASS | Standard Go layout, Makefile, gzip (proven, standard) |
| **Open Source Forever** | ✅ PASS (N/A) | No licensing changes |
| **Developer-First Design** | ✅ PASS | Clear structure, predictable locations, simple `make` commands |
| **Ship Fast, Validate Early** | ✅ PASS | Simple structure enables rapid development, no over-engineering |

### Technical Constraints ✅

| Constraint | Status | Verification |
|-----------|--------|-------------|
| **Performance** | ✅ PASS | Build < 2 min, binary < 50MB compressed |
| **Distribution** | ✅ PASS | Builds support npm, homebrew, Docker |
| **Database Support** | ✅ PASS (N/A) | Directory structure only |
| **Cost** | ✅ PASS | $0 (local build tooling) |
| **Code Quality** | ✅ PASS | Standard layout supports vet, lint, TDD |
| **License** | ✅ PASS (N/A) | No changes |
| **Platform Support** | ✅ PASS | All 5 platforms (macOS Intel/ARM, Linux x86/ARM, Windows x86) |

### Development Practices ✅

| Practice | Status | Verification |
|-----------|--------|-------------|
| **TDD Required** | ✅ PASS | Placeholder test verifies build system works |
| **Manual QA** | ✅ PASS | Build verification documented in quickstart.md |
| **Spec-Kit Workflow** | ✅ PASS | Following /speckit.plan workflow (this document) |
| **12-Week MVP** | ✅ PASS | Simple structure, no delays, fits timeline |
| **Open Source First** | ✅ PASS (N/A) | No launch changes |
| **Indie Constraints** | ✅ PASS | 10-15 hours work, personal equipment, no employer resources |

### Legal Constraints ✅ (CRITICAL)

| Constraint | Status | Verification |
|-----------|--------|-------------|
| **Independent Development** | ✅ PASS | Standard public tools only (Go, Make, gzip) |
| **No Employer References** | ✅ PASS | No employer references (infrastructure) |
| **Public Information Only** | ✅ PASS | golang-standards/project-layout (public) |
| **Open Source Protection** | ✅ PASS (N/A) | No licensing changes |
| **Illustrative Examples Only** | ✅ PASS (N/A) | No company references |

### Anti-Patterns Avoided ✅

| Anti-Pattern | Status | Verification |
|-----------|--------|-------------|
| **Feature Bloat** | ✅ AVOIDED | Minimal Makefile, only essential targets |
| **Enterprise-First** | ✅ AVOIDED | Simple structure, no complex build system |
| **Complex Pricing** | ✅ AVOIDED (N/A) | Infrastructure only |
| **Shiny Tech** | ✅ AVOIDED | Make, Go, gzip (all boring and standard) |
| **Over-Engineering** | ✅ AVOIDED | Standard layout, no custom build tools |
| **Generic Data** | ✅ AVOIDED (N/A) | No data generation in this feature |
| **Premature Optimization** | ✅ AVOIDED | Standard Go cache, no custom caching |
| **Cloud-First** | ✅ AVOIDED | Build system 100% offline |

**GATE RESULT**: ✅ PASS - All constitutional requirements met, no violations to justify

## Project Structure

### Documentation (this feature)

```
specs/002-f004-project-directory/
├── spec.md              # Feature specification (input)
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command) ✅ COMPLETE
├── quickstart.md        # Phase 1 output (/speckit.plan command) ✅ COMPLETE
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)

Notes:
- data-model.md SKIPPED (no data entities in infrastructure setup)
- contracts/ SKIPPED (no APIs in this feature)
```

### Source Code (repository root)

```
sourcebox/
├── cmd/sourcebox/        # CLI entry point (main.go only)
│   ├── main.go          # Minimal bootstrap logic (delegates to /pkg)
│   └── main_test.go     # Placeholder test (verifies build system)
│
├── pkg/                  # Public library code (organized by domain)
│   ├── generators/       # Data generation logic (F011-F020)
│   ├── schema/          # Schema parsing (F008)
│   └── database/        # Database connectors (F023-F024)
│
├── schemas/             # Schema JSON definitions (F007)
├── docker/              # Dockerfiles for databases (F031-F036)
├── docs/                # Docusaurus documentation site (F037)
├── examples/            # Usage examples (F006+)
│
├── dist/                # Build artifacts (gitignored)
│   ├── sourcebox                      # Current platform (uncompressed)
│   ├── sourcebox-darwin-amd64.gz      # macOS Intel
│   ├── sourcebox-darwin-arm64.gz      # macOS Apple Silicon
│   ├── sourcebox-linux-amd64.gz       # Linux x86_64
│   ├── sourcebox-linux-arm64.gz       # Linux ARM64
│   └── sourcebox-windows-amd64.exe.gz # Windows x86_64
│
├── Makefile             # Build automation (6 targets: build, test, install, build-all, clean, help)
├── .gitignore           # Includes dist/, coverage.txt
│
└── [root files from F003]
    ├── go.mod           # Module: github.com/jbeausoleil/sourcebox
    ├── go.sum           # (generated after dependencies added)
    ├── LICENSE          # MIT license
    ├── README.md        # Project overview with legal notice
    ├── CONTRIBUTING.md  # Contribution guidelines
    └── CODE_OF_CONDUCT.md  # Contributor Covenant v2.1
```

**Structure Decision**: Standard Go single-project layout (golang-standards/project-layout)

**Rationale**:
- `/cmd/sourcebox/`: Entry point with minimal logic (standard Go convention)
- `/pkg/`: Public library code organized by domain (supports future external use)
- NOT using `/internal/` initially (add when proprietary code exists, not before)
- Domain-specific directories (`/schemas`, `/docker`, `/docs`, `/examples`) for clear organization
- `/dist` for build artifacts (keeps root clean, easy to gitignore)

## Complexity Tracking

*No constitutional violations - this section intentionally left empty*

All decisions align with Boring Tech Wins, Developer-First Design, and Ship Fast principles. Standard Go project layout, standard Makefile, standard tooling (no exotic build systems).

---

## Planning Phase Deliverables

### Phase 0: Research ✅ COMPLETE

**Deliverable**: `research.md` (8 technical decisions documented)

All research decisions completed and documented:
1. ✅ Go Project Layout Standards (golang-standards/project-layout)
2. ✅ Makefile Structure (build, test, install, build-all, clean, help)
3. ✅ Cross-Compilation Configuration (5 platforms with naming conventions)
4. ✅ Binary Optimization Strategy (-s -w flags, gzip compression, < 50MB)
5. ✅ Version Injection Mechanism (git describe + ldflags -X)
6. ✅ Build Directory Organization (/dist with platform in filename)
7. ✅ Platform-Specific Considerations (.exe for Windows, ARM manual QA)
8. ✅ Build Performance Optimization (Make -j4, Go build cache, < 2 min)

**Location**: `specs/002-f004-project-directory/research.md`

### Phase 1: Design & Contracts ✅ COMPLETE

**Deliverables**:
1. ✅ `quickstart.md` - Build verification guide with step-by-step instructions
2. ✅ Updated `CLAUDE.md` - Agent context updated with build system information
3. ⏭️ `data-model.md` - SKIPPED (no data entities in infrastructure setup)
4. ⏭️ `contracts/` - SKIPPED (no APIs in this feature)

**Locations**:
- `specs/002-f004-project-directory/quickstart.md`
- `CLAUDE.md` (repository root)

### Phase 2: Task Generation ⏭️ NOT IN SCOPE

**Note**: `/speckit.plan` does NOT generate `tasks.md`. This is created separately by the `/speckit.tasks` command after planning is complete.

---

## Post-Design Constitution Re-Verification ✅

After Phase 1 design, all constitutional requirements remain satisfied:

- ✅ Directory structure follows standard Go project layout (Boring Tech principle)
- ✅ Build targets are simple and predictable (Developer-First principle)
- ✅ Binary optimization achieves < 50MB compressed (Performance constraint)
- ✅ Cross-compilation supports all 5 platforms (Platform Support constraint)
- ✅ Build time optimizations documented (Speed > Features principle)
- ✅ No over-engineering or unnecessary complexity (Ship Fast principle)
- ✅ Structure supports TDD workflow (Development Practice 1)

**No new violations introduced during design phase.**

---

## Success Criteria Summary

Planning phase is complete when:

- ✅ All 8 research decisions documented with clear rationale
- ✅ Directory structure aligns with golang-standards/project-layout
- ✅ Makefile design covers all essential targets
- ✅ Cross-compilation configuration verified for 5 platforms
- ✅ Binary optimization strategy achieves < 50MB compressed
- ✅ Quickstart provides clear build verification steps
- ✅ Constitution compliance verified (no violations)
- ✅ Agent context updated with build system capabilities
- ✅ Planning artifacts reference constitution and spec correctly

---

## Next Steps

**Immediate** (after planning complete):
1. Run `/speckit.tasks` to generate dependency-ordered tasks.md
2. Review generated tasks for accuracy and completeness
3. Begin implementation with `/speckit.implement`

**Implementation** (F004 tasks):
1. Create directory structure (/cmd, /pkg, /schemas, /docker, /docs, /examples, /dist)
2. Create Makefile with 6 targets (build, test, install, build-all, clean, help)
3. Create cmd/sourcebox/main.go with minimal bootstrap code
4. Create cmd/sourcebox/main_test.go with placeholder test
5. Update .gitignore to exclude dist/ and coverage.txt
6. Verify builds work for current platform (`make build`)
7. Verify cross-compilation works (`make build-all`)
8. Verify version injection works (`./dist/sourcebox --version`)
9. Run verification steps from quickstart.md

**Follow-up Features**:
- F006: Implement Cobra CLI framework in /cmd/sourcebox
- F008: Create schema parser in /pkg/schema
- F011-F020: Build data generators in /pkg/generators

---

## Notes for Implementation

### What NOT to Generate
- ❌ Don't generate data-model.md (no data entities in this feature)
- ❌ Don't generate API contracts (no APIs in this feature)
- ❌ Don't use non-standard Go project layout (confuses developers)
- ❌ Don't over-engineer the Makefile (keep targets simple)
- ❌ Don't use exotic build tools (Make is boring and standard)
- ❌ Don't create /internal directory yet (not needed until proprietary code exists)
- ❌ Don't optimize prematurely (baseline build first, optimize if needed)
- ❌ Don't skip platform testing (document manual QA requirement)
- ❌ Don't exceed 50MB compressed binary size (violates constitution)
- ❌ Don't create complex build scripts (Makefile should be readable)

### Key Implementation Principles
1. **Start simple**: Minimal main.go, simple Makefile, standard layout
2. **Test the build system**: Placeholder test verifies tooling works
3. **Keep it boring**: Standard tools, standard patterns, standard names
4. **Document clearly**: Quickstart explains how to verify everything works
5. **Version automatically**: git describe ensures version is always current

---

**Planning Status**: ✅ Complete
**Next Command**: `/speckit.tasks` (to generate implementation tasks)
**Branch**: `002-f004-project-directory`
**Estimated Implementation Time**: 1 day (small effort, foundational feature)
