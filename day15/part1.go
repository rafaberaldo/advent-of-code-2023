package day15

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day15/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var values = parseInput(input)
	var result int
	for _, v := range values {
		result += hash(v)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) []string {
	var lines = strings.Split(string(input), "\n")
	var values []string
	for _, line := range lines {
		if line == "" {
			continue
		}
		values = append(values, strings.Split(line, ",")...)
	}
	return values
}

func hash(str string) int {
	var result int
	for _, c := range str {
		result += int(c)
		result = result * 17 % 256
	}
	return result
}
