package day05

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() int {
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

	var fillMap = func(
		m *map[int]int,
		valuesStr []string,
		needed *[]int,
	) {
		var values = strToInt(valuesStr)
		var srcStart = values[1]
		var destStart = values[0]
		var length = values[2]

		slices.Sort(*needed)

		if slices.Max(*needed) < srcStart {
			return
		}

		for i := 0; i < length; i++ {
			if srcStart+i == destStart+i {
				continue
			}

			var _, ok = slices.BinarySearch(*needed, srcStart+i)
			if !ok {
				continue
			}

			*needed = append(*needed, destStart+i)
			slices.Sort(*needed)
			(*m)[srcStart+i] = destStart + i
		}
	}

	var seedToSoil = map[int]int{}
	var soilToFertilizer = map[int]int{}
	var fertilizerToWater = map[int]int{}
	var waterToLight = map[int]int{}
	var lightToTemperature = map[int]int{}
	var temperatureToHumidity = map[int]int{}
	var humidityToLocation = map[int]int{}
	var seeds []int
	var needed []int

	for i := 0; i < len(rows); i++ {
		// fmt.Printf("%#v\n", rows[i])

		if rows[i] == "" {
			ident = ""
			continue
		}

		if strings.Contains(rows[i], ":") {
			ident = strings.Split(rows[i], ":")[0]
			if ident == "seeds" {
				var value = strings.Split(rows[i], ":")[1]
				seeds = strToInt(strings.Split(value, " "))
				needed = seeds
			}
			continue
		}

		if ident == "seed-to-soil map" {
			fmt.Println("-----seed-to-soil map-----")
			var values = strings.Split(rows[i], " ")
			fillMap(&seedToSoil, values, &needed)
		}

		if ident == "soil-to-fertilizer map" {
			fmt.Println("-----soil-to-fertilizer map-----")
			var values = strings.Split(rows[i], " ")
			fillMap(&soilToFertilizer, values, &needed)
		}

		if ident == "fertilizer-to-water map" {
			fmt.Println("-----fertilizer-to-water map-----")
			var values = strings.Split(rows[i], " ")
			fillMap(&fertilizerToWater, values, &needed)
		}

		if ident == "water-to-light map" {
			fmt.Println("-----water-to-light map-----")
			var values = strings.Split(rows[i], " ")
			fillMap(&waterToLight, values, &needed)
		}

		if ident == "light-to-temperature map" {
			fmt.Println("-----light-to-temperature map-----")
			var values = strings.Split(rows[i], " ")
			fillMap(&lightToTemperature, values, &needed)
		}

		if ident == "temperature-to-humidity map" {
			fmt.Println("-----temperature-to-humidity map-----")
			var values = strings.Split(rows[i], " ")
			fillMap(&temperatureToHumidity, values, &needed)
		}

		if ident == "humidity-to-location map" {
			fmt.Println("-----humidity-to-location map-----")
			var values = strings.Split(rows[i], " ")
			fillMap(&humidityToLocation, values, &needed)
		}
	}

	var getValue = func(m *map[int]int, key int) int {
		if value, ok := (*m)[key]; ok {
			return value
		}
		return key
	}

	var minFinalLocation = 0
	for _, seed := range seeds {
		var soil = getValue(&seedToSoil, seed)
		var fertilizer = getValue(&soilToFertilizer, soil)
		var water = getValue(&fertilizerToWater, fertilizer)
		var light = getValue(&waterToLight, water)
		var temperature = getValue(&lightToTemperature, light)
		var humidity = getValue(&temperatureToHumidity, temperature)
		var location = getValue(&humidityToLocation, humidity)

		if minFinalLocation == 0 || location < minFinalLocation {
			minFinalLocation = location
		}
	}

	return minFinalLocation
}
