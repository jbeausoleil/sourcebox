/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Global flag variables for output control
var verbose bool
var quiet bool
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sourcebox",
	Short: "Generate realistic, verticalized demo data instantly",
	Long: `SourceBox generates production-like demo data for databases.

Built for developers who need realistic demo data in seconds, not hours.
Verticalized schemas for fintech, healthcare, retail, and more.

Works entirely offline - no cloud APIs, no authentication, no network calls.`,
	Example: `  # Seed MySQL with fintech loan data
  sourcebox seed mysql --schema=fintech-loans --records=1000

  # List all available schemas
  sourcebox list-schemas

  # Export data to SQL file instead of inserting
  sourcebox seed postgres --schema=healthcare-patients --output=data.sql`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// SetVersion sets the version string displayed by the version command
func SetVersion(v string) {
	rootCmd.Version = v
}

func init() {
	// Global flags available to all commands
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress non-error output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path (default: ~/.sourcebox.yaml)")
}
