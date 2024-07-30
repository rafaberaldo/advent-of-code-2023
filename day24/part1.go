package day24

import (
	"aoc2023/assert"
	_slices "aoc2023/lib"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

const (
	// P_MIN = 7
	// P_MAX = 27
	P_MIN = 2e14
	P_MAX = 4e14
)

func Part1() int {
	started := time.Now()
	input, err := os.ReadFile("day24/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	var vertices = parseInput(input)
	var pairs = createPairs(vertices)
	var result = simulate(pairs)

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) []Vertex {
	lines := strings.Split(string(input), "\n")
	var result []Vertex
	for _, line := range lines {
		if line == "" {
			continue
		}
		var posStr = strings.Split(line, "@")[0]
		var pos = _slices.StrToFloat(strings.Split(posStr, ", "))

		var velStr = strings.Split(line, "@")[1]
		var vel = _slices.StrToFloat(strings.Split(velStr, ", "))

		result = append(result, Vertex{
			pos[0], pos[1], pos[2], vel[0], vel[1], vel[2], line,
		})
	}

	return result
}

func createPairs(vertices []Vertex) []Pair {
	var pairs []Pair
	for _, vtx1 := range vertices {
		for _, vtx2 := range vertices {
			if vtx1 == vtx2 || slices.Contains(pairs, Pair{vtx2, vtx1}) {
				continue
			}
			pairs = append(pairs, Pair{vtx1, vtx2})
		}
	}
	var total = len(vertices) * (len(vertices) - 1) / 2
	assert.Assert(total == len(pairs), "wrong number of pairs!")
	return pairs
}

func simulate(pairs []Pair) int {
	var result = 0

	for _, pair := range pairs {
		if !segmentIntersect(pair.a.line(), pair.b.line()) {
			continue
		}
		x, y, ok := intersection(pair.a.line(), pair.b.line())
		if !ok || !inRange(x, y) {
			continue
		}
		result++
	}

	return result
}

func inRange(x, y float64) bool {
	return P_MIN <= x && x <= P_MAX && P_MIN <= y && y <= P_MAX
}

func segmentIntersect(a, b Line) bool {
	ccw := func(A, B, C Point) bool {
		return (C.y-A.y)*(B.x-A.x) > (B.y-A.y)*(C.x-A.x)
	}

	A := Point{a.x1, a.y1}
	B := Point{a.x2, a.y2}
	C := Point{b.x1, b.y1}
	D := Point{b.x2, b.y2}

	return ccw(A, C, D) != ccw(B, C, D) && ccw(A, B, C) != ccw(A, B, D)
}

// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection
func intersection(a, b Line) (float64, float64, bool) {
	den := (a.x1-a.x2)*(b.y1-b.y2) - (a.y1-a.y2)*(b.x1-b.x2)
	if den == 0 {
		return 0, 0, false
	}

	x := ((a.x1*a.y2-a.y1*a.x2)*(b.x1-b.x2) - (a.x1-a.x2)*(b.x1*b.y2-b.y1*b.x2)) / den
	y := ((a.x1*a.y2-a.y1*a.x2)*(b.y1-b.y2) - (a.y1-a.y2)*(b.x1*b.y2-b.y1*b.x2)) / den

	return x, y, true
}
