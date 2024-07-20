package day12

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day12/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var puzzle = parseInput2(input)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	var result = 0
	for _, p := range puzzle {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mutex.Unlock()
			var memo = make(map[string]int)
			var count = findPattern(p, &memo)
			mutex.Lock()
			result += count
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput2(input []byte) []Puzzle {
	var lines = strings.Split(string(input), "\n")
	var puzzle []Puzzle
	for _, line := range lines {
		if line == "" {
			continue
		}
		var fields = strings.Fields(line)
		var condition = toIntSlice(strings.Split(fields[1], ","))
		var pl = fields[0]
		for range 4 {
			condition = append(condition, toIntSlice(strings.Split(fields[1], ","))...)
			pl += "?" + fields[0]
		}
		puzzle = append(puzzle, Puzzle{pl, condition})
	}
	return puzzle
}
