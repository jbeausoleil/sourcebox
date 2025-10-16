# SourceBox Development Guidelines

Auto-generated from all feature plans. Last updated: 2025-10-16

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
- Go 1.21+ (existing project configuration from F003/F004) (006-f008-schema-parser)
- Schemas loaded from files (development) or `embed.FS` (production binary) (006-f008-schema-parser)

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
- **Dependencies**: 6 core dependencies (configured in F009)
  - github.com/spf13/cobra@v1.8.0 - CLI framework
  - github.com/brianvoe/gofakeit/v6@v6.27.0 - Data generation
  - github.com/schollz/progressbar/v3@v3.14.1 - Progress bars
  - github.com/go-sql-driver/mysql@v1.7.1 - MySQL driver
  - github.com/lib/pq@v1.10.9 - PostgreSQL driver
  - github.com/fatih/color@v1.16.0 - Terminal colors

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

## Schema JSON Format (005-f007-schema-json)

### Overview
SourceBox uses JSON-based schema definitions to describe database schemas, data generators, and relationships. All schemas are stored in `schemas/` directory and follow the format specification in `schemas/schema-spec.md`.

### Schema Structure
```json
{
  "schema_version": "1.0",
  "name": "fintech-loans",
  "description": "Realistic fintech loan data",
  "author": "SourceBox Contributors",
  "version": "1.0.0",
  "database_type": ["mysql", "postgres"],
  "metadata": {
    "industry": "fintech",
    "tags": ["loans", "credit"],
    "total_records": 4950
  },
  "tables": [...],
  "relationships": [...],
  "generation_order": ["borrowers", "loans", "payments"],
  "validation_rules": [...]
}
```

### Generator Types

**Built-in Generators** (standard across all schemas):
- Personal Data: `first_name`, `last_name`, `full_name`, `email`, `phone`, `address`, `ssn`, `date_of_birth`
- Company Data: `company_name`, `job_title`, `company_email`, `domain`
- Date/Time: `timestamp_past`, `timestamp_future`, `date_between`
- Numeric: `int_range`, `float_range`, `decimal_range`

**Custom Generators** (schema-specific):
- Defined per schema (e.g., `credit_score`, `loan_amount`, `diagnosis_code`)
- Enable verticalized data generation (fintech, healthcare, retail)

### Distribution Types

**uniform**: Evenly distributed values
- Parameters: `{min, max}`
- Use case: Random IDs, equally likely categories

**normal**: Bell curve distribution
- Parameters: `{mean, std_dev, min?, max?}`
- Use case: Credit scores, heights, test scores

**lognormal**: Right-skewed distribution
- Parameters: `{median, min, max}`
- Use case: Loan amounts, income, prices (naturally skewed data)

**weighted**: Specific values with probabilities
- Parameters: `{values: [{value, weight}, ...]}`
- Use case: Categorical data (loan status: 70% active, 25% paid, 5% delinquent)

**ranges**: Multiple ranges with weights
- Parameters: `{ranges: [{min, max, weight}, ...]}`
- Use case: Tiered distributions (interest rates by risk tier)

### Foreign Key Relationships

**Dual Representation** (inline + explicit):

**Inline** (in column definition, what parser uses):
```json
{
  "name": "borrower_id",
  "type": "int",
  "foreign_key": {
    "table": "borrowers",
    "column": "id",
    "on_delete": "CASCADE",
    "on_update": "CASCADE"
  }
}
```

**Explicit** (separate section, for documentation):
```json
{
  "from_table": "loans",
  "from_column": "borrower_id",
  "to_table": "borrowers",
  "to_column": "id",
  "relationship_type": "many_to_one",
  "description": "Each loan belongs to one borrower"
}
```

### Schema Versioning (Semantic Versioning)

- **schema_version**: Format version (1.0 for MVP, 2.0 if format changes)
- **version**: Content version (semver: 1.0.0, 1.1.0, 2.0.0)

**Major** (1.0.0 → 2.0.0): Breaking changes (table removal, type changes)
**Minor** (1.0.0 → 1.1.0): Backward-compatible additions (new tables, new columns)
**Patch** (1.0.0 → 1.0.1): Bug fixes (distribution improvements, docs)

### Validation Rules

**Schema-level**: Unique name, at least one table, valid database_type, generation_order includes all tables
**Table-level**: Unique name, at least one column, exactly one primary key, record_count > 0
**Column-level**: Valid type, valid generator, params match requirements, foreign keys reference existing tables
**Relationship-level**: Foreign keys reference primary keys, valid integrity actions (CASCADE, SET NULL, RESTRICT)
**Generation order**: Parent tables first, no circular dependencies

### Tier 1 Complexity Targets

- **Records**: <5,000 total records
- **Generation time**: <30 seconds (constitutional constraint)
- **Tables**: 2-5 tables with simple relationships
- **Use case**: MVP schemas (fintech, healthcare, retail)

### Creating New Schemas

1. Copy example schema (`schemas/example-schema.json`)
2. Update metadata (name, description, author, version)
3. Define tables with columns and generators
4. Specify foreign keys (inline + explicit)
5. Set generation_order (parent tables first)
6. Validate JSON syntax (`cat schema.json | jq '.'`)
7. Verify Tier 1 compliance (<5,000 records)

### Supported Data Types (MySQL/PostgreSQL common subset)

- Integer: `int`, `bigint`, `smallint`, `tinyint`
- Decimal: `decimal(p,s)`, `float`, `double`
- String: `varchar(n)`, `text`, `char(n)`
- Date/Time: `date`, `datetime`, `timestamp`
- Boolean: `boolean` (PostgreSQL), `bit` (MySQL)
- JSON: `json` (both), `jsonb` (PostgreSQL only)
- Enum: `enum('val1','val2')` (MySQL format)

## Dependency Management (007-f009-dependency-management)

### Dependency Commands
- `go get <dependency>@<version>` - Add or update specific dependency
- `go mod tidy` - Clean up go.mod and generate go.sum checksums
- `go mod verify` - Verify dependency checksums (security check)
- `go mod download` - Download all dependencies to module cache
- `go list -m all` - List all dependencies (direct + transitive)
- `go mod graph` - Show dependency graph
- `go mod why <dependency>` - Explain why dependency is needed

### Version Selection
- Use **exact semantic versions** (e.g., `v1.8.0`, not `v1.8.x` or `latest`)
- Pin to **stable releases** (not pre-release, not `@main`)
- Update cadence: **Quarterly review** (every 3 months)
- Emergency updates: Critical security vulnerabilities (CVE)

### License Requirements
- All dependencies MUST be MIT-compatible
- Acceptable: MIT, Apache 2.0, BSD-2, BSD-3, ISC, MPL 2.0
- Prohibited: GPL, LGPL, AGPL, SSPL, Proprietary
- Verification: Manual for MVP, automated scanning in Phase 2 (F005 CI/CD)

### go.mod Organization
- Single require block with alphabetical order (enforced by `go mod tidy`)
- No comments in go.mod (use README.md for documentation)
- Indirect dependencies in separate block (automatic, marked with `// indirect`)
- Never manually edit indirect dependency versions

### Build Performance
- First build (cold cache): 15-20 seconds
- Subsequent builds (warm cache): 5-10 seconds
- Threshold: **< 30 seconds** (constitutional constraint)
- Build caching: Automatic via Go build cache
- Dependencies rarely change → cache highly effective

### Documentation Location
- README.md "Dependencies" section (after "Installation", before "Usage")
- Grouped by category: CLI & UX, Data Generation, Database Drivers
- Each dependency: Name (GitHub link), version (backticks), one-line purpose
- License compatibility statement at bottom

### Dependency Update Process
1. Check release notes for breaking changes
2. Update one dependency at a time (isolate issues)
3. Run `go get <dependency>@latest` or `@vX.Y.Z`
4. Run `go mod tidy` after each update
5. Run full test suite (`go test ./...`)
6. Test key CLI commands manually
7. Commit go.mod and go.sum together
8. Update README.md dependency versions

### Known Dependency Licenses
- Cobra: Apache 2.0 ✅
- gofakeit: MIT ✅
- progressbar: MIT ✅
- go-sql-driver/mysql: MPL 2.0 ✅
- lib/pq: MIT ✅
- fatih/color: MIT ✅

### Transitive Dependencies
- Let Go manage automatically (don't manually add/edit)
- Trust major dependencies (Cobra, gofakeit) have vetted their deps
- Spot-check major transitive deps during quarterly review
- Full audit in Phase 2 (CI/CD with automated scanning)

### Offline Support
- All dependencies work offline after initial download
- Module cache: `~/go/pkg/mod` (persistent, local)
- Build cache: `~/.cache/go-build` (persistent, local)
- No network calls after initial `go mod download`

### Security
- go.sum MUST be committed to version control (reproducible builds)
- `go mod verify` checks checksums (detects tampering)
- Quarterly manual vulnerability check with `govulncheck`
- Phase 2: Automated security scanning in CI/CD (F005 follow-up)

## Recent Changes
- 007-f009-dependency-management: Added 6 core dependencies (Cobra, gofakeit, progressbar, MySQL driver, PostgreSQL driver, color)
- 007-f009-dependency-management: Configured dependency management with exact versions, MIT-compatible licenses
- 006-f008-schema-parser: Added Go 1.21+ (existing project configuration from F003/F004)
- 006-f008-schema-parser: Added [if applicable, e.g., PostgreSQL, CoreData, files or N/A]
- 005-f007-schema-json: Completed schema JSON format specification (schema_version 1.0)

## Known Technical Debt (004-f006-cobra-cli)
- T054: Test state pollution (tests pass individually, some fail in full suite) - decision pending
- T055: Coverage methodology needs documentation (76.2% vs 84.2% reporting)
- Integration tests require pre-built binary (accepted design pattern)

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
