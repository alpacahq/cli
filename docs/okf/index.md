# Alpaca CLI — knowledge bundle (OKF)

Generated draft (layered-docs 2026-09) — review required. Concept id = path without `.md`.

- [system](system.md) — System: single-binary CLI for the Alpaca Trading API and Market Data API, generated from OpenAPI specifications and aimed at agents and automation.
- [decisions/ADR-0001-generate-cli-from-openapi-specs](decisions/ADR-0001-generate-cli-from-openapi-specs.md) — Architecture Decision: the command surface, typed clients and response schemas are generated from vendored OpenAPI documents.
- [decisions/ADR-0002-agent-first-output-contract](decisions/ADR-0002-agent-first-output-contract.md) — Architecture Decision: JSON on stdout for API commands, human-readable text plus exit codes for operational commands.
- [decisions/ADR-0003-paper-trading-default](decisions/ADR-0003-paper-trading-default.md) — Architecture Decision: credentials resolve as an atomic bundle and routing defaults to paper trading unless live is explicitly selected.
- [decisions/ADR-0004-oauth-public-client-paper-only](decisions/ADR-0004-oauth-public-client-paper-only.md) — Architecture Decision: the CLI is a public native OAuth client with a localhost callback, and OAuth login is paper-only.
- [decisions/ADR-0005-distribution-goreleaser-homebrew](decisions/ADR-0005-distribution-goreleaser-homebrew.md) — Architecture Decision: tagged releases are cross-compiled by GoReleaser and published with a Homebrew tap formula, with install-method-aware self-update.

Architecture (LikeC4): [../architecture/cli.c4](../architecture/cli.c4), views `cli_context` and `cli_containers`.
