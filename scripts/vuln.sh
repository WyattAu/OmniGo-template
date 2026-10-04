#!/usr/bin/env bash
# govulncheck — symbol-level vulnerability analysis of the dep graph.
set -euo pipefail
cd "$(dirname "$0")/.."
exec govulncheck ./...
