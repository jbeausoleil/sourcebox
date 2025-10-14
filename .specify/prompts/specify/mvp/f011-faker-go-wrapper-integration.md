# Feature Specification Prompt: F011 - Faker.js Go Wrapper Integration

## Feature Metadata
- **Feature ID**: F011
- **Name**: Faker.js Go Wrapper Integration
- **Category**: Data Generation
- **Phase**: Week 5
- **Priority**: P0 (Must-have)
- **Estimated Effort**: Small (2 days)
- **Dependencies**: F009 (gofakeit dependency installed)

## Constitutional Alignment

### Core Principles
- ✅ **Boring Tech Wins**: Use proven library (gofakeit), don't reinvent
- ✅ **Speed > Features**: Fast data generation (100+ records/sec)
- ✅ **Developer-First**: Simple API, easy to use

### Technical Constraints
- ✅ **Performance**: Fast generation rates
- ✅ **Code Quality**: TDD required for core generators

### Development Practices
- ✅ **TDD Required**: Base generators are core functionality

## User Story
**US-MVP-006**: "As a data generation engine, I need reliable base generators for common data types (names, emails, dates) so I can build realistic datasets quickly."

## Problem Statement
SourceBox needs basic data generators for common fields like names, emails, addresses, dates, and numbers. Building these from scratch is time-consuming and error-prone. The gofakeit library provides battle-tested generators, but needs to be wrapped in a clean API that:
- Provides consistent interface for all generators
- Handles edge cases (nil values, empty strings)
- Is easily testable
- Supports future custom generators (F012)

## Solution Overview
Create a wrapper module (`pkg/generators/base.go`) that integrates gofakeit library and provides clean, tested generator functions. Implement generators for personal data (names, emails, phones), company data, dates/times, and numeric values. Write comprehensive unit tests to ensure reliability.

## Detailed Requirements

### Acceptance Criteria
1. **Base Generator Module**: `pkg/generators/base.go` created
2. **Personal Data Generators**: Name(), Email(), Address(), Phone(), DateOfBirth()
3. **Company Data Generators**: Company(), JobTitle(), Domain()
4. **Date/Time Generators**: Date(), Timestamp(), DateBetween(start, end), TimestampPast(daysAgo)
5. **Numeric Generators**: Int(), IntRange(min, max), Float(), FloatRange(min, max), Decimal(precision, scale)
6. **All Generators Return Realistic Data**: No "test@test.com" or "John Doe"
7. **Unit Tests**: `go test ./pkg/generators/base_test.go` passes with >90% coverage

### Technical Specifications

#### Base Generators: `pkg/generators/base.go`

```go
package generators

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// Initialize gofakeit with seed for reproducibility (optional)
func init() {
	// Use deterministic seed for testing, random for production
	// gofakeit.Seed(time.Now().UnixNano())
}

// Personal Data Generators

// FirstName generates a realistic first name
func FirstName() string {
	return gofakeit.FirstName()
}

// LastName generates a realistic last name
func LastName() string {
	return gofakeit.LastName()
}

// FullName generates a realistic full name
func FullName() string {
	return gofakeit.Name()
}

// Email generates a realistic email address
func Email() string {
	return gofakeit.Email()
}

// Phone generates a realistic phone number
func Phone() string {
	return gofakeit.Phone()
}

// Address generates a realistic street address
func Address() string {
	return gofakeit.Address().Address
}

// City generates a realistic city name
func City() string {
	return gofakeit.City()
}

// State generates a realistic state
func State() string {
	return gofakeit.State()
}

// ZipCode generates a realistic ZIP code
func ZipCode() string {
	return gofakeit.Zip()
}

// SSNLast4 generates last 4 digits of SSN
func SSNLast4() string {
	ssn := gofakeit.SSN()
	if len(ssn) >= 4 {
		return ssn[len(ssn)-4:]
	}
	return ssn
}

// DateOfBirth generates a date of birth for given age range
func DateOfBirth(minAge, maxAge int) time.Time {
	yearsAgo := gofakeit.Number(minAge, maxAge)
	return time.Now().AddDate(-yearsAgo, 0, 0)
}

// Company Data Generators

// Company generates a realistic company name
func Company() string {
	return gofakeit.Company()
}

// JobTitle generates a realistic job title
func JobTitle() string {
	return gofakeit.JobTitle()
}

// Domain generates a realistic domain name
func Domain() string {
	return gofakeit.DomainName()
}

// CompanyEmail generates a company email based on person's name
func CompanyEmail(firstName, lastName, company string) string {
	domain := gofakeit.DomainName()
	return fmt.Sprintf("%s.%s@%s",
		gofakeit.FirstName(),
		gofakeit.LastName(),
		domain)
}

// Date and Time Generators

// Date generates a random date
func Date() time.Time {
	return gofakeit.Date()
}

// Timestamp generates a random timestamp
func Timestamp() time.Time {
	return gofakeit.Date()
}

// DateBetween generates a date between start and end
func DateBetween(start, end time.Time) time.Time {
	return gofakeit.DateRange(start, end)
}

// TimestampPast generates a timestamp N days in the past
func TimestampPast(daysAgo int) time.Time {
	daysAgoMin := 1
	daysAgoMax := daysAgo
	randomDays := gofakeit.Number(daysAgoMin, daysAgoMax)
	return time.Now().AddDate(0, 0, -randomDays)
}

// TimestampFuture generates a timestamp N days in the future
func TimestampFuture(daysAhead int) time.Time {
	randomDays := gofakeit.Number(1, daysAhead)
	return time.Now().AddDate(0, 0, randomDays)
}

// Numeric Generators

// Int generates a random integer
func Int() int {
	return gofakeit.Number(0, 1000000)
}

// IntRange generates an integer between min and max (inclusive)
func IntRange(min, max int) int {
	return gofakeit.Number(min, max)
}

// Float generates a random float
func Float() float64 {
	return gofakeit.Float64()
}

// FloatRange generates a float between min and max
func FloatRange(min, max float64) float64 {
	return gofakeit.Float64Range(min, max)
}

// Decimal generates a decimal with given precision and scale
func Decimal(precision, scale int) float64 {
	// Generate decimal with specific precision and scale
	// Example: precision=10, scale=2 -> 12345678.90
	maxValue := 1.0
	for i := 0; i < precision-scale; i++ {
		maxValue *= 10
	}
	return FloatRange(0, maxValue)
}

// Boolean generates a random boolean
func Boolean() bool {
	return gofakeit.Bool()
}

// Utility Generators

// UUID generates a UUID v4
func UUID() string {
	return gofakeit.UUID()
}

// Username generates a realistic username
func Username() string {
	return gofakeit.Username()
}

// Password generates a password
func Password(length int) string {
	return gofakeit.Password(true, true, true, true, false, length)
}
```

#### Unit Tests: `pkg/generators/base_test.go`

```go
package generators

import (
	"strings"
	"testing"
	"time"
)

func TestFirstName(t *testing.T) {
	name := FirstName()
	if name == "" {
		t.Error("FirstName should not return empty string")
	}
	if len(name) < 2 {
		t.Errorf("FirstName should be at least 2 characters, got: %s", name)
	}
}

func TestLastName(t *testing.T) {
	name := LastName()
	if name == "" {
		t.Error("LastName should not return empty string")
	}
	if len(name) < 2 {
		t.Errorf("LastName should be at least 2 characters, got: %s", name)
	}
}

func TestFullName(t *testing.T) {
	name := FullName()
	if name == "" {
		t.Error("FullName should not return empty string")
	}
	if !strings.Contains(name, " ") {
		t.Errorf("FullName should contain a space, got: %s", name)
	}
}

func TestEmail(t *testing.T) {
	email := Email()
	if email == "" {
		t.Error("Email should not return empty string")
	}
	if !strings.Contains(email, "@") {
		t.Errorf("Email should contain @, got: %s", email)
	}
	if !strings.Contains(email, ".") {
		t.Errorf("Email should contain a dot, got: %s", email)
	}
}

func TestPhone(t *testing.T) {
	phone := Phone()
	if phone == "" {
		t.Error("Phone should not return empty string")
	}
	// Phone numbers vary in format, just check it's not empty
}

func TestDateOfBirth(t *testing.T) {
	dob := DateOfBirth(21, 75)
	now := time.Now()

	// Should be in the past
	if dob.After(now) {
		t.Error("DateOfBirth should be in the past")
	}

	// Should be at least 21 years ago
	minDate := now.AddDate(-75, 0, 0)
	maxDate := now.AddDate(-21, 0, 0)

	if dob.Before(minDate) || dob.After(maxDate) {
		t.Errorf("DateOfBirth should be between 21-75 years ago, got: %v", dob)
	}
}

func TestCompany(t *testing.T) {
	company := Company()
	if company == "" {
		t.Error("Company should not return empty string")
	}
}

func TestIntRange(t *testing.T) {
	min, max := 10, 20
	for i := 0; i < 100; i++ {
		val := IntRange(min, max)
		if val < min || val > max {
			t.Errorf("IntRange(%d, %d) returned %d, out of range", min, max, val)
		}
	}
}

func TestFloatRange(t *testing.T) {
	min, max := 10.0, 20.0
	for i := 0; i < 100; i++ {
		val := FloatRange(min, max)
		if val < min || val > max {
			t.Errorf("FloatRange(%.2f, %.2f) returned %.2f, out of range", min, max, val)
		}
	}
}

func TestTimestampPast(t *testing.T) {
	ts := TimestampPast(365)
	now := time.Now()

	if ts.After(now) {
		t.Error("TimestampPast should return a date in the past")
	}

	// Should be within 365 days ago
	minDate := now.AddDate(0, 0, -365)
	if ts.Before(minDate) {
		t.Errorf("TimestampPast(365) should be within last 365 days, got: %v", ts)
	}
}

func TestTimestampFuture(t *testing.T) {
	ts := TimestampFuture(365)
	now := time.Now()

	if ts.Before(now) {
		t.Error("TimestampFuture should return a date in the future")
	}

	// Should be within next 365 days
	maxDate := now.AddDate(0, 0, 365)
	if ts.After(maxDate) {
		t.Errorf("TimestampFuture(365) should be within next 365 days, got: %v", ts)
	}
}

func TestBoolean(t *testing.T) {
	// Generate many booleans and ensure we get both true and false
	trueCount := 0
	falseCount := 0

	for i := 0; i < 100; i++ {
		if Boolean() {
			trueCount++
		} else {
			falseCount++
		}
	}

	if trueCount == 0 || falseCount == 0 {
		t.Errorf("Boolean should return both true and false, got %d true and %d false", trueCount, falseCount)
	}
}

// Benchmark tests
func BenchmarkFirstName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstName()
	}
}

func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Email()
	}
}

func BenchmarkIntRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntRange(1, 1000)
	}
}
```

### Performance Considerations
- gofakeit is highly optimized (100K+ generations/sec)
- No external API calls (all local)
- Minimal memory allocation

### Testing Strategy
- Unit tests for all generator functions
- Verify output format (email has @, name is not empty)
- Verify ranges (IntRange, FloatRange stay in bounds)
- Verify dates (past dates are before now, future dates are after)
- Benchmark tests to ensure performance targets (>100 records/sec)

## Dependencies
- **Upstream**: F009 (gofakeit dependency must be installed)
- **Downstream**: F012 (Custom fintech generators build on these base generators)

## Deliverables
1. Base generator module (`pkg/generators/base.go`)
2. Comprehensive unit tests (`pkg/generators/base_test.go`)
3. Benchmark tests for performance validation
4. Test coverage report (>90%)

## Success Criteria
- ✅ All base generators implemented
- ✅ Tests pass: `go test ./pkg/generators/...`
- ✅ Test coverage > 90%
- ✅ Benchmarks show > 100 generations/second
- ✅ All generators return realistic data

## Anti-Patterns to Avoid
- ❌ Reinventing faker functionality (use gofakeit)
- ❌ Returning empty strings or nil values
- ❌ Hard-coded test data (e.g., "test@test.com")
- ❌ Poor test coverage
- ❌ Slow generators (< 100/sec)

## Implementation Notes
- gofakeit provides deterministic output with seeded random
- Wrapper pattern allows us to extend/customize behavior
- Tests ensure generators work before building engine (F013)

## TDD Requirements
**REQUIRED** - Base generators are core functionality:

1. Write test for FirstName generator
2. Run test - should fail (function doesn't exist)
3. Implement FirstName generator
4. Run test - should pass
5. Refactor if needed
6. Repeat for each generator

## Related Constitution Sections
- **Boring Tech Wins (Principle IV)**: Use proven gofakeit library
- **Performance (Technical Constraint 1)**: 100+ records/sec
- **TDD Required (Development Practice 1)**: Core data generation logic
- **Code Quality (Technical Constraint 5)**: >80% coverage
