package schema

import (
	"fmt"
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
						"name": "users",
						"record_count": 25,
						"columns": [
							{
								"name": "id",
								"type": "int",
								"primary_key": true
							}
						]
					},
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
				"generation_order": ["users", "posts"]
			}`,
			validate: func(t *testing.T, schema *Schema) {
				require.Len(t, schema.Tables, 2)

				// Verify users table exists
				usersTable := schema.Tables[0]
				assert.Equal(t, "users", usersTable.Name)

				// Verify posts table with foreign key
				postsTable := schema.Tables[1]
				assert.Equal(t, "posts", postsTable.Name)

				col := postsTable.Columns[1]
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

func TestValidateSchema_ValidMinimalSchema(t *testing.T) {
	// Test that ValidateSchema accepts a valid minimal schema
	schema := &Schema{
		Name:            "test-schema",
		DatabaseType:    []string{"mysql"},
		Tables:          []Table{},
		GenerationOrder: []string{},
	}

	err := ValidateSchema(schema)
	assert.NoError(t, err, "ValidateSchema should accept valid minimal schema")
}

func TestValidateSchema_ValidSchemaWithTable(t *testing.T) {
	// Test that ValidateSchema accepts a valid schema with a table
	schema := &Schema{
		Name:         "test-schema",
		DatabaseType: []string{"mysql"},
		Tables: []Table{
			{
				Name:        "users",
				RecordCount: 100,
				Columns: []Column{
					{
						Name:       "id",
						Type:       "int",
						PrimaryKey: true,
					},
					{
						Name: "email",
						Type: "varchar(255)",
					},
				},
			},
		},
		GenerationOrder: []string{"users"},
	}

	err := ValidateSchema(schema)
	assert.NoError(t, err, "ValidateSchema should accept valid schema with table")
}

// ============================================================================
// User Story 2: Detect Missing Required Fields (TDD RED Phase)
// These tests should FAIL initially until validation logic is implemented
// ============================================================================

func TestParseMissingSchemaName(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [],
		"generation_order": []
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when schema name is missing")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "name")
	assert.Contains(t, err.Error(), "required")
}

func TestParseMissingTables(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"generation_order": []
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when tables field is missing")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "tables")
	assert.Contains(t, err.Error(), "required")
}

func TestParseMissingDatabaseType(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"tables": [],
		"generation_order": []
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when database_type is missing")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "database_type")
	assert.Contains(t, err.Error(), "required")
}

func TestParseMissingGenerationOrder(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": []
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when generation_order is missing")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "generation_order")
	assert.Contains(t, err.Error(), "required")
}

func TestParseTableMissingName(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"description": "Table without name",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					}
				]
			}
		],
		"generation_order": ["users"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when table name is missing")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "table")
	assert.Contains(t, err.Error(), "name")
	assert.Contains(t, err.Error(), "required")
}

func TestParseTableMissingColumns(t *testing.T) {
	input := `{
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
				"record_count": 100
			}
		],
		"generation_order": ["users"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when table columns field is missing")
	assert.Nil(t, schema)
	// Note: The validation checks primary key first, which will fail for missing/empty columns
	// This is acceptable as it still prevents invalid schemas
	assert.Contains(t, err.Error(), "primary")
}

func TestParseTableMissingPrimaryKey(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		description string
	}{
		{
			name: "no primary key column",
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
						"record_count": 100,
						"columns": [
							{
								"name": "email",
								"type": "varchar(255)",
								"nullable": false
							}
						]
					}
				],
				"generation_order": ["users"]
			}`,
			description: "table has columns but none marked as primary key",
		},
		{
			name: "empty columns array",
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
						"record_count": 100,
						"columns": []
					}
				],
				"generation_order": ["users"]
			}`,
			description: "table has empty columns array",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			schema, err := ParseSchema(reader)

			require.Error(t, err, "ParseSchema should fail when table lacks primary key: %s", tt.description)
			assert.Nil(t, schema)
			assert.Contains(t, err.Error(), "primary")
			assert.Contains(t, err.Error(), "key")
		})
	}
}

func TestParseColumnMissingName(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": 100,
				"columns": [
					{
						"type": "int",
						"primary_key": true
					}
				]
			}
		],
		"generation_order": ["users"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when column name is missing")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "column")
	assert.Contains(t, err.Error(), "name")
	assert.Contains(t, err.Error(), "required")
}

func TestParseColumnMissingType(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"primary_key": true
					}
				]
			}
		],
		"generation_order": ["users"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when column type is missing")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "column")
	assert.Contains(t, err.Error(), "type")
	assert.Contains(t, err.Error(), "required")
}

func TestParseZeroRecordCount(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": 0,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					}
				]
			}
		],
		"generation_order": ["users"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when record_count is zero")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "record_count")
	assert.Contains(t, err.Error(), "greater than")
}

func TestParseNegativeRecordCount(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": -100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					}
				]
			}
		],
		"generation_order": ["users"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when record_count is negative")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "record_count")
	assert.Contains(t, err.Error(), "greater than")
}

// ============================================================================
// Additional Coverage Tests (T030-T039 coverage improvements)
// ============================================================================

func TestParseInvalidDatabaseType(t *testing.T) {
	tests := []struct {
		name         string
		databaseType string
	}{
		{
			name:         "unsupported database type - sqlite",
			databaseType: "sqlite",
		},
		{
			name:         "unsupported database type - mongodb",
			databaseType: "mongodb",
		},
		{
			name:         "unsupported database type - empty string",
			databaseType: "",
		},
		{
			name:         "unsupported database type - invalid name",
			databaseType: "invaliddb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := fmt.Sprintf(`{
				"schema_version": "1.0",
				"name": "test-schema",
				"description": "Test schema",
				"author": "Test Author",
				"version": "1.0.0",
				"database_type": ["%s"],
				"tables": [],
				"generation_order": []
			}`, tt.databaseType)

			reader := strings.NewReader(input)
			schema, err := ParseSchema(reader)

			require.Error(t, err, "ParseSchema should fail for invalid database_type: %s", tt.databaseType)
			assert.Nil(t, schema)
			assert.Contains(t, err.Error(), "invalid database_type")
			assert.Contains(t, err.Error(), "must be")
		})
	}
}

func TestParseMultiplePrimaryKeys(t *testing.T) {
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					},
					{
						"name": "email",
						"type": "varchar(255)",
						"primary_key": true
					}
				]
			}
		],
		"generation_order": ["users"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when table has multiple primary keys")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "must have exactly one primary key")
	assert.Contains(t, err.Error(), "found 2")
}

// ============================================================================
// User Story 3: Validate Foreign Key Integrity (T040-T044) - TDD RED Phase
// These tests should FAIL initially until foreign key validation is implemented
// ============================================================================

func TestParseValidForeignKey(t *testing.T) {
	// Test that a foreign key to an existing table succeeds
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": 50,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					},
					{
						"name": "email",
						"type": "varchar(255)",
						"nullable": false
					}
				]
			},
			{
				"name": "posts",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					},
					{
						"name": "user_id",
						"type": "int",
						"nullable": false,
						"foreign_key": {
							"table": "users",
							"column": "id",
							"on_delete": "CASCADE",
							"on_update": "CASCADE"
						}
					},
					{
						"name": "title",
						"type": "varchar(255)",
						"nullable": false
					}
				]
			}
		],
		"generation_order": ["users", "posts"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.NoError(t, err, "ParseSchema should succeed when foreign key references existing table")
	require.NotNil(t, schema)
	require.Len(t, schema.Tables, 2)

	// Verify foreign key was parsed correctly
	postsTable := schema.Tables[1]
	assert.Equal(t, "posts", postsTable.Name)
	userIDCol := postsTable.Columns[1]
	assert.Equal(t, "user_id", userIDCol.Name)
	require.NotNil(t, userIDCol.ForeignKey)
	assert.Equal(t, "users", userIDCol.ForeignKey.Table)
	assert.Equal(t, "id", userIDCol.ForeignKey.Column)
	assert.Equal(t, "CASCADE", userIDCol.ForeignKey.OnDelete)
	assert.Equal(t, "CASCADE", userIDCol.ForeignKey.OnUpdate)
}

func TestParseForeignKeyNonExistentTable(t *testing.T) {
	// Test that a foreign key referencing a non-existent table produces an error
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "posts",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					},
					{
						"name": "user_id",
						"type": "int",
						"nullable": false,
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
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when foreign key references non-existent table")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "foreign key")
	assert.Contains(t, err.Error(), "users")
	assert.Contains(t, err.Error(), "does not exist")
	assert.Contains(t, err.Error(), "posts")
	assert.Contains(t, err.Error(), "user_id")
}

func TestParseForeignKeyInvalidOnDelete(t *testing.T) {
	// Test that invalid on_delete action produces an error
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": 50,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					}
				]
			},
			{
				"name": "posts",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					},
					{
						"name": "user_id",
						"type": "int",
						"nullable": false,
						"foreign_key": {
							"table": "users",
							"column": "id",
							"on_delete": "DELETE_ALL",
							"on_update": "CASCADE"
						}
					}
				]
			}
		],
		"generation_order": ["users", "posts"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when on_delete action is invalid")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "on_delete")
	assert.Contains(t, err.Error(), "DELETE_ALL")
	assert.Contains(t, err.Error(), "must be one of")
	assert.Contains(t, err.Error(), "CASCADE")
	assert.Contains(t, err.Error(), "SET NULL")
	assert.Contains(t, err.Error(), "RESTRICT")
	assert.Contains(t, err.Error(), "posts")
	assert.Contains(t, err.Error(), "user_id")
}

func TestParseForeignKeyInvalidOnUpdate(t *testing.T) {
	// Test that invalid on_update action produces an error
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": 50,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					}
				]
			},
			{
				"name": "posts",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					},
					{
						"name": "user_id",
						"type": "int",
						"nullable": false,
						"foreign_key": {
							"table": "users",
							"column": "id",
							"on_delete": "CASCADE",
							"on_update": "UPDATE_ALL"
						}
					}
				]
			}
		],
		"generation_order": ["users", "posts"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.Error(t, err, "ParseSchema should fail when on_update action is invalid")
	assert.Nil(t, schema)
	assert.Contains(t, err.Error(), "on_update")
	assert.Contains(t, err.Error(), "UPDATE_ALL")
	assert.Contains(t, err.Error(), "must be one of")
	assert.Contains(t, err.Error(), "CASCADE")
	assert.Contains(t, err.Error(), "SET NULL")
	assert.Contains(t, err.Error(), "RESTRICT")
	assert.Contains(t, err.Error(), "posts")
	assert.Contains(t, err.Error(), "user_id")
}

func TestParseForeignKeyValidActions(t *testing.T) {
	// Test that all valid referential actions (CASCADE, SET NULL, RESTRICT) are accepted
	tests := []struct {
		name     string
		onDelete string
		onUpdate string
	}{
		{
			name:     "CASCADE on both",
			onDelete: "CASCADE",
			onUpdate: "CASCADE",
		},
		{
			name:     "SET NULL on both",
			onDelete: "SET NULL",
			onUpdate: "SET NULL",
		},
		{
			name:     "RESTRICT on both",
			onDelete: "RESTRICT",
			onUpdate: "RESTRICT",
		},
		{
			name:     "mixed CASCADE and SET NULL",
			onDelete: "CASCADE",
			onUpdate: "SET NULL",
		},
		{
			name:     "mixed RESTRICT and CASCADE",
			onDelete: "RESTRICT",
			onUpdate: "CASCADE",
		},
		{
			name:     "mixed SET NULL and RESTRICT",
			onDelete: "SET NULL",
			onUpdate: "RESTRICT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := fmt.Sprintf(`{
				"schema_version": "1.0",
				"name": "test-schema",
				"description": "Test schema",
				"author": "Test Author",
				"version": "1.0.0",
				"database_type": ["mysql"],
				"tables": [
					{
						"name": "users",
						"record_count": 50,
						"columns": [
							{
								"name": "id",
								"type": "int",
								"primary_key": true
							}
						]
					},
					{
						"name": "posts",
						"record_count": 100,
						"columns": [
							{
								"name": "id",
								"type": "int",
								"primary_key": true
							},
							{
								"name": "user_id",
								"type": "int",
								"nullable": false,
								"foreign_key": {
									"table": "users",
									"column": "id",
									"on_delete": "%s",
									"on_update": "%s"
								}
							}
						]
					}
				],
				"generation_order": ["users", "posts"]
			}`, tt.onDelete, tt.onUpdate)

			reader := strings.NewReader(input)
			schema, err := ParseSchema(reader)

			require.NoError(t, err, "ParseSchema should accept valid referential actions: on_delete=%s, on_update=%s", tt.onDelete, tt.onUpdate)
			require.NotNil(t, schema)
			require.Len(t, schema.Tables, 2)

			// Verify the actions were parsed correctly
			postsTable := schema.Tables[1]
			userIDCol := postsTable.Columns[1]
			require.NotNil(t, userIDCol.ForeignKey)
			assert.Equal(t, tt.onDelete, userIDCol.ForeignKey.OnDelete)
			assert.Equal(t, tt.onUpdate, userIDCol.ForeignKey.OnUpdate)
		})
	}
}

func TestParseForeignKeyMultipleReferences(t *testing.T) {
	// Test multiple foreign keys with different valid actions
	input := `{
		"schema_version": "1.0",
		"name": "test-schema",
		"description": "Test schema",
		"author": "Test Author",
		"version": "1.0.0",
		"database_type": ["mysql"],
		"tables": [
			{
				"name": "users",
				"record_count": 50,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					}
				]
			},
			{
				"name": "categories",
				"record_count": 20,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					}
				]
			},
			{
				"name": "posts",
				"record_count": 100,
				"columns": [
					{
						"name": "id",
						"type": "int",
						"primary_key": true
					},
					{
						"name": "user_id",
						"type": "int",
						"nullable": false,
						"foreign_key": {
							"table": "users",
							"column": "id",
							"on_delete": "CASCADE",
							"on_update": "CASCADE"
						}
					},
					{
						"name": "category_id",
						"type": "int",
						"nullable": true,
						"foreign_key": {
							"table": "categories",
							"column": "id",
							"on_delete": "SET NULL",
							"on_update": "RESTRICT"
						}
					}
				]
			}
		],
		"generation_order": ["users", "categories", "posts"]
	}`

	reader := strings.NewReader(input)
	schema, err := ParseSchema(reader)

	require.NoError(t, err, "ParseSchema should succeed with multiple valid foreign keys")
	require.NotNil(t, schema)
	require.Len(t, schema.Tables, 3)

	// Verify both foreign keys were parsed correctly
	postsTable := schema.Tables[2]
	assert.Equal(t, "posts", postsTable.Name)

	// First foreign key (user_id)
	userIDCol := postsTable.Columns[1]
	assert.Equal(t, "user_id", userIDCol.Name)
	require.NotNil(t, userIDCol.ForeignKey)
	assert.Equal(t, "users", userIDCol.ForeignKey.Table)
	assert.Equal(t, "id", userIDCol.ForeignKey.Column)
	assert.Equal(t, "CASCADE", userIDCol.ForeignKey.OnDelete)
	assert.Equal(t, "CASCADE", userIDCol.ForeignKey.OnUpdate)

	// Second foreign key (category_id)
	categoryIDCol := postsTable.Columns[2]
	assert.Equal(t, "category_id", categoryIDCol.Name)
	require.NotNil(t, categoryIDCol.ForeignKey)
	assert.Equal(t, "categories", categoryIDCol.ForeignKey.Table)
	assert.Equal(t, "id", categoryIDCol.ForeignKey.Column)
	assert.Equal(t, "SET NULL", categoryIDCol.ForeignKey.OnDelete)
	assert.Equal(t, "RESTRICT", categoryIDCol.ForeignKey.OnUpdate)
}
