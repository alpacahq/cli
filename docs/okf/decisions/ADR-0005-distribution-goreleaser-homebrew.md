---
type: Architecture Decision
title: ADR-0005 Distribute as prebuilt binaries via GoReleaser, a Homebrew tap and go install
description: Tagged releases are cross-compiled and published by GoReleaser with a Homebrew tap formula; the CLI detects its install method and offers the matching upgrade command.
status: accepted
date: unknown
deciders: []
supersedes: []
affects: [cli, cli.alpaca, ext_github]
allium: []
evidence: [.goreleaser.yaml, .github/workflows/release.yml, Makefile, internal/cmd/update.go, internal/cmd/update_check.go, README.md]
tags: [alpacahq, adr, generated]
timestamp: 2026-09-12T00:00:00+00:00
generated_by: claude-opus-5 / layered-docs 2026-09
source_commit: 53606273aa230a40c64b783425dcb3f4423ede30
source_branch: docs/layered-2026-09
generated_at: 2026-09-12T00:00:00+00:00
confidence: medium
review_status: draft-needs-review
---

# ADR-0005 Distribute as prebuilt binaries via GoReleaser, a Homebrew tap and go install

## Context

The CLI is a single static binary intended to be installed by developers, scripts and agents on
several operating systems. The README offers exactly two install routes, `go install` and Homebrew
(`README.md`). Because the project is pre-1.0 and releases may change behaviour, users and agents
need a reliable way to learn that a newer version exists and how to upgrade the copy they actually
have (`README.md`, `AGENTS.md`).

## Decision

Releases are driven by version tags. A release workflow runs GoReleaser, which cross-compiles for
Linux, macOS and Windows on amd64 and arm64 with CGO disabled, stamps the version into the binary
at link time, publishes archives (zip on Windows, tar.gz elsewhere) with a checksum file and a
filtered changelog, and pushes a formula to the organisation Homebrew tap
(`.github/workflows/release.yml`, `.goreleaser.yaml`).

At runtime the CLI queries the latest published release to decide whether an update is available,
detects whether it was installed by Homebrew or `go install`, and offers the matching upgrade
command (`internal/cmd/update.go`, `internal/cmd/update_check.go`).

## Consequences

- The release workflow is tag-triggered and is the only job with write permission to repository contents (`.github/workflows/release.yml`).
- Tap publication is conditional: a short-lived app token is minted only when the tap application identifiers are configured, and the formula upload is skipped when no tap token is present, so a release still succeeds without it (`.github/workflows/release.yml`, `.goreleaser.yaml`).
- The published formula installs the binary and generates shell completions from the executable, and its smoke test runs the version command — which is why shell completion and a cheap `version` command are part of the contract (`.goreleaser.yaml`, `internal/cmd/root.go`).
- A `make release` helper computes the next patch tag from existing tags, refuses to run on a dirty working tree, prompts for confirmation and pushes the tag (`Makefile`).
- The update check is a plain HTTPS request to the release metadata of this repository, sent with the CLI's own user agent and a short timeout; it stays silent on error and when already current, so the bare `alpaca` invocation does not become noisy (`internal/cmd/update.go`).
- Version comparison is a numeric major, minor and patch comparison that strips a leading "v" and any pre-release suffix (`internal/cmd/update_check.go`).
- Install-method detection is heuristic: it resolves the executable path through symlinks and looks for Homebrew path markers, then the Go bin and Go path locations, defaulting to the `go install` upgrade command. A binary placed somewhere unusual will be told to upgrade with `go install` (`internal/cmd/update_check.go`).
- The self-upgrade shells out to the same command a user would type, so that the command shown and the command run are identical (`internal/cmd/update.go`).

## Alternatives considered

No alternative distribution mechanism is documented in the repository. The in-place self-replacing
binary update that some CLIs use is implicitly not adopted, since the upgrade path delegates to the
detected package manager instead. Confidence: low on that reading — it is inferred from the
implementation, not stated.
