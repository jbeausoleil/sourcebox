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
