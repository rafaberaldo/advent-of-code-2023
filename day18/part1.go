package day18

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day18/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var instructions = parseInput(input)
	var result = calculateMap(instructions)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

type Direction struct {
	dx int
	dy int
}

type Instruction struct {
	dir Direction
	len int
}

type Point struct {
	x int
	y int
}

func parseInput(input []byte) []Instruction {
	var lines = strings.Split(string(input), "\n")
	var instructions []Instruction

	var getDirection = func(str string) Direction {
		switch str {
		case "U":
			return Direction{0, -1}
		case "D":
			return Direction{0, +1}
		case "L":
			return Direction{-1, 0}
		case "R":
			return Direction{+1, 0}
		}
		panic("couldn't find direction!")
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		var fields = strings.Fields(line)
		var length, _ = strconv.Atoi(fields[1])
		instructions = append(instructions, Instruction{getDirection(fields[0]), length})
	}

	return instructions
}

func calculateMap(instructions []Instruction) int {
	var path = []Point{{0, 0}}
	for _, inst := range instructions {
		for range inst.len {
			var last = path[len(path)-1]
			path = append(path, Point{last.x + inst.dir.dx, last.y + inst.dir.dy})
		}
	}
	path = path[:len(path)-1] // remove last (0,0)
	return getInnerPoints(path) + len(path)
}

// Shoelace formula
func getArea(path []Point) int {
	var first = path[0]
	var last = path[len(path)-1]
	var sum = last.x*first.y - last.y*first.x

	for i := 0; i < len(path)-1; i++ {
		var curr = path[i]
		var next = path[i+1]
		sum += curr.x*next.y - curr.y*next.x
	}

	return int(math.Abs(float64(sum)) / 2)
}

// Pick's theorem
func getInnerPoints(path []Point) int {
	return getArea(path) - len(path)/2 + 1
}
