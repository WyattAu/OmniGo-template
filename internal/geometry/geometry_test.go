package geometry

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name string
		a, b Point
		want float64
	}{
		{"3-4-5", Point{0, 0}, Point{3, 4}, 5},
		{"identity", Point{1, 1}, Point{1, 1}, 0},
		{"negative quadrant", Point{-1, -1}, Point{2, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Distance(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-12 {
				t.Fatalf("Distance(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	square := []Point{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
	if got := Perimeter(square); math.Abs(got-4) > 1e-12 {
		t.Fatalf("Perimeter(unit square) = %v, want 4", got)
	}
	if got := Perimeter(nil); got != 0 {
		t.Fatalf("Perimeter(nil) = %v, want 0", got)
	}
}

// REQ-001: total function property — never negative, NaN only from NaN input.
func FuzzDistance(f *testing.F) {
	f.Add(0.0, 0.0, 3.0, 4.0)
	f.Add(-1.5, -2.5, 2.5, 3.5)
	f.Fuzz(func(t *testing.T, x1, y1, x2, y2 float64) {
		if math.IsNaN(x1) || math.IsNaN(y1) || math.IsNaN(x2) || math.IsNaN(y2) {
			t.Skip("NaN input: NaN propagation is expected")
		}
		d := Distance(Point{x1, y1}, Point{x2, y2})
		if math.IsNaN(d) || d < 0 {
			t.Fatalf("Distance produced %v for finite inputs", d)
		}
	})
}
