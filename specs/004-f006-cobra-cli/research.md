# Research & Technical Decisions: Cobra CLI Framework Integration

**Feature**: F006 - Cobra CLI Framework Integration
**Branch**: `004-f006-cobra-cli`
**Date**: 2025-10-14

## Overview

This document captures all technical research and decisions made during Phase 0 planning for integrating the Cobra CLI framework into SourceBox. Each decision is documented with rationale, alternatives considered, and sources.

---

## 1. Cobra Initialization Structure

### Decision

Standard directory structure after `cobra-cli init`:

```
cmd/sourcebox/
├── main.go          # Entry point, version variable, calls cmd.Execute()
└── cmd/
    └── root.go      # Root command definition, global flags, Execute() function
```

Future commands added with `cobra-cli add <commandname>`:
```
cmd/sourcebox/cmd/
├── root.go
├── seed.go          # Added via: cobra-cli add seed
└── list_schemas.go  # Added via: cobra-cli add list-schemas
```

### Purpose of Each File

- **main.go**: Minimal bootstrap file
  - Declares `version` variable for build-time injection
  - Calls `cmd.Execute()` to run CLI
  - Handles process exit codes
  - No business logic (keep it simple)

- **cmd/root.go**: Root command configuration
  - Defines `rootCmd` (Cobra Command struct)
  - Sets up persistent flags (--verbose, --quiet, --config)
  - Implements `Execute()` function called from main.go
  - Contains command description, help text, examples
  - Root for all subcommands

- **cmd/<subcommand>.go**: Individual command files
  - Generated via `cobra-cli add <name>`
  - Defines command-specific logic
  - Local flags (only for that command)
  - Registers with rootCmd via init() function

### Rationale

This structure follows Cobra's official conventions and is used by production tools like kubectl, Hugo, and GitHub CLI. It provides:

- **Clear separation of concerns**: Bootstrap (main.go) vs command logic (cmd/)
- **Scalability**: Easy to add new commands without modifying main.go
- **Maintainability**: Each command in its own file
- **Convention over configuration**: Developers familiar with Cobra recognize this structure immediately

### Alternatives Considered

**Alternative 1**: Flat structure (all commands in main.go)
- Rejected: Doesn't scale beyond 2-3 commands, hard to maintain, no separation of concerns

**Alternative 2**: Deep nesting (cmd/commands/, cmd/flags/, cmd/utils/)
- Rejected: Over-engineering for MVP scope, adds unnecessary complexity

### Sources

- [Cobra official documentation](https://github.com/spf13/cobra#getting-started)
- [cobra-cli tool documentation](https://github.com/spf13/cobra-cli/blob/main/README.md)
- [kubectl source code](https://github.com/kubernetes/kubectl/tree/master/pkg/cmd) (reference implementation)

---

## 2. Root Command Architecture

### Decision

**main.go responsibilities** (MINIMAL):
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

**cmd/root.go responsibilities** (COMMAND LOGIC):
```go
package cmd

import (
    "github.com/spf13/cobra"
)

var (
    verbose bool
    quiet   bool
    cfgFile string
)

var rootCmd = &cobra.Command{
    Use:   "sourcebox",
    Short: "Generate realistic, verticalized demo data instantly",
    Long:  `SourceBox generates production-like demo data for databases...`,
    // Command logic, flag definitions
}

func Execute() error {
    return rootCmd.Execute()
}

func SetVersion(v string) {
    rootCmd.Version = v
}

func init() {
    rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
    rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet mode")
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path")
}
```

### Version Injection Mechanism

**Build-time injection**:
```bash
go build -ldflags="-X main.version=v1.0.0" ./cmd/sourcebox
```

**Variable pattern**: `-X main.version` (package.variable format)
- Package: `main` (from main.go)
- Variable: `version` (exported or unexported, both work)

**Display options**:
1. `sourcebox --version` (Cobra automatic flag)
2. `sourcebox version` (optional subcommand, add if needed later)

### Rationale

**Why minimal main.go**:
- Version variable MUST be in main package for ldflags injection
- main.go should only bootstrap, not contain logic
- Easier to test cmd package separately (no dependency on main)

**Why export Execute() from cmd package**:
- Clean separation: main.go calls cmd.Execute(), not direct rootCmd.Execute()
- Allows setting version before execution
- Standard pattern across Cobra applications

**Why SetVersion() function**:
- Cannot directly set rootCmd.Version from main package (rootCmd is private)
- Exported function provides clean API
- Allows version to be set at runtime (useful for testing)

### Alternatives Considered

**Alternative 1**: Version variable in cmd package
- Rejected: ldflags cannot inject into non-main packages reliably
- Would require: `-X github.com/jbeausoleil/sourcebox/cmd/sourcebox/cmd.version`
- Fragile, breaks if package path changes

**Alternative 2**: Hardcode version in rootCmd
- Rejected: Violates build-time injection requirement (F006 spec)
- Would require code changes for every release

**Alternative 3**: Read version from embedded file
- Rejected: More complex than ldflags, adds unnecessary I/O
- ldflags is standard Go practice for version injection

### Sources

- [Go ldflags documentation](https://pkg.go.dev/cmd/link)
- [Cobra version flag support](https://github.com/spf13/cobra#version-flag)
- [GitHub CLI version implementation](https://github.com/cli/cli/blob/trunk/cmd/gh/main.go) (reference)

---

## 3. Global vs Persistent Flags

### Decision

**PersistentFlags**: Available to command AND all subcommands
```go
rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet mode")
rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path")
```

**Local Flags**: Only for specific command
```go
seedCmd.Flags().StringP("schema", "s", "", "schema name (required)")
seedCmd.Flags().IntP("records", "n", 1000, "number of records to generate")
seedCmd.MarkFlagRequired("schema")
```

### Flag Cascade Behavior

- PersistentFlags defined on rootCmd cascade to ALL subcommands
- Local flags ONLY apply to the command they're defined on
- If local flag has same name as persistent flag, local overrides
- Access pattern: `cmd.Flags().GetString("flagname")` retrieves value with precedence

### Precedence Order (highest to lowest)

1. Local flag value (if set)
2. Persistent flag value (if set)
3. Default value (from flag definition)

### Use Cases

**Persistent Flags** (defined on rootCmd):
- `--config`: Configuration file path (affects all commands)
- `--verbose`: Increase output detail (applies to all operations)
- `--quiet`: Suppress output (applies to all operations)

**Local Flags** (command-specific):
- `seed --schema`: Schema name (only for seed command)
- `seed --records`: Record count (only for seed command)
- `seed --host`: Database host (only for seed command)

### Rationale

**Why use PersistentFlags for --verbose, --quiet, --config**:
- These flags modify behavior across ALL commands
- Users expect `sourcebox seed --verbose` and `sourcebox list-schemas --verbose` to both work
- Avoids code duplication (don't redefine --verbose on every command)

**Why use local Flags for --schema, --records**:
- These flags only make sense for specific commands
- `sourcebox list-schemas --schema=fintech` would be nonsensical
- Clearer command interface (only show relevant flags in help)

**Why allow local override**:
- Flexibility for edge cases (e.g., command-specific config file)
- Standard Cobra behavior, users expect it

### Alternatives Considered

**Alternative 1**: Everything as local flags (no persistent flags)
- Rejected: Massive code duplication, every command redefines --verbose, --quiet, --config
- User confusion (why doesn't --verbose work on this command?)

**Alternative 2**: Everything as persistent flags
- Rejected: Clutters help output with irrelevant flags
- `sourcebox list-schemas --help` would show database connection flags (nonsensical)

### Sources

- [Cobra flags documentation](https://github.com/spf13/cobra/blob/main/user_guide.md#persistent-flags)
- [Cobra flag precedence](https://github.com/spf13/cobra/blob/main/user_guide.md#local-flag-on-parent-commands)
- kubectl flag patterns (reference implementation)

---

## 4. Subcommand Scaffolding

### Decision

**Scaffold command pattern**:
```bash
cobra-cli add seed          # Creates cmd/seed.go
cobra-cli add list-schemas  # Creates cmd/list_schemas.go (hyphen → underscore)
```

**Generated boilerplate** (cmd/seed.go):
```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
    Use:   "seed",
    Short: "A brief description of your command",
    Long: `A longer description...`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("seed called")
    },
}

func init() {
    rootCmd.AddCommand(seedCmd)
    // Local flags defined here
}
```

### Customization Required

After generation, update:

1. **Use field** → Command signature with args:
   ```go
   Use: "seed <database>",  // Shows required argument
   ```

2. **Short field** → One-line description (<60 chars):
   ```go
   Short: "Seed a database with realistic demo data",
   ```

3. **Long field** → Detailed multi-paragraph description:
   ```go
   Long: `Seed a database with verticalized, production-like demo data.

   SourceBox generates realistic data based on industry-specific schemas
   (fintech, healthcare, retail) with proper relationships, distributions,
   and edge cases. Data is deterministic and reproducible.`,
   ```

4. **Example field** → Practical usage examples:
   ```go
   Example: `  # Seed MySQL with fintech loans data
     sourcebox seed mysql --schema=fintech-loans --records=1000

     # Seed Postgres with healthcare patient records
     sourcebox seed postgres --schema=healthcare-patients --records=5000`,
   ```

5. **Args field** → Argument validation:
   ```go
   Args: cobra.ExactArgs(1),  // Requires exactly 1 argument (database type)
   ```

6. **Run function** → Command implementation:
   ```go
   Run: func(cmd *cobra.Command, args []string) {
       // Phase 1 (F006): Placeholder
       fmt.Println("Seed command - implementation coming in F021")

       // Phase 2 (F021): Full implementation
       // database := args[0]
       // schema, _ := cmd.Flags().GetString("schema")
       // records, _ := cmd.Flags().GetInt("records")
       // ... seeding logic
   },
   ```

7. **Local flags** → Command-specific flags:
   ```go
   func init() {
       rootCmd.AddCommand(seedCmd)

       seedCmd.Flags().StringP("schema", "s", "", "schema name (required)")
       seedCmd.Flags().IntP("records", "n", 1000, "number of records")
       seedCmd.Flags().String("host", "localhost", "database host")
       seedCmd.MarkFlagRequired("schema")
   }
   ```

### Aliases Configuration

For commands with common short forms:
```go
var listSchemasCmd = &cobra.Command{
    Use:     "list-schemas",
    Aliases: []string{"ls"},  // Allows 'sourcebox ls'
    Short:   "List available data schemas",
}
```

### Registration Pattern

**Automatic** via init() function:
```go
func init() {
    rootCmd.AddCommand(seedCmd)  // Registers with root
}
```

- init() runs automatically when package is imported
- No manual registration needed in main.go or root.go
- Order-independent (all init() functions run before main())

### Rationale

**Why use cobra-cli add**:
- Generates correct boilerplate (reduces errors)
- Ensures consistent structure across commands
- Faster than writing from scratch

**Why customize after generation**:
- Generated boilerplate is intentionally minimal
- Real-world commands need clear help text, examples, validation
- Implementation comes in feature-specific phases (F021, F022)

**Why use init() for registration**:
- Standard Go pattern for package initialization
- Declarative (just import the command, registration happens automatically)
- No risk of forgetting to register command

### Alternatives Considered

**Alternative 1**: Manual command file creation (no cobra-cli)
- Rejected: Error-prone, easy to forget boilerplate, inconsistent structure

**Alternative 2**: Manual registration in root.go
- Rejected: Requires modifying root.go for every new command, scales poorly

**Alternative 3**: Dynamic command discovery (reflect package)
- Rejected: Over-engineering, adds complexity, no clear benefit over init()

### Sources

- [cobra-cli add documentation](https://github.com/spf13/cobra-cli#add-commands-to-a-project)
- [Cobra command structure](https://github.com/spf13/cobra/blob/main/user_guide.md#create-your-main)
- Hugo command structure (reference: extensive use of subcommands)

---

## 5. Version Injection Mechanism

### Decision

**Build flag pattern**:
```bash
go build -ldflags="-X main.version=$(git describe --tags --always --dirty)" ./cmd/sourcebox
```

**Variable declaration** (main.go):
```go
var version = "dev"  // Default for local builds
```

**Makefile integration** (update F004 Makefile):
```makefile
VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"

.PHONY: build
build: ## Build binary for current platform
	@echo "Building sourcebox $(VERSION)..."
	@mkdir -p dist
	@go build $(LDFLAGS) -o dist/sourcebox ./cmd/sourcebox
```

### Version String Format

**Tagged release** (e.g., v1.0.0):
```
$ git describe --tags --always --dirty
v1.0.0
```

**Commits after tag** (e.g., 3 commits after v1.0.0):
```
$ git describe --tags --always --dirty
v1.0.0-3-g5e8f2a1
```

**No tags** (falls back to commit SHA):
```
$ git describe --tags --always --dirty
5e8f2a1
```

**Uncommitted changes** (adds -dirty suffix):
```
$ git describe --tags --always --dirty
v1.0.0-dirty
```

### Cobra Integration

**Set version at runtime** (main.go):
```go
func main() {
    cmd.SetVersion(version)  // Pass to Cobra before Execute()
    cmd.Execute()
}
```

**Export SetVersion** (cmd/root.go):
```go
func SetVersion(v string) {
    rootCmd.Version = v  // Cobra displays this for --version
}
```

### Display Options

1. **Automatic flag**: `sourcebox --version`
   - Cobra generates this automatically when rootCmd.Version is set
   - Output: `sourcebox version {version string}`

2. **Optional subcommand**: `sourcebox version` (add later if needed)
   - Not required for F006 (MVP)
   - Can add in future if users request it

### Rationale

**Why ldflags injection**:
- Standard Go practice for version metadata
- Zero runtime overhead (compiled into binary)
- No external files needed (self-contained binary)
- Works with cross-compilation (each binary has correct version)

**Why `git describe --tags --always --dirty`**:
- Provides human-readable version (v1.0.0, not SHA)
- Falls back to SHA if no tags (always works)
- Includes commit count since tag (useful for debugging)
- Shows -dirty for uncommitted changes (critical for support)

**Why default to "dev"**:
- Clear indication this is a local build (not a release)
- Doesn't break if ldflags not provided (go run, go test)
- Easy to test (no git dependency in tests)

**Why NOT hardcode version**:
- Violates F006 spec requirement (build-time injection)
- Requires code change for every release (error-prone)
- Can't distinguish builds (all say same version)

### Alternatives Considered

**Alternative 1**: Version in VERSION file
- Rejected: Requires embedding file, adds I/O at runtime, breaks cross-compilation

**Alternative 2**: Version in go.mod
- Rejected: go.mod version is module version, not binary version (different concepts)

**Alternative 3**: Manual version updates in code
- Rejected: Error-prone, easy to forget, violates spec

### Sources

- [Go ldflags documentation](https://pkg.go.dev/cmd/link#hdr-Command_Line)
- [git describe documentation](https://git-scm.com/docs/git-describe)
- [Semantic versioning](https://semver.org/)
- GitHub CLI, Hugo, kubectl version patterns (reference implementations)

---

## 6. Help Text Quality Standards

### Decision

**Help text template**:
```go
var exampleCmd = &cobra.Command{
    Use:   "seed <database>",  // Command signature with args
    Short: "Seed a database with realistic demo data",  // <60 chars
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
        // Implementation
    },
}
```

### Field Requirements

**Use** (command signature):
- Format: `command [required-arg] <optional-arg>`
- Square brackets `[]`: Optional arguments
- Angle brackets `<>`: Required arguments (Cobra convention varies, we use `<>`)
- Example: `seed <database>` means database is required

**Short** (one-line description):
- Maximum 60 characters
- Appears in command list (`sourcebox --help`)
- No period at end (Cobra convention)
- Clear, actionable (start with verb: "Seed", "List", "Generate")

**Long** (detailed description):
- Multi-paragraph explanation
- What the command does
- When to use it
- Key concepts (what is a schema? what databases supported?)
- Prerequisites or assumptions
- Use backticks for multi-line (Go raw string literal)

**Example** (practical usage):
- 2-3 real-world examples
- Include comments explaining each example
- Show common flags and variations
- Copyable (users can paste directly)
- Start each line with 2 spaces (Cobra formatting)

### Tone Guidelines

**DO**:
- Clear, concise language
- Developer-friendly terminology
- Practical examples users can copy-paste
- Explain what, why, and when to use

**DON'T**:
- Marketing fluff ("Amazing tool!", "Revolutionary!")
- Unnecessary verbosity
- Jargon without explanation
- Vague descriptions ("Does stuff", "Makes things")

### Rationale

**Why prioritize Examples**:
- Developers learn by example (not by reading walls of text)
- Copyable examples reduce time-to-first-success
- Shows real-world usage patterns

**Why Short must be <60 chars**:
- Appears in command list with multiple commands
- Long descriptions break terminal formatting
- Forces clarity (can't be vague in 60 chars)

**Why Long can be multi-paragraph**:
- Displayed only when user requests help for specific command
- Space to explain concepts (what is verticalized data?)
- Includes prerequisites (what databases must be installed?)

### Alternatives Considered

**Alternative 1**: Minimal help text (Short only, no Long/Example)
- Rejected: Violates F006 spec (comprehensive help system required)
- Users can't self-serve learning

**Alternative 2**: External documentation (link to website)
- Rejected: Requires internet (violates local-first principle)
- Slows down users (context switch to browser)

**Alternative 3**: Interactive help (prompts, wizards)
- Rejected: Over-engineering for MVP, adds complexity
- Developers prefer static help (faster, scriptable)

### Sources

- [Cobra help text best practices](https://github.com/spf13/cobra/blob/main/user_guide.md#help-command)
- kubectl help text (gold standard for CLI help)
- [GitHub CLI help patterns](https://cli.github.com/manual/) (excellent examples)

---

## 7. Flag Validation Patterns

### Decision

**Required flags**:
```go
func init() {
    seedCmd.Flags().StringP("schema", "s", "", "schema name (required)")
    seedCmd.MarkFlagRequired("schema")  // Cobra validates automatically
}
```

**Mutually exclusive flags** (custom validation):
```go
var seedCmd = &cobra.Command{
    PreRunE: func(cmd *cobra.Command, args []string) error {
        verbose, _ := cmd.Flags().GetBool("verbose")
        quiet, _ := cmd.Flags().GetBool("quiet")

        if verbose && quiet {
            return fmt.Errorf("--verbose and --quiet are mutually exclusive")
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {
        // Main logic runs only if PreRunE succeeds
    },
}
```

**Value validation** (in Run function):
```go
Run: func(cmd *cobra.Command, args []string) {
    database := args[0]
    if database != "mysql" && database != "postgres" {
        return fmt.Errorf("unsupported database: %s (supported: mysql, postgres)", database)
    }

    records, _ := cmd.Flags().GetInt("records")
    if records < 1 || records > 1000000 {
        return fmt.Errorf("records must be between 1 and 1,000,000 (got: %d)", records)
    }

    // Valid, proceed with logic
}
```

### Error Message Template

```
❌ Error: {what went wrong}
💡 Fix: {how to fix it}

Examples:
  {correct usage example 1}
  {correct usage example 2}
```

**Example**:
```
❌ Error: --verbose and --quiet are mutually exclusive
💡 Fix: Use only one output mode

Examples:
  sourcebox seed mysql --schema=fintech --verbose
  sourcebox seed mysql --schema=fintech --quiet
```

### Flag Precedence

When conflicting flags are provided:

1. **Document precedence in help text**:
   ```go
   Long: `...

   Note: If both --verbose and --quiet are specified, --quiet takes precedence.`
   ```

2. **Implement precedence in code**:
   ```go
   if quiet {
       // Quiet mode overrides verbose
       setLogLevel("error")
   } else if verbose {
       setLogLevel("debug")
   } else {
       setLogLevel("info")  // Default
   }
   ```

### Rationale

**Why use MarkFlagRequired**:
- Cobra validates automatically (no manual checking needed)
- Clear error message ("required flag(s) \"schema\" not set")
- Prevents invalid command execution

**Why use PreRunE for mutual exclusivity**:
- Runs before main logic (fail fast)
- Returns error (Cobra handles error display + exit code)
- Keeps validation separate from business logic

**Why document precedence**:
- Users will provide conflicting flags (typos, scripts)
- Clear precedence prevents confusion ("which flag wins?")
- Predictable behavior (scripts can rely on it)

### Alternatives Considered

**Alternative 1**: No validation (trust users)
- Rejected: Results in cryptic errors deep in business logic
- Violates "fail gracefully" UX principle

**Alternative 2**: Validation in Run function only
- Rejected: Mixes validation with business logic (harder to maintain)
- No early exit (could do partial work before failing)

**Alternative 3**: Reject conflicting flags entirely
- Rejected: Too strict, users make mistakes, better to have predictable precedence

### Sources

- [Cobra flag validation](https://github.com/spf13/cobra/blob/main/user_guide.md#positional-and-custom-arguments)
- [Cobra PreRun hooks](https://github.com/spf13/cobra/blob/main/user_guide.md#prerun-and-postrun-hooks)
- kubectl validation patterns (reference implementation)

---

## 8. Command Testing Strategies

### Decision

**Test categories**:

1. **Help output tests**:
```go
func TestSeedHelpOutput(t *testing.T) {
    buf := new(bytes.Buffer)
    rootCmd.SetOut(buf)
    rootCmd.SetErr(buf)
    rootCmd.SetArgs([]string{"seed", "--help"})

    err := rootCmd.Execute()
    require.NoError(t, err)

    output := buf.String()
    assert.Contains(t, output, "Seed a database")
    assert.Contains(t, output, "--schema")
    assert.Contains(t, output, "fintech-loans")  // Example in help
}
```

2. **Flag parsing tests**:
```go
func TestSeedFlagParsing(t *testing.T) {
    tests := []struct {
        name    string
        args    []string
        wantErr bool
        errMsg  string
    }{
        {
            name:    "valid flags",
            args:    []string{"seed", "mysql", "--schema=fintech-loans", "--records=1000"},
            wantErr: false,
        },
        {
            name:    "missing required schema flag",
            args:    []string{"seed", "mysql"},
            wantErr: true,
            errMsg:  "required flag(s) \"schema\" not set",
        },
        {
            name:    "conflicting verbose and quiet",
            args:    []string{"seed", "mysql", "--schema=fintech", "--verbose", "--quiet"},
            wantErr: true,
            errMsg:  "mutually exclusive",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            rootCmd.SetArgs(tt.args)
            err := rootCmd.Execute()

            if tt.wantErr {
                require.Error(t, err)
                assert.Contains(t, err.Error(), tt.errMsg)
            } else {
                require.NoError(t, err)
            }
        })
    }
}
```

3. **Command registration tests**:
```go
func TestCommandsRegistered(t *testing.T) {
    commands := rootCmd.Commands()
    commandNames := make([]string, len(commands))
    for i, cmd := range commands {
        commandNames[i] = cmd.Name()
    }

    assert.Contains(t, commandNames, "seed")
    assert.Contains(t, commandNames, "list-schemas")
}
```

4. **Version display tests**:
```go
func TestVersionFlag(t *testing.T) {
    buf := new(bytes.Buffer)
    rootCmd.SetOut(buf)
    rootCmd.SetArgs([]string{"--version"})

    SetVersion("v1.0.0-test")

    err := rootCmd.Execute()
    require.NoError(t, err)
    assert.Contains(t, buf.String(), "v1.0.0-test")
}
```

### Output Capture Pattern

```go
func executeCommand(args []string) (output string, err error) {
    buf := new(bytes.Buffer)
    rootCmd.SetOut(buf)
    rootCmd.SetErr(buf)
    rootCmd.SetArgs(args)

    err = rootCmd.Execute()
    return buf.String(), err
}

// Usage in tests
output, err := executeCommand([]string{"seed", "--help"})
assert.NoError(t, err)
assert.Contains(t, output, "expected text")
```

### Coverage Target

- **Overall cmd/ package**: >80% coverage
- **Command Run functions**: 100% (TDD required for business logic)
- **Help text generation**: Not counted (static text, no logic)
- **init() functions**: Not counted (registration, no business logic)

### Rationale

**Why test help output**:
- Help is the primary discovery mechanism (users learn via help)
- Changes to help text are breaking changes (documentation)
- Ensures examples stay accurate as code evolves

**Why use table-driven tests**:
- Many flag combinations to test (exponential)
- Easy to add new test cases (just append to table)
- Clear test names (self-documenting)

**Why test command registration**:
- Catches typos in AddCommand() calls
- Ensures init() functions actually run
- Fast test (no execution, just introspection)

**Why NOT test Cobra internals**:
- Cobra itself is well-tested (don't re-test framework)
- Focus on OUR logic (validation, help text, business logic)

### Alternatives Considered

**Alternative 1**: No command tests (rely on manual testing)
- Rejected: Violates TDD requirement, makes refactoring risky

**Alternative 2**: Integration tests only (execute real commands)
- Rejected: Slow, hard to test edge cases, requires full setup

**Alternative 3**: Mock Cobra framework
- Rejected: Over-engineering, Cobra is stable, no need to mock

### Sources

- [Testing Cobra commands](https://github.com/spf13/cobra/blob/main/command_test.go) (Cobra's own tests)
- [Table-driven tests in Go](https://go.dev/wiki/TableDrivenTests)
- kubectl test patterns (extensive Cobra testing examples)

---

## 9. Error Handling Standards

### Decision

**Error return pattern** (use RunE, not Run):
```go
var seedCmd = &cobra.Command{
    Use:   "seed <database>",
    RunE: func(cmd *cobra.Command, args []string) error {  // RunE returns error
        database := args[0]

        if err := validateDatabase(database); err != nil {
            return fmt.Errorf("validation failed: %w", err)  // Wrap with context
        }

        if err := seedDatabase(database); err != nil {
            return fmt.Errorf("seeding failed: %w", err)
        }

        return nil  // Success
    },
    SilenceUsage: true,  // Don't print usage on errors
}
```

### Exit Codes

- **0**: Success
- **1**: General error (Cobra default for returned errors)
- **Custom codes**: Call `os.Exit(code)` directly if needed

```go
// Cobra automatic (recommended)
func main() {
    if err := cmd.Execute(); err != nil {
        os.Exit(1)  // All errors → exit code 1
    }
}

// Custom exit codes (only if needed)
func main() {
    if err := cmd.Execute(); err != nil {
        if errors.Is(err, ErrDatabaseConnection) {
            os.Exit(2)  // Database connection errors → exit code 2
        }
        os.Exit(1)  // All other errors → exit code 1
    }
}
```

### Error Message Format

**Template**:
```
Error: {clear description of what went wrong}
{optional: additional context}
```

**Examples**:
```
Error: unsupported database: mongodb (supported: mysql, postgres)

Error: failed to connect to mysql://localhost:3306
Connection refused - ensure MySQL is running

Error: schema not found: invalid-schema
Available schemas: fintech-loans, healthcare-patients, retail-orders
```

**User-facing errors** (avoid technical jargon):
```go
// Good
return fmt.Errorf("database connection failed: ensure MySQL is running on port 3306")

// Bad
return fmt.Errorf("dial tcp 127.0.0.1:3306: connect: connection refused")
```

### Error Wrapping

Use `%w` for error wrapping (preserves error chain):
```go
if err := db.Connect(); err != nil {
    return fmt.Errorf("failed to connect to database: %w", err)
}

// Later, can check with errors.Is()
if errors.Is(err, sql.ErrConnDone) {
    // Handle connection closed error
}
```

### Silence Usage

Set `SilenceUsage: true` to prevent printing full command usage on errors:

```go
var seedCmd = &cobra.Command{
    Use:          "seed <database>",
    SilenceUsage: true,  // Don't print usage on business logic errors
    RunE: func(cmd *cobra.Command, args []string) error {
        // Business logic errors don't print usage
        return fmt.Errorf("seeding failed: %w", err)
    },
}
```

**Why**: Usage should only print for syntax errors (wrong flags, missing args), not business logic errors (database connection failed).

### Rationale

**Why use RunE instead of Run**:
- Cobra handles error display + exit code automatically
- No need to print errors manually
- Consistent error formatting across all commands

**Why wrap errors with %w**:
- Preserves error chain (can check with errors.Is/As)
- Adds context ("failed to connect" vs raw "connection refused")
- Stack trace preserved for debugging

**Why use SilenceUsage**:
- Usage is for syntax errors, not runtime errors
- Database connection failure doesn't mean user needs to see full usage
- Cleaner output (error message only, not 50 lines of help)

**Why default to exit code 1**:
- Simple, standard across Unix tools
- Custom codes only if users need to distinguish error types in scripts

### Alternatives Considered

**Alternative 1**: Print errors manually in Run function
- Rejected: Inconsistent formatting, easy to forget, duplicates Cobra functionality

**Alternative 2**: Panic on errors
- Rejected: Not idiomatic Go, makes testing harder, prevents graceful recovery

**Alternative 3**: Complex exit code system
- Rejected: Over-engineering, users rarely need fine-grained exit codes

### Sources

- [Cobra error handling](https://github.com/spf13/cobra/blob/main/user_guide.md#suggestions-when-unknown-command-happens)
- [Go error wrapping](https://go.dev/blog/go1.13-errors)
- [Exit codes best practices](https://tldp.org/LDP/abs/html/exitcodes.html)

---

## 10. Shell Completion Integration

### Decision

**Cobra automatic support**: Shell completions are generated automatically by Cobra when `rootCmd.Version` is set.

**Supported shells**:
- bash
- zsh
- fish
- PowerShell

**Generation commands** (add later if needed):
```bash
# Generate bash completion script
sourcebox completion bash

# Generate zsh completion script
sourcebox completion zsh

# Generate fish completion script
sourcebox completion fish

# Generate PowerShell completion script
sourcebox completion powershell
```

**Installation** (user workflow):
```bash
# Bash (add to ~/.bashrc or ~/.bash_profile)
sourcebox completion bash > /usr/local/etc/bash_completion.d/sourcebox

# Zsh (add to ~/.zshrc)
sourcebox completion zsh > "${fpath[1]}/_sourcebox"

# Fish (add to ~/.config/fish/completions/)
sourcebox completion fish > ~/.config/fish/completions/sourcebox.fish
```

### Current Scope (F006 - MVP)

**NOT required for F006**:
- No completion subcommand needed
- No installation instructions needed
- Cobra provides basic completion automatically

**Add later** (F007 or based on user demand):
- `sourcebox completion` subcommand
- Installation documentation
- Custom completion logic (dynamic flag values, schema names)

### Future Enhancement Possibilities

**Static completions** (Cobra automatic):
- Command names (seed, list-schemas)
- Flag names (--schema, --records, --verbose)

**Dynamic completions** (custom implementation):
- Schema names (complete `--schema=` with available schemas)
- Database types (complete `seed <tab>` with mysql, postgres)
- Host completion (read from known hosts file)

### Rationale

**Why defer to Phase 2**:
- Not required for F006 MVP (foundation only)
- Basic Cobra completion already works (command/flag names)
- Dynamic completions require schema discovery (implemented in F022)
- Time-boxed feature (12-week deadline, prioritize core value)

**Why document for future**:
- Users will request this feature (convenience improvement)
- Cobra makes it easy to add later (minimal implementation)
- Shows we've considered UX completeness

**Why NOT implement now**:
- Doesn't affect core functionality (data generation)
- Completion is nice-to-have, not must-have for MVP
- Can validate user demand first (how many users request it?)

### Alternatives Considered

**Alternative 1**: Implement full completion in F006
- Rejected: Scope creep, delays MVP, not required for validation

**Alternative 2**: Never implement completion
- Rejected: Professional tools have completion, users expect it

**Alternative 3**: Use external completion tool (bash-completion, zsh-completions)
- Rejected: Cobra provides this out-of-box, no need for external tools

### Sources

- [Cobra completion documentation](https://github.com/spf13/cobra/blob/main/shell_completions.md)
- [Bash completion guide](https://github.com/scop/bash-completion)
- kubectl completion (reference: excellent dynamic completion)

---

## Summary

All 10 research decisions documented with:
- ✅ Decision: What was chosen (Cobra patterns, flag types, testing strategies)
- ✅ Rationale: Why chosen (constitutional alignment, best practices, developer UX)
- ✅ Alternatives considered: What else was evaluated and why rejected
- ✅ Sources: Where information came from (Cobra docs, Go docs, reference implementations)

**Next Phase**: Generate design artifacts (quickstart.md) based on these research decisions.

---

**Constitutional Compliance Check**:
- ✅ Boring Tech: Cobra is proven, standard Go CLI framework (used by kubectl, Hugo, gh)
- ✅ Speed: <100ms response time for help/version (statically compiled, no I/O)
- ✅ Developer-First: Comprehensive help, clear errors, zero config
- ✅ Code Quality: TDD patterns documented, >80% coverage target
- ✅ Platform Support: Cross-platform compatible (Cobra works everywhere)
- ✅ Local-First: No network dependencies (all local operations)
- ✅ Ship Fast: Standard patterns (no reinventing the wheel, use Cobra conventions)

**Violations**: None. All decisions align with SourceBox constitution.
