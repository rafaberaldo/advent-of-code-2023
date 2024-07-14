package day01

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	input, err := os.ReadFile("day01/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(input), "\n")
	count := 0

	for row := 0; row < len(rows); row++ {
		rowValue := ""

		for i := 0; i < len(rows[row]); i++ {
			if dig, err := strconv.Atoi(string(rows[row][i])); err == nil {
				rowValue = strconv.Itoa(dig)
				break
			}
		}

		for i := len(rows[row]) - 1; i >= 0; i-- {
			if dig, err := strconv.Atoi(string(rows[row][i])); err == nil {
				rowValue += strconv.Itoa(dig)
				break
			}
		}

		if dig, err := strconv.Atoi(rowValue); err == nil {
			count += dig
		}
	}

	return count
}
