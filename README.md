# alpaca

CLI for [Alpaca](https://alpaca.markets) Trading API. Trade stocks & crypto, access market data, and manage your account from the command line.

## Install

```bash
go install github.com/alpacahq/cli/cmd/alpaca@latest
```

Or download a prebuilt binary from [Releases](https://github.com/alpacahq/cli/releases).

## Post-Install Setup

After installing, run `setup` to install shell completions and man pages:

```bash
alpaca setup
```

This auto-detects your shell and installs to user-level directories (no `sudo` needed). It also runs automatically after `alpaca update`.

## Quick Start

```bash
# Authenticate
alpaca profile login

# Check your account
alpaca account get

# Submit an order
alpaca order submit AAPL --side buy --qty 10 --type market

# List your positions
alpaca position list

# List open orders
alpaca order list --status open

# Get market data
alpaca data bars AAPL --start 2025-01-01 --timeframe 1Day

# Check if market is open
alpaca clock
```

## Safety

**Paper trading is the default.** When you run `alpaca profile login`, credentials are stored for paper trading (`paper-api.alpaca.markets`) unless you explicitly pass `--live`.

**Credential safety:**

- For interactive use, `alpaca profile login` stores credentials in `~/.config/alpaca/profiles/` with restricted file permissions (0600).
- For CI/automation, use environment variables instead of stored profiles — no secrets touch disk:

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
alpaca position list --json
```

- Passing `--secret` via flags is discouraged (shell history exposure). Use interactive login or env vars.

## Commands

### Trading

| Command | Description |
|---------|-------------|
| `alpaca order submit <symbol>` | Submit any order type |
| `alpaca order list` | List orders |
| `alpaca order get <id>` | Get order details |
| `alpaca order cancel <id>` | Cancel an order |
| `alpaca order cancel-all` | Cancel all open orders |
| `alpaca order replace <id>` | Replace an existing order |
| `alpaca position list` | List open positions |
| `alpaca position get <symbol>` | Get position for a symbol |
| `alpaca position close <symbol>` | Close a position |
| `alpaca position close-all` | Close all positions |
| `alpaca option chain <symbol>` | Options chain (contracts) |
| `alpaca option get <id>` | Option contract details |
| `alpaca option exercise <id>` | Exercise an option |
| `alpaca option do-not-exercise <id>` | Mark option as do-not-exercise |

### Market Data

| Command | Description |
|---------|-------------|
| `alpaca data bars <symbol>` | Historical price bars (stock/crypto) |
| `alpaca data quotes <symbol>` | Historical quotes |
| `alpaca data trades <symbol>` | Historical trades |
| `alpaca data snapshot <symbol>` | Full snapshot |
| `alpaca data latest trade <symbol>` | Latest trade |
| `alpaca data latest quote <symbol>` | Latest quote |
| `alpaca data latest bar <symbol>` | Latest bar |
| `alpaca data option bars` | Option historical bars |
| `alpaca data option trades` | Option historical trades |
| `alpaca data option snapshot` | Option snapshots |
| `alpaca data option chain <symbol>` | Option chain (greeks and pricing) |
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
| `alpaca asset get <symbol>` | Asset details |
| `alpaca asset treasury` | US Treasury bonds |
| `alpaca asset bond` | US Corporate bonds |
| `alpaca portfolio history` | Portfolio value history |
| `alpaca corporate-action list` | Corporate action announcements |
| `alpaca corporate-action get <id>` | Get a specific announcement |
| `alpaca watchlist list` | List all watchlists |
| `alpaca watchlist get <id>` | Get watchlist details |
| `alpaca watchlist create <name>` | Create a watchlist |
| `alpaca watchlist update <id>` | Update a watchlist |
| `alpaca watchlist delete <id>` | Delete a watchlist |
| `alpaca watchlist add <id> <symbol>` | Add symbol to watchlist |
| `alpaca watchlist remove <id> <symbol>` | Remove symbol from watchlist |
| `alpaca wallet list` | List crypto wallets |
| `alpaca wallet transfer list` | List crypto transfers |
| `alpaca wallet transfer get <id>` | Get a crypto transfer |
| `alpaca wallet transfer create` | Create a crypto transfer |
| `alpaca wallet whitelist list` | List whitelisted addresses |
| `alpaca wallet whitelist add` | Add a whitelisted address |
| `alpaca wallet whitelist delete <id>` | Remove a whitelisted address |

### Utilities

| Command | Description |
|---------|-------------|
| `alpaca clock` | Market clock (supports `--markets` for v3) |
| `alpaca calendar` | Trading calendar (supports `--market` for v3) |
| `alpaca profile login` | Authenticate with API key/secret |
| `alpaca profile logout [name]` | Remove a profile |
| `alpaca profile status` | Show the active profile |
| `alpaca profile list` | List all profiles |
| `alpaca profile switch <name>` | Switch between profiles |
| `alpaca profile set <key> <value>` | Update a profile setting |
| `alpaca api get <path>` | GET request to any endpoint |
| `alpaca api post <path>` | POST request to any endpoint |
| `alpaca api patch <path>` | PATCH request to any endpoint |
| `alpaca api delete <path>` | DELETE request to any endpoint |
| `alpaca setup` | Install completions and man pages |
| `alpaca update` | Self-update |
| `alpaca version` | Print version |

Every command supports `--help` for full flag documentation.

## Output Formats

```bash
alpaca position list              # Pretty table (default)
alpaca position list --json       # JSON for scripts and agents
alpaca position list --csv        # CSV for spreadsheets
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
| `ALPACA_DEBUG` | Full HTTP request/response bodies on stderr (implies verbose) |

Global flags: `--json`, `--csv`, `--profile`, `--verbose`, `--debug`, `--quiet`, `--timeout`.

Precedence: flags > env vars > profile config > defaults.

## Shell Completions

### Bash

```bash
mkdir -p ~/.bash_completion.d
alpaca completion bash > ~/.bash_completion.d/alpaca
source ~/.bash_completion.d/alpaca
```

Add `source ~/.bash_completion.d/alpaca` to your `~/.bashrc` (or `~/.bash_profile` on macOS) to load on every session.

### Zsh

```bash
# Ensure completions directory is in fpath (add to ~/.zshrc before compinit)
alpaca completion zsh > "${fpath[1]}/_alpaca"
```

If `echo $fpath` is empty or the directory doesn't exist, create one:

```bash
mkdir -p ~/.zsh/completions
echo 'fpath=(~/.zsh/completions $fpath)' >> ~/.zshrc
echo 'autoload -Uz compinit && compinit' >> ~/.zshrc
alpaca completion zsh > ~/.zsh/completions/_alpaca
source ~/.zshrc
```

### Fish

```bash
alpaca completion fish > ~/.config/fish/completions/alpaca.fish
```

Fish loads completions automatically from this directory — no restart required.

### PowerShell

```powershell
alpaca completion powershell > alpaca.ps1
. ./alpaca.ps1
```

To persist, add the output to your PowerShell profile (`$PROFILE`).

### Verify it works

After installing, open a new shell and type:

```bash
alpaca <TAB>            # Should show subcommands
alpaca order submit AAPL --side <TAB>  # Should show buy/sell
```

If completions don't appear, check that your shell is sourcing the file and that `compinit` (zsh) or `complete` (bash) is loaded.

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
alpaca position list --json --quiet
alpaca data latest trade AAPL --json --quiet
```

### Structured Errors

When `--json` or `--quiet` is set, errors are JSON on stderr:

```json
{"error":"rate limited","code":0,"status":429,"hint":"Rate limited. Reduce request frequency or add delays between calls."}
```

### Unattended Operations

The CLI is fully non-interactive — no TTY detection, no interactive prompts.

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

For full request/response headers and bodies, use `--debug`:

```bash
alpaca account get --debug
# stderr: → GET https://paper-api.alpaca.markets/v2/account
# stderr: → User-Agent: alpaca-cli/0.1.0
# stderr: GET https://... → 200 (142ms)
# stderr: ← Content-Type: application/json
# stderr: ← {"id":"...","equity":"10000.00",...}
```

Credentials are always scrubbed from debug output.

### Dry Run

Preview an order without submitting:

```bash
alpaca order submit AAPL --side buy --qty 10 --type limit --limit-price 185.00 --dry-run
```

### Stdin Pipe Support

Pipe JSON payloads into `alpaca api post/patch`:

```bash
echo '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}' \
  | alpaca api post /v2/orders --json
```

If both `--data` and stdin are provided, `--data` takes precedence.

### Timeout

Override the default 30-second HTTP timeout:

```bash
alpaca data bars AAPL --start 2020-01-01 --end 2025-01-01 --timeout 120
```

### Resilience

The CLI automatically retries on 429 (rate limit) and 5xx errors with exponential backoff (max 3 attempts). The `Retry-After` header is respected for rate limits.

### JSON Output Stability

JSON output mirrors the Alpaca API response directly. Fields will not be removed or renamed without a major version bump. New fields may be added at any time (treat JSON output as forward-compatible).

### Example: Agent Workflow

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...

# Check if market is open
clock=$(alpaca clock --json --quiet)
is_open=$(echo "$clock" | jq -r '.is_open')

# Place order if open
if [ "$is_open" = "true" ]; then
  alpaca order submit AAPL --side buy --qty 10 --type market --json --quiet
fi

# Preview before submitting
alpaca order submit AAPL --side buy --qty 10 --type limit --limit-price 185.00 --dry-run

# Pipe complex payloads
cat order.json | alpaca api post /v2/orders --json --quiet

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
