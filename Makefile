# SourceBox Build Configuration
# Standard Makefile for Go project build automation

# Binary configuration
BINARY_NAME = sourcebox

# Version will be injected at build time from git
VERSION = $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

# Build output directory
BUILD_DIR = dist

# Installation directory
INSTALL_DIR = $(shell go env GOPATH)/bin

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
	@echo "Building $(BINARY_NAME) for current platform..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/sourcebox
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

.PHONY: test
test: ## Run tests with race detection and coverage
	@echo "Running tests with race detection and coverage..."
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo "Tests complete. Coverage report: coverage.txt"

.PHONY: install
install: build ## Install binary to $GOPATH/bin
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@mkdir -p $(INSTALL_DIR)
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installation complete: $(INSTALL_DIR)/$(BINARY_NAME)"
	@echo "Ensure $(INSTALL_DIR) is in your PATH to use '$(BINARY_NAME)' globally"

.PHONY: build-all
build-all: ## Cross-compile for all 5 platforms
	@echo "Build-all target not yet implemented"

.PHONY: clean
clean: ## Remove build artifacts and coverage files
	@echo "Clean target not yet implemented"
