# Alpaca CLI

A single-binary CLI for the Alpaca Trading API. Think `gh` for GitHub, `stripe` for Stripe — but for trading.

## Core Principle: Generate Everything Possible

The CLI is driven by OpenAPI specs. Maximize what's generated, minimize what's hand-written.

### What's generated

`api/specs/*.json` → `cmd/generate/main.go` → `internal/api/`

- **Types**: all request/response structs, enums, aliases (~170 types)
- **Typed clients**: `TradingClient` (61 methods), `MarketDataClient` (47 methods)
- **Param structs**: with `Values() url.Values` for query string encoding
- **Enum value slices**: `var <Name>Values = []string{...}` for every enum (schema-level and parameter-level) — used for shell completions via `cobra.FixedCompletions`
- **Parameter defaults**: `var <Params>Defaults = map[string]string{...}` — spec-defined default values for query parameters
- **Mutation metadata**: `var TradingMutatingMethods = map[string]bool{...}` — which client methods mutate state (POST/PUT/PATCH/DELETE). A test in `spec_test.go` verifies every command calling a mutating method has `warnLive()` or `requireConfirmation()`.
- **Request body validation**: `Validate() error` methods on request body structs with `required` fields — checks for zero-value strings

Trigger: `make generate` (or `go generate ./internal/api/...`)

Update specs from upstream: `make spec-update`

### What's NOT generated (and why)

- **Cobra commands** — flag names, help text, examples, and column definitions are UX decisions that require human judgment. Generated CLIs are always worse.
- **Flag registration** — commands manually populate generated param structs via `cmdutil` helpers. This keeps flag descriptions (UX) co-located with command definitions. The *binding* of which enum to which flag is a UX decision; the enum *values* come from the generated `<Name>Values` slices when available.
- **Output columns** — which fields to show, column headers, and formatting (P&L coloring, dollar formatting) are presentation choices, not API concerns.

### The boundary

```
Hand-written                          Generated
─────────────                         ─────────
Cobra commands (internal/cmd/)   →    Param/body structs (internal/api/)
Flag definitions + help text     →    Client methods (internal/api/)
Column definitions (columns.go)  →    Type definitions (internal/api/)
Output rendering logic           →    URL encoding (Values())
Flag-to-enum binding             →    Enum value slices (<Name>Values)
                                      Parameter defaults (<Params>Defaults)
                                      Mutation metadata (MutatingMethods)
                                      Request body Validate() methods
```

Commands bridge the two worlds: they read flags via `cmdutil`, populate generated structs, call generated client methods, and render responses through `output.Render`/`output.PrintSingle`.

**Do not edit `internal/api/` directly.** Change the specs or the generator, then regenerate.

## After Every Change

Always run tests and linting before considering a change complete:

```
make test      # go test ./...
make lint      # golangci-lint run ./...
```

Fix any failures you introduce before moving on.
