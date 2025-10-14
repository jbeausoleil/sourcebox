# Feature Specification Prompt: F034 - Docker Image Testing & Validation

## Feature Metadata
- **Feature ID**: F034
- **Name**: Docker Image Testing & Validation
- **Category**: Docker Images
- **Phase**: Week 11
- **Priority**: P0
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F033

## Solution Overview
Test all Docker images for functionality and data integrity. Each image starts successfully (<10s), database accessible, correct record counts, FK constraints work, sample queries return expected results. Automated test script.

## Acceptance Criteria
- Test: each image starts in <10s
- Test: database accessible on correct port
- Test: correct record counts in each table
- Test: FK constraints work
- Test: sample queries return results
- Automated test script: docker/test-all.sh

## Related Constitution: **Manual QA (Development Practice 3)**, **Quality Standards (Technical Constraint 5)**
