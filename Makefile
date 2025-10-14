# SourceBox Build Configuration
# Standard Makefile for Go project build automation

# Binary configuration
BINARY_NAME = sourcebox

# Version will be injected at build time from git
VERSION = $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

# Build output directory
BUILD_DIR = dist

# Linker flags for binary optimization (strip debug symbols)
LDFLAGS = -ldflags="-s -w -X main.version=$(VERSION)"

# Target platforms for cross-compilation
# Format: OS/ARCH
PLATFORMS = \
	darwin/amd64 \
	darwin/arm64 \
	linux/amd64 \
	linux/arm64 \
	windows/amd64

# Default target - show help when user runs 'make' with no arguments
.DEFAULT_GOAL := help

##@ General

.PHONY: help
help: ## Show this help message
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*##"; printf ""} /^[a-zA-Z_-]+:.*?##/ { printf "  %-15s %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build for current platform → dist/sourcebox
	@echo "Build target not yet implemented"

.PHONY: test
test: ## Run tests with race detection and coverage
	@echo "Test target not yet implemented"

.PHONY: install
install: ## Install binary to $GOPATH/bin
	@echo "Install target not yet implemented"

.PHONY: build-all
build-all: ## Cross-compile for all 5 platforms
	@echo "Build-all target not yet implemented"

.PHONY: clean
clean: ## Remove build artifacts and coverage files
	@echo "Clean target not yet implemented"
