package day22

import (
	"aoc2023/assert"
	_slices "aoc2023/lib"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func Part1() int {
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
		if _slices.Compare(bricks1, bricks2) {
			result++
		}
	}

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) []Brick {
	lines := strings.Split(string(input), "\n")

	var bricks []Brick
	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, "~")
		start := _slices.StrToInt(strings.Split(split[0], ","))
		end := _slices.StrToInt(strings.Split(split[1], ","))
		brick := Brick{Point{start[0], start[1], start[2]}, Point{end[0], end[1], end[2]}}
		bricks = append(bricks, brick)
	}

	return bricks
}

func settleBricks(bricks []Brick) []Brick {
	var newBricks []Brick
	for _, brick := range bricks {
		brick := settle(brick, newBricks)
		newBricks = append(newBricks, brick)
	}
	return newBricks
}

func sortByZ(a, b Brick) int {
	return cmp.Compare(a.start.z, b.end.z)
}

func settle(current Brick, prevs []Brick) Brick {
	if current.start.z < 1 {
		current.start.z++
		current.end.z++
		return current
	}

	for _, prev := range prevs {
		if collide(current, prev) {
			current.start.z++
			current.end.z++
			return current
		}
	}

	current.start.z--
	current.end.z--
	return settle(current, prevs)
}

// https://developer.mozilla.org/en-US/docs/Games/Techniques/2D_collision_detection
func collide(a Brick, b Brick) bool {
	if a.start.z != b.end.z {
		return false
	}

	var aWidth = a.end.x - a.start.x
	var aHeight = a.end.y - a.start.y
	var bWidth = b.end.x - b.start.x
	var bHeight = b.end.y - b.start.y

	if a.start.x <= b.start.x+bWidth &&
		a.start.x+aWidth >= b.start.x &&
		a.start.y <= b.start.y+bHeight &&
		a.start.y+aHeight >= b.start.y {
		return true
	}

	return false
}
