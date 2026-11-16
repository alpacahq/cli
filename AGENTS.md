# Alpaca CLI

A single-binary CLI for the Alpaca Trading API. Think `gh` for GitHub, `stripe` for Stripe — but for trading.

[`CONTRIBUTING.md`](CONTRIBUTING.md) is the canonical engineering guide. Read it before changing code. It covers generated versus handwritten code, command design, the output contract, integration test rules, financial safety, and which docs to keep in sync.

## Rules you must not miss

- **Generate everything**: do not edit generated files (`*.gen.go`) or the OpenAPI specs in `api/specs/`. Change the generator, or fix the spec upstream and re-import, then `make generate`.
- **Agent-first**: all parameters are explicit `--flag value`, with no positional arguments (except `alpaca api [METHOD] <path>`).
- **No backward compatibility**: pre-1.0. No aliases, shims, or deprecation wrappers.
- **Output contract**: API commands emit JSON. Operational commands emit text, and their signal is the exit code. Details are in `CONTRIBUTING.md`.

## After every change

```
make check             # lint + test + build
make test-integration  # live paper API; `make check` does not run these
```

Fix any failures you introduce. When a change renames commands or flags, or changes output shape, update `test/integration/` and any stale docs (`README.md`, `.agents/skills/alpaca-cli/SKILL.md`). Code is the source of truth.
