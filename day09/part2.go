package day09

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day09/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var histories = parseInput(input, true)
	var result = calculateNextHistorySum(histories)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}
