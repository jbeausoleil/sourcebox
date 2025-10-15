# Implementation Plan: Schema JSON Format Specification (F007)

**Branch**: `005-f007-schema-json` | **Date**: 2025-10-15 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/005-f007-schema-json/spec.md`

**Note**: This is a **design-only feature** (specification and documentation, no code implementation). Code implementation happens in F008 (Schema Parser).

## Summary

Define a comprehensive JSON-based schema format for describing database schemas, data generators, relationships, and validation rules. The format must support verticalized data generation (industry-specific realism), multiple tables with foreign keys, custom generators with distribution parameters, and schema versioning. Deliverables are:
1. Specification document (`schemas/schema-spec.md`)
2. Example schema demonstrating all features (`schemas/example-schema.json`)
3. Updated agent context (CLAUDE.md)

This specification will be implemented by F008 (Schema Parser).

## Technical Context

**Language/Version**: N/A (JSON specification document, language-agnostic)
**Primary Dependencies**: None (design work, no code dependencies)
**Storage**: Schemas will be stored as JSON files, embedded in Go binary via embed.FS (F008 implementation)
**Testing**: Manual validation (create example schema, verify JSON syntax, verify completeness)
**Target Platform**: JSON format is cross-platform (macOS, Linux, Windows)
**Project Type**: Specification document (design work, not implementation)
**Performance Goals**:
  - JSON parsing: Fast in Go standard library (microseconds for typical schema)
  - Schema validation: Once at load time (startup overhead negligible)
  - Format must enable fast generator lookup (O(1) or O(log n))
**Constraints**:
  - Must be human-readable (developers will write and edit schemas)
  - Must be machine-parseable (Go code will load and validate)
  - Must support MySQL and PostgreSQL (Phase 1 databases)
  - Must support custom generators (verticalized data)
  - Must be extensible (new generators, new databases, new features)
  - Must include validation rules (what's valid, what's invalid)
  - Must follow semver for schema versioning
  - Example schema must be Tier 1 complexity (<30s generation, MVP focus)
**Scale/Scope**: Foundation for all schemas (F010 fintech, F016 healthcare, F018 retail), community contributions (50+ schemas by Phase 2)

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

### Core Principles Verification ✅

✅ **Verticalized > Generic**: Schema format MUST support industry-specific data generators and realistic distributions (custom generators for fintech, healthcare, retail)
✅ **Speed > Features**: N/A (design work) - However, format design enables fast parsing in F008 (JSON parsing is fast in Go)
✅ **Local-First, Cloud Optional**: N/A (specification document)
✅ **Boring Tech Wins**: JSON format is proven, standard, boring - maximum compatibility, extensive tooling, well-understood
✅ **Open Source Forever**: Schema format will be MIT licensed, enables community schema contributions
✅ **Developer-First Design**: Clear specification document enables developers to contribute schemas, format is human-readable and machine-parseable
✅ **Ship Fast, Validate Early**: Design-only feature (2 days), no implementation delay, spec validated by creating example schema

### Technical Constraints Verification ✅

✅ **Performance**: JSON parsing is fast in Go standard library, schema loading happens once at startup
✅ **Distribution**: Schemas will be embedded in binary (no runtime file reads), format supports this
✅ **Database Support**: Format must support both MySQL and PostgreSQL column types, relationships, constraints
✅ **Cost**: $0 (design work, no infrastructure)
✅ **Code Quality**: N/A (documentation feature) - However, specification must be clear, unambiguous, implementable
✅ **License**: Schema format is MIT, schemas themselves are MIT (community can contribute)
✅ **Platform Support**: JSON is cross-platform, format design is platform-agnostic

### Legal Constraints Verification ✅

✅ **Independent Development**: JSON schema design using public best practices only
✅ **No Employer References**: N/A (technical specification)
✅ **Public Information Only**: All patterns from public JSON schema documentation, Go JSON parsing docs, database documentation
✅ **Open Source Protection**: Format will be MIT licensed
✅ **Illustrative Examples Only**: Example schemas use generic fintech terminology (loans, credit scores) from public domain

### Complexity Justification

**No violations identified**. This feature is design work only, no code complexity.

## Project Structure

### Documentation (this feature)

```
specs/005-f007-schema-json/
├── spec.md              # Feature specification (user stories, requirements, success criteria)
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output: 10 schema design decisions ✅ COMPLETE
├── quickstart.md        # Phase 1 output: Spec validation guide ✅ COMPLETE
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)

NOTE: data-model.md and contracts/ are SKIPPED for this feature because:
- F007 IS the data model format definition (recursive to define format for defining formats)
- Schema format itself IS the contract for F008 parser implementation
- No REST/GraphQL APIs (no need for API contracts)
```

### Source Code (repository root)

```
sourcebox/
├── schemas/                   # Schema definitions (F007 deliverables)
│   ├── schema-spec.md        # Comprehensive schema format specification (TO BE CREATED)
│   └── example-schema.json   # Example demonstrating all features (TO BE CREATED)
├── cmd/sourcebox/             # CLI (F006 - already exists)
├── pkg/
│   ├── schema/               # Schema parser (F008 - future implementation)
│   ├── generators/           # Data generators (F009-F020 - future)
│   └── database/             # Database connectors (F023-F024 - future)
└── CLAUDE.md                 # Agent context (updated with schema format guidelines) ✅ COMPLETE
```

**Structure Decision**: F007 is documentation-only. Primary deliverables are `schemas/schema-spec.md` and `schemas/example-schema.json`. No code implementation in this feature.

## Phase 0: Research (COMPLETE ✅)

**Status**: All 10 research decisions documented in `research.md`

### Research Decisions Documented:

1. **JSON Schema Design Best Practices**: Top-level structure, flat design, dual versioning (schema_version + version)
2. **Data Generator Type Patterns**: Built-in (personal, company, date, numeric) vs Custom (per-schema)
3. **Distribution Specification Format**: uniform, normal, lognormal, weighted, ranges
4. **Foreign Key Relationship Syntax**: Dual representation (inline for code, explicit for docs)
5. **Column Type Mappings**: MySQL/PostgreSQL common subset (int, varchar, decimal, timestamp, json, enum)
6. **Schema Versioning Strategy**: Semantic versioning (major, minor, patch)
7. **Validation Rules Specification**: Schema-level, table-level, column-level, relationship-level, generation order
8. **Extensibility Patterns**: Generator extensibility, database extensibility, column extensibility, schema_version for format evolution
9. **Generator Parameter Schema**: JSON objects with typed parameters, distribution params documented
10. **Example Schema Complexity Level**: Tier 1 (3 tables, 4,950 records, <30s, fintech domain)

**Constitutional Compliance**: All research decisions align with core principles (Verticalized > Generic, Boring Tech Wins, Developer-First, Open Source Forever)

**Source**: All decisions based on public information (JSON Schema docs, SQL docs, Faker.js patterns, Synthea patterns, semver specification)

## Phase 1: Design & Contracts (COMPLETE ✅)

### 1. Data Model (SKIPPED - See rationale above)

**Rationale**: F007 IS the data model format definition. Generating a data-model.md would be recursive (defining the format for defining formats). Skip this deliverable.

### 2. API Contracts (SKIPPED - See rationale above)

**Rationale**: F007 has no REST/GraphQL APIs. The JSON schema format itself IS the contract for F008 parser. No separate contracts/ directory needed. The schema-spec.md document serves as the contract.

### 3. Quickstart Guide (COMPLETE ✅)

**Status**: Generated `quickstart.md` with specification validation steps

**Contents**:
- Manual validation steps (10 verification checks)
- Specification completeness checklist
- Example schema JSON syntax validation
- Feature demonstration verification
- Tier 1 complexity compliance check
- Versioning documentation verification
- Validation rules documentation verification
- Generator documentation verification
- Data type documentation verification
- Foreign key relationship documentation verification
- F008 implementability verification
- Common issues and fixes
- Design quality checklist (simplicity, extensibility, clarity, realism, implementability)

### 4. Agent Context Update (COMPLETE ✅)

**Status**: Updated `CLAUDE.md` with schema format guidelines

**Additions**:
- Schema JSON format overview
- Schema structure example
- Generator types (built-in and custom)
- Distribution types (uniform, normal, lognormal, weighted, ranges)
- Foreign key dual representation (inline + explicit)
- Schema versioning (semantic versioning)
- Validation rules (schema/table/column/relationship/generation order)
- Tier 1 complexity targets (<5,000 records, <30s)
- Creating new schemas workflow
- Supported data types (MySQL/PostgreSQL common subset)

## Constitution Re-verification (Post-Design)

### Core Principles Alignment ✅

✅ **Verticalized > Generic**: Custom generators (credit_score, loan_amount) enable fintech realism without "John Doe" data
✅ **Speed > Features**: JSON format is simple, enables fast parsing (<1ms), no feature bloat
✅ **Local-First**: JSON format is offline-compatible, no network dependencies
✅ **Boring Tech Wins**: JSON is proven, standard, widely understood (not YAML, not custom DSL)
✅ **Open Source Forever**: MIT license, community can contribute schemas
✅ **Developer-First**: Clear specification, human-readable format, copyable examples
✅ **Ship Fast, Validate Early**: Design work complete in 2 days, no implementation delay

### Technical Constraints Compliance ✅

✅ **Performance**: Tier 1 example schema <5,000 records (constitutional <30s target)
✅ **Distribution**: JSON format embeds in Go binary via embed.FS
✅ **Database Support**: Format supports MySQL and PostgreSQL common types
✅ **Cost**: $0 (design work, no infrastructure)
✅ **Code Quality**: Specification is clear, unambiguous, implementable by F008
✅ **License**: MIT (schema format and example schemas)
✅ **Platform Support**: JSON is cross-platform

### Anti-Patterns Avoided ✅

✅ **No Feature Bloat**: Format is as simple as possible (flat structure, standard JSON)
✅ **No Over-Engineering**: No custom DSL, no AI/ML, no complex plugin system
✅ **No Generic Data**: Example schema demonstrates fintech realism (credit scores, loan amounts)
✅ **No Shiny Tech**: JSON is boring, proven (not YAML, not TOML, not custom format)
✅ **No Premature Optimization**: Start with 3 schemas (fintech, healthcare, retail), expand via community

### Legal Protection ✅

✅ **Independent Development**: JSON schema design using public best practices only
✅ **No Employer References**: No employer-specific information
✅ **Public Information Only**: All patterns from public JSON Schema docs, SQL docs, Faker.js, Synthea
✅ **Open Source Protection**: MIT license
✅ **Illustrative Examples**: Generic fintech terminology (loans, credit scores) from public domain

### Schema Complexity Tier Compliance ✅

✅ **Tier 1 Target**: Example schema has 4,950 records (within <5,000 Tier 1 constitutional constraint)
✅ **Generation Time**: <30 seconds (constitutional requirement for Tier 1)
✅ **Tables**: 3 tables (borrowers, loans, payments) - simple relationships
✅ **Use Case**: MVP focus (fintech schema for Phase 1 validation)

**No constitutional violations identified**. Proceeding to task generation.

## Deliverables Summary

**Generated by /speckit.plan** (Phase 0 + Phase 1):
1. ✅ `specs/005-f007-schema-json/plan.md` - This file
2. ✅ `specs/005-f007-schema-json/research.md` - Phase 0 output (10 schema design decisions)
3. ✅ `specs/005-f007-schema-json/quickstart.md` - Phase 1 output (spec validation guide)
4. ⏭️ `specs/005-f007-schema-json/data-model.md` - SKIPPED (F007 defines the model format itself)
5. ⏭️ `specs/005-f007-schema-json/contracts/` - SKIPPED (schema format IS the contract)
6. ✅ Updated CLAUDE.md with schema format guidelines

**NOT generated by /speckit.plan** (created later during implementation):
- `specs/005-f007-schema-json/tasks.md` - Phase 2, separate command (/speckit.tasks)
- `schemas/schema-spec.md` - Created during implementation (comprehensive format specification)
- `schemas/example-schema.json` - Created during implementation (demonstrates all features)

## Next Steps

1. **Run `/speckit.tasks`** to generate `tasks.md` from this plan
2. **Execute tasks** to create specification document and example schema
3. **Validate** using quickstart.md checklist
4. **Verify** constitutional compliance (Tier 1 complexity, verticalized data, implementable)
5. **F008 implementation** can begin after F007 specification is complete

## Success Criteria

Planning phase is complete when:

✅ All 10 research decisions documented with clear rationale
✅ JSON schema structure follows best practices (simplicity, clarity, extensibility)
✅ Generator types and distribution formats are well-documented
✅ Foreign key relationship format is unambiguous (inline + explicit)
✅ Schema versioning strategy follows semver
✅ Validation rules are comprehensive and implementable
✅ Example schema design is Tier 1 complexity (<5,000 records, <30s target)
✅ Quickstart provides clear specification validation steps
✅ Constitution compliance verified (no violations)
✅ Agent context updated with schema format guidelines
✅ Ready to proceed to task generation (/speckit.tasks)
✅ Specification is implementable by F008 parser
✅ Format enables community schema contributions

---

**Planning Status**: COMPLETE ✅
**Constitutional Violations**: None
**Ready for Task Generation**: Yes
**Next Command**: `/speckit.tasks`
