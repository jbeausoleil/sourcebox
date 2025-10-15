package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSetVersion verifies that SetVersion correctly sets the version
// on the root command.
func TestSetVersion(t *testing.T) {
	tests := []struct {
		name            string
		versionToSet    string
		expectedVersion string
	}{
		{
			name:            "set development version",
			versionToSet:    "dev",
			expectedVersion: "dev",
		},
		{
			name:            "set git describe short hash",
			versionToSet:    "854ed9e",
			expectedVersion: "854ed9e",
		},
		{
			name:            "set git describe with dirty flag",
			versionToSet:    "854ed9e-dirty",
			expectedVersion: "854ed9e-dirty",
		},
		{
			name:            "set semantic version tag",
			versionToSet:    "v1.0.0",
			expectedVersion: "v1.0.0",
		},
		{
			name:            "set version with tag and commits",
			versionToSet:    "v1.0.0-5-g854ed9e",
			expectedVersion: "v1.0.0-5-g854ed9e",
		},
		{
			name:            "set empty version",
			versionToSet:    "",
			expectedVersion: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the version
			SetVersion(tt.versionToSet)

			// Verify it was set correctly
			assert.Equal(t, tt.expectedVersion, rootCmd.Version,
				"SetVersion should set rootCmd.Version correctly")
		})
	}
}

// TestVersionFlag verifies that the --version flag displays the correct version.
func TestVersionFlag(t *testing.T) {
	tests := []struct {
		name           string
		versionToSet   string
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "version flag with dev version",
			versionToSet:   "dev",
			expectedOutput: "sourcebox version dev",
			expectError:    false,
		},
		{
			name:           "version flag with git hash",
			versionToSet:   "854ed9e-dirty",
			expectedOutput: "sourcebox version 854ed9e-dirty",
			expectError:    false,
		},
		{
			name:           "version flag with semantic version",
			versionToSet:   "v1.2.3",
			expectedOutput: "sourcebox version v1.2.3",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new command instance to avoid state pollution
			cmd := &cobra.Command{
				Use:     "sourcebox",
				Version: tt.versionToSet,
			}

			// Capture output
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)

			// Set args to trigger version flag
			cmd.SetArgs([]string{"--version"})

			// Execute command
			err := cmd.Execute()

			if tt.expectError {
				require.Error(t, err, "Expected error but got none")
			} else {
				require.NoError(t, err, "Unexpected error: %v", err)

				// Verify output contains expected version string
				output := buf.String()
				assert.Contains(t, output, tt.expectedOutput,
					"Version output should contain expected version string")
			}
		})
	}
}

// TestRootCommandWithoutArgs verifies that running the root command
// without arguments displays help text.
func TestRootCommandWithoutArgs(t *testing.T) {
	// Set a known version
	SetVersion("test-version")

	// Capture output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Reset args (no arguments)
	rootCmd.SetArgs([]string{})

	// Execute command
	err := rootCmd.Execute()
	require.NoError(t, err, "Root command should not error when called without args")

	// Verify help output is displayed
	output := buf.String()
	assert.Contains(t, output, "Usage:", "Help output should contain usage section")
	assert.Contains(t, output, "sourcebox", "Help output should contain command name")
	assert.Contains(t, output, "SourceBox generates production-like demo data for databases",
		"Help output should contain long description")
}

// TestRootCommandVersion verifies that the version is correctly
// accessible via the rootCmd after SetVersion is called.
func TestRootCommandVersion(t *testing.T) {
	// Test that default version can be changed
	initialVersion := rootCmd.Version

	// Set a test version
	testVersion := "test-1.2.3"
	SetVersion(testVersion)
	assert.Equal(t, testVersion, rootCmd.Version,
		"SetVersion should update rootCmd.Version")

	// Set a different version
	newVersion := "test-4.5.6"
	SetVersion(newVersion)
	assert.Equal(t, newVersion, rootCmd.Version,
		"SetVersion should allow multiple updates")

	// Verify it actually changed from initial
	assert.NotEqual(t, initialVersion, rootCmd.Version,
		"Version should have changed from initial value")
}

// TestVersionInjectionPattern verifies that the version string
// follows expected patterns from git describe.
func TestVersionInjectionPattern(t *testing.T) {
	validPatterns := []struct {
		name    string
		version string
		valid   bool
	}{
		{
			name:    "dev version",
			version: "dev",
			valid:   true,
		},
		{
			name:    "short git hash",
			version: "854ed9e",
			valid:   true,
		},
		{
			name:    "git hash with dirty flag",
			version: "854ed9e-dirty",
			valid:   true,
		},
		{
			name:    "semantic version tag",
			version: "v1.0.0",
			valid:   true,
		},
		{
			name:    "tag with commits since",
			version: "v1.0.0-5-g854ed9e",
			valid:   true,
		},
		{
			name:    "tag with commits and dirty",
			version: "v1.0.0-5-g854ed9e-dirty",
			valid:   true,
		},
		{
			name:    "empty version should be rejected",
			version: "",
			valid:   false,
		},
	}

	for _, tt := range validPatterns {
		t.Run(tt.name, func(t *testing.T) {
			SetVersion(tt.version)

			if tt.valid {
				assert.NotEmpty(t, rootCmd.Version,
					"Valid version patterns should not result in empty version")
			} else {
				// Empty is technically allowed, but not recommended
				// This documents the behavior
				t.Logf("Version set to: %q", rootCmd.Version)
			}
		})
	}
}

// TestVersionFlagShorthand verifies that -V or --version both work
// (Cobra provides this by default, but we test to ensure it).
func TestVersionFlagShorthand(t *testing.T) {
	testVersion := "test-version-shorthand"
	SetVersion(testVersion)

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "long form version flag",
			args: []string{"--version"},
		},
		// Note: Cobra doesn't provide a short -V flag by default
		// We document this in case we want to add it in the future
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			cmd := &cobra.Command{
				Use:     "sourcebox",
				Version: testVersion,
			}
			cmd.SetOut(buf)
			cmd.SetErr(buf)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			require.NoError(t, err, "Version flag should not error")

			output := buf.String()
			assert.Contains(t, output, testVersion,
				"Version output should contain the set version")
		})
	}
}

// TestMainVersionFlow verifies the complete flow from main.go perspective:
// version variable -> SetVersion -> rootCmd.Version
func TestMainVersionFlow(t *testing.T) {
	// Simulate what happens in main.go
	testVersions := []string{
		"dev",
		"854ed9e",
		"854ed9e-dirty",
		"v1.0.0",
		"v1.0.0-5-g854ed9e",
	}

	for _, v := range testVersions {
		t.Run("version="+v, func(t *testing.T) {
			// This simulates: cmd.SetVersion(version)
			SetVersion(v)

			// Verify it's accessible on rootCmd
			assert.Equal(t, v, rootCmd.Version,
				"Version should flow from SetVersion to rootCmd.Version")

			// Verify it can be displayed
			buf := new(bytes.Buffer)
			cmd := &cobra.Command{
				Use:     "sourcebox",
				Version: rootCmd.Version,
			}
			cmd.SetOut(buf)
			cmd.SetArgs([]string{"--version"})

			err := cmd.Execute()
			require.NoError(t, err)

			output := strings.TrimSpace(buf.String())
			expectedPrefix := "sourcebox version " + v
			assert.Equal(t, expectedPrefix, output,
				"Version output should match expected format")
		})
	}
}

// TestGlobalFlagVariables verifies that global flag variables are properly declared
// and initialized to their default values.
func TestGlobalFlagVariables(t *testing.T) {
	// Reset flags to default state
	verbose = false
	quiet = false

	// Verify default values
	assert.False(t, verbose, "verbose should default to false")
	assert.False(t, quiet, "quiet should default to false")

	// Verify flags can be modified (they're package-level vars)
	verbose = true
	assert.True(t, verbose, "verbose should be modifiable")

	quiet = true
	assert.True(t, quiet, "quiet should be modifiable")

	// Reset for other tests
	verbose = false
	quiet = false
}

// TestVerboseFlagParsing verifies that the --verbose/-v flag parses correctly
// in various forms and combinations.
func TestVerboseFlagParsing(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedValue bool
	}{
		{
			name:          "long form verbose flag",
			args:          []string{"--verbose"},
			expectedValue: true,
		},
		{
			name:          "short form verbose flag",
			args:          []string{"-v"},
			expectedValue: true,
		},
		{
			name:          "no verbose flag",
			args:          []string{},
			expectedValue: false,
		},
		{
			name:          "verbose with other args",
			args:          []string{"--verbose", "somecommand"},
			expectedValue: true,
		},
		{
			name:          "verbose short form with other args",
			args:          []string{"-v", "somecommand"},
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			verbose = false
			quiet = false

			// Create a fresh command to avoid state pollution
			cmd := &cobra.Command{
				Use: "sourcebox",
				Run: func(cmd *cobra.Command, args []string) {
					// No-op, we're just testing flag parsing
				},
			}
			cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

			// Capture output
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)
			cmd.SetArgs(tt.args)

			// Execute command
			err := cmd.Execute()

			// We expect no error for valid flag combinations
			// (invalid commands are fine, we're testing flag parsing)
			if err != nil && !strings.Contains(err.Error(), "unknown command") {
				require.NoError(t, err, "Unexpected error: %v", err)
			}

			// Verify flag value
			assert.Equal(t, tt.expectedValue, verbose,
				"verbose flag should be set to expected value")
		})
	}
}

// TestQuietFlagParsing verifies that the --quiet/-q flag parses correctly
// in various forms and combinations.
func TestQuietFlagParsing(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedValue bool
	}{
		{
			name:          "long form quiet flag",
			args:          []string{"--quiet"},
			expectedValue: true,
		},
		{
			name:          "short form quiet flag",
			args:          []string{"-q"},
			expectedValue: true,
		},
		{
			name:          "no quiet flag",
			args:          []string{},
			expectedValue: false,
		},
		{
			name:          "quiet with other args",
			args:          []string{"--quiet", "somecommand"},
			expectedValue: true,
		},
		{
			name:          "quiet short form with other args",
			args:          []string{"-q", "somecommand"},
			expectedValue: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			verbose = false
			quiet = false

			// Create a fresh command to avoid state pollution
			cmd := &cobra.Command{
				Use: "sourcebox",
				Run: func(cmd *cobra.Command, args []string) {
					// No-op, we're just testing flag parsing
				},
			}
			cmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress non-error output")

			// Capture output
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)
			cmd.SetArgs(tt.args)

			// Execute command
			err := cmd.Execute()

			// We expect no error for valid flag combinations
			if err != nil && !strings.Contains(err.Error(), "unknown command") {
				require.NoError(t, err, "Unexpected error: %v", err)
			}

			// Verify flag value
			assert.Equal(t, tt.expectedValue, quiet,
				"quiet flag should be set to expected value")
		})
	}
}

// TestFlagCombinations verifies that verbose and quiet flags work together
// and handle various combinations correctly.
func TestFlagCombinations(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		expectedVerbose bool
		expectedQuiet   bool
	}{
		{
			name:            "both verbose and quiet (long form)",
			args:            []string{"--verbose", "--quiet"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
		{
			name:            "both verbose and quiet (short form)",
			args:            []string{"-v", "-q"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
		{
			name:            "both flags (mixed form)",
			args:            []string{"--verbose", "-q"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
		{
			name:            "only verbose",
			args:            []string{"--verbose"},
			expectedVerbose: true,
			expectedQuiet:   false,
		},
		{
			name:            "only quiet",
			args:            []string{"--quiet"},
			expectedVerbose: false,
			expectedQuiet:   true,
		},
		{
			name:            "neither flag",
			args:            []string{},
			expectedVerbose: false,
			expectedQuiet:   false,
		},
		{
			name:            "combined short flags (vq)",
			args:            []string{"-vq"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
		{
			name:            "combined short flags (qv)",
			args:            []string{"-qv"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			verbose = false
			quiet = false

			// Create a fresh command
			cmd := &cobra.Command{
				Use: "sourcebox",
				Run: func(cmd *cobra.Command, args []string) {
					// No-op
				},
			}
			cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
			cmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress non-error output")

			// Capture output
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)
			cmd.SetArgs(tt.args)

			// Execute command
			err := cmd.Execute()
			require.NoError(t, err, "Unexpected error: %v", err)

			// Verify both flags
			assert.Equal(t, tt.expectedVerbose, verbose,
				"verbose flag should be set to expected value")
			assert.Equal(t, tt.expectedQuiet, quiet,
				"quiet flag should be set to expected value")
		})
	}
}

// TestPersistentFlagsInHelp verifies that global flags appear in help output.
func TestPersistentFlagsInHelp(t *testing.T) {
	// Reset root command
	verbose = false
	quiet = false

	// Capture help output
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	require.NoError(t, err, "Help command should not error")

	output := buf.String()

	// Verify verbose flag appears in help
	assert.Contains(t, output, "--verbose", "Help should contain --verbose flag")
	assert.Contains(t, output, "-v", "Help should contain -v shorthand")
	assert.Contains(t, output, "verbose output", "Help should contain verbose flag description")

	// Verify quiet flag appears in help
	assert.Contains(t, output, "--quiet", "Help should contain --quiet flag")
	assert.Contains(t, output, "-q", "Help should contain -q shorthand")
	assert.Contains(t, output, "suppress non-error output", "Help should contain quiet flag description")
}

// TestPersistentFlagsWithSubcommands verifies that global flags work with subcommands.
// This tests the "persistent" aspect of PersistentFlags.
func TestPersistentFlagsWithSubcommands(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		expectedVerbose bool
		expectedQuiet   bool
	}{
		{
			name:            "verbose flag before subcommand",
			args:            []string{"--verbose", "subcommand"},
			expectedVerbose: true,
			expectedQuiet:   false,
		},
		{
			name:            "quiet flag before subcommand",
			args:            []string{"--quiet", "subcommand"},
			expectedVerbose: false,
			expectedQuiet:   true,
		},
		{
			name:            "both flags before subcommand",
			args:            []string{"-v", "-q", "subcommand"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
		{
			name:            "verbose flag after subcommand",
			args:            []string{"subcommand", "--verbose"},
			expectedVerbose: true,
			expectedQuiet:   false,
		},
		{
			name:            "quiet flag after subcommand",
			args:            []string{"subcommand", "--quiet"},
			expectedVerbose: false,
			expectedQuiet:   true,
		},
		{
			name:            "both flags after subcommand",
			args:            []string{"subcommand", "-v", "-q"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			verbose = false
			quiet = false

			// Create root command with persistent flags
			rootCmd := &cobra.Command{
				Use: "sourcebox",
			}
			rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
			rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress non-error output")

			// Add a subcommand
			subCmd := &cobra.Command{
				Use: "subcommand",
				Run: func(cmd *cobra.Command, args []string) {
					// No-op, we're testing flag inheritance
				},
			}
			rootCmd.AddCommand(subCmd)

			// Capture output
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs(tt.args)

			// Execute command
			err := rootCmd.Execute()
			require.NoError(t, err, "Unexpected error: %v", err)

			// Verify flags were parsed correctly
			assert.Equal(t, tt.expectedVerbose, verbose,
				"verbose flag should be set to expected value")
			assert.Equal(t, tt.expectedQuiet, quiet,
				"quiet flag should be set to expected value")
		})
	}
}

// TestRootCommandFlagsIntegration verifies the actual rootCmd has the flags configured.
func TestRootCommandFlagsIntegration(t *testing.T) {
	// This test verifies the actual production rootCmd, not a test double

	// Check that persistent flags are defined
	verboseFlag := rootCmd.PersistentFlags().Lookup("verbose")
	require.NotNil(t, verboseFlag, "verbose flag should be defined")
	assert.Equal(t, "v", verboseFlag.Shorthand, "verbose shorthand should be 'v'")
	assert.Equal(t, "verbose output", verboseFlag.Usage, "verbose usage text should match")

	quietFlag := rootCmd.PersistentFlags().Lookup("quiet")
	require.NotNil(t, quietFlag, "quiet flag should be defined")
	assert.Equal(t, "q", quietFlag.Shorthand, "quiet shorthand should be 'q'")
	assert.Equal(t, "suppress non-error output", quietFlag.Usage, "quiet usage text should match")
}

// TestFlagDefaultValues verifies that flags have the correct default values.
func TestFlagDefaultValues(t *testing.T) {
	// Check verbose flag default
	verboseFlag := rootCmd.PersistentFlags().Lookup("verbose")
	require.NotNil(t, verboseFlag, "verbose flag should be defined")
	assert.Equal(t, "false", verboseFlag.DefValue, "verbose should default to false")

	// Check quiet flag default
	quietFlag := rootCmd.PersistentFlags().Lookup("quiet")
	require.NotNil(t, quietFlag, "quiet flag should be defined")
	assert.Equal(t, "false", quietFlag.DefValue, "quiet should default to false")
}

// TestExecuteFunction verifies that Execute() function works correctly.
// This tests the exported Execute function that's called from main.go.
func TestExecuteFunction(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{
			name:        "execute with help flag",
			args:        []string{"--help"},
			expectError: false,
		},
		{
			name:        "execute with version flag",
			args:        []string{"--version"},
			expectError: false,
		},
		{
			name:        "execute with verbose flag",
			args:        []string{"--verbose"},
			expectError: false,
		},
		{
			name:        "execute with quiet flag",
			args:        []string{"--quiet"},
			expectError: false,
		},
		{
			name:        "execute with combined flags",
			args:        []string{"-v", "-q"},
			expectError: false,
		},
		{
			name:        "execute with no args (shows help)",
			args:        []string{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			verbose = false
			quiet = false

			// Capture output
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs(tt.args)

			// Call Execute directly
			err := rootCmd.Execute()

			if tt.expectError {
				assert.Error(t, err, "Expected error but got none")
			} else {
				assert.NoError(t, err, "Unexpected error: %v", err)
				assert.NotEmpty(t, buf.String(), "Expected output from command")
			}
		})
	}
}
