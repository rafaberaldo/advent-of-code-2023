package day02

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part2() int {
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
		count += parseGameP2(rows[row])
	}
	return count
}

func parseGameP2(str string) int {
	var getMax = func(str string, regex string) int {
		var re = regexp.MustCompile(regex)
		var matches = re.FindAllStringSubmatch(str, -1)
		var max = 0
		for i := 0; i < len(matches); i++ {
			for j := 1; j < len(matches[i]); j++ {
				var value, _ = strconv.Atoi(matches[i][j])
				if value > max {
					max = value
				}
			}
		}
		return max
	}

	var game = strings.Split(str, ":")[1]
	var power = 1
	power *= getMax(game, `(\d+) red`)
	power *= getMax(game, `(\d+) green`)
	power *= getMax(game, `(\d+) blue`)
	return power
}
