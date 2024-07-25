package day19

import (
	"aoc2023/assert"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day19/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var instMap, parts = parseInput(input)

	var result = 0
	for _, part := range parts {
		if isPartAccepted(instMap, part) {
			result += part.sum()
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

type Instruction string

func (inst Instruction) process(part Part) (string, bool) {
	var str = string(inst)
	if strings.ContainsAny(str, "<>") {
		var split = strings.Split(str, ":")
		var result = split[1]
		var value, err = strconv.Atoi(split[0][2:])
		assert.Assert(err == nil, "error while converting value to int")

		if str[1] == '<' && part.get(str[0]) < value {
			return result, true
		}
		if str[1] == '>' && part.get(str[0]) > value {
			return result, true
		}
		return "", false
	}

	return str, true
}

type Part struct {
	x int
	m int
	a int
	s int
}

func (p Part) get(c byte) int {
	switch c {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	}
	panic("could not find attr to return!")
}

func (p Part) sum() int {
	return p.x + p.m + p.a + p.s
}

func parseInput(input []byte) (map[string][]Instruction, []Part) {
	var lines = strings.Split(string(input), "\n")
	var instructionsMap = make(map[string][]Instruction)
	var parts []Part

	var parsingState = "instruction"
	for _, line := range lines {
		if line == "" {
			parsingState = "part"
			continue
		}

		if parsingState == "instruction" {
			var firstBracketIdx = strings.IndexByte(line, '{')
			var key = line[:firstBracketIdx]
			var value = line[firstBracketIdx+1 : len(line)-1]
			var insts []Instruction
			for _, inst := range strings.Split(value, ",") {
				insts = append(insts, Instruction(inst))
			}
			instructionsMap[key] = insts
			continue
		}

		line = line[1 : len(line)-1] // remove leading and trailing braces
		var fields = strings.Split(line, ",")
		var part = Part{}
		for _, field := range fields {
			var split = strings.Split(field, "=")
			var value, err = strconv.Atoi(split[1])
			assert.Assert(err == nil, "couldn't convert part value!")
			switch split[0] {
			case "x":
				part.x = value
			case "m":
				part.m = value
			case "a":
				part.a = value
			case "s":
				part.s = value
			}
		}
		parts = append(parts, part)
	}

	return instructionsMap, parts
}

func isPartAccepted(instMap map[string][]Instruction, part Part) bool {
	var currentInsts = instMap["in"]

	for {
		for _, inst := range currentInsts {
			var res, ok = inst.process(part)
			if !ok {
				continue
			}
			if res == "A" {
				return true
			}
			if res == "R" {
				return false
			}
			currentInsts = instMap[res]
			break
		}
	}
}
