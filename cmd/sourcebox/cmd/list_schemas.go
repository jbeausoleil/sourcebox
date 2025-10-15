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
		fmt.Println("List-schemas command - implementation coming in F022")
		fmt.Println("Available schemas:")
		fmt.Println("  - fintech-loans")
		fmt.Println("  - healthcare-patients")
		fmt.Println("  - retail-orders")
	},
}

func init() {
	rootCmd.AddCommand(listSchemasCmd)
}
