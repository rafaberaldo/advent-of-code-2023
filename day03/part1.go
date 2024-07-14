package day03

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	row   int
	index int
}

func Part1() int {
	input, err := os.ReadFile("day03/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(input), "\n")

	var count = 0
	var intersectMap = make(map[Coord]bool)
	var reNum = regexp.MustCompile(`\d+|[^\.]`)

	for row := 0; row < len(rows); row++ {
		var iMatches = reNum.FindAllStringIndex(rows[row], -1)
		// fmt.Println(iMatches)

		var horizontalIntersect = func(a []int, b []int) bool {
			return a[0] == b[len(b)-1] || a[len(a)-1] == b[0]
		}

		var verticalIntersect = func(a []int, b []int) bool {
			return a[0] >= b[0] && a[0] <= b[len(b)-1] ||
				a[len(a)-1] >= b[0] && a[len(a)-1] <= b[len(b)-1] ||
				a[0] <= b[0] && a[len(a)-1] >= b[0]
		}

		for i := 0; i < len(iMatches); i++ {
			if i > 0 && horizontalIntersect(iMatches[i], iMatches[i-1]) {
				intersectMap[Coord{row, i}] = true
				intersectMap[Coord{row, i - 1}] = true
			}

			if row > 0 {
				var iMatchesPrev = reNum.FindAllStringIndex(rows[row-1], -1)
				for j := 0; j < len(iMatchesPrev); j++ {
					if verticalIntersect(iMatches[i], iMatchesPrev[j]) {
						intersectMap[Coord{row, i}] = true
						intersectMap[Coord{row - 1, j}] = true
					}
				}
			}
		}
	}

	// fmt.Println(intersectMap)

	for coord := range intersectMap {
		var matches = reNum.FindAllString(rows[coord.row], -1)
		// fmt.Println(coord.row, matches[coord.index])
		if value, err := strconv.Atoi(matches[coord.index]); err == nil {
			count += value
		}
	}

	return count
}
