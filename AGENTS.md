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

## Design Notes

- **`FlagDef.OASName` must stay**: Flag names are kebab-case (`page-token`), OAS names are snake_case (`page_token`). The mapping `_ → -` is lossy — if an upstream OAS param ever uses a hyphen, runtime reversal (`- → _`) would silently send the wrong query key. Keep both fields.

## Keep docs in sync

When a change affects CLI behavior, update any stale docs:

- `README.md` — user-facing documentation
- `skills/alpaca-cli/SKILL.md` — agent-facing skill

Code is the source of truth. If a doc contradicts the code, update the doc.
