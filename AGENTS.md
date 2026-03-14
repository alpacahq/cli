# Alpaca CLI

A single-binary CLI for the Alpaca Trading API. Think `gh` for GitHub, `stripe` for Stripe — but for trading.

## Core Principle: Generate Everything

The CLI is driven by OpenAPI specs. Maximize what's generated, minimize what's hand-written.

```
api/specs/*.json  →  cmd/generate/main.go  →  internal/api/     (types, clients, param structs, flag defs)
                                            →  internal/cmd/     (params_generated.go — FromFlags bridge)
```

**Do not edit generated files directly.** Change the specs or the generator, then `make generate`.

### Generated files

| File | Contents |
|---|---|
| `internal/api/trading_types.go` | Trading API Go types |
| `internal/api/trading_client.go` | Trading API client methods + Params structs + Values() |
| `internal/api/marketdata_types.go` | Market data Go types |
| `internal/api/marketdata_client.go` | Market data client methods + Params structs + Values() |
| `internal/api/descriptions.go` | FlagDef slices, typed Op structs with summaries |
| `internal/cmd/params_generated.go` | `FromFlags` functions: cobra flags → Params structs |

### FromFlags pattern

Generated `xParamsFromFlags(cmd)` functions use `flags.Changed()` guards so only flags the user explicitly passes populate the struct. Unset flags stay at zero value — the server applies its own defaults.

CLI-specific defaults (different from server defaults) are applied post-FromFlags:

```go
params := getAllOrdersParamsFromFlags(cmd)
if params.Status == "" {
    params.Status = "open"
}
```

`FlagOpts.Defaults` controls what `--help` shows, not what gets sent. If a default must be sent, add a post-FromFlags fallback as above.

Request body construction stays hand-written (enum casts, complex types, `Changed()` for PATCH).

## What stays hand-written

Cobra commands (`internal/cmd/`) are hand-written — they're UX decisions:
- `FlagOpts.Defaults` overrides
- `RequireStr` / `RequireAll` validation
- Request body construction (enum casts, complex types, JSON unmarshal for complex flags)
- `Changed()` checks for PATCH semantics (`order replace`, `watchlist update`)
- Output format overrides (`jsonOnly` for complex nested responses)
- Custom flags not from OAS (`--dry-run`, `--client-id`, `--market`)

## No flag exclusions

Do **not** use `FlagOpts.Exclude` to hide OAS-supported parameters. Every parameter the API supports should be exposed as a flag. The primary consumer is an agent, not a human — agents prefer explicit flags over positional args.

For complex/nested OAS fields (e.g. `advanced_instructions`, `legs`), register the flag and accept JSON input:

```go
if cmdutil.Changed(cmd, "advanced-instructions") {
    if err := json.Unmarshal([]byte(cmdutil.Str(cmd, "advanced-instructions")), &body.AdvancedInstructions); err != nil {
        return fmt.Errorf("--advanced-instructions: %w", err)
    }
}
```

Path parameters (like `order-id` in `order replace <order-id>`) remain positional — they identify the resource, not a query/body parameter.

## No backward compatibility

Pre-1.0. No aliases, shims, or deprecation wrappers. Just make the change and update docs.

## Flag naming invariant

Flags use OAS-generated kebab-case names, no aliases. `FlagOpts` enforces this at the type level. To change a flag name, change it in the spec.

## OAS specs are read-only inputs

The files in `api/specs/` are copies of upstream specs. **Never edit them in this repo.** Fix bugs at the source and re-import. The generator must work around spec bugs until upstream is corrected.

## Self-healing workarounds

When upstream OAS bugs require workarounds, record them here. After any spec update, check each entry and remove landed fixes.

| Workaround | File | What to check | Remove when |
|---|---|---|---|

## User-Agent invariant

Every outbound HTTP request must set `User-Agent: alpaca-cli/<version>`. When creating a standalone `http.NewRequest` anywhere in the codebase, set this header. Use `"alpaca-cli/" + version` in `internal/cmd/` or `oauth.UserAgent` in `internal/oauth/`.

## Update notifications

Background check (once per 24h), cached in `~/.config/alpaca/update-state.json`. A stderr notice appears when an update is available.

- `alpaca update --check` — live check, structured JSON output
- `alpaca version` — cached update info (no network call)
- `ALPACA_NO_UPDATE_NOTIFY=1` or `--quiet` — suppresses the notice

Install method is auto-detected (Homebrew, go install, binary). For Homebrew/go install users, `alpaca update` redirects to the appropriate package manager.

## After every change

```
make check     # lint + test + build
```

Fix any failures you introduce before moving on.

## Keep docs in sync

When a change affects CLI behavior, update any stale docs:

- `README.md` — user-facing documentation
- `skills/alpaca-cli/SKILL.md` — agent-facing skill ([Agent Skills](https://agentskills.io) format)

Code is the source of truth. If a doc contradicts the code, update the doc.
