package day15

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day15/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var hashmapArr = parseInput2(input)
	var result = calculateResult(hashmapArr)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

type Hashmap struct {
	key   string
	value int
}

func parseInput2(input []byte) [][]Hashmap {
	var lines = strings.Split(string(input), "\n")
	var instructions []string
	for _, line := range lines {
		if line == "" {
			continue
		}
		instructions = append(instructions, strings.Split(line, ",")...)
	}

	var hashmapArr = make([][]Hashmap, 256)
	for _, instruction := range instructions {
		if strings.Contains(instruction, "=") {
			var split = strings.Split(instruction, "=")
			var key = split[0]
			var value, _ = strconv.Atoi(split[1])
			var idx = hash(key)
			var hashmap = Hashmap{key, value}

			if keyIdx := findInHashmap(hashmapArr[idx], hashmap.key); keyIdx > -1 {
				hashmapArr[idx][keyIdx] = hashmap
			} else {
				hashmapArr[idx] = append(hashmapArr[idx], hashmap)
			}
		}

		if strings.Contains(instruction, "-") {
			var key = strings.Split(instruction, "-")[0]
			var idx = hash(key)

			if keyIdx := findInHashmap(hashmapArr[idx], key); keyIdx > -1 {
				hashmapArr[idx] = slices.Delete(hashmapArr[idx], keyIdx, keyIdx+1)
			}
		}
	}

	return hashmapArr
}

func findInHashmap(hashmaps []Hashmap, key string) int {
	for i, h := range hashmaps {
		if h.key == key {
			return i
		}
	}
	return -1
}

func calculateResult(hashmapArr [][]Hashmap) int {
	var result int
	for i, hashmaps := range hashmapArr {
		for j, hashmap := range hashmaps {
			result += (i + 1) * (j + 1) * hashmap.value
		}
	}
	return result
}
