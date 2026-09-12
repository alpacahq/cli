---
type: System
title: Alpaca CLI
description: Single-binary command-line client for the Alpaca Trading API and Market Data API, generated from OpenAPI specifications and aimed at AI agents and automation.
resource: likec4://alpacahq/cli
tags: [alpacahq, cli, trading, market-data, generated]
status: active
owners: []
timestamp: 2026-09-12T00:00:00+00:00
generated_by: claude-opus-5 / layered-docs 2026-09
source_commit: 53606273aa230a40c64b783425dcb3f4423ede30
source_branch: docs/layered-2026-09
generated_at: 2026-09-12T00:00:00+00:00
confidence: medium
review_status: draft-needs-review
links:
  repository: https://github.com/alpacahq/cli
  architecture: ../architecture/cli.c4
---

# Alpaca CLI

## Purpose

`alpaca` is a single Go binary that exposes the Alpaca Trading API and Market Data API on the
command line, so that orders, positions, options, account data, watchlists and market data can be
driven from a shell (`README.md`). The repository states that the primary consumer is an AI agent,
script or automation pipeline rather than a human at an interactive terminal: there are no
confirmation prompts or interactive guardrails, every parameter is an explicit flag, and destructive
commands execute immediately (`README.md`, `AGENTS.md`).

The command surface is not hand-maintained. It is generated from the Alpaca OpenAPI documents
vendored in `api/specs/`, so flags, enum completions, validation and response schemas follow the
published API (`AGENTS.md`, `cmd/generate/`, `Makefile`). The README labels the project an alpha
preview whose commands, flags and output formats may change between releases (`README.md`).

## Capabilities

- Trading commands: orders, positions, options, locate, clock, calendar — evidence: `internal/cmd/root.go`, `internal/cmd/order.go`, `internal/cmd/commands.gen.go`
- Account and asset commands: account, asset, watchlist, wallet, corporate actions — evidence: `internal/cmd/root.go`, `internal/cmd/watchlist.go`
- Market data commands across stocks, crypto, options, forex, fixed income, screener, news and metadata — evidence: `api/specs/market-data-api.json`, `internal/api/marketdata_client.gen.go`
- Raw request escape hatch `alpaca api METHOD path`, including piped JSON bodies — evidence: `internal/cmd/api.go`, `README.md`
- Output control: JSON by default, CSV, inline jq filtering, quiet mode, response-schema printing without an API call, per-request timeout — evidence: `internal/output/output.go`, `internal/output/jq.go`, `internal/cmd/root.go`
- Profile and credential management, including a browser OAuth login and named profiles — evidence: `internal/cmd/auth.go`, `internal/config/config.go`, `internal/oauth/oauth.go`
- Diagnostics: `alpaca doctor`, plus verbose, debug and trace modes with credentials scrubbed from output — evidence: `internal/cmd/doctor.go`, `internal/client/client.go`
- Shell completion scripts and self-update against the latest published release — evidence: `internal/cmd/update.go`, `.goreleaser.yaml`

## Interfaces

**Inbound:** a command-line interface only. The cobra root command `alpaca` with trading, account
and utility command groups is the single entrypoint (`cmd/alpaca/main.go`, `internal/cmd/root.go`).
Behaviour is also driven by environment variables and by a local profile file
(`internal/config/config.go`, `README.md`). During OAuth login the CLI briefly listens on a
localhost callback port from a fixed list (`internal/oauth/config.go`). No network service is
exposed.

**Outbound:** HTTPS calls to the Alpaca trading endpoints (paper or live, selected at runtime) and
to the market data endpoint (`internal/config/config.go`, `internal/client/client.go`); the Alpaca
OAuth authorize and token endpoints during login (`internal/oauth/config.go`); and the GitHub
releases API to determine whether a newer version exists (`internal/cmd/update.go`). At development
time `make spec-update` and `make spec-check` fetch the published OpenAPI documents (`Makefile`).

## Dependencies

- Alpaca Trading API — the API the CLI exists to wrap; paper and live base URLs are compiled in and selected at runtime — evidence `internal/config/config.go`
- Alpaca Market Data API — market data commands — evidence `internal/config/config.go`, `api/specs/market-data-api.json`
- Alpaca OAuth authorization service — `alpaca profile login` authorization-code flow — evidence `internal/oauth/config.go`, `internal/oauth/oauth.go`
- GitHub releases API — update checks and release distribution — evidence `internal/cmd/update.go`, `.github/workflows/release.yml`
- `spf13/cobra` and `spf13/pflag` — command and flag framework — evidence `go.mod`
- `itchyny/gojq` — inline jq filtering of responses — evidence `go.mod`, `internal/output/jq.go`
- `fatih/color`, `yarlson/tap`, `golang.org/x/term` — terminal output and prompts — evidence `go.mod`
- `gopkg.in/yaml.v3` — config and profile serialisation — evidence `go.mod`, `internal/config/config.go`
- `oasdiff` (`github.com/oasdiff/oasdiff`) — development-time OpenAPI changelog tool; `make spec-check` aborts with an install instruction when it is absent — evidence `Makefile`
- `curl`, `python3` — used by `make spec-update` to fetch and reformat the published OpenAPI documents — evidence `Makefile`

## Data & storage

No server-side datastore. Two local stores are evidenced:

- Vendored OpenAPI documents `api/specs/trading-api.json` and `api/specs/market-data-api.json`, treated as read-only inputs to the generator and not shipped with or read by the binary — there is no `go:embed` in the module and `api/specs` is opened only by `cmd/generate` and a test helper (`AGENTS.md`, `Makefile`, `cmd/generate/main.go`, `internal/cmd/spec_test.go`).
- A per-user configuration directory containing `config.yaml` and one YAML file per named profile, created with restricted directory and file permissions and holding API keys or an OAuth access token. The location defaults to `~/.config/alpaca` on all platforms — `config.Dir()` joins the user home directory with `.config/alpaca` rather than using the OS user-config directory — and is overridden by `ALPACA_CONFIG_DIR` (`internal/config/config.go`, `README.md`, `.agents/skills/alpaca-cli/SKILL.md`).

Credentials resolve as an atomic bundle, never mixed across sources: environment key and secret
first, then a profile access token, then profile key and secret (`internal/config/config.go`,
`README.md`).

## Operations

Build, test and release are the only runtime concerns; nothing in the repository deploys a service.

- `make build`, `make test`, `make lint`, `make check`, `make generate`, `make test-integration`, `make spec-update`, `make spec-check` — evidence `Makefile`
- CI on pushes and pull requests to the default branch: build, vet, a generated-code freshness check that fails when `make generate` produces a diff, race-enabled tests with coverage, a binary build and a smoke run of `alpaca version`; a separate job runs the integration suite against the paper API using repository secrets; a third job runs the linter — evidence `.github/workflows/ci.yml`
- Release on tags matching a version pattern: GoReleaser cross-compiles for Linux, macOS and Windows on amd64 and arm64, publishes archives and checksums, and pushes a formula to the organisation Homebrew tap when a tap token is available — evidence `.github/workflows/release.yml`, `.goreleaser.yaml`
- Distribution and upgrade paths are `go install` and Homebrew; `alpaca update` detects which was used and offers the matching upgrade command — evidence `README.md`, `internal/cmd/update_check.go`
- No Dockerfile, compose file, Kubernetes manifest, Terraform or serverless configuration is present in the repository, so no deployment model is recorded — evidence: repository file listing at the recorded commit

## Behaviour (Allium)

none — no Allium specification files exist in the repository at the recorded commit, and no accepted
specification was available to this pass.

## Decisions

- [ADR-0001 Generate the CLI from the Alpaca OpenAPI specifications](decisions/ADR-0001-generate-cli-from-openapi-specs.md)
- [ADR-0002 Agent-first output contract: JSON for API commands, text for operational commands](decisions/ADR-0002-agent-first-output-contract.md)
- [ADR-0003 Paper trading is the default; live trading requires an explicit opt-in](decisions/ADR-0003-paper-trading-default.md)
- [ADR-0004 OAuth login uses an embedded public client and is restricted to paper trading](decisions/ADR-0004-oauth-public-client-paper-only.md)
- [ADR-0005 Distribute as prebuilt binaries via GoReleaser, a Homebrew tap and go install](decisions/ADR-0005-distribution-goreleaser-homebrew.md)

## Evidence

- `README.md` — purpose, agent-first framing, authentication, command areas, output rules, configuration, development and support
- `AGENTS.md` — generate-everything principle, design philosophy, output contract, integration-test rules
- `go.mod` — Go version and direct dependencies
- `Makefile` — build, test, lint, generate, spec-update and spec-check targets and the spec source URL
- `.github/workflows/ci.yml` — build, generated-code freshness, tests, integration job, lint
- `.github/workflows/release.yml`, `.goreleaser.yaml` — tag-triggered cross-platform release and Homebrew tap publication
- `cmd/alpaca/main.go`, `internal/cmd/root.go` — CLI entrypoint, command groups, global flags, error and schema printing
- `cmd/generate/main.go`, `cmd/generate/commands.go` — the generator
- `internal/api/trading_client.gen.go`, `internal/api/marketdata_client.gen.go`, `internal/api/descriptions.gen.go` — generated API layer
- `internal/client/client.go` — HTTP transport, auth headers, retry and backoff, credential scrubbing, tracing
- `internal/config/config.go` — credential and base-URL resolution, profile storage
- `internal/oauth/config.go`, `internal/oauth/oauth.go` — OAuth endpoints, callback ports, login flow
- `internal/output/output.go`, `internal/output/jq.go` — JSON and CSV rendering, jq filtering
- `internal/cmd/update.go`, `internal/cmd/update_check.go` — release lookup and install-method detection
- `api/specs/trading-api.json`, `api/specs/market-data-api.json` — vendored OpenAPI inputs
- `.agents/skills/alpaca-cli/SKILL.md` — agent-facing usage guidance: `--quiet` in automation, exit codes 0/1/2, JSON errors on stderr, credential precedence, environment variables, anti-patterns
- `.agents/skills/alpaca-cli-regenerate/SKILL.md` — maintainer regeneration pipeline: `make spec-update`, `make generate`, generator registry rules, golden-file updates, `make check`

## Open questions

- Ownership is not recorded. There is no CODEOWNERS file and neither the README nor AGENTS.md names an owning team, so `owners` is left empty. Confidence: low on ownership.
- The repository history available at the recorded commit is a single commit, so decision rationale had to be mined from the README, AGENTS.md and in-code comments rather than from commit or pull-request discussion. Confidence: low on decision dates and deciders.
- `status: active` is inferred from the alpha-preview notice, the tag-triggered release workflow and the self-update command; no lifecycle statement exists in the repository. Confidence: medium.
- The four outbound Alpaca dependencies evidenced here — the trading endpoints, the market data endpoint, the OAuth authorize/token endpoints, and the published OpenAPI documents fetched at development time by `make spec-update` / `make spec-check` — are not represented in the LikeC4 model. The org external dictionary is canonical and currently has no matching id, and a repository pass must not mint one, so four ids remain to be reserved by the org pass (the published-document one being development-time only). Until then the LikeC4 context view shows only the release/update dependency. Confidence: low on the eventual ids and on how these endpoints are classified at the org merge.
- `oasdiff` is invoked only by `make spec-check`; whether it is installed in CI was not verified, as no workflow step references it. Confidence: low.

> Unverified: the set of Alpaca API endpoints actually reachable is determined by the vendored
> OpenAPI documents, which were not read in full. Only the endpoint hosts and the command areas
> named in `README.md` and `internal/config/config.go` were verified.
