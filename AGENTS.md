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

Query param structs are populated via generated `xParamsFromFlags(cmd)` functions. Each uses `flags.Lookup()` + `flags.Changed()` guards: excluded flags are safely skipped, and unset flags (even those with OAS defaults) are left at their zero value. Only flags the user explicitly provides on the command line populate the struct. This ensures OAS defaults are never sent as explicit query parameters — the server applies its own defaults.

Commands that need a CLI-specific default (different from the server default) apply it post-FromFlags:

```go
params := getAllOrdersParamsFromFlags(cmd)
if params.Status == "" {
    params.Status = "open"
}
```

`FlagOpts.Defaults` controls what `--help` shows, not what gets sent to the API. If a `FlagOpts.Defaults` value must be sent, add a post-FromFlags fallback as above.

Request body construction stays hand-written (enum casts, complex types, `Changed()` for PATCH).

## What stays hand-written

Cobra commands (`internal/cmd/`) are hand-written because they're UX decisions: which flags to expose, how to group subcommands, what examples to show, which columns to display, and how to format output.

Hand-written concerns per command:
- `FlagOpts.Exclude` decisions (positional args, deprecated params)
- `FlagOpts.Defaults` overrides
- `RequireStr` / `RequireAll` validation
- Request body construction (enum casts, complex types, positional arg overrides)
- `Changed()` checks for PATCH semantics (`order replace`, `watchlist update`)
- Output column selection and rendering
- Custom flags not from OAS (`--dry-run`, `--client-id`, `--market`)

## Flag naming invariant

Flags always use OAS-generated kebab-case names. No aliases. `FlagOpts` does not support aliases — this is enforced at the type level. If a flag name needs to change, change it in the spec.

## After every change

```
make check     # lint + test + build
```

Fix any failures you introduce before moving on.

## Keep docs in sync

When a change affects CLI behavior — new commands, changed flags, modified output formats, auth flow changes, or environment variable additions — check these files and update any that are now stale:

- `README.md` — canonical user-facing documentation
- `skills/alpaca-cli/SKILL.md` — agent-facing install, auth, and usage skill ([Agent Skills](https://agentskills.io) format)

Code is always the source of truth. If a doc contradicts the code, update the doc.
