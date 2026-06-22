---
name: alpaca-cli-regenerate
description: >-
  Keep the CLI in sync with upstream OpenAPI specs. Pulls latest specs,
  runs the code generator, registers new operations, updates golden files,
  and runs all checks. Use when specs changed upstream, a new API endpoint
  was added, generation is stale, or the user says "regenerate",
  "generate", "sync specs", "update specs", or "spec-update".
---

# Generate - Keep the CLI Up to Date

## Pipeline

Run each phase in order. Stop and fix before moving to the next.

### Phase 1: Pull latest specs

`make spec-update` requires network access (it curls `docs.alpaca.markets`).
Request `full_network` permissions if running in a sandbox.

```bash
make spec-update
```

Downloads `trading-api.json` and `market-data-api.json` from
`docs.alpaca.markets/openapi/` into `api/specs/`.

After pulling, check what changed:

```bash
git diff --stat api/specs/
```

If nothing changed, specs are already current - skip to Phase 4 to
verify generated code is still fresh.

### Phase 2: Generate code

```bash
make generate
```

This runs `go run ./cmd/generate` which reads the specs and writes:

| Generated file | Contents |
|---|---|
| `internal/api/trading_types.gen.go` | Structs and enums from trading spec |
| `internal/api/trading_client.gen.go` | Trading API client methods |
| `internal/api/marketdata_types.gen.go` | Structs and enums from market data spec |
| `internal/api/marketdata_client.gen.go` | Market data API client methods |
| `internal/api/descriptions.gen.go` | Op metadata, flags, response schemas |
| `internal/cmd/commands.gen.go` | Cobra command tree |

**Never edit `*.gen.go` files directly.** Change the specs or generator,
then re-run `make generate`.

### Phase 3: Fix generator failures

The generator enforces exhaustiveness. Common failures and fixes:

#### New operation not registered

```
unregistered operation "FooBar" — add to cmdRegistry or cmdSkip
```

Every `operationId` in the specs must appear in either `cmdRegistry` or
`cmdSkip` in `cmd/generate/commands.go`.

To register a new command, add an entry to `cmdRegistry`:

```go
"FooBar": {
    parent:   "<parent-group>",   // key in cmdParents
    use:      "<subcommand-name>",
    examples: "  alpaca <parent> <subcommand> --flag value",
},
```

Rules for registry entries:
- `parent` must be a key in `cmdParents` (or create a new parent group)
- `examples` is **required** - the generator fails without it
- Use kebab-case for `use` names
- Add `defaults` for sensible default flag values
- Add `bodyAliases` if a body field name collides with a query/path param
- Add `bodyHook` / `bodySkipFields` for complex body construction
- Add `configureFunc` for hand-written configure hooks (e.g. order submit's bracket legs)
- Add `normalize` for path params that need URL normalization (e.g. symbols with `/`)

If the command IS its parent (e.g. `alpaca clock` runs directly, not
`alpaca clock get`), use `self: true` and omit `use`:

```go
"LegacyClock": {
    parent:   "clock",
    self:     true,
    examples: "  alpaca clock",
},
```

To skip an operation (intentionally exclude from the CLI):

```go
var cmdSkip = map[string]string{
    "FooBar": "reason for skipping",
}
```

#### Body field collision

```
body field "name" collides with a query/path param — add bodyAliases
```

Fix by adding a `bodyAliases` entry that renames the body flag:

```go
"UpdateFoo": {
    parent:      "foo",
    use:         "update",
    bodyAliases: map[string]string{"name": "new-name"},
    examples:    "  alpaca foo update --name old --new-name new",
},
```

#### New parent group needed

Add to `cmdParents` in `cmd/generate/commands.go`:

```go
"myGroup": {
    use:   "my-group",
    short: "Short description",
    long:  "Longer description for --help.",
},
```

For nested groups, set `parent`:

```go
"mySubGroup": {
    use:    "sub",
    short:  "Sub-group description",
    parent: "myGroup",
},
```

**Top-level parent groups also need root wiring.** If the new group is
not nested under an existing parent, add it to `addGroup()` in
`internal/cmd/root.go`. Pick the right group (`tradingGroup`,
`accountGroup`, or `utilGroup`):

```go
addGroup(rootCmd, tradingGroup.ID, orderCmd, ..., myGroupCmd)
```

Without this, the command won't appear in `--help` or `--help-all`.

After fixing, re-run `make generate` until it succeeds.

### Phase 4: Update golden files

If the command tree or op metadata changed, golden files will be stale:

```bash
go test ./internal/cmd -run TestCommandTreeGolden -update
go test ./internal/cmd -run TestOpsGolden -update
```

### Phase 5: Run all checks

```bash
make check
```

This runs `golangci-lint`, `go test -race ./...`, and `go build`. Fix any
failures before proceeding.

If tests fail due to golden drift you missed, re-run Phase 4.

### Phase 6: Verify completeness

Review what changed end-to-end:

```bash
git diff --stat
```

Verify:
- [ ] `api/specs/` - spec files updated
- [ ] `internal/api/*.gen.go` - types and clients regenerated
- [ ] `internal/cmd/commands.gen.go` - command tree regenerated
- [ ] `internal/cmd/testdata/*.golden` - golden files updated if needed
- [ ] `cmd/generate/commands.go` - new operations registered if needed
- [ ] `internal/cmd/root.go` - new top-level parent groups wired via `addGroup`
- [ ] `test/integration/` - integration tests added for new commands

If commands were added, add integration tests in `test/integration/`.
Follow the rules in the "Integration tests" section of `AGENTS.md` -
one file per feature area, `t.Parallel()` for read-only tests, cleanup
for writes. If the endpoint is unavailable on paper (see the
paper-unavailable list in `AGENTS.md`), write a test that accepts either
a valid response or a structured JSON error.

If commands or flags were added, removed, or renamed, follow the
"Keep docs in sync" section in `AGENTS.md`.

Keep `README.md` concise. It should describe command areas and point to
`alpaca --help-all`, `<command> --help`, and `<command> --schema` instead
of listing generated command tables. Add README detail only when user-facing
behavior, safety guidance, authentication, output, or install/update flow
changes.

## Quick reference

| Task | Command |
|---|---|
| Pull specs | `make spec-update` |
| Generate code | `make generate` |
| Update golden files | `go test ./internal/cmd -run TestCommandTreeGolden -update && go test ./internal/cmd -run TestOpsGolden -update` |
| Lint + test + build | `make check` |
| See what changed | `git diff --stat` |
| Integration tests | `make test-integration` (needs API keys) |

## Key files

| File | Role | Editable? |
|---|---|---|
| `api/specs/*.json` | OAS specs (read-only inputs) | No - pull from upstream |
| `cmd/generate/main.go` | Generator logic | Yes |
| `cmd/generate/commands.go` | Command registry and parent groups | Yes |
| `internal/api/*.gen.go` | Generated types and clients | No - regenerate |
| `internal/cmd/commands.gen.go` | Generated Cobra commands | No - regenerate |
| `internal/cmd/root.go` | Root command wiring (`addGroup` for top-level parents) | Yes |
| `internal/cmd/testdata/*.golden` | Golden test snapshots | Update via `-update` flag |

## Anti-Patterns

- **NEVER** edit `*.gen.go` files directly - they are overwritten by `make generate`. Change the specs or the generator instead.
- **NEVER** edit spec files in `api/specs/` - fix bugs upstream and re-import with `make spec-update`.
- **NEVER** skip `make check` after generating - generated code can introduce lint failures or test regressions.
- **NEVER** update golden files without reviewing the diff - golden updates should reflect intentional command tree changes, not mask bugs.
- **NEVER** add a `cmdRegistry` entry without `examples` - the generator enforces this and will fail.
- **NEVER** add an operation to `cmdSkip` without a reason - the reason documents why the endpoint is excluded from the CLI.
