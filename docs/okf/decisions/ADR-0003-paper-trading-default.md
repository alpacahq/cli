---
type: Architecture Decision
title: ADR-0003 Paper trading is the default; live trading requires an explicit opt-in
description: Credentials resolve as an atomic bundle and routing defaults to the paper trading endpoint unless something explicitly opts into live.
status: accepted
date: unknown
deciders: []
supersedes: []
affects: [cli, cli.alpaca, cli.alpaca.auth, cli.profile_store]
allium: []
evidence: [internal/config/config.go, internal/config/config_test.go, internal/config/profile_test.go, README.md, AGENTS.md]
tags: [alpacahq, adr, generated]
timestamp: 2026-09-12T00:00:00+00:00
generated_by: claude-opus-5 / layered-docs 2026-09
source_commit: 53606273aa230a40c64b783425dcb3f4423ede30
source_branch: docs/layered-2026-09
generated_at: 2026-09-12T00:00:00+00:00
confidence: medium
review_status: draft-needs-review
---

# ADR-0003 Paper trading is the default; live trading requires an explicit opt-in

## Context

The CLI has no confirmation prompts and executes destructive commands immediately; the README calls
out that closing all positions, cancelling all orders and requesting a short-sale locate are truly
destructive (`README.md`). Because agents and scripts drive it, a misconfiguration must not silently
place real orders. The code comments state the intent: "The paper default is deliberate: scripts and
agents that forget to opt into live should hit paper, not live" (`internal/config/config.go`).

## Decision

Two independent resolutions are performed at startup (`internal/config/config.go`):

- Credentials resolve as an atomic bundle, never mixed across sources, in a fixed order: environment key and secret together, then a profile access token, then profile key and secret together. The winning source is recorded so that it fully determines which auth headers are sent.
- Paper versus live resolves separately: the live-trade environment variable if set, then the profile's live-trade field when credentials came from the profile, then the paper default. Both switches use "live" polarity so the unsafe path always requires an explicit opt-in.

The environment variable check is deliberately strict: only a case-insensitive "true" selects live,
so that typos fall back to paper rather than routing to the live API
(`internal/config/config.go`).

## Consequences

- Paper and live trading base URLs, and the market data URL, are compiled-in constants selected at runtime (`internal/config/config.go`).
- The profile's live-trade field is a pointer so that "not specified" is distinguishable from "explicitly paper"; paper profiles omit the field and live profiles set it to true (`internal/config/config.go`).
- A partial environment bundle does not half-apply: it falls through to the active profile (`README.md`, `internal/config/config.go`).
- Environment-supplied API keys default to paper, which is what makes the documented agent pattern of exporting key and secret safe by default (`README.md`).
- The resolution rules are covered by unit tests over configuration and profile handling (`internal/config/config_test.go`, `internal/config/profile_test.go`).
- The default profile name is itself the paper environment name, so a first run with no configuration resolves to a paper profile (`internal/config/config.go`).

## Alternatives considered

The code comment rules out a permissive parse of the live-trade flag (accepting values such as
"yes" or "1") on the grounds that typos must fail safe. No broader alternative, such as an
interactive confirmation for live trading, is documented in the repository; the stated agent-first
design excludes interactive prompts (`AGENTS.md`, `README.md`).
