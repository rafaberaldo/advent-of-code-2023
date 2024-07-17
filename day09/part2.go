package day09

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day09/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var histories = parseInput2(input)
	var result = calculateNextHistorySum(histories)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput2(input []byte) []History {
	var lines = strings.Split(string(input), "\n")
	var histories []History

	for _, line := range lines {
		if line == "" {
			continue
		}

		var fields = strings.Fields(line)
		var history = History{ident: line}
		var values []int
		for _, field := range fields {
			var n, _ = strconv.Atoi(field)
			values = append(values, n)
		}
		slices.Reverse(values)
		history.result = values[len(values)-1] // save the last value as result to be summed with others
		history.valuesList = append(history.valuesList, values)
		histories = append(histories, history)
	}

	return histories
}
