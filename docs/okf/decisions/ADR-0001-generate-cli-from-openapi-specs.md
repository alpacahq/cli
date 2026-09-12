---
type: Architecture Decision
title: ADR-0001 Generate the CLI from the Alpaca OpenAPI specifications
description: The command surface, typed clients and response schemas are generated from vendored OpenAPI documents rather than hand-written.
status: accepted
date: unknown
deciders: []
supersedes: []
affects: [cli, cli.generator, cli.specs, cli.alpaca]
allium: []
evidence: [AGENTS.md, Makefile, cmd/generate/main.go, cmd/generate/commands.go, internal/api/trading_client.gen.go, internal/api/marketdata_client.gen.go, internal/cmd/commands.gen.go, .github/workflows/ci.yml, README.md]
tags: [alpacahq, adr, generated]
timestamp: 2026-09-12T00:00:00+00:00
generated_by: claude-opus-5 / layered-docs 2026-09
source_commit: 53606273aa230a40c64b783425dcb3f4423ede30
source_branch: docs/layered-2026-09
generated_at: 2026-09-12T00:00:00+00:00
confidence: medium
review_status: draft-needs-review
---

# ADR-0001 Generate the CLI from the Alpaca OpenAPI specifications

## Context

The CLI wraps two large, evolving HTTP APIs. `AGENTS.md` states the core principle directly:
"The CLI is driven by OpenAPI specs. Maximize what's generated, minimize what's hand-written.
**Do not edit generated files directly.**" It also records that the specs are read-only inputs:
"never edit the specs in this repo. Fix bugs upstream and re-import." The README repeats the
consequence for users: "The CLI is generated from Alpaca OpenAPI specs, so the installed binary is
the source of truth for commands, flags, enum completions, validation, and response schemas"
(`README.md`).

## Decision

The Alpaca trading and market-data OpenAPI documents are vendored at `api/specs/trading-api.json`
and `api/specs/market-data-api.json` and are treated as read-only inputs. A generator in
`cmd/generate/` reads them and emits the typed API clients and command metadata
(`internal/api/*.gen.go`, `internal/cmd/commands.gen.go`). Changes to the command surface are made
by updating the specs or the generator and running `make generate`, never by editing generated
files (`AGENTS.md`, `Makefile`).

## Consequences

- The generator is a first-class build step: `make generate` runs `go run ./cmd/generate` (`Makefile`).
- Drift is enforced in CI. The build job runs `make generate` and then `git diff --exit-code internal/api/ internal/cmd/commands.gen.go`, failing with a message telling the author to regenerate and commit (`.github/workflows/ci.yml`).
- Spec refresh is a scripted operation: `make spec-update` fetches the published documents and `make spec-check` reports the changelog between the vendored copy and the published one using `oasdiff` (`github.com/oasdiff/oasdiff`), a third-party tool the target requires and whose absence aborts it with an install instruction (`Makefile`).
- Response schemas and enum completions are available offline because they are compiled into the binary as Go source, not read from the spec files at runtime: `internal/api/descriptions.gen.go` carries inline completion lists, and no `go:embed` or runtime read of `api/specs` exists. This is what lets `alpaca ... --schema` print response fields without an API call and lets shell completion offer valid enum values (`internal/api/descriptions.gen.go`, `internal/cmd/root.go`, `README.md`).
- A deliberate mapping wart is documented rather than removed: flag names are kebab-case while OAS parameter names are snake_case, so the generated flag definition keeps both names because reversing the transformation at runtime would be lossy (`AGENTS.md`).

## Alternatives considered

`AGENTS.md` records the rejected alternative for spec defects — editing the vendored specs locally —
and rules it out in favour of fixing the upstream document and re-importing. No other alternative is
documented in the repository.
