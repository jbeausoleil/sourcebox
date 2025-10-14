# Feature Specification Prompt: F014 - Edge Case & Null Value Generation

## Feature Metadata
- **Feature ID**: F014
- **Name**: Edge Case & Null Value Generation
- **Category**: Data Generation
- **Phase**: Week 6
- **Priority**: P1 (Should-have)
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F013 (Data generation engine must exist)

## Constitutional Alignment

### Core Principles
- ✅ **Verticalized > Generic**: Edge cases reflect real-world data issues
- ✅ **Developer-First**: Realistic test data includes edge cases
- ✅ **Quality Standards**: Data quality includes realistic imperfections

### Technical Constraints
- ✅ **Code Quality**: TDD for edge case logic
- ✅ **Performance**: Minimal performance impact

### Development Practices
- ✅ **TDD Required**: Core generation logic

## User Story
**US-MVP-008**: "As a QA engineer, I need realistic edge cases in my test data (negative balances, very high/low values, null optionals) so I can test how my application handles unusual scenarios."

## Problem Statement
Perfect test data is unrealistic. Production databases contain:
- **Edge cases**: Extreme values (very high/low credit scores, outlier loan amounts)
- **Null values**: Optional fields that are sometimes empty
- **Unusual combinations**: Valid but uncommon data patterns
- **Boundary values**: Min/max values for fields

Without edge cases, developers miss bugs that only surface with unusual data. Test data should include ~10% edge cases and respect nullable fields.

## Solution Overview
Enhance the data generation engine (F013) to support:
1. **Null probability**: Schema specifies nullable fields, engine respects it
2. **Edge case generation**: 10% of records include extreme values
3. **Boundary testing**: Generate min/max values occasionally
4. **Configurable via schema**: `null_probability`, `edge_case_probability` parameters

## Detailed Requirements

### Acceptance Criteria
1. **Null Value Support**: Nullable columns can generate null values
2. **Null Probability**: Configurable per column (default: 10% for nullable fields)
3. **Edge Case Generation**: 10% of records include extreme values
4. **Edge Case Types**:
   - Very high/low credit scores (300, 850)
   - Extreme loan amounts ($5K, $500K)
   - Future dates (for testing date validation)
   - Negative balances (for financial edge cases)
5. **Schema Configuration**: Support `null_probability`, `edge_case_probability` in generator params
6. **Unit Tests**: Verify edge cases appear ~10% of the time, nulls appear at specified probability

### Technical Specifications

#### Enhanced Engine: Update `pkg/generators/engine.go`

```go
// Add to generateColumnValue function

func (e *Engine) generateColumnValue(table *schema.Table, col *schema.Column, rowID int) (interface{}, error) {
	// Check if we should generate null for nullable column
	if col.Nullable {
		nullProb := e.getNullProbability(col)
		if rand.Float64() < nullProb {
			return nil, nil
		}
	}

	// Check if we should generate edge case
	edgeCaseProb := e.getEdgeCaseProbability(col)
	if rand.Float64() < edgeCaseProb {
		return e.generateEdgeCase(col)
	}

	// Normal generation flow...
	// [existing code]
}

// getNullProbability returns null probability for column (0.0-1.0)
func (e *Engine) getNullProbability(col *schema.Column) float64 {
	// Check generator params for custom null probability
	if prob, ok := col.GeneratorParams["null_probability"].(float64); ok {
		return prob
	}

	// Default: 10% null for nullable fields
	if col.Nullable {
		return 0.10
	}

	return 0.0
}

// getEdgeCaseProbability returns edge case probability (0.0-1.0)
func (e *Engine) getEdgeCaseProbability(col *schema.Column) float64 {
	// Check generator params
	if prob, ok := col.GeneratorParams["edge_case_probability"].(float64); ok {
		return prob
	}

	// Default: 10% edge cases
	return 0.10
}

// generateEdgeCase generates an edge case value for column
func (e *Engine) generateEdgeCase(col *schema.Column) (interface{}, error) {
	switch col.Generator {
	case "credit_score":
		// Extreme credit scores (300 or 850)
		if rand.Float64() < 0.5 {
			return 300, nil // Very low
		}
		return 850, nil // Perfect score

	case "loan_amount":
		// Extreme loan amounts ($5K or $500K)
		if rand.Float64() < 0.5 {
			return 5000.0, nil // Minimum
		}
		return 500000.0, nil // Maximum

	case "interest_rate":
		// Extreme rates (3% or 15%)
		if rand.Float64() < 0.5 {
			return 3.0, nil // Best rate
		}
		return 15.0, nil // Worst rate

	case "annual_income":
		// Extreme incomes
		if rand.Float64() < 0.5 {
			return 25000.0, nil // Minimum
		}
		return 500000.0, nil // Very high

	case "timestamp_past", "date_past":
		// Very old date (edge case for historical data)
		daysAgo := 3650 // 10 years ago
		return TimestampPast(daysAgo), nil

	case "int_range":
		// Min or max from params
		min, _ := col.GeneratorParams["min"].(float64)
		max, _ := col.GeneratorParams["max"].(float64)
		if rand.Float64() < 0.5 {
			return int(min), nil
		}
		return int(max), nil

	default:
		// For other generators, fall back to normal generation
		return e.invokeGenerator(col.Generator, col.GeneratorParams)
	}
}
```

#### Edge Case Generators: Add to `pkg/generators/fintech.go`

```go
// NegativeBalance generates a negative balance (edge case)
func NegativeBalance() float64 {
	// Small negative value (-$1 to -$500)
	return -1.0 - rand.Float64()*499.0
}

// FutureDate generates a date in the future (for testing validation)
func FutureDate() time.Time {
	daysAhead := rand.Intn(365) + 1
	return time.Now().AddDate(0, 0, daysAhead)
}

// ExtremelyLargeLoanAmount generates unrealistic loan amount (edge case)
func ExtremelyLargeLoanAmount() float64 {
	// $1M - $10M (unrealistic for most lending)
	return 1000000.0 + rand.Float64()*9000000.0
}
```

#### Unit Tests: `pkg/generators/edge_cases_test.go`

```go
package generators

import (
	"testing"
)

func TestNullValueGeneration(t *testing.T) {
	// Load schema with nullable columns
	s, err := schema.LoadSchema("../../schemas/fintech-loans.json")
	if err != nil {
		t.Fatalf("failed to load schema: %v", err)
	}

	engine := NewEngine(s)
	dataset, err := engine.Generate()
	if err != nil {
		t.Fatalf("generation failed: %v", err)
	}

	// Check borrowers table - phone is nullable
	borrowers := dataset.Tables["borrowers"]
	nullCount := 0

	for _, row := range borrowers.Rows {
		if row["phone"] == nil {
			nullCount++
		}
	}

	// Should have some nulls (approximately 10% = 25 out of 250)
	// Allow tolerance: 5-50 nulls
	if nullCount < 5 || nullCount > 50 {
		t.Errorf("Expected ~10%% nulls for phone field (25 out of 250), got %d", nullCount)
	}
}

func TestEdgeCaseGeneration(t *testing.T) {
	// Generate many credit scores and check for edge cases
	edgeCount := 0
	samples := 1000

	for i := 0; i < samples; i++ {
		score := CreditScore()

		// Edge cases: 300 or 850
		if score == 300 || score == 850 {
			edgeCount++
		}
	}

	// Should have approximately 10% edge cases (100 out of 1000)
	// Allow tolerance: 50-150
	if edgeCount < 50 || edgeCount > 150 {
		t.Errorf("Expected ~10%% edge cases (100 out of 1000), got %d", edgeCount)
	}
}

func TestNegativeBalance(t *testing.T) {
	balance := NegativeBalance()

	if balance >= 0 {
		t.Error("NegativeBalance should return negative value")
	}

	if balance < -500 || balance > -1 {
		t.Errorf("NegativeBalance should be between -$500 and -$1, got %.2f", balance)
	}
}

func TestFutureDate(t *testing.T) {
	futureDate := FutureDate()
	now := time.Now()

	if !futureDate.After(now) {
		t.Error("FutureDate should return date in the future")
	}

	// Should be within 365 days
	maxDate := now.AddDate(0, 0, 365)
	if futureDate.After(maxDate) {
		t.Error("FutureDate should be within next 365 days")
	}
}
```

#### Schema Example (F010 enhancement)

```json
{
  "name": "phone",
  "type": "varchar(20)",
  "generator": "phone",
  "nullable": true,
  "generator_params": {
    "null_probability": 0.10
  }
}
```

### Performance Considerations
- Minimal overhead (single random check per column)
- No significant performance impact
- Still meets > 100 records/sec target

### Testing Strategy
- Verify null values appear at specified probability
- Verify edge cases appear ~10% of the time
- Verify edge case values are realistic extremes
- Performance tests ensure no slowdown

## Dependencies
- **Upstream**: F013 (Engine must exist to enhance)
- **Downstream**: F020 (Integration tests verify edge cases)

## Deliverables
1. Enhanced engine with null and edge case support
2. Edge case generator functions
3. Unit tests for null/edge case generation
4. Test coverage for edge case logic

## Success Criteria
- ✅ Nullable columns generate nulls ~10% of time
- ✅ Edge cases appear ~10% of records
- ✅ Edge case values are realistic extremes
- ✅ Tests pass with proper distribution
- ✅ No performance degradation

## Anti-Patterns to Avoid
- ❌ Too many edge cases (> 20% unrealistic)
- ❌ Unrealistic edge cases (e.g., negative credit scores)
- ❌ Ignoring nullable columns
- ❌ Performance degradation
- ❌ No tests for edge case logic

## Implementation Notes
- Edge cases are valuable for QA testing
- 10% is realistic (not too much, not too little)
- Null probability should be configurable per column
- Edge cases should still respect data type constraints

## TDD Requirements
**Required** - Edge case logic is core functionality:

1. Write test: Verify nulls appear at specified rate
2. Run test - should fail
3. Implement: Null probability logic
4. Run test - should pass
5. Repeat for edge cases

## Related Constitution Sections
- **Verticalized > Generic (Principle I)**: Edge cases reflect real-world issues
- **Quality Standards (Technical Constraint 5)**: Include realistic imperfections
- **TDD Required (Development Practice 1)**: Core generation logic
