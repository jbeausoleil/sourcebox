# F004 Research: Project Directory Structure & Build System

**Feature Branch**: `002-f004-project-directory`
**Research Date**: 2025-10-14
**Status**: Complete

## Overview

This document captures all technical research and decisions for F004 - Project Directory Structure & Build System. All decisions prioritize boring tech, speed, and developer-first design as mandated by the constitution.

---

## Research Decision 1: Go Project Layout Standards

### Decision Point
Which directories are required for a Go CLI project? What goes in /cmd vs /pkg vs /internal?

### Decision
Adopt the **golang-standards/project-layout** pattern with the following structure:

```
sourcebox/
├── cmd/sourcebox/        # CLI entry point (main.go only)
├── pkg/                  # Public library code
│   ├── generators/       # Data generation logic (F011-F020)
│   ├── schema/          # Schema parsing (F008)
│   └── database/        # Database connectors (F023-F024)
├── schemas/             # Schema JSON definitions (F007)
├── docker/              # Dockerfiles (F031-F036)
├── docs/                # Documentation (F037)
├── examples/            # Usage examples
└── dist/                # Build artifacts (gitignored)
```

### Rationale
- **golang-standards/project-layout** is the de facto community standard for Go projects
- `/cmd/sourcebox/` contains only main.go with minimal bootstrap logic (aligns with Developer-First principle)
- `/pkg/` contains all reusable library code organized by domain (supports future external use)
- **NOT using /internal/** initially (Boring Tech principle: don't add structure until needed)
- `/schemas`, `/docker`, `/docs`, `/examples` are domain-specific to SourceBox (clear purpose)
- `/dist` for build artifacts keeps root clean (Developer-First: predictable locations)

### Alternatives Considered
1. **Flat structure** (src/ with everything mixed) → Rejected: doesn't scale, hard to navigate
2. **Monorepo with multiple modules** → Rejected: over-engineering for single-binary CLI
3. **Nest everything under /internal/** → Rejected: premature optimization, add when external packages exist

### Source
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout) (public repository, 48K+ stars)
- Popular Go CLI projects: cobra, hugo, kubernetes CLI (public repositories)
- Go official documentation on project structure (public documentation)

### Constitutional Alignment
- ✅ **Boring Tech Wins**: Standard Go layout, widely documented, proven
- ✅ **Developer-First Design**: Clear structure, predictable file locations
- ✅ **Ship Fast**: Simple structure enables rapid development

---

## Research Decision 2: Makefile Structure

### Decision Point
What targets are essential for a Go CLI project? What variables? What PHONY declarations?

### Decision
Create a Makefile with the following structure:

**Variables**:
```makefile
BINARY_NAME=sourcebox
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DIR=dist
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION)"
PLATFORMS=darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64
```

**Targets** (all declared as .PHONY):
- `build` - Build for current platform → dist/sourcebox
- `test` - Run all tests with race detection and coverage
- `install` - Install binary to $GOPATH/bin
- `build-all` - Cross-compile for all 5 platforms
- `clean` - Remove dist/ and coverage files
- `help` - Show available targets (set as .DEFAULT_GOAL)

**PHONY Declaration**:
```makefile
.PHONY: build test install build-all clean help
.DEFAULT_GOAL := help
```

### Rationale
- **Simple, predictable targets**: Developers can guess `make build`, `make test`, `make install` (Developer-First)
- **Version from git**: `git describe` provides semantic versioning without manual updates (Speed > Features)
- **LDFLAGS optimization**: `-s -w` strips debug symbols to reduce binary size (Performance constraint: < 50MB)
- **PHONY declarations**: Prevents conflicts with files named "build", "test", etc. (standard practice)
- **help as default**: `make` shows available targets without needing to read Makefile (Developer-First)

### Alternatives Considered
1. **Task/Just instead of Make** → Rejected: Make is boring, standard, universally available (Boring Tech principle)
2. **Shell scripts instead of Makefile** → Rejected: Make handles dependencies and parallel builds better
3. **Complex build system (Bazel, Buck)** → Rejected: massive over-engineering for single-binary CLI

### Source
- Common Makefile patterns from Go CLI projects: cobra, hugo, docker CLI (public repositories)
- Go build documentation (public documentation)
- Make manual and best practices (public documentation)

### Constitutional Alignment
- ✅ **Boring Tech Wins**: Makefile is standard, proven, available on all systems
- ✅ **Speed > Features**: Simple targets, no complex build logic
- ✅ **Developer-First Design**: Predictable commands, help target

---

## Research Decision 3: Cross-Compilation Configuration

### Decision Point
Exact GOOS/GOARCH values for 5 platforms? Output naming convention?

### Decision
**Platform Matrix**:

| Platform | GOOS | GOARCH | Output Filename |
|----------|------|--------|----------------|
| macOS Intel | darwin | amd64 | sourcebox-darwin-amd64.gz |
| macOS Apple Silicon | darwin | arm64 | sourcebox-darwin-arm64.gz |
| Linux x86_64 | linux | amd64 | sourcebox-linux-amd64.gz |
| Linux ARM64 | linux | arm64 | sourcebox-linux-arm64.gz |
| Windows x86_64 | windows | amd64 | sourcebox-windows-amd64.exe.gz |

**Build Command Pattern**:
```makefile
GOOS=darwin GOARCH=amd64 go build -o dist/sourcebox-darwin-amd64 $(LDFLAGS) ./cmd/sourcebox
gzip dist/sourcebox-darwin-amd64
```

**Windows Special Handling**:
- Binary name MUST include `.exe` extension before compression: `sourcebox-windows-amd64.exe`
- After gzip: `sourcebox-windows-amd64.exe.gz`

### Rationale
- **5 platforms cover 95%+ of developers**: macOS (Intel + ARM), Linux (x86 + ARM), Windows (x86)
- **Naming convention includes platform**: Clear which binary to download (Developer-First)
- **Windows .exe extension**: Required for execution on Windows (platform-specific handling)
- **gzip compression**: Standard, widely supported, reduces binary size to meet < 50MB requirement

### Alternatives Considered
1. **Support 32-bit platforms** → Rejected: obsolete, <1% of developers use 32-bit systems
2. **Support FreeBSD, OpenBSD** → Rejected: niche platforms, add after MVP validation
3. **Use tar.gz instead of gzip** → Rejected: single binary doesn't need tar, gzip is simpler

### Source
- Go cross-compilation documentation (official public docs)
- Platform naming standards from Docker, Homebrew (public repositories)
- GitHub release asset naming patterns (public repositories)

### Constitutional Alignment
- ✅ **Platform Support**: All 5 platforms (macOS Intel/ARM, Linux x86/ARM, Windows x86)
- ✅ **Developer-First Design**: Clear naming, predictable output
- ✅ **Performance**: gzip meets < 50MB compressed requirement

---

## Research Decision 4: Binary Optimization Strategy

### Decision Point
Which ldflags reduce binary size most? Compression algorithm? Target size?

### Decision
**Optimization Strategy**:

1. **Compiler Flags** (`-ldflags`):
   - `-s`: Strip debug symbols (reduces symbol table size)
   - `-w`: Strip DWARF debug information (removes debugging metadata)
   - Combined: `-ldflags="-s -w"`

2. **Compression**:
   - Algorithm: gzip (standard, widely available)
   - Compression level: default (gzip -9 for maximum compression if needed)

3. **Size Targets**:
   - Uncompressed binary: < 100MB (typical Go binary with stdlib)
   - Compressed binary: < 50MB (constitutional requirement)
   - Expected reduction: 50-70% compression ratio with gzip

**Build Command**:
```makefile
go build -ldflags="-s -w -X main.version=$(VERSION)" -o dist/sourcebox ./cmd/sourcebox
gzip dist/sourcebox
```

### Rationale
- **`-s -w` flags**: Remove debug information not needed for production binaries (standard Go practice)
- **gzip compression**: Standard algorithm, available on all platforms, good compression ratio
- **Version injection via -X**: Embeds git version at compile time without runtime overhead
- **Tradeoff**: Stripped binaries can't be debugged, but users can rebuild with debug symbols if needed

### Alternatives Considered
1. **UPX compression** → Rejected: not standard, antivirus false positives, exotic (violates Boring Tech)
2. **brotli/zstd compression** → Rejected: not universally available, adds complexity
3. **No compression** → Rejected: binaries exceed 50MB requirement

### Source
- Go linker documentation (`go tool link`) (official public docs)
- Go binary size optimization guides (public blog posts, Stack Overflow)
- gzip compression benchmarks (public data)

### Constitutional Alignment
- ✅ **Performance**: Achieves < 50MB compressed (technical constraint #1)
- ✅ **Boring Tech Wins**: Standard flags, standard compression algorithm
- ✅ **Speed > Features**: No exotic optimizations that add complexity

---

## Research Decision 5: Version Injection Mechanism

### Decision Point
How to embed version? Format from git? Fallback for dirty builds?

### Decision
**Version Extraction**:
```makefile
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
```

**Version Injection**:
```makefile
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION)"
```

**Version Display** (in main.go):
```go
var version = "dev" // Overridden at build time

func main() {
    if len(os.Args) > 1 && os.Args[1] == "--version" {
        fmt.Printf("sourcebox %s\n", version)
        os.Exit(0)
    }
    // ... rest of CLI logic
}
```

**Version Formats**:
- Tagged release: `v1.0.0` (from `git tag v1.0.0`)
- Commits after tag: `v1.0.0-5-g1234abcd` (5 commits past v1.0.0, commit hash 1234abcd)
- Dirty working tree: `v1.0.0-dirty` (uncommitted changes)
- No tags: `1234abcd` (commit hash only)
- No git: `dev` (fallback for non-git builds)

### Rationale
- **`git describe`**: Standard way to generate semantic versions from git tags
- **`--tags`**: Include all tags (annotated and lightweight)
- **`--always`**: Show commit hash if no tags exist (prevents empty version)
- **`--dirty`**: Append "-dirty" if working tree has uncommitted changes (transparency)
- **Fallback to "dev"**: Handles builds without git (e.g., downloaded source zip)
- **`-X main.version`**: Injects version string at compile time (no runtime overhead)

### Alternatives Considered
1. **Hardcoded version in code** → Rejected: manual updates are error-prone, not automated
2. **Read version from VERSION file** → Rejected: requires maintaining separate file, not DRY
3. **Runtime `git describe` execution** → Rejected: requires git at runtime, slow, brittle

### Source
- `git describe` documentation (official git docs)
- Go linker `-X` flag documentation (official Go docs)
- Semantic versioning specification (public semver.org)

### Constitutional Alignment
- ✅ **Developer-First Design**: Version automatically derived from git tags, no manual updates
- ✅ **Speed > Features**: Compile-time injection, no runtime overhead
- ✅ **Boring Tech Wins**: Standard git describe pattern used by many Go projects

---

## Research Decision 6: Build Directory Organization

### Decision Point
How to organize /dist for multiple platforms? Naming convention?

### Decision
**Directory Structure**:
```
dist/
├── sourcebox                      # Current platform (uncompressed)
├── sourcebox-darwin-amd64.gz      # macOS Intel (compressed)
├── sourcebox-darwin-arm64.gz      # macOS Apple Silicon (compressed)
├── sourcebox-linux-amd64.gz       # Linux x86_64 (compressed)
├── sourcebox-linux-arm64.gz       # Linux ARM64 (compressed)
└── sourcebox-windows-amd64.exe.gz # Windows x86_64 (compressed)
```

**Naming Convention**:
- Pattern: `{binary}-{os}-{arch}.gz` (compressed distribution binaries)
- Pattern: `{binary}.exe` for Windows (platform-specific extension)
- Pattern: `{binary}` (current platform, uncompressed, created by `make build`)

**Makefile Organization**:
```makefile
# Single platform (make build)
dist/sourcebox: cmd/sourcebox/main.go
    go build -o dist/sourcebox $(LDFLAGS) ./cmd/sourcebox

# All platforms (make build-all)
build-all: dist/sourcebox-darwin-amd64.gz dist/sourcebox-darwin-arm64.gz ...

dist/sourcebox-darwin-amd64.gz:
    GOOS=darwin GOARCH=amd64 go build -o dist/sourcebox-darwin-amd64 $(LDFLAGS) ./cmd/sourcebox
    gzip dist/sourcebox-darwin-amd64
```

### Rationale
- **Single /dist directory**: All build artifacts in one place, easy to gitignore, easy to clean
- **Platform in filename**: Developers can identify correct binary without extracting
- **Compressed for distribution**: Meets < 50MB requirement, standard for releases
- **Uncompressed for local dev**: `make build` creates uncompressed binary for fast iteration
- **Make targets as files**: Enables incremental builds (Make rebuilds only if source changed)

### Alternatives Considered
1. **Separate directories per platform** (dist/darwin/, dist/linux/) → Rejected: unnecessary hierarchy
2. **Store in release/ instead of dist/** → Rejected: "dist" is more standard in Go projects
3. **Include version in filename** (sourcebox-v1.0.0-darwin-amd64.gz) → Rejected: complicates automation, version in metadata is sufficient

### Source
- Go project conventions from kubernetes, docker, hugo (public repositories)
- GitHub release asset naming patterns (public repositories)
- Makefile best practices (public documentation)

### Constitutional Alignment
- ✅ **Developer-First Design**: Clear naming, predictable structure, easy cleanup
- ✅ **Speed > Features**: Simple flat structure, no unnecessary hierarchy
- ✅ **Boring Tech Wins**: Standard /dist directory, standard naming patterns

---

## Research Decision 7: Platform-Specific Considerations

### Decision Point
.exe extension for Windows? ARM64 compatibility testing approach?

### Decision
**Windows Handling**:
- Binary MUST have `.exe` extension: `sourcebox.exe`, `sourcebox-windows-amd64.exe`
- Build command: `go build -o dist/sourcebox-windows-amd64.exe $(LDFLAGS) ./cmd/sourcebox`
- Compression: `gzip dist/sourcebox-windows-amd64.exe` → `sourcebox-windows-amd64.exe.gz`
- Execution: Windows requires `.exe` extension, Go compiler doesn't add it automatically

**ARM64 Handling**:
- Cross-compilation: Works on any platform (Go toolchain supports all targets)
- Testing: **Manual QA required on ARM hardware** (M1/M2 Mac, ARM Linux server, Raspberry Pi)
- CI/CD: GitHub Actions supports ARM runners, but manual testing MUST happen before releases
- Fallback: If ARM binaries don't work, users can build from source on ARM platforms

**Platform-Specific Gotchas**:
- **Windows**: Path separators (`\` vs `/`), line endings (CRLF vs LF) → Go stdlib handles automatically
- **macOS**: Code signing NOT required for CLI tools (only GUI apps)
- **Linux**: glibc vs musl → Go produces static binaries, no dependency issues

### Rationale
- **Windows .exe requirement**: Windows won't execute files without .exe extension (OS constraint)
- **ARM testing constraint**: Can't fully test ARM binaries without ARM hardware (manual QA required)
- **Go cross-compilation**: Works reliably for all platforms without platform-specific build machines
- **Static binaries**: Go produces self-contained binaries, no dynamic linking issues

### Alternatives Considered
1. **Skip ARM support initially** → Rejected: ARM is growing (Apple Silicon, cloud ARM instances)
2. **Use Docker to test ARM binaries** → Rejected: Docker ARM emulation is slow and unreliable
3. **Add Windows code signing** → Rejected: costs money ($200+/year), not required for CLI tools

### Source
- Go cross-compilation documentation (official public docs)
- Windows executable requirements (public Microsoft docs)
- ARM64 testing best practices (public blog posts, Stack Overflow)

### Constitutional Alignment
- ✅ **Platform Support**: All 5 platforms (macOS Intel/ARM, Linux x86/ARM, Windows x86)
- ✅ **Manual QA Before Every Release**: ARM testing on real hardware (development practice #3)
- ✅ **Boring Tech Wins**: Standard Go cross-compilation, no exotic tooling

---

## Research Decision 8: Build Performance Optimization

### Decision Point
Can builds run in parallel? Caching strategies?

### Decision
**Parallel Builds**:
```makefile
build-all:
    @echo "Building for all platforms (can run with make -j4 build-all for parallel builds)"
    $(MAKE) dist/sourcebox-darwin-amd64.gz
    $(MAKE) dist/sourcebox-darwin-arm64.gz
    $(MAKE) dist/sourcebox-linux-amd64.gz
    $(MAKE) dist/sourcebox-linux-arm64.gz
    $(MAKE) dist/sourcebox-windows-amd64.exe.gz
```

**Usage**: `make -j4 build-all` (4 parallel builds)

**Caching Strategy**:
1. **Go build cache**: Enabled by default, stores compiled packages in `$GOCACHE`
2. **Make incremental builds**: Only rebuild if source files changed (dependency tracking)
3. **No custom caching**: Rely on Go's built-in cache, don't add complexity

**Performance Targets**:
- Single platform (`make build`): < 30 seconds (typical: 5-10 seconds on 2020 MacBook Pro)
- All platforms (`make -j4 build-all`): < 2 minutes (typical: 30-60 seconds with parallelization)
- Incremental rebuild: < 5 seconds (if only one file changed)

**Optimization Techniques**:
- **Parallel compilation**: Go compiler uses all CPU cores by default (`GOMAXPROCS`)
- **Minimal dependencies**: Only standard library for now (no external deps to compile)
- **Small codebase**: MVP has minimal code, build times naturally fast

### Rationale
- **Make's -j flag**: Built-in parallelization, no custom scripting needed (Boring Tech)
- **Go build cache**: Automatic caching of compiled packages, works across builds
- **No custom caching**: Premature optimization, Go's cache is sufficient for MVP
- **Performance targets**: Meet constitutional requirement (< 2 minutes per platform)

### Alternatives Considered
1. **Custom build cache (Bazel, Buck)** → Rejected: massive over-engineering for single-binary CLI
2. **Docker build caching** → Rejected: not using Docker for builds (local-first)
3. **Pre-compiled dependencies** → Rejected: no external dependencies in MVP

### Source
- Make parallel execution documentation (public Make manual)
- Go build cache documentation (official Go docs)
- Build performance best practices from large Go projects (public repositories)

### Constitutional Alignment
- ✅ **Speed > Features**: Build time < 2 minutes (technical constraint #1)
- ✅ **Boring Tech Wins**: Standard Make parallelization, standard Go caching
- ✅ **Ship Fast, Validate Early**: Fast builds enable rapid iteration

---

## Summary of Research Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| **1. Go Project Layout** | golang-standards/project-layout | Standard, proven, widely documented |
| **2. Makefile Structure** | build, test, install, build-all, clean, help | Simple, predictable, developer-friendly |
| **3. Cross-Compilation** | 5 platforms (macOS, Linux, Windows) | Covers 95%+ of developers |
| **4. Binary Optimization** | -ldflags="-s -w", gzip compression | Meets < 50MB requirement |
| **5. Version Injection** | git describe + ldflags -X | Automated, no manual updates |
| **6. Build Directory** | /dist with platform in filename | Clear, predictable, easy to clean |
| **7. Platform-Specific** | .exe for Windows, ARM manual QA | Platform requirements, manual testing |
| **8. Build Performance** | Make -j4, Go build cache | < 2 minute all-platform builds |

## Next Steps (Phase 1)

- ✅ Research complete (this document)
- ⏭️ Generate quickstart.md with build verification steps
- ⏭️ Update CLAUDE.md with build system context
- ⏭️ Skip data-model.md (no data entities in this feature)
- ⏭️ Skip contracts/ (no APIs in this feature)

## Constitutional Compliance Verification

### Core Principles
- ✅ **Verticalized > Generic**: N/A (infrastructure setup)
- ✅ **Speed > Features**: Build time < 2 minutes, optimized flags
- ✅ **Local-First**: Build system works 100% offline
- ✅ **Boring Tech Wins**: Standard Go layout, Makefile, gzip
- ✅ **Open Source Forever**: N/A (no licensing changes)
- ✅ **Developer-First Design**: Clear structure, simple commands
- ✅ **Ship Fast**: Simple structure enables rapid development

### Technical Constraints
- ✅ **Performance**: Build < 2 minutes, binary < 50MB compressed
- ✅ **Distribution**: Supports all channels (npm, homebrew, Docker)
- ✅ **Database Support**: N/A (directory structure only)
- ✅ **Cost**: $0 (local build tooling)
- ✅ **Code Quality**: Standard Go layout supports vet, lint, TDD
- ✅ **License**: N/A (no changes)
- ✅ **Platform Support**: All 5 platforms (macOS, Linux, Windows)

### Legal Constraints
- ✅ **Independent Development**: Uses standard public tools only
- ✅ **No Employer References**: N/A (infrastructure)
- ✅ **Public Information Only**: All patterns from public standards
- ✅ **Open Source Protection**: N/A (no licensing changes)
- ✅ **Illustrative Examples Only**: N/A (no company references)

### Anti-Patterns Avoided
- ✅ **No Shiny Tech**: Make, Go, gzip (all boring and standard)
- ✅ **No Over-Engineering**: Simple Makefile, standard layout
- ✅ **No Premature Optimization**: Standard caching, no custom systems
- ✅ **No Complex Build System**: Makefile only, no Bazel/Buck

---

**Research Status**: ✅ Complete
**Next Phase**: Generate quickstart.md (Phase 1)
