package day05

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	srcMin  int
	srcMax  int
	destMin int
	destMax int
}

type Seed struct {
	min int
	max int
}

func Part2() int {
	input, err := os.ReadFile("day05/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(input), "\n")
	var ident string

	var strToInt = func(str []string) []int {
		var result []int
		for i := 0; i < len(str); i++ {
			if v, err := strconv.Atoi(str[i]); err == nil {
				result = append(result, v)
			}
		}
		return result
	}

	var seedToSoil []Range
	var soilToFertilizer []Range
	var fertilizerToWater []Range
	var waterToLight []Range
	var lightToTemperature []Range
	var temperatureToHumidity []Range
	var humidityToLocation []Range
	var seeds []Seed

	var createRange = func(values []int) Range {
		var length = values[2]
		return Range{
			srcMin:  values[1],
			srcMax:  values[1] + length - 1,
			destMin: values[0],
			destMax: values[0] + length - 1,
		}
	}

	for i := 0; i < len(rows); i++ {
		// fmt.Printf("%#v\n", rows[i])

		if rows[i] == "" {
			ident = ""
			continue
		}

		if strings.Contains(rows[i], ":") {
			ident = strings.Split(rows[i], ":")[0]
			if ident == "seeds" {
				var valuesStr = strings.Split(rows[i], ":")[1]
				var values = strToInt(strings.Split(valuesStr, " "))
				for j := 0; j < len(values); j = j + 2 {
					seeds = append(seeds, Seed{min: values[j], max: values[j] + values[j+1] - 1})
				}
			}
			continue
		}

		if ident == "seed-to-soil map" {
			var values = strings.Split(rows[i], " ")
			seedToSoil = append(seedToSoil, createRange(strToInt(values)))
		}

		if ident == "soil-to-fertilizer map" {
			var values = strings.Split(rows[i], " ")
			soilToFertilizer = append(soilToFertilizer, createRange(strToInt(values)))
		}

		if ident == "fertilizer-to-water map" {
			var values = strings.Split(rows[i], " ")
			fertilizerToWater = append(fertilizerToWater, createRange(strToInt(values)))
		}

		if ident == "water-to-light map" {
			var values = strings.Split(rows[i], " ")
			waterToLight = append(waterToLight, createRange(strToInt(values)))
		}

		if ident == "light-to-temperature map" {
			var values = strings.Split(rows[i], " ")
			lightToTemperature = append(lightToTemperature, createRange(strToInt(values)))
		}

		if ident == "temperature-to-humidity map" {
			var values = strings.Split(rows[i], " ")
			temperatureToHumidity = append(temperatureToHumidity, createRange(strToInt(values)))
		}

		if ident == "humidity-to-location map" {
			var values = strings.Split(rows[i], " ")
			humidityToLocation = append(humidityToLocation, createRange(strToInt(values)))
		}
	}
	fmt.Println("Parsing done.")

	var getValue = func(ran []Range, value int) int {
		for i := 0; i < len(ran); i++ {
			if ran[i].destMin <= value && value <= ran[i].destMax {
				var offset = value - ran[i].destMin
				return ran[i].srcMin + offset
			}
		}
		return value
	}

	var hasSeed = func(s []Seed, value int) bool {
		for i := 0; i < len(s); i++ {
			if s[i].min <= value && value <= s[i].max {
				return true
			}
		}
		return false
	}

	var findMinFinalLocation = func(min int, max int, c chan int) {
		fmt.Printf("Starting goroutine (%#v)\n", min)
		for i := min; i <= max; i++ {
			// fmt.Printf("Testing location %#v\n", i)
			var humidity = getValue(humidityToLocation, i)
			var temperature = getValue(temperatureToHumidity, humidity)
			var light = getValue(lightToTemperature, temperature)
			var water = getValue(waterToLight, light)
			var fertilizer = getValue(fertilizerToWater, water)
			var soil = getValue(soilToFertilizer, fertilizer)
			var seed = getValue(seedToSoil, soil)
			if hasSeed(seeds, seed) {
				fmt.Printf("Stopping goroutne (%#v)\n", min)
				c <- i
				break
			}
		}
	}

	var found []int
	var c = make(chan int)
	var maxInt = math.MaxInt32 / 4
	// Not sure why it's taking longer when using a for loop
	go findMinFinalLocation(0, maxInt, c)
	go findMinFinalLocation(maxInt, maxInt*2, c)
	go findMinFinalLocation(maxInt*2, maxInt*3, c)
	go findMinFinalLocation(maxInt*3, maxInt*4, c)
	found = append(found, <-c)
	found = append(found, <-c)
	found = append(found, <-c)
	found = append(found, <-c)

	return slices.Min(found)
}
