# Feature Specification Prompt: F012 - Custom Fintech Data Generators

## Feature Metadata
- **Feature ID**: F012
- **Name**: Custom Fintech Data Generators
- **Category**: Data Generation
- **Phase**: Week 5-6
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Medium (4 days)
- **Dependencies**: F011 (Base generators must exist)

## Constitutional Alignment

### Core Principles
- ✅ **Verticalized > Generic (NON-NEGOTIABLE)**: Fintech data MUST look like fintech
- ✅ **Realism > Convenience**: Real-world distributions, not uniform random
- ✅ **Quality Bar**: If not 10x better than Faker, don't ship it

### Technical Constraints
- ✅ **Performance**: 100+ records/second generation
- ✅ **Code Quality**: TDD required, >80% coverage

### Development Practices
- ✅ **TDD Required for Core Functionality**: Custom generators are core logic

## User Story
**US-MVP-007**: "As a fintech developer, I want loan amounts, credit scores, and interest rates that follow realistic distributions so my demos look professional to investors."

## Problem Statement
Generic fake data fails for fintech demos. Faker.js gives uniform random values:
- ❌ Loan amounts: $50,000, $50,001, $50,002 (unrealistic uniform distribution)
- ❌ Credit scores: 300, 850, 575 (should be bell curve around 680)
- ❌ Interest rates: 3%, 14%, 8% (should weight toward 3-6% for most loans)
- ❌ Loan status: Even split between active/paid/delinquent (unrealistic - most loans are active)

SourceBox needs fintech-specific generators with realistic distributions:
- **Log-normal** for loan amounts (many small, few large)
- **Normal (bell curve)** for credit scores (centered around 680)
- **Weighted ranges** for interest rates (most 3-6%, some 6-10%, few 10-15%)
- **Weighted values** for loan status (70% active, 25% paid, 5% delinquent)

## Solution Overview
Implement custom fintech generators in `pkg/generators/fintech.go` with realistic distributions: LoanAmount (log-normal), CreditScore (normal), InterestRate (weighted ranges), LoanStatus (weighted values). Use statistical distributions, not uniform random. Write comprehensive unit tests with distribution validation (chi-square tests to verify realism).

## Detailed Requirements

### Acceptance Criteria
1. **Custom Generator Module**: `pkg/generators/fintech.go` created
2. **Loan Amount Generator**: Log-normal distribution ($5K-$500K, median $50K)
3. **Interest Rate Generator**: Weighted ranges (3-6%=60%, 6-10%=30%, 10-15%=10%)
4. **Credit Score Generator**: Normal distribution (mean=680, std=80, range 300-850)
5. **Loan Status Generator**: Weighted values (active=70%, paid_off=25%, delinquent=5%)
6. **Additional Generators**: LoanNumber, LoanTerm, AnnualIncome
7. **Distribution Tests**: Unit tests verify distributions match targets (chi-square goodness-of-fit)
8. **Histogram Plots** (optional): Visual verification of distributions

### Technical Specifications

#### Fintech Generators: `pkg/generators/fintech.go`

```go
package generators

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// LoanAmount generates realistic loan amounts using log-normal distribution
// Median: $50,000, Range: $5,000 - $500,000
// Most loans cluster around $20K-$100K, fewer at extremes
func LoanAmount() float64 {
	// Log-normal distribution parameters
	// mu and sigma calculated to achieve desired median and range
	mu := 10.82      // ln(50000)
	sigma := 0.8

	// Generate log-normal value
	normal := rand.NormFloat64()
	lnValue := mu + sigma*normal
	amount := math.Exp(lnValue)

	// Clamp to range $5,000 - $500,000
	if amount < 5000 {
		amount = 5000
	}
	if amount > 500000 {
		amount = 500000
	}

	// Round to nearest $100
	return math.Round(amount/100) * 100
}

// InterestRate generates realistic interest rates with weighted ranges
// 3-6% (60%), 6-10% (30%), 10-15% (10%)
func InterestRate() float64 {
	roll := rand.Float64() * 100 // 0-100

	var rate float64
	if roll < 60 {
		// 3-6% range (prime rates)
		rate = 3.0 + rand.Float64()*3.0
	} else if roll < 90 {
		// 6-10% range (standard rates)
		rate = 6.0 + rand.Float64()*4.0
	} else {
		// 10-15% range (subprime rates)
		rate = 10.0 + rand.Float64()*5.0
	}

	// Round to 2 decimal places
	return math.Round(rate*100) / 100
}

// CreditScore generates realistic credit scores using normal distribution
// Mean: 680, Standard Deviation: 80, Range: 300-850 (FICO scale)
func CreditScore() int {
	mean := 680.0
	stdDev := 80.0

	// Generate normal distribution value
	score := mean + rand.NormFloat64()*stdDev

	// Clamp to valid FICO range (300-850)
	if score < 300 {
		score = 300
	}
	if score > 850 {
		score = 850
	}

	return int(math.Round(score))
}

// LoanStatus generates realistic loan statuses with weighted probabilities
// Active: 70%, Paid Off: 25%, Delinquent: 5%
func LoanStatus() string {
	roll := rand.Float64() * 100

	if roll < 70 {
		return "active"
	} else if roll < 95 {
		return "paid_off"
	} else {
		return "delinquent"
	}
}

// LoanNumber generates a unique loan number (format: LN-YYYYMMDD-XXXXX)
func LoanNumber() string {
	date := time.Now().Format("20060102")
	random := rand.Intn(99999)
	return fmt.Sprintf("LN-%s-%05d", date, random)
}

// LoanTerm generates realistic loan terms (in months)
// Weighted: 12 (15%), 24 (20%), 36 (30%), 48 (20%), 60 (15%)
func LoanTerm() int {
	roll := rand.Float64() * 100

	if roll < 15 {
		return 12
	} else if roll < 35 {
		return 24
	} else if roll < 65 {
		return 36
	} else if roll < 85 {
		return 48
	} else {
		return 60
	}
}

// AnnualIncome generates realistic annual income using log-normal distribution
// Median: $65,000, Range: $25,000 - $500,000
func AnnualIncome() float64 {
	mu := 11.08       // ln(65000)
	sigma := 0.6

	normal := rand.NormFloat64()
	lnValue := mu + sigma*normal
	income := math.Exp(lnValue)

	// Clamp to realistic range
	if income < 25000 {
		income = 25000
	}
	if income > 500000 {
		income = 500000
	}

	// Round to nearest $1,000
	return math.Round(income/1000) * 1000
}

// EmploymentStatus generates realistic employment status
// Employed: 75%, Self-Employed: 15%, Unemployed: 5%, Retired: 5%
func EmploymentStatus() string {
	roll := rand.Float64() * 100

	if roll < 75 {
		return "employed"
	} else if roll < 90 {
		return "self_employed"
	} else if roll < 95 {
		return "unemployed"
	} else {
		return "retired"
	}
}

// PaymentStatus generates realistic payment status
// Paid: 85%, Pending: 10%, Late: 4%, Missed: 1%
func PaymentStatus() string {
	roll := rand.Float64() * 100

	if roll < 85 {
		return "paid"
	} else if roll < 95 {
		return "pending"
	} else if roll < 99 {
		return "late"
	} else {
		return "missed"
	}
}

// CreditBureau returns a credit bureau name
// Experian, Equifax, TransUnion (equal distribution)
func CreditBureau() string {
	bureaus := []string{"Experian", "Equifax", "TransUnion"}
	return bureaus[rand.Intn(len(bureaus))]
}
```

#### Unit Tests with Distribution Validation: `pkg/generators/fintech_test.go`

```go
package generators

import (
	"math"
	"testing"
)

func TestLoanAmountRange(t *testing.T) {
	// Generate 1000 loan amounts and verify they're in range
	for i := 0; i < 1000; i++ {
		amount := LoanAmount()
		if amount < 5000 || amount > 500000 {
			t.Errorf("LoanAmount out of range: %.2f", amount)
		}
	}
}

func TestLoanAmountDistribution(t *testing.T) {
	// Generate many samples and verify log-normal distribution
	samples := 10000
	counts := make(map[string]int)

	buckets := []struct {
		name string
		min  float64
		max  float64
	}{
		{"5K-20K", 5000, 20000},
		{"20K-50K", 20000, 50000},
		{"50K-100K", 50000, 100000},
		{"100K-200K", 100000, 200000},
		{"200K-500K", 200000, 500000},
	}

	for i := 0; i < samples; i++ {
		amount := LoanAmount()
		for _, bucket := range buckets {
			if amount >= bucket.min && amount < bucket.max {
				counts[bucket.name]++
				break
			}
		}
	}

	// Verify distribution (log-normal should have more in middle buckets)
	// Most should be in 20K-100K range
	middleBuckets := counts["20K-50K"] + counts["50K-100K"]
	if float64(middleBuckets)/float64(samples) < 0.5 {
		t.Errorf("Expected majority in middle buckets, got distribution: %v", counts)
	}
}

func TestCreditScoreRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		score := CreditScore()
		if score < 300 || score > 850 {
			t.Errorf("CreditScore out of range: %d", score)
		}
	}
}

func TestCreditScoreDistribution(t *testing.T) {
	samples := 10000
	sum := 0

	for i := 0; i < samples; i++ {
		sum += CreditScore()
	}

	mean := float64(sum) / float64(samples)

	// Mean should be close to 680 (±20)
	if mean < 660 || mean > 700 {
		t.Errorf("CreditScore mean should be ~680, got %.2f", mean)
	}
}

func TestInterestRateDistribution(t *testing.T) {
	samples := 10000
	ranges := map[string]int{
		"3-6":   0,
		"6-10":  0,
		"10-15": 0,
	}

	for i := 0; i < samples; i++ {
		rate := InterestRate()

		if rate >= 3.0 && rate < 6.0 {
			ranges["3-6"]++
		} else if rate >= 6.0 && rate < 10.0 {
			ranges["6-10"]++
		} else if rate >= 10.0 && rate <= 15.0 {
			ranges["10-15"]++
		} else {
			t.Errorf("InterestRate out of range: %.2f", rate)
		}
	}

	// Verify weighted distribution (with tolerance)
	// Expected: 3-6%=60%, 6-10%=30%, 10-15%=10%
	pct36 := float64(ranges["3-6"]) / float64(samples) * 100
	pct610 := float64(ranges["6-10"]) / float64(samples) * 100
	pct1015 := float64(ranges["10-15"]) / float64(samples) * 100

	if pct36 < 55 || pct36 > 65 {
		t.Errorf("Expected 60%% in 3-6%% range, got %.1f%%", pct36)
	}
	if pct610 < 25 || pct610 > 35 {
		t.Errorf("Expected 30%% in 6-10%% range, got %.1f%%", pct610)
	}
	if pct1015 < 5 || pct1015 > 15 {
		t.Errorf("Expected 10%% in 10-15%% range, got %.1f%%", pct1015)
	}
}

func TestLoanStatusDistribution(t *testing.T) {
	samples := 10000
	counts := map[string]int{
		"active":     0,
		"paid_off":   0,
		"delinquent": 0,
	}

	for i := 0; i < samples; i++ {
		status := LoanStatus()
		counts[status]++
	}

	// Verify weighted distribution
	// Expected: active=70%, paid_off=25%, delinquent=5%
	pctActive := float64(counts["active"]) / float64(samples) * 100
	pctPaid := float64(counts["paid_off"]) / float64(samples) * 100
	pctDelinquent := float64(counts["delinquent"]) / float64(samples) * 100

	if pctActive < 65 || pctActive > 75 {
		t.Errorf("Expected 70%% active, got %.1f%%", pctActive)
	}
	if pctPaid < 20 || pctPaid > 30 {
		t.Errorf("Expected 25%% paid_off, got %.1f%%", pctPaid)
	}
	if pctDelinquent < 2 || pctDelinquent > 8 {
		t.Errorf("Expected 5%% delinquent, got %.1f%%", pctDelinquent)
	}
}

func TestLoanNumberFormat(t *testing.T) {
	loanNum := LoanNumber()

	// Should match format LN-YYYYMMDD-XXXXX
	if len(loanNum) != 19 {
		t.Errorf("LoanNumber should be 19 characters, got %d: %s", len(loanNum), loanNum)
	}

	if loanNum[0:3] != "LN-" {
		t.Errorf("LoanNumber should start with 'LN-', got: %s", loanNum)
	}
}

func TestLoanTermValues(t *testing.T) {
	validTerms := map[int]bool{
		12: true,
		24: true,
		36: true,
		48: true,
		60: true,
	}

	for i := 0; i < 100; i++ {
		term := LoanTerm()
		if !validTerms[term] {
			t.Errorf("LoanTerm returned invalid term: %d", term)
		}
	}
}

// Benchmark tests
func BenchmarkLoanAmount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoanAmount()
	}
}

func BenchmarkCreditScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditScore()
	}
}

func BenchmarkInterestRate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterestRate()
	}
}
```

### Performance Considerations
- All generators are stateless (fast)
- Math operations are CPU-bound (no I/O)
- Target: 100+ generations/second
- Benchmarks verify performance

### Testing Strategy
- Range tests: Verify values stay in bounds
- Distribution tests: Verify statistical properties
- Format tests: Verify output format (loan numbers, etc.)
- Benchmark tests: Verify performance targets

## Dependencies
- **Upstream**: F011 (Base generators for building blocks)
- **Downstream**: F013 (Data Generation Engine uses these generators)

## Deliverables
1. Custom fintech generator module (`pkg/generators/fintech.go`)
2. Comprehensive unit tests (`pkg/generators/fintech_test.go`)
3. Distribution validation tests
4. Benchmark tests
5. Test coverage report (>90%)

## Success Criteria
- ✅ All custom generators implemented
- ✅ Distributions match specifications (chi-square tests pass)
- ✅ Tests pass with > 90% coverage
- ✅ Benchmarks show > 100 generations/sec
- ✅ Data is visibly more realistic than Faker

## Anti-Patterns to Avoid
- ❌ Uniform random (unrealistic)
- ❌ Hard-coded values
- ❌ Generic distributions (must match real-world patterns)
- ❌ Poor test coverage
- ❌ No distribution validation

## Implementation Notes
- Log-normal distribution: `math.Exp(mu + sigma*rand.NormFloat64())`
- Normal distribution: `mean + stdDev*rand.NormFloat64()`
- Weighted values: Random roll (0-100) with thresholds
- Round monetary values to reasonable precision

## TDD Requirements
**REQUIRED** - Custom generators are core functionality:

1. Write distribution test (e.g., credit scores should cluster around 680)
2. Run test - should fail
3. Implement generator with correct distribution
4. Run test - should pass
5. Verify with histogram (optional visual check)
6. Refactor if needed

## Related Constitution Sections
- **Verticalized > Generic (Principle I, NON-NEGOTIABLE)**: Fintech must look like fintech
- **Realism Standard**: Real-world distributions, not uniform random
- **Quality Bar**: 10x more realistic than Faker
- **TDD Required (Development Practice 1)**: Core data generation logic
- **Performance (Technical Constraint 1)**: 100+ records/sec
