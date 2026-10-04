// Command omni measures the distance between two points — a minimal,
// complete demonstration of the cmd/ + internal/ pattern.
package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wyattau/omnigo-template/internal/geometry"
)

func parsePoint(s string) (geometry.Point, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return geometry.Point{}, fmt.Errorf("want x,y, got %q", s)
	}
	x, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return geometry.Point{}, fmt.Errorf("x: %w", err)
	}
	y, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return geometry.Point{}, fmt.Errorf("y: %w", err)
	}
	return geometry.Point{X: x, Y: y}, nil
}

func main() {
	from := flag.String("from", "0,0", "origin point as x,y")
	to := flag.String("to", "3,4", "target point as x,y")
	flag.Parse()

	a, err := parsePoint(*from)
	if err != nil {
		log.Fatalf("bad --from: %v", err)
	}
	b, err := parsePoint(*to)
	if err != nil {
		log.Fatalf("bad --to: %v", err)
	}
	fmt.Printf("%.6f\n", geometry.Distance(a, b))
}
