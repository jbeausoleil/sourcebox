# SourceBox Development Guidelines

Auto-generated from all feature plans. Last updated: 2025-10-14

## Project Overview
- **Repository**: github.com/jbeausoleil/sourcebox
- **Language**: Go 1.21+
- **License**: MIT (open source forever)
- **Development Status**: Foundational setup (Week 3 - F003)

## Active Technologies
- Go 1.21+ (CLI development)
- Git (version control)
- GitHub (repository hosting, CI/CD)
- Make (build automation)
- Go 1.21 and 1.22 (matrix testing across both versions) (003-f005-github-actions)
- GitHub Actions runners (standard ubuntu-latest, macos-latest, windows-latest) (003-f005-github-actions)

## Project Structure
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
├── .gitignore          # Go-specific patterns
├── LICENSE             # MIT license
├── README.md           # Project overview with legal notice
├── CONTRIBUTING.md     # Contribution guidelines
├── CODE_OF_CONDUCT.md  # Contributor Covenant v2.1
├── go.mod              # Go module configuration
└── go.sum              # (generated after dependencies added)
```

## Go Module
- **Module Path**: github.com/jbeausoleil/sourcebox
- **Go Version**: 1.21+ (minimum)
- **Dependencies**: None yet (to be added in F009)

## Code Style
- Follow standard Go conventions (gofmt)
- Run `go vet ./...` before committing
- Run `gofmt -w .` before committing
- Write clear, self-documenting code
- Keep functions focused and small

## Testing Requirements
- **Core logic**: TDD required (write test first)
- **Coverage**: Aim for >80% on core packages
- **Manual QA**: Test on macOS, Linux, Windows for releases

## Build System

### Makefile Targets
- `make build` - Build for current platform → dist/sourcebox
- `make test` - Run all tests with race detection and coverage
- `make install` - Install binary to $GOPATH/bin
- `make build-all` - Cross-compile for all 5 platforms (darwin/amd64, darwin/arm64, linux/amd64, linux/arm64, windows/amd64)
- `make clean` - Remove dist/ and coverage files
- `make help` - Show available targets (default)

### Cross-Compilation Platforms
- macOS Intel (darwin/amd64)
- macOS Apple Silicon (darwin/arm64)
- Linux x86_64 (linux/amd64)
- Linux ARM64 (linux/arm64)
- Windows x86_64 (windows/amd64)

### Build Artifacts
- **Location**: `/dist` directory (gitignored)
- **Naming**: `sourcebox-{os}-{arch}.gz` (compressed distribution binaries)
- **Size Constraint**: < 50MB compressed per binary
- **Optimization**: Binaries built with `-ldflags="-s -w"` (strip debug symbols)
- **Version Injection**: Version embedded at compile time via `git describe --tags --always --dirty`

### Performance Targets
- Single platform build: < 30 seconds
- All platforms build: < 2 minutes (use `make -j4 build-all` for parallel builds)
- Binary size: < 50MB compressed (constitutional requirement)

## Legal Notice
**CRITICAL**: This project is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information. This notice MUST appear prominently in README.md.

## Recent Changes
- 003-f005-github-actions: Added Go 1.21 and 1.22 (matrix testing across both versions)
- 002-f004-project-directory: Directory structure and build system with Makefile
- 001-f003-initialize-git: Foundational Git repository and Go module setup

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
