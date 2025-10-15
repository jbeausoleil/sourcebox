/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/jbeausoleil/sourcebox/cmd/sourcebox/cmd"

var (
	// version will be set at build time via ldflags
	version = "dev"
)

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
