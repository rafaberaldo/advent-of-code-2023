package day22

import (
	"aoc2023/assert"
	_slices "aoc2023/lib"
	"fmt"
	"os"
	"slices"
	"time"
)

func Part2() int {
	started := time.Now()
	input, err := os.ReadFile("day22/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	var bricks = parseInput(input)
	slices.SortFunc(bricks, sortByZ)
	var settledBricks = settleBricks(bricks)

	var result = 0
	for i := 0; i < len(settledBricks); i++ {
		var bricks1 = _slices.Delete(settledBricks, i, i+1)
		var bricks2 = settleBricks(bricks1)
		for i := range bricks1 {
			if bricks1[i].start.z > bricks2[i].start.z {
				result++
			}
		}
	}

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}
