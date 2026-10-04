# 0001 — Single module anchored by go.work

Date: 2026-10-04

## Status

Accepted

## Context

Go monorepos start as one module and drift into accidental multi-module
setups with replace directives and version skew.

## Decision

One root module (`cmd/` + `internal/` + nothing exported at root);
`go.work` is committed as the *anchor* and documents the split policy:
a module breaks out into `modules/<name>` only when it needs an
independent release cadence, and `internal/` stays the privacy boundary.

## Consequences

- Zero replace-directive churn early; goreleaser config stays single-build.
- Splitting later is a documented, mechanical move.
