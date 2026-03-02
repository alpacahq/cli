# alpaca

CLI for [Alpaca](https://alpaca.markets) Trading API. Trade stocks & crypto, access market data, and manage your account from the command line.

## Install

```bash
go install github.com/alpacahq/cli/cmd/alpaca@latest
```

Or download a prebuilt binary from [Releases](https://github.com/alpacahq/cli/releases).

## Quick Start

```bash
# Authenticate
alpaca profile login

# Check your account
alpaca account get

# Buy 10 shares of AAPL
alpaca buy AAPL 10

# Get a quick price check
alpaca price AAPL

# List your positions
alpaca positions

# List open orders
alpaca orders

# Get market data
alpaca data bars --symbol AAPL --start 2025-01-01 --timeframe 1Day

# Check if market is open
alpaca clock
```

## Safety

**Paper trading is the default.** When you run `alpaca profile login`, credentials are stored for paper trading (`paper-api.alpaca.markets`) unless you explicitly pass `--live`.

- Destructive operations (`order cancel-all`, `position close-all`) on live accounts require `--confirm` to proceed.
- `buy` and `sell` commands print a warning when targeting a live account.
- Suppress informational warnings with `suppress_warnings: true` in your profile config. Confirmation prompts for destructive operations cannot be suppressed.

**Credential safety:**

- For interactive use, `alpaca profile login` stores credentials in `~/.config/alpaca/profiles/` with restricted file permissions (0600).
- For CI/automation, use environment variables instead of stored profiles — no secrets touch disk:

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
alpaca positions --json
```

- Passing `--secret` via flags is discouraged (shell history exposure). Use interactive login or env vars.

## Commands

### Trading

| Command | Description |
|---------|-------------|
| `alpaca buy <sym> <qty>` | Buy shares (market order) |
| `alpaca sell <sym> <qty>` | Sell shares (market order) |
| `alpaca order submit` | Submit any order type |
| `alpaca order list` | List orders |
| `alpaca order get <id>` | Get order details |
| `alpaca order cancel <id>` | Cancel an order |
| `alpaca order cancel-all` | Cancel all open orders |
| `alpaca order replace <id>` | Replace an existing order |
| `alpaca position list` | List open positions |
| `alpaca position get <sym>` | Get position for a symbol |
| `alpaca position close <sym>` | Close a position |
| `alpaca position close-all` | Close all positions |
| `alpaca option chain <sym>` | Options chain (trading API) |
| `alpaca option get <id>` | Option contract details |
| `alpaca option exercise <id>` | Exercise an option |

### Market Data

| Command | Description |
|---------|-------------|
| `alpaca price <sym>` | Latest price |
| `alpaca data bars` | Historical price bars (stock/crypto) |
| `alpaca data quotes` | Historical quotes |
| `alpaca data trades` | Historical trades |
| `alpaca data snapshot <sym>` | Full snapshot |
| `alpaca data latest trade <sym>` | Latest trade |
| `alpaca data latest quote <sym>` | Latest quote |
| `alpaca data latest bar <sym>` | Latest bar |
| `alpaca data option bars` | Option historical bars |
| `alpaca data option trades` | Option historical trades |
| `alpaca data option snapshot` | Option snapshots |
| `alpaca data option chain <sym>` | Option chain (market data) |
| `alpaca data option latest-quotes` | Latest option quotes |
| `alpaca data option latest-trades` | Latest option trades |
| `alpaca data forex rates` | Historical forex rates |
| `alpaca data forex latest` | Latest forex rates |
| `alpaca data crypto-orderbook` | Latest crypto orderbooks |
| `alpaca data auctions` | Stock auction data |
| `alpaca data corporate-actions` | Corporate actions (market data) |
| `alpaca data fixed-income` | Fixed income prices |
| `alpaca screener most-actives` | Most active stocks |
| `alpaca screener movers` | Top market movers |
| `alpaca news` | Market news |

### Account & Assets

| Command | Description |
|---------|-------------|
| `alpaca account get` | Account details (equity, buying power) |
| `alpaca account config get` | Account configuration |
| `alpaca account config set` | Update account settings |
| `alpaca activity list` | Account activity (fills, dividends, etc.) |
| `alpaca asset list` | Browse equities and crypto |
| `alpaca asset get <sym>` | Asset details |
| `alpaca asset treasury` | US Treasury bonds |
| `alpaca asset bond` | US Corporate bonds |
| `alpaca portfolio history` | Portfolio value history |
| `alpaca corporate-action list` | Corporate actions announcements |
| `alpaca watchlist create/list/get/add/remove/update/delete` | Watchlist management |
| `alpaca wallet list/transfer/transfers/whitelist` | Crypto funding |

### Utilities

| Command | Description |
|---------|-------------|
| `alpaca clock` | Market clock (supports `--markets` for v3) |
| `alpaca calendar` | Trading calendar (supports `--market` for v3) |
| `alpaca profile login` | Authenticate with API key/secret |
| `alpaca profile switch <name>` | Switch between profiles |
| `alpaca api get <path>` | Raw API access |
| `alpaca update` | Self-update |
| `alpaca version` | Print version |

Every command supports `--help` for full flag documentation.

## Output Formats

```bash
alpaca positions              # Pretty table (default)
alpaca positions --json       # JSON for scripts and agents
alpaca positions --csv        # CSV for spreadsheets
```

## Configuration

### Profiles

```bash
alpaca profile login                                 # Paper trading (default)
alpaca profile login --name live --live               # Live trading
alpaca profile login --name staging --base-url https://staging-api.example.com
alpaca profile switch live                            # Switch default profile
```

Credentials are stored in `~/.config/alpaca/profiles/`.

### Environment Variables

| Variable | Description |
|----------|-------------|
| `ALPACA_API_KEY` | API key (overrides profile) |
| `ALPACA_SECRET_KEY` | Secret key (overrides profile) |
| `ALPACA_BASE_URL` | Trading API base URL |
| `ALPACA_DATA_URL` | Market data API base URL |
| `ALPACA_PROFILE` | Profile name to use |
| `ALPACA_OUTPUT` | Default output format (`table`, `json`, `csv`) |
| `ALPACA_CONFIG_DIR` | Config directory (default: `~/.config/alpaca`) |
| `ALPACA_VERBOSE` | Enable verbose HTTP tracing (any non-empty value) |

Precedence: flags > env vars > profile config > defaults.

## Shell Completions

```bash
alpaca completion bash > ~/.bash_completion.d/alpaca  # Bash
alpaca completion zsh > "${fpath[1]}/_alpaca"          # Zsh
alpaca completion fish > ~/.config/fish/completions/alpaca.fish  # Fish
```

Enum-valued flags auto-complete with valid values (e.g. `--side` → `buy`/`sell`, `--type` → `market`/`limit`/`stop`/etc.).

## Agent & Automation

Designed for scripting, CI pipelines, and AI agent integration.

### Auth (no disk, no prompts)

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
```

### Clean Output

Use `--json` for structured data and `--quiet` to suppress all non-data output (warnings, hints, color):

```bash
alpaca positions --json --quiet
alpaca data latest trade AAPL --json --quiet
```

### Structured Errors

When `--json` or `--quiet` is set, errors are JSON on stderr:

```json
{"error":"rate limited","code":0,"status":429,"hint":"Rate limited. Reduce request frequency or add delays between calls."}
```

### Unattended Operations

The CLI is fully non-interactive. Destructive operations on live accounts require `--confirm`:

```bash
alpaca order cancel-all --confirm
alpaca position close-all --confirm
```

### Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | API or general error |
| `2` | Authentication error (401/403) |

### Verbose Tracing

Debug API calls with `--verbose` or `ALPACA_VERBOSE=1`:

```bash
alpaca account get --verbose
# stderr: GET https://paper-api.alpaca.markets/v2/account → 200 (142ms)
```

### Resilience

The CLI automatically retries on 429 (rate limit) and 5xx errors with exponential backoff (max 3 attempts). The `Retry-After` header is respected for rate limits.

### Example: Agent Workflow

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...

# Check if market is open
clock=$(alpaca clock --json --quiet)
is_open=$(echo "$clock" | jq -r '.is_open')

# Place order if open
if [ "$is_open" = "true" ]; then
  alpaca buy AAPL 10 --json --quiet --confirm
fi

# Handle errors programmatically
if ! result=$(alpaca order get abc123 --json --quiet 2>err.json); then
  status=$(jq .status err.json)
  echo "Failed with HTTP $status"
fi
```

## Development

```bash
make build            # Build binary to bin/alpaca
make install          # Install to $GOPATH/bin
make test             # Run unit tests
make test-integration # Run integration tests (requires ALPACA_TEST_API_KEY)
make lint             # Run linter
make generate         # Regenerate typed API clients from OpenAPI specs
make spec-update      # Fetch latest OpenAPI specs from Alpaca docs
```

### Integration Tests

```bash
export ALPACA_TEST_API_KEY=PK...
export ALPACA_TEST_SECRET_KEY=...
export ALPACA_TEST_BASE_URL=https://paper-api.alpaca.markets  # optional
make test-integration
```

## Self-Update

```bash
alpaca update          # Download and install latest version
alpaca update --check  # Check without installing
```

## License

Apache 2.0
