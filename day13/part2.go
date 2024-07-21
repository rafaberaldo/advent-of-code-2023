package day13

import (
	"aoc2023/assert"
	"fmt"
	"log"
	"os"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day13/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var blocks = parseInput(input)

	var p1ResultMap = make(map[int]int)
	for i, lines := range blocks {
		var result = findReflection(lines, false)
		p1ResultMap[i] = result
	}

	var result = 0
	for i, lines := range blocks {
		var res = findReflection2Pre(lines, false, p1ResultMap[i]/100)
		if res == 0 {
			res = findReflection2Pre(rotate90Deg(lines), true, p1ResultMap[i])
		}
		assert.Assert(res > 0, "should have found the result!")
		result += res
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

type Point struct {
	x int
	y int
}

func findReflection2Pre(lines []string, rotated bool, excludeIdx int) int {
	var linesOffByOne = make(map[Point]bool)

	for y, line1 := range lines {
		for _, line2 := range lines {
			if x := isOffByOne(line1, line2); x != -1 {
				linesOffByOne[Point{x, y}] = true
			}
		}
	}

	var result int
	for point := range linesOffByOne {
		var res = findReflection2(swapChar(lines, point), rotated, excludeIdx)
		if res > 0 {
			result += res
			break
		}
	}

	return result
}

func findReflection2(lines []string, rotated bool, excludeIdx int) int {
	var lastLine string
outer:
	for i, line := range lines {
		if lastLine == line && i != excludeIdx {
			var k = i
			for j := i - 1; j >= 0; j-- {
				if k >= len(lines) {
					break
				}
				if lines[j] != lines[k] {
					continue outer
				}
				k++
			}

			if rotated {
				return i
			} else {
				return i * 100
			}
		}
		lastLine = line
	}

	return 0
}

func isOffByOne(a string, b string) int {
	if a == b {
		return -1
	}

	var count = 0
	var charIndex = -1
	for i := range a {
		if a[i] != b[i] {
			count++
			charIndex = i
		}
		if count > 1 {
			return -1
		}
	}

	assert.Assert(count == 1, "should not be diff than 1!")

	return charIndex
}

func swapChar(lines []string, point Point) []string {
	var newLines = make([]string, len(lines))
	copy(newLines, lines)

	if newLines[point.y][point.x] == '#' {
		newLines[point.y] = replaceAtIndex(newLines[point.y], '.', point.x)
	} else {
		newLines[point.y] = replaceAtIndex(newLines[point.y], '#', point.x)
	}

	return newLines
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
