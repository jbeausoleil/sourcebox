# Feature Specification: Project Directory Structure & Build System

**Feature Branch**: `002-f004-project-directory`
**Created**: 2025-10-14
**Status**: Draft
**Input**: User description: "F004 - Project Directory Structure & Build System"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Developer Navigates Codebase (Priority: P1)

A developer new to the SourceBox project opens the repository and needs to quickly locate where to add schema parsing logic. They find a clear `/pkg/schema` directory with expected structure, allowing them to contribute within minutes of cloning.

**Why this priority**: Without clear organization, every developer task requires learning a custom structure. Standard layout enables immediate productivity and reduces onboarding friction.

**Independent Test**: Clone the repository, review the directory structure, and successfully locate the appropriate package directory for schema, database, or generator logic within 2 minutes.

**Acceptance Scenarios**:

1. **Given** a developer clones the repository, **When** they explore the project root, **Then** they see standard Go directories (cmd/, pkg/, etc.) with clear purposes
2. **Given** a developer needs to add schema logic, **When** they look for the schema package, **Then** they find /pkg/schema with clear organization
3. **Given** a developer wants to understand project organization, **When** they review README or docs/, **Then** they find directory structure documentation

---

### User Story 2 - Developer Builds Application (Priority: P1)

A developer wants to test their changes by building the application for their local platform (macOS, Linux, or Windows). They run a single `make build` command and receive a working binary in under 30 seconds.

**Why this priority**: Build capability is essential for any development work. Developers must be able to build and test their changes locally before committing.

**Independent Test**: Run `make build` from the project root and verify a working binary is created in the /dist directory that executes without errors.

**Acceptance Scenarios**:

1. **Given** a developer runs `make build`, **When** the build completes, **Then** a binary appears in /dist directory
2. **Given** a developer runs the built binary, **When** it executes, **Then** it runs without errors (even if minimal functionality)
3. **Given** a developer wants to see available commands, **When** they run `make help`, **Then** they see all available Makefile targets

---

### User Story 3 - Developer Runs Test Suite (Priority: P1)

A developer wants to verify their changes haven't broken existing functionality. They run `make test` and receive clear test results with coverage information.

**Why this priority**: Testing is fundamental to quality. Developers need immediate feedback on whether their changes maintain existing behavior.

**Independent Test**: Run `make test` and verify tests execute successfully with coverage reporting, even if only placeholder tests exist initially.

**Acceptance Scenarios**:

1. **Given** a developer runs `make test`, **When** tests complete, **Then** they see test results with pass/fail status
2. **Given** tests complete successfully, **When** coverage reports generate, **Then** coverage.txt file is created
3. **Given** a developer adds new code, **When** they run tests, **Then** race conditions are detected (via -race flag)

---

### User Story 4 - Developer Installs CLI Locally (Priority: P2)

A developer wants to use the SourceBox CLI from anywhere on their system. They run `make install` and can execute `sourcebox` commands from any directory.

**Why this priority**: Local installation enables dogfooding and integration testing. While important, it's not required for basic development.

**Independent Test**: Run `make install`, then open a new terminal and execute `sourcebox --version` from any directory to verify installation.

**Acceptance Scenarios**:

1. **Given** a developer runs `make install`, **When** installation completes, **Then** binary is placed in $GOPATH/bin
2. **Given** the binary is installed, **When** developer opens a new terminal, **Then** `sourcebox` command is available globally
3. **Given** the binary executes, **When** user runs `sourcebox --version`, **Then** the current git version is displayed

---

### User Story 5 - Maintainer Creates Release Binaries (Priority: P2)

A project maintainer wants to create distribution binaries for all supported platforms (macOS Intel/ARM, Linux x86_64/ARM64, Windows x86_64). They run `make build-all` and receive compressed binaries for each platform, all under 50MB.

**Why this priority**: Release capability is essential for distribution but not required for daily development. Can be added after core build system works.

**Independent Test**: Run `make build-all`, verify 5 compressed binaries (.gz files) are created in /dist, and confirm each is under 50MB.

**Acceptance Scenarios**:

1. **Given** a maintainer runs `make build-all`, **When** cross-compilation completes, **Then** binaries for 5 platforms are created
2. **Given** binaries are created, **When** they are compressed with gzip, **Then** each .gz file is under 50MB
3. **Given** platform-specific binaries exist, **When** tested on respective platforms, **Then** they execute correctly

---

### User Story 6 - Developer Cleans Build Artifacts (Priority: P3)

A developer wants to clean up build artifacts before committing or to free disk space. They run `make clean` and all generated files are removed.

**Why this priority**: Cleanup is a convenience feature that improves developer experience but isn't critical for core functionality.

**Independent Test**: Run `make build` to create artifacts, then run `make clean` and verify /dist directory and coverage files are removed.

**Acceptance Scenarios**:

1. **Given** build artifacts exist in /dist, **When** developer runs `make clean`, **Then** /dist directory is removed
2. **Given** coverage.txt exists, **When** developer runs `make clean`, **Then** coverage.txt is removed
3. **Given** the repository is clean, **When** developer runs `git status`, **Then** no build artifacts appear as untracked files

---

### Edge Cases

- What happens when a developer runs `make build` without Go installed?
- What happens when cross-compilation encounters missing platform tools?
- How does the system handle missing /dist directory on first build?
- What happens when git tags don't exist for version injection?
- How are build errors communicated clearly to developers?
- What happens when running `make build-all` on Windows (shell script compatibility)?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST provide a /cmd/sourcebox directory containing main.go and CLI entry point
- **FR-002**: System MUST provide /pkg directory with subdirectories for generators, schema, and database packages
- **FR-003**: System MUST provide /schemas directory for schema JSON definitions
- **FR-004**: System MUST provide /docker directory for container-related files
- **FR-005**: System MUST provide /docs directory for documentation source
- **FR-006**: System MUST provide /examples directory for usage examples
- **FR-007**: System MUST provide a Makefile with build, test, install, build-all, clean, and help targets
- **FR-008**: System MUST build binaries for current platform using `make build`
- **FR-009**: System MUST run all tests with race detection and coverage using `make test`
- **FR-010**: System MUST install binary to $GOPATH/bin using `make install`
- **FR-011**: System MUST cross-compile for 5 platforms (darwin/amd64, darwin/arm64, linux/amd64, linux/arm64, windows/amd64) using `make build-all`
- **FR-012**: System MUST compress all cross-compiled binaries with gzip
- **FR-013**: System MUST inject git version into binary at compile time using ldflags
- **FR-014**: System MUST strip debug symbols and symbol table to reduce binary size
- **FR-015**: System MUST output all build artifacts to /dist directory
- **FR-016**: System MUST remove all build artifacts when `make clean` is executed
- **FR-017**: System MUST display available Makefile targets when `make help` is executed
- **FR-018**: System MUST include placeholder test file in /cmd/sourcebox/main_test.go to verify build system

### Key Entities

- **Directory Structure**: Organized hierarchy of folders following golang-standards/project-layout conventions
  - /cmd: Entry points for applications
  - /pkg: Public library code that can be used by external applications
  - /schemas: Data schema definitions (JSON files)
  - /docker: Container configurations
  - /docs: Documentation files
  - /examples: Sample code and usage demonstrations

- **Build Artifacts**: Generated binary files for different platforms
  - Platform-specific binaries (uncompressed executables)
  - Compressed distribution files (.gz)
  - Test coverage reports (coverage.txt)

- **Build Configuration**: Variables and settings that control compilation
  - Binary name (sourcebox)
  - Version (derived from git)
  - Build flags (optimization, stripping)
  - Target platforms (OS/architecture combinations)

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Developers can locate the appropriate package directory for a given feature (schema, database, generator) within 2 minutes of viewing the repository
- **SC-002**: Build completes for current platform in under 30 seconds on standard development hardware
- **SC-003**: All compressed distribution binaries are under 50MB in size
- **SC-004**: Test suite executes successfully with coverage reporting enabled
- **SC-005**: Cross-compilation produces 5 working binaries (one per supported platform) in under 2 minutes
- **SC-006**: Developers can successfully install and execute the CLI from any directory after running `make install`
- **SC-007**: All build artifacts are removed after running `make clean` with no manual intervention required
- **SC-008**: New contributors can understand the project structure without external documentation beyond README

## Assumptions *(optional)*

- Developers have Go 1.21+ installed as specified in project requirements
- Developers are using Unix-like systems (macOS, Linux) or WSL on Windows for development
- Git is installed and repository has at least one commit for version generation
- Standard shell tools (gzip) are available on the build system
- $GOPATH is configured correctly for `make install` to function
- Cross-compilation tools are available via Go toolchain (no additional installation required)

## Out of Scope *(optional)*

- Docker image building (placeholder target only, actual implementation in later features)
- CI/CD pipeline configuration (handled in F005)
- Actual functional CLI commands (handled in F006 with Cobra integration)
- Schema file content (handled in F008 and beyond)
- Documentation site implementation (Docusaurus setup in F037)
- Build caching or optimization beyond standard Go compiler capabilities
- Support for platforms beyond the specified 5 OS/architecture combinations
