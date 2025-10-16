package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// ParseSchema parses a schema from an io.Reader.
// Returns the parsed Schema or an error if parsing fails.
// Uses strict parsing to catch unknown fields in the JSON.
func ParseSchema(r io.Reader) (*Schema, error) {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	var schema Schema
	if err := decoder.Decode(&schema); err != nil {
		return nil, fmt.Errorf("ParseSchema: failed to decode JSON: %w", err)
	}

	// Validate the schema after parsing
	if err := ValidateSchema(&schema); err != nil {
		return nil, fmt.Errorf("ParseSchema: %w", err)
	}

	return &schema, nil
}

// LoadSchema loads and parses a schema from a file path.
// Returns the parsed Schema or an error if loading or parsing fails.
func LoadSchema(path string) (*Schema, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("LoadSchema: failed to open file %q: %w", path, err)
	}
	defer f.Close()

	schema, err := ParseSchema(f)
	if err != nil {
		return nil, fmt.Errorf("LoadSchema: failed to parse schema from %q: %w", path, err)
	}

	return schema, nil
}

// ValidateSchema validates a schema's structure and semantic rules.
// Implements multi-phase fail-fast validation for User Stories 2-6:
// - User Story 2: Detect missing required fields
// - User Story 3: Validate foreign key references
// - User Story 4: Validate generation order
// - User Story 5: Detect duplicate names
// - User Story 6: Validate data type consistency
//
// Returns the first validation error encountered, or nil if valid.
func ValidateSchema(s *Schema) error {
	// User Story 2: Detect Missing Required Fields

	// T030: Check schema name is required
	if s.Name == "" {
		return fmt.Errorf("schema name is required")
	}

	// T032: Check database_type is required and valid
	if len(s.DatabaseType) == 0 {
		return fmt.Errorf("database_type is required")
	}

	// Validate each database type is either "mysql" or "postgres"
	for _, dbType := range s.DatabaseType {
		if dbType != "mysql" && dbType != "postgres" {
			return fmt.Errorf("invalid database_type %q: must be \"mysql\" or \"postgres\"", dbType)
		}
	}

	// T031: Check tables field is present (not nil)
	// Note: Empty tables array is allowed for minimal schemas
	if s.Tables == nil {
		return fmt.Errorf("tables field is required")
	}

	// T033: Check generation_order field is present (not nil)
	// Note: Empty generation_order array is allowed for minimal schemas
	if s.GenerationOrder == nil {
		return fmt.Errorf("generation_order is required")
	}

	// T038: Integrate table and column validation
	// T045: Build tableNames map for downstream validation (User Story 3)
	tableNames := make(map[string]bool)

	for i, table := range s.Tables {
		// Validate each table
		if err := ValidateTable(&table, i); err != nil {
			return err
		}

		// Track table names for foreign key validation
		tableNames[table.Name] = true
	}

	// T048: User Story 3: Validate Foreign Key Integrity
	if err := ValidateForeignKeys(s.Tables, tableNames); err != nil {
		return err
	}

	return nil
}

// ValidateTable validates a single table's structure and constraints.
// Returns the first validation error encountered, or nil if valid.
func ValidateTable(t *Table, tableIndex int) error {
	// T034: Check table name is required
	if t.Name == "" {
		return fmt.Errorf("table %d: table name is required", tableIndex)
	}

	// T037: Check record_count is positive (> 0)
	if t.RecordCount <= 0 {
		return fmt.Errorf("table %d (%s): record_count must be greater than 0", tableIndex, t.Name)
	}

	// T036: Exactly one primary key per table (checked before empty columns)
	// This catches both empty columns and columns without primary key
	pkCount := 0
	for _, col := range t.Columns {
		if col.PrimaryKey {
			pkCount++
		}
	}

	if pkCount == 0 {
		return fmt.Errorf("table %d (%s): must have exactly one primary key", tableIndex, t.Name)
	}

	if pkCount > 1 {
		return fmt.Errorf("table %d (%s): must have exactly one primary key, found %d", tableIndex, t.Name, pkCount)
	}

	// T034: Check columns array is non-empty (after primary key check)
	if len(t.Columns) == 0 {
		return fmt.Errorf("table %d (%s): columns are required", tableIndex, t.Name)
	}

	// T038: Validate each column
	for j, col := range t.Columns {
		if err := ValidateColumn(&col, tableIndex, t.Name, j); err != nil {
			return err
		}
	}

	return nil
}

// ValidateColumn validates a single column's structure and constraints.
// Returns the first validation error encountered, or nil if valid.
func ValidateColumn(c *Column, tableIndex int, tableName string, colIndex int) error {
	// T035: Check column name is required
	if c.Name == "" {
		return fmt.Errorf("table %d (%s): column %d: column name is required", tableIndex, tableName, colIndex)
	}

	// T035: Check column type is required
	if c.Type == "" {
		return fmt.Errorf("table %d (%s): column %d (%s): column type is required", tableIndex, tableName, colIndex, c.Name)
	}

	return nil
}

// ValidateForeignKeys validates all foreign key references in the schema.
// T046: Checks that foreign keys reference tables that exist in tableNames map.
// Returns the first validation error encountered, or nil if all foreign keys are valid.
func ValidateForeignKeys(tables []Table, tableNames map[string]bool) error {
	for _, table := range tables {
		for _, col := range table.Columns {
			// Skip columns without foreign keys
			if col.ForeignKey == nil {
				continue
			}

			fk := col.ForeignKey

			// T046: Check that referenced table exists
			if !tableNames[fk.Table] {
				return fmt.Errorf("table '%s': column '%s': foreign key references table '%s' which does not exist in schema", table.Name, col.Name, fk.Table)
			}

			// T047: Validate on_delete action
			if err := ValidateReferentialAction(fk.OnDelete, "on_delete", table.Name, col.Name); err != nil {
				return err
			}

			// T047: Validate on_update action
			if err := ValidateReferentialAction(fk.OnUpdate, "on_update", table.Name, col.Name); err != nil {
				return err
			}
		}
	}

	return nil
}

// ValidateReferentialAction validates a foreign key referential action.
// T047: Checks that action is one of: CASCADE, SET NULL, RESTRICT (case-sensitive).
// Returns an error if the action is invalid, or nil if valid.
func ValidateReferentialAction(action string, actionType string, tableName string, colName string) error {
	// Valid actions according to SQL standard and schema spec (F007)
	validActions := []string{"CASCADE", "SET NULL", "RESTRICT"}

	// Normalize to uppercase for case-insensitive comparison
	normalizedAction := strings.ToUpper(action)

	// Check if action is valid
	for _, valid := range validActions {
		if normalizedAction == valid {
			return nil
		}
	}

	// T049: Return error with full context
	return fmt.Errorf("table '%s': column '%s': invalid %s action '%s': must be one of: %s",
		tableName, colName, actionType, action, strings.Join(validActions, ", "))
}
