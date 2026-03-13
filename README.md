# Alpaca CLI

CLI for [Alpaca](https://alpaca.markets) Trading API. Trade stocks & crypto, access market data, and manage your account from the command line.

> **Alpha Preview** — This CLI is under active development and available as an early alpha release for testing and feedback. Commands, flags, and output formats are subject to change. Use in production workflows at your own risk.

## Design Philosophy

The CLI is driven by OpenAPI specs — types, clients, param structs, flag definitions, enum completions, and validation are all generated from `api/specs/*.json`.

## Install

**Homebrew** (macOS / Linux):

> Homebrew tap (`alpacahq/homebrew-tap`) is being set up. Until it's published, use `go install` or download a binary from Releases.

```bash
brew install alpacahq/tap/alpaca
```

**Go**:

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

**Supported platforms:**

| OS | Shells | Man pages |
|----|--------|-----------|
| macOS | bash, zsh, fish, PowerShell | Yes |
| Linux | bash, zsh, fish, PowerShell | Yes |
| Windows | PowerShell | No |

Override auto-detection with `--shell`:

```bash
alpaca setup --shell fish
```

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
alpaca position list --json
```

- Passing `--secret` via flags is discouraged (shell history exposure). Use interactive login or env vars.

**OAuth security model:** The CLI is a first-party public OAuth client — the embedded `client_id` and `client_secret` serve only as app identifiers, not security credentials. This is standard for native CLI apps (GitHub CLI, Google Cloud CLI, Stripe CLI all do the same). The user always controls access through the browser consent screen; the client credentials alone grant no access to any account. For details, see:

- [RFC 8252 Section 8.5](https://datatracker.ietf.org/doc/html/rfc8252#section-8.5) — *"Secrets that are statically included as part of an app distributed to multiple users should not be treated as confidential secrets."*
- [GitHub CLI approach](https://github.com/cli/oauth/issues/1#issuecomment-749151295) — GitHub's `cli/oauth` maintainer on why embedding client secrets in open-source CLIs is acceptable.

Note: OAuth login is restricted to paper trading while the flow does not support PKCE (pending Alpaca OAuth server support). Live trading requires API keys. The authorization code grant with state parameter and localhost redirect URI validation provides adequate security for paper trading.

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
| `alpaca option contracts <symbol>` | List option contracts |
| `alpaca option get <id>` | Option contract details |
| `alpaca option exercise <id>` | Exercise an option |
| `alpaca option do-not-exercise <id>` | Mark option as do-not-exercise |
| `alpaca clock` | Market clock (supports `--markets` for v3) |
| `alpaca calendar` | Trading calendar (supports `--market` for v3) |

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
| `alpaca data option exchanges` | Option exchanges |
| `alpaca data option conditions` | Option trade conditions |
| `alpaca data forex rates` | Historical forex rates |
| `alpaca data forex latest` | Latest forex rates |
| `alpaca data crypto-orderbook` | Latest crypto orderbooks |
| `alpaca data auctions` | Stock auction data |
| `alpaca data corporate-actions` | Corporate actions (market data) |
| `alpaca data fixed-income` | Fixed income prices |
| `alpaca data logo <symbol>` | Company logo URL |
| `alpaca data meta exchanges` | Exchange code reference |
| `alpaca data meta conditions <ticktype>` | Trade/quote condition codes |
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
| `alpaca account portfolio` | Portfolio equity and P&L history |
| `alpaca asset list` | Browse equities and crypto |
| `alpaca asset get <symbol>` | Asset details |
| `alpaca asset treasury` | US Treasury bonds |
| `alpaca asset bond` | US Corporate bonds |
| `alpaca corporate-action list` | Corporate action announcements |
| `alpaca corporate-action get <id>` | Get a specific announcement |
| `alpaca watchlist list` | List all watchlists |
| `alpaca watchlist get <id>` | Get watchlist details |
| `alpaca watchlist create <name>` | Create a watchlist |
| `alpaca watchlist update <id>` | Update a watchlist |
| `alpaca watchlist delete <id>` | Delete a watchlist |
| `alpaca watchlist add <id> <symbol>` | Add symbol to watchlist |
| `alpaca watchlist remove <id> <symbol>` | Remove symbol from watchlist |
| `alpaca watchlist get-by-name <name>` | Get watchlist by name |
| `alpaca watchlist update-by-name <name>` | Update watchlist by name |
| `alpaca watchlist delete-by-name <name>` | Delete watchlist by name |
| `alpaca watchlist add-by-name <name> <symbol>` | Add symbol to watchlist by name |
| `alpaca wallet list` | List crypto wallets |
| `alpaca wallet transfer list` | List crypto transfers |
| `alpaca wallet transfer get <id>` | Get a crypto transfer |
| `alpaca wallet transfer create` | Create a crypto transfer |
| `alpaca wallet transfer estimate` | Estimate transfer fees |
| `alpaca wallet whitelist list` | List whitelisted addresses |
| `alpaca wallet whitelist add` | Add a whitelisted address |
| `alpaca wallet whitelist delete <id>` | Remove a whitelisted address |

### Utilities

| Command | Description |
|---------|-------------|
| `alpaca profile login` | Authenticate via browser OAuth (or `--api-key`) |
| `alpaca profile logout [name]` | Remove a profile |
| `alpaca profile status` | Show the active profile |
| `alpaca profile list` | List all profiles |
| `alpaca profile switch <name>` | Switch between profiles |
| `alpaca profile set <key> <value>` | Update a profile setting |
| `alpaca api get <path>` | GET request to any endpoint |
| `alpaca api post <path>` | POST request to any endpoint |
| `alpaca api patch <path>` | PATCH request to any endpoint |
| `alpaca api delete <path>` | DELETE request to any endpoint |
| `alpaca doctor` | Check config and API connectivity |
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
| `ALPACA_OUTPUT` | Default output format (`table`, `json`, `csv`) |
| `ALPACA_CONFIG_DIR` | Config directory (default: `~/.config/alpaca`) |
| `ALPACA_VERBOSE` | Show HTTP request summaries on stderr (any non-empty value) |
| `ALPACA_DEBUG` | Show HTTP request/response headers and bodies on stderr (any non-empty value) |
| `ALPACA_TRACE` | Show HTTP timing breakdown on stderr — DNS, TLS, TTFB (any non-empty value) |
| `ALPACA_NO_UPDATE_NOTIFY` | Suppress background update notices (any non-empty value) |

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

Designed for scripting, CI pipelines, and AI agent integration. For AI agents, see the [`alpaca-cli` Agent Skill](skills/alpaca-cli/SKILL.md) for structured install, auth, and usage guidance in [Agent Skills](https://agentskills.io) format.

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
# stderr: → User-Agent: alpaca-cli/0.1.0
# stderr: ← Content-Type: application/json
# stderr: ← {"id":"...","equity":"10000.00",...}
```

Credentials are always scrubbed from stderr output.

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

### Pagination

Data commands support auto-pagination with `--all`:

```bash
alpaca data bars AAPL --start 2025-01-01 --all
alpaca data trades AAPL --start 2025-01-01 --all --max 5000
alpaca data news --symbols AAPL --all --max 100
```

`--max` limits the total number of items returned (default: 10,000).

### Response Schemas

Any command with OAS-generated flags supports `--schema` to show the response fields without making an API call:

```bash
alpaca order list --schema       # Show Order response fields
alpaca asset list --schema       # Show Asset response fields
alpaca data bars --schema        # Show bars response fields
```

### Diagnostics

Check your CLI setup:

```bash
alpaca doctor
```

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

The CLI checks for updates in the background (once every 24 hours) and shows a notice on stderr when a newer version is available. The upgrade command is tailored to your install method:

| Install method | Upgrade command |
|---|---|
| Homebrew | `brew upgrade alpaca` |
| go install | `go install github.com/alpacahq/cli/cmd/alpaca@latest` |
| Binary download | `alpaca update` |

```bash
alpaca update          # Download and install latest version (binary installs)
alpaca update --check  # Check without installing
```

For scripts and agents, use the machine-readable check:

```bash
alpaca update --check --json
```

Suppress update notices with `ALPACA_NO_UPDATE_NOTIFY=1` or `--quiet`.

## License

Apache 2.0
