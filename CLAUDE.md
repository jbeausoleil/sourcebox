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
- Cobra v1.8+ (CLI framework) (004-f006-cobra-cli)
- cobra-cli v1.3+ (CLI scaffolding tool) (004-f006-cobra-cli)
- testify (testing assertions) (004-f006-cobra-cli)

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

## Cobra CLI Framework (004-f006-cobra-cli)

### Command Structure
```
cmd/sourcebox/
├── main.go              # Entry point, version injection
└── cmd/
    ├── root.go          # Root command, global flags
    ├── seed.go          # Seed command (scaffolded)
    └── list_schemas.go  # List-schemas command (scaffolded)
```

### Global Flags (Persistent)
- `--verbose, -v`: Enable verbose output (applies to all commands)
- `--quiet, -q`: Suppress non-error output (applies to all commands)
- `--config <file>`: Custom config file path (applies to all commands)

### Command Patterns
- **Root command**: `sourcebox` (shows help, version)
- **Subcommands**: `seed`, `list-schemas` (implementation in F021, F022)
- **Help system**: `--help` flag on all commands, comprehensive examples
- **Version display**: `--version` flag, injected at build time via ldflags

### Version Injection
```makefile
VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"
```

### Help Text Standards
- **Use**: Command signature with argument placeholders (`seed <database>`)
- **Short**: One-line description (<60 chars)
- **Long**: Multi-paragraph explanation with use cases
- **Example**: 2-3 practical copyable examples
- **Tone**: Clear, developer-friendly, no marketing fluff

### Testing Patterns
- Table-driven tests for flag combinations
- Output buffer capture (`cmd.SetOut(buf)`, `cmd.SetErr(buf)`)
- Flag parsing verification (`cmd.SetArgs()`, `cmd.Execute()`)
- Coverage target: >80% for cmd/ package
- TDD required for all command logic

### Error Handling
- Use `RunE` (not `Run`) for error returns
- Return errors from Run function (Cobra handles display + exit code)
- Error format: `fmt.Errorf("context: %w", err)` for wrapping
- User-facing errors: Clear, actionable (what's wrong + how to fix)
- SilenceUsage: true (don't print usage on business logic errors)

## Recent Changes
- 004-f006-cobra-cli: Added Cobra v1.8+ CLI framework integration
- 004-f006-cobra-cli: Added cobra-cli v1.3+ scaffolding tool
- 004-f006-cobra-cli: Added testify for unit test assertions
- 004-f006-cobra-cli: Deferred output helper functions to F021 (VerbosePrintf, QuietPrintf)
- 004-f006-cobra-cli: Deferred color output support to F021
- 004-f006-cobra-cli: Deferred verbosity levels (-vv, -vvv) to future releases
- 003-f005-github-actions: Added Go 1.21 and 1.22 (matrix testing across both versions)

## Known Technical Debt (004-f006-cobra-cli)
- T054: Test state pollution (tests pass individually, some fail in full suite) - decision pending
- T055: Coverage methodology needs documentation (76.2% vs 84.2% reporting)
- Integration tests require pre-built binary (accepted design pattern)

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
