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

##@ Build Targets

.PHONY: build
build: ## Build for current platform → dist/sourcebox
	@echo "Building $(BINARY_NAME) for current platform..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/sourcebox
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

.PHONY: build-darwin-amd64
build-darwin-amd64: ## Build for macOS Intel (darwin/amd64)
	@echo "Building $(BINARY_NAME) for macOS Intel (darwin/amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/sourcebox
	@echo "Compressing binary..."
	@gzip -f -9 $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64.gz"

.PHONY: build-darwin-arm64
build-darwin-arm64: ## Build for macOS Apple Silicon (darwin/arm64)
	@echo "Building $(BINARY_NAME) for macOS Apple Silicon (darwin/arm64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/sourcebox
	@echo "Compressing binary..."
	@gzip -f -9 $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64.gz"

.PHONY: build-linux-amd64
build-linux-amd64: ## Build for Linux x86_64 (linux/amd64)
	@echo "Building $(BINARY_NAME) for Linux x86_64 (linux/amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/sourcebox
	@echo "Compressing binary..."
	@gzip -f -9 $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64.gz"

.PHONY: build-linux-arm64
build-linux-arm64: ## Build for Linux ARM64 (linux/arm64)
	@echo "Building $(BINARY_NAME) for Linux ARM64 (linux/arm64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/sourcebox
	@echo "Compressing binary..."
	@gzip -f -9 $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64.gz"

.PHONY: build-windows-amd64
build-windows-amd64: ## Build for Windows x86_64 (windows/amd64)
	@echo "Building $(BINARY_NAME) for Windows x86_64 (windows/amd64)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/sourcebox
	@echo "Compressing binary..."
	@gzip -f -9 $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe.gz"

.PHONY: build-all
build-all: ## Cross-compile for all 5 platforms with compression
	@echo "========================================"
	@echo "Building $(BINARY_NAME) for all platforms"
	@echo "Version: $(VERSION)"
	@echo "========================================"
	@echo ""
	@$(MAKE) build-darwin-amd64
	@echo ""
	@$(MAKE) build-darwin-arm64
	@echo ""
	@$(MAKE) build-linux-amd64
	@echo ""
	@$(MAKE) build-linux-arm64
	@echo ""
	@$(MAKE) build-windows-amd64
	@echo ""
	@echo "========================================"
	@echo "Verifying compressed binary sizes..."
	@echo "========================================"
	@$(MAKE) verify-sizes
	@echo ""
	@echo "========================================"
	@echo "Build complete! All binaries:"
	@echo "========================================"
	@ls -lh $(BUILD_DIR)/*.gz | awk '{printf "  %s  %s\n", $$5, $$9}'

.PHONY: verify-sizes
verify-sizes: ## Verify all compressed binaries are under 50MB
	@echo "Checking binary sizes (must be < 50MB compressed)..."
	@for file in $(BUILD_DIR)/*.gz; do \
		if [ -f "$$file" ]; then \
			size=$$(stat -f%z "$$file" 2>/dev/null || stat -c%s "$$file" 2>/dev/null); \
			size_mb=$$(echo "scale=2; $$size / 1048576" | bc); \
			filename=$$(basename "$$file"); \
			if [ "$$size" -gt 52428800 ]; then \
				echo "  ERROR: $$filename is $$size_mb MB (exceeds 50MB limit)"; \
				exit 1; \
			else \
				echo "  OK: $$filename is $$size_mb MB"; \
			fi; \
		fi; \
	done

##@ Development

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

.PHONY: clean
clean: ## Remove build artifacts and coverage files
	@echo "Cleaning build artifacts and coverage files..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.txt
	@echo "Clean complete"
