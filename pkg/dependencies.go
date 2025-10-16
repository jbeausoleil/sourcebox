// Package pkg provides temporary dependency imports for F009 (Dependency Management Setup).
//
// This file ensures dependencies are present in go.mod before their actual usage
// in future features (F013 Data Generation Engine, F021 Seed Command).
//
// TODO(F021): Remove this file once seed command implementation imports these packages directly.
package pkg

import (
	// Data generation dependency (used in F013: Data Generation Engine)
	_ "github.com/brianvoe/gofakeit/v6"

	// CLI UX dependencies (used in F021: Seed Command Implementation)
	_ "github.com/fatih/color"
	_ "github.com/schollz/progressbar/v3"

	// Database drivers (used in F021: Seed Command Implementation)
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)
