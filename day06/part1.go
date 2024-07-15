package day06

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	record   int
	winnable int
}

const VELOCITY = 1

func Part1() int {
	input, err := os.ReadFile("day06/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var races = parseInput(input)
	for i := range races {
		simulateRace(&races[i])
	}
	// fmt.Printf("%#v", races)

	return getResult(&races)
}

func parseInput(input []byte) []Race {
	rows := strings.Split(string(input), "\n")
	var times = strings.Fields(rows[0])
	var distances = strings.Fields(rows[1])
	var races []Race
	for i := 1; i < len(times); i++ {
		var time, _ = strconv.Atoi(times[i])
		var record, _ = strconv.Atoi(distances[i])
		races = append(races, Race{time: time, record: record})
	}
	return races
}

func simulateRace(race *Race) {
	for timePressing := 1; timePressing < race.time; timePressing++ {
		var distance = (race.time - timePressing) * (VELOCITY * timePressing)
		if distance > race.record {
			race.winnable++
		}
	}
}

func getResult(races *[]Race) int {
	var result = 1
	for _, race := range *races {
		result *= int(math.Max(0, float64(race.winnable)))
	}
	return result
}
