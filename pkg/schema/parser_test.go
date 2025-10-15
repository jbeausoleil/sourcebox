package schema

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSchema_ValidJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		validate func(t *testing.T, schema *Schema)
	}{
		{
			name: "minimal valid schema",
			input: `{
				"schema_version": "1.0",
				"name": "test-schema",
				"description": "Test schema",
				"author": "Test Author",
				"version": "1.0.0",
				"database_type": ["mysql"],
				"tables": [],
				"generation_order": []
			}`,
			validate: func(t *testing.T, schema *Schema) {
				assert.Equal(t, "1.0", schema.SchemaVersion)
				assert.Equal(t, "test-schema", schema.Name)
				assert.Equal(t, "Test schema", schema.Description)
				assert.Equal(t, "Test Author", schema.Author)
				assert.Equal(t, "1.0.0", schema.Version)
				assert.Equal(t, []string{"mysql"}, schema.DatabaseType)
				assert.Empty(t, schema.Tables)
				assert.Empty(t, schema.GenerationOrder)
			},
		},
		{
			name: "schema with metadata",
			input: `{
				"schema_version": "1.0",
				"name": "test-schema",
				"description": "Test schema",
				"author": "Test Author",
				"version": "1.0.0",
				"database_type": ["mysql", "postgres"],
				"metadata": {
					"industry": "fintech",
					"tags": ["loans", "credit"],
					"total_records": 1000,
					"complexity_tier": 1
				},
				"tables": [],
				"generation_order": []
			}`,
			validate: func(t *testing.T, schema *Schema) {
				assert.Equal(t, "fintech", schema.Metadata.Industry)
				assert.Equal(t, []string{"loans", "credit"}, schema.Metadata.Tags)
				assert.Equal(t, 1000, schema.Metadata.TotalRecords)
				assert.Equal(t, 1, schema.Metadata.ComplexityTier)
			},
		},
		{
			name: "schema with table",
			input: `{
				"schema_version": "1.0",
				"name": "test-schema",
				"description": "Test schema",
				"author": "Test Author",
				"version": "1.0.0",
				"database_type": ["mysql"],
				"tables": [
					{
						"name": "users",
						"description": "User table",
						"record_count": 100,
						"columns": [
							{
								"name": "id",
								"type": "int",
								"nullable": false,
								"primary_key": true
							},
							{
								"name": "email",
								"type": "varchar(255)",
								"nullable": false,
								"unique": true,
								"generator": "email"
							}
						]
					}
				],
				"generation_order": ["users"]
			}`,
			validate: func(t *testing.T, schema *Schema) {
				require.Len(t, schema.Tables, 1)
				table := schema.Tables[0]
				assert.Equal(t, "users", table.Name)
				assert.Equal(t, "User table", table.Description)
				assert.Equal(t, 100, table.RecordCount)
				require.Len(t, table.Columns, 2)

				// Verify first column
				col1 := table.Columns[0]
				assert.Equal(t, "id", col1.Name)
				assert.Equal(t, "int", col1.Type)
				assert.False(t, col1.Nullable)
				assert.True(t, col1.PrimaryKey)

				// Verify second column
				col2 := table.Columns[1]
				assert.Equal(t, "email", col2.Name)
				assert.Equal(t, "varchar(255)", col2.Type)
				assert.False(t, col2.Nullable)
				assert.True(t, col2.Unique)
				assert.Equal(t, "email", col2.Generator)
			},
		},
		{
			name: "schema with foreign key",
			input: `{
				"schema_version": "1.0",
				"name": "test-schema",
				"description": "Test schema",
				"author": "Test Author",
				"version": "1.0.0",
				"database_type": ["mysql"],
				"tables": [
					{
						"name": "posts",
						"record_count": 50,
						"columns": [
							{
								"name": "id",
								"type": "int",
								"primary_key": true
							},
							{
								"name": "user_id",
								"type": "int",
								"foreign_key": {
									"table": "users",
									"column": "id",
									"on_delete": "CASCADE",
									"on_update": "CASCADE"
								}
							}
						]
					}
				],
				"generation_order": ["posts"]
			}`,
			validate: func(t *testing.T, schema *Schema) {
				require.Len(t, schema.Tables, 1)
				table := schema.Tables[0]
				require.Len(t, table.Columns, 2)

				col := table.Columns[1]
				assert.Equal(t, "user_id", col.Name)
				require.NotNil(t, col.ForeignKey)
				assert.Equal(t, "users", col.ForeignKey.Table)
				assert.Equal(t, "id", col.ForeignKey.Column)
				assert.Equal(t, "CASCADE", col.ForeignKey.OnDelete)
				assert.Equal(t, "CASCADE", col.ForeignKey.OnUpdate)
			},
		},
		{
			name: "schema with relationships",
			input: `{
				"schema_version": "1.0",
				"name": "test-schema",
				"description": "Test schema",
				"author": "Test Author",
				"version": "1.0.0",
				"database_type": ["mysql"],
				"tables": [],
				"relationships": [
					{
						"from_table": "posts",
						"from_column": "user_id",
						"to_table": "users",
						"to_column": "id",
						"relationship_type": "many_to_one",
						"description": "Each post belongs to one user"
					}
				],
				"generation_order": []
			}`,
			validate: func(t *testing.T, schema *Schema) {
				require.Len(t, schema.Relationships, 1)
				rel := schema.Relationships[0]
				assert.Equal(t, "posts", rel.FromTable)
				assert.Equal(t, "user_id", rel.FromColumn)
				assert.Equal(t, "users", rel.ToTable)
				assert.Equal(t, "id", rel.ToColumn)
				assert.Equal(t, "many_to_one", rel.RelationshipType)
				assert.Equal(t, "Each post belongs to one user", rel.Description)
			},
		},
		{
			name: "schema with validation rules",
			input: `{
				"schema_version": "1.0",
				"name": "test-schema",
				"description": "Test schema",
				"author": "Test Author",
				"version": "1.0.0",
				"database_type": ["mysql"],
				"tables": [],
				"validation_rules": [
					{
						"rule": "users.email UNIQUE",
						"description": "Email must be unique",
						"severity": "error"
					}
				],
				"generation_order": []
			}`,
			validate: func(t *testing.T, schema *Schema) {
				require.Len(t, schema.ValidationRules, 1)
				rule := schema.ValidationRules[0]
				assert.Equal(t, "users.email UNIQUE", rule.Rule)
				assert.Equal(t, "Email must be unique", rule.Description)
				assert.Equal(t, "error", rule.Severity)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			schema, err := ParseSchema(reader)
			require.NoError(t, err)
			require.NotNil(t, schema)
			tt.validate(t, schema)
		})
	}
}

func TestParseSchema_InvalidJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr string
	}{
		{
			name:        "malformed JSON",
			input:       `{invalid json}`,
			expectedErr: "failed to decode JSON",
		},
		{
			name:        "empty input",
			input:       ``,
			expectedErr: "failed to decode JSON",
		},
		{
			name: "unknown field",
			input: `{
				"schema_version": "1.0",
				"name": "test",
				"unknown_field": "value",
				"database_type": ["mysql"],
				"tables": [],
				"generation_order": []
			}`,
			expectedErr: "unknown field",
		},
		{
			name: "wrong type for field",
			input: `{
				"schema_version": "1.0",
				"name": "test",
				"database_type": "mysql",
				"tables": [],
				"generation_order": []
			}`,
			expectedErr: "failed to decode JSON",
		},
		{
			name:        "incomplete JSON",
			input:       `{"schema_version": "1.0", "name": "test"`,
			expectedErr: "failed to decode JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			schema, err := ParseSchema(reader)
			require.Error(t, err)
			assert.Nil(t, schema)
			assert.Contains(t, err.Error(), tt.expectedErr)
			assert.Contains(t, err.Error(), "ParseSchema")
		})
	}
}

func TestParseSchema_ComplexGeneratorParams(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "data",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					},
					{
						"name": "score",
						"type": "int",
						"generator": "int_range",
						"generator_params": {
							"min": 0,
							"max": 100
						}
					}
				]
			}
		],
		"generation_order": ["data"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)
	require.NoError(t, err)
	require.NotNil(t, schema)

	require.Len(t, schema.Tables, 1)
	table := schema.Tables[0]
	require.Len(t, table.Columns, 2)

	col := table.Columns[1]
	assert.Equal(t, "score", col.Name)
	require.NotNil(t, col.GeneratorParams)

	// Verify generator params are parsed as map
	assert.Contains(t, col.GeneratorParams, "min")
	assert.Contains(t, col.GeneratorParams, "max")
}

func TestLoadSchema_ValidFile(t *testing.T) {
	// Use the actual example schema file
	schemaPath := filepath.Join("..", "..", "schemas", "example-schema.json")

	schema, err := LoadSchema(schemaPath)
	require.NoError(t, err)
	require.NotNil(t, schema)

	// Validate basic structure
	assert.Equal(t, "1.0", schema.SchemaVersion)
	assert.Equal(t, "fintech-loans", schema.Name)
	assert.Equal(t, "Realistic fintech loan data with borrowers, loans, and payments", schema.Description)
	assert.Equal(t, "SourceBox Contributors", schema.Author)
	assert.Equal(t, "1.0.0", schema.Version)
	assert.Equal(t, []string{"mysql", "postgres"}, schema.DatabaseType)

	// Validate metadata
	assert.Equal(t, "fintech", schema.Metadata.Industry)
	assert.Equal(t, 4950, schema.Metadata.TotalRecords)
	assert.Equal(t, 1, schema.Metadata.ComplexityTier)

	// Validate tables
	require.Len(t, schema.Tables, 3)
	assert.Equal(t, "borrowers", schema.Tables[0].Name)
	assert.Equal(t, "loans", schema.Tables[1].Name)
	assert.Equal(t, "payments", schema.Tables[2].Name)

	// Validate generation order
	assert.Equal(t, []string{"borrowers", "loans", "payments"}, schema.GenerationOrder)

	// Validate relationships
	require.Len(t, schema.Relationships, 2)

	// Validate validation rules
	require.Len(t, schema.ValidationRules, 2)
}

func TestLoadSchema_FileNotFound(t *testing.T) {
	schema, err := LoadSchema("/nonexistent/path/to/schema.json")
	require.Error(t, err)
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "LoadSchema")
	assert.Contains(t, err.Error(), "failed to open file")
}

func TestLoadSchema_InvalidJSON(t *testing.T) {
	// Create temporary file with invalid JSON
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "invalid.json")
	err := os.WriteFile(tmpFile, []byte(`{invalid json}`), 0644)
	require.NoError(t, err)

	schema, err := LoadSchema(tmpFile)
	require.Error(t, err)
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "LoadSchema")
	assert.Contains(t, err.Error(), "failed to parse schema")
}

func TestLoadSchema_EmptyFile(t *testing.T) {
	// Create temporary empty file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "empty.json")
	err := os.WriteFile(tmpFile, []byte(``), 0644)
	require.NoError(t, err)

	schema, err := LoadSchema(tmpFile)
	require.Error(t, err)
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "LoadSchema")
}

func TestValidateSchema_Stub(t *testing.T) {
	// Test that ValidateSchema stub returns nil
	schema := &Schema{
		SchemaVersion: "1.0",
		Name:          "test",
	}

	err := ValidateSchema(schema)
	assert.NoError(t, err, "ValidateSchema stub should return nil")
}

func TestValidateSchema_NilSchema(t *testing.T) {
	// Test that ValidateSchema handles nil input gracefully
	// (stub implementation doesn't validate, so it should just return nil)
	err := ValidateSchema(nil)
	assert.NoError(t, err, "ValidateSchema stub should return nil even for nil input")
}
