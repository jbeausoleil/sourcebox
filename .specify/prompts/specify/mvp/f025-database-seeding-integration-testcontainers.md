# Feature Specification Prompt: F025 - Database Seeding Integration (Testcontainers)

## Feature Metadata
- **Feature ID**: F025
- **Name**: Database Seeding Integration (Testcontainers)
- **Category**: CLI Tool
- **Phase**: Week 10
- **Priority**: P0
- **Estimated Effort**: Medium (4 days)
- **Dependencies**: F023 (MySQL), F024 (Postgres)

## Constitutional Alignment
- ✅ **TDD Required**: Integration tests for database seeding
- ✅ **Quality Standards**: Comprehensive testing

## Solution Overview
Integration tests using Testcontainers for real MySQL/Postgres. Test: seed fintech to MySQL (verify 4,950 records), seed healthcare to Postgres (verify 5,000 records), verify FK integrity, query performance.

## Acceptance Criteria
- Testcontainers for MySQL and Postgres
- Integration tests for all 3 schemas
- Verify record counts match expectations
- Verify FK integrity after seeding
- Query performance: <100ms for simple SELECT

## Related Constitution: **TDD Required (Development Practice 1)**, **Manual QA (Development Practice 3)**
