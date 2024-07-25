package day18

import (
	"aoc2023/assert"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day18/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var instructions = parseInput2(input)
	var result = calculateMap(instructions)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput2(input []byte) []Instruction {
	var lines = strings.Split(string(input), "\n")
	var instructions []Instruction

	var getDirection = func(b byte) Direction {
		switch b {
		case '3': // U
			return Direction{0, -1}
		case '1': // D
			return Direction{0, +1}
		case '2': // L
			return Direction{-1, 0}
		case '0': // R
			return Direction{+1, 0}
		}
		panic("couldn't find direction!")
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		var color = strings.Fields(line)[2]
		var length, err = strconv.ParseInt(color[2:7], 16, 0)
		assert.Assert(err == nil, "error converting length!")
		instructions = append(instructions, Instruction{getDirection(color[7]), int(length)})
	}

	return instructions
}
