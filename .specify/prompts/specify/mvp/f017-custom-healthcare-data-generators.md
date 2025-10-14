# Feature Specification Prompt: F017 - Custom Healthcare Data Generators

## Feature Metadata
- **Feature ID**: F017
- **Name**: Custom Healthcare Data Generators
- **Category**: Data Generation
- **Phase**: Week 7
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (4 days)
- **Dependencies**: F011 (Base generators), F016 (Healthcare schema)

## Constitutional Alignment
- ✅ **Verticalized > Generic**: Medical data must use real codes and terms
- ✅ **Realism**: Real-world medical distributions
- ✅ **TDD Required**: Core generators need tests

## User Story
**US-MVP-011**: "As a data generation engine, I need generators for ICD-10 codes, medications, insurance providers, and visit types with realistic distributions."

## Solution Overview
Implement `pkg/generators/healthcare.go` with: Diagnosis() (real ICD-10 codes), Medication() (real drug names), VisitType() (weighted distribution), InsuranceProvider() (real insurers), realistic age distributions.

## Acceptance Criteria
- Diagnosis generator with real ICD-10 codes (E11, I10, J44, etc.)
- Medication generator with real names (Lipitor, Metformin, etc.)
- VisitType: weighted (routine=60%, emergency=20%, follow-up=20%)
- InsuranceProvider: real names (Blue Cross, Aetna, UnitedHealthcare)
- Age distribution: 0-100, weighted toward 40-65
- Unit tests verify distributions

## Related Constitution Sections
- **Verticalized > Generic (Principle I)**
- **Industry Terminology**: Real medical codes
- **TDD Required (Development Practice 1)**
