# F004 Quickstart: Directory Structure & Build Verification

**Feature**: F004 - Project Directory Structure & Build System
**Branch**: `002-f004-project-directory`
**Date**: 2025-10-14

## Purpose

This quickstart guide verifies that the directory structure and build system are correctly implemented. Follow these steps after implementing F004 to confirm all functionality works as expected.

---

## Prerequisites

Before beginning, ensure:
- ✅ F003 completed (Git repository and Go module initialized)
- ✅ Go 1.21+ installed (`go version` shows 1.21 or higher)
- ✅ Make installed (standard on macOS/Linux, installable on Windows via chocolatey)
- ✅ Git installed with at least one commit in repository

**Quick Check**:
```bash
go version      # Should show: go version go1.21 or higher
make --version  # Should show: GNU Make 3.81 or higher
git status      # Should show: On branch 002-f004-project-directory
```

---

## Directory Structure Overview

After F004 implementation, the repository structure should match:

```
sourcebox/
├── cmd/
│   └── sourcebox/
│       ├── main.go           # CLI entry point (minimal bootstrap logic)
│       └── main_test.go      # Placeholder test (verifies build system)
│
├── pkg/                      # Public library code (empty initially)
│   ├── generators/           # Data generation logic (added in F011-F020)
│   ├── schema/              # Schema parsing (added in F008)
│   └── database/            # Database connectors (added in F023-F024)
│
├── schemas/                  # Schema JSON definitions (added in F007)
├── docker/                   # Dockerfiles (added in F031-F036)
├── docs/                     # Documentation (added in F037)
├── examples/                 # Usage examples (added later)
│
├── dist/                     # Build artifacts (gitignored)
│   └── .gitkeep             # Ensures directory exists in git
│
├── Makefile                  # Build automation
├── .gitignore               # Includes dist/, coverage.txt
│
└── [root files from F003]
    ├── go.mod
    ├── go.sum
    ├── LICENSE
    ├── README.md
    ├── CONTRIBUTING.md
    └── CODE_OF_CONDUCT.md
```

### Directory Purposes

| Directory | Purpose | When Populated |
|-----------|---------|----------------|
| `/cmd/sourcebox/` | CLI entry point (main.go only) | F004 (now) |
| `/pkg/generators/` | Data generation logic | F011-F020 |
| `/pkg/schema/` | Schema parsing | F008 |
| `/pkg/database/` | Database connectors | F023-F024 |
| `/schemas/` | Schema JSON definitions | F007 |
| `/docker/` | Dockerfiles for databases | F031-F036 |
| `/docs/` | Docusaurus site source | F037 |
| `/examples/` | Usage examples | F006+ |
| `/dist/` | Build artifacts (gitignored) | F004 (now) |

---

## Build System Verification

### Step 1: Verify Directory Structure

```bash
# Check all directories exist
ls -la cmd/sourcebox/
ls -la pkg/
ls -la schemas/
ls -la docker/
ls -la docs/
ls -la examples/
ls -la dist/

# Verify main.go exists
cat cmd/sourcebox/main.go

# Verify placeholder test exists
cat cmd/sourcebox/main_test.go
```

**Expected**:
- All directories exist
- `main.go` contains minimal CLI bootstrap code
- `main_test.go` contains placeholder test (e.g., `TestPlaceholder`)

---

### Step 2: Verify Makefile Exists

```bash
# Check Makefile exists
cat Makefile

# Run help target (should be default)
make
```

**Expected Output** (from `make` or `make help`):
```
Available targets:
  make build      - Build for current platform → dist/sourcebox
  make test       - Run all tests with coverage
  make install    - Install binary to $GOPATH/bin
  make build-all  - Cross-compile for all 5 platforms
  make clean      - Remove dist/ and coverage files
  make help       - Show this help message
```

---

### Step 3: Build for Current Platform

```bash
# Build binary for current platform
make build

# Verify binary exists
ls -lh dist/sourcebox

# Check binary size (should be reasonable, < 100MB uncompressed)
du -h dist/sourcebox

# Test binary executes
./dist/sourcebox
```

**Expected**:
- Binary created in `dist/sourcebox` (or `dist/sourcebox.exe` on Windows)
- Binary executes without errors (may show help, version, or minimal output)
- Build completes in < 30 seconds

**Success Criteria**:
- ✅ Binary exists in dist/
- ✅ Binary is executable
- ✅ Build time < 30 seconds

---

### Step 4: Run Test Suite

```bash
# Run all tests with coverage
make test

# Verify coverage file created
cat coverage.txt
```

**Expected Output**:
```
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
=== RUN   TestPlaceholder
--- PASS: TestPlaceholder (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/jbeausoleil/sourcebox/cmd/sourcebox  0.123s
```

**Success Criteria**:
- ✅ Placeholder test passes
- ✅ Coverage report generated (coverage.txt)
- ✅ No race conditions detected

---

### Step 5: Install Binary Locally

```bash
# Install binary to $GOPATH/bin
make install

# Verify binary installed
which sourcebox

# Test installed binary
sourcebox --version
```

**Expected Output**:
```
# which sourcebox
/Users/[username]/go/bin/sourcebox

# sourcebox --version
sourcebox [version-from-git-describe]
```

**Success Criteria**:
- ✅ Binary installed to $GOPATH/bin
- ✅ `sourcebox` command available globally (open new terminal to verify)
- ✅ `sourcebox --version` shows git-derived version

---

### Step 6: Cross-Compile All Platforms

```bash
# Cross-compile for all 5 platforms (can take 1-2 minutes)
make build-all

# List all platform binaries
ls -lh dist/
```

**Expected Output**:
```
dist/
├── sourcebox                      # Current platform (uncompressed)
├── sourcebox-darwin-amd64.gz      # macOS Intel
├── sourcebox-darwin-arm64.gz      # macOS Apple Silicon
├── sourcebox-linux-amd64.gz       # Linux x86_64
├── sourcebox-linux-arm64.gz       # Linux ARM64
└── sourcebox-windows-amd64.exe.gz # Windows x86_64
```

**Verify Binary Sizes**:
```bash
# Check all compressed binaries are < 50MB
ls -lh dist/*.gz

# Should show sizes like:
# -rw-r--r--  1 user  staff   15M Oct 14 10:30 sourcebox-darwin-amd64.gz
# -rw-r--r--  1 user  staff   14M Oct 14 10:30 sourcebox-darwin-arm64.gz
# -rw-r--r--  1 user  staff   15M Oct 14 10:30 sourcebox-linux-amd64.gz
# -rw-r--r--  1 user  staff   14M Oct 14 10:30 sourcebox-linux-arm64.gz
# -rw-r--r--  1 user  staff   15M Oct 14 10:30 sourcebox-windows-amd64.exe.gz
```

**Success Criteria**:
- ✅ 5 compressed binaries created (darwin amd64/arm64, linux amd64/arm64, windows amd64)
- ✅ All .gz files < 50MB (constitutional requirement)
- ✅ Windows binary includes .exe extension before .gz
- ✅ Build time for all platforms < 2 minutes (with parallel builds: `make -j4 build-all`)

---

### Step 7: Verify Version Injection

```bash
# Check version embedded in binary
./dist/sourcebox --version

# Should show one of:
# - Tagged release: "sourcebox v1.0.0"
# - After tag: "sourcebox v1.0.0-5-g1234abcd"
# - Dirty working tree: "sourcebox v1.0.0-dirty"
# - No tags: "sourcebox 1234abcd"
# - No git: "sourcebox dev"
```

**Expected**:
- Version is automatically derived from git (no hardcoded version)
- Format follows semantic versioning or git commit hash
- If working tree is dirty, version includes "-dirty" suffix

**Success Criteria**:
- ✅ `--version` flag shows git-derived version
- ✅ Version format matches expected patterns above
- ✅ No hardcoded version strings in code

---

### Step 8: Clean Build Artifacts

```bash
# Remove all build artifacts
make clean

# Verify dist/ and coverage.txt removed
ls dist/        # Should be empty or show only .gitkeep
ls coverage.txt # Should show: No such file or directory
```

**Expected**:
- `dist/` directory is empty (or contains only .gitkeep if tracked)
- `coverage.txt` is removed
- No other files are removed (Makefile only cleans build artifacts)

**Success Criteria**:
- ✅ `make clean` removes dist/ contents
- ✅ `make clean` removes coverage.txt
- ✅ `git status` shows no leftover build artifacts

---

## Makefile Targets Reference

### Quick Reference

```bash
make             # Show help message (default target)
make build       # Build for current platform → dist/sourcebox
make test        # Run tests with race detection and coverage
make install     # Install to $GOPATH/bin
make build-all   # Cross-compile for all 5 platforms
make clean       # Remove build artifacts
make help        # Show this help message
```

### Advanced Usage

**Parallel builds** (faster cross-compilation):
```bash
make -j4 build-all  # Build 4 platforms in parallel
```

**Verbose builds** (show compiler output):
```bash
go build -v -o dist/sourcebox ./cmd/sourcebox
```

**Custom build flags** (for debugging):
```bash
# Build with debug symbols (not stripped)
go build -o dist/sourcebox-debug ./cmd/sourcebox

# Build with race detector (slow, for testing only)
go build -race -o dist/sourcebox-race ./cmd/sourcebox
```

---

## Verification Checklist

Use this checklist to confirm F004 is fully implemented:

### Directory Structure
- [ ] All directories exist (`/cmd`, `/pkg`, `/schemas`, `/docker`, `/docs`, `/examples`)
- [ ] `/cmd/sourcebox/main.go` contains minimal CLI bootstrap code
- [ ] `/cmd/sourcebox/main_test.go` contains placeholder test
- [ ] `/pkg` subdirectories created (generators, schema, database)
- [ ] `/dist` directory exists and is gitignored

### Makefile
- [ ] Makefile exists in repository root
- [ ] `make` (no args) shows help message
- [ ] All targets are declared as .PHONY
- [ ] VERSION variable uses `git describe`
- [ ] LDFLAGS include `-s -w` (strip symbols)

### Build System
- [ ] `make build` produces working binary in < 30 seconds
- [ ] `make test` runs successfully (placeholder test passes)
- [ ] `make install` installs binary to $GOPATH/bin
- [ ] `make build-all` creates 5 platform binaries
- [ ] All compressed binaries < 50MB
- [ ] `make clean` removes dist/ and coverage.txt

### Version Injection
- [ ] Binary includes `--version` flag
- [ ] Version is derived from `git describe`
- [ ] Version injection uses `-X main.version` ldflags
- [ ] Fallback to "dev" if git unavailable

### Platform Support
- [ ] macOS Intel binary created (darwin/amd64)
- [ ] macOS Apple Silicon binary created (darwin/arm64)
- [ ] Linux x86_64 binary created (linux/amd64)
- [ ] Linux ARM64 binary created (linux/arm64)
- [ ] Windows x86_64 binary created (windows/amd64 with .exe extension)

### Git Configuration
- [ ] `/dist/` is gitignored
- [ ] `coverage.txt` is gitignored
- [ ] No build artifacts in git status after `make build`

---

## Troubleshooting

### Problem: `make: command not found`

**Solution**: Install Make
```bash
# macOS (via Xcode Command Line Tools)
xcode-select --install

# macOS (via Homebrew)
brew install make

# Linux (Debian/Ubuntu)
sudo apt-get install build-essential

# Windows (via Chocolatey)
choco install make
```

### Problem: `go: command not found`

**Solution**: Install Go 1.21+
```bash
# macOS (via Homebrew)
brew install go

# Linux (via official tarball)
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Windows (via installer)
# Download from https://go.dev/dl/
```

### Problem: Binary exceeds 50MB compressed

**Cause**: External dependencies added or large assets embedded

**Solution**:
1. Check for unnecessary dependencies: `go mod graph`
2. Verify ldflags include `-s -w`: `cat Makefile | grep LDFLAGS`
3. Check binary size before compression: `ls -lh dist/sourcebox`
4. If still too large, consider using UPX compression (last resort, not standard)

### Problem: `git describe` fails with "fatal: no tags"

**Expected**: Version falls back to commit hash or "dev"

**Solution** (optional, not required for F004):
```bash
# Create initial tag (do this before first release, not now)
git tag v0.1.0
git push origin v0.1.0
```

### Problem: Cross-compilation fails for ARM platforms

**Cause**: Go toolchain missing ARM support (rare)

**Solution**:
```bash
# Reinstall Go with ARM support (usually included by default)
brew reinstall go  # macOS
```

### Problem: Windows binary missing .exe extension

**Cause**: Makefile not adding .exe suffix for Windows builds

**Solution**: Check Makefile build target for Windows:
```makefile
dist/sourcebox-windows-amd64.exe:
    GOOS=windows GOARCH=amd64 go build -o dist/sourcebox-windows-amd64.exe $(LDFLAGS) ./cmd/sourcebox
    gzip dist/sourcebox-windows-amd64.exe
```

---

## Next Steps

After verifying F004 is complete:

1. **F006**: Implement Cobra CLI framework in `/cmd/sourcebox`
   - Add CLI commands (seed, list-schemas, version, help)
   - Cobra provides flags, subcommands, help generation

2. **F008**: Create schema parser in `/pkg/schema`
   - Parse JSON schema definitions
   - Validate schema structure
   - Generate sample data from schemas

3. **F011-F020**: Build data generators in `/pkg/generators`
   - Implement verticalized data generators (fintech, healthcare, retail)
   - Use schemas from F008 to generate realistic data
   - Write TDD tests for distribution accuracy

---

## Success Criteria Summary

F004 is complete when:

- ✅ All directories exist with clear purposes
- ✅ Makefile present with 6 working targets (build, test, install, build-all, clean, help)
- ✅ `make build` produces working binary in < 30 seconds
- ✅ `make test` runs successfully (placeholder test passes)
- ✅ `make build-all` creates 5 platform binaries in < 2 minutes
- ✅ All compressed binaries < 50MB
- ✅ Version injection works (--version shows git-derived version)
- ✅ Directory structure documented (this guide or README)
- ✅ No constitutional violations (boring tech, speed, developer-first)

---

**Quickstart Status**: Ready for verification
**Next Feature**: F006 (Cobra CLI framework)
