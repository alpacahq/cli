# Alpaca CLI

[![Forum](https://img.shields.io/badge/Forum-Alpaca%20Community-00C805?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAACXBIWXMAAAsTAAALEwEAmpwYAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAL0SURBVHgBtVY/T1NRFP+d+1pojAoxRiMuj0WBDj4WEyfALwAmrsbdRVx0cBASTdyERVdI/AD1AxhkMtHBOkDRpU00khgVFCW09d3jube0KbR979Y+fsnt+9Pz//7ueYfgCC4Gg+FeZYZAEyAErMkH8WDtT9omxSUw8gxe9TJ9ORrOb7vYpVjHhcAPuTILqJsNh26GlxTS8zSaL8XIdXAsGety9QEzZtEDiLCg+iWQDhVpG4DJWqO6woCPBCBOSlKNqXbVUIdfVNaCIEnnBsaWsVlZGwvQGlyTYMKZt3HWUokDAYSFbLEb5+rMPXD5A+CdAKkT0N+exurYIDLp8TonUg3nG9knQjgf3SAzAnXqBrj6BZQeAkkwvPMyUsVuh5Bbbu/YJOxLKX03bLfOBmZs5vVne+2/6KRvfBmf5t5WIOTqXHxHgC0zDT0CHb+KXlHrLZglficdLlPdclFS58T5wHSkjP58G/w7ehssCNtef3pYhZnKDBxQK/t0rJwaeggnMGxrV9LHJ50Ujl12EoPZJiGnC0hhQsnPpVjB07egzt6tPegdhB+vQH99jKggnMAUpFjL0YsgYMu+i3HvwmskAXMklTiP/MLx7lscGYQHKlbmZ07WCxwVhISIHRz05n2w2fPyhjz8QmKQo2hIWHKR1T+eIyxel1a7gqRgvwvM+n1XWvttN6EI8sIBetWVjud4xBzAmleVt9eXc+FBAwlWwAyvisbNd5mXnbXUSSQB4d6SmQnsMfRILzhpJZi90uG8vdpoRjdkpqfFOCVKn0ciYF60PtE0lHrl1ByBSzhiGB+psfXG8NMIwHBBQU9FBtFjBYxt46P53YFWbMoSalzrFAT1wAFj09reL33bAAz6suv5jpVQ/9cD6pkb2y0m2ypIlN7o+nALMV23wGs6qkI4ldkdP5y5M7gw4v8tZJdksf7zhp1Q+bQVfl9e4M1nfpx9h1l4P5CiPwi/aObHSVlmivJl1WcJ00lLskyJV2XliMipu/4DGLZL7Y7ZLckAAAAASUVORK5CYII=)](https://forum.alpaca.markets/)
[![Slack](https://img.shields.io/badge/Slack-Alpaca%20Community-4A154B?logo=slack&logoColor=white)](https://alpaca.markets/slack)

CLI for the [Alpaca](https://alpaca.markets) Trading API. Trade stocks and crypto, access market data, and manage your account from the command line.

> [!WARNING]
> **Alpha Preview** - This CLI is under active development. Commands, flags, and output formats may change or be removed without notice between releases. Do not depend on current behavior in production workflows.

## Built For Agents

Alpaca CLI is designed for AI agents, scripts, and automation pipelines. It is not an interactive trading terminal: there are no confirmation prompts, "are you sure?" dialogs, or interactive guardrails. Every command executes immediately.

Destructive commands are truly destructive:

- `alpaca position close-all` liquidates your entire portfolio.
- `alpaca order cancel-all` cancels every open order without listing them first.
- `alpaca locate create` requests shares for a short sale and may incur locate fees.

Paper trading is the default. Live trading requires an explicit opt-in.

## Install

```bash
# Go
go install github.com/alpacahq/cli/cmd/alpaca@latest

# Homebrew
brew install alpacahq/tap/cli
```

## Quick Start

```bash
# Authenticate with paper trading OAuth
alpaca profile login

# Check your account
alpaca account get

# Submit an order
alpaca order submit --symbol AAPL --side buy --qty 10 --type market

# List positions and open orders
alpaca position list
alpaca order list --status open

# Get market data
alpaca data bars --symbol AAPL --start 2025-01-01 --timeframe 1Day

# Check if the market is open
alpaca clock
```

## Authentication

`alpaca profile login` opens a browser for OAuth authorization and stores a paper trading profile in `~/.config/alpaca/profiles/` with restricted file permissions.

```bash
alpaca profile login                              # OAuth, paper trading
alpaca profile login --api-key                    # API key and secret, paper trading
alpaca profile login --api-key --live             # API key and secret, live trading
alpaca profile login --api-key --name prod --live # Named live profile
alpaca profile switch prod                        # Switch active profile
```

OAuth is currently paper-only. Live trading requires API keys.

For scripts, CI, and agents, prefer environment variables so secrets do not touch disk:

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
alpaca account get --quiet
```

Environment API keys default to paper trading. To use live trading:

```bash
export ALPACA_LIVE_TRADE=true
```

Credential lookup uses the first complete credential bundle:

1. `ALPACA_API_KEY` and `ALPACA_SECRET_KEY`
2. Profile `access_token`
3. Profile `api_key` and `secret_key`

A partial environment bundle falls through to the active profile. OAuth tokens cannot be supplied through environment variables.

## Commands

The CLI is generated from Alpaca OpenAPI specs, so the installed binary is the source of truth for commands, flags, enum completions, validation, and response schemas.

```bash
alpaca --help-all              # Full command reference
alpaca order submit --help     # Flags for one command
alpaca order submit --schema   # Response fields without an API call
```

Top-level command areas include:

- Trading: `order`, `position`, `option`, `locate`, `clock`, `calendar`
- Account and assets: `account`, `asset`, `watchlist`, `wallet`, `corporate-action`
- Market data: `data`, `data crypto`, `data option`, `data forex`, `data fixed-income`, `data meta`, `data screener`, `data news`
- Utilities: `profile`, `api`, `doctor`, `update`, `version`, `completion`

Use `alpaca api [METHOD] <path>` as a raw escape hatch for direct API requests.

## Output

API commands return JSON on stdout by default. They also support CSV output, inline jq filtering, quiet mode, response schemas, and request timeouts.

```bash
alpaca position list
alpaca position list --csv
alpaca position list --jq '.[0].symbol'
alpaca order list --quiet
alpaca data bars --symbol AAPL --start 2020-01-01 --timeout 120
```

Operational commands such as `version`, `doctor`, `profile`, `update`, `completion`, and help emit human-readable text. The exception is `alpaca update --check`, which emits JSON for automation.

Errors are JSON on stderr:

```json
{"error":"rate limited","code":0,"status":429,"hint":"Rate limited. Reduce request frequency or add delays between calls."}
```

Exit codes:

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | API or general error |
| `2` | Authentication error |

## Configuration

| Variable | Description |
|----------|-------------|
| `ALPACA_API_KEY` | API key. Must be set with `ALPACA_SECRET_KEY`. |
| `ALPACA_SECRET_KEY` | Secret key. Must be set with `ALPACA_API_KEY`. |
| `ALPACA_LIVE_TRADE` | `true` routes to live trading. Anything else routes to paper trading. |
| `ALPACA_PROFILE` | Profile name to use. |
| `ALPACA_OUTPUT` | Default output format: `json` or `csv`. |
| `ALPACA_CONFIG_DIR` | Config directory. Defaults to `~/.config/alpaca`. |
| `ALPACA_QUIET` | Suppress non-data output such as warnings, hints, and color. |
| `ALPACA_VERBOSE` | Show HTTP request summaries on stderr. |
| `ALPACA_DEBUG` | Show HTTP request and response headers and bodies on stderr. |
| `ALPACA_TRACE` | Show HTTP timing breakdown on stderr. |

Global flags: `--csv`, `--jq`, `--profile`, `--verbose`, `--debug`, `--trace`, `--quiet`, `--schema`, `--timeout`.

## Automation Notes

For AI agents, see the [`alpaca-cli` Agent Skill](.agents/skills/alpaca-cli/SKILL.md) for structured installation, authentication, and usage guidance.

For unattended order submission, pass `--client-order-id` so retries after ambiguous failures do not create duplicate orders:

```bash
client_order_id="$(uuidgen)"
alpaca order submit \
  --symbol AAPL \
  --side buy \
  --qty 10 \
  --type market \
  --client-order-id "$client_order_id" \
  --quiet
```

Use `--dry-run` to preview an order without submitting:

```bash
alpaca order submit --symbol AAPL --side buy --qty 10 --type limit --limit-price 185.00 --dry-run
```

Pipe JSON payloads into raw `POST` and `PATCH` API requests:

```bash
echo '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}' \
  | alpaca api POST /v2/orders
```

The CLI retries 429 and 5xx responses with exponential backoff up to 3 attempts. `Retry-After` headers are respected.

## Diagnostics

```bash
alpaca doctor                  # Check config and API connectivity
alpaca account get --verbose   # Request summary on stderr
alpaca account get --trace     # DNS, TLS, TTFB, and total timing on stderr
alpaca account get --debug     # Headers and bodies on stderr
```

Credentials are always scrubbed from diagnostic output.

## Shell Completions

```bash
alpaca completion bash
alpaca completion zsh
alpaca completion fish
alpaca completion powershell
```

Save the generated completion script where your shell expects it, then open a new shell. Enum-valued flags complete with valid values from the OpenAPI specs.

## Self-Update

```bash
alpaca update          # Check for updates and prompt to install
alpaca update --yes    # Check and install without prompting
alpaca update --check  # Machine-readable update check
```

Upgrade commands are tailored to the detected install method:

| Install method | Upgrade command |
|---|---|
| Go | `go install github.com/alpacahq/cli/cmd/alpaca@latest` |
| Homebrew | `brew upgrade alpacahq/tap/cli` |

## Development

The CLI is driven by OpenAPI specs in `api/specs/*.json`. Do not edit generated files directly. Change the specs or generator, then run `make generate`.

```bash
make build            # Build binary to bin/alpaca
make test             # Run unit tests
make lint             # Run linter
make check            # Lint, test, and build
make test-integration # Run live paper API integration tests
make generate         # Regenerate API clients and command metadata
make spec-update      # Fetch latest OpenAPI specs
```

Integration tests require paper trading credentials:

```bash
export ALPACA_TEST_API_KEY=PK...
export ALPACA_TEST_SECRET_KEY=...
export ALPACA_TEST_BASE_URL=https://paper-api.alpaca.markets # optional
make test-integration
```

## Support

- **CLI issues:** Bugs, feature requests, or questions specific to this CLI → [GitHub Issues](https://github.com/alpacahq/cli/issues).
- **General Alpaca support & API discussion:** Account questions, platform issues, or broader API topics → [Alpaca Community Forum](https://forum.alpaca.markets/).
- **Slack community:** Chat with other developers and the Alpaca community on Slack ([join here](https://alpaca.markets/slack)).

## License

Apache 2.0
