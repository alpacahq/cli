# Alpaca CLI

A single-binary CLI for the Alpaca Trading API. Think `gh` for GitHub, `stripe` for Stripe — but for trading.

## Core Principle: Generate Everything

The CLI is driven by OpenAPI specs. Maximize what's generated, minimize what's hand-written.

```
api/specs/*.json  →  cmd/generate/main.go  →  internal/api/
```

**Do not edit `internal/api/` directly.** Change the specs or the generator, then `make generate`.

## What stays hand-written

Cobra commands (`internal/cmd/`) are hand-written because they're UX decisions: which flags to expose, how to group subcommands, what examples to show, which columns to display, and how to format output. Generated CLIs are always worse.

Everything else — types, clients, param structs, flag definitions, enum completions, validation, mutation metadata — is generated.

## After every change

```
make test      # go test ./...
make lint      # golangci-lint run ./...
```

Fix any failures you introduce before moving on.

## Keep docs in sync

When a change affects CLI behavior — new commands, changed flags, modified output formats, auth flow changes, or environment variable additions — check these files and update any that are now stale:

- `README.md` — canonical user-facing documentation
- `skills/alpaca-cli/SKILL.md` — agent-facing install, auth, and usage skill ([Agent Skills](https://agentskills.io) format)
- `skills/alpaca-cli/references/commands.md` — command reference with examples

Code is always the source of truth. If a doc contradicts the code, update the doc.
