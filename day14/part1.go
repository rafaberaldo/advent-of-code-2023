package day14

import (
	"aoc2023/assert"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day14/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var matrix, points = parseInput(input)
	var newPoints []Point
	for i := range points {
		var newPoint = walk(matrix, points[i], points[i], Point{0, -1})
		newPoints = append(newPoints, newPoint)
		updateMatrix(&matrix, points[i], newPoint)
	}

	var result = calculateResult(len(matrix), newPoints)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) ([][]string, []Point) {
	var lines = strings.Split(string(input), "\n")
	var matrix [][]string
	var points []Point

	for y, line := range lines {
		if line == "" {
			continue
		}
		var chars []string
		for x, c := range line {
			chars = append(chars, string(c))
			if c == 'O' {
				points = append(points, Point{x, y})
			}
		}
		matrix = append(matrix, chars)
	}

	return matrix, points
}

type Point struct {
	x int
	y int
}

func walk(matrix [][]string, current Point, start Point, dir Point) Point {
	if current.y < 0 || current.y > len(matrix)-1 ||
		current.x < 0 || current.x > len(matrix[0])-1 {
		return Point{current.x - dir.x, current.y - dir.y}
	}

	if slices.Contains([]string{"#", "O"}, matrix[current.y][current.x]) && current != start {
		return Point{current.x - dir.x, current.y - dir.y}
	}

	return walk(matrix, Point{current.x + dir.x, current.y + dir.y}, start, dir)
}

func updateMatrix(matrix *[][]string, old Point, new Point) {
	if old == new {
		return
	}

	assert.Assert((*matrix)[old.y][old.x] == "O", "old point should be an 'O'!")
	assert.Assert((*matrix)[new.y][new.x] != "#", "new point should not be a '#'!")
	assert.Assert((*matrix)[new.y][new.x] == ".", "new point should be a '.'!")

	(*matrix)[old.y][old.x] = "."
	(*matrix)[new.y][new.x] = "O"
}

func calculateResult(height int, points []Point) int {
	var result = 0
	for _, point := range points {
		result += height - point.y
	}
	return result
}
