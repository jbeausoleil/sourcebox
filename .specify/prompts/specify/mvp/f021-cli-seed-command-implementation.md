# Feature Specification Prompt: F021 - CLI 'seed' Command Implementation

## Feature Metadata
- **Feature ID**: F021
- **Name**: CLI 'seed' Command Implementation
- **Category**: CLI Tool
- **Phase**: Week 9
- **Priority**: P0
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F006 (Cobra CLI)

## Constitutional Alignment
- ✅ **Developer-First**: Zero config, works out of the box
- ✅ **Speed > Features**: <30s seeding time
- ✅ **Clear Errors**: Actionable error messages

## Solution Overview
Implement main `seed` command with flags: --schema (required), --records, --host, --port, --user, --password, --db-name, --output (SQL export), --dry-run, --quiet. Comprehensive help text, flag validation, clear errors.

## Acceptance Criteria
- Command: `sourcebox seed <database> --schema=<name> --records=<count>`
- All flags working with validation
- Help text comprehensive
- Unit tests for flag parsing
- Clear error messages for missing/invalid flags

## Related Constitution: **Zero Config (UX Principle 2)**, **Fail Gracefully (UX Principle 4)**
