package day14

import (
	"aoc2023/assert"
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
	var matrix, points = parseInput(input)

	var dirs = []Point{
		{0, -1}, // north
		{-1, 0}, // west
		{0, +1}, // south
		{+1, 0}, // east
	}

	// key = result, value = array of step numbers (idx+1)
	// used to calculate the cycle/steps number
	var stepsMap = make(map[int][]int)

	// key = step, value = result
	// used to get the result w/o having to run again
	var resultMap = make(map[int]int)

	// can probably find the result in the first 350 steps
	// increase this if it's not finding the result
	for inc := range 350 {
		for iDir, dir := range dirs {
			var newPoints []Point
			for i := range points {
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
		stepsMap[res] = append(stepsMap[res], inc+1)
		resultMap[inc+1] = res
	}

	var steps, first = findCycle(stepsMap)
	assert.Assert(steps > 0, "could not find cycle!")

	var result, ok = resultMap[findResultStep(steps, first)]
	assert.Assert(ok, "should have found the result! increase step count")

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
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

func findCycle(m map[int][]int) (int, int) {
	var find = func(values []int) int {
		assert.Assert(len(values) >= 3, "slice length should be >= 3!")

		var steps = values[len(values)-1] - values[len(values)-2]
		for i := 1; i < len(values); i++ {
			if values[i]-values[i-1] != steps {
				return -1
			}
		}
		return steps
	}

	for _, values := range m {
		if len(values) < 3 {
			continue
		}

		if steps := find(values); steps > 0 {
			return steps, values[0]
		}
	}

	return -1, -1
}

func findResultStep(steps int, first int) int {
	var result = 1_000_000_000 % steps
	for result < first {
		result += steps
	}
	return result
}
