# Contributing to SourceBox

Thank you for your interest in contributing to SourceBox! This document provides guidelines for contributing to the project. All participants must follow the [SourceBox Code of Conduct](CODE_OF_CONDUCT.md); report any concerns via [GitHub Issues](https://github.com/jbeausoleil/sourcebox/issues).

## Table of Contents
- [Bug Reports](#bug-reports)
- [Feature Requests](#feature-requests)
- [Pull Requests](#pull-requests)
- [Code Style](#code-style)
- [Testing Requirements](#testing-requirements)

## Bug Reports

Found a bug? Please help us fix it by providing detailed information.

**Before submitting:** Search existing issues to avoid duplicates.

**Required information:**
- **Description**: Clear summary of the bug
- **Steps to Reproduce**: Numbered list of exact steps
- **Environment**:
  - OS and version (e.g., macOS 14.2, Ubuntu 22.04, Windows 11)
  - Go version (`go version`)
  - SourceBox version (`sourcebox --version` or commit SHA)
- **Expected Behavior**: What should happen
- **Actual Behavior**: What actually happened
- **Logs**: Relevant error messages or stack traces
- **Additional Context**: Screenshots, configuration files, etc. (if applicable)

**Example:**
```
### Description
CLI crashes when processing files with special characters

### Steps to Reproduce
1. Create file named `test@file.txt`
2. Run `sourcebox process test@file.txt`
3. Observe crash

### Environment
- OS: macOS 14.2
- Go: 1.21.5
- SourceBox: commit abc123

### Expected Behavior
File should be processed successfully

### Actual Behavior
Panic: invalid character '@' in filename

### Logs
[paste error output here]
```

## Feature Requests

We welcome ideas for new features! Please focus on describing the problem you're trying to solve rather than prescribing a specific solution.

**Required information:**
- **Problem Statement**: What problem are you trying to solve? What's the use case?
- **Current Workaround**: How do you currently handle this (if at all)?
- **Expected Benefit**: Who benefits and how? (users, developers, both)
- **Additional Context**: Examples, mockups, or references to similar features in other tools

**Example:**
```
### Problem Statement
As a developer working with multiple projects, I need to quickly switch between different source configurations, but currently I have to manually edit config files each time.

### Current Workaround
I maintain separate config files (config-project1.yaml, config-project2.yaml) and manually copy them to config.yaml when switching projects.

### Expected Benefit
Users managing multiple projects would save time and reduce errors from manual config editing.

### Additional Context
Similar to how `git` handles profiles or how `kubectl` manages contexts.
```

## Pull Requests

Ready to contribute code? Follow this process:

### 1. Fork and Branch
- Fork the repository to your GitHub account
- Clone your fork locally
- Create a feature branch: `git checkout -b feature/your-feature-name`
- Use descriptive branch names: `fix/bug-description` or `feature/feature-description`

### 2. Make Changes
- Follow the [Code Style](#code-style) guidelines
- Follow the [Testing Requirements](#testing-requirements)
- Keep commits focused and atomic
- Write clear commit messages following this format:
  ```
  type(scope): brief description

  Detailed explanation of what changed and why.

  Fixes #issue-number (if applicable)
  ```
  - **Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
  - **Scope**: package or component name (optional)
  - **Example**: `feat(config): add profile switching support`

### 3. Before Submitting
- Run `gofmt -w .` to format code
- Run `go vet ./...` to check for issues
- Run `go test ./...` to ensure all tests pass
- Update documentation if needed
- Rebase on latest main: `git pull --rebase upstream main`

### 4. Submit PR
- Push to your fork: `git push origin feature/your-feature-name`
- Open a Pull Request on GitHub
- Provide a clear description of changes
- Link related issues
- Wait for review and address feedback

### 5. Review Process
- Maintainers will review your PR
- Address requested changes by pushing new commits
- Once approved, maintainers will merge your PR

## Code Style

SourceBox follows standard Go conventions and best practices.

### Formatting
- **Required**: Run `gofmt -w .` before committing
- **Required**: Run `go vet ./...` and fix all issues before committing
- Use tabs for indentation (Go standard)
- Keep line length reasonable (aim for <100 characters)

### Naming Conventions
- Use `MixedCaps` or `mixedCaps` (not underscores)
- Exported names start with uppercase
- Unexported names start with lowercase
- Interface names: single-method interfaces end in `-er` (e.g., `Reader`, `Writer`)

### Code Organization
- Keep functions focused and small (aim for <50 lines)
- Write self-documenting code with clear variable names
- Add comments for non-obvious logic or complex algorithms
- Use godoc-style comments for exported functions and types:
  ```go
  // ProcessFile reads and processes the given file.
  // It returns an error if the file cannot be read or is invalid.
  func ProcessFile(path string) error {
      // implementation
  }
  ```

### Error Handling
- Always check errors; don't ignore them
- Wrap errors with context: `fmt.Errorf("failed to process file: %w", err)`
- Return errors rather than panicking (except in truly exceptional cases)

### Package Structure
- One package per directory
- Keep package names short and lowercase
- Avoid `util` or `common` packages; use descriptive names

## Testing Requirements

Quality is important. Follow these testing guidelines:

### Test-Driven Development (TDD)
- **Required for core logic**: Write tests before implementation
- **Process**: Write failing test → Implement feature → Test passes → Refactor
- Example: Business logic, data processing, algorithms

### Coverage Goals
- **Target**: >80% coverage for core packages
- **Minimum**: 60% coverage for all packages
- Run `go test -cover ./...` to check coverage
- Run `go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out` to view detailed coverage

### Test Organization
- Test files: `*_test.go` in the same package
- Table-driven tests for multiple similar cases
- Use `t.Run()` for subtests
- Example:
  ```go
  func TestProcessFile(t *testing.T) {
      tests := []struct {
          name    string
          input   string
          want    string
          wantErr bool
      }{
          {"valid file", "test.txt", "processed", false},
          {"missing file", "missing.txt", "", true},
      }

      for _, tt := range tests {
          t.Run(tt.name, func(t *testing.T) {
              got, err := ProcessFile(tt.input)
              if (err != nil) != tt.wantErr {
                  t.Errorf("ProcessFile() error = %v, wantErr %v", err, tt.wantErr)
                  return
              }
              if got != tt.want {
                  t.Errorf("ProcessFile() = %v, want %v", got, tt.want)
              }
          })
      }
  }
  ```

### Manual QA
- **Required for releases**: Test on macOS, Linux, and Windows
- Test critical user workflows end-to-end
- Verify installation and upgrade paths
- Check documentation matches behavior

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test -run TestProcessFile ./pkg/processor

# Run tests with race detection
go test -race ./...
```

---

## Questions?

If you have questions about contributing, feel free to:
- Open a discussion on GitHub
- Ask in an issue or pull request
- Check existing documentation in the repository

Thank you for contributing to SourceBox!
