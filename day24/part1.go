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
		if !hasIntersection(pair.lines()) {
			continue
		}
		if !inRange(intersection(pair.lines())) {
			continue
		}
		result++
	}

	return result
}

func inRange(x, y float64) bool {
	return P_MIN <= x && x <= P_MAX && P_MIN <= y && y <= P_MAX
}

func ccw(a, b, c Point) bool {
	return (c.y-a.y)*(b.x-a.x) > (b.y-a.y)*(c.x-a.x)
}

// https://bryceboe.com/2006/10/23/line-segment-intersection-algorithm/
func hasIntersection(a, b, c, d Point) bool {
	return ccw(a, c, d) != ccw(b, c, d) && ccw(a, b, c) != ccw(a, b, d)
}

// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection
func intersection(a, b, c, d Point) (float64, float64) {
	den := (a.x-b.x)*(c.y-d.y) - (a.y-b.y)*(c.x-d.x)
	assert.Assert(den != 0, "lines do not intersect!")

	n1 := a.x*b.y - a.y*b.x
	n2 := c.x*d.y - c.y*d.x
	x := (n1*(c.x-d.x) - (a.x-b.x)*n2) / den
	y := (n1*(c.y-d.y) - (a.y-b.y)*n2) / den

	return x, y
}
