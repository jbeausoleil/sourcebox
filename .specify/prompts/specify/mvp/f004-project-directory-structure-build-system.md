# F004 - Project Directory Structure & Build System

## Feature Metadata
- **Feature ID**: F004
- **Name**: Project Directory Structure & Build System
- **Category**: Foundation
- **Phase**: Week 3
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (1 day)
- **Dependencies**: F003 (Git repository must exist)

## Constitutional Alignment

### Core Principles
- ✅ **Boring Tech Wins**: Standard Go project layout, Makefile for build automation
- ✅ **Ship Fast, Validate Early**: Simple structure that enables rapid development
- ✅ **Developer-First Design**: Clear organization, predictable file locations
- ✅ **Speed > Features**: Optimized build process, compressed binaries

### Technical Constraints
- ✅ **Platform Support**: Cross-compilation for macOS (Intel + Apple Silicon), Linux (x86_64 + ARM64), Windows (x86_64)
- ✅ **Binary size**: < 50MB compressed per platform
- ✅ **Build time**: Reasonable compile times

### Development Practices
- ✅ **Simple > Complex**: Follow standard Go project layout, avoid over-engineering

## User Story
**US-MVP-002**: "As a developer, I want a well-organized project structure that follows Go best practices so I can quickly locate files and understand the codebase."

## Problem Statement
SourceBox needs a clear, maintainable directory structure that follows Go community conventions. The build system must support cross-platform compilation and produce small, efficient binaries for distribution. Developers should be able to build, test, and install the CLI tool using simple `make` commands.

## Solution Overview
Create a standard Go project layout with directories for commands (`/cmd`), packages (`/pkg`), schemas (`/schemas`), Docker files (`/docker`), documentation (`/docs`), and examples (`/examples`). Implement a Makefile with targets for building, testing, installing, and Docker operations. Configure cross-compilation for all supported platforms with binary compression.

## Detailed Requirements

### Acceptance Criteria
1. **Directory Structure Created**:
   ```
   sourcebox/
   ├── cmd/              # CLI commands and main.go
   ├── pkg/              # Internal packages
   │   ├── generators/   # Data generation logic
   │   ├── schema/       # Schema parsing and validation
   │   └── database/     # Database connectors
   ├── schemas/          # Schema JSON definitions
   ├── docker/           # Dockerfiles and Docker scripts
   ├── docs/             # Documentation source
   ├── examples/         # Usage examples
   ├── .github/          # GitHub Actions workflows (F005)
   ├── Makefile          # Build automation
   └── [root files from F003]
   ```

2. **Makefile Created** with the following targets:
   - `make build` - Build for current platform
   - `make test` - Run all tests
   - `make install` - Install binary to $GOPATH/bin
   - `make docker-build` - Build Docker images (placeholder)
   - `make clean` - Remove build artifacts
   - `make build-all` - Cross-compile for all platforms
   - `make help` - Show available targets

3. **Cross-Compilation Working**:
   - macOS Intel (darwin/amd64)
   - macOS Apple Silicon (darwin/arm64)
   - Linux x86_64 (linux/amd64)
   - Linux ARM64 (linux/arm64)
   - Windows x86_64 (windows/amd64)

4. **Binary Size < 50MB**: Compressed binaries meet size requirement

5. **Build Output Directory**: Builds go to `/dist` directory, organized by platform

### Technical Specifications

#### Directory Purpose & Organization

**/cmd/sourcebox/**
- Contains `main.go` and CLI command setup
- Minimal business logic - delegates to /pkg
- Cobra command definitions

**/pkg/generators/**
- Data generation engine
- Base and custom generators
- Distribution logic

**/pkg/schema/**
- Schema parser and validator
- Schema data structures
- Schema loading (embedded and file-based)

**/pkg/database/**
- MySQL and PostgreSQL connectors
- Database abstraction layer
- Seeding logic

**/schemas/**
- fintech-loans.json
- healthcare-patients.json
- retail-ecommerce.json
- schema-spec.md (documentation)

**/docker/**
- mysql/Dockerfile
- postgres/Dockerfile
- build-all.sh
- test-all.sh
- docker-compose.yml

**/docs/**
- Docusaurus site source (F037)
- Architecture documentation
- API documentation

**/examples/**
- Quick start examples
- Integration examples
- Sample code

#### Makefile Implementation

```makefile
# Key variables
BINARY_NAME=sourcebox
VERSION=$(shell git describe --tags --always --dirty)
BUILD_DIR=dist
PLATFORMS=darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64

# Default target
.DEFAULT_GOAL := help

# Build for current platform
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/sourcebox

# Run tests
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Install binary
install:
	go install -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/sourcebox

# Cross-compile for all platforms
build-all:
	@for platform in $(PLATFORMS); do \
		GOOS=$${platform%/*} GOARCH=$${platform#*/} \
		go build -o $(BUILD_DIR)/$(BINARY_NAME)-$${platform%/*}-$${platform#*/} \
		-ldflags="-s -w -X main.version=$(VERSION)" ./cmd/sourcebox; \
		gzip $(BUILD_DIR)/$(BINARY_NAME)-$${platform%/*}-$${platform#*/}; \
	done

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR) coverage.txt

# Help target
help:
	@echo "Available targets:"
	@echo "  build       - Build for current platform"
	@echo "  test        - Run all tests"
	@echo "  install     - Install binary to GOPATH/bin"
	@echo "  build-all   - Cross-compile for all platforms"
	@echo "  clean       - Remove build artifacts"
```

#### Build Flags Explanation
- `-ldflags="-s -w"`: Strip debug info and symbol table (reduces binary size)
- `-X main.version=$(VERSION)`: Inject version at compile time
- `gzip`: Compress binaries to meet <50MB requirement

### Performance Considerations
- **Build Speed**: Parallel compilation where possible
- **Binary Size**: Use build flags to minimize size
- **Compression**: gzip all distribution binaries

### Testing Strategy
**Manual Verification**:
1. Run `make build` - binary created in /dist
2. Run binary - verify it executes
3. Run `make test` - placeholder test passes
4. Run `make build-all` - all platform binaries created
5. Check binary sizes - all < 50MB compressed
6. Verify cross-compiled binaries on respective platforms (macOS, Linux, Windows)

**Automated Tests**:
- Placeholder test in `/cmd/sourcebox/main_test.go` to verify build system works
- More comprehensive tests added in subsequent features

## Dependencies
- **Upstream**: F003 (Git repository and Go module must exist)
- **Downstream**:
  - F006 (Cobra CLI) will use /cmd structure
  - F008 (Schema Parser) will use /pkg/schema
  - F011-F020 (Data Generation) will use /pkg/generators
  - F023-F024 (Database Connectors) will use /pkg/database

## Deliverables
1. Complete directory structure with all required folders
2. Makefile with all specified targets
3. Cross-compilation configuration
4. Build script that produces binaries for all platforms
5. Placeholder main.go in /cmd/sourcebox
6. Documentation of directory structure (in README or docs/)

## Success Criteria
- ✅ All directories created and documented
- ✅ Makefile works with all targets
- ✅ Cross-compilation produces binaries for 5 platforms
- ✅ All compressed binaries < 50MB
- ✅ Build process is fast and reliable
- ✅ Structure follows Go community best practices

## Anti-Patterns to Avoid
- ❌ Overly complex directory structure (keep it simple)
- ❌ Non-standard Go layout (confuses developers)
- ❌ Large binaries (>50MB compressed)
- ❌ Platform-specific code in wrong locations
- ❌ Missing Makefile targets (should cover all common tasks)
- ❌ Slow build times (optimize with proper flags)

## Implementation Notes
- Follow golang-standards/project-layout (unofficial but widely adopted)
- Keep /cmd minimal - main.go should just bootstrap Cobra
- Use /pkg for all business logic
- /internal not needed yet (project will be open source)
- Version injection via ldflags enables `sourcebox --version`

## TDD Requirements
**Not applicable for project structure** - This is organizational setup. However, do create a placeholder test file to verify the build system works:

```go
// cmd/sourcebox/main_test.go
package main

import "testing"

func TestPlaceholder(t *testing.T) {
	// Verify build system works
	t.Log("Build system operational")
}
```

## Related Constitution Sections
- **Boring Tech Wins (Principle IV)**: Standard Go project layout, Makefile
- **Platform Support (Technical Constraint 7)**: macOS, Linux, Windows support
- **Binary size < 50MB** (Technical Constraint 1): Memory footprint requirement
- **Developer-First Design (Principle VI)**: Clear structure, predictable organization
