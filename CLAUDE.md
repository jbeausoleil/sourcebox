# SourceBox Development Guidelines

Auto-generated from all feature plans. Last updated: 2025-10-14

## Project Overview
- **Repository**: github.com/jbeausoleil/sourcebox
- **Language**: Go 1.21+
- **License**: MIT (open source forever)
- **Development Status**: Foundational setup (Week 3 - F003)

## Active Technologies
- Go 1.21+ (CLI development)
- Git (version control)
- GitHub (repository hosting, CI/CD)

## Project Structure
```
sourcebox/
├── .gitignore          # Go-specific patterns
├── LICENSE             # MIT license
├── README.md           # Project overview with legal notice
├── CONTRIBUTING.md     # Contribution guidelines
├── CODE_OF_CONDUCT.md  # Contributor Covenant v2.1
├── go.mod              # Go module configuration
└── go.sum              # (generated after dependencies added)
```

## Go Module
- **Module Path**: github.com/jbeausoleil/sourcebox
- **Go Version**: 1.21+ (minimum)
- **Dependencies**: None yet (to be added in F009)

## Code Style
- Follow standard Go conventions (gofmt)
- Run `go vet ./...` before committing
- Run `gofmt -w .` before committing
- Write clear, self-documenting code
- Keep functions focused and small

## Testing Requirements
- **Core logic**: TDD required (write test first)
- **Coverage**: Aim for >80% on core packages
- **Manual QA**: Test on macOS, Linux, Windows for releases

## Legal Notice
**CRITICAL**: This project is developed independently on personal equipment, outside of work hours, with no use of employer resources or proprietary information. This notice MUST appear prominently in README.md.

## Recent Changes
- 001-f003-initialize-git: Foundational Git repository and Go module setup

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->