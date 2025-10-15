package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
// This is a stub implementation that will be filled with validation logic
// in Phase 4 (User Stories 2-6) to validate:
// - Schema-level constraints (unique name, valid database_type, etc.)
// - Table-level constraints (unique names, record counts, primary keys)
// - Column-level constraints (valid types, generators, foreign keys)
// - Relationship integrity (foreign keys reference existing tables/columns)
// - Generation order (parent tables before children, no circular dependencies)
//
// For now, returns nil (no validation errors).
func ValidateSchema(s *Schema) error {
	return nil
}
