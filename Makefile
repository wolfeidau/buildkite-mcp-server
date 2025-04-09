default: build

PHONY: build
build:
	go build -o buildkite-mcp-server ./cmd/buildkite-mcp-server/main.go

build-snapshot:
	goreleaser build --snapshot --clean  --single-target

.PHONY: run
run:
	go run cmd/buildkite-mcp-server/main.go stdio

test:
	go test ./... -v
