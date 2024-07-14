package day03

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PartCoord struct {
	row       int
	spaces    []int
	value     string
	neighbors []string
}

func Part2() int {
	input, err := os.ReadFile("day03/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(input), "\n")

	var count = 0
	var partsCoord = []PartCoord{}
	var reNum = regexp.MustCompile(`\d+|[^\.]`)

	var intRange = func(min int, max int) []int {
		var result = []int{}
		for i := min; i <= max; i++ {
			result = append(result, i)
		}
		return result
	}

	for row := 0; row < len(rows); row++ {
		var iMatches = reNum.FindAllStringIndex(rows[row], -1)
		var matches = reNum.FindAllString(rows[row], -1)

		for i := 0; i < len(iMatches); i++ {
			var part = PartCoord{row, intRange(iMatches[i][0], iMatches[i][1]), matches[i], nil}
			partsCoord = append(partsCoord, part)
		}
	}

	var comparePartsCoord = func(a PartCoord, b PartCoord) bool {
		return a.row == b.row && a.value == b.value && a.spaces[0] == b.spaces[0]
	}

	var findNeighbors = func(part *PartCoord) {
		for _, p := range partsCoord {
			if part.row-p.row > 1 || part.row-p.row < -1 || comparePartsCoord(*part, p) {
				continue
			}

		outer:
			for i := 0; i < len(part.spaces); i++ {
				for j := 0; j < len(p.spaces); j++ {
					if part.spaces[i] == p.spaces[j] {
						part.neighbors = append(part.neighbors, p.value)
						break outer
					}
				}
			}
		}
	}

	for i := 0; i < len(partsCoord); i++ {
		findNeighbors(&partsCoord[i])

		// fmt.Printf("%+v\n", partsCoord[i])

		if len(partsCoord[i].neighbors) == 0 {
			continue
		}

		if partsCoord[i].value != "*" {
			continue
		}

		if len(partsCoord[i].neighbors) != 2 {
			continue
		}

		if value1, err := strconv.Atoi(partsCoord[i].neighbors[0]); err == nil {
			if value2, err := strconv.Atoi(partsCoord[i].neighbors[1]); err == nil {
				count += value1 * value2
			}
		}
	}

	return count
}
