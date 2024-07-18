package day10

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day10/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var maze, startPoint = parseInput(input)
	var visited = make(map[Point]bool)
	var path []Point
	walk(maze, startPoint, &visited, &path, "START")
	// fmt.Printf("Starting point: %#v\n", startPoint)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return len(path) / 2
}

type Point struct {
	x int
	y int
}

func parseInput(input []byte) ([][]string, Point) {
	var padded = addPadding(input)
	var lines = strings.Split(string(padded), "\n")

	var maze = make([][]string, len(lines))
	var startPoint Point

	for y, line := range lines {
		if line == "" {
			continue
		}

		maze[y] = make([]string, len(line))

		for x, c := range line {
			maze[y][x] = string(c)

			if c == 'S' {
				startPoint = Point{x, y}
			}
		}
	}

	return maze, startPoint
}

func addPadding(input []byte) string {
	var result = "." + strings.ReplaceAll(string(input), "\n", ".\n.")
	result = result[:len(result)-2] // remove last '\n.' from last line

	var lineLength = len(strings.Split(result, "\n")[0])
	var emptyLine = strings.Repeat(".", lineLength)
	result = emptyLine + "\n" + result + "\n" + emptyLine

	return result
}

var directions = [][]int{
	{0, -1}, // up
	{+1, 0}, // right
	{0, +1}, // down
	{-1, 0}, // left
}

func walk(maze [][]string, current Point, visited *map[Point]bool, path *[]Point, move string) bool {

	if len(*path) > 0 {
		var last = (*path)[len(*path)-1]
		if !isValidMove(maze[current.y][current.x], maze[last.y][last.x], move) {
			return false
		}

		// fmt.Printf("Last: %#v ---- Current: %#v (%v)\n", maze[last.y][last.x], maze[current.y][current.x], move)
	}

	if maze[current.y][current.x] == "S" && move != "START" {
		return true
	}

	if (*visited)[Point{current.x, current.y}] {
		return false
	}

	(*visited)[current] = true
	*path = append(*path, current)

	for i, dir := range directions {
		var next = Point{x: current.x + dir[0], y: current.y + dir[1]}
		var nextMove string
		switch i {
		case 0:
			nextMove = "UP"
		case 1:
			nextMove = "RIGHT"
		case 2:
			nextMove = "DOWN"
		case 3:
			nextMove = "LEFT"
		}

		if walk(maze, next, visited, path, nextMove) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func isValidMove(current string, last string, move string) bool {
	var up = []string{"|", "F", "7"}
	var down = []string{"|", "J", "L"}
	var left = []string{"-", "F", "L"}
	var right = []string{"-", "J", "7"}

	// win condition, inverted
	if current == "S" {
		switch move {
		case "UP":
			return slices.Contains(down, last)
		case "DOWN":
			return slices.Contains(up, last)
		case "LEFT":
			return slices.Contains(right, last)
		case "RIGHT":
			return slices.Contains(left, last)
		}
	}

	if last == "S" {
		switch move {
		case "UP":
			return slices.Contains(up, current)
		case "DOWN":
			return slices.Contains(down, current)
		case "LEFT":
			return slices.Contains(left, current)
		case "RIGHT":
			return slices.Contains(right, current)
		}
	}

	if last == "|" {
		switch move {
		case "UP":
			return slices.Contains(up, current)
		case "DOWN":
			return slices.Contains(down, current)
		case "LEFT":
			return false
		case "RIGHT":
			return false
		}
	}

	if last == "-" {
		switch move {
		case "UP":
			return false
		case "DOWN":
			return false
		case "LEFT":
			return slices.Contains(left, current)
		case "RIGHT":
			return slices.Contains(right, current)
		}
	}

	if last == "F" {
		switch move {
		case "UP":
			return false
		case "DOWN":
			return slices.Contains(down, current)
		case "LEFT":
			return false
		case "RIGHT":
			return slices.Contains(right, current)
		}
	}

	if last == "7" {
		switch move {
		case "UP":
			return false
		case "DOWN":
			return slices.Contains(down, current)
		case "LEFT":
			return slices.Contains(left, current)
		case "RIGHT":
			return false
		}
	}

	if last == "J" {
		switch move {
		case "UP":
			return slices.Contains(up, current)
		case "DOWN":
			return false
		case "LEFT":
			return slices.Contains(left, current)
		case "RIGHT":
			return false
		}
	}

	if last == "L" {
		switch move {
		case "UP":
			return slices.Contains(up, current)
		case "DOWN":
			return false
		case "LEFT":
			return false
		case "RIGHT":
			return slices.Contains(right, current)
		}
	}

	panic("No options are valid?!")
}
