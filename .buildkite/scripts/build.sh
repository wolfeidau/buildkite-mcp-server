#!/usr/bin/env bash

set -euo pipefail

echo "--- clean"
rm -rf dist
mkdir dist
echo "--- :golang: mod download"
go mod download
echo "--- :golang: build"
go build -v -o dist ./cmd/...
echo "--- :golang: build test"
go test -v -c -o dist/test/ ./...
