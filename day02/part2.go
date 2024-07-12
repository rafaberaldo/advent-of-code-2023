package day02

import (
	"regexp"
	"strconv"
	"strings"
)

func Part2() int {
	var count = 0
	var input = GetInput()
	for row := 0; row < len(input); row++ {
		count += parseGameP2(input[row])
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
