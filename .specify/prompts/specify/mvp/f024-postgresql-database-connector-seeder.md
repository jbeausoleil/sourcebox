# Feature Specification Prompt: F024 - PostgreSQL Database Connector & Seeder

## Feature Metadata
- **Feature ID**: F024
- **Name**: PostgreSQL Database Connector & Seeder
- **Category**: CLI Tool
- **Phase**: Week 10
- **Priority**: P0
- **Estimated Effort**: Large (5 days)
- **Dependencies**: F009 (Postgres driver), F021 (Seed command)

## Constitutional Alignment
- ✅ **Performance**: <30s seeding (NON-NEGOTIABLE)
- ✅ **TDD Required**: Core database logic

## Solution Overview
Implement Postgres connector in `pkg/database/postgres.go` with PostgreSQL-specific DDL, COPY protocol for performance, transaction support, connection pooling. Unit tests with mocks.

## Acceptance Criteria
- Postgres connection with error handling
- PostgreSQL-specific DDL
- COPY protocol for fast inserts
- Transaction support
- Connection pooling
- Unit tests with mocks
- Performance: <30s for 1,000 records

## Related Constitution: **Performance Gates (Technical Constraint 1)**, **Boring Tech (Principle IV)**
