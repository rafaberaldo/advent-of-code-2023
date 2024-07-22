package day16

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
}

type Point2 struct {
	x       int
	y       int
	visited map[Point]Point
}

func (p *Point2) point() Point {
	return Point{p.x, p.y}
}

var dirs = struct {
	UP    Point
	DOWN  Point
	LEFT  Point
	RIGHT Point
}{
	Point{0, -1},
	Point{0, +1},
	Point{-1, 0},
	Point{+1, 0},
}

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day16/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var maze = parseInput(input)
	var visited = make(map[Point]bool)
	walk(maze, Point2{0, 0, make(map[Point]Point)}, dirs.RIGHT, &visited)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return len(visited)
}

func parseInput(input []byte) [][]string {
	var lines = strings.Split(string(input), "\n")
	var maze [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}
		var chars []string
		for _, c := range line {
			chars = append(chars, string(c))
		}
		maze = append(maze, chars)
	}

	// for _, m := range maze {
	// 	fmt.Println(m)
	// }

	return maze
}

// Direction it is walking to direction it has to go
var walkingMap = map[string]map[Point][]Point{
	"|": {
		dirs.LEFT:  {dirs.UP, dirs.DOWN},
		dirs.RIGHT: {dirs.UP, dirs.DOWN},
	},
	"-": {
		dirs.UP:   {dirs.LEFT, dirs.RIGHT},
		dirs.DOWN: {dirs.LEFT, dirs.RIGHT},
	},
	"/": {
		dirs.RIGHT: {dirs.UP},
		dirs.LEFT:  {dirs.DOWN},
		dirs.UP:    {dirs.RIGHT},
		dirs.DOWN:  {dirs.LEFT},
	},
	"\\": {
		dirs.RIGHT: {dirs.DOWN},
		dirs.LEFT:  {dirs.UP},
		dirs.UP:    {dirs.LEFT},
		dirs.DOWN:  {dirs.RIGHT},
	},
}

func walk(maze [][]string, current Point2, dir Point, visited *map[Point]bool) {
	if current.y < 0 || current.y >= len(maze) ||
		current.x < 0 || current.x >= len(maze[0]) {
		return
	}

	// stop infinite loop
	if (current.visited)[current.point()] == dir {
		return
	}

	(*visited)[current.point()] = true
	current.visited[current.point()] = dir
	// printMaze(maze, current.point())

	var tile = maze[current.y][current.x]
	if nextDirs, ok := walkingMap[tile][dir]; ok {
		for _, nextDir := range nextDirs {
			var walkTo = Point2{current.x + nextDir.x, current.y + nextDir.y, current.visited}
			walk(maze, walkTo, nextDir, visited)
		}
	} else {
		// keep walking
		var walkTo = Point2{current.x + dir.x, current.y + dir.y, current.visited}
		walk(maze, walkTo, dir, visited)
	}
}

func printMaze(maze [][]string, current Point) {
	time.Sleep(time.Millisecond * 250)
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
