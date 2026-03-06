---
name: alpaca-cli
description: >
  Install, configure, and use the Alpaca CLI — a command-line tool for the
  Alpaca Trading API. Covers installation (Go, Homebrew, binary), API key
  authentication, profile management, and agent/automation integration.
  Use when the user asks to install the Alpaca CLI, set up Alpaca API
  credentials, trade stocks or crypto from the command line, get market data
  via CLI, or integrate the Alpaca CLI into scripts, CI pipelines, or AI
  agent workflows. Keywords: alpaca, trading, stocks, crypto, market data,
  brokerage, command line, CLI tool, API key setup.
compatibility: Requires Go (go install) or Homebrew (brew install) for installation. macOS, Linux, and Windows supported.
---

# Alpaca CLI

## Install

Check if already installed:

```bash
alpaca version
```

If not installed, choose one method:

**Go** (recommended — works on macOS, Linux, Windows):

```bash
go install github.com/alpacahq/cli/cmd/alpaca@latest
```

**Homebrew** (macOS / Linux):

```bash
brew install alpacahq/tap/alpaca
```

**Binary download**: Get a prebuilt binary from [GitHub Releases](https://github.com/alpacahq/cli/releases).

## Post-install setup

Install shell completions and man pages:

```bash
alpaca setup
```

This auto-detects the shell and installs to user-level directories (no `sudo`). Override with `--shell fish` or `--shell zsh` if needed.

## Authentication

**Paper trading is the default.** Credentials stored via `alpaca profile login` point to `paper-api.alpaca.markets` unless `--live` is explicitly passed.

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
alpaca profile login --name paper              # default paper
alpaca profile login --name live --live         # live trading
alpaca profile switch live                      # switch active profile
alpaca profile status                           # show active profile
```

## Verify installation

```bash
alpaca version
alpaca clock --json --quiet
```

If authenticated, also verify credentials work:

```bash
alpaca account get --json --quiet
```

Exit code `0` = success, `2` = auth error.

## Agent and automation usage

The CLI is fully non-interactive — no TTY detection, no prompts.

### Always use `--json --quiet`

`--json` gives structured output. `--quiet` suppresses all decorative output (warnings, hints, color). Combine both for machine-readable results:

```bash
alpaca position list --json --quiet
alpaca data latest trade AAPL --json --quiet
```

### Structured errors on stderr

When `--json` or `--quiet` is set, errors are JSON on stderr:

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
alpaca order submit AAPL --side buy --qty 10 --type limit --limit-price 185.00 --dry-run
```

### Pipe JSON payloads

```bash
echo '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}' \
  | alpaca api post /v2/orders --json
```

### Resilience

The CLI retries on 429 and 5xx with exponential backoff (max 3 attempts). `Retry-After` headers are respected.

### Debug API calls

```bash
alpaca account get --verbose   # one-line request/response summary on stderr
alpaca account get --debug     # full headers and bodies on stderr
```

Credentials are always scrubbed from debug output.

## Environment variables

| Variable | Description |
|----------|-------------|
| `ALPACA_API_KEY` | API key (overrides profile) |
| `ALPACA_SECRET_KEY` | Secret key (overrides profile) |
| `ALPACA_BASE_URL` | Trading API base URL |
| `ALPACA_DATA_URL` | Market data API base URL |
| `ALPACA_PROFILE` | Profile name to use |
| `ALPACA_OUTPUT` | Default output format (`table`, `json`, `csv`) |
| `ALPACA_CONFIG_DIR` | Config directory (default: `~/.config/alpaca`) |
| `ALPACA_VERBOSE` | Enable verbose HTTP tracing |
| `ALPACA_DEBUG` | Full HTTP request/response bodies on stderr |

Precedence: flags > env vars > profile config > defaults.

## Self-update

```bash
alpaca update              # download and install latest version
alpaca update --check      # check without installing
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

Data commands support auto-pagination:

```bash
alpaca data bars AAPL --start 2025-01-01 --all
alpaca data trades AAPL --start 2025-01-01 --all --max 5000
alpaca news --symbols AAPL --all --max 100
```

`--all` fetches all available pages. `--max` caps the number of items (default: 10,000).

## Troubleshooting

**`command not found: alpaca`** — Ensure `$GOPATH/bin` (usually `~/go/bin`) is in your `PATH`. For Homebrew, run `brew link alpacahq/tap/alpaca`.

**Exit code 2 on every command** — Credentials are missing or invalid. Re-run `alpaca profile login` or verify `ALPACA_API_KEY` / `ALPACA_SECRET_KEY` env vars.

**Rate limited (429)** — The CLI retries automatically, but if it persists, add delays between calls. Check `Retry-After` in `--verbose` output.

**Completions not working** — Run `alpaca setup` again, then open a new shell. For zsh, ensure `compinit` is loaded in `.zshrc`.

## Anti-patterns

- **NEVER** switch to live trading without explicit user intent. `alpaca profile login` defaults to paper trading — do not pass `--live` unless the user specifically asks for it.
- **NEVER** pass `--secret` as a CLI flag — it leaks into shell history. Use `alpaca profile login` interactively or set `ALPACA_SECRET_KEY` as an env var.
- **NEVER** omit `--json --quiet` in automation or agent workflows — without them, output includes decorative tables, color codes, and hints that break parsing.
- **NEVER** ignore exit code `2` — it means authentication failed. Do not retry; fix credentials first.
- **NEVER** hardcode API keys in scripts or committed files — use environment variables or profile-based auth.
- **NEVER** submit live orders without confirming the user's intent — use `--dry-run` to preview first when there is any ambiguity.

## Further reading

- Full documentation: [README](https://github.com/alpacahq/cli/blob/main/README.md)
- Every command supports `--help` for flag details
- Alpaca API docs: [docs.alpaca.markets](https://docs.alpaca.markets)
