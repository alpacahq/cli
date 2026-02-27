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

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
export ALPACA_BASE_URL=https://paper-api.alpaca.markets
```

Precedence: flags > env vars > profile config > defaults.

## Shell Completions

```bash
alpaca completion bash > ~/.bash_completion.d/alpaca  # Bash
alpaca completion zsh > "${fpath[1]}/_alpaca"          # Zsh
alpaca completion fish > ~/.config/fish/completions/alpaca.fish  # Fish
```

Enum-valued flags auto-complete with valid values (e.g. `--side` → `buy`/`sell`, `--type` → `market`/`limit`/`stop`/etc.).

## Agent / Automation

Designed for scripting and AI agent integration:

```bash
result=$(alpaca positions --json)
price=$(alpaca data latest trade AAPL --json)

# Exit codes: 0=success, 1=API error, 2=auth error (401/403)
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
