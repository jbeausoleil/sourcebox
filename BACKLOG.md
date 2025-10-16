# SourceBox Project Backlog

**Purpose**: Track cross-feature enhancements and technical debt not tied to specific features

**Last Updated**: 2025-10-15 (F008 post-MVP enhancements added)

---

## CLI Framework Enhancements

### From F006 (Cobra CLI Framework)

#### High Priority
_None currently identified_

#### Medium Priority
- **Graduated Verbosity Levels** (F006 deferred)
  - Support `-vv`, `-vvv` style verbosity
  - Requires: User feedback indicating need for more granular control
  - Estimate: 2-3 hours (flag parsing, output level logic)
  - Source: `docs/004-f006-cobra-cli/session-summary/t016-t019-summary.yaml:132`

- **Output Helper Utilities** (F006 → F021)
  - `VerbosePrintf()`, `QuietPrintf()` wrappers
  - Required by: F021 (Seed), F022 (List-Schemas)
  - Estimate: 1-2 hours
  - Source: `docs/004-f006-cobra-cli/session-summary/t016-t019-summary.yaml:133`

#### Low Priority
- **Environment Variable Overrides** (F006 deferred)
  - Support `SOURCEBOX_VERBOSE=1`, `SOURCEBOX_QUIET=1`
  - Use case: CI automation, scripts
  - Estimate: 1 hour
  - Source: `docs/004-f006-cobra-cli/session-summary/t016-t019-summary.yaml:136`

- **Color Output Support** (F006 → F021)
  - Colorized logs in verbose mode
  - Requires: Terminal capability detection
  - Estimate: 2-3 hours
  - Source: `docs/004-f006-cobra-cli/session-summary/t016-t019-summary.yaml:135`

---

## Schema Parser Enhancements

### From F008 (Schema Parser & Validator)

#### High Priority
_None currently identified_

#### Medium Priority
- **Refactor Test Fixtures to testdata/**
  - Move embedded test fixtures to testdata/ directory
  - Improves: Test maintainability and readability
  - Estimate: 1 hour
  - Source: `docs/006-f008-schema-parser/session-summary/t005-t008-summary.yaml`

- **Validate Foreign Key Type Compatibility**
  - Ensure FK column types match referenced PK column types
  - Prevents: Type mismatch errors at database creation
  - Estimate: 2 hours
  - Source: `docs/006-f008-schema-parser/session-summary/t045-t050-summary.yaml`

- **Validate FK References Primary Key Columns**
  - Verify foreign keys reference actual primary key columns
  - Prevents: Invalid FK constraints at database level
  - Estimate: 2 hours
  - Source: `docs/006-f008-schema-parser/session-summary/t045-t050-summary.yaml`

#### Low Priority
- **Enhance Error Messages with Line Numbers**
  - Use json.Decoder with Offset() to report line numbers
  - Improves: Developer experience when debugging malformed schemas
  - Estimate: 2 hours
  - Source: `docs/006-f008-schema-parser/session-summary/t075-t098-summary.yaml`

- **Add Enum Syntax Validation**
  - Validate enum value syntax (single quotes required: enum('a','b'))
  - Prevents: Invalid enum definitions like enum(val1,val2)
  - Estimate: 1 hour
  - Source: `docs/006-f008-schema-parser/session-summary/t075-t098-summary.yaml`

---

## CI/CD Enhancements

### From F008 (Schema Parser & Validator)
- **Benchmark CI Job** (F032 future feature)
  - Add performance regression detection to GitHub Actions
  - Prevents: Performance degradation going unnoticed
  - Estimate: 1 hour (as part of F032 CI/CD Enhancement)
  - Source: `docs/006-f008-schema-parser/session-summary/t075-t098-summary.yaml`

---

## Technical Debt

### Testing Infrastructure
- **Test State Pollution** (T054 decision pending)
  - Status: Identified, awaiting fix/accept decision
  - Location: `specs/004-f006-cobra-cli/tasks.md` (T054)

- **Coverage Methodology Documentation** (T055)
  - Status: Identified, pending documentation
  - Location: `specs/004-f006-cobra-cli/tasks.md` (T055)

---

## How to Use This Backlog

1. **Adding Items**: Include priority, estimate, source reference
2. **Scheduling**: Medium+ items with estimates should be scheduled within 2 releases
3. **Reviewing**: Review quarterly or when starting new feature planning
4. **Retiring**: Move completed items to feature changelogs
