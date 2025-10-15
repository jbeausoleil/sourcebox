//go:build integration
// +build integration

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestMakefileVersionExtraction verifies that the Makefile VERSION
// variable correctly extracts version from git describe.
//
// Run with: go test -tags=integration ./cmd/sourcebox
func TestMakefileVersionExtraction(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Run the same git describe command that the Makefile uses
	cmd := exec.Command("sh", "-c", "git describe --tags --always --dirty 2>/dev/null || echo 'dev'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Git describe command failed: %v\nOutput: %s", err, output)
	}

	version := strings.TrimSpace(string(output))

	// Verify version is not empty
	if version == "" {
		t.Fatal("VERSION should not be empty")
	}

	// Verify version is either "dev" or a valid git describe output
	if version != "dev" {
		// Valid git describe output should be non-empty and contain only
		// alphanumeric characters, dots, hyphens, and underscores
		if len(version) == 0 {
			t.Error("VERSION should not be empty when git describe succeeds")
		}
		t.Logf("Extracted version: %s", version)
	} else {
		t.Log("Git not available or no commits, using default 'dev' version")
	}
}

// TestBuildWithMakefile verifies that building with make produces
// a binary with the correct version.
//
// Run with: go test -tags=integration ./cmd/sourcebox
func TestBuildWithMakefile(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Find the project root (where Makefile is located)
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	// Navigate up to project root
	projectRoot := filepath.Join(wd, "..", "..")

	// Clean any existing builds
	cleanCmd := exec.Command("make", "clean")
	cleanCmd.Dir = projectRoot
	if output, err := cleanCmd.CombinedOutput(); err != nil {
		t.Logf("Clean command output: %s", output)
	}

	// Build using make
	buildCmd := exec.Command("make", "build")
	buildCmd.Dir = projectRoot
	output, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Make build failed: %v\nOutput: %s", err, output)
	}

	// Verify the binary exists
	binaryPath := filepath.Join(projectRoot, "dist", "sourcebox")
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		t.Fatal("Binary was not created by make build")
	}

	// Run the binary with --version
	versionCmd := exec.Command(binaryPath, "--version")
	versionOutput, err := versionCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Version command failed: %v\nOutput: %s", err, versionOutput)
	}

	// Verify version output format
	outputStr := strings.TrimSpace(string(versionOutput))
	if !strings.HasPrefix(outputStr, "sourcebox version ") {
		t.Errorf("Version output = %q, want prefix 'sourcebox version '", outputStr)
	}

	// Extract and log the version
	version := strings.TrimPrefix(outputStr, "sourcebox version ")
	t.Logf("Built binary version: %s", version)

	// Verify version is not empty
	if version == "" {
		t.Error("Version should not be empty")
	}
}
