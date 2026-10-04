# 0002 — goreleaser with homebrew tap

Date: 2026-10-04

## Status

Accepted

## Context

Estate CLIs already distribute via homebrew taps (homebrew-clawdius,
homebrew-suture-merge-driver); release chores (checksums, SBOM, attestation)
must not be manual.

## Decision

Tag `v*` triggers goreleaser: CGO-free multiplatform builds (linux/darwin/
windows × amd64/arm64), checksums, SBOM, build-provenance attestation, and a
commented homebrew tap scaffold to enable after the first release.

## Consequences

- Releases are one tag push; tap updates are one config uncomment.
