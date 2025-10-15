# F007 Quickstart: Schema Format Specification Validation

**Feature**: Schema JSON Format Specification
**Purpose**: Validate that the schema format specification is complete, implementable, and constitutional
**Date**: 2025-10-15

---

## Prerequisites

- Constitution read (lines 176-197 for schema complexity tiers)
- F007 spec read (understanding of requirements: specs/005-f007-schema-json/spec.md)
- JSON syntax validator available (jsonlint, jq, IDE, or online tool like https://jsonlint.com/)
- Understanding of Tier 1 complexity constraints (<5,000 records, <30s generation)

---

## Validation Overview

The Schema JSON Format Specification is validated by:
1. Creating comprehensive specification document (`schemas/schema-spec.md`)
2. Creating example schema demonstrating all features (`schemas/example-schema.json`)
3. Verifying JSON syntax is valid
4. Verifying all sections are documented
5. Verifying specification is implementable by F008 parser

This quickstart guides you through validating the specification, NOT implementing code. F007 is design work, F008 is implementation.

---

## Manual Validation Steps

### 1. Verify Specification Document Completeness

**Check**: Does `schemas/schema-spec.md` exist and contain all required sections?

```bash
# Check file exists
ls -lh schemas/schema-spec.md

# Check file size (should be substantial, >10KB)
wc -l schemas/schema-spec.md

# Verify section headers exist
grep -E "^##? " schemas/schema-spec.md
```

**Required sections checklist**:
- [ ] Overview (what is the format, why JSON)
- [ ] Schema Structure (top-level fields: schema_version, name, description, etc.)
- [ ] Table Definition (table object structure)
- [ ] Column Definition (column object structure, types, generators)
- [ ] Generator Types (built-in and custom generators)
- [ ] Distribution Types (uniform, normal, lognormal, weighted, ranges)
- [ ] Foreign Key Relationships (inline foreign_key object and explicit relationships section)
- [ ] Schema Versioning (semver rules: major, minor, patch)
- [ ] Validation Rules (what makes a schema valid)
- [ ] Supported Data Types (MySQL and PostgreSQL common types)
- [ ] Example Schema (reference to example-schema.json)

**Expected outcome**: All sections present with clear explanations, code examples, and practical guidance. No [TODO] or [NEEDS CLARIFICATION] markers remaining.

**Common issues**:
- Missing sections → Incomplete specification
- Vague language → F008 parser can't implement
- No examples → Developers can't copy patterns

---

### 2. Verify Example Schema JSON Syntax

**Check**: Is `schemas/example-schema.json` valid JSON?

```bash
# Validate JSON syntax with Python
cat schemas/example-schema.json | python -m json.tool > /dev/null && echo "✅ Valid JSON" || echo "❌ Invalid JSON"

# Or use jq
cat schemas/example-schema.json | jq '.' > /dev/null && echo "✅ Valid JSON" || echo "❌ Invalid JSON"

# Or copy-paste to https://jsonlint.com/ for browser-based validation
```

**Expected outcome**: Valid JSON with no syntax errors, properly formatted with indentation.

**Common issues**:
- Missing commas between array/object elements
- Trailing commas (not allowed in JSON)
- Unquoted keys or values
- Unescaped special characters in strings

---

### 3. Verify Example Schema Demonstrates All Features

**Check**: Does `schemas/example-schema.json` include examples of all documented features?

```bash
# Check top-level fields exist
cat schemas/example-schema.json | jq 'keys'

# Check for schema_version (format version)
cat schemas/example-schema.json | jq '.schema_version'

# Check for version (content version, semver format)
cat schemas/example-schema.json | jq '.version'

# Check for database_type array
cat schemas/example-schema.json | jq '.database_type'

# Check for metadata object
cat schemas/example-schema.json | jq '.metadata'

# Check for tables array with multiple tables
cat schemas/example-schema.json | jq '.tables | length'

# Check for relationships section
cat schemas/example-schema.json | jq '.relationships | length'

# Check for generation_order array
cat schemas/example-schema.json | jq '.generation_order'
```

**Feature checklist**:
- [ ] schema_version field (format version, should be "1.0")
- [ ] name, description, author, version (metadata fields)
- [ ] database_type array (should include "mysql" and "postgres")
- [ ] metadata object (industry, tags, total_records)
- [ ] tables array with multiple tables (2-4 tables for Tier 1)
- [ ] columns with various types (int, varchar, decimal, timestamp, etc.)
- [ ] Primary keys (auto_increment or similar)
- [ ] Foreign keys (inline foreign_key object)
- [ ] Indexes (unique and non-unique)
- [ ] Built-in generators (first_name, last_name, email, etc.)
- [ ] Custom generators (credit_score, loan_amount, etc.)
- [ ] Generator parameters (distribution, mean, std_dev, min, max, weights)
- [ ] Relationships section (explicit documentation)
- [ ] generation_order array (parent tables first)
- [ ] validation_rules section (optional soft constraints)

**Expected outcome**: Example schema demonstrates ALL documented features, nothing is documented but not shown.

**Common issues**:
- Feature documented but not demonstrated in example
- Example shows features not documented in spec
- Inconsistent naming (spec says "generator_params", example says "generatorParams")

---

### 4. Verify Tier 1 Complexity Compliance

**Check**: Does example schema comply with Tier 1 constitutional constraint (<5,000 records, <30s generation)?

```bash
# Check total_records in metadata
cat schemas/example-schema.json | jq '.metadata.total_records'

# Calculate total records by summing table record_count fields
cat schemas/example-schema.json | jq '[.tables[].record_count] | add'
```

**Expected outcome**: Total records <5,000 (Tier 1 complexity, constitutional constraint from lines 176-197)

**Example breakdown**:
- borrowers: 250 records
- loans: 1,000 records
- payments: 3,700 records
- Total: 4,950 records ✅ (within Tier 1)

**Common issues**:
- Total records >5,000 → Exceeds Tier 1 complexity
- metadata.total_records doesn't match sum of table record_counts → Inconsistency
- Complex relationships (10+ foreign keys) → May slow generation below 30s

---

### 5. Verify Schema Versioning Documentation

**Check**: Is the versioning strategy clearly documented?

```bash
# Check schema-spec.md has versioning section
grep -A 30 "Schema Versioning" schemas/schema-spec.md | head -35

# Or search for semver-related terms
grep -i "major\|minor\|patch" schemas/schema-spec.md
```

**Versioning checklist**:
- [ ] schema_version field explained (format version, currently 1.0)
- [ ] version field explained (content version, semver format)
- [ ] Major version rules documented (breaking changes: table removal, type changes)
- [ ] Minor version rules documented (backward-compatible: new tables, new columns)
- [ ] Patch version rules documented (bug fixes: distribution improvements, docs)
- [ ] Examples of version bumps (1.0.0 → 1.1.0 vs 1.0.0 → 2.0.0)

**Expected outcome**: Clear versioning strategy documented with examples, follows semantic versioning (semver) conventions.

**Common issues**:
- Versioning strategy not documented → Schema evolution is unclear
- Examples missing → Developers don't know when to bump versions
- Conflicting guidance (is adding a column major or minor?) → Ambiguous

---

### 6. Verify Validation Rules Documentation

**Check**: Are validation rules comprehensively documented?

```bash
# Check schema-spec.md has validation section
grep -A 50 "Validation Rules" schemas/schema-spec.md | head -55

# Or search for validation-related terms
grep -i "required\|optional\|validation" schemas/schema-spec.md
```

**Validation rules checklist**:
- [ ] Schema-level validation (unique name, at least one table, valid database_type)
- [ ] Table-level validation (unique name, at least one column, exactly one primary key, record_count > 0)
- [ ] Column-level validation (valid types, valid generators, generator params match, foreign keys reference existing tables/columns)
- [ ] Relationship validation (foreign keys reference existing tables, foreign keys reference primary keys or unique columns, referential integrity actions are valid)
- [ ] Generation order validation (parent tables first, no circular dependencies)
- [ ] Optional validation_rules section explained (soft constraints vs hard validation)

**Expected outcome**: Comprehensive validation rules that F008 parser can implement unambiguously.

**Common issues**:
- Validation rules too vague ("tables should be valid") → Not implementable
- Edge cases not documented (circular dependencies, missing foreign keys) → Incomplete
- Required vs optional fields unclear → Schema authors confused

---

### 7. Verify Generator Documentation

**Check**: Are all generator types documented with parameters and examples?

```bash
# Check schema-spec.md has generator section
grep -A 80 "Generator Types" schemas/schema-spec.md | head -85

# Or search for specific generator names
grep -i "first_name\|email\|credit_score\|timestamp" schemas/schema-spec.md
```

**Generator types checklist**:
- [ ] Built-in personal data generators (first_name, last_name, full_name, email, phone, address, ssn, date_of_birth)
- [ ] Built-in company data generators (company_name, job_title, company_email, domain)
- [ ] Built-in date/time generators (timestamp_past, timestamp_future, date_between)
- [ ] Built-in numeric generators (int_range, float_range, decimal_range)
- [ ] Custom generator pattern (how schemas define custom generators like credit_score, loan_amount)
- [ ] Generator parameter structure (generator_params object with distribution, mean, std_dev, etc.)
- [ ] Distribution types with parameters (uniform: min/max, normal: mean/std_dev/min/max, lognormal: median/min/max, weighted: values array, ranges: ranges array)
- [ ] Examples for each generator type

**Expected outcome**: All generator types documented with parameter schemas and copyable examples.

**Common issues**:
- Generator documented but parameters unclear → Developers can't configure
- No examples → Developers don't know how to use
- Inconsistent parameter names (mean vs average) → Confusing

---

### 8. Verify Data Type Documentation

**Check**: Are all supported SQL data types documented?

```bash
# Check schema-spec.md has data types section
grep -A 40 "Supported Data Types" schemas/schema-spec.md | head -45

# Or search for specific type names
grep -i "varchar\|int\|decimal\|timestamp\|json" schemas/schema-spec.md
```

**Data types checklist**:
- [ ] Integer types (int, bigint, smallint, tinyint)
- [ ] Decimal types (decimal(p,s), float, double)
- [ ] String types (varchar(n), text, char(n))
- [ ] Date/Time types (date, datetime, timestamp)
- [ ] Boolean type (boolean for PostgreSQL, bit for MySQL)
- [ ] JSON types (json for both, jsonb for PostgreSQL)
- [ ] Enum types (enum('val1','val2') format, platform differences documented)
- [ ] Type format specifications (varchar(n), decimal(p,s) - parameters explained)

**Expected outcome**: All supported SQL data types documented with format specifications and platform-specific guidance.

**Common issues**:
- Platform-specific types not documented (jsonb) → Developers confused
- Type parameters missing (varchar needs length) → Incomplete documentation
- MySQL vs PostgreSQL differences not explained → Cross-database compatibility unclear

---

### 9. Verify Foreign Key Relationship Documentation

**Check**: Is the foreign key format clearly documented?

```bash
# Check schema-spec.md has foreign key section
grep -A 60 "Foreign Key Relationships" schemas/schema-spec.md | head -65

# Or search for foreign key terms
grep -i "foreign_key\|on_delete\|on_update\|relationship" schemas/schema-spec.md
```

**Foreign key checklist**:
- [ ] Inline foreign_key object structure (table, column, on_delete, on_update)
- [ ] Referential integrity actions documented (CASCADE, SET NULL, RESTRICT)
- [ ] Explicit relationships section format (from_table, from_column, to_table, to_column, relationship_type, description)
- [ ] Relationship types explained (one_to_one, one_to_many, many_to_one, many_to_many)
- [ ] Why both inline and explicit? (rationale: inline for code, explicit for documentation)

**Expected outcome**: Foreign key format is clear, both inline and explicit representations documented with rationale.

**Common issues**:
- Only inline or only explicit → Incomplete pattern
- Rationale missing → Developers don't understand why both
- Referential integrity actions not explained → Unclear what CASCADE does

---

### 10. Verify F008 Implementability

**Check**: Can F008 parser implement this specification unambiguously?

**Manual review questions**:
1. Are all required fields clearly marked? (Look for "required" vs "optional" labels)
2. Are all optional fields clearly marked?
3. Are validation rules specific enough to implement? (Example: "at least one table" is specific, "tables should be valid" is not)
4. Are error messages for invalid schemas described? (What error for missing primary key? Missing foreign key target?)
5. Can you write pseudocode for validation from the spec? (If not, spec is too vague)
6. Are there any ambiguous phrases? ("some tables may...", "typically this is...", "consider doing...")
7. Are all examples consistent with rules? (Example schema follows all validation rules)

**Expected outcome**: Specification is unambiguous, implementable by F008 parser with clear error messages.

**Common issues**:
- Vague language ("typically", "usually", "consider") → Ambiguous
- Missing error message guidance → Poor developer experience
- Conflicting rules (spec says X, example shows Y) → Inconsistent
- Edge cases not handled (what if generation_order is missing?) → Incomplete

---

## Verification Checklist

After completing all validation steps, verify:

- [ ] `schemas/schema-spec.md` exists with all required sections
- [ ] `schemas/example-schema.json` exists and is valid JSON
- [ ] Example schema demonstrates all documented features
- [ ] Example schema is Tier 1 complexity (<5,000 records, <30s generation)
- [ ] Schema versioning strategy documented (semver: major, minor, patch)
- [ ] Validation rules comprehensively documented
- [ ] All generator types documented with parameters
- [ ] All supported data types documented (MySQL/PostgreSQL common subset)
- [ ] Foreign key relationships clearly documented (inline + explicit)
- [ ] Specification is unambiguous and implementable by F008
- [ ] Example schema can be used as template for F010, F016, F018
- [ ] Format supports custom generators (verticalized data)
- [ ] Format supports multiple databases (MySQL, PostgreSQL)
- [ ] Format is extensible (future generators, databases, features)
- [ ] No [TODO] or [NEEDS CLARIFICATION] markers remain
- [ ] All examples are copyable and executable (valid JSON)

---

## Common Issues and Fixes

### Issue: JSON syntax errors in example schema

**Symptoms**:
- `python -m json.tool` fails with "Expecting ',' delimiter" or similar
- jsonlint shows "Parse error on line X"
- jq fails with "parse error: Invalid numeric literal"

**Check**:
- Missing commas between array/object elements
- Trailing commas (JSON doesn't allow them)
- Unquoted keys (keys must be strings: `"name": "value"`)
- Unescaped special characters (use `\"` for quotes in strings)

**Fix**:
```bash
# Use JSON validator to identify exact error location
cat schemas/example-schema.json | python -m json.tool

# Or use jq with --indent for pretty-printing
cat schemas/example-schema.json | jq --indent 2 '.' > schemas/example-schema-formatted.json
```

---

### Issue: Example schema too complex (>30s generation)

**Symptoms**:
- Total records >10,000
- Many foreign key lookups (10+ foreign keys)
- Complex generators (nested lookups, correlated data)

**Check**:
```bash
# Count total records
cat schemas/example-schema.json | jq '[.tables[].record_count] | add'

# Count foreign keys
cat schemas/example-schema.json | jq '[.tables[].columns[] | select(.foreign_key != null)] | length'
```

**Fix**:
- Reduce record counts (target <5,000 for Tier 1)
- Simplify relationships (2-3 foreign keys maximum for example)
- Use simpler generators (built-in over custom)
- Target Tier 1 complexity (<30s, MVP focus)

---

### Issue: Specification ambiguous

**Symptoms**:
- "Can you implement a parser from this spec?" → Unclear/uncertain
- Required vs optional fields not marked
- Validation rules use vague language ("should", "typically", "consider")

**Check**:
- Search for ambiguous phrases: `grep -i "should\|typically\|usually\|consider\|may" schemas/schema-spec.md`
- Required fields marked? `grep -i "required" schemas/schema-spec.md`
- Examples for each rule? `grep -A 5 "Example:" schemas/schema-spec.md`

**Fix**:
- Replace "should" with "MUST" or "MAY" (RFC 2119 keywords)
- Add "Required" vs "Optional" labels to all field descriptions
- Add code examples for each validation rule
- Document error messages for invalid schemas

---

### Issue: Generator parameters unclear

**Symptoms**:
- Parameter names unclear (mean vs average?)
- Parameter types not documented (number vs string?)
- No examples provided

**Check**:
```bash
# Search for generator parameter examples
grep -A 10 "generator_params" schemas/schema-spec.md
```

**Fix**:
- Add generator parameter schema for each generator type:
  ```
  normal distribution:
  - mean (required, number): Mean value
  - std_dev (required, number): Standard deviation
  - min (optional, number): Minimum value
  - max (optional, number): Maximum value
  ```
- Add copyable JSON examples for each generator
- Document parameter types (number, string, array, object)

---

### Issue: Validation rules missing edge cases

**Symptoms**:
- Circular dependencies not mentioned
- Missing foreign key targets not mentioned
- Invalid generator parameters not mentioned

**Check**:
```bash
# Search for edge case handling
grep -i "circular\|missing\|invalid\|error" schemas/schema-spec.md
```

**Fix**:
- Add "Edge Cases" subsection to validation rules
- Document what happens when:
  - generation_order has circular dependencies (A → B → C → A)
  - Foreign key references non-existent table
  - Generator parameters are invalid (mean without std_dev)
  - metadata.total_records doesn't match sum of table record_counts
- Provide clear error messages for each edge case

---

### Issue: Foreign key syntax confusing

**Symptoms**:
- Inline vs explicit - are both needed?
- Purpose of each representation unclear
- Conflicting examples (inline says CASCADE, explicit says RESTRICT)

**Check**:
```bash
# Search for foreign key documentation
grep -A 20 "foreign_key" schemas/schema-spec.md
```

**Fix**:
- Document rationale for dual representation:
  - **Inline foreign_key**: What parser uses (part of column definition, SQL constraint)
  - **Explicit relationships**: Documentation for humans (relationship overview, visualization)
- Ensure examples are consistent (inline and explicit show same relationships)
- Add diagram showing relationship structure

---

## Performance Validation (Not Applicable for F007)

**Note**: F007 is specification work only. Performance validation happens in:
- **F008**: Schema parser implementation (parsing time, validation time)
- **F010/F016/F018**: Schema implementations (data generation time)

For F007, focus on **design quality** (simplicity, clarity, implementability), not runtime performance.

---

## Next Steps

After completing F007 validation:

1. **F008 (Schema Parser)**: Implement schema parser in Go using this specification
2. **F010 (Fintech Schema)**: Create fintech-loans schema using this format
3. **F016 (Healthcare Schema)**: Create healthcare schema using this format
4. **F018 (Retail Schema)**: Create retail schema using this format
5. **Community Contributions**: Enable schema contributions using this format (Phase 2)

---

## Design Quality Checklist

### Simplicity ✅
- [ ] JSON format is simple, readable, human-editable
- [ ] No unnecessary nesting or complexity
- [ ] Field names are clear, self-explanatory (generator_params, foreign_key, generation_order)
- [ ] Examples are easy to understand and copy

### Extensibility ✅
- [ ] New generators can be added without breaking existing schemas (custom generators)
- [ ] New databases can be added without format changes (database_type array)
- [ ] New column attributes can be added (forward compatibility, unknown attrs ignored)
- [ ] schema_version allows future format evolution (1.0 → 2.0)

### Clarity ✅
- [ ] Specification answers common questions (How do I...?)
- [ ] Required vs optional fields are explicit (marked in docs)
- [ ] Validation rules are unambiguous (specific, not vague)
- [ ] Error conditions are documented (what error for missing primary key?)

### Realism ✅
- [ ] Format enables verticalized data (custom generators: credit_score, loan_amount)
- [ ] Supports realistic distributions (normal, lognormal, weighted, ranges)
- [ ] Supports complex relationships (foreign keys, indexes, generation_order)
- [ ] Example schema demonstrates real-world fintech data (not "John Doe")

### Implementability ✅
- [ ] F008 parser can be implemented from this spec (unambiguous)
- [ ] Validation rules are specific enough to implement (concrete rules, not guidelines)
- [ ] Error messages for invalid schemas are described (what to show users)
- [ ] Edge cases are documented (circular deps, missing FKs, invalid params)

---

## Success Indicators

You know F007 is complete when:

✅ All 10 verification steps pass
✅ All checklists are 100% complete
✅ No [TODO] or [NEEDS CLARIFICATION] markers remain
✅ Example schema demonstrates ALL documented features
✅ Specification can be implemented by F008 without ambiguity
✅ Constitutional compliance verified (Tier 1 complexity, verticalized data, boring tech)
✅ Ready to proceed to task generation (/speckit.tasks)

---

**Validation Status**: Ready for implementation
**Constitutional Violations**: None expected
**Next Command**: `/speckit.tasks` to generate tasks.md from plan.md
