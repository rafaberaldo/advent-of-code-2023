package day23

import (
	"aoc2023/assert"
	"fmt"
	"os"
	"slices"
	"time"
)

func Part2() int {
	started := time.Now()
	input, err := os.ReadFile("day23/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	maze := parseInput(input)
	var finalists []Point
	// TODO: generate adjacent list from input/maze (for performance)
	walk(
		maze,
		Point{x: 1, y: 0},
		Point{x: len(maze[0]) - 2, y: len(maze) - 1},
		make(map[Point]bool),
		&finalists,
	)

	var result = slices.MaxFunc(finalists, sortBySteps).steps

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func walk(maze [][]string, current, end Point, visited map[Point]bool, finalists *[]Point) {
	if current.y < 0 || current.y > len(maze)-1 ||
		current.x < 0 || current.x > len(maze[0])-1 {
		return
	}

	if visited[current.state()] {
		return
	}

	if maze[current.y][current.x] == "#" {
		return
	}

	// printMaze(maze, current)

	if current.x == end.x && current.y == end.y {
		*finalists = append(*finalists, current)
		return
	}

	visited[current.state()] = true
	var backward = Direction{-current.dir.dx, -current.dir.dy}
	var dirs = dirsExcept(backward)
	for _, dir := range dirs {
		var edge = Point{current.x + dir.dx, current.y + dir.dy, current.steps + 1, dir}
		walk(maze, edge, end, visited, finalists)
	}
	// delete from visited after the loop to avoid copying multiple maps
	delete(visited, current.state())
}

func printMaze(maze [][]string, current Point) {
	time.Sleep(time.Millisecond * 25)
	var newMaze = make([][]string, len(maze))
	for i := range maze {
		newMaze[i] = make([]string, len(maze[i]))
		copy(newMaze[i], maze[i])
	}

	newMaze[current.y][current.x] = "O"
	for _, m := range newMaze {
		fmt.Println(m)
	}
	fmt.Printf("\n\n")
}
