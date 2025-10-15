package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// version will be set at build time via ldflags
	version = "dev"
)

func main() {
	// Define flags
	versionFlag := flag.Bool("version", false, "print version information")
	flag.BoolVar(versionFlag, "v", false, "print version information (shorthand)")

	// Parse flags
	flag.Parse()

	// Handle version flag
	if *versionFlag {
		fmt.Printf("sourcebox version %s\n", version)
		os.Exit(0)
	}

	// Placeholder output - will be replaced when pkg/ modules are implemented
	fmt.Println("SourceBox - Mock Data Generation Tool")
	fmt.Println("Coming soon: Run 'sourcebox --version' for version info")
	fmt.Println()
	fmt.Println("This is a placeholder. Core functionality will be implemented in upcoming features.")
}
