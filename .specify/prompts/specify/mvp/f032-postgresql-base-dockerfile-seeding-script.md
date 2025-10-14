# Feature Specification Prompt: F032 - PostgreSQL Base Dockerfile & Seeding Script

## Feature Metadata
- **Feature ID**: F032
- **Name**: PostgreSQL Base Dockerfile & Seeding Script
- **Category**: Docker Images
- **Phase**: Week 11
- **Priority**: P0
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F024 (Postgres connector), F026 (SQL export)

## Constitutional Alignment
- ✅ **Speed**: <10s Docker startup
- ✅ **Local-First**: Self-contained image

## Solution Overview
Create base PostgreSQL Dockerfile. FROM postgres:16, seed script in /docker-entrypoint-initdb.d/, environment vars, image size <500MB, startup <10s.

## Acceptance Criteria
- Dockerfile: docker/postgres/Dockerfile
- Base: FROM postgres:16
- Seed script: /docker-entrypoint-initdb.d/seed.sql
- Environment: POSTGRES_DB=demo, POSTGRES_PASSWORD=password
- Image size: <500MB compressed
- Startup: <10s to ready
- Test: docker run works

## Related Constitution: **Performance Gates (Technical Constraint 1)**
