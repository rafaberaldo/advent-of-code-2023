package day12

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day12/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var puzzle = parseInput(input)

	var result = 0
	for _, p := range puzzle {
		var memo = make(map[string]int)
		result += findPattern(p, &memo)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

type Puzzle struct {
	line    string
	damaged []int8
}

func (p *Puzzle) key() string {
	return p.line + toStringSlice(p.damaged)
}

func parseInput(input []byte) []Puzzle {
	var lines = strings.Split(string(input), "\n")
	var puzzle []Puzzle
	for _, line := range lines {
		if line == "" {
			continue
		}
		var fields = strings.Fields(line)
		var condition = toIntSlice(strings.Split(fields[1], ","))
		puzzle = append(puzzle, Puzzle{fields[0], condition})
	}
	return puzzle
}

func toIntSlice(str []string) []int8 {
	var result []int8
	for i := 0; i < len(str); i++ {
		if v, err := strconv.Atoi(str[i]); err == nil {
			result = append(result, int8(v))
		}
	}
	return result
}

func toStringSlice(numbers []int8) string {
	var result = make([]string, len(numbers))
	for i, v := range numbers {
		result[i] = strconv.Itoa(int(v))
	}

	return strings.Join(result, ",")
}

func findPattern(p Puzzle, memo *map[string]int) int {
	// fmt.Printf("Testing %#v -- %#v\n", p.line, p.damaged)

	if len(p.line) == 0 {
		return 0
	}

	if len(p.damaged) == 0 {
		if !strings.Contains(p.line, "#") {
			return 1
		}
		return 0
	}

	if v, ok := (*memo)[p.key()]; ok {
		return v
	}

	if p.line[0] == '.' {
		var next = Puzzle{p.line[1:], p.damaged}
		(*memo)[next.key()] = findPattern(next, memo)
		return (*memo)[next.key()]
	}

	if p.line[0] == '?' {
		var next1 = Puzzle{"#" + p.line[1:], p.damaged}
		var next2 = Puzzle{p.line[1:], p.damaged}
		(*memo)[next1.key()] = findPattern(next1, memo)
		(*memo)[next2.key()] = findPattern(next2, memo)
		return (*memo)[next1.key()] + (*memo)[next2.key()]
	}

	if p.line[0] == '#' {
		var brokenCount = countBroken(p.line)
		if brokenCount == int(p.damaged[0]) {
			if brokenCount+1 >= len(p.line) {
				if len(p.damaged) == 1 {
					return 1
				}
				return 0
			}
			var next = Puzzle{p.line[brokenCount+1:], p.damaged[1:]}
			return findPattern(next, memo)
		}

		if brokenCount < len(p.line) && p.line[brokenCount] == '?' {
			var nextLine = p.line[:brokenCount] + "#" + p.line[brokenCount+1:]
			var next = Puzzle{nextLine, p.damaged}
			return findPattern(next, memo)
		}
	}

	return 0
}

func countBroken(str string) int {
	var position = 0
	for position < len(str) && str[position] == '#' {
		position++
	}
	return position
}
