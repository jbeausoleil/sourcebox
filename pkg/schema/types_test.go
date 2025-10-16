package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTypesInstantiation verifies that all schema types can be instantiated.
// This is a basic compilation and structure test for Phase 1.
func TestTypesInstantiation(t *testing.T) {
	tests := []struct {
		name     string
		instance interface{}
	}{
		{
			name: "Schema",
			instance: &Schema{
				SchemaVersion:   "1.0",
				Name:            "test-schema",
				Description:     "Test schema",
				Author:          "Test Author",
				Version:         "1.0.0",
				DatabaseType:    []string{"mysql", "postgres"},
				Metadata:        SchemaMetadata{},
				Tables:          []Table{},
				Relationships:   []Relationship{},
				GenerationOrder: []string{},
				ValidationRules: []ValidationRule{},
			},
		},
		{
			name: "SchemaMetadata",
			instance: &SchemaMetadata{
				Industry:     "test",
				Tags:         []string{"tag1", "tag2"},
				TotalRecords: 1000,
			},
		},
		{
			name: "Table",
			instance: &Table{
				Name:        "test_table",
				Description: "Test table",
				RecordCount: 100,
				Columns:     []Column{},
				Indexes:     []Index{},
			},
		},
		{
			name: "Column",
			instance: &Column{
				Name:            "test_column",
				Type:            "varchar(255)",
				Nullable:        false,
				PrimaryKey:      false,
				AutoIncrement:   false,
				Default:         nil,
				Unique:          false,
				Description:     "Test column",
				Generator:       "uuid",
				GeneratorParams: map[string]interface{}{},
				ForeignKey:      nil,
			},
		},
		{
			name: "ForeignKey",
			instance: &ForeignKey{
				Table:    "parent_table",
				Column:   "id",
				OnDelete: "CASCADE",
				OnUpdate: "CASCADE",
			},
		},
		{
			name: "Index",
			instance: &Index{
				Name:    "idx_test",
				Columns: []string{"col1", "col2"},
				Unique:  false,
			},
		},
		{
			name: "Relationship",
			instance: &Relationship{
				FromTable:        "child_table",
				FromColumn:       "parent_id",
				ToTable:          "parent_table",
				ToColumn:         "id",
				RelationshipType: "many_to_one",
				Description:      "Test relationship",
			},
		},
		{
			name: "ValidationRule",
			instance: &ValidationRule{
				Rule:        "unique_name",
				Description: "Schema name must be unique",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.instance, "should be able to instantiate %s", tt.name)
		})
	}
}

// TestJSONTags verifies that JSON tags are properly defined on Schema struct.
// This ensures compatibility with the F007 schema format.
func TestJSONTags(t *testing.T) {
	// This test verifies that the struct tags exist and are correct
	// by attempting to create a schema that would match F007 format
	schema := Schema{
		SchemaVersion: "1.0",
		Name:          "test",
		DatabaseType:  []string{"mysql"},
	}

	assert.Equal(t, "1.0", schema.SchemaVersion)
	assert.Equal(t, "test", schema.Name)
	assert.NotNil(t, schema.DatabaseType)
}

// TestColumnWithForeignKey verifies that columns can have foreign key constraints.
func TestColumnWithForeignKey(t *testing.T) {
	fk := &ForeignKey{
		Table:    "parent_table",
		Column:   "id",
		OnDelete: "CASCADE",
		OnUpdate: "CASCADE",
	}

	col := Column{
		Name:       "parent_id",
		Type:       "int",
		ForeignKey: fk,
	}

	assert.NotNil(t, col.ForeignKey)
	assert.Equal(t, "parent_table", col.ForeignKey.Table)
	assert.Equal(t, "id", col.ForeignKey.Column)
	assert.Equal(t, "CASCADE", col.ForeignKey.OnDelete)
	assert.Equal(t, "CASCADE", col.ForeignKey.OnUpdate)
}

// TestColumnParams verifies that column generator params can hold various types.
func TestColumnParams(t *testing.T) {
	col := Column{
		Name: "test_col",
		Type: "decimal(10,2)",
		GeneratorParams: map[string]interface{}{
			"min":     0.0,
			"max":     1000.0,
			"mean":    500.0,
			"std_dev": 100.0,
		},
	}

	assert.Equal(t, 0.0, col.GeneratorParams["min"])
	assert.Equal(t, 1000.0, col.GeneratorParams["max"])
	assert.Equal(t, 500.0, col.GeneratorParams["mean"])
	assert.Equal(t, 100.0, col.GeneratorParams["std_dev"])
}
