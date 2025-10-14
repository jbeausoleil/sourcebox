# Feature Specification Prompt: F029 - CLI Binary Build & Distribution Setup

## Feature Metadata
- **Feature ID**: F029
- **Name**: CLI Binary Build & Distribution Setup
- **Category**: CLI Tool
- **Phase**: Week 10
- **Priority**: P0
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F004 (Build system), F021

## Solution Overview
Configure multi-platform binary builds. Makefile targets for each platform, cross-compilation with GOOS/GOARCH, binary size <50MB, gzip compression, GitHub Actions builds on release, upload to GitHub Releases.

## Acceptance Criteria
- Makefile targets: build-mac, build-linux, build-windows
- Cross-compilation working
- Binary size <50MB per platform
- gzip compression
- GitHub Actions: build on release tag
- Upload to GitHub Releases

## Related Constitution: **Platform Support (Technical Constraint 7)**, **Binary size <50MB (Technical Constraint 1)**
