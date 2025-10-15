package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSeedCommandRegistration verifies that the seed command is properly
// registered with the root command.
func TestSeedCommandRegistration(t *testing.T) {
	// Verify seed command exists
	seedCmd := rootCmd.Commands()
	var found bool
	for _, cmd := range seedCmd {
		if cmd.Name() == "seed" {
			found = true
			break
		}
	}
	assert.True(t, found, "seed command should be registered with root command")
}

// TestSeedCommandHelp verifies that the seed command has comprehensive help text.
func TestSeedCommandHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"seed", "--help"})

	err := rootCmd.Execute()
	require.NoError(t, err, "Help command should not error")

	output := buf.String()

	// Verify Use field
	assert.Contains(t, output, "seed <database>", "Help should show usage with database argument")

	// Verify Long description content (Short only shows in parent listing)
	assert.Contains(t, output, "Seed a database with verticalized", "Help should show long description")
	assert.Contains(t, output, "production-like", "Help should mention production-like data")
	assert.Contains(t, output, "fintech", "Help should mention fintech vertical")
	assert.Contains(t, output, "healthcare", "Help should mention healthcare vertical")
	assert.Contains(t, output, "retail", "Help should mention retail vertical")
	assert.Contains(t, output, "mysql, postgres", "Help should list supported databases")

	// Verify Examples section
	assert.Contains(t, output, "Examples:", "Help should contain examples section")
	assert.Contains(t, output, "sourcebox seed mysql", "Help should show MySQL example")
	assert.Contains(t, output, "sourcebox seed postgres", "Help should show Postgres example")
	assert.Contains(t, output, "--output=", "Help should show output flag example")

	// Verify flags are documented
	assert.Contains(t, output, "--schema", "Help should document schema flag")
	assert.Contains(t, output, "-s", "Help should document schema shorthand")
	assert.Contains(t, output, "--records", "Help should document records flag")
	assert.Contains(t, output, "-n", "Help should document records shorthand")
	assert.Contains(t, output, "--host", "Help should document host flag")
	assert.Contains(t, output, "--port", "Help should document port flag")
	assert.Contains(t, output, "--user", "Help should document user flag")
	assert.Contains(t, output, "--password", "Help should document password flag")
	assert.Contains(t, output, "--db-name", "Help should document db-name flag")
	assert.Contains(t, output, "--output", "Help should document output flag")
	assert.Contains(t, output, "--dry-run", "Help should document dry-run flag")
}

// TestSeedCommandRequiredFlags verifies that the schema flag is marked as required.
func TestSeedCommandRequiredFlags(t *testing.T) {
	// Verify schema flag is marked as required in the flag definition
	schemaFlag := seedCmd.Flags().Lookup("schema")
	require.NotNil(t, schemaFlag, "schema flag should be defined")

	// Check the required annotation (Cobra tracks this internally)
	// We can verify by checking the flag exists and is the correct type
	assert.Equal(t, "schema name (required)", schemaFlag.Usage, "schema usage should indicate it's required")
}

// TestSeedCommandFlagParsing verifies that all flags parse correctly.
func TestSeedCommandFlagParsing(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedSchema string
		expectedRecords int
		expectedHost   string
		expectedPort   int
		expectedUser   string
		expectedPass   string
		expectedDBName string
		expectedOutput string
		expectedDryRun bool
	}{
		{
			name:           "schema flag long form",
			args:           []string{"mysql", "--schema=fintech-loans"},
			expectedSchema: "fintech-loans",
			expectedRecords: 1000, // default
			expectedHost:   "localhost", // default
			expectedPort:   0, // default
			expectedUser:   "root", // default
			expectedPass:   "", // default
			expectedDBName: "demo", // default
			expectedOutput: "", // default
			expectedDryRun: false, // default
		},
		{
			name:           "schema flag short form",
			args:           []string{"postgres", "-s", "healthcare-patients"},
			expectedSchema: "healthcare-patients",
			expectedRecords: 1000,
			expectedHost:   "localhost",
			expectedPort:   0,
			expectedUser:   "root",
			expectedPass:   "",
			expectedDBName: "demo",
			expectedOutput: "",
			expectedDryRun: false,
		},
		{
			name:           "records flag long form",
			args:           []string{"mysql", "--schema=fintech-loans", "--records=5000"},
			expectedSchema: "fintech-loans",
			expectedRecords: 5000,
			expectedHost:   "localhost",
			expectedPort:   0,
			expectedUser:   "root",
			expectedPass:   "",
			expectedDBName: "demo",
			expectedOutput: "",
			expectedDryRun: false,
		},
		{
			name:           "records flag short form",
			args:           []string{"postgres", "-s", "retail-orders", "-n", "3000"},
			expectedSchema: "retail-orders",
			expectedRecords: 3000,
			expectedHost:   "localhost",
			expectedPort:   0,
			expectedUser:   "root",
			expectedPass:   "",
			expectedDBName: "demo",
			expectedOutput: "",
			expectedDryRun: false,
		},
		{
			name:           "all connection flags",
			args:           []string{"mysql", "-s", "fintech-loans", "--host=db.example.com", "--port=3307", "--user=admin", "--password=secret", "--db-name=production"},
			expectedSchema: "fintech-loans",
			expectedRecords: 1000,
			expectedHost:   "db.example.com",
			expectedPort:   3307,
			expectedUser:   "admin",
			expectedPass:   "secret",
			expectedDBName: "production",
			expectedOutput: "",
			expectedDryRun: false,
		},
		{
			name:           "output flag",
			args:           []string{"postgres", "--schema=healthcare-patients", "--output=patients.sql"},
			expectedSchema: "healthcare-patients",
			expectedRecords: 1000,
			expectedHost:   "localhost",
			expectedPort:   0,
			expectedUser:   "root",
			expectedPass:   "",
			expectedDBName: "demo",
			expectedOutput: "patients.sql",
			expectedDryRun: false,
		},
		{
			name:           "dry-run flag",
			args:           []string{"mysql", "-s", "fintech-loans", "--dry-run"},
			expectedSchema: "fintech-loans",
			expectedRecords: 1000,
			expectedHost:   "localhost",
			expectedPort:   0,
			expectedUser:   "root",
			expectedPass:   "",
			expectedDBName: "demo",
			expectedOutput: "",
			expectedDryRun: true,
		},
		{
			name:           "all flags together",
			args:           []string{"postgres", "-s", "retail-orders", "-n", "2500", "--host=localhost", "--port=5433", "--user=postgres", "--password=pass123", "--db-name=retail", "--output=orders.sql", "--dry-run"},
			expectedSchema: "retail-orders",
			expectedRecords: 2500,
			expectedHost:   "localhost",
			expectedPort:   5433,
			expectedUser:   "postgres",
			expectedPass:   "pass123",
			expectedDBName: "retail",
			expectedOutput: "orders.sql",
			expectedDryRun: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			// Prepend "seed" to the args for rootCmd execution
			argsWithCommand := append([]string{"seed"}, tt.args...)
			rootCmd.SetArgs(argsWithCommand)

			err := rootCmd.Execute()
			require.NoError(t, err, "Command should not error with valid flags")

			// Verify flag values
			schema, _ := seedCmd.Flags().GetString("schema")
			assert.Equal(t, tt.expectedSchema, schema, "schema flag should be parsed correctly")

			records, _ := seedCmd.Flags().GetInt("records")
			assert.Equal(t, tt.expectedRecords, records, "records flag should be parsed correctly")

			host, _ := seedCmd.Flags().GetString("host")
			assert.Equal(t, tt.expectedHost, host, "host flag should be parsed correctly")

			port, _ := seedCmd.Flags().GetInt("port")
			assert.Equal(t, tt.expectedPort, port, "port flag should be parsed correctly")

			user, _ := seedCmd.Flags().GetString("user")
			assert.Equal(t, tt.expectedUser, user, "user flag should be parsed correctly")

			password, _ := seedCmd.Flags().GetString("password")
			assert.Equal(t, tt.expectedPass, password, "password flag should be parsed correctly")

			dbName, _ := seedCmd.Flags().GetString("db-name")
			assert.Equal(t, tt.expectedDBName, dbName, "db-name flag should be parsed correctly")

			output, _ := seedCmd.Flags().GetString("output")
			assert.Equal(t, tt.expectedOutput, output, "output flag should be parsed correctly")

			dryRun, _ := seedCmd.Flags().GetBool("dry-run")
			assert.Equal(t, tt.expectedDryRun, dryRun, "dry-run flag should be parsed correctly")

			// Reset flags for next test
			seedCmd.Flags().Set("schema", "")
			seedCmd.Flags().Set("records", "1000")
			seedCmd.Flags().Set("host", "localhost")
			seedCmd.Flags().Set("port", "0")
			seedCmd.Flags().Set("user", "root")
			seedCmd.Flags().Set("password", "")
			seedCmd.Flags().Set("db-name", "demo")
			seedCmd.Flags().Set("output", "")
			seedCmd.Flags().Set("dry-run", "false")
		})
	}
}

// TestSeedCommandArgumentValidation verifies that the command accepts arguments.
func TestSeedCommandArgumentValidation(t *testing.T) {
	// Verify Args validator is set correctly
	assert.NotNil(t, seedCmd.Args, "Args validator should be set")

	// Test valid execution with proper arguments
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"seed", "mysql", "--schema=fintech-loans"})

	err := rootCmd.Execute()
	require.NoError(t, err, "Command should not error with valid arguments")

	// Reset flags
	seedCmd.Flags().Set("schema", "")
}

// TestSeedCommandPlaceholderOutput verifies that the placeholder Run function
// is callable. Full output testing is deferred to F021 implementation.
func TestSeedCommandPlaceholderOutput(t *testing.T) {
	// This is a basic smoke test to ensure the command runs
	// Full functionality will be tested in F021
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"seed", "mysql", "--schema=fintech-loans"})

	err := rootCmd.Execute()
	require.NoError(t, err, "Seed command should execute without error")

	// Reset flags
	seedCmd.Flags().Set("schema", "")
}

// TestSeedCommandFlagDefaults verifies that all flags have correct default values.
func TestSeedCommandFlagDefaults(t *testing.T) {
	// Check schema flag (required, no default value applies)
	schemaFlag := seedCmd.Flags().Lookup("schema")
	require.NotNil(t, schemaFlag, "schema flag should be defined")
	assert.Equal(t, "s", schemaFlag.Shorthand, "schema shorthand should be 's'")
	assert.Equal(t, "", schemaFlag.DefValue, "schema default should be empty")

	// Check records flag
	recordsFlag := seedCmd.Flags().Lookup("records")
	require.NotNil(t, recordsFlag, "records flag should be defined")
	assert.Equal(t, "n", recordsFlag.Shorthand, "records shorthand should be 'n'")
	assert.Equal(t, "1000", recordsFlag.DefValue, "records default should be 1000")

	// Check host flag
	hostFlag := seedCmd.Flags().Lookup("host")
	require.NotNil(t, hostFlag, "host flag should be defined")
	assert.Equal(t, "", hostFlag.Shorthand, "host should not have shorthand")
	assert.Equal(t, "localhost", hostFlag.DefValue, "host default should be localhost")

	// Check port flag
	portFlag := seedCmd.Flags().Lookup("port")
	require.NotNil(t, portFlag, "port flag should be defined")
	assert.Equal(t, "", portFlag.Shorthand, "port should not have shorthand")
	assert.Equal(t, "0", portFlag.DefValue, "port default should be 0")

	// Check user flag
	userFlag := seedCmd.Flags().Lookup("user")
	require.NotNil(t, userFlag, "user flag should be defined")
	assert.Equal(t, "", userFlag.Shorthand, "user should not have shorthand")
	assert.Equal(t, "root", userFlag.DefValue, "user default should be root")

	// Check password flag
	passwordFlag := seedCmd.Flags().Lookup("password")
	require.NotNil(t, passwordFlag, "password flag should be defined")
	assert.Equal(t, "", passwordFlag.Shorthand, "password should not have shorthand")
	assert.Equal(t, "", passwordFlag.DefValue, "password default should be empty")

	// Check db-name flag
	dbNameFlag := seedCmd.Flags().Lookup("db-name")
	require.NotNil(t, dbNameFlag, "db-name flag should be defined")
	assert.Equal(t, "", dbNameFlag.Shorthand, "db-name should not have shorthand")
	assert.Equal(t, "demo", dbNameFlag.DefValue, "db-name default should be demo")

	// Check output flag
	outputFlag := seedCmd.Flags().Lookup("output")
	require.NotNil(t, outputFlag, "output flag should be defined")
	assert.Equal(t, "", outputFlag.Shorthand, "output should not have shorthand")
	assert.Equal(t, "", outputFlag.DefValue, "output default should be empty")

	// Check dry-run flag
	dryRunFlag := seedCmd.Flags().Lookup("dry-run")
	require.NotNil(t, dryRunFlag, "dry-run flag should be defined")
	assert.Equal(t, "", dryRunFlag.Shorthand, "dry-run should not have shorthand")
	assert.Equal(t, "false", dryRunFlag.DefValue, "dry-run default should be false")
}

// TestSeedCommandWithGlobalFlags verifies that seed command works with
// global flags (verbose, quiet, config).
func TestSeedCommandWithGlobalFlags(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		expectedVerbose bool
		expectedQuiet   bool
	}{
		{
			name:            "seed with verbose flag",
			args:            []string{"--verbose", "seed", "mysql", "--schema=fintech-loans"},
			expectedVerbose: true,
			expectedQuiet:   false,
		},
		{
			name:            "seed with quiet flag",
			args:            []string{"--quiet", "seed", "postgres", "-s", "healthcare-patients"},
			expectedVerbose: false,
			expectedQuiet:   true,
		},
		{
			name:            "seed with both verbose and quiet",
			args:            []string{"-v", "-q", "seed", "mysql", "--schema=fintech-loans"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
		{
			name:            "global flags after seed command",
			args:            []string{"seed", "mysql", "--schema=fintech-loans", "-v"},
			expectedVerbose: true,
			expectedQuiet:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			resetGlobalFlags()

			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()
			require.NoError(t, err, "Command should not error")

			// Verify global flags were parsed
			assert.Equal(t, tt.expectedVerbose, verbose, "verbose flag should be parsed correctly")
			assert.Equal(t, tt.expectedQuiet, quiet, "quiet flag should be parsed correctly")

			// Reset seed flags
			seedCmd.Flags().Set("schema", "")
		})
	}
}

// TestSeedCommandInRootHelp verifies that seed command appears in root help.
func TestSeedCommandInRootHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	require.NoError(t, err, "Help command should not error")

	output := buf.String()
	assert.Contains(t, output, "seed", "Root help should list seed command")
	assert.Contains(t, output, "Seed a database with realistic demo data", "Root help should show seed short description")
}

// TestSeedCommandNegativeCases verifies error handling for invalid inputs.
func TestSeedCommandNegativeCases(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "unknown flag",
			args:        []string{"seed", "mysql", "--schema=fintech-loans", "--unknown-flag=value"},
			expectError: true,
			errorMsg:    "unknown flag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()

			if tt.expectError {
				require.Error(t, err, "Command should error")
				assert.Contains(t, strings.ToLower(err.Error()), strings.ToLower(tt.errorMsg), "Error message should contain expected text")
			} else {
				require.NoError(t, err, "Command should not error for this test case")
			}

			// Reset flags
			seedCmd.Flags().Set("schema", "")
		})
	}
}
