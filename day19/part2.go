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

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day19/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var instMap, _ = parseInput(input)
	var acceptedRanges []PartRange
	var partRange = PartRange{
		x: Range{1, 4001},
		m: Range{1, 4001},
		a: Range{1, 4001},
		s: Range{1, 4001},
	}
	findAcceptedRanges(instMap, "in", partRange, &acceptedRanges)

	var result = 0
	for _, pr := range acceptedRanges {
		result += pr.sum()
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func findAcceptedRanges(
	instMap map[string][]Instruction,
	instKey string,
	partRange PartRange,
	acceptedRanges *[]PartRange,
) {
	if instKey == "R" {
		return
	}
	if instKey == "A" {
		*acceptedRanges = append(*acceptedRanges, partRange)
		return
	}

	for _, inst := range instMap[instKey] {
		var exp, ok = inst.expression()
		if !ok {
			var next = inst.next()
			findAcceptedRanges(instMap, next, partRange, acceptedRanges)
			return
		}

		var pr1, pr2 = partRange.divide(exp)
		// one of them will succeed and recurse, the other keeps looping
		if next, ok := inst.process(pr1.avgPart()); ok {
			findAcceptedRanges(instMap, next, pr1, acceptedRanges)
			partRange = pr2
		}
		if next, ok := inst.process(pr2.avgPart()); ok {
			findAcceptedRanges(instMap, next, pr2, acceptedRanges)
			partRange = pr1
		}
	}

	panic("THIS IS ILLEGAL!")
}

type Expression struct {
	ident  byte
	symbol byte
	value  int
	next   string
}

func (inst Instruction) expression() (Expression, bool) {
	var str = string(inst)
	if strings.ContainsAny(str, "<>") {
		var split = strings.Split(str, ":")
		var next = split[1]
		var value, err = strconv.Atoi(split[0][2:])
		assert.Assert(err == nil, "error while converting value to int")
		return Expression{str[0], str[1], value, next}, true
	}

	return Expression{}, false
}

func (inst Instruction) next() string {
	var str = string(inst)
	assert.Assert(!strings.ContainsAny(str, "<>"), "should not be calling next() if its an expression!")

	return str
}

type Range struct {
	min int
	max int
}

type PartRange struct {
	x Range
	m Range
	a Range
	s Range
}

func (pr PartRange) avgPart() Part {
	var x = (pr.x.min + pr.x.max) / 2
	var m = (pr.m.min + pr.m.max) / 2
	var a = (pr.a.min + pr.a.max) / 2
	var s = (pr.s.min + pr.s.max) / 2
	return Part{x, m, a, s}
}

func (pr PartRange) divide(exp Expression) (PartRange, PartRange) {
	var pr1, pr2 = pr, pr
	var offset = 0
	if exp.symbol == '>' {
		offset = 1
	}

	switch exp.ident {
	case 'x':
		pr1.x.min = exp.value + offset
		pr2.x.max = exp.value + offset
	case 'm':
		pr1.m.min = exp.value + offset
		pr2.m.max = exp.value + offset
	case 'a':
		pr1.a.min = exp.value + offset
		pr2.a.max = exp.value + offset
	case 's':
		pr1.s.min = exp.value + offset
		pr2.s.max = exp.value + offset
	}

	return pr1, pr2
}

func (pr PartRange) sum() int {
	r := pr.x.max - pr.x.min
	r *= pr.m.max - pr.m.min
	r *= pr.a.max - pr.a.min
	r *= pr.s.max - pr.s.min
	return r
}
