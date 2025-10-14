# Repository Guidelines

## Project Structure & Module Organization
- Module root is `github.com/jbeausoleil/sourcebox`; add CLIs under `cmd/<tool>` and shared code under `internal/` or `pkg/` when implementation starts.
- `specs/` houses research, plans, and acceptance criteria; `docs/` stores session summaries; `internal-docs/` contains roadmaps. Extend these sources instead of scattering new context elsewhere.
- Keep generated datasets out of version control; use lightweight fixtures inside `testdata/` when examples are unavoidable.

## Build, Test, and Development Commands
- `gofmt -w .` — required formatter; run it before every commit.
- `go vet ./...` — lint pass to keep static analysis noise out of PRs.
- `go test ./...` (optionally `-race` or `-coverprofile=coverage.out`) — execute tests, race checks, and review coverage via `go tool cover -html=coverage.out`.
- `go mod tidy` — sync `go.mod` and `go.sum` after dependency changes.

## Coding Style & Naming Conventions
- Stick to idiomatic Go: tabs, sub-100-character lines, and zero manual formatting.
- Apply `MixedCaps` naming; exported items start uppercase, single-method interfaces use the `-er` suffix.
- Document exported packages and functions with godoc comments and favor small, purpose-specific packages over grab-bag utilities.

## Testing Guidelines
- Prefer table-driven tests in colocated `*_test.go` files; subtests (`t.Run`) keep scenarios readable.
- Aim for ≥80% coverage on core packages and never dip the repo below 60%; note exceptions in the PR body.
- Capture manual QA notes, platform checks, and golden-file references in `docs/session-summary/` so future contributors understand the exercised surface area.

## Commit & Pull Request Guidelines
- Follow the `type(scope): summary` message format described in `CONTRIBUTING.md`; keep changesets atomic.
- Before opening a PR, rebase on `main`, rerun `gofmt`, `go vet`, and `go test`, and list the commands executed in the description.
- Link issues with `Fixes #123`, attach user-facing evidence when applicable, and point reviewers to the relevant spec (e.g., `specs/001-f003-initialize-git/plan.md`) for context.
