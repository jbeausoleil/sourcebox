# Feature Specification Prompt: F009 - Dependency Management Setup

## Feature Metadata
- **Feature ID**: F009
- **Name**: Dependency Management Setup
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P1 (Should-have)
- **Estimated Effort**: Small (1 day)
- **Dependencies**: F003 (Go module initialized)

## Constitutional Alignment

### Core Principles
- ✅ **Boring Tech Wins**: Use standard Go dependency management (go modules)
- ✅ **Simple > Complex**: Minimal dependencies, standard tooling
- ✅ **Developer-First Design**: Clear dependency documentation

### Technical Constraints
- ✅ **Code Quality**: No unused dependencies
- ✅ **Security**: Verified checksums (go.sum)

### Development Practices
- ✅ **Indie Project Constraints**: Minimize dependency maintenance burden

## User Story
**US-MVP-002**: "As a developer, I want all dependencies clearly documented and properly managed so I can build the project without issues."

## Problem Statement
SourceBox requires external libraries for CLI framework (Cobra), data generation (Faker), progress display, and database drivers. Dependencies must be:
- **Documented**: Developers know what each dependency does
- **Verified**: go.sum ensures integrity
- **Minimal**: Avoid dependency bloat
- **Standard**: Use Go modules (not dep, glide, or vendor)

Without proper dependency management, builds fail, security is compromised, and maintenance burden grows.

## Solution Overview
Configure go.mod with all required dependencies, run `go mod tidy` to clean up, verify go.sum checksums, and document all dependencies in README.md with their purpose. Use standard Go tooling exclusively.

## Detailed Requirements

### Acceptance Criteria
1. **go.mod Updated**: All dependencies listed with versions
2. **go.sum Generated**: Checksums for all dependencies
3. **Key Dependencies Installed**:
   - `github.com/spf13/cobra` - CLI framework
   - `github.com/brianvoe/gofakeit/v6` - Data generation
   - `github.com/schollz/progressbar/v3` - Progress bars
   - `github.com/go-sql-driver/mysql` - MySQL driver
   - `github.com/lib/pq` - PostgreSQL driver
   - `github.com/fatih/color` - Terminal colors
4. **Dependency Documentation**: README.md lists all dependencies with purpose
5. **go mod tidy Succeeds**: No unused dependencies
6. **No Security Vulnerabilities**: Run `go list -m -json all | go-audit` (or similar)

### Technical Specifications

#### go.mod Contents

```go
module github.com/yourusername/sourcebox

go 1.21

require (
	github.com/spf13/cobra v1.8.0
	github.com/brianvoe/gofakeit/v6 v6.27.0
	github.com/schollz/progressbar/v3 v3.14.1
	github.com/go-sql-driver/mysql v1.7.1
	github.com/lib/pq v1.10.9
	github.com/fatih/color v1.16.0
)

require (
	// Indirect dependencies (automatically managed)
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
)
```

#### Dependency Documentation (README.md section)

```markdown
## Dependencies

SourceBox uses the following dependencies:

### CLI & UX
- **[Cobra](https://github.com/spf13/cobra)** (`v1.8.0`) - CLI framework (used by kubectl, Hugo, GitHub CLI)
- **[progressbar](https://github.com/schollz/progressbar)** (`v3.14.1`) - Progress bars for seeding operations
- **[color](https://github.com/fatih/color)** (`v1.16.0`) - Terminal color output

### Data Generation
- **[gofakeit](https://github.com/brianvoe/gofakeit)** (`v6.27.0`) - Realistic fake data generation

### Database Drivers
- **[go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)** (`v1.7.1`) - MySQL database driver
- **[lib/pq](https://github.com/lib/pq)** (`v1.10.9`) - PostgreSQL database driver

All dependencies are MIT or similarly permissive licenses compatible with our MIT license.
```

#### Installing Dependencies

```bash
# Install all dependencies
go mod download

# Clean up unused dependencies
go mod tidy

# Verify dependencies
go mod verify
```

#### Updating Dependencies (future)

```bash
# Update to latest versions
go get -u ./...

# Update specific dependency
go get github.com/spf13/cobra@latest

# Run tidy after updates
go mod tidy
```

### Dependency Rationale

**Cobra** - CLI Framework
- Why: Industry standard for Go CLIs (kubectl, gh, hugo all use it)
- Alternatives rejected: urfave/cli (less features), custom (too much work)
- License: Apache 2.0 (compatible)

**gofakeit** - Data Generation
- Why: Comprehensive fake data generators, well-maintained
- Alternatives rejected: faker (less comprehensive), custom (too much work)
- License: MIT (compatible)

**progressbar** - Progress Display
- Why: Simple, customizable progress bars
- Alternatives rejected: uiprogress (unmaintained), custom (reinventing wheel)
- License: MIT (compatible)

**go-sql-driver/mysql** - MySQL Driver
- Why: Standard Go MySQL driver
- Alternatives rejected: None (this is the standard)
- License: MPL 2.0 (compatible)

**lib/pq** - PostgreSQL Driver
- Why: Standard Go PostgreSQL driver
- Alternatives rejected: pgx (more complex than needed)
- License: MIT (compatible)

**fatih/color** - Terminal Colors
- Why: Simple, cross-platform color output
- Alternatives rejected: gookit/color (overkill), custom (not worth it)
- License: MIT (compatible)

### Performance Considerations
- All dependencies are lightweight
- No heavy frameworks or ORMs
- Minimal transitive dependencies
- Fast compile times maintained

### Security Considerations
- All dependencies have active maintenance
- No known CVEs in specified versions
- go.sum ensures integrity (checksums)
- Regular updates planned (quarterly review)

### Testing Strategy

**Manual Verification**:
```bash
# Verify go.mod is valid
go mod verify

# Check for vulnerabilities (if go-audit installed)
go list -m -json all | go-audit

# Ensure tidy doesn't change anything
go mod tidy
git diff go.mod go.sum  # Should be empty

# Build project to verify dependencies work
go build ./cmd/sourcebox

# Run tests to verify dependencies integrate correctly
go test ./...
```

**Automated Checks** (in CI/CD - F005):
- `go mod verify` in GitHub Actions
- Check for outdated dependencies
- Security scanning (e.g., govulncheck)

## Dependencies
- **Upstream**: F003 (Go module must be initialized)
- **Downstream**: All subsequent features use these dependencies

## Deliverables
1. go.mod with all required dependencies
2. go.sum with dependency checksums
3. Dependency documentation in README.md
4. Verification that `go mod tidy` passes
5. Verification that `go mod verify` passes

## Success Criteria
- ✅ All dependencies installed correctly
- ✅ go.mod and go.sum are clean and verified
- ✅ `go build` succeeds
- ✅ `go test` succeeds
- ✅ No unused dependencies
- ✅ All dependencies documented with rationale

## Anti-Patterns to Avoid
- ❌ Vendor directory (go modules handle this)
- ❌ Uncommitted go.sum (needed for reproducible builds)
- ❌ Unused dependencies (run `go mod tidy`)
- ❌ Outdated dependencies with known vulnerabilities
- ❌ Non-permissive licenses (must be MIT-compatible)
- ❌ Heavy frameworks (e.g., ORMs we don't need)

## Implementation Notes
- Go modules are standard since Go 1.11
- go.sum MUST be committed to Git
- Indirect dependencies are automatically managed
- Use specific versions, not `latest` (reproducibility)
- Review dependency licenses (ensure MIT-compatible)

## TDD Requirements
**Not applicable for dependency setup** - This is configuration work. However, verify dependencies work by:
1. Running `go build` successfully
2. Running existing tests successfully
3. Importing each dependency in code and verifying it compiles

## Related Constitution Sections
- **Boring Tech Wins (Principle IV)**: Standard Go modules, proven dependencies
- **Simple > Complex**: Minimal dependencies, no heavy frameworks
- **Cost Constraints (Technical Constraint 4)**: Free dependencies, no paid libraries
- **Open Source License (Technical Constraint 6)**: All dependencies MIT-compatible
- **Indie Project Constraints (Development Practice 7)**: Minimize maintenance burden
