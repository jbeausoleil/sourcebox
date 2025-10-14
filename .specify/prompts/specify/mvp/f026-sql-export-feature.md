# Feature Specification Prompt: F026 - SQL Export Feature

## Feature Metadata
- **Feature ID**: F026
- **Name**: SQL Export Feature
- **Category**: CLI Tool
- **Phase**: Week 10
- **Priority**: P1
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F021, F023, F024

## Solution Overview
Add --output flag to export SQL instead of direct insertion. Generated SQL includes CREATE TABLE + INSERT statements, formatted and readable, works for both MySQL and PostgreSQL syntax.

## Acceptance Criteria
- Flag: `--output=<file.sql>`
- Exports CREATE TABLE + INSERT statements
- Formatted and readable SQL
- MySQL and PostgreSQL syntax support
- Validates exported SQL by executing in test env

## Related Constitution: **Developer-First Design (Principle VI)**
