VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"
BINARY := alpaca

.PHONY: build install test test-integration lint check clean generate spec-update release

build:
	go build $(LDFLAGS) -o bin/$(BINARY) ./cmd/alpaca

install:
	go install $(LDFLAGS) ./cmd/alpaca

test:
	go test -race ./...

check: lint test build

test-integration:
	go test -v -tags integration -count=1 -timeout 5m ./test/integration/...

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/

generate:
	go run ./cmd/generate

spec-update:
	@echo "Fetching latest OpenAPI specs..."
	curl -sSfL "https://docs.alpaca.markets/openapi/trading-api.json" | python3 -m json.tool > api/specs/trading-api.json
	curl -sSfL "https://docs.alpaca.markets/openapi/market-data-api.json" | python3 -m json.tool > api/specs/market-data-api.json
	@echo "Specs updated. Run 'make generate' to regenerate client code."

release:
	@if [ -n "$$(git status --porcelain)" ]; then echo "error: working tree is dirty" >&2; exit 1; fi
	@git fetch --tags origin
	@LAST=$$(git tag -l 'v0.0.*' --sort=-v:refname | head -1); \
	if [ -z "$$LAST" ]; then NEXT=v0.0.1; \
	else NEXT=v0.0.$$((  $${LAST##*.} + 1  )); fi; \
	echo "$$LAST -> $$NEXT"; \
	read -p "Tag $$NEXT and push? [y/N] " confirm; \
	if [ "$$confirm" != "y" ] && [ "$$confirm" != "Y" ]; then echo "Aborted."; exit 1; fi; \
	git tag -a "$$NEXT" -m "Release $$NEXT" && git push origin "$$NEXT"; \
	echo "Tagged $$NEXT — release workflow will run at https://github.com/alpacahq/cli/actions"
