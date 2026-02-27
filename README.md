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
alpaca auth login

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

| Command | Description |
|---------|-------------|
| `alpaca auth login` | Authenticate with API key/secret |
| `alpaca account get` | Account details (equity, buying power) |
| `alpaca buy <sym> <qty>` | Buy shares (market order) |
| `alpaca sell <sym> <qty>` | Sell shares (market order) |
| `alpaca price <sym>` | Latest price |
| `alpaca positions` | List open positions |
| `alpaca orders` | List open orders |
| `alpaca order submit` | Submit any order type |
| `alpaca order cancel <id>` | Cancel an order |
| `alpaca position close <sym>` | Close a position |
| `alpaca data bars` | Historical price bars |
| `alpaca data latest trade <sym>` | Latest trade |
| `alpaca data latest quote <sym>` | Latest quote |
| `alpaca data snapshot <sym>` | Full snapshot |
| `alpaca news` | Market news |
| `alpaca asset list` | Browse assets |
| `alpaca watchlist create` | Create a watchlist |
| `alpaca option chain <sym>` | Options chain |
| `alpaca clock` | Market clock |
| `alpaca calendar` | Trading calendar |
| `alpaca api get <path>` | Raw API access |

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
alpaca auth login                                    # Paper trading (default)
alpaca auth login --profile live --environment live   # Live trading
alpaca auth switch live                              # Switch default profile
```

Credentials are stored in `~/.config/alpaca/profiles/`.

### Environment Variables

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
export ALPACA_BASE_URL=https://paper-api.alpaca.markets
```

Precedence: flags > env vars > profile config > defaults.

## Agent / Automation

Designed for scripting and AI agent integration:

```bash
result=$(alpaca positions --json)
price=$(alpaca data latest trade AAPL --json)

# Exit codes: 0=success, 1=API error, 2=auth error, 3=validation, 4=network
```

## Development

```bash
make build            # Build binary to bin/alpaca
make install          # Install to $GOPATH/bin
make test             # Run unit tests
make test-integration # Run integration tests (requires ALPACA_TEST_API_KEY)
make lint             # Run linter
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
