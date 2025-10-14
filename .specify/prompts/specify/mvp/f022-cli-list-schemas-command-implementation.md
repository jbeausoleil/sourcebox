# Feature Specification Prompt: F022 - CLI 'list-schemas' Command Implementation

## Feature Metadata
- **Feature ID**: F022
- **Name**: CLI 'list-schemas' Command Implementation
- **Category**: CLI Tool
- **Phase**: Week 9
- **Priority**: P1
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F006 (Cobra CLI), F008 (Schema parser)

## Constitutional Alignment
- ✅ **Developer-First**: Easy schema discovery
- ✅ **Clear Output**: Table format with descriptions

## Solution Overview
Implement `list-schemas` command (alias: `ls`) that displays all available schemas in table format: name, description, tables, record count. Reads from embedded schemas directory, colorized output.

## Acceptance Criteria
- Command: `sourcebox list-schemas` or `sourcebox ls`
- Table output: schema name | description | tables | records
- Reads embedded schemas
- Colorized for readability

## Related Constitution: **Developer-First Design (Principle VI)**
