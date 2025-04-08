PHONY: build
build:
	goreleaser build --snapshot --clean  --single-target

.PHONY: run
run:
	go run cmd/buildkite-mcp-server/main.go stdio
