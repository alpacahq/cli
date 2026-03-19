# Alpaca CLI

A single-binary CLI for the Alpaca Trading API. Think `gh` for GitHub, `stripe` for Stripe — but for trading.

## Core Principle: Generate Everything

The CLI is driven by OpenAPI specs. Maximize what's generated, minimize what's hand-written. **Do not edit generated files directly.** Change the specs or the generator, then `make generate`.

## Design Philosophy

- **Agent-first**: the primary consumer is an AI agent. All parameters are explicit `--flag value` — no positional arguments. Exception: `alpaca api [METHOD] <path>` uses positional args because it's a raw escape hatch, not a generated command.
- **OAS specs are read-only inputs**: never edit the specs in this repo. Fix bugs upstream and re-import.
- **No backward compatibility**: pre-1.0. No aliases, shims, or deprecation wrappers. Just make the change.

## After every change

```
make check     # lint + test + build
```

Fix any failures you introduce before moving on.

When a refactor changes command names, flags, or output shape, review `test/integration/` and update any affected tests so they stay in sync.

## Output contract

Commands fall into two categories with different output rules:

**API commands** (trading, data, account, watchlist, etc.) — call Alpaca endpoints and return structured JSON on stdout. These are the agent pipeline. They support `--csv`, `--jq`, `--quiet`, and `--schema`. Errors go to stderr as JSON.

**Operational commands** (`version`, `doctor`, `profile *`, `update`, `completion`, `--help`, `--help-all`, `--schema`) — manage the CLI itself. These emit human-readable text on stdout. The machine-readable signal is the **exit code** (0 = success, non-zero = failure), not the output format. Do not convert these to JSON — an agent that needs to verify connectivity runs `alpaca account get --quiet`, not `alpaca doctor`.

The one exception is `update --check`, which emits JSON because agents need to programmatically decide whether to upgrade.

**Rule of thumb:** if a command hits the Alpaca API and returns API data, it emits JSON. If it manages the CLI's own state or helps a human troubleshoot, it emits text.

## Design Notes

- **`FlagDef.OASName` must stay**: Flag names are kebab-case (`page-token`), OAS names are snake_case (`page_token`). The mapping `_ → -` is lossy — if an upstream OAS param ever uses a hyphen, runtime reversal (`- → _`) would silently send the wrong query key. Keep both fields.

## Integration tests (`test/integration/`)

Gated by `//go:build integration` and require `ALPACA_TEST_API_KEY`+`ALPACA_TEST_SECRET_KEY` or `ALPACA_TEST_ACCESS_TOKEN`. Default target is `paper-api.alpaca.markets`.

Rules:
- **Read-only tests must call `t.Parallel()`** — every test that only fetches data should run concurrently.
- **Write tests must clean up** — use `t.Cleanup()` to cancel orders, delete watchlists, close positions. Never leave side-effects.
- **Flat tests for independent calls, sub-tests for sequential chains** — don't use `t.Run` unless steps depend on prior state (e.g., create → get → delete).
- **Use the helpers** — `alpaca()`, `alpacaFail()`, `parseJSONMap()`, `requireFields()`, `daysAgo()`, `pollFor()`. Don't reinvent them.
- **Some endpoints are unavailable on paper** — wallet, crypto-perp trading, forex, fixed-income, logo, treasury, and bonds return 403/404. Don't write tests that require these unless you have a compatible test account.
- **One file per feature area** — `order_test.go`, `data_option_test.go`, `crypto_perp_data_test.go`, etc. Keep files focused.

## Keep docs in sync

When a change affects CLI behavior, update any stale docs:

- `README.md` — user-facing documentation
- `skills/alpaca-cli/SKILL.md` — agent-facing skill

Code is the source of truth. If a doc contradicts the code, update the doc.
