# Feature Specification Prompt: F036 - Docker Compose Multi-Database Example

## Feature Metadata
- **Feature ID**: F036
- **Name**: Docker Compose Multi-Database Example
- **Category**: Docker Images
- **Phase**: Week 12
- **Priority**: P1
- **Estimated Effort**: Small (1 day)
- **Dependencies**: F035

## Solution Overview
Create Docker Compose file for running multiple databases simultaneously. Services: mysql-fintech, postgres-healthcare, mysql-retail with unique ports, environment variables documented.

## Acceptance Criteria
- File: docker-compose.yml in repo root
- Services: mysql-fintech, postgres-healthcare, mysql-retail
- Port mapping: unique ports for each
- Environment variables documented
- Test: docker-compose up -d starts all
- Test: all databases accessible simultaneously
- Example queries documented

## Related Constitution: **Developer-First (Principle VI)**, **Documentation (Practice 6)**
