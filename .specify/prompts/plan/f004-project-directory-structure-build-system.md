# Implementation Planning Prompt: F004 - Project Directory Structure & Build System

## Feature Metadata
- **Feature ID**: F004
- **Name**: Project Directory Structure & Build System
- **Feature Branch**: `002-f004-project-directory`
- **Category**: Foundation
- **Phase**: Week 3
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (1 day)
- **Dependencies**: F003 (Git repository must exist)
- **Spec Location**: `.specify/prompts/specify/mvp/f004-project-directory-structure-build-system.md`

## Constitutional Alignment

### Core Principles Verification
- ✅ **Verticalized > Generic**: N/A (infrastructure setup)
- ✅ **Speed > Features**: Optimized build flags, parallel compilation, < 2 minute build time
- ✅ **Local-First, Cloud Optional**: Build system works 100% offline
- ✅ **Boring Tech Wins**: Standard Go project layout, Makefile (proven, boring, standard)
- ✅ **Open Source Forever**: N/A (no licensing changes)
- ✅ **Developer-First Design**: Clear directory structure, predictable file locations, simple `make` commands
- ✅ **Ship Fast, Validate Early**: Simple structure enables rapid development, no over-engineering

### Technical Constraints Verification
- ✅ **Performance**: Build time < 2 minutes for single platform, binary size < 50MB compressed
- ✅ **Distribution**: Builds support all distribution channels (npm, homebrew, Docker)
- ✅ **Database Support**: N/A (directory structure only)
- ✅ **Cost**: $0 (local build tooling)
- ✅ **Code Quality**: Standard Go layout supports `go vet`, `golangci-lint`, TDD workflow
- ✅ **License**: N/A (no changes)
- ✅ **Platform Support**: Cross-compilation for macOS (Intel + ARM), Linux (x86 + ARM), Windows

### Legal Constraints Verification (CRITICAL)
- ✅ **Independent Development**: Build system uses standard public tools only
- ✅ **No Employer References**: N/A (infrastructure)
- ✅ **Public Information Only**: All patterns from public golang-standards/project-layout
- ✅ **Open Source Protection**: N/A (no licensing changes)
- ✅ **Illustrative Examples Only**: N/A (no company references)

## Planning Context

### Feature Summary
Establish a well-organized Go project structure following community best practices with directories for CLI commands (`/cmd`), packages (`/pkg`), schemas (`/schemas`), Docker files (`/docker`), documentation (`/docs`), and examples (`/examples`). Implement a Makefile with targets for building, testing, installing, and cross-compiling. Support 5 platforms with optimized, compressed binaries.

### Key Technical Decisions Required

**Phase 0 Research Topics**:
1. **Go Project Layout**: What is the standard golang-standards/project-layout? Which directories are essential vs. optional?
2. **Makefile Best Practices**: What are common targets for Go CLI projects? What variables and patterns?
3. **Cross-Compilation Configuration**: What GOOS/GOARCH values for 5 platforms? How to organize output?
4. **Binary Optimization**: What ldflags reduce binary size? How effective is compression?
5. **Version Injection**: How to inject version at compile time? Format for `--version` flag?
6. **Build Directory Structure**: How to organize /dist for multi-platform builds?
7. **Platform-Specific Handling**: Any special considerations for Windows (.exe), ARM architectures?
8. **Build Speed**: Parallel compilation strategies? Caching approaches?

### Technical Context (Pre-filled)

**Language/Version**: Go 1.21+ (specified in spec and constitution)
**Primary Dependencies**: None yet (foundational feature, dependencies added in F009)
**Storage**: File system (directory structure, build artifacts in /dist)
**Testing**: Placeholder test to verify build system works
**Target Platform**: Cross-platform (macOS Intel + ARM, Linux x86 + ARM, Windows x86)
**Project Type**: Single project (CLI tool)
**Performance Goals**:
  - Build time < 2 minutes per platform
  - Binary size < 50MB compressed
  - Cross-compilation for all 5 platforms
**Constraints**:
  - Must use standard Go project layout (golang-standards/project-layout)
  - Must support TDD workflow (directory structure for tests)
  - Must be boring and predictable (no exotic build systems)
  - Build artifacts in /dist, not root directory
  - Version injection via ldflags
**Scale/Scope**: Foundation for all subsequent development (weeks 4-12)

## Planning Workflow

### Phase 0: Research & Technical Decisions

Generate `research.md` with documented decisions for:

#### 1. Go Project Layout Standards
- **Decision Point**: Which directories are required? What goes in /cmd vs /pkg vs /internal?
- **Research**: Study golang-standards/project-layout and popular Go CLI projects (cobra, hugo, terraform)
- **Output**: Final directory structure with purpose/contents for each directory

#### 2. Makefile Structure
- **Decision Point**: What targets are essential? What variables? What PHONY declarations?
- **Research**: Common patterns from Go CLI projects
- **Output**: Makefile outline with:
  - Variables: BINARY_NAME, VERSION, BUILD_DIR, PLATFORMS
  - Targets: build, test, install, build-all, clean, help
  - PHONY declarations

#### 3. Cross-Compilation Configuration
- **Decision Point**: Exact GOOS/GOARCH values for 5 platforms? Output naming convention?
- **Research**: Go cross-compilation documentation, platform naming standards
- **Output**: Platform matrix:
  - darwin/amd64 (macOS Intel)
  - darwin/arm64 (macOS Apple Silicon)
  - linux/amd64 (Linux x86_64)
  - linux/arm64 (Linux ARM64)
  - windows/amd64 (Windows x86_64)

#### 4. Binary Optimization Strategy
- **Decision Point**: Which ldflags reduce size most? Compression algorithm? Target size?
- **Research**: Go compiler optimization flags, compression effectiveness
- **Output**:
  - ldflags: `-s -w` (strip debug symbols and symbol table)
  - Compression: gzip (standard, widely supported)
  - Target: < 50MB compressed per binary

#### 5. Version Injection Mechanism
- **Decision Point**: How to embed version? Format from git? Fallback for dirty builds?
- **Research**: Go ldflags -X flag, git describe patterns
- **Output**:
  - Version extraction: `git describe --tags --always --dirty`
  - Injection: `-ldflags="-X main.version=$(VERSION)"`
  - Display: `sourcebox --version` shows git-derived version

#### 6. Build Directory Organization
- **Decision Point**: How to organize /dist for multiple platforms? Naming convention?
- **Research**: Common patterns for multi-platform distribution
- **Output**:
  ```
  dist/
  ├── sourcebox-darwin-amd64.gz
  ├── sourcebox-darwin-arm64.gz
  ├── sourcebox-linux-amd64.gz
  ├── sourcebox-linux-arm64.gz
  └── sourcebox-windows-amd64.exe.gz
  ```

#### 7. Platform-Specific Considerations
- **Decision Point**: .exe extension for Windows? ARM64 compatibility testing approach?
- **Research**: Go Windows build conventions, ARM cross-compilation validation
- **Output**:
  - Windows: Add .exe extension automatically
  - ARM: Note that ARM binaries need testing on ARM hardware (deferred to manual QA)

#### 8. Build Performance Optimization
- **Decision Point**: Can builds run in parallel? Caching strategies?
- **Research**: Make parallel execution, Go build cache
- **Output**:
  - Parallel builds: Use Make's parallel flag for build-all
  - Caching: Rely on Go's built-in build cache
  - Target: Single platform < 2 min, all platforms < 10 min

**Deliverable**: `specs/002-f004-project-directory/research.md`

### Phase 1: Design & Contracts

#### 1. Data Model (SKIP for this feature)
**Rationale**: F004 is directory structure and build system setup. No data entities, no models, no database schema. Skip data-model.md generation.

#### 2. API Contracts (SKIP for this feature)
**Rationale**: F004 has no APIs, no endpoints, no services. CLI commands come later in F006. Skip contracts/ directory generation.

#### 3. Quickstart Guide (REQUIRED)
Generate `quickstart.md` with:

```markdown
# F004 Quickstart: Directory Structure & Build Verification

## Prerequisites
- F003 completed (Git repository and Go module initialized)
- Go 1.21+ installed
- Make installed (standard on macOS/Linux, installable on Windows)

## Directory Structure Overview
```
sourcebox/
├── cmd/sourcebox/        # CLI entry point (main.go)
├── pkg/                  # Internal packages
│   ├── generators/       # Data generation logic (F011-F020)
│   ├── schema/          # Schema parsing (F008)
│   └── database/        # Database connectors (F023-F024)
├── schemas/             # Schema JSON definitions (F007)
├── docker/              # Dockerfiles (F031-F036)
├── docs/                # Documentation (F037)
├── examples/            # Usage examples
├── dist/                # Build artifacts (gitignored)
├── Makefile            # Build automation
└── [root files from F003]
```

## Build System Verification

### 1. Build for Current Platform
```bash
make build
# Expected: Binary created in dist/sourcebox
```

### 2. Run Placeholder Test
```bash
make test
# Expected: Placeholder test passes, confirms build system works
```

### 3. Install Locally
```bash
make install
# Expected: Binary installed to $GOPATH/bin
```

### 4. Cross-Compile All Platforms
```bash
make build-all
# Expected: 5 compressed binaries in dist/
# darwin-amd64, darwin-arm64, linux-amd64, linux-arm64, windows-amd64
```

### 5. Verify Binary Sizes
```bash
ls -lh dist/
# Expected: All .gz files < 50MB
```

### 6. Test Binary Execution
```bash
./dist/sourcebox
# Expected: Binary runs (may show help or version)
```

## Makefile Targets Reference

- `make build` - Build for current platform → dist/sourcebox
- `make test` - Run all tests with coverage
- `make install` - Install binary to $GOPATH/bin
- `make build-all` - Cross-compile for all 5 platforms
- `make clean` - Remove dist/ and coverage files
- `make help` - Show available targets

## Verification Checklist
- [ ] All directories created (/cmd, /pkg, /schemas, /docker, /docs, /examples)
- [ ] Makefile present with all targets
- [ ] `make build` produces working binary
- [ ] `make test` runs successfully
- [ ] `make build-all` creates 5 platform binaries
- [ ] All compressed binaries < 50MB
- [ ] Directory structure documented in README or docs/

## Next Steps
- F006: Implement Cobra CLI framework in /cmd/sourcebox
- F008: Create schema parser in /pkg/schema
- F011-F020: Build data generators in /pkg/generators
```

**Deliverable**: `specs/002-f004-project-directory/quickstart.md`

#### 4. Update Agent Context
Run: `.specify/scripts/bash/update-agent-context.sh claude`

This updates the Claude-specific context file (CLAUDE.md) with:
- Project structure overview
- Makefile targets and usage
- Build system capabilities
- Cross-compilation platforms

**Deliverable**: Updated `CLAUDE.md` with build system information

## Constitution Re-verification

After Phase 1 design, verify:
- [ ] Directory structure follows standard Go project layout (Boring Tech principle)
- [ ] Build targets are simple and predictable (Developer-First principle)
- [ ] Binary optimization achieves < 50MB compressed (Performance constraint)
- [ ] Cross-compilation supports all 5 platforms (Platform Support constraint)
- [ ] Build time optimizations documented (Speed > Features principle)
- [ ] No over-engineering or unnecessary complexity (Ship Fast principle)
- [ ] Structure supports TDD workflow (Development Practice 1)

## Deliverables Summary

**Generated by /speckit.plan**:
1. ✅ `specs/002-f004-project-directory/plan.md` - This file
2. ✅ `specs/002-f004-project-directory/research.md` - Phase 0 output (8 research decisions)
3. ✅ `specs/002-f004-project-directory/quickstart.md` - Phase 1 output (build verification guide)
4. ⏭️ `specs/002-f004-project-directory/data-model.md` - SKIP (N/A for infrastructure)
5. ⏭️ `specs/002-f004-project-directory/contracts/` - SKIP (N/A for build system)
6. ✅ Updated CLAUDE.md with build system context

**NOT generated by /speckit.plan** (created later by /speckit.tasks):
- `specs/002-f004-project-directory/tasks.md` - Phase 2, separate command

## Success Criteria for Planning Phase

- ✅ All 8 research decisions documented with clear rationale
- ✅ Directory structure aligns with golang-standards/project-layout
- ✅ Makefile design covers all essential targets
- ✅ Cross-compilation configuration verified for 5 platforms
- ✅ Binary optimization strategy achieves < 50MB compressed
- ✅ Quickstart provides clear build verification steps
- ✅ Constitution compliance verified (no violations)
- ✅ Agent context updated with build system capabilities
- ✅ Planning artifacts reference constitution and spec correctly

## Anti-Patterns to Avoid

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

## Implementation Notes

### For the AI Agent
When executing `/speckit.plan` with this prompt:

1. **Start with comprehensive research.md**: Document all 8 research decisions with:
   - Decision: What was chosen
   - Rationale: Why chosen (reference constitution/spec)
   - Alternatives considered: What else was evaluated
   - Source: Where information came from (public docs, standards)

2. **Be explicit about skips**: Clearly state why data-model.md and contracts/ are not needed for infrastructure setup

3. **Focus on quickstart.md**: This is the primary deliverable beyond research. Include:
   - Directory structure overview with descriptions
   - Step-by-step build verification
   - Makefile targets reference
   - Verification checklist
   - Next steps (F006, F008, F011-F020)

4. **Verify constitutional compliance**:
   - Boring Tech: Standard Go layout, standard Makefile
   - Speed: Build time optimizations, parallel builds
   - Platform Support: All 5 platforms verified
   - Developer-First: Clear structure, simple commands
   - No complexity violations

5. **Keep it standard**: This is intentionally boring infrastructure:
   - Use proven patterns from golang-standards/project-layout
   - Follow common Makefile conventions
   - No exotic tools or build systems
   - Predictable and maintainable

### Directory Organization Guidelines

**What goes in /cmd/sourcebox/**:
- main.go only
- Minimal logic (bootstrap Cobra, exit)
- Delegates to /pkg for all business logic

**What goes in /pkg/**:
- All business logic
- Organized by domain (/generators, /schema, /database)
- Each subdirectory is a self-contained package

**What goes in /schemas/**:
- JSON schema definitions (added in F007)
- Schema documentation
- No code, only data

**What goes in /docker/**:
- Dockerfiles for each database
- Docker build/test scripts
- docker-compose.yml
- No Go code

**What goes in /docs/**:
- Docusaurus site source (F037)
- Architecture documentation
- API documentation
- No implementation code

**What goes in /examples/**:
- Quickstart examples
- Integration examples
- Fully working sample code

### Makefile Design Principles

1. **PHONY targets**: Declare all targets as .PHONY to avoid file conflicts
2. **Variables**: Define BINARY_NAME, VERSION, BUILD_DIR, PLATFORMS at top
3. **Default target**: Set .DEFAULT_GOAL := help
4. **Version extraction**: Use `git describe --tags --always --dirty`
5. **Parallel builds**: Design build-all to support `make -j4 build-all`
6. **Clean target**: Remove dist/ and coverage files, nothing else
7. **Help target**: Show all available targets with descriptions

### Cross-Compilation Notes

**Platform Matrix**:
- darwin/amd64 → sourcebox-darwin-amd64.gz
- darwin/arm64 → sourcebox-darwin-arm64.gz
- linux/amd64 → sourcebox-linux-amd64.gz
- linux/arm64 → sourcebox-linux-arm64.gz
- windows/amd64 → sourcebox-windows-amd64.exe.gz

**Binary Naming**: Include platform in filename for clarity

**Compression**: gzip all binaries to meet < 50MB requirement

**Testing Strategy**: Manual QA on real hardware (macOS, Linux, Windows) required before release

### Related Constitution Sections
- **Core Principle IV**: Boring Tech Wins (standard Go layout, Makefile)
- **Core Principle VI**: Developer-First Design (clear structure, predictable organization)
- **Core Principle II**: Speed > Features (optimized builds, parallel compilation)
- **Technical Constraint 7**: Platform Support (macOS, Linux, Windows cross-compilation)
- **Technical Constraint 1**: Performance (< 50MB binary size, reasonable build times)
- **Development Practice 6**: Spec-Kit workflow (this planning phase)
- **Anti-Pattern 4**: Shiny Tech (reject exotic build systems)
- **Anti-Pattern 5**: Over-Engineering (keep structure simple, expand as needed)

## Drag-and-Drop Usage

**To use this prompt**:
1. Drag this file into Claude Code
2. Claude will execute the `/speckit.plan` workflow for F004
3. Expected outputs:
   - research.md with 8 documented research decisions
   - quickstart.md with directory overview and build verification steps
   - Updated CLAUDE.md with build system context
   - Constitution compliance verified

**Estimated time**: 10-15 minutes for complete planning phase

**Next command**: `/speckit.tasks` to generate tasks.md from this plan
