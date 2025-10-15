# F005 Quickstart: CI/CD Pipeline Verification

**Feature**: GitHub Actions CI/CD Pipeline
**Branch**: `003-f005-github-actions`
**Date**: 2025-10-14

## Prerequisites

Before proceeding with this quickstart, ensure:

- ✅ **F003 completed**: Git repository exists with remote configured (github.com/jbeausoleil/sourcebox)
- ✅ **F004 completed**: Build system and Makefile are functional (`make build`, `make test` work locally)
- ✅ **GitHub repository**: Repository created at github.com/jbeausoleil/sourcebox and pushed
- ✅ **GitHub Actions enabled**: Actions tab visible in repository (enabled by default for public repos)
- ✅ **Local testing passes**: `go test ./...` and `make build` complete successfully locally

---

## Workflow Overview

The CI/CD pipeline runs automatically on every push and pull request:

```
.github/workflows/ci.yml
│
├── Trigger: push to any branch OR pull request to main/develop
│
├── Test Job (matrix: 3 OS × 2 Go versions = 6 parallel runs)
│   ├── Checkout code (actions/checkout@v4)
│   ├── Setup Go 1.21 or 1.22 (actions/setup-go@v5)
│   ├── Cache Go modules (actions/cache@v3)
│   │   └── Cache key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
│   ├── Download dependencies (go mod download)
│   ├── Run tests with race detector and coverage
│   │   └── go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
│   ├── Upload coverage to Codecov (codecov/codecov-action@v3)
│   └── Verify build succeeds (make build)
│
└── Lint Job (ubuntu-latest, Go 1.22)
    ├── Checkout code (actions/checkout@v4)
    ├── Setup Go 1.22 (actions/setup-go@v5)
    └── Run golangci-lint (golangci/golangci-lint-action@v3)
```

**Expected Execution Time**:
- First run (no cache): 5-8 minutes
- Subsequent runs (cache hit): 2-4 minutes
- Lint job: 30-60 seconds

---

## Manual Workflow Trigger

### 1. Initial Workflow Creation

After the `.github/workflows/ci.yml` file is created:

```bash
# Verify workflow file exists
ls -la .github/workflows/ci.yml

# Stage and commit workflow
git add .github/workflows/ci.yml
git commit -m "Add GitHub Actions CI/CD pipeline (F005)

- Test matrix: macOS, Linux, Windows × Go 1.21, 1.22
- Enable race detector and coverage reporting
- Integrate golangci-lint for code quality
- Upload coverage to Codecov
- Target: <5 min workflow execution with caching"

# Push to feature branch
git push origin 003-f005-github-actions
```

**Expected Result**: Workflow triggers automatically on push.

### 2. Verify Workflow Triggers

Navigate to your repository on GitHub:

```
https://github.com/jbeausoleil/sourcebox/actions
```

**What to check**:
1. **Workflow appears**: You should see "CI" workflow run for your commit
2. **Status shows**: "In progress" (yellow circle) or "Completed" (green check / red X)
3. **Commit SHA matches**: Verify the workflow is running against your latest commit

**Expected output**:
```
CI
Triggered by: push
Branch: 003-f005-github-actions
Status: In progress...
```

### 3. Monitor Workflow Execution

Click on the workflow run to see detailed logs:

```
https://github.com/jbeausoleil/sourcebox/actions/runs/[RUN_ID]
```

**What to check**:
1. **All 7 jobs appear**:
   - Test (ubuntu-latest, 1.21)
   - Test (ubuntu-latest, 1.22)
   - Test (macos-latest, 1.21)
   - Test (macos-latest, 1.22)
   - Test (windows-latest, 1.21)
   - Test (windows-latest, 1.22)
   - Lint

2. **Jobs run in parallel**: Matrix jobs start simultaneously (not sequentially)

3. **Logs show progress**:
   - ✅ Checkout code: Complete in <10 seconds
   - ✅ Setup Go: Complete in <30 seconds (first run), <5 seconds (cached)
   - ✅ Cache restore: "Cache restored from key: ..." (subsequent runs)
   - ✅ Download dependencies: Complete in <1 minute (first run), skipped (cached)
   - ✅ Run tests: Verbose output shows test results
   - ✅ Upload coverage: "Codecov report uploaded successfully"
   - ✅ Build: "Build complete: dist/sourcebox"

**Expected output per job**:
```
✓ Test (ubuntu-latest, 1.21)
  ✓ Set up job (2s)
  ✓ Checkout code (5s)
  ✓ Setup Go 1.21 (8s)
  ✓ Cache Go modules (2s) - Cache restored
  ✓ Download dependencies (0s) - Skipped (cache hit)
  ✓ Run tests (45s)
    ok      github.com/jbeausoleil/sourcebox/cmd/sourcebox  0.012s  coverage: 85.7%
  ✓ Upload coverage (3s)
  ✓ Build (12s)
  ✓ Complete job (1s)
Total: 1m 18s
```

### 4. Verify Coverage Upload

**In workflow logs**, find the "Upload coverage to Codecov" step:

```
Run codecov/codecov-action@v3
  with:
    files: ./coverage.txt
    flags: unittests
    name: codecov-ubuntu-latest-go1.21
```

**Expected output**:
```
[info] Uploading coverage report...
[info] Codecov report uploaded successfully
[info] View report: https://codecov.io/gh/jbeausoleil/sourcebox
```

**Navigate to Codecov**:
```
https://codecov.io/gh/jbeausoleil/sourcebox
```

**What to check**:
1. **Coverage percentage**: Should be >0% (initially low, will increase as tests are added)
2. **Coverage per matrix job**: 6 separate reports (ubuntu, macos, windows × 1.21, 1.22)
3. **File-level coverage**: Drill down to see which files/functions are covered

**Expected Codecov dashboard**:
- **Overall coverage**: Shows aggregated coverage from all 6 jobs
- **Trend graph**: Shows coverage over time (will populate after multiple commits)
- **Files**: Shows coverage per file (e.g., `cmd/sourcebox/main.go: 75%`)

### 5. Check Status Badges

**After workflow completes**, view README.md on GitHub:

```
https://github.com/jbeausoleil/sourcebox/blob/main/README.md
```

**Expected badges** (add to README.md if not present):

```markdown
[![CI](https://github.com/jbeausoleil/sourcebox/actions/workflows/ci.yml/badge.svg)](https://github.com/jbeausoleil/sourcebox/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/jbeausoleil/sourcebox/branch/main/graph/badge.svg)](https://codecov.io/gh/jbeausoleil/sourcebox)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/jbeausoleil/sourcebox)](go.mod)
```

**Badge status**:
- **CI badge**: Shows "passing" (green) or "failing" (red)
- **Codecov badge**: Shows coverage percentage (e.g., "75%")
- **License**: Shows "MIT" (static)
- **Go Version**: Shows "1.21" (from go.mod)

---

## PR Workflow Testing

### 1. Create Test PR

Create a test PR to verify CI runs on pull requests:

```bash
# Make a trivial change to test PR workflow
echo "# Test PR for CI/CD verification" >> docs/test-pr.md
git add docs/test-pr.md
git commit -m "Test: Verify CI runs on PR"

# Push to feature branch
git push origin 003-f005-github-actions
```

**Navigate to GitHub and create PR**:
```
Base: main
Compare: 003-f005-github-actions
Title: "Add GitHub Actions CI/CD pipeline (F005)"
```

### 2. Verify PR Checks - Comprehensive Guide

Once you create or update a pull request, GitHub displays status checks that validate your changes before merging. This section explains what to look for and how to interpret the results.

#### 2.1 Locating the Checks Section

**In the PR interface**, the Checks section appears in three places:

1. **At the top of the PR**: Next to the PR title, you'll see a status icon:
   - Yellow circle: Checks are running
   - Green checkmark: All checks passed
   - Red X: Some checks failed

2. **Bottom of the PR description**: Scroll down to see the detailed "Checks" section showing all individual checks

3. **Checks tab**: Click the "Checks" tab at the top of the PR to see detailed logs for each check

#### 2.2 Expected Checks Display

**All 7 checks should appear**:

```
CI / Test (ubuntu-latest, 1.21)
CI / Test (ubuntu-latest, 1.22)
CI / Test (macos-latest, 1.21)
CI / Test (macos-latest, 1.22)
CI / Test (windows-latest, 1.21)
CI / Test (windows-latest, 1.22)
CI / Lint
```

**Important**: All 7 checks are required for this pipeline. If you see fewer checks, the workflow may not be configured correctly.

#### 2.3 Check Status: Running (In Progress)

**What you'll see when checks are running**:

```
⏳ Some checks haven't completed yet

  ⏳ CI / Test (ubuntu-latest, 1.21) — In progress...
  ⏳ CI / Test (ubuntu-latest, 1.22) — In progress...
  ⏳ CI / Test (macos-latest, 1.21) — In progress...
  ⏳ CI / Test (macos-latest, 1.22) — In progress...
  ⏳ CI / Test (windows-latest, 1.21) — In progress...
  ⏳ CI / Test (windows-latest, 1.22) — In progress...
  ⏳ CI / Lint — In progress...
```

**Visual indicators**:
- Yellow/orange circle icon next to each check
- "In progress..." or "Queued" status text
- Time elapsed may show (e.g., "Running for 30 seconds")

**Merge button state**:
```
⚠ Merging is blocked
Some checks haven't completed yet
[Merge pull request] button is DISABLED (greyed out)
```

**What to do**:
- Wait for checks to complete (typically 2-5 minutes)
- Click on individual checks to view live logs
- Monitor progress in the Actions tab: `https://github.com/jbeausoleil/sourcebox/actions`

**Expected behavior**:
- Checks should start within 10-30 seconds of pushing
- Multiple checks run in parallel (not one at a time)
- Total time: 2-5 minutes for all checks to complete

#### 2.4 Check Status: Passing (Success)

**What you'll see when all checks pass**:

```
✓ All checks have passed
7 successful checks

  ✓ CI / Test (ubuntu-latest, 1.21) — 1m 15s
  ✓ CI / Test (ubuntu-latest, 1.22) — 1m 18s
  ✓ CI / Test (macos-latest, 1.21) — 1m 42s
  ✓ CI / Test (macos-latest, 1.22) — 1m 45s
  ✓ CI / Test (windows-latest, 1.21) — 2m 03s
  ✓ CI / Test (windows-latest, 1.22) — 2m 05s
  ✓ CI / Lint — 35s
```

**Visual indicators**:
- Green checkmark icon next to each check
- Duration displayed for each check (e.g., "1m 15s")
- "All checks have passed" summary at top

**Merge button state**:
```
✓ All checks have passed
This branch has no conflicts with the base branch

[Merge pull request] button is ENABLED (green)
```

**What to verify**:
1. **All 7 checks are green**: No skipped or failed checks
2. **Timing is reasonable**: Checks complete in <5 minutes total
3. **Coverage uploaded**: Look for coverage report link in check details
4. **Build succeeded**: Each test job includes successful build step

**Click "Details" on any check** to see:
- Detailed test output
- Coverage percentages
- Build artifacts
- Any warnings (even if check passed)

**Expected test output in check details**:
```
Run go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
=== RUN   TestExample
--- PASS: TestExample (0.00s)
PASS
ok      github.com/jbeausoleil/sourcebox/cmd/sourcebox  0.012s  coverage: 85.7%
```

#### 2.5 Check Status: Failing (Error)

**What you'll see when checks fail**:

```
✗ Some checks were not successful
1 failing and 6 successful checks

  ✗ CI / Test (ubuntu-latest, 1.21) — Failed in 45s
  ✗ CI / Test (ubuntu-latest, 1.22) — Failed in 47s
  ✗ CI / Test (macos-latest, 1.21) — Failed in 1m 12s
  ✗ CI / Test (macos-latest, 1.22) — Failed in 1m 15s
  ✗ CI / Test (windows-latest, 1.21) — Failed in 1m 32s
  ✗ CI / Test (windows-latest, 1.22) — Failed in 1m 35s
  ✓ CI / Lint — 35s
```

**Visual indicators**:
- Red X icon next to failed checks
- Green checkmark for passing checks (in this example, Lint passed)
- "Some checks were not successful" summary at top

**Merge button state**:
```
✗ Merging is blocked
Some checks were not successful
1 failing check

[Merge pull request] button is DISABLED (greyed out)
```

**Common failure scenarios**:

**Scenario 1: Test failures**
```
Click "Details" on failed check to see:

--- FAIL: TestExample (0.00s)
    main_test.go:10: expected "hello", got "goodbye"
FAIL
FAIL    github.com/jbeausoleil/sourcebox/cmd/sourcebox  0.012s
FAIL
Error: Process completed with exit code 1.
```

**What to do**:
1. Click "Details" on the failed check
2. Read the test failure message
3. Fix the failing test locally
4. Run `go test ./...` to verify fix
5. Commit and push the fix
6. Checks will re-run automatically

**Scenario 2: Lint failures**
```
Click "Details" on failed Lint check to see:

cmd/sourcebox/main.go:42:2: ineffectual assignment to err (ineffassign)
    err := doSomething()
    ^
```

**What to do**:
1. Run `golangci-lint run` locally (install if needed)
2. Fix the linting issues
3. Run again to verify
4. Commit and push the fix

**Scenario 3: Build failures**
```
Click "Details" on failed check to see:

./main.go:15:9: undefined: MissingFunction
make: *** [build] Error 1
Error: Process completed with exit code 2.
```

**What to do**:
1. Run `make build` locally to reproduce
2. Fix the build error
3. Verify build succeeds locally
4. Commit and push the fix

**Scenario 4: Race condition detected**
```
Click "Details" on failed check to see:

==================
WARNING: DATA RACE
Read at 0x00c0001a2000 by goroutine 7:
  main.processData()
      /home/runner/work/sourcebox/sourcebox/pkg/generator/data.go:42

Previous write at 0x00c0001a2000 by goroutine 6:
  main.updateData()
      /home/runner/work/sourcebox/sourcebox/pkg/generator/data.go:35
==================
```

**What to do**:
1. Run `go test -race ./...` locally to reproduce
2. Fix the data race (add mutexes, use channels, etc.)
3. Verify race is resolved locally
4. Commit and push the fix

#### 2.6 Check Status: Partial Failures

**What you'll see when some checks pass and some fail**:

```
✗ Some checks were not successful
2 failing and 5 successful checks

  ✓ CI / Test (ubuntu-latest, 1.21) — 1m 15s
  ✓ CI / Test (ubuntu-latest, 1.22) — 1m 18s
  ✓ CI / Test (macos-latest, 1.21) — 1m 42s
  ✓ CI / Test (macos-latest, 1.22) — 1m 45s
  ✗ CI / Test (windows-latest, 1.21) — Failed in 1m 32s
  ✗ CI / Test (windows-latest, 1.22) — Failed in 1m 35s
  ✓ CI / Lint — 35s
```

**What this indicates**:
- Platform-specific issue (in this example, Windows-only failure)
- Could be file path separators, line endings, or platform-specific code

**What to do**:
1. Compare failed check logs with passing check logs
2. Look for platform-specific code differences
3. Use `filepath.Join()` instead of hardcoded paths
4. Add `.gitattributes` for line ending normalization
5. Test locally with `GOOS=windows go test ./...` if possible

#### 2.7 Verifying Check Details

**For each passing or failing check**, click "Details" to verify:

1. **Test execution**:
   ```
   Run go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
   ```
   - Should show verbose test output
   - Should show coverage percentage per package
   - Should include race detector results

2. **Coverage upload**:
   ```
   Run codecov/codecov-action@v3
   [info] Uploading coverage report...
   [info] Codecov report uploaded successfully
   ```
   - Should show successful upload
   - Should include link to coverage report

3. **Build verification**:
   ```
   Run make build
   Building sourcebox...
   Build complete: dist/sourcebox
   ```
   - Should show successful build
   - Should create binary in dist/ directory

4. **Cache performance**:
   ```
   Cache restored from key: ubuntu-latest-go-1.21-xxx
   ```
   - First run: No cache message (expected)
   - Subsequent runs: Should show cache hit

#### 2.8 Re-Running Checks

**If checks fail due to transient issues** (network errors, flaky tests):

1. **Re-run all checks**:
   - Click "Re-run jobs" button (top right of Checks tab)
   - Select "Re-run all jobs"

2. **Re-run failed checks only**:
   - Click "Re-run jobs" button
   - Select "Re-run failed jobs"

**When to re-run**:
- Network errors (Codecov upload failed)
- GitHub Actions infrastructure issues
- Timeout errors (rare)

**When NOT to re-run**:
- Test failures (fix the code instead)
- Lint errors (fix the code instead)
- Build errors (fix the code instead)

#### 2.9 Branch Protection and Merge Blocking

**If branch protection is configured** (see "Configuring Branch Protection Rules" section):

**Expected behavior**:
- Merge button is DISABLED until all checks pass
- Warning message: "Merging is blocked - Required status checks must pass"
- Cannot merge even with admin privileges (if "Include administrators" enabled)

**Workflow**:
1. Create PR → Checks run automatically
2. Wait for checks to complete
3. If any check fails → Fix code, push update
4. Checks re-run automatically on new push
5. Once all checks pass → Merge button becomes ENABLED
6. Click "Merge pull request" to merge

**Override options** (not recommended):
- Admins can bypass protection if "Include administrators" is disabled
- Temporary rule disable (Settings → Branches → Edit rule)
- Use only for emergency hotfixes

### 3. Test Failing Workflow

Intentionally break a test to verify CI catches failures:

```bash
# Create a failing test
cat > cmd/sourcebox/main_test.go <<'EOF'
package main

import "testing"

func TestFailure(t *testing.T) {
    t.Error("Intentional failure to test CI")
}
EOF

git add cmd/sourcebox/main_test.go
git commit -m "Test: Intentional failure for CI verification"
git push origin 003-f005-github-actions
```

**Expected PR output**:
```
✗ Some checks were not successful
  ✗ CI / Test (ubuntu-latest, 1.21) — Failed
    FAIL    github.com/jbeausoleil/sourcebox/cmd/sourcebox  0.012s
    --- FAIL: TestFailure (0.00s)
        main_test.go:6: Intentional failure to test CI
  ... (similar failures on all 6 test jobs)
  ✓ CI / Lint — 35s
```

**What to check**:
1. **PR is blocked from merging**: Merge button shows "Merging is blocked"
2. **Clear failure message**: Test failure details visible in PR checks
3. **All matrix jobs show failure**: Failure affects all 6 test jobs (not just one platform)

**After verification, revert the failing test**:
```bash
git rm cmd/sourcebox/main_test.go
git commit -m "Revert: Remove intentional test failure"
git push origin 003-f005-github-actions
```

---

## Configuring Branch Protection Rules

**Purpose**: Enforce that all CI checks must pass before code can be merged into protected branches (main, develop). This prevents broken code from reaching production and maintains code quality standards.

**When to configure**: After verifying that the CI/CD workflow is stable and all checks pass reliably. Recommended to configure after completing User Story 2 (PR Quality Gates).

### Step-by-Step Setup Guide

#### 1. Navigate to Repository Settings

```
1. Go to your repository: https://github.com/jbeausoleil/sourcebox
2. Click the "Settings" tab (top right, requires admin/write access)
3. In left sidebar, click "Branches" under "Code and automation"
```

**Screenshot reference**: Settings → Branches page shows "Branch protection rules" section

#### 2. Add Branch Protection Rule

```
1. Click "Add rule" button (or "Add branch protection rule")
2. In "Branch name pattern" field, enter: main
3. This creates a protection rule for the main branch
```

**Alternative patterns**:
- `main` - Protects only the main branch
- `develop` - Protects only the develop branch
- `main|develop` - Protects both branches (regex pattern)
- `release/*` - Protects all release branches

**Recommendation**: Start with `main` only, add `develop` later if using a develop branch workflow.

#### 3. Configure Required Status Checks

Scroll down to the "Protect matching branches" section and configure:

##### 3a. Enable Status Check Requirement

```
☑ Require status checks to pass before merging
```

**What this does**: Blocks merging if any required status check fails. PRs must have all checks passing (green) before the merge button is enabled.

##### 3b. Enable Branch Freshness Requirement

```
☑ Require branches to be up to date before merging
```

**What this does**: Ensures the PR branch has the latest changes from the target branch (main) merged in before allowing merge. Prevents integration issues where code passes tests on an old version of main but fails on the current version.

**When to enable**:
- **Recommended**: Yes, for multi-developer projects
- **Optional**: Can be disabled for solo development if rebasing is tedious
- **Trade-off**: Extra rebasing effort vs. higher confidence in integration

##### 3c. Select Required Status Checks

After enabling "Require status checks to pass before merging", a search box appears:

```
Search for status checks in the last week for this repository
```

**How to populate the list**:
1. The search box only shows status checks that have run at least once in the past week
2. If the list is empty, push a commit or create a PR first to trigger the workflow
3. Wait for the workflow to complete (all 7 jobs)
4. Refresh the branch protection settings page
5. The status checks should now appear in the search results

**Required checks to select** (select ALL 7):

```
☑ Test (ubuntu-latest, 1.21)
☑ Test (ubuntu-latest, 1.22)
☑ Test (macos-latest, 1.21)
☑ Test (macos-latest, 1.22)
☑ Test (windows-latest, 1.21)
☑ Test (windows-latest, 1.22)
☑ Lint
```

**Why all 7?**: Requiring all matrix combinations ensures cross-platform compatibility. If only one check is required (e.g., ubuntu-latest, 1.21), code could break on macOS or Windows without blocking the merge.

**Partial selection risk**: If you only require `ubuntu-latest` checks, code could merge with failures on macOS or Windows, defeating the purpose of multi-platform testing.

##### 3d. Additional Protection Options (Recommended)

```
☑ Require a pull request before merging
  ☑ Require approvals: 1 (for multi-developer teams)
  ☐ Dismiss stale pull request approvals when new commits are pushed (optional)
  ☐ Require review from Code Owners (if CODEOWNERS file exists)

☐ Require conversation resolution before merging (optional, good for teams)

☐ Require signed commits (optional, higher security requirement)

☑ Require linear history (prevents merge commits, enforces rebase/squash)
  - Recommendation: Enable for clean Git history

☐ Include administrators (optional, applies rules to admins too)
  - Recommendation: Disable for solo projects (you need override ability)
  - Recommendation: Enable for team projects (enforce rules on everyone)

☑ Allow force pushes (select "Specify who can force push")
  - Select "Nobody" to prevent force pushes to main
  - Prevents accidental history rewrites

☑ Allow deletions
  - Recommendation: Disable (prevent accidental branch deletion)
```

#### 4. Save Protection Rule

```
1. Scroll to bottom of page
2. Click "Create" (for new rule) or "Save changes" (for existing rule)
3. Verify green success banner appears: "Branch protection rule created"
```

#### 5. Verify Protection is Active

**Test 1: View protected branch indicator**

```
1. Navigate to repository home page
2. Click "main" branch dropdown
3. Look for shield icon next to "main" branch name
```

**Expected**: Shield icon indicates branch is protected

**Test 2: Attempt to merge PR with failing checks**

```
1. Create a PR with failing tests (see "Test Failing Workflow" section)
2. Navigate to PR page
3. Scroll to merge button
```

**Expected behavior**:
```
⚠ Merging is blocked
Some checks were not successful
7 failing and 0 successful checks

Required status checks must pass:
  ✗ CI / Test (ubuntu-latest, 1.21) — Failed
  ✗ CI / Test (ubuntu-latest, 1.22) — Failed
  ... (all 7 checks failed)

[Merge pull request] button is DISABLED (greyed out)
```

**Test 3: Attempt to merge PR with passing checks**

```
1. Fix the failing tests
2. Push to PR branch
3. Wait for all 7 checks to pass
4. Navigate to PR page
```

**Expected behavior**:
```
✓ All checks have passed
7 successful checks

  ✓ CI / Test (ubuntu-latest, 1.21) — 1m 15s
  ... (all 7 checks passed)

[Merge pull request] button is ENABLED (green)
```

**Test 4: Attempt to merge outdated PR (if "up to date" required)**

```
1. Have a PR open with passing checks
2. Merge another PR to main first (or push directly to main)
3. Navigate back to first PR
```

**Expected behavior**:
```
⚠ This branch is out of date with the base branch
Merge the latest changes from main into this branch

[Update branch] button appears
[Merge pull request] button is DISABLED until branch is updated
```

### Status Check Names Reference

**How to find exact check names**:

If you're unsure of the exact status check names, follow these steps:

```
1. Navigate to a recent PR or push that triggered the workflow
2. Go to the "Checks" tab or scroll to "Checks" section in PR
3. Copy the exact job names as they appear (case-sensitive)
4. Use these names when configuring branch protection
```

**Standard check names for this workflow**:
- Pattern: `CI / <job-name>` (note the space after slash)
- Test jobs: `CI / Test (<os>, <go-version>)`
- Lint job: `CI / Lint`

**Common mistakes**:
- ❌ `Test (ubuntu-latest, 1.21)` (missing "CI / " prefix)
- ❌ `CI/Test (ubuntu-latest, 1.21)` (no space after slash)
- ✅ `CI / Test (ubuntu-latest, 1.21)` (correct format)

### Expected Behavior After Configuration

#### For Pull Requests

**Before checks complete**:
```
Status: Some checks haven't completed yet
Merge button: DISABLED
Message: "Merging is blocked - Waiting for status to be reported"
```

**After checks pass**:
```
Status: ✓ All checks have passed
Merge button: ENABLED (green)
Message: "This branch has no conflicts with the base branch"
```

**If any check fails**:
```
Status: ✗ Some checks were not successful
Merge button: DISABLED
Message: "Merging is blocked - 1 failing check"
Developer action required: Fix code, push new commit, wait for checks to pass
```

**If branch is out of date (when "up to date" required)**:
```
Status: ⚠ This branch is out of date
Merge button: DISABLED
Message: "This branch is out-of-date with the base branch"
Developer action required: Click "Update branch" to merge main into PR branch
```

#### For Direct Pushes to Main

**With protection enabled**:
```
❌ Direct pushes to main are BLOCKED (even for admins, unless override configured)
Error message: "Branch 'main' is protected"
Required action: Create PR, get checks passing, then merge
```

**Without "Include administrators" enabled**:
```
✅ Admins can still push directly (bypass protection)
⚠ Warning: Use this sparingly (emergency hotfixes only)
Recommendation: Prefer PR workflow even for admins
```

### Troubleshooting Branch Protection

#### Issue: Status checks don't appear in search box

**Symptom**: Search box for status checks is empty or doesn't show CI jobs

**Solution**:
1. Verify workflow has run at least once in the past week
2. Push a commit or create a PR to trigger workflow
3. Wait for workflow to complete fully
4. Refresh branch protection settings page
5. Search again for status checks

**Root cause**: GitHub only shows status checks that have reported status in the last 7 days

#### Issue: Merge button still enabled despite failing checks

**Symptom**: PR can be merged even though status checks failed

**Possible causes**:
1. Branch protection not configured correctly
2. Wrong branch pattern (e.g., `main` vs `master`)
3. Required status checks not selected
4. User has admin override permissions ("Include administrators" not enabled)

**Solution**:
1. Navigate to Settings → Branches
2. Verify rule exists for correct branch name
3. Edit rule and verify "Require status checks" is checked
4. Verify all 7 checks are selected in required status checks list
5. Consider enabling "Include administrators" to enforce on all users

#### Issue: Checks pass but merge still blocked

**Symptom**: All 7 checks show green, but merge button is disabled

**Possible causes**:
1. Branch is out of date (if "Require branches to be up to date" is enabled)
2. Required approvals not met (if "Require approvals" is configured)
3. Unresolved conversations (if "Require conversation resolution" is enabled)
4. Different status check names (e.g., workflow renamed)

**Solution**:
1. Check for warning messages above merge button (explains why blocked)
2. Click "Update branch" if out of date
3. Request review/approval if required
4. Resolve conversations if required
5. Verify status check names match workflow job names exactly

#### Issue: Protection prevents emergency fixes

**Symptom**: Need to push critical hotfix but CI is failing due to unrelated issue

**Emergency bypass options**:
1. **Temporary rule disable** (admin only):
   - Settings → Branches → Edit rule
   - Uncheck "Require status checks"
   - Merge hotfix
   - Re-enable protection immediately after

2. **Admin override** (if "Include administrators" not enabled):
   - Admins can push directly to protected branch
   - Use with extreme caution
   - Create follow-up PR to fix properly

3. **Hotfix branch pattern exclusion**:
   - Configure protection rule to exclude `hotfix/*` pattern
   - Push to `hotfix/critical-fix` branch
   - Merge to main (bypasses protection)
   - Use only for true emergencies

**Recommendation**: Fix CI issues properly rather than bypassing protection. Emergency bypasses should be rare (<1% of merges).

### Maintenance and Updates

#### When to update protection rules

**Add new checks**:
- If workflow adds new required jobs (e.g., security scanning, integration tests)
- Add to required status checks list

**Rename checks**:
- If workflow job names change (e.g., `Test` → `Unit Test`)
- Remove old check names, add new check names
- Test with a PR to verify new names work

**Add protected branches**:
- If adding `develop`, `release/*`, or other protected branches
- Create separate rule or update pattern to `main|develop`

#### Reviewing protection effectiveness

**Monthly review checklist**:
- [ ] Are all required checks still running? (no deprecated checks)
- [ ] Are checks completing in reasonable time? (<5 min target)
- [ ] Are PRs being blocked appropriately? (failures prevent merge)
- [ ] Are admins following protection rules? (no excessive bypasses)
- [ ] Is branch protection preventing incidents? (broken code reaching main)

**Metrics to track**:
- % of PRs merged with all checks passing (target: 100%)
- % of commits to main via PR vs direct push (target: >95% via PR)
- Number of admin bypasses per month (target: <5%)
- Average time PR blocked due to failing checks (indicates test quality)

---

## Debugging Failed Workflows

### Common Issue 1: Workflow Doesn't Trigger

**Symptom**: No workflow run appears in Actions tab after pushing commit.

**Check**:
1. **Actions enabled**: Navigate to `Settings → Actions → General`. Verify "Allow all actions and reusable workflows" is selected.
2. **Workflow file location**: Verify file is at `.github/workflows/ci.yml` (not `github/workflows` or `.github/workflow`)
3. **YAML syntax**: Validate YAML at https://www.yamllint.com/. Incorrect indentation breaks parsing.

**Fix**:
- Enable Actions in repository settings
- Move workflow file to correct location
- Fix YAML syntax errors (check indentation, quotes)

---

### Common Issue 2: Tests Fail in CI but Pass Locally

**Symptom**: `go test ./...` passes locally, but CI shows failures.

**Check**:
1. **Platform-specific code**: Check if code uses platform-specific APIs (file paths, line endings)
2. **Race conditions**: Run `go test -race ./...` locally to reproduce
3. **Environment differences**: Check for hardcoded paths, missing environment variables

**Fix**:
- **File paths**: Use `filepath.Join()` instead of string concatenation (`path/to/file` → `filepath.Join("path", "to", "file")`)
- **Line endings**: Git may convert CRLF (Windows) ↔ LF (Linux/Mac). Add `.gitattributes` with `* text=auto`
- **Race conditions**: Fix data races detected by `-race` flag (add mutexes, use channels)

**Example**:
```bash
# Reproduce CI failure locally
GOOS=windows GOARCH=amd64 go test -race ./...
```

---

### Common Issue 3: Coverage Upload Fails

**Symptom**: "Upload coverage to Codecov" step fails with error.

**Check**:
1. **Codecov token**: For public repos, no token needed. For private repos, check if `CODECOV_TOKEN` secret is configured.
2. **Coverage file exists**: Verify `coverage.txt` is generated by `go test -coverprofile=coverage.txt`
3. **Network issues**: Transient Codecov API errors (retry workflow)

**Fix**:
- **Private repo**: Add `CODECOV_TOKEN` secret in `Settings → Secrets and variables → Actions`
- **Coverage file missing**: Check if tests are running successfully (coverage file only generated if tests pass)
- **Network error**: Re-run failed jobs in Actions tab

---

### Common Issue 4: Workflow Timeout (>5 minutes)

**Symptom**: Workflow exceeds 5-minute target, times out, or is very slow.

**Check**:
1. **Cache miss**: Check if "Cache restored" message appears in logs. No cache = full `go mod download`.
2. **Large dependencies**: Check `go.sum` for unexpected large dependencies.
3. **Slow tests**: Check test logs for slow tests (>10 seconds per package).

**Fix**:
- **Cache miss**: Verify `go.sum` hasn't changed. If intentional (new dependency), first run will be slow but subsequent runs will be fast.
- **Large dependencies**: Review dependencies with `go mod graph | grep <dep-name>`. Consider lighter alternatives.
- **Slow tests**: Optimize slow tests or use `t.Parallel()` to parallelize within package.

**Monitor workflow duration**:
```
Actions tab → Workflow run → View total duration in upper right
Target: <5 min for 90% of runs
```

---

### Common Issue 5: golangci-lint Failures

**Symptom**: Lint job fails with warnings/errors.

**Check**:
1. **New linting rules**: golangci-lint may enable new linters on updates.
2. **Code style issues**: Check specific warnings in lint job logs.

**Fix**:
- **Run locally**: `golangci-lint run` (install: https://golangci-lint.run/usage/install/)
- **Fix issues**: Address warnings (e.g., unused variables, ineffectual assignments)
- **Disable specific linter** (if false positive): Add `.golangci.yml`:
  ```yaml
  linters:
    disable:
      - errcheck  # Example: disable errcheck if too noisy
  ```

**Example lint failure**:
```
cmd/sourcebox/main.go:42:2: ineffectual assignment to err (ineffassign)
    err := doSomething()
    ^
```

**Fix**:
```go
// Before
err := doSomething()

// After (check error or remove if truly unused)
if err := doSomething(); err != nil {
    return err
}
```

---

## Verification Checklist

Use this checklist to verify F005 is fully functional:

- [ ] **Workflow file exists**: `.github/workflows/ci.yml` is present and committed
- [ ] **Workflow triggers on push**: Push to any branch triggers workflow
- [ ] **Workflow triggers on PR**: Creating/updating PR triggers workflow
- [ ] **All 6 test matrix jobs pass**: ubuntu, macos, windows × Go 1.21, 1.22
- [ ] **Lint job passes**: golangci-lint runs with zero warnings
- [ ] **Coverage uploads to Codecov**: Coverage reports visible at codecov.io/gh/jbeausoleil/sourcebox
- [ ] **Status badges display correctly**: CI badge shows "passing" (green) in README.md
- [ ] **Workflow completes in <5 minutes**: Cached runs finish in 2-4 minutes
- [ ] **PR checks appear**: PR shows all 7 checks (Test × 6, Lint × 1)
- [ ] **Failed tests block PR**: Intentional test failure prevents merge
- [ ] **Cache works**: "Cache restored" message appears on second run

**Additional verification** (optional but recommended):
- [ ] **Branch protection configured**: Required status checks prevent merging broken code
- [ ] **Coverage trends tracked**: Codecov shows coverage trend graph after 3+ commits
- [ ] **Platform-specific tests pass**: No failures unique to Windows/macOS/Linux
- [ ] **Race detector enabled**: Tests run with `-race` flag (check logs)

---

## Performance Monitoring

### Baseline Metrics (Expected)

- **First run** (no cache): 5-8 minutes total
  - Checkout: 5-10 seconds
  - Setup Go: 20-30 seconds
  - Download dependencies: 1-2 minutes
  - Run tests: 1-2 minutes
  - Upload coverage: 5-10 seconds
  - Build: 30-60 seconds

- **Subsequent runs** (cache hit): 2-4 minutes total
  - Checkout: 5-10 seconds
  - Setup Go: 5-10 seconds (cached)
  - Download dependencies: 0 seconds (skipped, cache hit)
  - Run tests: 1-2 minutes
  - Upload coverage: 5-10 seconds
  - Build: 30-60 seconds

- **Lint job**: 30-60 seconds total

### Tracking Workflow Duration

**In GitHub Actions tab**:
1. Navigate to `Actions → CI`
2. View workflow history (sorted by date)
3. Check duration for each run (upper right of each workflow)

**Monitor for regressions**:
- If duration consistently exceeds 5 minutes → investigate
- If cache stops working → check `go.sum` changes
- If specific job is slow → check logs for slow tests

**Example workflow history**:
```
✓ CI - Add feature X                   3m 42s
✓ CI - Fix bug Y                        2m 15s (cache hit)
✓ CI - Update dependencies              6m 03s (cache miss, go.sum changed)
✓ CI - Refactor Z                       2m 35s (cache hit)
```

**Target**: 90% of workflow runs complete in <5 minutes (with caching).

---

## Next Steps

After F005 is complete and verified:

1. **F006 - Cobra CLI Framework**: Tests will run via this CI/CD pipeline automatically
2. **F009 - Add Dependencies**: Workflow will cache new dependencies, reducing run time
3. **F020 - Data Generation Tests**: Workflow will enforce >80% coverage target for core packages

**Branch protection setup** (recommended after workflow is stable):
- See "Configuring Branch Protection Rules" section above for detailed setup guide
- Enforce all 7 status checks (Test × 6, Lint × 1) before merging
- Require branches to be up to date before merging
- Prevent direct pushes to main (require PR workflow)

**Expected behavior after branch protection**:
- PRs cannot be merged if any check fails
- PRs must be rebased/merged with latest `main` before merging
- Developers receive immediate feedback on code quality
- Main branch always contains tested, working code

---

## Troubleshooting Reference

### Workflow YAML Syntax Errors

**Symptom**: Workflow doesn't appear in Actions tab, or shows "Invalid workflow file".

**Debug**:
1. Validate YAML syntax: https://www.yamllint.com/
2. Check indentation (spaces only, no tabs)
3. Verify action versions exist (e.g., `actions/checkout@v4`)

**Common errors**:
```yaml
# ❌ Incorrect (tab indentation)
jobs:
	test:
		runs-on: ubuntu-latest

# ✅ Correct (space indentation)
jobs:
  test:
    runs-on: ubuntu-latest
```

---

### Cache Misses

**Symptom**: Workflow takes 5-8 minutes every run (no cache hit).

**Debug**:
1. Check workflow logs for "Cache restored from key: ..." message
2. If missing → cache miss (expected on first run or `go.sum` change)
3. If consistently missing → cache key misconfigured

**Common causes**:
- **`go.sum` changed**: New dependency or version update invalidates cache (expected)
- **Cache key incorrect**: Verify key is `${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}`
- **Cache corrupted**: Manually clear cache in `Settings → Actions → Caches`

**Fix**:
- **Intentional `go.sum` change**: First run will be slow, subsequent runs fast
- **Cache key mismatch**: Fix workflow file, push update
- **Corrupted cache**: Delete cache in Settings, re-run workflow

---

### Coverage Gaps

**Symptom**: Codecov shows <80% coverage for core packages.

**Debug**:
1. Run locally: `go test -coverprofile=coverage.txt ./...`
2. View coverage: `go tool cover -html=coverage.txt`
3. Identify untested functions/branches

**Fix**:
- Write tests for uncovered code paths
- Target: >80% for `pkg/*` packages (core logic)
- Acceptable: 60-80% for `cmd/*` (CLI glue code)

**Example**:
```bash
# Generate coverage report
go test -coverprofile=coverage.txt ./...

# View in browser
go tool cover -html=coverage.txt

# Look for red (uncovered) lines in HTML report
```

---

### Platform-Specific Failures

**Symptom**: Tests pass on Linux/macOS but fail on Windows (or vice versa).

**Common causes**:
1. **File path separators**: `/` (Linux/macOS) vs `\` (Windows)
2. **Line endings**: LF (Linux/macOS) vs CRLF (Windows)
3. **Case sensitivity**: Linux/macOS filesystem case-sensitive, Windows not
4. **File permissions**: Different permission models across platforms

**Fix**:
- **File paths**: Use `filepath.Join("path", "to", "file")` instead of `"path/to/file"`
- **Line endings**: Add `.gitattributes` with `* text=auto` to normalize line endings
- **Case sensitivity**: Ensure file name case matches in code and filesystem
- **Permissions**: Use `os.Chmod()` conditionally (check `runtime.GOOS`)

**Example**:
```go
// ❌ Platform-specific (breaks on Windows)
configPath := "/etc/sourcebox/config.yaml"

// ✅ Cross-platform
import "path/filepath"
configPath := filepath.Join("/", "etc", "sourcebox", "config.yaml")

// ✅ Even better (use user config dir)
import "os"
homeDir, _ := os.UserHomeDir()
configPath := filepath.Join(homeDir, ".sourcebox", "config.yaml")
```

---

### ARM Runners (Future Consideration)

**Note**: GitHub-hosted ARM runners are not part of the free tier. If ARM-specific testing is needed in the future:

**Options**:
1. **Self-hosted ARM runners**: Set up Raspberry Pi or ARM server as GitHub Actions runner
2. **Paid ARM runners**: GitHub Enterprise includes ARM runner minutes
3. **Cross-compilation only**: Build for ARM without running tests (acceptable for CLI tools)

**Current approach**: Test on standard x86_64 runners (ubuntu, macos, windows). Cross-compile for ARM64 in release workflow (future F-TBD).

---

## Summary

F005 establishes automated CI/CD pipeline with:

- ✅ **Automatic test execution** on every push and pull request
- ✅ **Multi-platform testing** (macOS, Linux, Windows × Go 1.21, 1.22)
- ✅ **Code quality enforcement** (golangci-lint with zero warnings)
- ✅ **Coverage tracking** (>80% target for core packages via Codecov)
- ✅ **Fast feedback** (<5 min workflow execution with caching)
- ✅ **Transparent status** (badges in README, PR checks)
- ✅ **Quality gates** (branch protection prevents merging broken code)

**Constitutional Alignment**:
- **Boring Tech Wins**: Standard GitHub Actions, golangci-lint, Codecov
- **Speed > Features**: <5 min workflow, caching enabled, parallel execution
- **Cost Constraints**: Free tier (GitHub Actions 2,000 min/month, Codecov free for open source)
- **Platform Support**: All 3 platforms, 2 Go versions
- **Code Quality**: >80% coverage, race detector, zero lint warnings
- **Developer-First**: Fast feedback, clear badges, actionable errors
- **Ship Fast, Validate Early**: Automated testing prevents regressions, enables rapid iteration

**Developer Workflow**:
1. Write code locally
2. Run `go test ./...` and `make build` locally (TDD)
3. Push to GitHub
4. CI runs automatically (< 5 min feedback)
5. Fix any failures shown in PR checks
6. Merge when all checks pass

This CI/CD pipeline is the foundation for all future development (F006-F020+).
