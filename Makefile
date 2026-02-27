VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"
BINARY := alpaca

.PHONY: build install test test-integration lint clean

build:
	go build $(LDFLAGS) -o bin/$(BINARY) ./cmd/alpaca

install:
	go install $(LDFLAGS) ./cmd/alpaca

test:
	go test ./...

test-integration:
	go test -v -tags integration -count=1 -timeout 5m ./test/integration/...

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/
