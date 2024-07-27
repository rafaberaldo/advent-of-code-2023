package day21

import (
	"aoc2023/assert"
	"fmt"
	"os"
	"strings"
	"time"
)

func Part1() int {
	started := time.Now()
	input, err := os.ReadFile("day21/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	start, maze := parseInput(input)
	var result = search(maze, start)

	// printMaze(maze, mapToSlice(result))

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return len(result)
}

func parseInput(input []byte) (Point, [][]string) {
	lines := strings.Split(string(input), "\n")
	maze := make([][]string, len(lines)-1)
	var start Point
	for y, line := range lines {
		if line == "" {
			continue
		}
		maze[y] = make([]string, len(line))
		for x, c := range line {
			if c == 'S' {
				start = Point{x, y}
			}
			maze[y][x] = string(c)
		}
	}
	return start, maze
}

// C S S   O R D E R
var dirs = []Direction{
	{0, -1}, // up
	{+1, 0}, // right
	{0, +1}, // down
	{-1, 0}, // left
}

const MAX_STEPS = 64

func search(maze [][]string, start Point) map[Point]bool {
	var visited = make(map[Point]bool)
	var steps = make(map[Point]int)
	var result = make(map[Point]bool)
	var queue = []Point{start}

	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true

		if steps[current]%2 == 0 {
			result[current] = true
		}

		if steps[current] == MAX_STEPS {
			continue
		}

		steps[current]++

		for _, dir := range dirs {
			var edge = Point{current.x + dir.dx, current.y + dir.dy}
			if edge.y < 0 || edge.y > len(maze)-1 ||
				edge.x < 0 || edge.x > len(maze[0])-1 {
				continue
			}

			if maze[edge.y][edge.x] == "#" {
				continue
			}

			if visited[edge] {
				continue
			}

			steps[edge] = steps[current]
			queue = append(queue, edge)
		}
	}

	return result
}

func printMaze(maze [][]string, points []Point) {
	for _, p := range points {
		maze[p.y][p.x] = "O"
	}

	for _, m := range maze {
		fmt.Println(m)
	}
	fmt.Printf("\n\n")
}

func mapToSlice[K comparable, V any](m map[K]V) []K {
	var result []K
	for k := range m {
		result = append(result, k)
	}
	return result
}
