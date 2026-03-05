VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"
BINARY := alpaca

.PHONY: build install test test-integration lint check clean generate spec-update man

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

man:
	go run ./cmd/doc

spec-update:
	@echo "Fetching latest OpenAPI specs..."
	curl -sSfL "https://docs.alpaca.markets/_mock/openapi/trading/bundled" | python3 -m json.tool > api/specs/trading-api.json
	curl -sSfL "https://docs.alpaca.markets/_mock/openapi/data/bundled" | python3 -m json.tool > api/specs/market-data-api.json
	@echo "Specs updated. Run 'make generate' to regenerate client code."
