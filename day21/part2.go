package day21

import (
	"aoc2023/assert"
	"fmt"
	"os"
	"time"
)

func Part2() int {
	started := time.Now()
	input, err := os.ReadFile("day21/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	start, maze := parseInput(input)
	var steps = search2(maze, start)

	var result = calculateResult(len(maze), steps, 26501365)
	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func search2(maze [][]string, start Point) map[Point]int {
	var visited = make(map[Point]bool)
	var steps = make(map[Point]int)
	var queue = []Point{start}
	steps[start] = 0

	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true

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

			steps[edge] = steps[current] + 1
			queue = append(queue, edge)
		}
	}

	return steps
}

// tbh could't find the solution for this by myself
// https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21
func calculateResult(gridSize int, steps map[Point]int, maxSteps int) int {
	var radius = gridSize / 2
	var radiusMaxSteps = (maxSteps - radius) / gridSize
	assert.Assert(radiusMaxSteps == 202300, "wrong grid size!")

	var perfectEven = radiusMaxSteps * radiusMaxSteps            // the num of even tiles if grid didn't have obstacles (#)
	var perfectOdd = (radiusMaxSteps + 1) * (radiusMaxSteps + 1) // the num of odd tiles if grid didn't have obstacles (#)

	var filterMap = func(fn func(steps int) bool) int {
		var result = 0
		for _, v := range steps {
			if fn(v) {
				result++
			}
		}
		return result
	}

	var allEven = filterMap(func(steps int) bool { return steps%2 == 0 })
	var allOdd = filterMap(func(steps int) bool { return steps%2 != 0 })
	var cornersEven = filterMap(func(steps int) bool { return steps > radius && steps%2 == 0 })
	var cornersOdd = filterMap(func(steps int) bool { return steps > radius && steps%2 != 0 })

	return perfectOdd*allOdd +
		perfectEven*allEven -
		((radiusMaxSteps + 1) * cornersOdd) +
		(radiusMaxSteps * cornersEven)
}
