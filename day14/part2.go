package day14

import (
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day14/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var dirs = []Point{
		{0, -1}, // north
		{-1, 0}, // west
		{0, +1}, // south
		{+1, 0}, // east
	}
	var matrix, points = parseInput(input)
	fmt.Println("total height", len(matrix))
	fmt.Println("total points", len(points))
	for inc := range 1000 {
		for iDir, dir := range dirs {
			var newPoints []Point
			for i := range points {
				// fmt.Println(points[i])
				var newPoint = walk(matrix, points[i], points[i], dir)
				newPoints = append(newPoints, newPoint)
				updateMatrix(&matrix, points[i], newPoint)
			}
			if iDir == 0 || iDir == 3 { // Before North or West
				slices.SortFunc(newPoints, sortUpDown)
			} else {
				slices.SortFunc(newPoints, sortDownUp)
			}
			points = newPoints
		}
		var res = calculateResult(len(matrix), points)
		fmt.Println(inc+1, " === ", res)
	}

	var result = calculateResult(len(matrix), points)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result // 93742
}

func sortUpDown(a Point, b Point) int {
	if a.y < b.y {
		return -1
	}
	if a.y > b.y {
		return +1
	}

	if a.x < b.x {
		return -1
	}
	if a.x > b.x {
		return +1
	}

	return 0
}

func sortDownUp(a Point, b Point) int {
	if a.y < b.y {
		return +1
	}
	if a.y > b.y {
		return -1
	}

	if a.x < b.x {
		return +1
	}
	if a.x > b.x {
		return -1
	}

	return 0
}
