package day02

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() int {
	input, err := os.ReadFile("day02/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(input), "\n")

	var count = 0
	for row := 0; row < len(rows); row++ {
		if rows[row] == "" {
			continue
		}
		count += parseGameP1(rows[row])
	}
	return count
}

func parseGameP1(str string) int {
	var maxRed = 12
	var maxGreen = 13
	var maxBlue = 14

	var id, _ = strconv.Atoi(strings.Split(strings.Split(str, ":")[0], " ")[1])

	var rounds = strings.Split(strings.Split(str, ":")[1], ";")
	var isValid = true

	var isRoundValid = func(str string, max int, regex string) bool {
		var re = regexp.MustCompile(regex)
		var matches = re.FindStringSubmatch(str)
		if len(matches) > 1 {
			var value, _ = strconv.Atoi(matches[1])
			return value <= max
		}
		return true
	}

	for i := 0; i < len(rounds); i++ {
		isValid = isRoundValid(rounds[i], maxRed, `(\d+) red`)
		if !isValid {
			break
		}
		isValid = isRoundValid(rounds[i], maxGreen, `(\d+) green`)
		if !isValid {
			break
		}
		isValid = isRoundValid(rounds[i], maxBlue, `(\d+) blue`)
		if !isValid {
			break
		}
	}

	if isValid {
		return id
	} else {
		return 0
	}
}
