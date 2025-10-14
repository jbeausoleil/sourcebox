# Implementation Planning Prompt: F003 - Initialize Git Repository & Go Module

## Feature Metadata
- **Feature ID**: F003
- **Name**: Initialize Git Repository & Go Module
- **Feature Branch**: `001-f003-initialize-git`
- **Category**: Foundation
- **Phase**: Week 3
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (1 day)
- **Dependencies**: None (foundational feature)
- **Spec Location**: `/specs/001-f003-initialize-git/spec.md`

## Constitutional Alignment

### Core Principles Verification
- ✅ **Verticalized > Generic**: N/A (infrastructure setup)
- ✅ **Speed > Features**: Repository setup completes in < 2 minutes
- ✅ **Local-First, Cloud Optional**: Git works 100% offline after clone
- ✅ **Boring Tech Wins**: Git + Go module (standard, proven, boring)
- ✅ **Open Source Forever**: MIT license, public repository (non-negotiable)
- ✅ **Developer-First Design**: Standard Go project conventions, clear documentation
- ✅ **Ship Fast, Validate Early**: Foundational step for 12-week MVP

### Technical Constraints Verification
- ✅ **Performance**: < 2 minute clone + setup time (SC-001 from spec)
- ✅ **Distribution**: GitHub (standard for open source)
- ✅ **Database Support**: N/A (no database in this feature)
- ✅ **Cost**: $0 (GitHub free tier)
- ✅ **Code Quality**: Manual verification sufficient for configuration
- ✅ **License**: MIT (non-negotiable)
- ✅ **Platform Support**: Cross-platform (Git + Go work everywhere)

### Legal Constraints Verification (CRITICAL)
- ✅ **Independent Development**: README includes legal notice
- ✅ **No Employer References**: Documentation avoids employer mentions
- ✅ **Public Information Only**: All patterns from public Go/Git standards
- ✅ **Open Source Protection**: MIT license in LICENSE file
- ✅ **Illustrative Examples Only**: No false claims or affiliations

## Planning Context

### Feature Summary
Initialize a public GitHub repository with proper Go module setup, MIT license, contribution guidelines, code of conduct, and legal independence notice. This is the foundational step that enables all subsequent development.

### Key Technical Decisions Required

**Phase 0 Research Topics**:
1. **Go .gitignore Patterns**: What are standard Go community patterns?
2. **Go Module Configuration**: What minimum Go version? (Spec says 1.21+)
3. **README Structure**: What badges, sections, and content?
4. **Legal Notice Placement**: Where and how to display independence notice?
5. **CONTRIBUTING.md Content**: What instructions for bugs, features, PRs?
6. **CODE_OF_CONDUCT Standard**: Use Contributor Covenant v2.1?

### Technical Context (Pre-filled)

**Language/Version**: Go 1.21+ (specified in spec)
**Primary Dependencies**: None (foundational feature, no code dependencies yet)
**Storage**: Git repository (local .git directory)
**Testing**: Manual verification checklist (no executable logic to test)
**Target Platform**: Cross-platform (macOS, Linux, Windows via Git + Go)
**Project Type**: Single project (CLI tool)
**Performance Goals**: Clone + setup < 2 minutes (SC-001)
**Constraints**:
  - Repository size < 1MB (SC-002)
  - Must work offline after initial clone (Principle III)
  - No proprietary information (Legal Constraint 3)
  - No employer references (Legal Constraint 2)
**Scale/Scope**: Foundational infrastructure (no scale considerations yet)

## Planning Workflow

### Phase 0: Research & Technical Decisions

Generate `research.md` with documented decisions for:

#### 1. Git Configuration
- **Decision Point**: What should be in .gitignore?
- **Research**: Standard Go .gitignore patterns from github.com/github/gitignore
- **Output**: Documented pattern list with rationale

#### 2. Go Module Setup
- **Decision Point**: Module path format, minimum Go version
- **Research**: Go module best practices from golang.org
- **Output**: Module path (github.com/jbeausoleil/sourcebox) and version (1.21+) with justification

#### 3. Documentation Structure
- **Decision Point**: What sections in README? What badges?
- **Research**: Common patterns from popular Go CLI projects
- **Output**: README outline with mandatory sections and badges

#### 4. Legal Independence Notice
- **Decision Point**: Exact wording and placement of legal notice
- **Research**: Constitution Section 6 (Legal Constraints)
- **Output**: Final notice text and location (README, above the fold)

#### 5. Contribution Guidelines
- **Decision Point**: What process for bugs, features, PRs?
- **Research**: Standard OSS contribution patterns
- **Output**: CONTRIBUTING.md structure with clear instructions

#### 6. Code of Conduct
- **Decision Point**: Which standard to use?
- **Research**: Contributor Covenant v2.1 (current standard)
- **Output**: Rationale for using Contributor Covenant

**Deliverable**: `specs/001-f003-initialize-git/research.md`

### Phase 1: Design & Contracts

#### 1. Data Model (SKIP for this feature)
**Rationale**: F003 is repository infrastructure setup. No data entities, no models. Skip data-model.md generation.

#### 2. API Contracts (SKIP for this feature)
**Rationale**: F003 has no APIs, no endpoints, no services. Skip contracts/ directory generation.

#### 3. Quickstart Guide (REQUIRED)
Generate `quickstart.md` with:

```markdown
# F003 Quickstart: Repository Setup & Verification

## Prerequisites
- Git 2.0+ installed
- Go 1.21+ installed
- GitHub account (for forking/contributing)

## Setup Steps
1. Clone repository
2. Verify Go module
3. Review documentation
4. Understand contribution process

## Verification Checklist
- [ ] `git clone` completes successfully
- [ ] `go mod download` runs without errors
- [ ] README renders correctly on GitHub
- [ ] Legal notice is visible in README
- [ ] All documentation files present (LICENSE, CONTRIBUTING, CODE_OF_CONDUCT)
- [ ] Repository size < 1MB

## Next Steps
- Review F004 (Project Directory Structure)
- Understand Spec-Kit workflow
```

**Deliverable**: `specs/001-f003-initialize-git/quickstart.md`

#### 4. Update Agent Context
Run: `.specify/scripts/bash/update-agent-context.sh claude`

This updates the Claude-specific context file with:
- Git repository URL
- Go module path
- Legal independence notice location

**Deliverable**: Updated `.claude/context.md` (or similar agent-specific file)

## Constitution Re-verification

After Phase 1 design, verify:
- [ ] Legal notice is documented in quickstart
- [ ] No employer references in any generated files
- [ ] MIT license approach is documented in research.md
- [ ] Speed constraint (< 2 min setup) is verifiable via quickstart checklist
- [ ] Open source principle (MIT, public repo) is maintained

## Deliverables Summary

**Generated by /speckit.plan**:
1. ✅ `specs/001-f003-initialize-git/plan.md` - This file
2. ✅ `specs/001-f003-initialize-git/research.md` - Phase 0 output
3. ✅ `specs/001-f003-initialize-git/quickstart.md` - Phase 1 output
4. ⏭️ `specs/001-f003-initialize-git/data-model.md` - SKIP (N/A)
5. ⏭️ `specs/001-f003-initialize-git/contracts/` - SKIP (N/A)
6. ✅ Updated agent context file

**NOT generated by /speckit.plan** (created later by /speckit.tasks):
- `specs/001-f003-initialize-git/tasks.md` - Phase 2, separate command

## Success Criteria for Planning Phase

- ✅ Research decisions are documented with clear rationale
- ✅ Quickstart provides clear verification steps
- ✅ Constitution compliance is verified
- ✅ No complexity violations (this is intentionally simple infrastructure)
- ✅ Agent context is updated with relevant technical details
- ✅ Planning artifacts reference the constitution and spec correctly

## Anti-Patterns to Avoid

- ❌ Don't generate data-model.md (no entities in this feature)
- ❌ Don't generate API contracts (no APIs in this feature)
- ❌ Don't over-engineer the .gitignore (use standard Go patterns)
- ❌ Don't add unnecessary badges to README (only: build, license, Go version)
- ❌ Don't include employer references in any documentation
- ❌ Don't make complex technical decisions (this should be boring and standard)

## Implementation Notes

### For the AI Agent
When executing `/speckit.plan` with this prompt:

1. **Start with research.md**: Document all 6 research decisions listed in Phase 0
2. **Be explicit about skips**: Clearly state why data-model.md and contracts/ are not needed
3. **Focus on quickstart.md**: This is the primary deliverable beyond research
4. **Verify legal compliance**: Double-check no employer references, legal notice is documented
5. **Keep it simple**: This is intentionally boring infrastructure - no complexity needed

### Related Constitution Sections
- **Core Principle V**: Open Source Forever (MIT license non-negotiable)
- **Core Principle IV**: Boring Tech Wins (Git + Go, standard patterns)
- **Legal Constraints (Section 6)**: Independent development, no employer resources (CRITICAL)
- **Development Practice 6**: Spec-Kit workflow (this planning phase)
- **Anti-Pattern 4**: Shiny Tech (reject non-standard solutions)

## Drag-and-Drop Usage

**To use this prompt**:
1. Drag this file into Claude Code
2. Claude will execute the `/speckit.plan` workflow for F003
3. Expected outputs:
   - research.md with 6 documented decisions
   - quickstart.md with setup verification steps
   - Updated agent context
   - Constitution compliance verified

**Estimated time**: 5-10 minutes for complete planning phase
