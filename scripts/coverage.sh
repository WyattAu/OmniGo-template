#!/usr/bin/env bash
# Coverage gate — >=80 default (Go-idiomatic tier a). Override: ./scripts/coverage.sh 70
set -euo pipefail
cd "$(dirname "$0")/.."
go test -coverprofile=coverage.out ./...
total=$(go tool cover -func=coverage.out | awk '/^total:/ {print substr($3, 1, length($3)-1)}')
echo "coverage: ${total}%"
awk -v t="$total" -v m="${1:-80}" 'BEGIN { exit (t+0 >= m+0) ? 0 : 1 }' \
  || { echo "FAIL: coverage ${total}% below gate ${1:-80}%"; exit 1; }
