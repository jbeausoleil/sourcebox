/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listSchemasCmd represents the list-schemas command
var listSchemasCmd = &cobra.Command{
	Use:     "list-schemas",
	Aliases: []string{"ls"},
	Short:   "List all available data schemas",
	Long: `List all available verticalized data schemas.

SourceBox provides industry-specific schemas for fintech, healthcare,
retail, and other verticals. Each schema includes realistic field
distributions, relationships, and edge cases.

Schemas are categorized by industry and use case.`,

	Example: `  # List all available schemas
  sourcebox list-schemas

  # Using short alias
  sourcebox ls`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "List-schemas command - implementation coming in F022")
		fmt.Fprintln(cmd.OutOrStdout(), "Available schemas:")
		fmt.Fprintln(cmd.OutOrStdout(), "  - fintech-loans")
		fmt.Fprintln(cmd.OutOrStdout(), "  - healthcare-patients")
		fmt.Fprintln(cmd.OutOrStdout(), "  - retail-orders")
	},
}

func init() {
	rootCmd.AddCommand(listSchemasCmd)
}
