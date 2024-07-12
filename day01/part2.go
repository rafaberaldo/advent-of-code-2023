package day01

import (
	"regexp"
	"strconv"
)

func Part2() int {

	count := 0

	input := GetInput()
	for row := 0; row < len(input); row++ {
		value, err := computeValue(input[row])
		if value < 10 || value > 99 {
			panic("??")
		}
		if err != nil {
			panic(err)
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
