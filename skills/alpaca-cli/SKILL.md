---
name: alpaca-cli
description: >
  Install, configure, and use the Alpaca CLI — a command-line tool for the
  Alpaca Trading API. Covers installation (Go), API key
  authentication, profile management, and agent/automation integration.
  Use when the user asks to install the Alpaca CLI, set up Alpaca API
  credentials, trade stocks or crypto from the command line, get market data
  via CLI, or integrate the Alpaca CLI into scripts, CI pipelines, or AI
  agent workflows. Keywords: alpaca, trading, stocks, crypto, market data,
  brokerage, command line, CLI tool, API key setup.
compatibility: Requires Go (go install) for installation. macOS, Linux, and Windows supported.
---

# Alpaca CLI

## Install

Check if already installed:

```bash
alpaca version
```

If not installed:

```bash
go install github.com/alpacahq/cli/cmd/alpaca@latest
```

## Authentication

**Paper trading is the default.** `alpaca profile login` uses OAuth and is paper-only. For live trading, use API keys: `alpaca profile login --api-key --live`.

### Interactive login (stores credentials on disk)

```bash
alpaca profile login
```

Credentials are stored in `~/.config/alpaca/profiles/` with 0600 permissions.

### Environment variables (for scripts, CI, agents)

No secrets touch disk. Preferred for automation:

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
```

Env vars override profile credentials.

### Multiple profiles

```bash
alpaca profile login --name paper              # OAuth, paper (default)
alpaca profile login --api-key --name live --live  # API keys, live trading
alpaca profile switch live                      # switch active profile
alpaca doctor                                   # show active profile + connectivity
```

## Verify installation

```bash
alpaca version
alpaca clock --quiet
```

If authenticated, also verify credentials work:

```bash
alpaca account get --quiet
```

Exit code `0` = success, `2` = auth error.

## Agent and automation usage

The CLI is fully non-interactive — no TTY detection, no prompts.

### Always use `--quiet`

JSON is the default output format. `--quiet` suppresses all non-data output (warnings, hints). Use it for machine-readable results:

```bash
alpaca position list --quiet
alpaca data latest-trade --symbol AAPL --quiet
```

### Structured errors on stderr

Errors are always JSON on stderr:

```json
{"error":"rate limited","code":0,"status":429,"hint":"Rate limited. Reduce request frequency or add delays between calls."}
```

### Exit codes

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | API or general error |
| `2` | Authentication error (401/403) |

### Dry run

Preview an order without submitting:

```bash
alpaca order submit --symbol AAPL --side buy --qty 10 --type limit --limit-price 185.00 --dry-run
```

### Pipe JSON payloads

```bash
echo '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}' \
  | alpaca api POST /v2/orders
```

### Resilience

The CLI retries on 429 and 5xx with exponential backoff (max 3 attempts). `Retry-After` headers are respected.

### Debug API calls

```bash
alpaca account get --verbose   # one-line request/response summary on stderr
alpaca account get --trace     # timing breakdown (DNS, TLS, TTFB) on stderr
alpaca account get --debug     # full headers and bodies on stderr
```

Credentials are always scrubbed from debug output.

### Filter JSON output

```bash
alpaca position list --jq '.[0].symbol'
alpaca order list --jq '[.[] | {id, symbol, side, qty}]'
```

`--jq` applies a jq expression to the JSON output without requiring an external `jq` install.

## Environment variables

| Variable | Description |
|----------|-------------|
| `ALPACA_API_KEY` | API key (overrides profile) |
| `ALPACA_SECRET_KEY` | Secret key (overrides profile) |
| `ALPACA_BASE_URL` | Trading API base URL |
| `ALPACA_DATA_URL` | Market data API base URL |
| `ALPACA_PROFILE` | Profile name to use |
| `ALPACA_OUTPUT` | Default output format (`json`, `csv`) |
| `ALPACA_CONFIG_DIR` | Config directory (default: `~/.config/alpaca`) |
| `ALPACA_VERBOSE` | Show HTTP request summaries on stderr |
| `ALPACA_DEBUG` | Show HTTP request/response headers and bodies on stderr |
| `ALPACA_TRACE` | Show HTTP timing breakdown on stderr (DNS, TLS, TTFB) |
| `ALPACA_NO_UPDATE_NOTIFY` | Suppress background update notices |

Precedence: flags > env vars > profile config > defaults.

## Self-update

The CLI checks for updates in the background and shows a notice when a newer version is available. Before starting work, check for updates programmatically:

```bash
alpaca update --check --quiet
```

This returns structured output:

```json
{"current":"0.0.1","latest":"0.0.2","update_available":true,"install_method":"goinstall","update_command":"go install github.com/alpacahq/cli/cmd/alpaca@latest"}
```

If `update_available` is `true`, run the `update_command` value to upgrade.

Manual commands:

```bash
alpaca update              # check for updates and show upgrade command
alpaca update --check      # machine-readable update check (JSON)
```

## Discovering commands

```bash
alpaca --help-all                    # dump all commands, subcommands, and flags
alpaca order --help                  # help for a command group
alpaca order submit --help           # help for a specific command
alpaca order list --schema           # show response fields for a command
alpaca doctor                        # check config and API connectivity
```

Use `--help-all` to find the right command. Use `<command> --help` for flag details. Use `<command> --schema` to see API response fields generated from the spec. These are always current — never rely on stale documentation when the CLI is installed.

## Pagination

Data commands return one page by default. Use `--limit` to control page size and `--page-token` to fetch subsequent pages:

```bash
alpaca data bars --symbol AAPL --start 2025-01-01 --limit 500
alpaca data bars --symbol AAPL --start 2025-01-01 --limit 500 --page-token <token>
```

The response includes a `next_page_token` field when more data is available.

## Troubleshooting

**`command not found: alpaca`** — Ensure `$GOPATH/bin` (usually `~/go/bin`) is in your `PATH`.

**Exit code 2 on every command** — Credentials are missing or invalid. Re-run `alpaca profile login` or verify `ALPACA_API_KEY` / `ALPACA_SECRET_KEY` env vars.

**Rate limited (429)** — The CLI retries automatically, but if it persists, add delays between calls. Check `Retry-After` in `--verbose` output.

**Completions not working** — Run `alpaca completion <shell>` to generate completions, then open a new shell. For zsh, ensure `compinit` is loaded in `.zshrc`.

## Anti-patterns

- **NEVER** switch to live trading without explicit user intent. `alpaca profile login` defaults to paper trading — do not pass `--live` unless the user specifically asks for it.
- **NEVER** pass `--secret` as a CLI flag — it leaks into shell history. Use `alpaca profile login` interactively or set `ALPACA_SECRET_KEY` as an env var.
- **NEVER** omit `--quiet` in automation or agent workflows — without it, output may include hints and warnings on stderr that break parsing.
- **NEVER** ignore exit code `2` — it means authentication failed. Do not retry; fix credentials first.
- **NEVER** hardcode API keys in scripts or committed files — use environment variables or profile-based auth.
- **NEVER** submit live orders without confirming the user's intent — use `--dry-run` to preview first when there is any ambiguity.

## Further reading

- Full documentation: [README](https://github.com/alpacahq/cli/blob/main/README.md)
- Every command supports `--help` for flag details
- Alpaca API docs: [docs.alpaca.markets](https://docs.alpaca.markets)
