# Implementation Planning Prompt: F005 - GitHub Actions CI/CD Pipeline

## Feature Metadata
- **Feature ID**: F005
- **Name**: GitHub Actions CI/CD Pipeline
- **Feature Branch**: `003-f005-github-actions`
- **Category**: Foundation
- **Phase**: Week 4
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (1 day)
- **Dependencies**: F003 (Git repository), F004 (Build system)
- **Spec Location**: `.specify/prompts/specify/mvp/f005-github-actions-ci-cd-pipeline.md`

## Constitutional Alignment

### Core Principles Verification
- ✅ **Verticalized > Generic**: N/A (infrastructure setup)
- ✅ **Speed > Features**: Fast feedback (<5 min), caching enabled, parallel execution
- ✅ **Local-First, Cloud Optional**: CI complements local testing, doesn't replace it
- ✅ **Boring Tech Wins**: Standard GitHub Actions, golangci-lint, Codecov (proven, boring, standard)
- ✅ **Open Source Forever**: Free GitHub Actions tier, free Codecov for open source
- ✅ **Developer-First Design**: Fast feedback, clear status badges, actionable errors
- ✅ **Ship Fast, Validate Early**: Automated testing prevents regressions, enables rapid iteration

### Technical Constraints Verification
- ✅ **Performance**: Workflow completes in <5 minutes, caching reduces execution time
- ✅ **Distribution**: N/A (CI/CD infrastructure)
- ✅ **Database Support**: N/A (testing infrastructure)
- ✅ **Cost**: $0 (free GitHub Actions tier: 2,000 minutes/month, free Codecov)
- ✅ **Code Quality**: Enforces >80% coverage, runs golangci-lint, race detector enabled
- ✅ **License**: N/A (no changes)
- ✅ **Platform Support**: Tests on macOS, Linux, Windows; Go 1.21 and 1.22

### Legal Constraints Verification (CRITICAL)
- ✅ **Independent Development**: Uses standard public GitHub Actions only
- ✅ **No Employer References**: N/A (infrastructure)
- ✅ **Public Information Only**: All patterns from public GitHub Actions documentation
- ✅ **Open Source Protection**: N/A (no licensing changes)
- ✅ **Illustrative Examples Only**: N/A (no company references)

## Planning Context

### Feature Summary
Establish automated continuous integration using GitHub Actions to run tests, verify builds, check code quality, and track coverage across multiple platforms (macOS, Linux, Windows) and Go versions (1.21, 1.22). Workflow executes on every push and pull request, provides fast feedback (<5 minutes), and displays status badges in README.md. Includes race detection, linting with golangci-lint, and coverage tracking with Codecov.

### Key Technical Decisions Required

**Phase 0 Research Topics**:
1. **GitHub Actions Workflow Syntax**: What is the standard YAML structure? What triggers (push, pull_request) are needed? Best practices for workflow organization?
2. **Matrix Strategy Configuration**: How to configure matrix for 3 OS × 2 Go versions? Best practices for parallel execution? How to handle matrix failures?
3. **Go Testing Best Practices in CI**: What test flags are essential (-v, -race, -coverprofile)? How to generate atomic coverage reports? What's the standard coverage format?
4. **Caching Strategies**: How to cache Go modules in GitHub Actions? What's the optimal cache key? How much time does caching save?
5. **golangci-lint Integration**: How to integrate golangci-lint action? What's the default configuration? Custom .golangci.yml needed? What timeout is reasonable?
6. **Codecov Integration**: How to upload coverage reports? What's the action syntax? How to configure per-matrix reporting? Is Codecov free for open source?
7. **Status Badge Generation**: How to generate GitHub Actions badge? Codecov badge? Go version badge? License badge? What's the URL format?
8. **Workflow Performance Optimization**: How to achieve <5 min execution? What steps are parallelizable? What can be cached? What's the baseline without optimization?
9. **PR Quality Gates**: How to configure branch protection rules? Should PRs be blocked on failing tests? How to require status checks?
10. **Flaky Test Detection**: How to identify flaky tests in CI? Should tests be retried? What's the best practice for handling intermittent failures?

### Technical Context (Pre-filled)

**Language/Version**: Go 1.21 and 1.22 (matrix testing across both versions)
**Primary Dependencies**:
  - GitHub Actions (actions/checkout@v4, actions/setup-go@v5, actions/cache@v3)
  - golangci-lint (golangci/golangci-lint-action@v3)
  - Codecov (codecov/codecov-action@v3)
**Storage**: GitHub Actions runners (standard ubuntu-latest, macos-latest, windows-latest)
**Testing**: `go test` with race detector (-race), coverage reporting (-coverprofile), atomic mode
**Target Platform**: Multi-platform CI (macOS Intel + ARM runners, Linux x86 runners, Windows x86 runners)
**Project Type**: Single project (CLI tool)
**Performance Goals**:
  - Workflow execution time <5 minutes (typical changes)
  - Cache hit reduces time from ~5 min to ~2 min
  - Matrix jobs run in parallel (faster feedback)
**Constraints**:
  - Must use free GitHub Actions tier (2,000 minutes/month)
  - Must test on all 3 platforms (macOS, Linux, Windows)
  - Must test on Go 1.21 and 1.22
  - Must enable race detector (catches concurrency bugs)
  - Must enforce >80% coverage for core packages
  - Must run golangci-lint with zero warnings
  - Must complete in <5 minutes (fast feedback)
**Scale/Scope**: Foundation for all subsequent development (weeks 4-12), critical quality gate

## Planning Workflow

### Phase 0: Research & Technical Decisions

Generate `research.md` with documented decisions for:

#### 1. GitHub Actions Workflow Structure
- **Decision Point**: What's the standard workflow file structure? What sections (name, on, jobs, steps)?
- **Research**: Study GitHub Actions documentation and popular Go project workflows
- **Output**: Workflow structure with:
  - Trigger configuration (push, pull_request)
  - Job definitions (test, lint)
  - Matrix strategy (os, go-version)
  - Step sequence (checkout, setup, cache, test, build)

#### 2. Matrix Strategy Design
- **Decision Point**: How to configure 3 OS × 2 Go versions = 6 test runs? Fail-fast or continue-on-error?
- **Research**: GitHub Actions matrix documentation, best practices for fail-fast
- **Output**: Matrix configuration:
  ```yaml
  strategy:
    matrix:
      os: [ubuntu-latest, macos-latest, windows-latest]
      go-version: ['1.21', '1.22']
  ```
  - Fail-fast: false (run all combinations even if one fails)
  - Run-on: ${{ matrix.os }}

#### 3. Go Testing Configuration
- **Decision Point**: What flags are essential? How to generate coverage in atomic mode?
- **Research**: `go test` documentation, coverage best practices
- **Output**:
  - Test command: `go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...`
  - Flags explained:
    - `-v`: Verbose output (see which tests run)
    - `-race`: Enable race detector (catches concurrency bugs)
    - `-coverprofile=coverage.txt`: Generate coverage report
    - `-covermode=atomic`: Atomic coverage mode (thread-safe, works with -race)
    - `./...`: Test all packages

#### 4. Caching Strategy
- **Decision Point**: What to cache? What cache key? How much time savings?
- **Research**: GitHub Actions cache documentation, Go module caching patterns
- **Output**:
  - Cache path: `~/go/pkg/mod` (Go modules directory)
  - Cache key: `${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}`
  - Restore keys: `${{ runner.os }}-go-`
  - Expected savings: ~5 min → ~2 min on cache hit

#### 5. golangci-lint Integration
- **Decision Point**: Use golangci-lint action or run manually? What configuration? What timeout?
- **Research**: golangci-lint-action documentation, default linters
- **Output**:
  - Use official action: `golangci/golangci-lint-action@v3`
  - Version: `latest` (auto-updates to latest stable)
  - Timeout: 5 minutes (prevent hanging)
  - Configuration: Default linters initially (can customize later with .golangci.yml)
  - Run on: ubuntu-latest only (one platform sufficient for linting)

#### 6. Codecov Integration
- **Decision Point**: How to upload coverage? Per-matrix or aggregated? Free for open source?
- **Research**: Codecov documentation, codecov-action usage
- **Output**:
  - Use official action: `codecov/codecov-action@v3`
  - Upload from all matrix jobs (separate coverage reports)
  - Naming: `codecov-${{ matrix.os }}-go${{ matrix.go-version }}`
  - Flags: `unittests`
  - Free tier: Yes, free for open source projects
  - Coverage tracking: Trends over time, PR comments with coverage delta

#### 7. Status Badge URLs
- **Decision Point**: What badges to display? What's the URL format?
- **Research**: GitHub Actions badge documentation, Shields.io badges
- **Output**: Badge URLs (replace `yourusername/sourcebox` with actual repo):
  ```markdown
  [![CI](https://github.com/yourusername/sourcebox/actions/workflows/ci.yml/badge.svg)](https://github.com/yourusername/sourcebox/actions/workflows/ci.yml)
  [![codecov](https://codecov.io/gh/yourusername/sourcebox/branch/main/graph/badge.svg)](https://codecov.io/gh/yourusername/sourcebox)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
  [![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/sourcebox)](go.mod)
  ```

#### 8. Workflow Performance Optimization
- **Decision Point**: How to achieve <5 min target? What's parallelizable?
- **Research**: GitHub Actions performance best practices
- **Output**:
  - Parallel execution: Matrix jobs run concurrently (default behavior)
  - Caching: Go modules cached (saves ~3 min)
  - Incremental testing: Go's built-in test cache (automatic)
  - Target baseline: 3-5 min with caching, 5-8 min without
  - Future optimization: Consider limiting matrix to recent Go version only

#### 9. PR Quality Gates
- **Decision Point**: Should failing tests block PR merges? What checks are required?
- **Research**: GitHub branch protection documentation
- **Output**:
  - Branch protection rule: Require status checks for main branch
  - Required checks:
    - Test job (all 6 matrix combinations)
    - Lint job
  - Merge restriction: Require checks to pass before merging
  - Configuration: Manual setup in GitHub repo settings (document in quickstart)

#### 10. Flaky Test Handling
- **Decision Point**: How to detect and handle flaky tests?
- **Research**: Go testing best practices, CI/CD flaky test patterns
- **Output**:
  - Detection: Monitor workflow history for intermittent failures
  - Prevention: Ensure tests are deterministic, avoid time/network dependencies
  - Handling: Don't use test retries initially (masks problems)
  - Process: Track flaky tests as bugs, fix root cause
  - Target: <5% flaky test rate (monitored via workflow history)

**Deliverable**: `specs/003-f005-github-actions/research.md`

### Phase 1: Design & Contracts

#### 1. Data Model (SKIP for this feature)
**Rationale**: F005 is CI/CD pipeline setup. No data entities, no models, no database schema. Skip data-model.md generation.

#### 2. API Contracts (CONDITIONAL for this feature)
**Rationale**: F005 has no REST/GraphQL APIs. However, the workflow YAML itself is a "contract" for CI behavior.

**Optional Contract Document**: `contracts/ci-workflow-schema.yaml`
```yaml
# GitHub Actions Workflow Contract
# This documents the expected structure and behavior of .github/workflows/ci.yml

workflow:
  name: CI
  triggers:
    - push (branches: main, develop)
    - pull_request (branches: main, develop)

  jobs:
    test:
      runs_on: matrix[ubuntu-latest, macos-latest, windows-latest]
      go_versions: matrix[1.21, 1.22]
      steps:
        - checkout_code
        - setup_go
        - cache_modules
        - download_dependencies
        - run_tests (with race detector, coverage)
        - upload_coverage
        - verify_build

    lint:
      runs_on: ubuntu-latest
      go_version: 1.22
      steps:
        - checkout_code
        - setup_go
        - run_golangci_lint

  expected_duration: <5 minutes (cached), <8 minutes (cold)
  required_status_checks:
    - test (all 6 matrix combinations)
    - lint
```

**Decision**: Optional - include only if helpful for documentation. Primary deliverable is quickstart.md.

#### 3. Quickstart Guide (REQUIRED)
Generate `quickstart.md` with:

```markdown
# F005 Quickstart: CI/CD Pipeline Verification

## Prerequisites
- F003 completed (Git repository with remote configured)
- F004 completed (Build system and Makefile functional)
- GitHub repository created and pushed
- GitHub account with Actions enabled (free tier)

## Workflow Overview

The CI/CD pipeline runs automatically on every push and pull request:

```
.github/workflows/ci.yml
├── Test Job (matrix: 3 OS × 2 Go versions = 6 runs)
│   ├── Checkout code
│   ├── Setup Go 1.21 or 1.22
│   ├── Cache Go modules
│   ├── Download dependencies
│   ├── Run tests (-race -cover)
│   ├── Upload coverage to Codecov
│   └── Build binary
└── Lint Job (ubuntu-latest, Go 1.22)
    ├── Checkout code
    ├── Setup Go 1.22
    └── Run golangci-lint
```

## Manual Workflow Trigger

### 1. Initial Workflow Creation
```bash
# Workflow file created in .github/workflows/ci.yml
git add .github/workflows/ci.yml
git commit -m "Add GitHub Actions CI/CD pipeline"
git push origin 003-f005-github-actions
```

### 2. Verify Workflow Triggers
```bash
# Navigate to GitHub repository
# Go to Actions tab
# Verify workflow run appears for your commit
```

Expected: Workflow run shows "CI" with status "In progress" or "Completed"

### 3. Monitor Workflow Execution
```bash
# In GitHub Actions tab:
# Click on workflow run
# View job logs (Test, Lint)
# Verify all 6 test matrix jobs complete successfully
```

Expected output:
- ✅ Test (ubuntu-latest, go 1.21)
- ✅ Test (ubuntu-latest, go 1.22)
- ✅ Test (macos-latest, go 1.21)
- ✅ Test (macos-latest, go 1.22)
- ✅ Test (windows-latest, go 1.21)
- ✅ Test (windows-latest, go 1.22)
- ✅ Lint

### 4. Verify Coverage Upload
```bash
# In workflow logs, find "Upload coverage to Codecov" step
# Verify upload succeeds (no errors)
# Navigate to codecov.io/gh/yourusername/sourcebox
# Verify coverage reports appear
```

Expected: Coverage report uploaded for each matrix job, overall coverage visible on Codecov

### 5. Check Status Badges
```bash
# View README.md on GitHub
# Verify CI badge shows "passing" (green)
# Verify Codecov badge shows coverage percentage
```

Expected badges:
- [![CI](https://github.com/yourusername/sourcebox/actions/workflows/ci.yml/badge.svg)](...)
- [![codecov](https://codecov.io/gh/yourusername/sourcebox/branch/main/graph/badge.svg)](...)

## PR Workflow Testing

### 1. Create Test PR
```bash
# Make a trivial change
echo "# Test" >> test-file.md
git add test-file.md
git commit -m "Test CI on PR"
git push origin 003-f005-github-actions

# Create PR via GitHub UI
# Target branch: main
```

### 2. Verify PR Checks
```bash
# In PR view on GitHub:
# Scroll to "Checks" section
# Verify "CI / Test (matrix)" and "CI / Lint" appear
# Verify status updates as jobs run
```

Expected: PR shows "All checks have passed" or "Some checks are pending"

### 3. Test Failing Workflow
```bash
# Intentionally break a test
# Push to PR branch
# Verify workflow fails
# Verify PR is blocked from merging
```

Expected: PR shows "Some checks were not successful" and merge button is disabled

## Debugging Failed Workflows

### Common Issues

**Issue**: Workflow doesn't trigger
- **Check**: Repository settings → Actions → Workflows enabled?
- **Fix**: Enable Actions in repository settings

**Issue**: Tests fail in CI but pass locally
- **Check**: Platform-specific code? Race conditions? Environment differences?
- **Fix**: Run locally with `go test -race ./...` to reproduce

**Issue**: Coverage upload fails
- **Check**: Codecov token configured? (not needed for public repos)
- **Fix**: For private repos, add CODECOV_TOKEN secret

**Issue**: Workflow timeout (>5 minutes)
- **Check**: Cache hit? Large dependency download?
- **Fix**: Verify cache configuration, check go.sum hasn't changed

**Issue**: golangci-lint failures
- **Check**: New linting rules? Code style issues?
- **Fix**: Run `golangci-lint run` locally, fix issues

## Verification Checklist
- [ ] Workflow file exists at .github/workflows/ci.yml
- [ ] Workflow triggers on push to main/develop branches
- [ ] Workflow triggers on pull requests
- [ ] All 6 test matrix jobs pass
- [ ] Lint job passes
- [ ] Coverage reports upload to Codecov
- [ ] Status badges show "passing" in README
- [ ] Workflow completes in <5 minutes (with caching)
- [ ] PR checks appear and update correctly
- [ ] Failed tests block PR merging

## Performance Monitoring

### Baseline Metrics
- **First run** (no cache): 5-8 minutes
- **Subsequent runs** (cache hit): 2-4 minutes
- **Per-platform test time**: ~1-2 minutes
- **Lint time**: ~30-60 seconds

### Tracking Workflow Duration
```bash
# In GitHub Actions tab:
# View workflow history
# Monitor duration trends
# Target: Stay under 5 minutes for 90% of runs
```

## Next Steps
- F006: Add Cobra CLI framework (tests will run via this pipeline)
- F009: Add dependencies (workflow will cache them)
- F020: Add data generation tests (workflow will enforce coverage >80%)

## Troubleshooting Reference

### Workflow YAML Syntax Errors
- Validate YAML: https://www.yamllint.com/
- Check indentation (spaces, not tabs)
- Verify action versions are valid

### Cache Misses
- Check if go.sum changed (invalidates cache)
- Verify cache key matches expected pattern
- Clear cache manually if corrupted (repo settings)

### Coverage Gaps
- Run locally: `go test -coverprofile=coverage.txt ./...`
- View coverage: `go tool cover -html=coverage.txt`
- Target: >80% for core packages (pkg/*)

### Platform-Specific Failures
- Windows path issues: Use `filepath.Join()`, not string concatenation
- macOS/Linux differences: Check file permissions, case sensitivity
- ARM runners: May have different performance characteristics
```

**Deliverable**: `specs/003-f005-github-actions/quickstart.md`

#### 4. Update Agent Context
Run: `.specify/scripts/bash/update-agent-context.sh claude`

This updates the Claude-specific context file (CLAUDE.md) with:
- CI/CD pipeline overview
- GitHub Actions workflow structure
- Testing requirements (race detector, coverage >80%)
- Linting requirements (golangci-lint, zero warnings)
- Manual QA supplement (CI doesn't replace manual testing)

**Deliverable**: Updated `CLAUDE.md` with CI/CD testing information

## Constitution Re-verification

After Phase 1 design, verify:
- [ ] Workflow uses standard GitHub Actions (Boring Tech principle)
- [ ] Fast feedback achieved (<5 min target) (Speed > Features principle)
- [ ] Free tier sufficient (Cost Constraints)
- [ ] Tests run on all 3 platforms (Platform Support constraint)
- [ ] Coverage tracking enforces >80% for core packages (Code Quality constraint)
- [ ] Race detector enabled (Code Quality constraint)
- [ ] golangci-lint runs with zero warnings (Code Quality constraint)
- [ ] Manual QA still required before releases (Development Practice 3)
- [ ] TDD workflow supported (tests run automatically) (Development Practice 1)
- [ ] No over-engineering or unnecessary complexity (Ship Fast principle)

## Deliverables Summary

**Generated by /speckit.plan**:
1. ✅ `specs/003-f005-github-actions/plan.md` - This file
2. ✅ `specs/003-f005-github-actions/research.md` - Phase 0 output (10 research decisions)
3. ✅ `specs/003-f005-github-actions/quickstart.md` - Phase 1 output (CI/CD verification guide)
4. ⏭️ `specs/003-f005-github-actions/data-model.md` - SKIP (N/A for CI/CD)
5. ⏭️ `specs/003-f005-github-actions/contracts/ci-workflow-schema.yaml` - OPTIONAL (workflow structure documentation)
6. ✅ Updated CLAUDE.md with CI/CD context

**NOT generated by /speckit.plan** (created later by /speckit.tasks):
- `specs/003-f005-github-actions/tasks.md` - Phase 2, separate command

## Success Criteria for Planning Phase

- ✅ All 10 research decisions documented with clear rationale
- ✅ GitHub Actions workflow structure follows best practices
- ✅ Matrix strategy covers all platforms and Go versions
- ✅ Caching strategy achieves <5 min workflow time
- ✅ golangci-lint integration configured correctly
- ✅ Codecov integration configured for free tier
- ✅ Status badge URLs documented
- ✅ Quickstart provides clear CI/CD verification steps
- ✅ Flaky test handling strategy documented
- ✅ Constitution compliance verified (no violations)
- ✅ Agent context updated with CI/CD testing requirements
- ✅ Planning artifacts reference constitution and spec correctly

## Anti-Patterns to Avoid

- ❌ Don't generate data-model.md (no data entities in this feature)
- ❌ Don't test on only one platform (must test macOS, Linux, Windows)
- ❌ Don't skip race detector (can hide concurrency bugs)
- ❌ Don't skip linting (leads to inconsistent code style)
- ❌ Don't use paid CI services (violates cost constraints)
- ❌ Don't exceed 5-minute workflow time (discourages frequent commits)
- ❌ Don't ignore failed workflows (defeats purpose of CI)
- ❌ Don't skip caching (wastes GitHub Actions minutes)
- ❌ Don't use test retries for flaky tests (masks root cause)
- ❌ Don't configure complex conditional workflows (keep it simple)
- ❌ Don't over-optimize prematurely (baseline first, optimize if needed)
- ❌ Don't forget branch protection rules (manual setup required)

## Implementation Notes

### For the AI Agent
When executing `/speckit.plan` with this prompt:

1. **Start with comprehensive research.md**: Document all 10 research decisions with:
   - Decision: What was chosen
   - Rationale: Why chosen (reference constitution/spec)
   - Alternatives considered: What else was evaluated
   - Source: Where information came from (GitHub Actions docs, Go docs)

2. **Be explicit about skips**: Clearly state why data-model.md is not needed for CI/CD infrastructure. Contracts are optional (workflow schema documentation).

3. **Focus on quickstart.md**: This is the primary deliverable beyond research. Include:
   - Workflow structure overview
   - Step-by-step verification (manual trigger, PR testing, debugging)
   - Verification checklist
   - Troubleshooting reference
   - Performance monitoring guidelines
   - Next steps (F006, F009, F020)

4. **Verify constitutional compliance**:
   - Boring Tech: Standard GitHub Actions, golangci-lint, Codecov
   - Speed: <5 min workflow time, caching enabled
   - Cost: Free tier (GitHub Actions, Codecov)
   - Platform Support: All 3 platforms, 2 Go versions
   - Code Quality: >80% coverage, race detector, zero lint warnings
   - Developer-First: Fast feedback, clear status badges
   - No complexity violations

5. **Keep it standard**: This is intentionally boring CI/CD infrastructure:
   - Use official GitHub Actions (actions/*, golangci/*, codecov/*)
   - Follow Go testing conventions (-race, -cover)
   - Standard matrix strategy (no exotic configurations)
   - Predictable and maintainable

### GitHub Actions Best Practices

**Workflow Organization**:
- One workflow file (ci.yml) for simplicity
- Separate jobs for different concerns (test, lint)
- Use matrix for cross-platform testing
- Cache dependencies to reduce execution time

**Test Job Design**:
- Matrix strategy for OS and Go version combinations
- Run on all combinations (fail-fast: false)
- Enable race detector (-race flag)
- Generate coverage in atomic mode (-covermode=atomic)
- Upload coverage to Codecov
- Verify build succeeds (quick smoke test)

**Lint Job Design**:
- Run on single platform (ubuntu-latest) - linting doesn't need multi-platform
- Use latest Go version (1.22)
- Use official golangci-lint action
- Set reasonable timeout (5 minutes)

**Caching Strategy**:
- Cache Go modules (~/go/pkg/mod)
- Cache key based on go.sum hash
- Restore keys for partial matches
- Expected savings: 50-60% execution time reduction

**Performance Optimization**:
- Parallel matrix execution (default)
- Module caching (explicit configuration)
- Incremental testing (Go built-in)
- Target: <5 min with cache, <8 min cold

### CI/CD vs Build System (F004) Differences

**F004 (Build System)**:
- Focus: Compiling binaries
- Tool: Makefile
- Execution: Local developer machine
- Output: Binaries in /dist
- Performance: <2 min build time

**F005 (CI/CD Pipeline)**:
- Focus: Running tests
- Tool: GitHub Actions
- Execution: GitHub-hosted runners
- Output: Test results, coverage reports
- Performance: <5 min workflow time

**Both enforce**:
- Cross-platform support
- Performance requirements
- Code quality standards
- Developer-first design

### Codecov Configuration

**For Public Repos** (SourceBox):
- No token required (automatic)
- Coverage reports public
- Trends tracked over time
- PR comments with coverage delta

**For Private Repos** (future):
- Add CODECOV_TOKEN secret
- Coverage reports private
- Same features as public

**Coverage Thresholds**:
- Target: >80% for core packages (pkg/*)
- Acceptable: 60-80% for CLI/examples
- Track trends, not absolute numbers

### Branch Protection Configuration

**Manual Setup Required** (after workflow is working):
1. Repository Settings → Branches
2. Add rule for `main` branch
3. Require status checks to pass:
   - Test (ubuntu-latest, go 1.21)
   - Test (ubuntu-latest, go 1.22)
   - Test (macos-latest, go 1.21)
   - Test (macos-latest, go 1.22)
   - Test (windows-latest, go 1.21)
   - Test (windows-latest, go 1.22)
   - Lint
4. Require branches to be up to date

**Rationale**: Prevents merging broken code, enforces quality gates

### Related Constitution Sections
- **Core Principle VII**: Ship Fast, Validate Early (automated testing prevents regressions)
- **Core Principle VI**: Developer-First Design (fast feedback, clear status)
- **Core Principle II**: Speed > Features (<5 min workflow, caching enabled)
- **Technical Constraint 5**: Code Quality Standards (>80% coverage, zero warnings)
- **Technical Constraint 7**: Platform Support (macOS, Linux, Windows testing)
- **Technical Constraint 4**: Cost Constraints (free GitHub Actions tier)
- **Development Practice 1**: TDD Required (CI runs tests automatically)
- **Development Practice 3**: Manual QA Before Release (CI supplements, doesn't replace)
- **Anti-Pattern 5**: Over-Engineering (keep workflow simple, standard)

## Drag-and-Drop Usage

**To use this prompt**:
1. Drag this file into Claude Code
2. Claude will execute the `/speckit.plan` workflow for F005
3. Expected outputs:
   - research.md with 10 documented research decisions
   - quickstart.md with CI/CD verification guide
   - Optional: contracts/ci-workflow-schema.yaml (workflow structure doc)
   - Updated CLAUDE.md with CI/CD testing context
   - Constitution compliance verified

**Estimated time**: 15-20 minutes for complete planning phase

**Next command**: `/speckit.tasks` to generate tasks.md from this plan
