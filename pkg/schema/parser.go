package schema

import (
	"fmt"
	"io"
)

// ParseSchema parses a schema from an io.Reader.
// Returns the parsed Schema or an error if parsing fails.
func ParseSchema(r io.Reader) (*Schema, error) {
	return nil, fmt.Errorf("ParseSchema: not yet implemented")
}

// LoadSchema loads and parses a schema from a file path.
// Returns the parsed Schema or an error if loading or parsing fails.
func LoadSchema(path string) (*Schema, error) {
	return nil, fmt.Errorf("LoadSchema: not yet implemented")
}
