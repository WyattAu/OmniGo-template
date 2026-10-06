#!/usr/bin/env bash
# Perf budget gate: measure, emit bench/current.tsv, compare to the committed
# baseline with the shared comparator. `make bench-update` re-baselines
# deliberately - it is the only way a baseline moves.
set -euo pipefail
cd "$(dirname "$0")/.."
BASELINE=bench/baseline.tsv
CURRENT=bench/current.tsv
THRESHOLD_PCT="${OMNI_BENCH_THRESHOLD_PCT:-10}"
UPDATE=()
[ "${1:-}" = "--update" ] && UPDATE=(--update)
mkdir -p bench

# `-run=^$` skips test bodies; -count=5 gives a sample spread so the gate can
# use the statistical path (mean + std-dev) instead of a coin-flip threshold.
go test -run='^$' -bench=. -benchtime=200ms -count=5 ./... >bench/go-bench.txt

python3 - <<'PY' >"$CURRENT"
import math
import pathlib
import re
import statistics
import sys

pattern = re.compile(r"^(Benchmark\S+)\s+\d+\s+([\d.]+)\s+ns/op")
samples: dict[str, list[float]] = {}
for line in pathlib.Path("bench/go-bench.txt").read_text().splitlines():
    match = pattern.match(line.strip())
    if match:
        samples.setdefault(match.group(1), []).append(float(match.group(2)))
if not samples:
    raise SystemExit("bench: no `ns/op` measurements parsed from bench/go-bench.txt")

for name, values in sorted(samples.items()):
    mean = statistics.fmean(values)
    # One sample has no spread to report; 0 noise means "compare by budget only".
    noise = statistics.stdev(values) if len(values) > 1 else 0.0
    if not math.isfinite(mean) or mean <= 0:
        print(f"bench: {name} produced no usable measurement", file=sys.stderr)
        raise SystemExit(1)
    print(f"{name}\t{mean:.3f}\tns\tgate\t{noise:.3f}")
PY

python3 scripts/compare-bench.py "$BASELINE" "$CURRENT" \
  --threshold-pct "$THRESHOLD_PCT" "${UPDATE[@]+"${UPDATE[@]}"}"
