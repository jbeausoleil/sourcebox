# Feature Specification Prompt: F005 - GitHub Actions CI/CD Pipeline

## Feature Metadata
- **Feature ID**: F005
- **Name**: GitHub Actions CI/CD Pipeline
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (1 day)
- **Dependencies**: F003 (Git repository), F004 (Build system)

## Constitutional Alignment

### Core Principles
- ✅ **Ship Fast, Validate Early**: Automated testing prevents regressions
- ✅ **Developer-First Design**: Fast feedback on commits and PRs
- ✅ **Quality Standards**: Continuous testing ensures code quality

### Technical Constraints
- ✅ **Platform Support**: Test on macOS, Linux, Windows
- ✅ **Code Quality**: >80% coverage for core logic
- ✅ **Cost Constraints**: Use free GitHub Actions tier

### Development Practices
- ✅ **TDD Required**: CI runs all tests automatically
- ✅ **Manual QA Before Release**: CI is supplement, not replacement

## User Story
**US-MVP-002**: "As a developer, I want automated tests to run on every commit so I can catch bugs early and maintain code quality."

## Problem Statement
SourceBox needs continuous integration to automatically run tests, check code quality, and verify builds across multiple platforms and Go versions. This prevents regressions, maintains quality standards, and provides fast feedback to contributors. Without CI/CD, manual testing is error-prone and doesn't scale as the project grows.

## Solution Overview
Implement a GitHub Actions workflow that runs on every push and pull request. The workflow will execute tests across multiple Go versions (1.21, 1.22) and operating systems (Ubuntu, macOS, Windows), generate coverage reports, and verify builds succeed. A status badge in README.md will show build health at a glance.

## Detailed Requirements

### Acceptance Criteria
1. **GitHub Actions Workflow Created**: `.github/workflows/ci.yml` exists
2. **Tests Run on Every Commit**: Triggered by push to any branch
3. **Tests Run on Pull Requests**: Triggered by PR creation/update
4. **Multiple Go Versions Tested**: Go 1.21 and 1.22
5. **Multiple OS Tested**: ubuntu-latest, macos-latest, windows-latest
6. **Build Verification**: `go build` succeeds for all platforms
7. **Status Badge Added**: README.md shows build status badge
8. **Fast Feedback**: Workflow completes in < 5 minutes for typical changes

### Technical Specifications

#### Workflow File: `.github/workflows/ci.yml`

```yaml
name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  test:
    name: Test
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
        go-version: ['1.21', '1.22']

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: ${{ matrix.go-version }}

    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        restore-keys: |
          ${{ runner.os }}-go-

    - name: Download dependencies
      run: go mod download

    - name: Run tests
      run: go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.txt
        flags: unittests
        name: codecov-${{ matrix.os }}-go${{ matrix.go-version }}

    - name: Build
      run: go build -v ./cmd/sourcebox

  lint:
    name: Lint
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.22'

    - name: Run golangci-lint
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest
        args: --timeout=5m
```

#### Status Badge for README.md

Add to README.md:
```markdown
[![CI](https://github.com/yourusername/sourcebox/actions/workflows/ci.yml/badge.svg)](https://github.com/yourusername/sourcebox/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/yourusername/sourcebox/branch/main/graph/badge.svg)](https://codecov.io/gh/yourusername/sourcebox)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/sourcebox)](go.mod)
```

#### Workflow Features

**Test Job**:
- Matrix strategy: 3 OS × 2 Go versions = 6 test runs
- Race detector enabled (`-race` flag)
- Coverage report generated
- Coverage uploaded to Codecov for tracking trends

**Lint Job**:
- Runs golangci-lint with comprehensive checks
- Catches common Go issues (unused vars, error handling, etc.)
- Enforces code style consistency

**Caching**:
- Go modules cached to speed up workflow
- Cache key based on go.sum hash
- Reduces workflow time from ~5 min to ~2 min on cache hit

### Performance Considerations
- **Parallel Execution**: Matrix jobs run in parallel (faster feedback)
- **Caching**: Go modules cached to reduce download time
- **Incremental Testing**: Only affected tests run where possible
- **Timeout**: Set 5-minute timeout for lint to prevent hanging

### Testing Strategy
**Manual Verification**:
1. Push commit to repository - verify workflow triggers
2. Create PR - verify workflow runs and shows status
3. Check Actions tab - verify all jobs pass
4. Verify badge shows "passing" status
5. Break a test intentionally - verify workflow fails and shows error

**Ongoing Validation**:
- Monitor workflow duration (should stay < 5 minutes)
- Check coverage trends (should stay > 80% for core packages)
- Review failed workflows to identify flaky tests

## Dependencies
- **Upstream**:
  - F003 (Git repository must exist)
  - F004 (Build system and Makefile must exist)
- **Downstream**: All future features benefit from automated testing

## Deliverables
1. GitHub Actions workflow file (`.github/workflows/ci.yml`)
2. Status badges added to README.md
3. Codecov integration configured (optional but recommended)
4. golangci-lint configuration (`.golangci.yml` if custom settings needed)
5. Documentation of CI/CD process in CONTRIBUTING.md

## Success Criteria
- ✅ Workflow runs on every commit and PR
- ✅ Tests pass on all platforms and Go versions
- ✅ Build succeeds for all platforms
- ✅ Status badge shows "passing"
- ✅ Workflow completes in < 5 minutes
- ✅ Coverage reports uploaded successfully

## Anti-Patterns to Avoid
- ❌ Only testing on one platform (must test cross-platform)
- ❌ Skipping race detector (can hide concurrency bugs)
- ❌ No linting (leads to inconsistent code style)
- ❌ Slow workflows (> 10 minutes discourages frequent commits)
- ❌ Ignoring failed workflows (defeats the purpose of CI)
- ❌ Not caching dependencies (wastes time and GitHub Actions minutes)

## Implementation Notes
- GitHub Actions free tier: 2,000 minutes/month (sufficient for this project)
- Matrix strategy provides comprehensive coverage without manual setup
- golangci-lint is the standard Go linter (replaces multiple tools)
- Codecov is free for open source projects
- Badge URLs need to be updated with actual GitHub username

## TDD Requirements
**Not applicable for CI/CD setup** - This feature tests the testing infrastructure itself. However, verify it works by:
1. Adding a failing test temporarily
2. Pushing to GitHub
3. Verifying workflow fails
4. Fixing test
5. Verifying workflow passes

## Related Constitution Sections
- **Code Quality Standards (Technical Constraint 5)**: >80% coverage, strict mode, zero warnings
- **TDD Required for Core Functionality (Development Practice 1)**: CI enforces test execution
- **Manual QA Before Every Release (Development Practice 3)**: CI supplements, doesn't replace
- **Platform Support (Technical Constraint 7)**: Test on macOS, Linux, Windows
- **Cost Constraints (Technical Constraint 4)**: Free GitHub Actions tier
