package day23

import (
	"aoc2023/assert"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func Part1() int {
	started := time.Now()
	input, err := os.ReadFile("day23/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	maze := parseInput(input)
	result := search(maze, Point{x: 1, y: 0}, Point{x: len(maze[0]) - 2, y: len(maze) - 1})

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) [][]string {
	lines := strings.Split(string(input), "\n")
	maze := make([][]string, len(lines)-1)
	for y, line := range lines {
		if line == "" {
			continue
		}
		maze[y] = strings.Split(line, "")
	}

	return maze
}

var UP = Direction{0, -1}
var RIGHT = Direction{+1, 0}
var DOWN = Direction{0, +1}
var LEFT = Direction{-1, 0}

func search(maze [][]string, start, end Point) int {
	var visited = make(map[Point]bool)
	var queue = []Point{start}
	var finalists []Point

	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true

		if current.x == end.x && current.y == end.y {
			finalists = append(finalists, current)
			continue
		}

		var backward = Direction{-current.dir.dx, -current.dir.dy}
		var dirs = dirsExcept(backward)
		for _, dir := range dirs {
			var edge = Point{current.x + dir.dx, current.y + dir.dy, current.steps + 1, dir}

			if edge.y < 0 || edge.y > len(maze)-1 ||
				edge.x < 0 || edge.x > len(maze[0])-1 {
				continue
			}

			if !canWalkInto(maze, edge, dir) {
				continue
			}

			if visited[edge] {
				continue
			}

			queue = append(queue, edge)
		}
	}

	return slices.MaxFunc(finalists, sortBySteps).steps
}

func dirsExcept(dir Direction) []Direction {
	return slices.DeleteFunc(
		[]Direction{UP, RIGHT, DOWN, LEFT},
		func(d Direction) bool {
			return d == dir
		})
}

func sortBySteps(a, b Point) int {
	return cmp.Compare(a.steps, b.steps)
}

func canWalkInto(maze [][]string, point Point, dir Direction) bool {
	var tile = maze[point.y][point.x]
	if tile == "#" {
		return false
	}

	if dir == UP && tile == "v" {
		return false
	}

	if dir == DOWN && tile == "^" {
		return false
	}

	if dir == LEFT && tile == ">" {
		return false
	}

	if dir == RIGHT && tile == "<" {
		return false
	}

	return true
}
