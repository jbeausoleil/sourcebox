package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListSchemasCommandRegistration verifies that the list-schemas command
// is properly registered with the root command.
func TestListSchemasCommandRegistration(t *testing.T) {
	// Verify list-schemas command exists
	commands := rootCmd.Commands()
	var found bool
	for _, cmd := range commands {
		if cmd.Name() == "list-schemas" {
			found = true
			break
		}
	}
	assert.True(t, found, "list-schemas command should be registered with root command")
}

// TestListSchemasCommandAlias verifies that the ls alias works.
func TestListSchemasCommandAlias(t *testing.T) {
	// Check that alias is defined
	aliases := listSchemasCmd.Aliases
	require.Len(t, aliases, 1, "list-schemas should have one alias")
	assert.Equal(t, "ls", aliases[0], "alias should be 'ls'")

	// Test that alias actually works by executing through root
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"ls"})

	err := rootCmd.Execute()
	require.NoError(t, err, "ls alias should execute without error")

	output := buf.String()
	assert.Contains(t, output, "Available schemas:", "ls alias should produce same output as list-schemas")
}

// TestListSchemasCommandHelp verifies that the list-schemas command has
// comprehensive help text.
func TestListSchemasCommandHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"list-schemas", "--help"})

	err := rootCmd.Execute()
	require.NoError(t, err, "Help command should not error")

	output := buf.String()

	// Verify Use field
	assert.Contains(t, output, "list-schemas", "Help should show command name")

	// Verify Short description
	assert.Contains(t, output, "List all available data schemas", "Help should contain short description")

	// Verify Long description content
	assert.Contains(t, output, "verticalized", "Help should mention verticalized schemas")
	assert.Contains(t, output, "fintech", "Help should mention fintech vertical")
	assert.Contains(t, output, "healthcare", "Help should mention healthcare vertical")
	assert.Contains(t, output, "retail", "Help should mention retail vertical")
	assert.Contains(t, output, "industry", "Help should mention industry categorization")

	// Verify Examples section
	assert.Contains(t, output, "Examples:", "Help should contain examples section")
	assert.Contains(t, output, "sourcebox list-schemas", "Help should show full command example")
	assert.Contains(t, output, "sourcebox ls", "Help should show alias example")

	// Verify Aliases section
	assert.Contains(t, output, "Aliases:", "Help should list aliases")
	assert.Contains(t, output, "ls", "Help should show ls alias")
}

// TestListSchemasCommandExecution verifies that the command executes and
// produces expected placeholder output.
func TestListSchemasCommandExecution(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		expectedInOutput []string
	}{
		{
			name: "list-schemas command",
			args: []string{"list-schemas"},
			expectedInOutput: []string{
				"List-schemas command",
				"implementation coming in F022",
				"Available schemas:",
				"fintech-loans",
				"healthcare-patients",
				"retail-orders",
			},
		},
		{
			name: "ls alias",
			args: []string{"ls"},
			expectedInOutput: []string{
				"List-schemas command",
				"implementation coming in F022",
				"Available schemas:",
				"fintech-loans",
				"healthcare-patients",
				"retail-orders",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()
			require.NoError(t, err, "Command should not error")

			output := buf.String()
			for _, expected := range tt.expectedInOutput {
				assert.Contains(t, output, expected, "Output should contain expected text")
			}
		})
	}
}

// TestListSchemasCommandNoFlags verifies that list-schemas has no local flags.
func TestListSchemasCommandNoFlags(t *testing.T) {
	// list-schemas should not have any local flags
	localFlags := listSchemasCmd.LocalFlags()
	assert.Equal(t, 0, localFlags.NFlag(), "list-schemas should have no local flags")
}

// TestListSchemasCommandNoArguments verifies that list-schemas accepts no arguments.
func TestListSchemasCommandNoArguments(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "no arguments (valid)",
			args:        []string{"list-schemas"},
			expectError: false,
		},
		{
			name:        "with extra arguments (should be ignored or error)",
			args:        []string{"list-schemas", "extra", "args"},
			expectError: false, // Cobra by default ignores extra args unless Args validator is set
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
				require.NoError(t, err, "Command should not error")
			}
		})
	}
}

// TestListSchemasCommandWithGlobalFlags verifies that list-schemas works with
// global flags (verbose, quiet, config).
func TestListSchemasCommandWithGlobalFlags(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		expectedVerbose bool
		expectedQuiet   bool
	}{
		{
			name:            "list-schemas with verbose flag",
			args:            []string{"--verbose", "list-schemas"},
			expectedVerbose: true,
			expectedQuiet:   false,
		},
		{
			name:            "list-schemas with quiet flag",
			args:            []string{"--quiet", "list-schemas"},
			expectedVerbose: false,
			expectedQuiet:   true,
		},
		{
			name:            "list-schemas with both verbose and quiet",
			args:            []string{"-v", "-q", "list-schemas"},
			expectedVerbose: true,
			expectedQuiet:   true,
		},
		{
			name:            "global flags after list-schemas",
			args:            []string{"list-schemas", "-v"},
			expectedVerbose: true,
			expectedQuiet:   false,
		},
		{
			name:            "ls alias with verbose flag",
			args:            []string{"-v", "ls"},
			expectedVerbose: true,
			expectedQuiet:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags
			verbose = false
			quiet = false

			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()
			require.NoError(t, err, "Command should not error")

			// Verify global flags were parsed
			assert.Equal(t, tt.expectedVerbose, verbose, "verbose flag should be parsed correctly")
			assert.Equal(t, tt.expectedQuiet, quiet, "quiet flag should be parsed correctly")
		})
	}
}

// TestListSchemasCommandInRootHelp verifies that list-schemas command appears
// in root help.
func TestListSchemasCommandInRootHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	require.NoError(t, err, "Help command should not error")

	output := buf.String()
	assert.Contains(t, output, "list-schemas", "Root help should list list-schemas command")
	assert.Contains(t, output, "List all available data schemas", "Root help should show list-schemas short description")
}

// TestListSchemasVsLsAlias verifies that both command forms produce identical output.
func TestListSchemasVsLsAlias(t *testing.T) {
	// Run with full command name
	buf1 := new(bytes.Buffer)
	rootCmd.SetOut(buf1)
	rootCmd.SetErr(buf1)
	rootCmd.SetArgs([]string{"list-schemas"})
	err1 := rootCmd.Execute()
	require.NoError(t, err1, "list-schemas should not error")
	output1 := buf1.String()

	// Run with alias
	buf2 := new(bytes.Buffer)
	rootCmd.SetOut(buf2)
	rootCmd.SetErr(buf2)
	rootCmd.SetArgs([]string{"ls"})
	err2 := rootCmd.Execute()
	require.NoError(t, err2, "ls should not error")
	output2 := buf2.String()

	// Verify identical output
	assert.Equal(t, output1, output2, "list-schemas and ls should produce identical output")
}

// TestListSchemasPlaceholderContent verifies the specific content of the
// placeholder output.
func TestListSchemasPlaceholderContent(t *testing.T) {
	buf := new(bytes.Buffer)
	listSchemasCmd.SetOut(buf)
	listSchemasCmd.SetErr(buf)
	listSchemasCmd.SetArgs([]string{})

	err := listSchemasCmd.Execute()
	require.NoError(t, err, "Command should not error")

	output := buf.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Verify output structure
	require.GreaterOrEqual(t, len(lines), 5, "Output should have at least 5 lines")

	// First line should be the implementation notice
	assert.Contains(t, lines[0], "List-schemas command", "First line should mention the command")
	assert.Contains(t, lines[0], "F022", "First line should mention F022")

	// Second line should introduce available schemas
	assert.Contains(t, lines[1], "Available schemas:", "Second line should introduce schemas")

	// Verify all three schemas are listed
	schemaLines := lines[2:]
	schemaText := strings.Join(schemaLines, "\n")
	assert.Contains(t, schemaText, "fintech-loans", "Output should list fintech-loans schema")
	assert.Contains(t, schemaText, "healthcare-patients", "Output should list healthcare-patients schema")
	assert.Contains(t, schemaText, "retail-orders", "Output should list retail-orders schema")
}

// TestListSchemasCommandStructure verifies the command structure matches spec.
func TestListSchemasCommandStructure(t *testing.T) {
	// Verify Use field
	assert.Equal(t, "list-schemas", listSchemasCmd.Use, "Use field should be 'list-schemas'")

	// Verify aliases
	require.Len(t, listSchemasCmd.Aliases, 1, "Should have exactly one alias")
	assert.Equal(t, "ls", listSchemasCmd.Aliases[0], "Alias should be 'ls'")

	// Verify Short description exists and is reasonable length
	assert.NotEmpty(t, listSchemasCmd.Short, "Short description should not be empty")
	assert.Less(t, len(listSchemasCmd.Short), 100, "Short description should be concise")

	// Verify Long description exists and is longer than Short
	assert.NotEmpty(t, listSchemasCmd.Long, "Long description should not be empty")
	assert.Greater(t, len(listSchemasCmd.Long), len(listSchemasCmd.Short), "Long description should be longer than Short")

	// Verify Example exists
	assert.NotEmpty(t, listSchemasCmd.Example, "Example should not be empty")

	// Verify Run function is set
	assert.NotNil(t, listSchemasCmd.Run, "Run function should be set")
}

// TestListSchemasHelpVerbose verifies that verbose help includes all details.
func TestListSchemasHelpVerbose(t *testing.T) {
	buf := new(bytes.Buffer)
	listSchemasCmd.SetOut(buf)
	listSchemasCmd.SetErr(buf)
	listSchemasCmd.SetArgs([]string{"--help"})

	err := listSchemasCmd.Execute()
	require.NoError(t, err, "Help should not error")

	output := buf.String()

	// Should contain all major sections
	assert.Contains(t, output, "Usage:", "Help should have Usage section")
	assert.Contains(t, output, "Aliases:", "Help should have Aliases section")
	assert.Contains(t, output, "Examples:", "Help should have Examples section")

	// Long description should be present
	assert.Contains(t, output, "verticalized data schemas", "Help should show long description")

	// Global flags should be inherited and shown
	assert.Contains(t, output, "Global Flags:", "Help should show global flags section")
	assert.Contains(t, output, "--verbose", "Help should show verbose flag")
	assert.Contains(t, output, "--quiet", "Help should show quiet flag")
	assert.Contains(t, output, "--config", "Help should show config flag")
}

// TestListSchemasCommandIntegration verifies integration with root command.
func TestListSchemasCommandIntegration(t *testing.T) {
	// Verify command is properly integrated
	var listCmd *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "list-schemas" {
			listCmd = cmd
			break
		}
	}
	require.NotNil(t, listCmd, "list-schemas command should be added to root")

	// Verify it's the same command we're testing
	assert.Equal(t, listSchemasCmd, listCmd, "Registered command should be the same as the module variable")

	// Verify parent is root
	assert.Equal(t, rootCmd, listCmd.Parent(), "Parent command should be root")
}
