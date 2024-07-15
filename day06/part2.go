package day06

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2() int {
	input, err := os.ReadFile("day06/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var races = parseInput2(input)
	for i := range races {
		simulateRace(&races[i])
	}
	// fmt.Printf("%#v", races)

	return getResult(&races)
}

func parseInput2(input []byte) []Race {
	rows := strings.Split(string(input), "\n")
	var times = strings.ReplaceAll(strings.Split(rows[0], ":")[1], " ", "")
	var distances = strings.ReplaceAll(strings.Split(rows[1], ":")[1], " ", "")
	var races []Race
	var time, _ = strconv.Atoi(times)
	var record, _ = strconv.Atoi(distances)
	races = append(races, Race{time: time, record: record})
	return races
}
