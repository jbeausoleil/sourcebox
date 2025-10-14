# Feature Specification Prompt: F006 - Cobra CLI Framework Integration

## Feature Metadata
- **Feature ID**: F006
- **Name**: Cobra CLI Framework Integration
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F004 (Project directory structure)

## Constitutional Alignment

### Core Principles
- ✅ **Boring Tech Wins**: Cobra is the standard Go CLI framework (used by kubectl, Hugo, GitHub CLI)
- ✅ **Developer-First Design**: CLI-first approach, not web UI
- ✅ **Zero Config**: Works out of the box with sane defaults

### Technical Constraints
- ✅ **Platform Support**: Cobra works on all platforms
- ✅ **Code Quality**: Standard Go conventions

### Development Practices
- ✅ **Simple > Complex**: Use proven framework, don't reinvent

## User Story
**US-MVP-003**: "As a developer, I want a well-structured CLI with intuitive commands, helpful flags, and comprehensive help text so I can quickly learn and use SourceBox."

## Problem Statement
SourceBox needs a robust command-line interface with subcommands (`seed`, `list-schemas`, `version`), global and local flags, comprehensive help system, and user-friendly error messages. Building this from scratch is time-consuming and error-prone. Cobra provides battle-tested CLI infrastructure used by major projects like Kubernetes, allowing SourceBox to focus on core functionality rather than CLI plumbing.

## Solution Overview
Integrate the Cobra CLI framework to provide command structure, flag parsing, help generation, and shell completions. Initialize Cobra with `cobra-cli init`, create the root command with global flags (`--verbose`, `--quiet`, `--config`), scaffold subcommands (`seed`, `list-schemas`), and implement the version command with build-time version injection.

## Detailed Requirements

### Acceptance Criteria
1. **Cobra CLI Initialized**: `cobra-cli init` executed successfully
2. **Root Command Created**: Main command structure in `/cmd/sourcebox/`
3. **Help System Working**: `sourcebox --help` shows comprehensive help
4. **Version Command**: `sourcebox --version` or `sourcebox version` shows version info
5. **Global Flags Implemented**:
   - `--verbose` or `-v`: Enable verbose output
   - `--quiet` or `-q`: Suppress non-error output
   - `--config <file>`: Use custom config file (future use)
6. **Subcommands Scaffolded**:
   - `sourcebox seed` (implementation in F021)
   - `sourcebox list-schemas` (implementation in F022)
7. **Help Text Quality**: Clear descriptions, usage examples, flag documentation

### Technical Specifications

#### Install Cobra CLI
```bash
go install github.com/spf13/cobra-cli@latest
```

#### Initialize Cobra Structure
```bash
cd sourcebox
cobra-cli init
```

#### Root Command: `cmd/sourcebox/main.go`

```go
package main

import (
	"os"
	"github.com/yourusername/sourcebox/cmd/sourcebox/cmd"
)

var version = "dev" // Injected at build time

func main() {
	cmd.SetVersion(version)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
```

#### Root Command: `cmd/sourcebox/cmd/root.go`

```go
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
	quiet   bool
	version string
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "sourcebox",
	Short: "Generate realistic, verticalized demo data instantly",
	Long: `SourceBox generates production-like demo data for fintech, healthcare,
retail, and more. Seed databases in <30 seconds or spin up pre-loaded
Docker containers in <10 seconds.

Perfect for:
  • Demo environments and presentations
  • Local development with realistic data
  • Sales engineers building proof-of-concepts
  • QA and testing with production-like scenarios`,
	Version: version,
}

// Execute adds all child commands to the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sourcebox.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress non-error output")
}

// SetVersion sets the version string (injected at build time)
func SetVersion(v string) {
	version = v
	rootCmd.Version = v
}

// IsVerbose returns whether verbose mode is enabled
func IsVerbose() bool {
	return verbose
}

// IsQuiet returns whether quiet mode is enabled
func IsQuiet() bool {
	return quiet
}
```

#### Subcommand Scaffolding: `seed` command

```bash
cobra-cli add seed
```

Creates: `cmd/sourcebox/cmd/seed.go`

```go
package cmd

import (
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed [database]",
	Short: "Seed a database with realistic data",
	Long: `Seed a MySQL or PostgreSQL database with verticalized demo data.

Examples:
  # Seed MySQL with fintech loans data (1000 records)
  sourcebox seed mysql --schema=fintech-loans --records=1000

  # Seed PostgreSQL with healthcare data
  sourcebox seed postgres --schema=healthcare-patients --host=localhost --port=5432

  # Export to SQL file instead of direct insertion
  sourcebox seed mysql --schema=retail-ecommerce --output=seed.sql`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation in F021
		cmd.Println("Seed command - implementation coming in F021")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Local flags (specific to seed command)
	seedCmd.Flags().String("schema", "", "schema name (required)")
	seedCmd.Flags().Int("records", 1000, "number of records to generate")
	seedCmd.Flags().String("host", "localhost", "database host")
	seedCmd.Flags().Int("port", 0, "database port (default: 3306 for MySQL, 5432 for Postgres)")
	seedCmd.Flags().String("user", "root", "database user")
	seedCmd.Flags().String("password", "", "database password")
	seedCmd.Flags().String("db-name", "demo", "database name")
	seedCmd.Flags().String("output", "", "export to SQL file instead of inserting")
	seedCmd.Flags().Bool("dry-run", false, "show what would be done without executing")

	// Required flags
	seedCmd.MarkFlagRequired("schema")
}
```

#### Subcommand Scaffolding: `list-schemas` command

```bash
cobra-cli add list-schemas
```

Creates: `cmd/sourcebox/cmd/list_schemas.go`

```go
package cmd

import (
	"github.com/spf13/cobra"
)

var listSchemasCmd = &cobra.Command{
	Use:     "list-schemas",
	Aliases: []string{"ls"},
	Short:   "List all available schemas",
	Long: `List all available data schemas with descriptions and statistics.

Examples:
  # List all schemas
  sourcebox list-schemas

  # Short alias
  sourcebox ls`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation in F022
		cmd.Println("List schemas command - implementation coming in F022")
	},
}

func init() {
	rootCmd.AddCommand(listSchemasCmd)
}
```

#### Version Information

Version is injected at build time via ldflags (configured in F004 Makefile):
```bash
go build -ldflags="-X main.version=v1.0.0" ./cmd/sourcebox
```

User can see version with:
```bash
sourcebox --version
# or
sourcebox version
```

### Performance Considerations
- Cobra adds minimal overhead (<1ms command parsing)
- Help text generation is lazy (only computed when --help is used)
- Flag parsing is efficient

### Testing Strategy

**Unit Tests: `cmd/sourcebox/cmd/root_test.go`**
```go
package cmd

import (
	"bytes"
	"testing"
)

func TestRootCommand(t *testing.T) {
	// Test help output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("root command failed: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("SourceBox")) {
		t.Error("help text should contain 'SourceBox'")
	}
}

func TestVersionFlag(t *testing.T) {
	SetVersion("v1.2.3-test")

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"--version"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("version command failed: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("v1.2.3-test")) {
		t.Errorf("version output should contain 'v1.2.3-test', got: %s", output)
	}
}

func TestGlobalFlags(t *testing.T) {
	rootCmd.SetArgs([]string{"--verbose", "--quiet"})
	err := rootCmd.Execute()

	// Should succeed even with conflicting flags (implementation decides precedence)
	if err != nil {
		t.Fatalf("global flags parsing failed: %v", err)
	}
}
```

**Manual Testing**:
```bash
# Test help system
sourcebox --help
sourcebox seed --help
sourcebox list-schemas --help

# Test version
sourcebox --version

# Test subcommands exist (will show placeholder)
sourcebox seed mysql --schema=fintech-loans
sourcebox list-schemas

# Test flags parse correctly (no errors)
sourcebox seed mysql --schema=test --verbose --dry-run
```

## Dependencies
- **Upstream**: F004 (Directory structure must exist)
- **Downstream**:
  - F021 (Seed command implementation)
  - F022 (List-schemas command implementation)
  - F027 (CLI output formatting)

## Deliverables
1. Cobra framework integrated and initialized
2. Root command with global flags
3. Version command with build-time injection
4. Scaffolded `seed` and `list-schemas` subcommands
5. Comprehensive help text
6. Unit tests for root command and flags
7. Documentation of CLI structure

## Success Criteria
- ✅ `sourcebox --help` shows clear, comprehensive help
- ✅ `sourcebox --version` shows correct version
- ✅ Global flags (`--verbose`, `--quiet`, `--config`) work
- ✅ Subcommands are registered and show help
- ✅ Command structure is extensible for future commands
- ✅ Tests pass: `go test ./cmd/sourcebox/cmd/...`

## Anti-Patterns to Avoid
- ❌ Reinventing CLI framework (use Cobra, it's battle-tested)
- ❌ Poor help text (should be clear and include examples)
- ❌ Inconsistent flag naming (follow Cobra conventions)
- ❌ Missing version command (users need to know which version they're running)
- ❌ No tests for CLI commands (makes refactoring risky)
- ❌ Hardcoded version string (should be injected at build time)

## Implementation Notes
- Cobra is the de facto standard for Go CLIs (kubectl, hugo, gh all use it)
- Use `cobra-cli` generator for consistent command structure
- Keep command logic in `cmd/` minimal - delegate to `/pkg` for business logic
- Follow Cobra naming conventions: Use/Short/Long/Example
- Aliases are useful: `ls` is shorter than `list-schemas`

## TDD Requirements
**Required for this feature** - Core CLI infrastructure needs tests:

1. **Test root command help**: Verify help text renders correctly
2. **Test version injection**: Verify version string is displayed
3. **Test flag parsing**: Verify global flags work
4. **Test subcommand registration**: Verify commands are properly added

Follow TDD workflow:
1. Write test (e.g., `TestRootCommand`)
2. Run test - should fail (command not implemented)
3. Implement command
4. Run test - should pass
5. Refactor if needed

## Related Constitution Sections
- **Boring Tech Wins (Principle IV)**: Cobra is proven, standard Go CLI framework
- **Developer-First Design (Principle VI)**: CLI-first, not web UI
- **Zero Config (UX Principle 2)**: Works out of the box with sane defaults
- **TDD Required (Development Practice 1)**: Core CLI logic must be tested
- **Code Quality Standards (Technical Constraint 5)**: Unit tests required
