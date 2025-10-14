# Feature Specification Prompt: F031 - MySQL Base Dockerfile & Seeding Script

## Feature Metadata
- **Feature ID**: F031
- **Name**: MySQL Base Dockerfile & Seeding Script
- **Category**: Docker Images
- **Phase**: Week 11
- **Priority**: P0
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F023 (MySQL connector), F026 (SQL export)

## Constitutional Alignment
- ✅ **Speed**: <10s Docker startup (NON-NEGOTIABLE)
- ✅ **Local-First**: Self-contained, no external downloads

## Solution Overview
Create base MySQL Dockerfile with automatic seeding on startup. FROM mysql:8.0, seed script in /docker-entrypoint-initdb.d/, environment vars for config, image size <500MB, startup <10s.

## Acceptance Criteria
- Dockerfile: docker/mysql/Dockerfile
- Base: FROM mysql:8.0
- Seed script: /docker-entrypoint-initdb.d/seed.sql
- Environment: MYSQL_DATABASE=demo, MYSQL_ROOT_PASSWORD=password
- Image size: <500MB compressed
- Startup: <10s to ready state
- Test: docker run works, database accessible immediately

## Related Constitution: **Speed > Features (Principle II)**, **Performance Gates (Technical Constraint 1)**
