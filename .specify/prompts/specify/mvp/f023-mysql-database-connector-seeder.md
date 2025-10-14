# Feature Specification Prompt: F023 - MySQL Database Connector & Seeder

## Feature Metadata
- **Feature ID**: F023
- **Name**: MySQL Database Connector & Seeder
- **Category**: CLI Tool
- **Phase**: Week 9-10
- **Priority**: P0
- **Estimated Effort**: Large (5 days)
- **Dependencies**: F009 (MySQL driver), F021 (Seed command)

## Constitutional Alignment
- ✅ **Performance**: <30s seeding for 1,000 records (NON-NEGOTIABLE)
- ✅ **TDD Required**: Core database logic

## Solution Overview
Implement MySQL connector in `pkg/database/mysql.go` with connection handling, DDL execution (CREATE TABLE), batched INSERT statements (100 records/batch), transaction support, connection pooling. Comprehensive error handling, unit tests with mocks.

## Acceptance Criteria
- MySQL connection with error handling
- DDL execution from schema
- Batched INSERTs (100/batch)
- Transaction support for rollback
- Connection pooling
- Unit tests with mock database
- Performance: Seed 1,000 records in <30s

## Related Constitution: **Performance Gates (Technical Constraint 1)**, **TDD Required (Development Practice 1)**
