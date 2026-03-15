# Alpaca CLI

A single-binary CLI for the Alpaca Trading API. Think `gh` for GitHub, `stripe` for Stripe — but for trading.

## Core Principle: Generate Everything

The CLI is driven by OpenAPI specs. Maximize what's generated, minimize what's hand-written. **Do not edit generated files directly.** Change the specs or the generator, then `make generate`.

## Design Philosophy

- **Agent-first**: the primary consumer is an AI agent. All parameters are explicit `--flag value` — no positional arguments.
- **OAS specs are read-only inputs**: never edit the specs in this repo. Fix bugs upstream and re-import.
- **No backward compatibility**: pre-1.0. No aliases, shims, or deprecation wrappers. Just make the change.

## After every change

```
make check     # lint + test + build
```

Fix any failures you introduce before moving on.

## Keep docs in sync

When a change affects CLI behavior, update any stale docs:

- `README.md` — user-facing documentation
- `skills/alpaca-cli/SKILL.md` — agent-facing skill

Code is the source of truth. If a doc contradicts the code, update the doc.
