## Summary

<!-- What changed? -->

## Motivation

<!-- Why is this change needed? Link the issue or prior discussion when applicable. -->

## Verification

<!-- List the checks you ran and their results. -->

## Checklist

- [ ] The change is focused and follows the contribution policy.
- [ ] Any substantial API, dependency, or third-party work was discussed with maintainers before implementation.
- [ ] Generated files and the OpenAPI specs under `api/specs/` were not edited by hand; `make generate` was run when specs or the generator changed.
- [ ] New or changed behavior has focused tests when applicable.
- [ ] Default tests are deterministic and credential-free and cannot place live orders; integration tests target paper trading, use unique symbols, and clean up after themselves; and fixtures use synthetic data.
- [ ] `README.md` and `.agents/skills/alpaca-cli/SKILL.md` are updated when CLI behavior changes.
- [ ] `make check` passes, and `make test-integration` was run when CLI behavior changed.
- [ ] No credentials, account data, request or trace IDs, or other secrets are included.
- [ ] Any provider-specific integration or promotional contribution has written Alpaca approval confirmed by repository maintainers.
