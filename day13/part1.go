package day13

import (
	"aoc2023/assert"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day13/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var blocks = parseInput(input)
	var result = 0
	for _, lines := range blocks {
		result += findReflection(lines, false)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) [][]string {
	var blocks = strings.Split(string(input), "\n\n")
	var result [][]string
	for _, block := range blocks {
		var lines = strings.Split(block, "\n")

		if lines[len(lines)-1] == "" {
			result = append(result, lines[:len(lines)-1])
		} else {
			result = append(result, lines)
		}
	}
	return result
}

func findReflection(lines []string, rotated bool) int {
	var lastLine string
outer:
	for i, line := range lines {
		if lastLine == line {
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

	assert.Assert(!rotated, "should have found the result!")

	return findReflection(rotate90Deg(lines), true)
}

func rotate90Deg(lines []string) []string {
	var temp = create2dMatrix(len(lines), len(lines[0]))

	for y := len(lines) - 1; y >= 0; y-- {
		for x := range lines[y] {
			var c = string(lines[y][x])
			temp[x][len(lines)-1-y] = c
		}
	}

	var newLines []string
	for _, v := range temp {
		newLines = append(newLines, strings.Join(v, ""))
	}

	return newLines
}

func create2dMatrix(dx int, dy int) [][]string {
	var result = make([][]string, dy)

	for y := range dy {
		result[y] = make([]string, dx)
	}

	return result
}
