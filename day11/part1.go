package day11

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day11/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var points = parseInput(input, 2)
	var pairs = getPairs(points)
	var result = 0
	for _, pair := range pairs {
		result += getDistance(pair.a, pair.b)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

type Point struct {
	x     int
	y     int
	ident string
}

func parseInput(input []byte, expand int) []Point {
	var lines = expandEmptyLines(input, expand-1)
	var tilted = rotate90Deg(lines)
	lines = expandEmptyLines(tilted, expand-1)

	var image = make([][]string, len(lines))
	var galaxies []Point

	for y, line := range lines {
		image[y] = make([]string, len(line))

		for x, c := range line {
			image[y][x] = string(c)

			if c == '#' {
				galaxies = append(galaxies, Point{x, y, string(c)})
			}
		}
	}

	return galaxies
}

func expandEmptyLines(input []byte, times int) []string {
	var lines = strings.Split(string(input), "\n")
	var emptyLine = strings.Repeat(".", len(lines[0]))

	// Remove last empty line
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	for i := 0; i < len(lines); i++ {
		if lines[i] == emptyLine {
			for range times {
				lines = slices.Insert(lines, i+1, emptyLine)
				i++
			}
		}
	}

	return lines
}

func rotate90Deg(lines []string) []byte {
	var temp = create2dSlice(len(lines), len(lines[0]))

	for y := len(lines) - 1; y >= 0; y-- {
		for x := range lines[y] {
			var c = string(lines[y][x])
			temp[x][len(lines)-1-y] = c
		}
	}

	var newLines []string
	for _, v := range temp {
		newLines = append(newLines, strings.Join(v, ""))
	}

	return []byte(strings.Join(newLines, "\n"))
}

func create2dSlice(dx int, dy int) [][]string {
	var result = make([][]string, dy)

	for y := range dy {
		result[y] = make([]string, dx)
	}

	return result
}

type Pair struct {
	a Point
	b Point
}

func getPairs(points []Point) []Pair {
	var pairs []Pair
	for _, pointA := range points {
		for _, pointB := range points {
			if pointA == pointB {
				continue
			}
			if slices.Contains(pairs, Pair{pointB, pointA}) {
				continue
			}
			pairs = append(pairs, Pair{pointA, pointB})
		}
	}

	var total = len(points) * (len(points) - 1) / 2
	if total != len(pairs) {
		log.Fatalf("Wrong pair number! Expected %v, received %v", total, len(pairs))
	}

	return pairs
}

// Manhattan Distance / Taxicab geometry
func getDistance(pointA Point, pointB Point) int {
	var a = math.Abs(float64(pointA.x - pointB.x))
	var b = math.Abs(float64(pointA.y - pointB.y))
	return int(a + b)
}
