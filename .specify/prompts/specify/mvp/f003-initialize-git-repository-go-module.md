# Feature Specification Prompt: F003 - Initialize Git Repository & Go Module

## Feature Metadata
- **Feature ID**: F003
- **Name**: Initialize Git Repository & Go Module
- **Category**: Foundation
- **Phase**: Week 3
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (1 day)
- **Dependencies**: None (foundational feature)

## Constitutional Alignment

### Core Principles
- ✅ **Open Source Forever**: Public MIT-licensed repository
- ✅ **Boring Tech Wins**: Go for CLI (single binary, cross-platform, standard)
- ✅ **Ship Fast, Validate Early**: Foundation for 12-week MVP
- ✅ **Developer-First Design**: Standard Go project structure
- ✅ **Legal Protection**: Independent development notice included

### Technical Constraints
- ✅ **Platform Support**: Git repository accessible on all platforms
- ✅ **Open Source License**: MIT license (non-negotiable)
- ✅ **Cost Constraints**: Free GitHub hosting

### Development Practices
- ✅ **Indie Project Constraints**: Personal equipment, outside work hours
- ✅ **Legal Independence**: No employer resources, public information only

## User Story
**US-MVP-002**: "As a developer, I want to clone the SourceBox repository and have a clear understanding of the project structure, licensing, and how to contribute."

## Problem Statement
SourceBox needs a foundational Git repository with proper Go module initialization, clear licensing (MIT), contribution guidelines, and code of conduct. This is the first technical step that enables all subsequent development work. The repository must be public, well-documented, and include a legal notice about independent development to protect the project's legal status.

## Solution Overview
Create a public GitHub repository with MIT license, initialize as a Go module, add essential documentation files (README, CONTRIBUTING, CODE_OF_CONDUCT), configure .gitignore for Go projects, and include a legal notice stating the project is "Developed independently on personal equipment."

## Detailed Requirements

### Acceptance Criteria
1. **Git Repository Created**: Repository exists at `github.com/yourusername/sourcebox`
2. **Repository is Public**: No authentication required to clone or view
3. **MIT License Added**: LICENSE file with MIT license text
4. **README.md Created**: Project overview with badges (build status, license, version)
5. **CONTRIBUTING.md Created**: Contribution guidelines including how to:
   - Report bugs
   - Submit feature requests
   - Create pull requests
   - Follow coding standards
6. **CODE_OF_CONDUCT.md Created**: Standard Contributor Covenant code of conduct
7. **Go Module Initialized**: `go mod init github.com/yourusername/sourcebox` executed successfully
8. **.gitignore Configured**: Properly configured for Go projects, excluding:
   - Binary files
   - Build artifacts
   - IDE files (.vscode/, .idea/, *.swp)
   - OS files (.DS_Store)
   - Test coverage files
9. **Legal Notice Added**: README includes: "SourceBox is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information."

### Technical Specifications

#### Repository Structure (Initial)
```
sourcebox/
├── .gitignore
├── LICENSE
├── README.md
├── CONTRIBUTING.md
├── CODE_OF_CONDUCT.md
├── go.mod
└── go.sum (generated after first dependency)
```

#### README.md Contents
Must include:
- Project name and tagline
- Badges: Build Status, License, Go Version
- Quick description (1-2 sentences)
- Problem statement (why SourceBox exists)
- Features (brief bulleted list)
- Installation instructions (placeholder for future)
- Quick start example (placeholder for future)
- Contributing link
- License badge/link
- Legal independence notice

#### Go Module Configuration
- Module path: `github.com/yourusername/sourcebox`
- Go version: 1.21 or higher
- No dependencies yet (will be added in F009)

#### .gitignore Patterns
Must exclude:
```
# Binaries
sourcebox
*.exe
*.dll
*.so
*.dylib

# Test & coverage
*.test
*.out
coverage.txt
coverage.html

# Build artifacts
/dist/
/build/

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Environment
.env
.env.local
```

### Performance Considerations
- Repository clone should complete quickly (< 30 seconds on typical connection)
- No large binary files checked into Git (use .gitignore)
- Keep repository size minimal in early stages

### Security & Privacy
- No secrets or credentials in repository (add .env to .gitignore)
- No proprietary information or employer references
- Legal independence notice protects project status

### Testing Strategy
**Manual Verification (no automated tests needed for this feature)**:
1. Clone repository in clean directory
2. Verify `go mod download` works (even with no dependencies)
3. Verify README renders correctly on GitHub
4. Verify all links in documentation work
5. Verify legal notice is visible and clear

## Dependencies
- **Upstream**: None (this is the first technical feature)
- **Downstream**:
  - F004 (Project Directory Structure) depends on this
  - F009 (Dependency Management) depends on go.mod
  - All subsequent features depend on repository existing

## Deliverables
1. Public GitHub repository at specified URL
2. MIT LICENSE file
3. README.md with project overview and legal notice
4. CONTRIBUTING.md with contribution guidelines
5. CODE_OF_CONDUCT.md with community standards
6. Initialized Go module (go.mod file)
7. Properly configured .gitignore for Go projects

## Success Criteria
- ✅ Repository is publicly accessible
- ✅ Go module can be referenced in other projects
- ✅ Documentation clearly explains project purpose
- ✅ Legal independence is documented
- ✅ Contribution process is clear
- ✅ MIT license is properly applied

## Anti-Patterns to Avoid
- ❌ Private repository (violates open source principle)
- ❌ Proprietary license or no license
- ❌ Employer references or proprietary information
- ❌ Missing documentation files
- ❌ Checking in binary files or build artifacts
- ❌ Complex initial structure (keep it simple)

## Implementation Notes
- Use standard Go community practices (follow golang.org/doc conventions)
- Keep initial structure minimal - directory organization comes in F004
- Focus on legal clarity - independence notice is critical
- Badge URLs will need updating after CI/CD setup in F005

## TDD Requirements
**Not applicable for this feature** - This is project setup/configuration work, not executable logic. Manual verification is sufficient.

## Related Constitution Sections
- **Open Source Forever (Principle V)**: MIT license non-negotiable
- **Legal Constraints (Section 6)**: Independent development, no employer resources
- **Boring Tech Wins (Principle IV)**: Go module for CLI
- **Developer-First Design (Principle VI)**: Clear documentation and contribution guidelines
