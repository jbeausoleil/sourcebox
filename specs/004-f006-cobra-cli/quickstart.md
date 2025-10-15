# F006 Quickstart: Cobra CLI Framework Verification

**Feature**: F006 - Cobra CLI Framework Integration
**Branch**: `004-f006-cobra-cli`
**Date**: 2025-10-14

## Prerequisites

Before starting, verify:
- ✅ F004 completed (Project directory structure and Makefile functional)
- ✅ Go 1.21+ installed: `go version`
- ✅ Git repository initialized with F003 baseline

## Installation

### 1. Install Cobra CLI Tool

```bash
# Install cobra-cli scaffolding tool
go install github.com/spf13/cobra-cli@latest

# Verify installation
cobra-cli --version
```

**Expected output**: `Cobra CLI v1.3.0` (or newer)

**Troubleshooting**:
- Command not found? Add `$GOPATH/bin` to PATH:
  ```bash
  export PATH=$PATH:$(go env GOPATH)/bin
  ```
- Add to shell profile (~/.zshrc, ~/.bashrc):
  ```bash
  echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
  source ~/.zshrc
  ```

---

## Workflow Overview

The Cobra CLI framework provides this structure:

```
cmd/sourcebox/
├── main.go              # Entry point, version injection
└── cmd/
    ├── root.go          # Root command, global flags
    ├── seed.go          # Seed command (scaffolded)
    └── list_schemas.go  # List-schemas command (scaffolded)
```

**Phase 0**: Initialize Cobra structure (root command)
**Phase 1**: Scaffold subcommands (seed, list-schemas)
**Phase 2**: Verify help system and version display
**Phase 3**: Run unit tests

---

## Phase 0: Initialize Cobra Structure

### Step 1: Initialize Cobra

```bash
# From repository root
cd /Users/jbeausoleil/Projects/03_projects/personal/sourcebox

# Initialize Cobra (creates cmd/sourcebox/ directory)
cobra-cli init --config .cobra.yaml
```

**Expected**: Creates `cmd/sourcebox/main.go` and `cmd/sourcebox/cmd/root.go`

### Step 2: Verify Structure

```bash
# List created files
ls -la cmd/sourcebox/
ls -la cmd/sourcebox/cmd/
```

**Expected output**:
```
cmd/sourcebox/:
main.go

cmd/sourcebox/cmd:
root.go
```

### Step 3: Review Generated Files

**main.go** (entry point):
```bash
cat cmd/sourcebox/main.go
```

Expected content:
```go
package main

import "github.com/jbeausoleil/sourcebox/cmd/sourcebox/cmd"

func main() {
    cmd.Execute()
}
```

**cmd/root.go** (root command):
```bash
cat cmd/sourcebox/cmd/root.go
```

Expected: Cobra boilerplate with rootCmd definition

---

## Phase 1: Customize Root Command

### Step 4: Update Root Command Metadata

Edit `cmd/sourcebox/cmd/root.go`:

```go
var rootCmd = &cobra.Command{
    Use:   "sourcebox",
    Short: "Generate realistic, verticalized demo data instantly",
    Long: `SourceBox generates production-like demo data for databases.

Built for developers who need realistic demo data in seconds, not hours.
Verticalized schemas for fintech, healthcare, retail, and more.

Works entirely offline - no cloud APIs, no authentication, no network calls.`,
    Version: "dev",  // Will be injected at build time
}
```

### Step 5: Add Global Flags

Add to `init()` function in `cmd/root.go`:

```go
func init() {
    rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
    rootCmd.PersistentFlags().BoolP("quiet", "q", false, "suppress non-error output")
    rootCmd.PersistentFlags().String("config", "", "config file path (default: ~/.sourcebox.yaml)")
}
```

### Step 6: Update main.go for Version Injection

Edit `cmd/sourcebox/main.go`:

```go
package main

import (
    "os"
    "github.com/jbeausoleil/sourcebox/cmd/sourcebox/cmd"
)

var version = "dev"  // Overridden at build time via ldflags

func main() {
    cmd.SetVersion(version)
    if err := cmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

### Step 7: Export SetVersion Function

Add to `cmd/root.go`:

```go
func SetVersion(v string) {
    rootCmd.Version = v
}
```

---

## Phase 2: Build and Test Root Command

### Step 8: Install Cobra Dependency

```bash
# Add Cobra to go.mod
go get github.com/spf13/cobra@latest

# Tidy dependencies
go mod tidy
```

### Step 9: Build Binary

```bash
# Build using existing Makefile (from F004)
make build
```

**Expected output**:
```
Building sourcebox...
Built: dist/sourcebox
```

### Step 10: Test Help Output

```bash
# Test root help
./dist/sourcebox --help
```

**Expected output**:
```
SourceBox generates production-like demo data for databases.

Built for developers who need realistic demo data in seconds, not hours.
Verticalized schemas for fintech, healthcare, retail, and more.

Works entirely offline - no cloud APIs, no authentication, no network calls.

Usage:
  sourcebox [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
      --config string   config file path (default: ~/.sourcebox.yaml)
  -h, --help            help for sourcebox
  -q, --quiet           suppress non-error output
  -v, --verbose         verbose output
      --version         version for sourcebox

Use "sourcebox [command] --help" for more information about a command.
```

### Step 11: Test Version Display

```bash
# Test version flag
./dist/sourcebox --version
```

**Expected output**:
```
sourcebox version dev
```

**Note**: Shows "dev" for local builds. After Makefile update, will show actual git version.

### Step 12: Test Global Flags

```bash
# Test verbose flag
./dist/sourcebox --verbose --help

# Test quiet flag
./dist/sourcebox --quiet --help

# Test config flag (should parse without error)
./dist/sourcebox --config=/tmp/test.yaml --help
```

**Expected**: All commands execute without errors, help displays correctly.

---

## Phase 3: Scaffold Subcommands

### Step 13: Add Seed Command

```bash
# Generate seed command
cobra-cli add seed
```

**Expected**: Creates `cmd/sourcebox/cmd/seed.go`

### Step 14: Customize Seed Command

Edit `cmd/sourcebox/cmd/seed.go`:

```go
var seedCmd = &cobra.Command{
    Use:   "seed <database>",
    Short: "Seed a database with realistic demo data",
    Long: `Seed a database with verticalized, production-like demo data.

SourceBox generates realistic data based on industry-specific schemas
(fintech, healthcare, retail) with proper relationships, distributions,
and edge cases. Data is deterministic and reproducible.

Supported databases: mysql, postgres
Supported schemas: fintech-loans, healthcare-patients, retail-orders`,

    Example: `  # Seed MySQL with 1000 fintech loan records
  sourcebox seed mysql --schema=fintech-loans --records=1000

  # Seed Postgres with healthcare patient data
  sourcebox seed postgres --schema=healthcare-patients --records=5000

  # Export to SQL file instead of inserting
  sourcebox seed mysql --schema=fintech-loans --output=loans.sql`,

    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Seed command - implementation coming in F021")
        fmt.Printf("  Database: %s\n", args[0])
        schema, _ := cmd.Flags().GetString("schema")
        records, _ := cmd.Flags().GetInt("records")
        fmt.Printf("  Schema: %s\n", schema)
        fmt.Printf("  Records: %d\n", records)
    },
}

func init() {
    rootCmd.AddCommand(seedCmd)

    seedCmd.Flags().StringP("schema", "s", "", "schema name (required)")
    seedCmd.Flags().IntP("records", "n", 1000, "number of records to generate")
    seedCmd.Flags().String("host", "localhost", "database host")
    seedCmd.Flags().Int("port", 0, "database port (auto-detect by database type)")
    seedCmd.Flags().String("user", "root", "database user")
    seedCmd.Flags().String("password", "", "database password")
    seedCmd.Flags().String("db-name", "demo", "database name")
    seedCmd.Flags().String("output", "", "export to SQL file instead of inserting")
    seedCmd.Flags().Bool("dry-run", false, "show what would be done without executing")

    seedCmd.MarkFlagRequired("schema")
}
```

### Step 15: Add List-Schemas Command

```bash
# Generate list-schemas command
cobra-cli add list-schemas
```

**Expected**: Creates `cmd/sourcebox/cmd/list_schemas.go`

### Step 16: Customize List-Schemas Command

Edit `cmd/sourcebox/cmd/list_schemas.go`:

```go
var listSchemasCmd = &cobra.Command{
    Use:     "list-schemas",
    Aliases: []string{"ls"},
    Short:   "List all available data schemas",
    Long: `List all available verticalized data schemas.

SourceBox provides industry-specific schemas for fintech, healthcare,
retail, and other verticals. Each schema includes realistic field
distributions, relationships, and edge cases.

Schemas are categorized by industry and use case.`,

    Example: `  # List all available schemas
  sourcebox list-schemas

  # Using short alias
  sourcebox ls`,

    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("List-schemas command - implementation coming in F022")
        fmt.Println("Available schemas:")
        fmt.Println("  - fintech-loans")
        fmt.Println("  - healthcare-patients")
        fmt.Println("  - retail-orders")
    },
}

func init() {
    rootCmd.AddCommand(listSchemasCmd)
}
```

---

## Phase 4: Test Subcommands

### Step 17: Rebuild Binary

```bash
# Rebuild with new commands
make build
```

### Step 18: Test Seed Command Help

```bash
# Test seed help
./dist/sourcebox seed --help
```

**Expected output**:
```
Seed a database with verticalized, production-like demo data.

SourceBox generates realistic data based on industry-specific schemas
(fintech, healthcare, retail) with proper relationships, distributions,
and edge cases. Data is deterministic and reproducible.

Supported databases: mysql, postgres
Supported schemas: fintech-loans, healthcare-patients, retail-orders

Usage:
  sourcebox seed <database> [flags]

Examples:
  # Seed MySQL with 1000 fintech loan records
  sourcebox seed mysql --schema=fintech-loans --records=1000

  # Seed Postgres with healthcare patient data
  sourcebox seed postgres --schema=healthcare-patients --records=5000

  # Export to SQL file instead of inserting
  sourcebox seed mysql --schema=fintech-loans --output=loans.sql

Flags:
      --db-name string     database name (default "demo")
      --dry-run            show what would be done without executing
  -h, --help               help for seed
      --host string        database host (default "localhost")
      --output string      export to SQL file instead of inserting
      --password string    database password
      --port int           database port (auto-detect by database type)
  -n, --records int        number of records to generate (default 1000)
  -s, --schema string      schema name (required)
      --user string        database user (default "root")

Global Flags:
      --config string   config file path (default: ~/.sourcebox.yaml)
  -q, --quiet           suppress non-error output
  -v, --verbose         verbose output
```

### Step 19: Test List-Schemas Command Help

```bash
# Test list-schemas help
./dist/sourcebox list-schemas --help

# Test short alias
./dist/sourcebox ls --help
```

**Expected**: Both commands show the same help output.

### Step 20: Test Subcommand Registration

```bash
# Verify commands appear in root help
./dist/sourcebox --help
```

**Expected**: "Available Commands" section includes:
```
Available Commands:
  completion    Generate the autocompletion script for the specified shell
  help          Help about any command
  list-schemas  List all available data schemas
  seed          Seed a database with realistic demo data
```

### Step 21: Test Placeholder Execution

```bash
# Execute seed command (placeholder)
./dist/sourcebox seed mysql --schema=fintech-loans --records=1000

# Execute list-schemas command (placeholder)
./dist/sourcebox list-schemas

# Test short alias
./dist/sourcebox ls
```

**Expected output**:

Seed:
```
Seed command - implementation coming in F021
  Database: mysql
  Schema: fintech-loans
  Records: 1000
```

List-schemas:
```
List-schemas command - implementation coming in F022
Available schemas:
  - fintech-loans
  - healthcare-patients
  - retail-orders
```

---

## Phase 5: Update Build System for Version Injection

### Step 22: Update Makefile

Edit `Makefile` to add version injection:

```makefile
# Version from git tags (fallback to commit SHA)
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

# Build flags
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"

.PHONY: build
build: ## Build binary for current platform
	@echo "Building sourcebox $(VERSION)..."
	@mkdir -p dist
	@go build $(LDFLAGS) -o dist/sourcebox ./cmd/sourcebox
	@echo "Built: dist/sourcebox"
```

### Step 23: Test Version Injection

```bash
# Rebuild with version injection
make build

# Check version (should show git commit SHA or tag)
./dist/sourcebox --version
```

**Expected output** (example):
```
sourcebox version 82479a7-dirty
```

- Shows commit SHA if no tags exist
- Shows tag if tagged (e.g., v1.0.0)
- Adds "-dirty" suffix if uncommitted changes

---

## Phase 6: Unit Tests

### Step 24: Create Test File

Create `cmd/sourcebox/cmd/root_test.go`:

```go
package cmd

import (
    "bytes"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestRootCommandHelp(t *testing.T) {
    buf := new(bytes.Buffer)
    rootCmd.SetOut(buf)
    rootCmd.SetErr(buf)
    rootCmd.SetArgs([]string{"--help"})

    err := rootCmd.Execute()
    require.NoError(t, err)

    output := buf.String()
    assert.Contains(t, output, "SourceBox")
    assert.Contains(t, output, "Generate realistic, verticalized demo data")
    assert.Contains(t, output, "--verbose")
    assert.Contains(t, output, "--quiet")
    assert.Contains(t, output, "--config")
}

func TestVersionFlag(t *testing.T) {
    buf := new(bytes.Buffer)
    rootCmd.SetOut(buf)
    rootCmd.SetArgs([]string{"--version"})

    SetVersion("v1.0.0-test")

    err := rootCmd.Execute()
    require.NoError(t, err)
    assert.Contains(t, buf.String(), "v1.0.0-test")
}

func TestCommandsRegistered(t *testing.T) {
    commands := rootCmd.Commands()
    commandNames := make([]string, 0)
    for _, cmd := range commands {
        commandNames = append(commandNames, cmd.Name())
    }

    assert.Contains(t, commandNames, "seed")
    assert.Contains(t, commandNames, "list-schemas")
}
```

### Step 25: Install Test Dependencies

```bash
# Install testify for assertions
go get github.com/stretchr/testify@latest
go mod tidy
```

### Step 26: Run Tests

```bash
# Run all tests in cmd package
go test ./cmd/sourcebox/cmd/... -v

# Check coverage
go test ./cmd/sourcebox/cmd/... -coverprofile=coverage.txt
go tool cover -func=coverage.txt
```

**Expected output**:
```
=== RUN   TestRootCommandHelp
--- PASS: TestRootCommandHelp (0.00s)
=== RUN   TestVersionFlag
--- PASS: TestVersionFlag (0.00s)
=== RUN   TestCommandsRegistered
--- PASS: TestCommandsRegistered (0.00s)
PASS
coverage: 85.2% of statements
```

---

## Verification Checklist

After completing all phases, verify:

- ✅ cobra-cli tool installed and functional
- ✅ Cobra structure initialized (main.go, cmd/root.go)
- ✅ Root command displays help correctly
- ✅ Version flag shows version information
- ✅ Global flags parse without errors (--verbose, --quiet, --config)
- ✅ Seed command scaffolded with comprehensive help text
- ✅ List-schemas command scaffolded with alias (ls)
- ✅ Subcommands appear in root help output
- ✅ All commands have clear, useful help text with examples
- ✅ Placeholder execution works for all commands
- ✅ Build system updated (version injection functional)
- ✅ Unit tests pass with >80% coverage

---

## Performance Verification

### Response Time Checks

```bash
# Test help response time
time ./dist/sourcebox --help

# Test version response time
time ./dist/sourcebox --version

# Test subcommand help response time
time ./dist/sourcebox seed --help
```

**Expected**:
- Help display: <1 second
- Version display: <100ms
- No noticeable startup delay

**Success criteria** (from spec):
- SC-006: Command-line interface responds to user input in under 100ms for all help and version requests

---

## Common Issues & Debugging

### Issue: cobra-cli command not found

**Symptom**: `command not found: cobra-cli`

**Check**:
```bash
# Is cobra-cli installed?
go install github.com/spf13/cobra-cli@latest

# Is $GOPATH/bin in PATH?
echo $PATH | grep "$(go env GOPATH)/bin"
```

**Fix**:
```bash
# Add to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Make permanent (add to ~/.zshrc or ~/.bashrc)
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

---

### Issue: Commands not appearing in help

**Symptom**: `sourcebox --help` doesn't show seed or list-schemas commands

**Check**:
```bash
# Does cmd/seed.go have init() function?
grep -n "func init()" cmd/sourcebox/cmd/seed.go

# Does init() call AddCommand?
grep -n "rootCmd.AddCommand" cmd/sourcebox/cmd/seed.go
```

**Fix**:
Ensure `cmd/seed.go` has:
```go
func init() {
    rootCmd.AddCommand(seedCmd)
}
```

---

### Issue: Version shows "dev" instead of git version

**Symptom**: `sourcebox --version` shows "dev" even after Makefile update

**Check**:
```bash
# Is VERSION being set correctly in Makefile?
make -n build | grep ldflags

# Does git describe work?
git describe --tags --always --dirty
```

**Fix**:
```bash
# If no tags exist, git describe shows commit SHA (expected)
# To test with a tag:
git tag v0.1.0
make build
./dist/sourcebox --version  # Should show v0.1.0
```

---

### Issue: Flags not parsing correctly

**Symptom**: `sourcebox seed --schema=test` gives error "unknown flag: --schema"

**Check**:
```bash
# Is flag defined in seed command?
grep -n "Flags().String" cmd/sourcebox/cmd/seed.go

# Is seed command initialized?
go build ./cmd/sourcebox && ./sourcebox seed --help
```

**Fix**:
- PersistentFlags go in root.go (global)
- Local Flags go in specific command file (seed.go)
- Verify flag is defined BEFORE MarkFlagRequired

---

### Issue: Tests failing to capture output

**Symptom**: Tests fail with "expected output not found"

**Check**:
```go
// Are you setting output buffers?
buf := new(bytes.Buffer)
rootCmd.SetOut(buf)
rootCmd.SetErr(buf)
```

**Fix**:
Always set output buffers BEFORE calling Execute():
```go
func TestExample(t *testing.T) {
    buf := new(bytes.Buffer)
    rootCmd.SetOut(buf)  // Capture stdout
    rootCmd.SetErr(buf)  // Capture stderr
    rootCmd.SetArgs([]string{"--help"})

    err := rootCmd.Execute()
    require.NoError(t, err)

    output := buf.String()
    assert.Contains(t, output, "expected text")
}
```

---

### Issue: Help text not showing

**Symptom**: `sourcebox seed --help` shows minimal output

**Check**:
```bash
# Are Use, Short, Long, Example fields set?
grep -A 10 "var seedCmd" cmd/sourcebox/cmd/seed.go
```

**Fix**:
Ensure seedCmd has all fields:
```go
var seedCmd = &cobra.Command{
    Use:     "seed <database>",  // Required
    Short:   "Brief description",  // Required
    Long:    `Detailed description`,  // Recommended
    Example: `Examples here`,  // Recommended
}
```

---

## Next Steps

After F006 completion:

1. **F021: Implement Seed Command**
   - Replace placeholder in `cmd/seed.go`
   - Add database connection logic
   - Implement data generation
   - Full TDD with integration tests

2. **F022: Implement List-Schemas Command**
   - Replace placeholder in `cmd/list_schemas.go`
   - Read schema definitions from files
   - Format output (table or JSON)
   - Add filtering/search functionality

3. **F027: CLI Output Formatting**
   - Add progress bars for seeding
   - Implement table formatting
   - Add color output (optional)
   - Improve error message formatting

---

## Summary

**F006 establishes the CLI foundation**:
- ✅ Cobra framework integrated
- ✅ Root command with global flags (--verbose, --quiet, --config)
- ✅ Version injection from git tags
- ✅ Subcommands scaffolded (seed, list-schemas)
- ✅ Comprehensive help system
- ✅ Unit tests with >80% coverage
- ✅ Cross-platform build support

**Constitutional compliance**:
- ✅ Boring Tech: Cobra is standard, proven framework
- ✅ Speed: <100ms response time for help/version
- ✅ Developer-First: Comprehensive help, clear structure
- ✅ Local-First: No network dependencies
- ✅ TDD: Unit tests for all command logic
- ✅ Platform Support: Works on macOS, Linux, Windows

**Ready for**:
- F021: Full seed command implementation
- F022: Full list-schemas implementation
- F027: Enhanced output formatting
