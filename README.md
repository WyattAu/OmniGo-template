# OmniGo-template

Maximalist Go **monorepo** template: `cmd/` + `internal/` single module
anchored by `go.work`, golangci-lint (max config), govulncheck, native
fuzzing, race-detector CI, goreleaser with homebrew tap scaffold — full
IDE/OS integration (nix flake, dual devcontainers, VS Code). Part of the
[WyattAu Omni template family](https://github.com/WyattAu?tab=repositories&q=omni-).

## Start here (after "Use this template")

1. Rename: module path `github.com/wyattau/omnigo-template` → your module,
   `cmd/omni` → your binary, `internal/geometry` → your domain.
2. Pick a door — all resolve to identical toolchains:

   | Door | Command |
   |---|---|
   | nix + direnv (host) | `direnv allow` |
   | Devcontainer (image) | VS Code → *Reopen in Container* |
   | Devcontainer (nix)  | palette → *Rebuild in Container* → pick `.devcontainer/nix/` |

   No nix, no docker? `./scripts/bootstrap.sh` prints the manual path.
3. `make ci` — must be green before your first push.

## Make targets

| Target | Gate |
|---|---|
| `make build` | `go build ./...` |
| `make test` | `go test -race ./...` |
| `make lint` | golangci-lint (gosec, gocritic, revive, exhaustive, …) |
| `make fmt` / `fmt-check` | gofumpt |
| `make coverage` | `go test -cover` with ≥80% gate |
| `make vuln` | govulncheck (symbol-level) |
| `make fuzz` | native fuzzing (30s default) |
| `make contract` | Omni Core Contract structural checks |
| `make ci` | contract + fmt-check + lint + test + vuln |

## What is inside

```
cmd/omni               example binary (flags + error handling, no panics)
internal/geometry      example domain package: table tests + Go fuzz target
go.work                monorepo anchor (ADR-0001: split policy)
.golangci.yml          maximal-but-explainable lint stack (v2 config)
.goreleaser.yaml       tag → multiplatform release + SBOM + tap scaffold
scripts/  + Makefile   the gates (make ci == CI)
docs/adr/              decision log
.github/workflows/     ci (matrix + tip leg + race + vuln + fuzz), release, devcontainers
.forgejo/              thin self-hosted mirror (scripts are canonical)
```

## Release flow

Push tag `v*` → goreleaser builds, SBOMs, checksums, and attests; enable the
homebrew tap block after creating `homebrew-omni` (ADR-0002).

## Estate pointers

- Gates, policies: [engineering-standards](https://github.com/WyattAu/engineering-standards)
- Omni Core Contract: [OMNI-CORE.md](https://github.com/WyattAu/engineering-standards/blob/main/OMNI-CORE.md)

## License

Apache-2.0 — commercial use expressly permitted.


## Performance budgets

Performance is a gate, not a hope. `make bench` measures, writes
`bench/current.tsv`, and compares it against the committed
`bench/baseline.tsv`; anything more than the threshold worse fails. The
comparator (`scripts/compare-bench.py`) is identical across the whole Omni
estate, so the policy is auditable in one place.

| Verb | What it does |
|---|---|
| `make bench` | measure + compare (advisory job in CI: `perf`) |
| `make bench-update` | deliberately re-baseline; the only way a baseline moves |

The first run on a fresh clone records the baseline instead of failing, so the
gate is meaningful from the second run onwards. Override the budget per run
with `OMNI_BENCH_THRESHOLD_PCT=15 make bench`. Rationale and per-template
metrics: `docs/adr/0005-performance-budget-gate.md`.


## Determinism

`make repro` builds twice from a clean state with a pinned `SOURCE_DATE_EPOCH`
and compares artifact hashes. Toolchains that are deterministic gate the build;
toolchains that embed timestamps or build ids by design report the difference
and explain why, rather than pretending to be reproducible. Rationale and the
per-toolchain split: `docs/adr/0006-determinism-verification.md`.
