# Feature Specification Prompt: F028 - Error Handling & User Guidance

## Feature Metadata
- **Feature ID**: F028
- **Name**: Error Handling & User Guidance
- **Category**: CLI Tool
- **Phase**: Week 10
- **Priority**: P1
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F021, F023, F024

## Solution Overview
Comprehensive error handling with actionable messages. Connection errors show cause and solution, schema not found lists available schemas, invalid flags show correct usage, database errors suggest fixes.

## Acceptance Criteria
- Connection errors: show connection string + error + fix
- Schema not found: list available schemas
- Invalid flags: show correct usage examples
- Database errors: suggest fixes
- All errors include help link

## Related Constitution: **Fail Gracefully (UX Principle 4)**, **Actionable Error Messages**
