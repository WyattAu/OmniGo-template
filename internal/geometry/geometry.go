// Package geometry demonstrates the internal-package pattern: consumed by
// cmd/ via the module path, invisible to importers outside the module.
package geometry

import "math"

// Point is a 2-D coordinate in float space.
type Point struct {
	X, Y float64
}

// Distance returns the Euclidean distance between a and b. It is total:
// for any finite inputs the result is finite and non-negative (REQ-001).
func Distance(a, b Point) float64 {
	dx := b.X - a.X
	dy := b.Y - a.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Perimeter returns the perimeter of the polygon formed by pts (closing
// back to the first point). Fewer than two points has zero perimeter.
func Perimeter(pts []Point) float64 {
	if len(pts) < 2 {
		return 0
	}
	total := 0.0
	for i := range pts {
		total += Distance(pts[i], pts[(i+1)%len(pts)])
	}
	return total
}
