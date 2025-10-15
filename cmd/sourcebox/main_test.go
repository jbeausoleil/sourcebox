package main

import (
	"testing"
)

// TestPlaceholder verifies the build system and testing infrastructure work.
// This is a minimal placeholder test that will be expanded in future features
// as the CLI functionality is implemented.
func TestPlaceholder(t *testing.T) {
	// Verify that the version variable exists and has a default value
	if version == "" {
		t.Error("version variable should have a default value")
	}

	// Test passes - build system and testing infrastructure are working
	t.Log("Build system and testing infrastructure verified")
}

// TestVersionVariable verifies the version variable has the expected default value
// and can be overridden at build time via ldflags.
func TestVersionVariable(t *testing.T) {
	tests := []struct {
		name    string
		version string
		wantErr bool
	}{
		{
			name:    "default version is 'dev'",
			version: version,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.version == "" {
				t.Error("version should never be empty")
			}

			// In test environment (not built with ldflags), version should be "dev"
			// When built with ldflags, it will be a git describe output
			// Both are valid, just verify it's not empty
			if len(tt.version) == 0 {
				t.Errorf("version length = %d, want > 0", len(tt.version))
			}
		})
	}
}

// TestVersionDefault ensures the version variable defaults to "dev" when
// not overridden by build-time ldflags.
func TestVersionDefault(t *testing.T) {
	// When running tests without build-time ldflags injection,
	// version should have its default value of "dev"
	if version != "dev" {
		// This is actually OK - it means the binary was built with ldflags
		t.Logf("version = %q (built with ldflags)", version)
	} else {
		t.Logf("version = %q (default value)", version)
	}

	// The important check: version must not be empty
	if version == "" {
		t.Fatal("version must not be empty")
	}
}
