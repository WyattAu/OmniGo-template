#!/usr/bin/env bash
# Race-detector test run — the Go gate workhorse.
set -euo pipefail
cd "$(dirname "$0")/.."
exec go test -race ./...
