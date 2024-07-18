package day10

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day10/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var maze, startPoint = parseInput(input)
	var visited = make(map[Point]bool)
	var path []Point
	walk(maze, startPoint, &visited, &path, "START")

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return getInnerPoints(path)
}

// Shoelace formula
func getArea(path []Point) int {
	var first = path[0]
	var last = path[len(path)-1]
	var sum = last.x*first.y - last.y*first.x

	for i := 0; i < len(path)-1; i++ {
		var curr = path[i]
		var next = path[i+1]
		sum += curr.x*next.y - curr.y*next.x
	}

	return int(math.Abs(float64(sum)) / 2)
}

// Pick's theorem
func getInnerPoints(path []Point) int {
	return getArea(path) - len(path)/2 + 1
}
