# Research: F003 - Initialize Git Repository & Go Module

**Feature**: F003 - Initialize Git Repository & Go Module
**Branch**: `001-f003-initialize-git`
**Date**: 2025-10-14
**Status**: Complete

## Overview

This document captures technical research and decisions for initializing the SourceBox Git repository with proper Go module configuration, documentation, and legal protections.

---

## Decision 1: Go .gitignore Patterns

### Decision
Use standard Go community .gitignore patterns from [github/gitignore](https://github.com/github/gitignore/blob/main/Go.gitignore) with additions for common developer tools and environment files.

### Rationale
- **Community Standard**: The github/gitignore repository is the de facto standard used by GitHub's own template system
- **Comprehensive Coverage**: Includes patterns for Go binaries, test artifacts, coverage files, and build outputs
- **Proven in Production**: Used by thousands of Go projects successfully
- **Aligns with Constitution Principle IV**: "Boring Tech Wins" - use proven, standard patterns

### Implementation
```gitignore
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib
sourcebox

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool
*.out
coverage.txt
coverage.html

# Dependency directories (Go modules)
vendor/

# Build artifacts
/dist/
/build/

# IDE files
.vscode/
.idea/
*.swp
*.swo
*~

# OS files
.DS_Store
Thumbs.db

# Environment files (CRITICAL: prevent credential leaks)
.env
.env.local
.env.*.local
```

### Alternatives Considered
- **Custom minimal .gitignore**: Rejected because incomplete coverage risks committing build artifacts
- **Language-agnostic .gitignore**: Rejected because Go-specific patterns (vendor/, *.test) would be missing
- **No .gitignore**: Rejected because accidental commits of binaries and secrets would bloat repository

---

## Decision 2: Go Module Configuration

### Decision
- **Module Path**: `github.com/jbeausoleil/sourcebox`
- **Minimum Go Version**: `1.21`
- **Initial Dependencies**: None (foundational feature)

### Rationale

**Module Path**:
- Standard format: `domain/username/project`
- Maps to actual GitHub repository location
- Enables future `go get github.com/jbeausoleil/sourcebox` imports
- Follows golang.org module path conventions

**Go Version 1.21**:
- Released August 2023 (stable, well-adopted)
- Includes modern features: min/max builtins, clear() builtin, improved type inference
- Balances modern features with reasonable adoption (not bleeding-edge)
- Supported by all major Go tooling (GoLand, VS Code, gopls)
- Current LTS version with active security support

**No Dependencies Initially**:
- Foundational feature focuses on repository structure only
- Dependencies will be added in F009 (Dependency Management)
- Keeps go.mod minimal and focused
- Avoids premature dependency decisions

### Implementation
```go
module github.com/jbeausoleil/sourcebox

go 1.21
```

### Alternatives Considered
- **Go 1.22/1.23**: Rejected as too new; adoption not widespread enough (violates "boring tech wins")
- **Go 1.20**: Rejected as missing useful modern features; 1.21 is better balanced
- **Generic module path**: Rejected because non-standard paths cause import issues

---

## Decision 3: README Structure & Badges

### Decision
README.md with the following structure and badges:

**Badges** (in order):
1. Build Status (GitHub Actions) - placeholder until F005
2. License (MIT)
3. Go Version (1.21+)

**Sections**:
1. Project name + tagline
2. Legal independence notice (prominently placed)
3. Problem statement (why SourceBox exists)
4. Features (brief bulleted list)
5. Installation (placeholder for future)
6. Quick start (placeholder for future)
7. Contributing
8. License

### Rationale

**Badge Selection**:
- **Build Status**: Standard for open source projects; shows health
- **License Badge**: Critical for developers evaluating legal usage
- **Go Version Badge**: Helps developers verify compatibility
- **Excluded**: Code coverage (premature), downloads (no releases yet), stars (vanity metric)

**Legal Notice Placement**:
- **Above the fold** (near top of README)
- Visible without scrolling on GitHub
- Protects project's independent development status
- Satisfies Constitution Legal Constraint 1

**Structure Rationale**:
- **Problem first**: Developers need to understand "why" before "how"
- **Features second**: Quick scan of capabilities
- **Installation/Quickstart placeholders**: Signals future functionality
- **Contributing/License**: Standard OSS expectation

### Implementation
```markdown
# SourceBox

[![Build Status](https://github.com/jbeausoleil/sourcebox/workflows/CI/badge.svg)](https://github.com/jbeausoleil/sourcebox/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/dl/)

> Verticalized demo data for developers in 30 seconds

## Legal Notice

**SourceBox is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information.**

[... rest of README content ...]
```

### Alternatives Considered
- **Legal notice in footer**: Rejected because less visible (constitution requires prominence)
- **More badges** (coverage, downloads): Rejected as premature and distracting
- **Detailed installation now**: Rejected because no actual installation exists yet (F004-F010 implement this)

---

## Decision 4: Legal Independence Notice

### Decision
**Exact Wording**:
> SourceBox is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information.

**Placement**:
- README.md (immediately after tagline, before problem statement)
- In `## Legal Notice` section with heading for prominence

### Rationale

**Wording Justification**:
- **"Independently developed"**: Establishes autonomous creation
- **"Personal equipment"**: No employer hardware/cloud resources used
- **"Outside work hours"**: No time theft from employer
- **"No employer resources"**: No proprietary tools, cloud credits, internal libraries
- **"No proprietary information"**: Clean-room development, public info only

**Placement Justification**:
- **Above the fold**: Visible without scrolling
- **Dedicated section**: Gives appropriate weight to legal protection
- **Before features**: Establishes context before technical details
- Satisfies Constitution Section 6 (Legal Constraints), particularly:
  - Legal Constraint 1: Independent Development Only
  - Legal Constraint 2: No Employer References
  - Legal Constraint 4: Open Source Protection

### Implementation
Integrated into README template with clear visual separation.

### Alternatives Considered
- **Footer placement**: Rejected because less visible and easier to miss
- **Verbose legal language**: Rejected because intimidating to contributors
- **No notice**: Rejected because violates Constitution (NON-NEGOTIABLE legal protection)
- **LICENSE file only**: Rejected because README is more visible than LICENSE

---

## Decision 5: CONTRIBUTING.md Content

### Decision
Structure CONTRIBUTING.md with three primary sections:

1. **Bug Reports**: How to report bugs with required information
2. **Feature Requests**: How to propose new features
3. **Pull Requests**: How to contribute code with standards

### Rationale

**Bug Reports Section**:
- Reduces noise in issue tracker
- Ensures actionable bug reports with reproduction steps
- Template-style guidance (without actual GitHub issue templates yet)
- Aligns with Constitution Principle VI: Developer-First Design

**Feature Requests Section**:
- Encourages problem-focused proposals (not solution-focused)
- Helps prioritize community needs
- Sets expectation that features must align with constitution

**Pull Request Section**:
- Coding standards: gofmt, golint, go vet required
- Testing requirements: TDD for core logic (per Constitution Practice 1)
- Review process: maintainer approval required
- Commit message format: conventional commits style

### Implementation

```markdown
# Contributing to SourceBox

## Reporting Bugs

Before creating bug reports, please check existing issues. When creating a bug report, include:

- **Clear description**: What did you expect vs. what happened?
- **Steps to reproduce**: Exact commands/actions to trigger the bug
- **Environment**: Go version, OS, architecture
- **Logs/errors**: Relevant error messages or stack traces

## Requesting Features

Feature requests should focus on the problem, not the solution. Include:

- **Problem statement**: What pain point are you experiencing?
- **Current workaround**: How are you solving this today?
- **Expected benefit**: How would this feature improve your workflow?

## Submitting Pull Requests

1. **Fork and branch**: Create a feature branch from `main`
2. **Code standards**:
   - Run `gofmt -w .` before committing
   - Run `go vet ./...` and fix issues
   - Follow TDD for core functionality (write tests first)
3. **Commit messages**: Use conventional commits format
   - `feat:` for new features
   - `fix:` for bug fixes
   - `docs:` for documentation
   - `refactor:` for code improvements
4. **Testing**: Ensure all tests pass with `go test ./...`
5. **Pull request**: Target `main` branch, describe changes clearly

### Code Style

- Follow standard Go conventions (use `gofmt`)
- Write clear, self-documenting code
- Add comments for non-obvious logic
- Keep functions focused and small

### Testing Requirements

- **Core logic**: TDD required (write test first, then implementation)
- **Coverage**: Aim for >80% coverage on core packages
- **Manual QA**: Test on macOS, Linux, Windows for releases
```

### Alternatives Considered
- **GitHub issue templates**: Rejected for F003 scope (can add in future)
- **Detailed style guide**: Rejected because gofmt handles formatting
- **CLA requirement**: Rejected because MIT license doesn't require it (Constitution Legal Constraint 4)
- **Strict review process**: Rejected because solo project initially (can add when team grows)

---

## Decision 6: Code of Conduct Standard

### Decision
Use **Contributor Covenant v2.1** as the code of conduct.

### Rationale

**Industry Standard**:
- Most widely adopted code of conduct in open source (used by 100,000+ projects)
- Recognized by GitHub, GitLab, and Linux Foundation
- Clear, concise, and enforceable

**Content Coverage**:
- Expected behavior (inclusive, respectful, professional)
- Unacceptable behavior (harassment, discrimination, trolling)
- Enforcement responsibilities (maintainer role)
- Reporting process (email contact)
- Enforcement guidelines (warning → temporary ban → permanent ban)

**Alignment with Constitution**:
- Supports Principle VI: Developer-First Design (welcoming community)
- Supports Go-to-Market Philosophy 4: Community-Driven (healthy collaboration)
- Low maintenance burden (standard text, no custom enforcement needed initially)

### Implementation
Copy Contributor Covenant v2.1 text from [contributor-covenant.org](https://www.contributor-covenant.org/version/2/1/code_of_conduct/) with customizations:

- **Contact email**: To be determined (personal email or GitHub issues)
- **Project name**: SourceBox
- **Scope**: Applies to all project spaces (repo, issues, PRs, discussions)

### Alternatives Considered
- **Custom code of conduct**: Rejected because reinventing wheel, lower trust
- **No code of conduct**: Rejected because GitHub recommends it for healthy OSS projects
- **Django Code of Conduct**: Rejected because more complex than needed for solo project
- **Go Community Code of Conduct**: Rejected because Contributor Covenant is more widely recognized

---

## Constitution Compliance Verification

### Core Principles
- ✅ **Boring Tech Wins**: All decisions use standard, proven patterns (Go modules, github/gitignore, Contributor Covenant)
- ✅ **Open Source Forever**: MIT license chosen (Decision 3), prominently displayed
- ✅ **Developer-First Design**: Clear documentation structure, standard Go conventions
- ✅ **Ship Fast**: Simple, minimal decisions enable quick implementation

### Technical Constraints
- ✅ **Performance**: Repository setup completes in < 2 minutes (minimal files, no large binaries)
- ✅ **Platform Support**: Git + Go work on all platforms (macOS, Linux, Windows)
- ✅ **Cost**: $0 (GitHub free tier, no external services)
- ✅ **License**: MIT (Decision 3)

### Legal Constraints (CRITICAL)
- ✅ **Independent Development**: Legal notice prominently placed (Decision 4)
- ✅ **No Employer References**: All documentation avoids employer mentions
- ✅ **Public Information Only**: All patterns from public sources (github/gitignore, golang.org, contributor-covenant.org)
- ✅ **Open Source Protection**: MIT license in LICENSE file

### Anti-Patterns Avoided
- ✅ No shiny tech (avoided exotic tools, stuck with Git + Go)
- ✅ No over-engineering (minimal .gitignore, simple README structure)
- ✅ No complex processes (straightforward contribution guidelines)
- ✅ No proprietary information (all research from public sources)

---

## Summary of Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| .gitignore | github/gitignore Go template + IDE/env additions | Community standard, comprehensive |
| Go Module | github.com/jbeausoleil/sourcebox, Go 1.21+ | Standard path format, balanced version |
| README | Badges (build, license, Go) + 8 sections | Standard OSS structure, legal notice prominent |
| Legal Notice | Independence statement in README above fold | Protects legal status, visible |
| CONTRIBUTING.md | Bugs, features, PRs with clear guidance | Reduces noise, sets expectations |
| CODE_OF_CONDUCT | Contributor Covenant v2.1 | Industry standard, widely trusted |

---

## Next Steps

With research complete, proceed to Phase 1:
1. ✅ Skip data-model.md (no data entities in this feature)
2. ✅ Skip contracts/ (no APIs in this feature)
3. ⏭️ Generate quickstart.md (setup and verification guide)
4. ⏭️ Update agent context (.claude/context.md or similar)
5. ⏭️ Fill plan.md with technical context and constitution check

---

**Research Complete**: 2025-10-14
**All NEEDS CLARIFICATION Resolved**: Yes
**Ready for Phase 1**: Yes
