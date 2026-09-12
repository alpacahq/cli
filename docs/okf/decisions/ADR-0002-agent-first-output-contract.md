---
type: Architecture Decision
title: ADR-0002 Agent-first output contract — JSON for API commands, text for operational commands
description: API commands emit structured JSON on stdout and JSON errors on stderr; commands that manage the CLI itself emit human-readable text and signal via exit code.
status: accepted
date: unknown
deciders: []
supersedes: []
affects: [cli, cli.alpaca, cli.alpaca.renderer, cli.alpaca.commands]
allium: []
evidence: [AGENTS.md, README.md, internal/output/output.go, internal/output/jq.go, internal/cmd/root.go, internal/cmd/update.go, internal/client/client.go]
tags: [alpacahq, adr, generated]
timestamp: 2026-09-12T00:00:00+00:00
generated_by: claude-opus-5 / layered-docs 2026-09
source_commit: 53606273aa230a40c64b783425dcb3f4423ede30
source_branch: docs/layered-2026-09
generated_at: 2026-09-12T00:00:00+00:00
confidence: medium
review_status: draft-needs-review
---

# ADR-0002 Agent-first output contract — JSON for API commands, text for operational commands

## Context

`AGENTS.md` names the primary consumer: "the primary consumer is an AI agent. All parameters are
explicit flag value pairs — no positional arguments", with the raw `alpaca api` escape hatch as the
single stated exception. The README makes the same point for users and adds that there are no
confirmation prompts or interactive guardrails (`README.md`). An agent-driven CLI therefore needs a
stable, machine-parseable contract, but the commands that manage the CLI itself are read by humans
troubleshooting their setup.

## Decision

Commands are split into two output categories (`AGENTS.md`, section "Output contract"):

- API commands — trading, data, account, watchlist and similar — return structured JSON on stdout and support `--csv`, `--jq`, `--quiet` and `--schema`. Errors are JSON on stderr.
- Operational commands — `version`, `doctor`, `profile`, `update`, `completion`, help and `--schema` — emit human-readable text; the machine-readable signal is the exit code, not the output format.

One exception is recorded: `alpaca update --check` emits JSON "because agents need to
programmatically decide whether to upgrade" (`AGENTS.md`, `README.md`).

## Consequences

- A single render pipeline applies the jq filter and then the chosen format for API commands, and CSV headers are derived from the generated response schema (`internal/cmd/root.go`, `internal/output/output.go`).
- Errors are emitted as a JSON object on stderr carrying message, code, HTTP status and a hint, plus method, path and request id when known (`internal/cmd/root.go`, `internal/client/client.go`).
- Exit codes are fixed and narrow: success, API or general error, and a distinct authentication error returned for HTTP 401 (`internal/client/client.go`, `README.md`).
- Hints are curated per status code so that an agent reading stderr gets an actionable next step for validation errors, rate limiting, invalid credentials and forbidden responses (`internal/client/client.go`).
- The guidance explicitly rejects converting operational commands to JSON, and records the intended agent pattern for connectivity checks: run an API command in quiet mode rather than parsing `doctor` output (`AGENTS.md`).
- `update --check` implements the exception by encoding current version, latest version, availability, install method and upgrade command as JSON (`internal/cmd/update.go`).

## Alternatives considered

`AGENTS.md` records and rejects the alternative of making every command JSON, on the grounds that
operational commands help a human troubleshoot and their machine signal is the exit code. No other
alternative is documented in the repository.
