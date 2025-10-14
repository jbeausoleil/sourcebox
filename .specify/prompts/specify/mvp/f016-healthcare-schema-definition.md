# Feature Specification Prompt: F016 - Healthcare Schema Definition (JSON)

## Feature Metadata
- **Feature ID**: F016
- **Name**: Healthcare Schema Definition (JSON)
- **Category**: Data Generation
- **Phase**: Week 7
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (3 days)
- **Dependencies**: F007 (Schema spec), F008 (Parser)

## Constitutional Alignment
- ✅ **Verticalized > Generic**: Healthcare must look like healthcare
- ✅ **Realism**: Real ICD-10 codes, actual medications, proper insurance
- ✅ **Quality Bar**: 10x more realistic than Faker

## User Story
**US-MVP-010**: "As a healthcare application developer, I want realistic patient, visit, prescription, and claims data with proper ICD-10 codes and medication names."

## Problem Statement
Healthcare demos need realistic medical data: proper diagnoses (ICD-10), real medication names, insurance providers, visit types. Generic data looks unprofessional.

## Solution Overview
Create `schemas/healthcare-patients.json` with 4 tables: patients (500), visits (1500), prescriptions (2000), claims (1000). Use real ICD-10 codes, actual medication names, realistic visit distributions.

## Acceptance Criteria
- Schema file: `schemas/healthcare-patients.json`
- 4 tables with proper relationships
- Real ICD-10 codes (E11=diabetes, I10=hypertension)
- Real medications (Lipitor, Metformin, Lisinopril)
- Realistic visit types (routine=60%, emergency=20%, follow-up=20%)
- Total: 5,000 records

## Related Constitution Sections
- **Verticalized > Generic (Principle I, NON-NEGOTIABLE)**
- **Industry Terminology**: Real medical codes and terms
- **Quality Bar**: 10x more realistic than Faker
