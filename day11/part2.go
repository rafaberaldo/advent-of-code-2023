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

	var c = make(chan int)
	defer close(c)

	var expand1 = 1
	go getResult(input, expand1, c)
	var result1 = <-c

	var expand2 = 2
	go getResult(input, expand2, c)
	var result2 = <-c

	var slope = (result2 - result1) / (expand2 - expand1)
	var diff = result1 - slope
	var result = slope*1_000_000 + diff

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func getResult(input []byte, expand int, c chan int) {
	var result int
	var points = parseInput(input, expand)
	var pairs = getPairs(points)
	for _, pair := range pairs {
		result += getDistance(pair.a, pair.b)
	}
	c <- result
}
