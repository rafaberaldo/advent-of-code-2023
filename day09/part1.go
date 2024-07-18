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

type History struct {
	ident      string
	valuesList [][]int
	result     int
}

func Part1() int {
	var start = time.Now()
	input, err := os.ReadFile("day09/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var histories = parseInput(input, false)
	var result = calculateNextHistorySum(histories)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte, reverse bool) []History {
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
		if reverse {
			slices.Reverse(values)
		}
		history.result = values[len(values)-1] // save the last value as result to be summed with others
		history.valuesList = append(history.valuesList, values)
		histories = append(histories, history)
	}

	return histories
}

func calculateNextHistorySum(histories []History) int {
	var historiesCopy = make([]History, len(histories))
	copy(historiesCopy, histories)

	for historyIndex := range historiesCopy {
		var history = &historiesCopy[historyIndex]

		for vListIndex := 0; vListIndex < len(history.valuesList); vListIndex++ {
			var valueList = history.valuesList[vListIndex]
			var newValues []int

			if sameElements(valueList) {
				break
			}

			for i := 0; i < len(valueList)-1; i++ {
				newValues = append(newValues, valueList[i+1]-valueList[i])
			}

			history.result += newValues[len(newValues)-1]
			history.valuesList = append(history.valuesList, newValues)
		}
	}

	var sum = 0
	for _, h := range historiesCopy {
		// fmt.Printf("%#v\n", h)
		sum += h.result
	}
	return sum
}

func sameElements(values []int) bool {
	for i := 0; i < len(values)-1; i++ {
		if values[i+1] != values[i] {
			return false
		}
	}
	return true
}
