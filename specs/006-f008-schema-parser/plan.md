# Implementation Plan: F008 - Schema Parser & Validator

**Branch**: `006-f008-schema-parser` | **Date**: 2025-10-15 | **Spec**: [`spec.md`](./spec.md)
**Input**: Feature specification from `/specs/006-f008-schema-parser/spec.md`

## Summary

Build a robust schema parser in `/pkg/schema/parser.go` that loads JSON schema files (defined by F007), unmarshals into Go structs, validates all constraints (required fields, valid types, foreign key integrity, generation order), and returns clear, actionable error messages. Use Go's standard `encoding/json` for parsing. Implement comprehensive validation with TDD discipline targeting 100% test coverage. Foundation for all data generation features (F010 fintech, F016 healthcare, F018 retail).

**Technical Approach** (from research):
- Go structs with JSON tags mirror F007 format exactly
- Strict parsing with `DisallowUnknownFields()` to catch typos early
- Multi-phase fail-fast validation (schema → table → column → foreign keys → generation order)
- Structured error messages with context (table name, column name, array index)
- Linear-time validation (O(n) complexity) using map lookups for <100ms target
- TDD workflow: Write test (RED) → Implement (GREEN) → Refactor → Repeat

---

## Technical Context

**Language/Version**: Go 1.21+ (existing project configuration from F003/F004)

**Primary Dependencies**:
- `encoding/json` - Go standard library for JSON parsing
- `github.com/stretchr/testify` - Testing assertions (already in project from F006)

**Storage**: Schemas loaded from files (development) or `embed.FS` (production binary)

**Testing**: TDD required with 100% coverage target for `pkg/schema/parser.go` and `types.go`

**Target Platform**: Cross-platform (macOS Intel/ARM, Linux x86_64/ARM64, Windows x86_64)

**Project Type**: Single project (CLI tool library - `pkg/schema` package)

**Performance Goals**:
- Schema parsing: <100ms for typical schemas (<10 tables, <50 columns)
- JSON unmarshaling: O(n) where n = schema size
- Validation: O(n) where n = number of tables + columns
- Target measured on 2020 MacBook Pro (constitutional hardware baseline)

**Constraints**:
- Must validate all F007 schema rules (required fields, valid types, foreign key integrity)
- Must provide clear, actionable error messages (UX Principle 4: Fail Gracefully)
- Must use Go standard library `encoding/json` (no external JSON libraries)
- Must support both file and embedded schema loading
- TDD required (100% coverage for core parser logic)
- Must be cross-platform compatible (no platform-specific file operations)
- Error messages must indicate exactly what's wrong and where

**Scale/Scope**: Foundation for all schema-based features (F010 fintech, F016 healthcare, F018 retail, F013 data generation engine)

---

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

### Core Principles Verification

✅ **I. Verticalized > Generic**: N/A (foundation infrastructure) - Parser validates verticalized schemas
✅ **II. Speed > Features**: Fast schema loading (<100ms) is critical for CLI startup time
✅ **III. Local-First, Cloud Optional**: Pure local operation, no network dependencies
✅ **IV. Boring Tech Wins**: Go standard library (`encoding/json`), proven patterns, no exotic dependencies
✅ **V. Open Source Forever**: Parser is core CLI functionality, MIT licensed
✅ **VI. Developer-First Design**: Clear error messages enable self-service debugging
✅ **VII. Ship Fast, Validate Early**: TDD required, 3-day implementation, foundational for all data generation

### Technical Constraints Verification

✅ **1. Performance**: <100ms schema loading (constitutional requirement), <30s total data generation
✅ **2. Distribution Channels**: Parser embedded in binary, schemas loaded from `embed.FS`
✅ **3. Database Support**: Validates schemas for both MySQL and PostgreSQL
✅ **4. Cost Constraints**: $0 (standard library, no external dependencies)
✅ **5. Code Quality Standards**: TDD required, 100% coverage target for core parser logic
✅ **6. Open Source License**: MIT, uses only standard library (no license conflicts)
✅ **7. Platform Support**: Cross-platform (Go stdlib `encoding/json` works everywhere)

### Development Practices Verification

✅ **1. TDD Required for Core Functionality**: Parser is core logic, TDD non-negotiable
✅ **2. Test-After OK for Non-Critical Paths**: N/A (all parser logic is critical)
✅ **3. Manual QA Before Every Release**: Manual testing on Mac/Linux/Windows before parser release
✅ **4. Ship CLI + Docker MVP in 12 Weeks**: Parser is foundational (Week 4), on track
✅ **5. Open Source Launch First**: Parser part of open source CLI
✅ **6. Spec-Kit Driven Workflow**: Using `/speckit.plan` workflow (this document)
✅ **7. Indie Project Constraints**: 3-day implementation fits 10-15 hours/week constraint

### Legal Constraints Verification (CRITICAL)

✅ **1. Independent Development Only**: Go standard library patterns, public JSON parsing documentation
✅ **2. No Employer References**: N/A (technical implementation)
✅ **3. Public Information Only**: All patterns from public Go documentation, JSON parsing best practices
✅ **4. Open Source Protection**: MIT license, Go standard library (BSD license, compatible)
✅ **5. Illustrative Examples Only**: Parser validates generic schema format (fintech, healthcare, retail)

### Anti-Patterns Verification

✅ **1. Feature Bloat**: Rejected custom JSON parser, complex graph algorithms (MVP focus)
✅ **2. Enterprise-First**: Simple parser for MVP, no enterprise features
✅ **3. Complex Pricing in Phase 1**: Parser is free forever (CLI tool)
✅ **4. Shiny Tech**: Using Go stdlib, not Rust or exotic libraries
✅ **5. Over-Engineering**: Simple structs, standard unmarshaling, no interfaces
✅ **6. Generic Data**: Parser enables verticalized data (validates custom generators)
✅ **7. Premature Optimization**: Get correctness first, optimize only if needed
✅ **8. Cloud-First**: Parser works offline, no internet required

**GATE STATUS**: ✅ PASS - All constitutional requirements met

---

## Project Structure

### Documentation (this feature)

```
specs/006-f008-schema-parser/
├── plan.md              # This file (/speckit.plan command output)
├── spec.md              # Feature specification (user stories, requirements)
├── research.md          # Phase 0 output (10 technical decisions)
├── data-model.md        # Phase 1 output (Go struct design)
├── quickstart.md        # Phase 1 output (TDD workflow guide)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created yet)
```

### Source Code (repository root)

```
pkg/schema/
├── types.go           # Go structs (Schema, Table, Column, ForeignKey, etc.)
├── parser.go          # Parser functions (LoadSchema, ParseSchema, ValidateSchema)
└── parser_test.go     # TDD test suite (table-driven tests)
```

**Structure Decision**: Single project structure (Option 1). Parser is a library package (`pkg/schema`) used by CLI (`cmd/sourcebox`) and data generation engine (F013). No separate frontend/backend or mobile components needed.

---

## Phase 0: Research & Technical Decisions

**Status**: ✅ Complete
**Output**: [`research.md`](./research.md)

### Decisions Made

1. **Go Struct Design**: Structs mirror F007 JSON format with pointer/value type distinction
2. **JSON Unmarshaling**: Strict parsing with `DisallowUnknownFields()`, standard library only
3. **Validation Architecture**: Multi-phase fail-fast validation (schema → table → column → foreign keys → generation order)
4. **Error Message Quality**: Structured errors with context (table name, column name, index)
5. **File vs Embedded Loading**: Two functions (`LoadSchema` file path, `ParseSchema` io.Reader)
6. **Data Type Validation**: Case-insensitive prefix matching against supported types list
7. **Foreign Key Integrity**: Table-level validation in MVP, column-level in Phase 2
8. **Generation Order**: Completeness validation in MVP, circular dependency detection in Phase 2
9. **TDD Patterns**: Table-driven tests, strict TDD workflow, 100% coverage target
10. **Performance Optimization**: Simple linear-time algorithms with pre-allocated maps

All decisions align with constitutional principles (Boring Tech Wins, Developer-First, Ship Fast).

---

## Phase 1: Design & Contracts

**Status**: ✅ Complete
**Outputs**:
- [`data-model.md`](./data-model.md) - Go struct architecture
- [`quickstart.md`](./quickstart.md) - TDD workflow and verification guide
- Agent context updated: `CLAUDE.md`

### Data Model Summary

**Core Structs** (defined in `pkg/schema/types.go`):
- `Schema` - Top-level container (required fields as values, optional as pointers)
- `SchemaMetadata` - Metadata container
- `Table` - Table definition with columns and indexes
- `Column` - Column definition with type, constraints, generator, foreign key
- `ForeignKey` - Foreign key constraint (used as pointer in Column)
- `Index` - Table index (optional)
- `Relationship` - Explicit relationship documentation
- `ValidationRule` - Optional validation rule

**Design Principles**:
- Exact JSON format mirroring (F007 compliance)
- Pointer vs value types for required/optional distinction
- `map[string]interface{}` for flexible generator parameters
- JSON tags with snake_case (`json:"schema_version"`)

### API Contracts

**Public Functions**:

```go
// LoadSchema loads and parses a schema from file path
func LoadSchema(path string) (*Schema, error)

// ParseSchema parses a schema from any io.Reader
func ParseSchema(r io.Reader) (*Schema, error)
```

**Validation Functions** (internal):

```go
func ValidateSchema(s *Schema) error
func ValidateTable(t *Table, existingTables map[string]bool) error
func ValidateColumn(c *Column, existingColumns map[string]bool) error
func ValidateDataType(dataType string) error
func ValidateGenerationOrder(s *Schema, tableNames map[string]bool) error
```

**Error Handling**:
- All errors wrapped with context: `fmt.Errorf("context: %w", err)`
- Structured messages: "table 2 (loans): column 5 (borrower_id): foreign key references non-existent table 'users'"
- Fail-fast strategy: Return first error encountered

### Quickstart Summary

**TDD Workflow**:
1. Phase 1: Setup test structure (write first test before implementation)
2. Phase 2: Implement basic parser (make first test pass)
3. Phase 3: Add validation (TDD cycle: test → implement → refactor)
4. Phase 4: Edge cases (empty files, malformed JSON, unknown fields)

**Verification Checklist**:
- All tests pass (`go test ./pkg/schema/...`)
- Coverage is 100% (`go test -cover`)
- F007 example schema parses successfully
- Invalid schemas return clear, actionable errors
- Performance is <100ms for typical schemas

---

## Phase 2: Task Generation

**Status**: ⏸️ Pending
**Command**: `/speckit.tasks` (separate command, not part of `/speckit.plan`)

**Expected Tasks** (from TDD workflow):
1. Create `pkg/schema/types.go` with all Go structs
2. Create `pkg/schema/parser_test.go` with first test (valid schema)
3. Create `pkg/schema/parser.go` with basic parser (make test pass)
4. Add validation tests (TDD: one test at a time)
5. Implement validation functions (make each test pass)
6. Add edge case tests (empty files, malformed JSON, unknown fields)
7. Verify coverage is 100%
8. Manual QA on Mac/Linux/Windows

---

## Constitutional Re-Verification (Post-Design)

*Re-check after Phase 1 design to ensure no violations introduced*

### Design Alignment

✅ **Go Struct Design**: Simple, no over-engineering (Anti-Pattern 5: Over-Engineering)
✅ **JSON Parsing**: Standard library only (Core Principle IV: Boring Tech Wins)
✅ **Validation**: Fail-fast with clear errors (UX Principle 4: Fail Gracefully)
✅ **Performance**: Linear-time algorithms (<100ms target) (Core Principle II: Speed > Features)
✅ **Testing**: TDD workflow defined (Development Practice 1: TDD Required)
✅ **Error Messages**: Structured with context (Core Principle VI: Developer-First Design)

### No New Violations

✅ No external dependencies added (Go stdlib only)
✅ No complex algorithms introduced (simple map lookups)
✅ No premature optimization (correctness first)
✅ No cloud dependencies (works offline)
✅ No platform-specific code (cross-platform compatible)

**RE-VERIFICATION STATUS**: ✅ PASS - Design maintains constitutional compliance

---

## Complexity Tracking

*This section is empty because there are NO constitution violations to justify.*

All design decisions align with constitutional principles:
- Simple > Complex (standard library, no interfaces)
- Boring Tech Wins (Go stdlib `encoding/json`)
- Speed > Features (linear-time validation)
- Developer-First (clear error messages)
- TDD Required (100% coverage target)

---

## Implementation Notes

### For the Developer

**TDD Discipline** (NON-NEGOTIABLE):
1. Write failing test (RED)
2. Implement minimal code to pass (GREEN)
3. Refactor if needed (keep tests GREEN)
4. Repeat for each validation rule

**Test Coverage Target**: 100% for `pkg/schema/parser.go`

**Development Time Estimate**: 3 days (10-15 hours total)
- Day 1: Setup + basic parser + first validations (5 hours)
- Day 2: Complete all validations (5 hours)
- Day 3: Edge cases + manual QA + documentation (5 hours)

**Integration Points**:
- F013 Data Generation Engine: Uses `ParseSchema` to load schemas before generating data
- F010 Fintech Schema: Validated by parser during development
- F016 Healthcare Schema: Validated by parser during development
- F018 Retail Schema: Validated by parser during development

### For the AI Agent

When executing `/speckit.implement`:

1. **Start with tests**: Create `parser_test.go` before `parser.go`
2. **One validation at a time**: Don't implement all validations at once
3. **Use table-driven tests**: Pattern shown in quickstart.md
4. **Verify error messages**: Each test should check error message content
5. **Check coverage**: Run `go test -cover` after each validation added
6. **Manual verification**: Test with F007 example schema before completing

**Common Pitfalls to Avoid**:
- ❌ Don't skip TDD (100% coverage required)
- ❌ Don't use external JSON libraries (Go stdlib sufficient)
- ❌ Don't accumulate errors (fail-fast for clarity)
- ❌ Don't provide vague errors ("validation error" not acceptable)
- ❌ Don't skip edge cases (empty files, malformed JSON)

---

## Success Criteria

### Phase 0 (Research) ✅ Complete

- [x] All 10 technical decisions documented
- [x] Constitutional compliance verified
- [x] Alternatives considered for each decision
- [x] Source references provided

### Phase 1 (Design) ✅ Complete

- [x] Data model documented (Go struct architecture)
- [x] API contracts defined (public functions)
- [x] Quickstart guide created (TDD workflow)
- [x] Agent context updated (CLAUDE.md)
- [x] Constitutional re-verification passed

### Phase 2 (Tasks) ⏸️ Pending

- [ ] Task breakdown generated (`/speckit.tasks` command)
- [ ] Dependencies identified
- [ ] Time estimates provided
- [ ] Implementation order defined

### Implementation (Future) ⏸️ Pending

- [ ] All tests pass (`go test ./pkg/schema/...`)
- [ ] Coverage is 100% (`go test -cover`)
- [ ] F007 example schema parses successfully
- [ ] Invalid schemas return clear, actionable errors
- [ ] Performance is <100ms for typical schemas
- [ ] Manual QA passed on Mac/Linux/Windows

---

## Related Documentation

- **F007 Schema Specification**: `/schemas/schema-spec.md` - JSON format definition
- **Feature Specification**: [`spec.md`](./spec.md) - User stories and requirements
- **Research Decisions**: [`research.md`](./research.md) - Technical decision rationale
- **Data Model**: [`data-model.md`](./data-model.md) - Go struct architecture
- **Quickstart Guide**: [`quickstart.md`](./quickstart.md) - TDD workflow and verification
- **Constitution**: `/.specify/memory/constitution.md` - Project principles and constraints

---

## Next Steps

1. **Execute `/speckit.tasks`**: Generate task breakdown for implementation
2. **Execute `/speckit.implement`**: Begin TDD implementation
3. **Integrate with F013**: Use parser in data generation engine
4. **Validate F010/F016/F018**: Use parser to validate verticalized schemas

---

**Version**: 1.0.0
**Status**: Design Complete, Ready for Task Generation
**Last Updated**: 2025-10-15
