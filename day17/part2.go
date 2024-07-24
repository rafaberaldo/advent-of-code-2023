package day17

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Part2() int {
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
		4,
		10,
	)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}
