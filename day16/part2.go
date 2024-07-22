package day16

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day16/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var result int
	var maze = parseInput(input)

	for y := range len(maze) {
		var visitedLeft = make(map[Point]bool)
		var visitedRight = make(map[Point]bool)
		walk(maze, Point2{0, y, make(map[Point]Point)}, dirs.RIGHT, &visitedRight)
		walk(maze, Point2{len(maze[0][0]) - 1, y, make(map[Point]Point)}, dirs.LEFT, &visitedLeft)
		if len(visitedLeft) > result {
			result = len(visitedLeft)
		}
		if len(visitedRight) > result {
			result = len(visitedRight)
		}
	}

	for x := range len(maze[0]) {
		var visitedUp = make(map[Point]bool)
		var visitedDown = make(map[Point]bool)
		walk(maze, Point2{x, 0, make(map[Point]Point)}, dirs.DOWN, &visitedDown)
		walk(maze, Point2{x, len(maze) - 1, make(map[Point]Point)}, dirs.UP, &visitedUp)
		if len(visitedUp) > result {
			result = len(visitedUp)
		}
		if len(visitedDown) > result {
			result = len(visitedDown)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}
