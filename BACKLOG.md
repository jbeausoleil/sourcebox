# SourceBox Project Backlog

**Purpose**: Track cross-feature enhancements and technical debt not tied to specific features

**Last Updated**: 2025-10-15

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
