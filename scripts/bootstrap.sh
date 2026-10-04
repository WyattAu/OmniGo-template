#!/usr/bin/env bash
# Non-nix fallback instructions. Canonical env: flake.nix (go + lint stack).
set -euo pipefail
cat <<'MSG'
Manual toolchain (no nix):
  1. Go >= 1.24 (https://go.dev/dl/)
  2. go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
     go install golang.org/x/vuln/cmd/govulncheck@latest
     go install mvdan.cc/gofumpt@latest
  3. make ci
Prefer zero setup? Open the repo in a devcontainer, or `nix develop`.
MSG
