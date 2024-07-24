package day17

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
}

type Vertex struct {
	point Point
	dist  int
	dir   Direction
	prev  VertexState
	steps int
}

type VertexState struct {
	point Point
	dir   Direction
	steps int
}

func (v *Vertex) state() VertexState {
	return VertexState{v.point, v.dir, v.steps}
}

type Direction struct {
	dx int
	dy int
}

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day17/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var maze = parseInput(input)
	var result = findPath(
		maze,
		Point{0, 0},
		Point{len(maze[0]) - 1, len(maze) - 1},
	)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) [][]int {
	var lines = strings.Split(string(input), "\n")
	var maze [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}
		var chars []int
		for _, c := range line {
			var num, _ = strconv.Atoi(string(c))
			chars = append(chars, num)
		}
		maze = append(maze, chars)
	}

	return maze
}

var dirs = []Direction{
	{0, -1}, // up
	{+1, 0}, // right
	{0, +1}, // down
	{-1, 0}, // left
}

func findPath(maze [][]int, init Point, end Point) int {
	var visited = make(map[VertexState]bool)
	var dests []Vertex
	dests = append(dests, Vertex{point: init})

	var endVertex Vertex
	for {
		slices.SortFunc(dests, func(a Vertex, b Vertex) int { return cmp.Compare(a.dist, b.dist) })
		var current = dests[0]
		dests = dests[1:]

		if visited[current.state()] {
			continue
		}

		visited[current.state()] = true

		if current.point == end {
			endVertex = current
			break
		}

		for _, dir := range dirs {
			if current.dir == dir && current.steps == 3 {
				continue
			}

			var inverseDir = Direction{-current.dir.dx, -current.dir.dy}
			// cannot go backwards
			if dir == inverseDir {
				continue
			}

			var point = Point{current.point.x + dir.dx, current.point.y + dir.dy}
			if point.y < 0 || point.y > len(maze)-1 ||
				point.x < 0 || point.x > len(maze[0])-1 {
				continue
			}

			var steps = 1
			if dir == current.dir {
				steps = current.steps + 1
			}
			var edge = Vertex{
				point: point,
				dist:  current.dist + maze[point.y][point.x],
				dir:   dir,
				prev:  current.state(),
				steps: steps,
			}

			if visited[edge.state()] {
				continue
			}

			dests = append(dests, edge)
		}
	}

	return endVertex.dist
}
