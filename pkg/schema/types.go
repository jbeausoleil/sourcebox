package schema

// Schema represents the top-level schema definition for a database schema.
// It matches the JSON schema format defined in F007.
type Schema struct {
	SchemaVersion   string           `json:"schema_version"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Author          string           `json:"author"`
	Version         string           `json:"version"`
	DatabaseType    []string         `json:"database_type"`
	Metadata        SchemaMetadata   `json:"metadata"`
	Tables          []Table          `json:"tables"`
	Relationships   []Relationship   `json:"relationships"`
	GenerationOrder []string         `json:"generation_order"`
	ValidationRules []ValidationRule `json:"validation_rules"`
}

// SchemaMetadata contains metadata about the schema.
type SchemaMetadata struct {
	Industry       string   `json:"industry"`
	Tags           []string `json:"tags"`
	TotalRecords   int      `json:"total_records"`
	ComplexityTier int      `json:"complexity_tier"`
}

// Table represents a database table definition.
type Table struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	RecordCount int      `json:"record_count"`
	Columns     []Column `json:"columns"`
	Indexes     []Index  `json:"indexes"`
}

// Column represents a database column definition.
type Column struct {
	Name            string                 `json:"name"`
	Type            string                 `json:"type"`
	Nullable        bool                   `json:"nullable"`
	PrimaryKey      bool                   `json:"primary_key"`
	AutoIncrement   bool                   `json:"auto_increment"`
	Default         *string                `json:"default,omitempty"`
	Unique          bool                   `json:"unique"`
	Description     string                 `json:"description"`
	Generator       string                 `json:"generator"`
	GeneratorParams map[string]interface{} `json:"generator_params"`
	ForeignKey      *ForeignKey            `json:"foreign_key,omitempty"`
}

// ForeignKey represents a foreign key constraint on a column.
type ForeignKey struct {
	Table    string `json:"table"`
	Column   string `json:"column"`
	OnDelete string `json:"on_delete"`
	OnUpdate string `json:"on_update"`
}

// Index represents a database index definition.
type Index struct {
	Name    string   `json:"name"`
	Columns []string `json:"columns"`
	Type    string   `json:"type"`
	Unique  bool     `json:"unique"`
}

// Relationship represents an explicit relationship between tables.
// This is for documentation purposes; the actual foreign key constraint
// is defined inline in the Column struct.
type Relationship struct {
	FromTable        string `json:"from_table"`
	FromColumn       string `json:"from_column"`
	ToTable          string `json:"to_table"`
	ToColumn         string `json:"to_column"`
	RelationshipType string `json:"relationship_type"`
	Description      string `json:"description"`
}

// ValidationRule represents a validation rule for the schema.
type ValidationRule struct {
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}
