#!/usr/bin/env bash
# Short local fuzz. CI runs the corpus on every PR; schedule long runs via
# workflow_dispatch. Seed corpus lives in testdata/fuzz/.
set -euo pipefail
cd "$(dirname "$0")/.."
exec go test -fuzz="FuzzDistance" -fuzztime="${1:-30s}" ./internal/geometry
