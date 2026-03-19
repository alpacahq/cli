# Alpaca CLI

CLI for [Alpaca](https://alpaca.markets) Trading API. Trade stocks & crypto, access market data, and manage your account from the command line.

> [!WARNING]
> **Alpha Preview** — This CLI is under active development. Commands, flags, and output formats may change or be removed without notice between releases. Do not depend on current behavior in production workflows.

### Built for agents

This CLI is designed primarily as a tool for AI agents and automation pipelines, not as a general-purpose interactive terminal tool. It has no confirmation prompts, no "are you sure?" dialogs, and no interactive mode. Every command executes immediately and returns structured output.

**This means destructive commands are truly destructive.** For example:

- `alpaca position close-all` will liquidate your entire portfolio instantly
- `alpaca order cancel-all` will cancel every open order without listing them first

There are no guardrails — the CLI trusts that the caller (human or agent) knows what it's asking for. If you're using this with live trading credentials, understand that a single mistyped command can have real financial consequences.

## Design Philosophy

The CLI is driven by OpenAPI specs — types, clients, param structs, flag definitions, enum completions, and validation are all generated from `api/specs/*.json`.

## Install

**Go** (recommended):

```bash
go install github.com/alpacahq/cli/cmd/alpaca@latest
```

**Homebrew** (macOS / Linux) — *coming soon*:

```bash
brew install alpacahq/tap/alpaca
```

## Quick Start

```bash
# Authenticate
alpaca profile login

# Check your account
alpaca account get

# Submit an order
alpaca order submit --symbol AAPL --side buy --qty 10 --type market

# List your positions
alpaca position list

# List open orders
alpaca order list --status open

# Get market data
alpaca data bars --symbol AAPL --start 2025-01-01 --timeframe 1Day

# Check if market is open
alpaca clock
```

## Safety

**Paper trading is the default.** When you run `alpaca profile login`, credentials are stored for paper trading (`paper-api.alpaca.markets`). Live trading requires API keys.

**Authentication methods:**

- **OAuth (default, paper only)** — `alpaca profile login` opens a browser for OAuth authorization. No keys to copy/paste. The CLI starts a temporary localhost server to receive the callback, exchanges the authorization code for a token, and stores it locally. OAuth is currently restricted to paper trading while the OAuth flow is hardened (PKCE support pending).
- **API keys (paper and live)** — `alpaca profile login --api-key` prompts for API key + secret. Required for live trading. Recommended for CI/automation where a browser isn't available.

**Credential safety:**

- Credentials are stored in `~/.config/alpaca/profiles/` with restricted file permissions (0600).
- For CI/automation, use environment variables instead of stored profiles — no secrets touch disk:

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
alpaca position list
```

- Passing `--secret` via flags is discouraged (shell history exposure). Use interactive login or env vars.

**OAuth security model:** The CLI is a first-party public OAuth client — the embedded `client_id` and `client_secret` are app identifiers, not security credentials. This is standard for native CLI apps ([RFC 8252 §8.5](https://datatracker.ietf.org/doc/html/rfc8252#section-8.5)). The user controls access through the browser consent screen. OAuth is restricted to paper trading while PKCE support is pending; live trading requires API keys.

## Commands

### Trading

| Command | Description |
|---------|-------------|
| `alpaca order submit` | Submit any order type |
| `alpaca order list` | List orders |
| `alpaca order get` | Get order by ID |
| `alpaca order get-by-client-id` | Get order by client order ID |
| `alpaca order cancel` | Cancel an order |
| `alpaca order cancel-all` | Cancel all open orders |
| `alpaca order replace` | Replace an existing order |
| `alpaca position list` | List open positions |
| `alpaca position get` | Get position for a symbol |
| `alpaca position close` | Close a position |
| `alpaca position close-all` | Close all positions |
| `alpaca option contracts` | List option contracts |
| `alpaca option get` | Option contract details |
| `alpaca option exercise` | Exercise an option |
| `alpaca option do-not-exercise` | Mark option as do-not-exercise |
| `alpaca clock` | US market clock |
| `alpaca clock markets` | Multi-market clock |
| `alpaca calendar` | US trading calendar |
| `alpaca calendar market` | Market-specific calendar |

### Market Data

| Command | Description |
|---------|-------------|
| `alpaca data bars` | Historical bars (single symbol) |
| `alpaca data quotes` | Historical quotes (single symbol) |
| `alpaca data trades` | Historical trades (single symbol) |
| `alpaca data snapshot` | Stock snapshot (single symbol) |
| `alpaca data latest-bar` | Latest bar (single symbol) |
| `alpaca data latest-quote` | Latest quote (single symbol) |
| `alpaca data latest-trade` | Latest trade (single symbol) |
| `alpaca data multi-bars` | Historical bars (multi-symbol) |
| `alpaca data multi-quotes` | Historical quotes (multi-symbol) |
| `alpaca data multi-trades` | Historical trades (multi-symbol) |
| `alpaca data multi-snapshots` | Snapshots (multi-symbol) |
| `alpaca data latest-bars` | Latest bars (multi-symbol) |
| `alpaca data latest-quotes` | Latest quotes (multi-symbol) |
| `alpaca data latest-trades` | Latest trades (multi-symbol) |
| `alpaca data auction` | Stock auction (single symbol) |
| `alpaca data auctions` | Stock auctions (multi-symbol) |
| `alpaca data crypto bars` | Crypto historical bars |
| `alpaca data crypto quotes` | Crypto historical quotes |
| `alpaca data crypto trades` | Crypto historical trades |
| `alpaca data crypto snapshots` | Crypto snapshots |
| `alpaca data crypto latest-bars` | Latest crypto bars |
| `alpaca data crypto latest-quotes` | Latest crypto quotes |
| `alpaca data crypto latest-trades` | Latest crypto trades |
| `alpaca data crypto-orderbook` | Latest crypto orderbooks |
| `alpaca data option bars` | Option historical bars |
| `alpaca data option trades` | Option historical trades |
| `alpaca data option snapshot` | Option snapshots |
| `alpaca data option chain` | Option chain (greeks and pricing) |
| `alpaca data option latest-quotes` | Latest option quotes |
| `alpaca data option latest-trades` | Latest option trades |
| `alpaca data option exchanges` | Option exchanges |
| `alpaca data option conditions` | Option trade conditions |
| `alpaca data forex rates` | Historical forex rates |
| `alpaca data forex latest` | Latest forex rates |
| `alpaca data corporate-actions` | Corporate actions (market data) |
| `alpaca data fixed-income` | Fixed income prices |
| `alpaca data logo` | Company logo URL |
| `alpaca data meta exchanges` | Exchange code reference |
| `alpaca data meta conditions` | Trade/quote condition codes |
| `alpaca data screener most-actives` | Most active stocks |
| `alpaca data screener movers` | Top market movers |
| `alpaca data news` | Market news |

### Account & Assets

| Command | Description |
|---------|-------------|
| `alpaca account get` | Account details (equity, buying power) |
| `alpaca account config get` | Account configuration |
| `alpaca account config set` | Update account settings |
| `alpaca account activity list` | Account activity (fills, dividends, etc.) |
| `alpaca account activity list-by-type` | Activity filtered by single type |
| `alpaca account portfolio` | Portfolio equity and P&L history |
| `alpaca asset list` | Browse equities and crypto |
| `alpaca asset get` | Asset details |
| `alpaca asset treasury` | US Treasury bonds |
| `alpaca asset bond` | US Corporate bonds |
| `alpaca corporate-action list` | Corporate action announcements |
| `alpaca corporate-action get` | Get a specific announcement |
| `alpaca watchlist list` | List all watchlists |
| `alpaca watchlist get` | Get watchlist details |
| `alpaca watchlist create` | Create a watchlist |
| `alpaca watchlist update` | Update a watchlist |
| `alpaca watchlist delete` | Delete a watchlist |
| `alpaca watchlist add` | Add symbol to watchlist |
| `alpaca watchlist remove` | Remove symbol from watchlist |
| `alpaca watchlist get-by-name` | Get watchlist by name |
| `alpaca watchlist update-by-name` | Update watchlist by name |
| `alpaca watchlist delete-by-name` | Delete watchlist by name |
| `alpaca watchlist add-by-name` | Add symbol to watchlist by name |
| `alpaca watchlist remove-by-name` | Remove symbol from watchlist by name |
| `alpaca wallet list` | List crypto wallets |
| `alpaca wallet transfer list` | List crypto transfers |
| `alpaca wallet transfer get` | Get a crypto transfer |
| `alpaca wallet transfer create` | Create a crypto transfer |
| `alpaca wallet transfer estimate` | Estimate transfer fees |
| `alpaca wallet whitelist list` | List whitelisted addresses |
| `alpaca wallet whitelist add` | Add a whitelisted address |
| `alpaca wallet whitelist delete` | Remove a whitelisted address |

### Crypto Perpetuals

| Command | Description |
|---------|-------------|
| `alpaca crypto-perp vitals` | Account vitals |
| `alpaca crypto-perp leverage` | Get leverage settings |
| `alpaca crypto-perp set-leverage` | Set leverage for an asset |
| `alpaca crypto-perp wallet list` | List perp funding wallets |
| `alpaca crypto-perp wallet transfer list` | List perp transfers |
| `alpaca crypto-perp wallet transfer get` | Get a perp transfer |
| `alpaca crypto-perp wallet transfer create` | Create a perp transfer |
| `alpaca crypto-perp wallet transfer estimate` | Estimate perp transfer fees |
| `alpaca crypto-perp wallet whitelist list` | List whitelisted perp addresses |
| `alpaca crypto-perp wallet whitelist add` | Add a whitelisted perp address |
| `alpaca crypto-perp wallet whitelist delete` | Remove a whitelisted perp address |
| `alpaca crypto-perp data latest-bars` | Latest perp bars |
| `alpaca crypto-perp data latest-futures-pricing` | Latest futures pricing |
| `alpaca crypto-perp data latest-orderbooks` | Latest perp orderbooks |
| `alpaca crypto-perp data latest-quotes` | Latest perp quotes |
| `alpaca crypto-perp data latest-trades` | Latest perp trades |

### Utilities

| Command | Description |
|---------|-------------|
| `alpaca profile login` | Authenticate via browser OAuth (or `--api-key`) |
| `alpaca profile logout [name]` | Remove a profile |
| `alpaca profile list` | List all profiles |
| `alpaca profile switch <name>` | Switch between profiles |
| `alpaca api [METHOD] <path>` | Raw API request (GET, POST, PATCH, DELETE) |
| `alpaca doctor` | Check config and API connectivity |
| `alpaca update` | Check for updates and show upgrade instructions |
| `alpaca version` | Print version (`alpaca --version` also works) |

Every command supports `--help` for full flag documentation.

## Output Formats

```bash
alpaca position list                          # JSON (default)
alpaca position list --csv                    # CSV for spreadsheets
alpaca position list --jq '.[0].symbol'       # Filter JSON with jq expressions
```

## Configuration

### Profiles

```bash
alpaca profile login                                 # OAuth via browser, paper (default)
alpaca profile login --api-key                       # API key/secret, paper
alpaca profile login --api-key --live                # API key/secret, live trading
alpaca profile login --api-key --name live --live    # API key for live with custom name
alpaca profile login --name staging --base-url https://staging-api.example.com
alpaca profile switch live                           # Switch default profile
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
| `ALPACA_OUTPUT` | Default output format (`json`, `csv`) |
| `ALPACA_CONFIG_DIR` | Config directory (default: `~/.config/alpaca`) |
| `ALPACA_VERBOSE` | Show HTTP request summaries on stderr (any non-empty value) |
| `ALPACA_DEBUG` | Show HTTP request/response headers and bodies on stderr (any non-empty value) |
| `ALPACA_TRACE` | Show HTTP timing breakdown on stderr — DNS, TLS, TTFB (any non-empty value) |
Global flags: `--csv`, `--jq`, `--profile`, `--verbose`, `--debug`, `--trace`, `--quiet`, `--schema`, `--timeout`.

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
alpaca order submit --symbol AAPL --side <TAB>  # Should show buy/sell
```

If completions don't appear, check that your shell is sourcing the file and that `compinit` (zsh) or `complete` (bash) is loaded.

Enum-valued flags auto-complete with valid values (e.g. `--side` → `buy`/`sell`, `--type` → `market`/`limit`/`stop`/etc.).

## Agent & Automation

Designed for scripting, CI pipelines, and AI agent integration. For AI agents, see the [`alpaca-cli` Agent Skill](skills/alpaca-cli/SKILL.md) for structured install, auth, and usage guidance in [Agent Skills](https://agentskills.io) format.

### Auth (no disk, no prompts)

```bash
export ALPACA_API_KEY=PK...
export ALPACA_SECRET_KEY=...
```

### Clean Output

JSON is the default output format. Use `--quiet` to suppress all non-data output (warnings, hints). Use `--jq` to filter JSON inline without piping to `jq`:

```bash
alpaca position list --quiet
alpaca data latest-trade --symbol AAPL --quiet
alpaca order list --jq '[.[] | {id, symbol, side, qty}]' --quiet
```

### Structured Errors

Errors are always JSON on stderr:

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
| `2` | Authentication error (401) |

### Diagnostics

Three orthogonal flags for different debugging needs. Combine as needed.

`--verbose` — request summaries (what happened):

```bash
alpaca account get --verbose
# stderr: GET https://paper-api.alpaca.markets/v2/account → 200 (142ms)
```

`--trace` — timing breakdown (why is it slow):

```bash
alpaca account get --trace
# stderr: trace: GET https://paper-api.alpaca.markets/v2/account
# stderr:   dns:     4ms
# stderr:   tcp:     98ms  (35.194.67.18:443)
# stderr:   tls:     137ms
# stderr:   ttfb:    125ms
# stderr:   total:   365ms → 200
```

`--debug` — wire-level detail (what was sent/received):

```bash
alpaca account get --debug
# stderr: → GET https://paper-api.alpaca.markets/v2/account
# stderr: → User-Agent: alpaca-cli/0.0.1
# stderr: ← Content-Type: application/json
# stderr: ← {"id":"...","equity":"10000.00",...}
```

Credentials are always scrubbed from stderr output.

### Dry Run

Preview an order without submitting:

```bash
alpaca order submit --symbol AAPL --side buy --qty 10 --type limit --limit-price 185.00 --dry-run
```

### Stdin Pipe Support

Pipe JSON payloads into `alpaca api POST/PATCH`:

```bash
echo '{"symbol":"AAPL","qty":"1","side":"buy","type":"market","time_in_force":"day"}' \
  | alpaca api POST /v2/orders
```

If both `--body` and stdin are provided, `--body` takes precedence.

### Response Schemas

Any command with OAS-generated flags supports `--schema` to show the response fields without making an API call:

```bash
alpaca order list --schema       # Show Order response fields
alpaca asset list --schema       # Show Asset response fields
alpaca data bars --schema        # Show bars response fields
```

### Timeout

Override the default 30-second HTTP timeout:

```bash
alpaca data bars --symbol AAPL --start 2020-01-01 --end 2025-01-01 --timeout 120
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
clock=$(alpaca clock --quiet)
is_open=$(echo "$clock" | jq -r '.is_open')

# Place order if open
if [ "$is_open" = "true" ]; then
  alpaca order submit --symbol AAPL --side buy --qty 10 --type market --quiet
fi

# Preview before submitting
alpaca order submit --symbol AAPL --side buy --qty 10 --type limit --limit-price 185.00 --dry-run

# Pipe complex payloads
cat order.json | alpaca api POST /v2/orders --quiet

# Handle errors programmatically
if ! result=$(alpaca order get --order-id abc123 --quiet 2>err.json); then
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

Check for updates explicitly when you want upgrade guidance. The upgrade command is tailored to your install method:

| Install method | Upgrade command |
|---|---|
| go install | `go install github.com/alpacahq/cli/cmd/alpaca@latest` |
| Homebrew | `brew upgrade alpaca` *(coming soon)* |

```bash
alpaca update          # Check for updates and show upgrade command
alpaca update --check  # Machine-readable update check (JSON)
```

## License

Apache 2.0
