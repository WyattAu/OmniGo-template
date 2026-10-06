package geometry

import "testing"

// Benchmarks are the perf gate input: `make bench` runs them through
// scripts/bench-budget.sh and compares them with the committed baseline in
// bench/baseline.tsv (mean of five runs, gated against the sample spread).
//
// Batched loops, not single calls: a nanosecond-scale call is dominated by loop
// and timer overhead, so the numbers would not be comparable.
var benchSink float64

func BenchmarkDistance(b *testing.B) {
	p1, p2 := Point{0, 0}, Point{3, 4}
	b.ReportAllocs()
	for b.Loop() {
		benchSink = Distance(p1, p2)
	}
}

func BenchmarkDistanceBatch(b *testing.B) {
	p1, p2 := Point{0, 0}, Point{3, 4}
	b.ReportAllocs()
	total := 0.0
	for b.Loop() {
		for i := 0; i < 1_000; i++ {
			total += Distance(p1, p2)
		}
	}
	benchSink = total
}
