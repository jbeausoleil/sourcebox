# F003 Quickstart: Repository Setup & Verification

**Feature**: F003 - Initialize Git Repository & Go Module
**Purpose**: Verify the foundational Git repository and Go module setup is correct
**Target**: Developers cloning SourceBox for the first time

---

## Prerequisites

Before starting, ensure you have the following installed:

- **Git 2.0+**: `git --version`
- **Go 1.21+**: `go version`
- **GitHub account**: For forking and contributing (optional for read-only access)

### Installation Links
- Git: [https://git-scm.com/downloads](https://git-scm.com/downloads)
- Go: [https://golang.org/dl/](https://golang.org/dl/)

---

## Setup Steps

### 1. Clone Repository

```bash
# HTTPS (recommended for read-only)
git clone https://github.com/jbeausoleil/sourcebox.git
cd sourcebox

# SSH (recommended for contributors)
git clone git@github.com:jbeausoleil/sourcebox.git
cd sourcebox
```

**Expected**: Clone completes in < 30 seconds on typical broadband connection

**Troubleshooting**:
- If SSH fails: Ensure SSH keys are configured in GitHub
- If HTTPS fails: Check network connectivity

### 2. Verify Go Module

```bash
# Download module dependencies (even with no deps, verifies module configuration)
go mod download

# Verify module path
go list -m
# Expected output: github.com/jbeausoleil/sourcebox

# Check Go version requirement
grep "^go " go.mod
# Expected output: go 1.21
```

**Expected**: All commands complete successfully with no errors

**Troubleshooting**:
- If `go mod download` fails: Ensure Go 1.21+ is installed
- If module path is wrong: Check go.mod file syntax

### 3. Review Documentation

Open and read the following files to understand the project:

```bash
# Essential reading
cat README.md              # Project overview and legal notice
cat LICENSE                # MIT license terms
cat CONTRIBUTING.md        # How to contribute
cat CODE_OF_CONDUCT.md     # Community standards

# Or use your preferred editor/viewer
```

**Key things to verify**:
- README displays legal independence notice prominently
- LICENSE contains MIT license text
- CONTRIBUTING.md has clear instructions for bugs, features, and PRs
- CODE_OF_CONDUCT.md contains Contributor Covenant v2.1

### 4. Understand Contribution Process

After reading CONTRIBUTING.md, you should be able to answer:

1. **How do I report a bug?**
   - Create GitHub issue with reproduction steps, environment, and error logs

2. **How do I request a feature?**
   - Create GitHub issue with problem statement, current workaround, and expected benefit

3. **How do I submit code?**
   - Fork repo → create feature branch → follow TDD → run tests → create PR

---

## Verification Checklist

Run through this checklist to verify the repository setup is correct:

### Git Repository
- [ ] `git clone` completed successfully
- [ ] Repository size is < 1MB (verify with `du -sh .git`)
- [ ] `.gitignore` file exists and contains Go-specific patterns
- [ ] No binary files, build artifacts, or IDE files are committed

### Go Module
- [ ] `go mod download` runs without errors
- [ ] `go list -m` outputs `github.com/jbeausoleil/sourcebox`
- [ ] `go.mod` specifies Go 1.21 or higher
- [ ] No dependencies listed in go.mod yet (foundational feature)

### Documentation
- [ ] README.md renders correctly on GitHub (check on web)
- [ ] README.md contains legal independence notice near top
- [ ] LICENSE file contains full MIT license text
- [ ] CONTRIBUTING.md has clear instructions for contributing
- [ ] CODE_OF_CONDUCT.md contains Contributor Covenant v2.1

### Legal Compliance
- [ ] Legal notice is visible in README (above the fold)
- [ ] No employer references in any documentation files
- [ ] No proprietary information in any files
- [ ] No credentials, secrets, or .env files committed (check with `git log --all --full-history -- .env`)

### Repository Health
- [ ] All documentation files use proper markdown formatting
- [ ] No broken links in README, CONTRIBUTING, or other docs
- [ ] Repository is publicly accessible (no authentication required to clone)
- [ ] Default branch is `main` or `master` (verify with `git branch`)

---

## Quick Reference Commands

```bash
# Verify repository state
git status                    # Check for uncommitted changes
git log --oneline -10         # View recent commits
git remote -v                 # Verify remote URL

# Verify Go setup
go version                    # Check Go version (should be 1.21+)
go mod verify                 # Verify module integrity
go list -m all               # List all modules (should be just sourcebox initially)

# Run basic checks
go vet ./...                  # Static analysis (no code yet, so should pass quickly)
go fmt ./...                  # Format code (no code yet, so no changes)

# Repository size check
du -sh .git                   # Should be < 1MB
du -sh .                      # Total size should be < 1MB

# Search for potential issues
git log --all --full-history -- .env     # Should find nothing (no secrets)
grep -r "TODO\|FIXME" .                  # Check for pending work markers
```

---

## Performance Benchmarks

As per spec success criteria (SC-001), verify performance meets targets:

| Operation | Target | How to Measure |
|-----------|--------|----------------|
| Clone repository | < 2 minutes | Time `git clone` command (includes network transfer) |
| Setup verification | < 30 seconds | Time running verification checklist commands |
| Repository size | < 1MB | Run `du -sh .git` |

**Example timing**:
```bash
# Measure clone time
time git clone https://github.com/jbeausoleil/sourcebox.git
# Expected: real < 2m (typically < 10s on good connection)

# Measure verification time
time (go mod download && go list -m && cat README.md LICENSE CONTRIBUTING.md CODE_OF_CONDUCT.md > /dev/null)
# Expected: real < 30s (typically < 5s)
```

---

## Common Issues & Solutions

### Issue: `go mod download` fails with "invalid version"
**Solution**: Ensure Go 1.21+ is installed. Run `go version` and upgrade if needed.

### Issue: Git clone is very slow
**Solution**: Check network connection. Try HTTPS instead of SSH or vice versa.

### Issue: Can't see legal notice in README on GitHub
**Solution**: Clear browser cache or view raw README.md. Legal notice should be near top.

### Issue: `.gitignore` is not ignoring IDE files
**Solution**: Verify `.gitignore` contains `.vscode/`, `.idea/`, etc. Restart IDE if needed.

### Issue: Repository size exceeds 1MB
**Solution**: Check for accidentally committed binaries with `git log --all --stat`. Remove with `git filter-branch` if needed.

---

## Next Steps

After completing this quickstart and verification:

1. **For Contributors**:
   - Read CONTRIBUTING.md thoroughly
   - Review open issues on GitHub
   - Join discussions and propose features
   - Follow the Spec-Kit workflow when contributing code

2. **For Development**:
   - Review F004 (Project Directory Structure) to understand source code layout
   - Review F005 (CI/CD Setup) to understand automated testing
   - Review F009 (Dependency Management) to understand how dependencies are added

3. **For Understanding SourceBox**:
   - Read the constitution at `.specify/memory/constitution.md`
   - Review feature specifications in `/specs/` directory
   - Understand the 12-week MVP plan in documentation

---

## Related Documentation

- **Feature Spec**: `specs/001-f003-initialize-git/spec.md`
- **Research**: `specs/001-f003-initialize-git/research.md`
- **Implementation Plan**: `specs/001-f003-initialize-git/plan.md`
- **Constitution**: `.specify/memory/constitution.md`
- **Spec-Kit Commands**: `.claude/commands/speckit.*.md`

---

## Success Criteria Met?

After completing this quickstart, you should be able to confidently answer:

- ✅ **What is SourceBox?** → Verticalized demo data for developers in 30 seconds
- ✅ **What license is it under?** → MIT (free forever, commercial use allowed)
- ✅ **How do I contribute?** → Follow CONTRIBUTING.md (bugs, features, PRs)
- ✅ **Is it independently developed?** → Yes, legal notice in README confirms this
- ✅ **Can I use it in my project?** → Yes, MIT license permits commercial use
- ✅ **Where do I report issues?** → GitHub issues following CONTRIBUTING.md guidelines

If you can answer all these questions, you've successfully verified F003! 🎉

---

**Quickstart Version**: 1.0.0
**Last Updated**: 2025-10-14
**Feedback**: Report issues or improvements in GitHub issues
