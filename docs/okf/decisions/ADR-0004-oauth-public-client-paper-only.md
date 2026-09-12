---
type: Architecture Decision
title: ADR-0004 OAuth login uses an embedded public client and is restricted to paper trading
description: The CLI ships as a public native OAuth client with a localhost callback, and OAuth login is limited to paper trading until the flow is hardened.
status: accepted
date: unknown
deciders: []
supersedes: []
affects: [cli, cli.alpaca, cli.alpaca.auth, cli.profile_store]
allium: []
evidence: [internal/oauth/config.go, internal/oauth/oauth.go, internal/cmd/auth.go, internal/cmd/auth_test.go, README.md]
tags: [alpacahq, adr, generated]
timestamp: 2026-09-12T00:00:00+00:00
generated_by: claude-opus-5 / layered-docs 2026-09
source_commit: 53606273aa230a40c64b783425dcb3f4423ede30
source_branch: docs/layered-2026-09
generated_at: 2026-09-12T00:00:00+00:00
confidence: medium
review_status: draft-needs-review
---

# ADR-0004 OAuth login uses an embedded public client and is restricted to paper trading

## Context

`alpaca profile login` opens a browser for authorization and stores a paper trading profile
(`README.md`). A distributed command-line binary cannot keep a client credential confidential, and
the repository documents this explicitly: the CLI's OAuth application identifiers are embedded in
the binary and are to be treated as public, because any user can extract them
(`internal/oauth/config.go`). The same comment cites RFC 8252 on native apps as public clients and
notes that another widely used CLI embeds its credential in the same way.

## Decision

The CLI registers as a first-party native OAuth application whose identifiers are compiled into the
binary and documented as public. Security is not taken from the confidentiality of that credential
but from three server-side and flow-level controls named in the code: user consent in the browser,
server-side redirect URI validation against pre-registered localhost callback ports, and a state
parameter for CSRF protection (`internal/oauth/config.go`).

Because the flow does not yet use PKCE or the device authorization grant, OAuth login is restricted
to paper trading. Live trading requires API keys (`internal/oauth/config.go`, `README.md`).

## Consequences

- The authorize endpoint, token endpoint and default scope set are fixed constants (`internal/oauth/config.go`).
- A short list of candidate localhost callback ports is tried in order; each must be pre-registered on the OAuth application, which bounds where a token can be delivered (`internal/oauth/config.go`).
- The OAuth access token is written into a named profile file in the user configuration directory with restricted permissions, and OAuth tokens cannot be supplied through environment variables (`internal/config/config.go`, `README.md`).
- A profile access token wins over profile key and secret in credential resolution, and the client sends a bearer authorization header in that case instead of the API-key headers (`internal/config/config.go`, `internal/client/client.go`).
- Scripts, CI and agents are steered away from the OAuth path and towards environment API keys so that secrets do not touch disk (`README.md`).
- The login flow and its edge cases are covered by unit tests (`internal/cmd/auth_test.go`).
- A hardening path is recorded in the code: adopt PKCE or the device authorization grant, after which the paper-only restriction could be revisited (`internal/oauth/config.go`).

## Alternatives considered

The code comment considers and rejects treating the embedded credential as a security boundary, and
records PKCE and the device authorization grant as the intended future flows rather than current
ones. API-key authentication is retained in parallel as the only route to live trading
(`internal/oauth/config.go`, `README.md`).
