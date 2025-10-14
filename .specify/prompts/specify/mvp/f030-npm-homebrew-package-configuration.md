# Feature Specification Prompt: F030 - npm & Homebrew Package Configuration

## Feature Metadata
- **Feature ID**: F030
- **Name**: npm & Homebrew Package Configuration
- **Category**: CLI Tool
- **Phase**: Week 11
- **Priority**: P0
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F029

## Solution Overview
Prepare CLI for npm and Homebrew distribution. npm package.json downloads correct platform binary, Homebrew formula installs from GitHub releases, test installations work.

## Acceptance Criteria
- npm package.json for 'sourcebox' package
- npm install script downloads correct platform binary
- Homebrew formula: sourcebox.rb
- Formula installs from GitHub releases
- Test: `npm install -g sourcebox` works
- Test: `brew install sourcebox` works (macOS)

## Related Constitution: **Distribution Channels (Technical Constraint 2)**, **Developer-First (Principle VI)**
