/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed <database>",
	Short: "Seed a database with realistic demo data",
	Long: `Seed a database with verticalized, production-like demo data.

SourceBox generates realistic data based on industry-specific schemas
(fintech, healthcare, retail) with proper relationships, distributions,
and edge cases. Data is deterministic and reproducible.

Supported databases: mysql, postgres
Supported schemas: fintech-loans, healthcare-patients, retail-orders`,

	Example: `  # Seed MySQL with 1000 fintech loan records
  sourcebox seed mysql --schema=fintech-loans --records=1000

  # Seed Postgres with healthcare patient data
  sourcebox seed postgres --schema=healthcare-patients --records=5000

  # Export to SQL file instead of inserting
  sourcebox seed mysql --schema=fintech-loans --output=loans.sql`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "Seed command - implementation coming in F021")
		fmt.Fprintf(cmd.OutOrStdout(), "  Database: %s\n", args[0])
		schema, _ := cmd.Flags().GetString("schema")
		records, _ := cmd.Flags().GetInt("records")
		fmt.Fprintf(cmd.OutOrStdout(), "  Schema: %s\n", schema)
		fmt.Fprintf(cmd.OutOrStdout(), "  Records: %d\n", records)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Local flags for seed command
	seedCmd.Flags().StringP("schema", "s", "", "schema name (required)")
	seedCmd.Flags().IntP("records", "n", 1000, "number of records to generate")
	seedCmd.Flags().String("host", "localhost", "database host")
	seedCmd.Flags().Int("port", 0, "database port (auto-detect by database type)")
	seedCmd.Flags().String("user", "root", "database user")
	seedCmd.Flags().String("password", "", "database password")
	seedCmd.Flags().String("db-name", "demo", "database name")
	seedCmd.Flags().String("output", "", "export to SQL file instead of inserting")
	seedCmd.Flags().Bool("dry-run", false, "show what would be done without executing")

	// Mark schema flag as required
	_ = seedCmd.MarkFlagRequired("schema")
}
