# Feature Specification Prompt: F020 - Data Generation Integration Tests

## Feature Metadata
- **Feature ID**: F020
- **Name**: Data Generation Integration Tests
- **Category**: Data Generation
- **Phase**: Week 8
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F013-F019 (All data generation features)

## Constitutional Alignment
- ✅ **TDD Required**: Integration tests verify full workflow
- ✅ **Quality Standards**: Comprehensive testing
- ✅ **Data Quality**: Verify FK integrity, distributions

## User Story
**US-MVP-014**: "As a QA engineer, I need comprehensive integration tests that verify all three schemas generate realistic, valid data with proper relationships."

## Solution Overview
Create comprehensive integration test suite that generates full datasets for all 3 schemas, verifies record counts, FK integrity, data quality, distributions, and performance.

## Acceptance Criteria
- Integration test: generate full fintech dataset (4,950 records)
- Integration test: generate full healthcare dataset (5,000 records)
- Integration test: generate full retail dataset (8,500 records)
- Verify FK integrity for all schemas
- Verify record counts match targets
- Verify data quality (no nulls in required fields, valid formats)
- All tests pass: `go test ./pkg/generators/... -v`
- Performance: All schemas generate in < 5 seconds

## Deliverables
1. Integration test suite (`pkg/generators/integration_test.go`)
2. FK integrity validation
3. Data quality checks
4. Performance benchmarks
5. Test coverage report

## Related Constitution Sections
- **TDD Required (Development Practice 1)**
- **Code Quality Standards (Technical Constraint 5)**
- **Performance Gates (Technical Constraint 1)**
