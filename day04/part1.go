package day04

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	input, err := os.ReadFile("day04/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(input), "\n")

	var count = 0

	for row := 0; row < len(rows); row++ {
		if rows[row] == "" {
			continue
		}
		var numbers = strings.Split(strings.Split(rows[row], ":")[1], "|")
		var realNumbers = strings.Split(numbers[0], " ")
		var myNumbers = strings.Split(numbers[1], " ")

		var winningNumbers = []int{}
		for i := 0; i < len(realNumbers); i++ {
			for j := 0; j < len(myNumbers); j++ {
				if realNumbers[i] == myNumbers[j] {
					if value, err := strconv.Atoi(myNumbers[j]); err == nil {
						winningNumbers = append(winningNumbers, value)
					}
				}
			}
		}

		if len(winningNumbers) >= 1 {
			count += int(math.Pow(2, float64(len(winningNumbers)-1)))
		}
	}

	return count
}
