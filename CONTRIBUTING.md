# Contributing

Thanks for helping improve Alpaca CLI. By participating, you agree to follow our [Code of Conduct](./CODE_OF_CONDUCT.md).

## Choose the right channel

- **CLI issues:** Bugs, feature requests, or questions specific to this CLI → [GitHub Issues](https://github.com/alpacahq/cli/issues/new/choose).
- **Security vulnerabilities:** Do not report them publicly. Follow the private reporting process in our [Security Policy](./SECURITY.md).
- **Alpaca API bugs:** Bugs in the Alpaca API itself, its OpenAPI specs, or its documentation → [alpacahq/alpaca-api](https://github.com/alpacahq/alpaca-api/issues/new/choose).
- **Community:** API feature requests, general questions, and dashboard bugs → [Alpaca Community Forum](https://forum.alpaca.markets/) or [Slack](https://alpaca.markets/slack).
- **Account help:** Questions about your account, funding, or orders → [Support request](https://alpaca.markets/support/request).
- **Outages:** Service status and incident updates → [Status page](https://status.alpaca.markets/).

Please do not include credentials, account details, request or trace IDs, or other sensitive information in an issue, discussion, pull request, test, or example.

## Contribution policy

### Contributions we welcome

Focused contributions are welcome when they improve this CLI, including:

- bug fixes and compatibility improvements;
- support for documented Alpaca API capabilities;
- focused tests and reliability improvements;
- documentation and examples; and
- maintainability improvements that preserve public behavior.

### Discuss substantial changes first

Open an issue before investing in a substantial CLI design, new dependency, broad refactor, or third-party integration. Early discussion helps confirm that the proposal fits the CLI's scope and avoids duplicated or unusable work.

Bugs in the Alpaca API itself, its OpenAPI specs, or its documentation must be fixed upstream, not patched in this CLI. Report them in [alpacahq/alpaca-api](https://github.com/alpacahq/alpaca-api/issues/new/choose). API feature requests, general questions, and dashboard bugs belong on the [Alpaca Community Forum](https://forum.alpaca.markets/) or [Slack](https://alpaca.markets/slack). Contributions that add support for already documented API capabilities are welcome after the CLI design is agreed with maintainers.

### Out of scope

This repository does not accept:

- trading strategies, algorithm design, or investment advice;
- bugs in the Alpaca API, its OpenAPI specs, or docs.alpaca.markets;
- API feature requests, dashboard bugs, or outage reports;
- questions that aren't about this CLI, including questions about your account, funding, or orders;
- credentials, private account data, request or trace IDs, or other secrets;
- referral links or promotional content; or
- unapproved integrations with third-party products or services.

Contributions that add or promote third-party products or services, including provider-specific clients, API calls, hosted endpoints, referral links, or promotional content, require prior written approval from Alpaca before implementation. Open an issue describing the technical scope; repository maintainers will coordinate the necessary internal review and confirm whether the proposal may proceed. Contributors do not need to identify or contact internal approvers.

Using an ordinary open-source dependency or improving generic interoperability does not by itself make a contribution a provider-specific integration. Maintainers will explain when a proposal falls outside this policy. A pull request that proceeds without the required approval may be closed as out of scope, so open an issue first if you are unsure whether a proposal qualifies.

## License and contribution terms

The project is licensed under the [Apache License 2.0](./LICENSE). Subject to its full terms, the license permits use, reproduction, modification, and distribution of the software. It does not grant permission to use Alpaca trade names, trademarks, service marks, or product names except as described in the license.

Unless you explicitly state otherwise, a contribution intentionally submitted for inclusion in this project is provided under Apache-2.0 without additional terms, as described in Section 5 of the license. This project does not currently require a separate Contributor License Agreement (CLA) or Developer Certificate of Origin (DCO) sign-off.

## Contribution workflow

1. Search existing issues and pull requests.
2. Open an issue first for work described under [Discuss substantial changes first](#discuss-substantial-changes-first).
3. Fork the repository and create a focused branch.
4. Make the change in hand-written source files. For generated code, change the generator or re-import upstream specs with `make spec-update`, then run `make generate`.
5. Add focused tests for new or changed behavior and update relevant documentation.
6. Run the required checks listed below.
7. Open a focused pull request, explain why the change is needed, and respond to review feedback.

Use clear, descriptive commit messages and pull request titles. Conventional Commits are not required, and contributors do not need to squash commits before opening a pull request. Maintainers use GitHub's squash-merge workflow for the final merge.

## Prerequisites

- Go, using the version declared in [`go.mod`](./go.mod).
- `make` for the repository's documented convenience targets.
- [`golangci-lint`](https://golangci-lint.run/) v2.11.4 for `make lint`, matching the version pinned in CI.
- Alpaca paper trading credentials when running the integration tests.
- `curl` and `python3` when running `make spec-update`.
- Optionally, [`oasdiff`](https://github.com/oasdiff/oasdiff) for `make spec-check`.

## Setup and checks

Run the required checks from the repository root after every change:

```bash
make check
make test-integration
```

`make check` runs the linter, the unit tests with the race detector, and builds the binary to `bin/alpaca`. `make test-integration` verifies live paper API behavior and requires paper trading credentials.

Regenerate generated API clients and command metadata when specs or the generator change:

```bash
make generate
```

Integration tests are gated by the `integration` build tag and always target the paper trading API. They require paper trading credentials, supplied as either an API key and secret or an OAuth access token:

```bash
export ALPACA_TEST_API_KEY=PK...
export ALPACA_TEST_SECRET_KEY=...
# or: export ALPACA_TEST_ACCESS_TOKEN=...
make test-integration
```

CI also verifies that generated code is up to date by running `make generate` and failing on any diff in `internal/api/` and `internal/cmd/commands.gen.go`.

There is no numeric coverage threshold, but changed behavior must have focused tests.

## Testing and financial safety

- Default repository tests and CI must never place live orders. Unit tests (`make test`) must be deterministic and must not depend on Alpaca credentials or network access to the Alpaca API.
- Integration tests always target the paper trading API, never live trading. Tests that write data must clean up with `t.Cleanup()` (cancel orders, delete watchlists, close positions) and use a symbol no other parallel test uses.
- Paper trading is the CLI's default. Examples in documentation must use paper trading by default and require an explicit opt-in (`profile login --live` or `ALPACA_LIVE_TRADE=true`) before targeting live trading.
- Real credentials, OAuth tokens, account identifiers, account data, and production payloads must never appear in tests, fixtures, examples, or committed logs and artifacts. Remove request and trace IDs from public or committed material; share private diagnostics only through Alpaca Support channels when needed.

## Coding guidelines

### Generated versus handwritten code

The CLI is driven by the OpenAPI specs in `api/specs/*.json`. Maximize what is generated and minimize what is hand-written.

Never hand-edit generated files, including `internal/api/*.gen.go` and `internal/cmd/commands.gen.go`. Other files in `internal/api/` (such as `api.go` and `types.go`) are handwritten. Change the generator in `cmd/generate`, or fix specification bugs upstream in [alpacahq/alpaca-api](https://github.com/alpacahq/alpaca-api/issues/new/choose) and re-import with `make spec-update`; then run `make generate`. Command help text for generated commands comes from the OpenAPI descriptions, so fix wording problems there upstream. Do not edit the specs under `api/specs/` by hand.

Handwritten code lives in `cmd/alpaca`, the generator in `cmd/generate`, and the packages under `internal/` (for example `internal/cmd`, `internal/client`, `internal/config`, `internal/oauth`, and `internal/output`).

### Command design

The primary consumer of the CLI is an AI agent. All parameters are explicit `--flag value` pairs, with no positional arguments. The one exception is `alpaca api [METHOD] <path>`, which is a raw escape hatch rather than a generated command.

Flag names are kebab-case (`page-token`) while OpenAPI names are snake_case (`page_token`). Keep both `FlagDef.Name` and `FlagDef.OASName`: the `_` to `-` mapping is lossy, so if an upstream parameter ever contains a hyphen, reversing it at runtime would silently send the wrong query key.

There is no backward-compatibility layer before 1.0: make the change directly without aliases, shims, or deprecation wrappers. When a change renames commands or flags, or changes output shape, update `test/integration/` so it stays in sync.

### Output contract

- **API commands** call Alpaca endpoints and return structured JSON on stdout. They support `--csv`, `--jq`, `--quiet`, and `--schema`. Errors go to stderr as JSON.
- **Operational commands** (`version`, `doctor`, `profile *`, `update`, `completion`, `--help`, `--help-all`, and `--schema`) manage the CLI itself and emit human-readable text. The machine-readable signal is the exit code (0 for success, non-zero for failure). Do not convert them to JSON: an agent that needs to verify connectivity runs `alpaca account get --quiet`, not `alpaca doctor`. The one exception is `update --check`, which emits JSON so agents can decide whether to upgrade.

If a command hits the Alpaca API and returns API data, it emits JSON. If it manages the CLI's own state or helps a human troubleshoot, it emits text.

### Integration tests

- Read-only tests must call `t.Parallel()`.
- Use flat tests for independent calls and sub-tests only for sequential chains that depend on prior state.
- Use the existing helpers (`alpaca()`, `alpacaFail()`, `parseJSONMap()`, `requireFields()`, `daysAgo()`, `pollFor()`) instead of reinventing them.
- Give each parallel test a unique symbol with `submitTestOrder(t, "SYMBOL")`. The Alpaca API rejects orders as a potential wash trade when buy and sell orders coexist on the same symbol, and bracket orders create sell-side child legs.
- Some endpoints are unavailable on paper trading (wallet, forex, logo, and tokenization). Don't write tests that require them without a compatible test account.
- Keep one file per feature area, such as `order_test.go` or `data_stock_test.go`. Cross-cutting end-to-end flows go in `e2e_test.go`.

### Documentation

Code is the source of truth. When a change affects CLI behavior, update any stale documentation in [`README.md`](./README.md) and in the agent-facing skill at [`.agents/skills/alpaca-cli/SKILL.md`](./.agents/skills/alpaca-cli/SKILL.md).

## Pull request checklist

Before requesting review, confirm that:

- the change is focused and follows the contribution policy;
- any substantial API, dependency, or third-party work was discussed with maintainers before implementation;
- generated files and the OpenAPI specs were not edited by hand, and `make generate` was run when specs or the generator changed;
- new or changed behavior has focused tests when applicable;
- default tests are deterministic and credential-free and cannot place live orders; integration tests target paper trading, use unique symbols, and clean up after themselves; and fixtures use synthetic data;
- `README.md` and `.agents/skills/alpaca-cli/SKILL.md` are updated when CLI behavior changes;
- `make check` and `make test-integration` pass;
- no credentials, account data, request or trace IDs, or other secrets are included; and
- any provider-specific integration or promotional contribution has written Alpaca approval confirmed by repository maintainers.
