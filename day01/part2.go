package day01

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part2() int {
	input, err := os.ReadFile("day01/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(input), "\n")
	count := 0

	for row := 0; row < len(rows); row++ {
		if rows[row] == "" {
			continue
		}

		value, err := computeValue(rows[row])
		if value < 10 || value > 99 {
			log.Fatal("??")
		}
		if err != nil {
			log.Fatal(err)
		}
		count += value
	}

	return count
}

func computeValue(str string) (int, error) {
	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	re, _ := regexp.Compile(`\d|one|two|three|four|five|six|seven|eight|nine`)

	rowValue := ""

	value := re.FindString(str)
	if len(value) > 1 {
		rowValue = numbers[value]
	} else {
		rowValue = value
	}

	for i := len(str) - 1; i >= 0; i-- {
		value := re.FindString(string(str[i:]))
		if value == "" {
			continue
		}

		if len(value) > 1 {
			rowValue += numbers[value]
			break
		} else {
			rowValue += value
			break
		}
	}

	return strconv.Atoi(rowValue)
}
