package day11

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day11/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var result1 int
	var expand1 = 1
	var points = parseInput(input, expand1)
	var pairs = getPairs(points)
	for _, pair := range pairs {
		result1 += getDistance(pair.a, pair.b)
	}

	var result2 int
	var espand2 = 2
	points = parseInput(input, espand2)
	pairs = getPairs(points)
	for _, pair := range pairs {
		result2 += getDistance(pair.a, pair.b)
	}

	var slope = (result2 - result1) / (espand2 - expand1)
	var diff = result1 - slope
	var result = slope*1_000_000 + diff

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}
