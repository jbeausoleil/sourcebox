# Implementation Planning Prompt: F006 - Cobra CLI Framework Integration

## Feature Metadata
- **Feature ID**: F006
- **Name**: Cobra CLI Framework Integration
- **Feature Branch**: `004-f006-cobra-cli`
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F004 (Project directory structure and build system)
- **Spec Location**: `specs/004-f006-cobra-cli/spec.md`

## Constitutional Alignment

### Core Principles Verification
- ✅ **Verticalized > Generic**: N/A (CLI infrastructure, not data)
- ✅ **Speed > Features**: Command parsing <100ms, help generation <1s, zero startup delay
- ✅ **Local-First, Cloud Optional**: Pure local CLI tool, no network dependencies
- ✅ **Boring Tech Wins**: Cobra is the de facto standard (kubectl, Hugo, gh all use it)
- ✅ **Open Source Forever**: Cobra is MIT licensed, integrates seamlessly with SourceBox
- ✅ **Developer-First Design**: CLI-first tool, intuitive commands, comprehensive help, clear errors
- ✅ **Ship Fast, Validate Early**: Battle-tested framework accelerates development, TDD required

### Technical Constraints Verification
- ✅ **Performance**: <100ms response time for help/version, <1ms parsing overhead
- ✅ **Distribution**: Cobra compiles into single Go binary (no additional dependencies)
- ✅ **Database Support**: N/A (CLI framework)
- ✅ **Cost**: $0 (Cobra is free, open source)
- ✅ **Code Quality**: Unit tests required for all commands, >80% coverage target
- ✅ **License**: Cobra is MIT (compatible with SourceBox MIT license)
- ✅ **Platform Support**: Cobra works on all platforms (macOS, Linux, Windows)

### Legal Constraints Verification (CRITICAL)
- ✅ **Independent Development**: Cobra is public, standard Go framework
- ✅ **No Employer References**: N/A (CLI infrastructure)
- ✅ **Public Information Only**: All patterns from public Cobra documentation
- ✅ **Open Source Protection**: Cobra MIT license, no proprietary dependencies
- ✅ **Illustrative Examples Only**: N/A (no company references)

## Planning Context

### Feature Summary
Integrate the Cobra CLI framework to provide professional command structure, flag parsing, help generation, and shell completions for SourceBox. Initialize Cobra with `cobra-cli init`, create root command with global flags (`--verbose`, `--quiet`, `--config`), scaffold subcommands (`seed`, `list-schemas`), implement version command with build-time injection, and establish comprehensive help system. Foundation for all CLI functionality in F021 (Seed command) and F022 (List-schemas command).

### Key Technical Decisions Required

**Phase 0 Research Topics**:
1. **Cobra Initialization Pattern**: What is the standard `cobra-cli init` structure? How does it organize cmd/ directory? What files are generated?
2. **Root Command Architecture**: How to structure main.go vs cmd/root.go? What belongs in each file? How is version injection handled?
3. **Global vs Persistent Flags**: When to use PersistentFlags vs local Flags? How do flags cascade to subcommands? What's the precedence order?
4. **Subcommand Registration**: How to use `cobra-cli add` to scaffold commands? What's the init() pattern for AddCommand? How are aliases configured?
5. **Version Injection Mechanism**: How to inject version at build time via ldflags? What variable name pattern (-X main.version)? How to display version?
6. **Help Text Best Practices**: What fields are required (Use, Short, Long, Example)? How to write clear descriptions? What examples to include?
7. **Flag Validation Patterns**: How to mark required flags? How to validate flag values? How to provide clear error messages for invalid flags?
8. **Command Testing Strategies**: How to test command execution? How to capture output (stdout/stderr)? How to test flag parsing? Table-driven test patterns?
9. **Error Handling Standards**: How to return errors from Run functions? How to customize error messages? How to set exit codes?
10. **Shell Completion Integration**: How does Cobra generate completions? Do we need to configure anything? Is it automatic for bash/zsh/fish?

### Technical Context (Pre-filled)

**Language/Version**: Go 1.21+ (existing project configuration from F003/F004)
**Primary Dependencies**:
  - github.com/spf13/cobra (CLI framework, ~v1.8+)
  - github.com/spf13/cobra-cli (scaffolding tool, ~v1.3+)
**Storage**: N/A (CLI framework, no database)
**Testing**: Table-driven unit tests, output buffer capture, flag parsing verification
**Target Platform**: Cross-platform CLI (macOS, Linux, Windows)
**Project Type**: Single binary CLI tool
**Performance Goals**:
  - Command parsing: <1ms overhead
  - Help display: <1 second response time
  - Version display: <100ms response time
  - Zero startup delay (statically compiled binary)
**Constraints**:
  - Must work out of box with zero configuration
  - Must provide comprehensive help for all commands
  - Must support both long (--verbose) and short (-v) flags
  - Must integrate with existing build system (F004 Makefile)
  - Version must be injected at build time (not hardcoded)
  - TDD required for all command logic
  - Must be cross-platform compatible
**Scale/Scope**: Foundation for all CLI functionality (weeks 4-12), every feature will extend this structure

## Planning Workflow

### Phase 0: Research & Technical Decisions

Generate `research.md` with documented decisions for:

#### 1. Cobra Initialization Structure
- **Decision Point**: What's the standard directory structure after `cobra-cli init`? What files are created? What's the purpose of each?
- **Research**: Cobra documentation, example projects using Cobra
- **Output**: Directory structure explanation:
  ```
  cmd/sourcebox/
  ├── main.go (entry point, version injection)
  └── cmd/
      └── root.go (root command definition, global flags)
  ```
  - main.go: Entry point, version variable, calls cmd.Execute()
  - cmd/root.go: Root command definition, global flags, Execute() function
  - Future commands: cmd/seed.go, cmd/list_schemas.go (added with `cobra-cli add`)

#### 2. Root Command Architecture
- **Decision Point**: How to split responsibilities between main.go and cmd/root.go? How does version injection work?
- **Research**: Cobra best practices, version injection patterns
- **Output**:
  - **main.go**: Minimal - declares version variable, calls cmd.Execute(), handles exit codes
  - **cmd/root.go**: Root command definition, flag bindings, helper functions (IsVerbose, IsQuiet)
  - **Version injection**: Build flag `-ldflags="-X main.version=v1.0.0"` sets version variable
  - **SetVersion function**: Exported from cmd package to set rootCmd.Version at runtime

#### 3. Global vs Persistent Flags
- **Decision Point**: PersistentFlags vs Flags - when to use each? How do they cascade?
- **Research**: Cobra flag documentation, flag precedence
- **Output**:
  - **PersistentFlags**: Flags available to command AND all subcommands (--verbose, --quiet, --config)
  - **Local Flags**: Flags only for specific command (--records for seed command)
  - **Precedence**: Local flags override persistent flags if same name
  - **Access pattern**: Define on rootCmd.PersistentFlags() for global flags
  - Use cases:
    - Global config: --config (affects all commands)
    - Output control: --verbose, --quiet (applies to all command output)
    - Command-specific: --schema, --records (only for seed command)

#### 4. Subcommand Scaffolding
- **Decision Point**: How to use `cobra-cli add` to create new commands? What's generated? How to customize?
- **Research**: cobra-cli documentation, generated command structure
- **Output**:
  - **Command**: `cobra-cli add <commandname>` (e.g., `cobra-cli add seed`)
  - **Generated file**: `cmd/<commandname>.go` with command boilerplate
  - **Customization needed**:
    - Update Use field (command signature: "seed [database]")
    - Update Short and Long descriptions
    - Add practical Examples
    - Configure Args (cobra.ExactArgs, cobra.MinimumNArgs, cobra.NoArgs)
    - Add local flags (seedCmd.Flags().String(...))
    - Implement Run function (placeholder initially, full implementation in F021)
  - **Aliases**: Configure in Use field (e.g., `Aliases: []string{"ls"}`)
  - **Registration**: Automatic via init() function calling rootCmd.AddCommand(cmd)

#### 5. Version Injection Mechanism
- **Decision Point**: How to inject version at build time? What variable pattern? How to display?
- **Research**: Go build ldflags, Cobra version support
- **Output**:
  - **Build flag**: `-ldflags="-X main.version=$(git describe --tags --always --dirty)"`
  - **Variable declaration**: `var version = "dev"` in main.go (default for local builds)
  - **Cobra integration**: `rootCmd.Version = version` in cmd/root.go
  - **Display options**:
    - `sourcebox --version` (flag, Cobra automatic)
    - `sourcebox version` (subcommand, add if needed in future)
  - **Makefile integration**: Update F004 Makefile to include version ldflags

#### 6. Help Text Quality Standards
- **Decision Point**: What fields are required? How to write clear, useful help? What examples to include?
- **Research**: Cobra help generation, best practices from kubectl, gh CLI
- **Output**: Help text template:
  ```go
  var exampleCmd = &cobra.Command{
      Use:   "command [args]",       // Command signature
      Short: "Brief one-line description",
      Long: `Detailed multi-line description.

  Explain what the command does, when to use it, and key concepts.
  Include any important prerequisites or assumptions.`,
      Example: `  # Example 1: Common use case
    sourcebox command --flag=value

    # Example 2: Advanced use case
    sourcebox command --flag1=val1 --flag2=val2`,
      Args: cobra.ExactArgs(1),      // Argument validation
      Run: func(cmd *cobra.Command, args []string) {
          // Implementation
      },
  }
  ```
  - **Use**: Command name + argument placeholders ([required] <optional>)
  - **Short**: Single sentence (<60 chars), appears in command list
  - **Long**: Detailed explanation, use cases, prerequisites
  - **Example**: 2-3 practical examples with comments
  - **Tone**: Clear, concise, developer-friendly (no marketing fluff)

#### 7. Flag Validation Patterns
- **Decision Point**: How to validate required flags? How to provide clear errors? Best practices?
- **Research**: Cobra flag validation, custom validation functions
- **Output**:
  - **Required flags**: `cmd.MarkFlagRequired("flagname")` in init()
  - **Mutual exclusivity**: Custom PreRunE validation:
    ```go
    PreRunE: func(cmd *cobra.Command, args []string) error {
        if verbose && quiet {
            return errors.New("--verbose and --quiet are mutually exclusive")
        }
        return nil
    }
    ```
  - **Value validation**: Use flag binding + validation in Run function
  - **Error messages**: Clear, actionable (what's wrong + how to fix)
  - **Flag precedence**: When conflicting flags provided, document which takes precedence (e.g., --quiet overrides --verbose)

#### 8. Command Testing Strategies
- **Decision Point**: How to test command execution? How to capture output? What to test?
- **Research**: Cobra testing patterns, Go testing best practices
- **Output**: Testing patterns:
  ```go
  func TestRootCommand(t *testing.T) {
      // Capture output
      buf := new(bytes.Buffer)
      rootCmd.SetOut(buf)
      rootCmd.SetErr(buf)

      // Set args
      rootCmd.SetArgs([]string{"--help"})

      // Execute
      err := rootCmd.Execute()

      // Assert
      assert.NoError(t, err)
      assert.Contains(t, buf.String(), "SourceBox")
  }
  ```
  - **Test categories**:
    - Help output (--help flag, content verification)
    - Version display (--version flag, version string match)
    - Flag parsing (valid and invalid flags)
    - Error handling (missing required flags, invalid values)
    - Command registration (subcommands appear in help)
  - **Table-driven tests**: For testing multiple flag combinations
  - **Coverage target**: >80% for command package (cmd/)

#### 9. Error Handling Standards
- **Decision Point**: How to return errors from commands? How to set exit codes? Error message format?
- **Research**: Cobra error handling, Go error patterns
- **Output**:
  - **Error return**: Return error from Run function, Cobra handles exit code:
    ```go
    Run: func(cmd *cobra.Command, args []string) error {
        if err := validateInput(args); err != nil {
            return fmt.Errorf("validation failed: %w", err)
        }
        return nil
    }
    ```
  - **Exit codes**:
    - 0: Success
    - 1: General error (Cobra default for returned errors)
    - Custom: Call os.Exit(code) directly if needed
  - **Error format**: Use fmt.Errorf with %w for error wrapping
  - **User-facing errors**: Clear, actionable messages (avoid technical jargon)
  - **Silencing usage**: `SilenceUsage: true` to not print usage on errors

#### 10. Shell Completion Integration
- **Decision Point**: How does Cobra handle completions? Configuration needed? Which shells?
- **Research**: Cobra completion documentation
- **Output**:
  - **Built-in support**: Cobra generates completions automatically
  - **Shells supported**: bash, zsh, fish, PowerShell
  - **Generation commands** (add later if needed):
    - `sourcebox completion bash` → bash completion script
    - `sourcebox completion zsh` → zsh completion script
  - **Installation**: User runs completion command, pipes to shell config
  - **Current scope**: Not required for F006 (foundation only)
  - **Future**: Add completion subcommand in F007 or later if user demand exists

**Deliverable**: `specs/004-f006-cobra-cli/research.md`

### Phase 1: Design & Contracts

#### 1. Data Model (SKIP for this feature)
**Rationale**: F006 is CLI framework integration. No data entities, no models, no database schema. Skip data-model.md generation.

#### 2. API Contracts (CONDITIONAL for this feature)
**Rationale**: F006 has no REST/GraphQL APIs. However, CLI commands are a form of "contract" for user interaction.

**Optional Contract Document**: `contracts/cli-command-structure.md`
```markdown
# CLI Command Structure Contract

## Root Command
**Command**: `sourcebox`
**Description**: Generate realistic, verticalized demo data instantly
**Global Flags**:
  - `--config <file>` (string): Custom config file path
  - `--verbose, -v` (bool): Enable verbose output
  - `--quiet, -q` (bool): Suppress non-error output
**Version Flag**: `--version`
**Help Flag**: `--help`

## Subcommands

### seed
**Command**: `sourcebox seed [database]`
**Description**: Seed a database with realistic data
**Arguments**:
  - `database` (required): Database type (mysql, postgres)
**Flags**:
  - `--schema <name>` (string, required): Schema name (e.g., fintech-loans)
  - `--records <num>` (int, default: 1000): Number of records
  - `--host <host>` (string, default: localhost): Database host
  - `--port <port>` (int): Database port (auto-detect by database type)
  - `--user <user>` (string, default: root): Database user
  - `--password <pass>` (string): Database password
  - `--db-name <name>` (string, default: demo): Database name
  - `--output <file>` (string): Export to SQL file instead of inserting
  - `--dry-run` (bool): Show what would be done without executing
**Example**: `sourcebox seed mysql --schema=fintech-loans --records=1000`
**Status**: Scaffolded in F006, implemented in F021

### list-schemas
**Command**: `sourcebox list-schemas`
**Aliases**: `ls`
**Description**: List all available data schemas
**Example**: `sourcebox list-schemas` or `sourcebox ls`
**Status**: Scaffolded in F006, implemented in F022

## Exit Codes
- `0`: Success
- `1`: General error (validation failure, execution error)

## Output Format
- **Default**: Human-readable text
- **Verbose mode**: Detailed progress information
- **Quiet mode**: Only errors displayed
```

**Decision**: Optional - include only if helpful for documentation. Primary deliverable is quickstart.md.

#### 3. Quickstart Guide (REQUIRED)
Generate `quickstart.md` with:

```markdown
# F006 Quickstart: Cobra CLI Framework Verification

## Prerequisites
- F004 completed (Project directory structure and Makefile functional)
- Go 1.21+ installed
- cobra-cli tool installed: `go install github.com/spf13/cobra-cli@latest`

## Workflow Overview

The Cobra CLI framework provides:

```
cmd/sourcebox/
├── main.go (Entry point, version injection)
└── cmd/
    ├── root.go (Root command, global flags)
    ├── seed.go (Seed command scaffold)
    └── list_schemas.go (List-schemas command scaffold)
```

## Manual Verification Steps

### 1. Install Cobra CLI Tool
```bash
go install github.com/spf13/cobra-cli@latest

# Verify installation
cobra-cli --version
```

Expected: cobra-cli version information displayed

### 2. Initialize Cobra Structure
```bash
# From repo root
cobra-cli init

# Verify structure
ls -R cmd/sourcebox/
```

Expected output:
```
cmd/sourcebox/:
main.go

cmd/sourcebox/cmd:
root.go
```

### 3. Build and Test Root Command
```bash
# Build binary
make build

# Test help
./dist/sourcebox --help

# Test version
./dist/sourcebox --version
```

Expected help output:
- Tool description (SourceBox generates production-like demo data...)
- Available commands list
- Global flags (--config, --verbose, --quiet)
- Usage examples

Expected version output:
- Version string (e.g., "dev" for local builds, "v1.0.0" for tagged releases)

### 4. Verify Global Flags
```bash
# Test verbose flag
./dist/sourcebox --verbose --help

# Test quiet flag
./dist/sourcebox --quiet --help

# Test config flag parsing
./dist/sourcebox --config=test.yaml --help
```

Expected: No errors, help displays correctly with flags parsed

### 5. Scaffold Subcommands
```bash
# Add seed command
cobra-cli add seed

# Add list-schemas command
cobra-cli add list-schemas

# Verify files created
ls cmd/sourcebox/cmd/
```

Expected files:
- root.go
- seed.go
- list_schemas.go

### 6. Test Subcommand Registration
```bash
# Rebuild after adding commands
make build

# Test seed command help
./dist/sourcebox seed --help

# Test list-schemas command help
./dist/sourcebox list-schemas --help

# Test short alias
./dist/sourcebox ls --help
```

Expected: Each command shows its specific help text with flags and examples

### 7. Test Placeholder Execution
```bash
# Execute seed command (placeholder)
./dist/sourcebox seed mysql --schema=fintech-loans

# Execute list-schemas command (placeholder)
./dist/sourcebox list-schemas
```

Expected: Placeholder message (e.g., "Seed command - implementation coming in F021")

### 8. Run Unit Tests
```bash
# Run command tests
go test ./cmd/sourcebox/cmd/... -v

# Check coverage
go test ./cmd/sourcebox/cmd/... -coverprofile=coverage.txt
go tool cover -func=coverage.txt
```

Expected: All tests pass, coverage >80% for cmd/ package

## Verification Checklist
- [ ] cobra-cli tool installed and functional
- [ ] Cobra structure initialized (main.go, cmd/root.go)
- [ ] Root command displays help correctly
- [ ] Version flag shows version information
- [ ] Global flags parse without errors (--verbose, --quiet, --config)
- [ ] Seed command scaffolded with help text
- [ ] List-schemas command scaffolded with alias (ls)
- [ ] Subcommands appear in root help output
- [ ] All commands have clear, useful help text
- [ ] Unit tests pass with >80% coverage
- [ ] Build system updated (version injection working)

## Debugging Common Issues

### Issue: cobra-cli command not found
- **Check**: Is $GOPATH/bin in your PATH?
- **Fix**: Add `export PATH=$PATH:$(go env GOPATH)/bin` to shell profile

### Issue: Commands not appearing in help
- **Check**: Is init() function calling rootCmd.AddCommand()?
- **Fix**: Verify cmd/seed.go has `init() { rootCmd.AddCommand(seedCmd) }`

### Issue: Version shows "dev" instead of actual version
- **Check**: Is version being injected at build time?
- **Fix**: Update Makefile to include `-ldflags="-X main.version=..."`

### Issue: Flags not parsing correctly
- **Check**: PersistentFlags vs Flags - using correct method?
- **Fix**: Use rootCmd.PersistentFlags() for global flags, cmd.Flags() for local

### Issue: Tests failing to capture output
- **Check**: Using cmd.SetOut() and cmd.SetErr()?
- **Fix**: Set output buffers before calling Execute()

## Performance Verification

### Response Time Checks
```bash
# Test help response time
time ./dist/sourcebox --help

# Test version response time
time ./dist/sourcebox --version
```

Expected:
- Help display: <1 second
- Version display: <100ms
- No noticeable startup delay

## Next Steps
- F021: Implement seed command (uses this CLI structure)
- F022: Implement list-schemas command (uses this CLI structure)
- F027: Add CLI output formatting (progress bars, tables)

## Troubleshooting Reference

### Cobra Initialization Issues
- Ensure you're in the correct directory (repo root)
- Verify go.mod exists (from F003)
- Check Go version: `go version` (must be 1.21+)

### Flag Conflicts
- Global flags (PersistentFlags) cascade to subcommands
- Local flags override global flags with same name
- Use clear naming to avoid conflicts

### Help Text Not Showing
- Verify Use, Short, Long fields are set
- Check for syntax errors in command definition
- Rebuild after changes: `make build`

### Version Injection Not Working
- Check Makefile has correct ldflags pattern
- Verify variable name matches: `main.version`
- Test with: `go build -ldflags="-X main.version=test" ./cmd/sourcebox`
```

**Deliverable**: `specs/004-f006-cobra-cli/quickstart.md`

#### 4. Update Agent Context
Run: `.specify/scripts/bash/update-agent-context.sh claude`

This updates the Claude-specific context file (CLAUDE.md) with:
- Cobra CLI framework overview
- Command structure patterns (root command, subcommands)
- Global flag conventions (--verbose, --quiet, --config)
- Help text quality standards
- TDD requirements for CLI commands
- Version injection mechanism

**Deliverable**: Updated `CLAUDE.md` with CLI framework guidelines

## Constitution Re-verification

After Phase 1 design, verify:
- [ ] Cobra framework is boring, proven tech (Boring Tech principle)
- [ ] Command parsing is fast (<100ms) (Speed principle)
- [ ] CLI works entirely offline (Local-First principle)
- [ ] Help system enables self-service learning (Developer-First principle)
- [ ] Zero configuration required out of box (Zero Config UX principle)
- [ ] TDD coverage >80% for command package (Code Quality constraint)
- [ ] Cross-platform support verified (Platform Support constraint)
- [ ] No unnecessary complexity introduced (Ship Fast principle)
- [ ] Help text includes practical examples (Developer-First principle)
- [ ] Version injection follows build-time pattern (no hardcoding)

## Deliverables Summary

**Generated by /speckit.plan**:
1. ✅ `specs/004-f006-cobra-cli/plan.md` - This file
2. ✅ `specs/004-f006-cobra-cli/research.md` - Phase 0 output (10 research decisions)
3. ✅ `specs/004-f006-cobra-cli/quickstart.md` - Phase 1 output (Cobra verification guide)
4. ⏭️ `specs/004-f006-cobra-cli/data-model.md` - SKIP (N/A for CLI framework)
5. ⏭️ `specs/004-f006-cobra-cli/contracts/cli-command-structure.md` - OPTIONAL (CLI specification)
6. ✅ Updated CLAUDE.md with CLI framework context

**NOT generated by /speckit.plan** (created later by /speckit.tasks):
- `specs/004-f006-cobra-cli/tasks.md` - Phase 2, separate command

## Success Criteria for Planning Phase

- ✅ All 10 research decisions documented with clear rationale
- ✅ Cobra initialization pattern identified and explained
- ✅ Root command architecture follows best practices
- ✅ Flag hierarchy (global vs local) clearly defined
- ✅ Subcommand scaffolding workflow documented
- ✅ Version injection mechanism specified
- ✅ Help text quality standards established
- ✅ Command testing patterns identified
- ✅ Error handling conventions defined
- ✅ Quickstart provides clear verification steps
- ✅ Constitution compliance verified (no violations)
- ✅ Agent context updated with CLI development guidelines
- ✅ Planning artifacts reference constitution and spec correctly

## Anti-Patterns to Avoid

- ❌ Don't generate data-model.md (no data entities in this feature)
- ❌ Don't reinvent CLI framework (use Cobra, it's battle-tested)
- ❌ Don't hardcode version strings (must be build-time injection)
- ❌ Don't skip help text examples (users need practical guidance)
- ❌ Don't skip command tests (makes refactoring risky)
- ❌ Don't over-engineer command structure (keep it simple)
- ❌ Don't use inconsistent flag naming (follow Cobra conventions)
- ❌ Don't create commands without clear help text
- ❌ Don't ignore error handling (every error needs clear message)
- ❌ Don't bypass TDD for command logic (core infrastructure needs tests)
- ❌ Don't add unnecessary global flags (only --verbose, --quiet, --config for now)
- ❌ Don't implement command logic in F006 (scaffold only, implement in F021/F022)

## Implementation Notes

### For the AI Agent
When executing `/speckit.plan` with this prompt:

1. **Start with comprehensive research.md**: Document all 10 research decisions with:
   - Decision: What was chosen (Cobra pattern, flag type, etc.)
   - Rationale: Why chosen (reference constitution, best practices)
   - Alternatives considered: What else was evaluated
   - Source: Where information came from (Cobra docs, Go docs, example projects)

2. **Be explicit about skips**: Clearly state why data-model.md is not needed for CLI framework. Contracts are optional (CLI command specification documentation).

3. **Focus on quickstart.md**: This is the primary deliverable beyond research. Include:
   - Cobra installation and initialization steps
   - Root command verification
   - Subcommand scaffolding workflow
   - Help system testing
   - Version display verification
   - Global flag testing
   - Unit test execution
   - Troubleshooting common issues
   - Performance verification
   - Next steps (F021, F022, F027)

4. **Verify constitutional compliance**:
   - Boring Tech: Cobra is proven, standard Go CLI framework
   - Speed: <100ms response time for help/version
   - Developer-First: Comprehensive help, clear errors, zero config
   - Code Quality: TDD required, >80% coverage
   - Platform Support: Cross-platform compatible
   - No complexity violations (keep command structure simple)

5. **Keep it standard**: This is intentionally boring CLI infrastructure:
   - Use official Cobra patterns (cobra-cli init, cobra-cli add)
   - Follow Cobra naming conventions (Use/Short/Long/Example)
   - Standard flag patterns (PersistentFlags for global, Flags for local)
   - Predictable and maintainable

### Cobra CLI Best Practices

**Project Organization**:
- main.go: Minimal entry point, version variable, calls Execute()
- cmd/root.go: Root command definition, global flags, Execute() function
- cmd/*.go: One file per subcommand, generated with cobra-cli add
- Clear separation of concerns (main for bootstrap, cmd for logic)

**Command Definition**:
- Use: Command signature with argument placeholders
- Short: One-line description (<60 chars)
- Long: Multi-paragraph explanation with use cases
- Example: 2-3 practical examples with comments
- Args: Validation function (ExactArgs, MinimumNArgs, NoArgs)
- Run/RunE: Command implementation (RunE for error return)

**Flag Patterns**:
- PersistentFlags: Available to command and all children
- Flags: Only for specific command
- StringVar/BoolVar/IntVar: Direct variable binding
- MarkFlagRequired: Validation for required flags
- Flag names: kebab-case (--my-flag), avoid camelCase

**Help Text Quality**:
- Clear, concise, developer-friendly language
- Practical examples that users can copy-paste
- Explain what, why, and when to use
- No marketing fluff, no unnecessary verbosity
- Include common pitfalls or gotchas

**Testing Strategies**:
- Test help output (content and formatting)
- Test flag parsing (valid and invalid values)
- Test command registration (subcommands appear)
- Test error handling (missing flags, invalid args)
- Use bytes.Buffer to capture output
- Table-driven tests for multiple scenarios

### CLI vs Build System (F004) Relationship

**F004 (Build System)**:
- Focus: Compiling binaries with version injection
- Tool: Makefile
- Execution: Developer machine (local builds)
- Output: Binaries in /dist
- Performance: <30s build time

**F006 (CLI Framework)**:
- Focus: Command structure and user interaction
- Tool: Cobra framework
- Execution: Runtime (user runs commands)
- Output: Help text, command execution
- Performance: <100ms response time

**Integration**:
- F004 Makefile uses ldflags to inject version into F006 CLI
- F006 CLI commands will invoke F004 build logic (future features)
- Both enforce cross-platform compatibility
- Both follow developer-first design principles

### Version Injection Pattern

**Build-Time Injection**:
```bash
go build -ldflags="-X main.version=$(git describe --tags --always --dirty)" ./cmd/sourcebox
```

**Runtime Display**:
```go
// main.go
var version = "dev"  // Default for local builds

func main() {
    cmd.SetVersion(version)  // Pass to Cobra
    cmd.Execute()
}

// cmd/root.go
func SetVersion(v string) {
    version = v
    rootCmd.Version = v  // Cobra displays this
}
```

**Makefile Integration** (from F004):
```makefile
VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"

build:
	go build $(LDFLAGS) -o dist/sourcebox ./cmd/sourcebox
```

### Related Constitution Sections
- **Core Principle IV**: Boring Tech Wins (Cobra is proven, standard)
- **Core Principle VI**: Developer-First Design (CLI-first, comprehensive help)
- **Core Principle VII**: Ship Fast, Validate Early (TDD required for CLI)
- **Technical Constraint 5**: Code Quality Standards (unit tests, >80% coverage)
- **Technical Constraint 7**: Platform Support (cross-platform CLI)
- **UX Principle 2**: Zero Config (works out of box, sane defaults)
- **UX Principle 5**: Boring CLI (functional > flashy, clear > clever)
- **Development Practice 1**: TDD Required (command logic must be tested)
- **Development Practice 6**: Spec-Kit Driven (this planning workflow)
- **Anti-Pattern 1**: Feature Bloat (only essential flags/commands)
- **Anti-Pattern 5**: Over-Engineering (keep command structure simple)

## Drag-and-Drop Usage

**To use this prompt**:
1. Drag this file into Claude Code
2. Claude will execute the `/speckit.plan` workflow for F006
3. Expected outputs:
   - research.md with 10 documented research decisions
   - quickstart.md with Cobra CLI verification guide
   - Optional: contracts/cli-command-structure.md (CLI specification)
   - Updated CLAUDE.md with CLI framework context
   - Constitution compliance verified

**Estimated time**: 20-30 minutes for complete planning phase

**Next command**: `/speckit.tasks` to generate tasks.md from this plan

**Success indicators**:
- Research decisions are clear and actionable
- Quickstart provides step-by-step verification
- CLI command structure is well-documented
- Testing patterns are established
- No constitutional violations identified
- Ready to proceed to task generation (/speckit.tasks)
