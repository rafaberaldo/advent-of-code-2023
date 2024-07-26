package day20

import (
	"aoc2023/assert"
	"fmt"
	"os"
	"strings"
	"time"
)

func Part1() int {
	var start = time.Now()
	var input, err = os.ReadFile("day20/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	var broadcaster, modulesMap = parseInput(input)
	var lowCount = 0
	var highCount = 0
	for range 1000 {
		var lc, hc = broadcast(broadcaster, &modulesMap)
		lowCount += lc
		highCount += hc
	}

	var result = lowCount * highCount

	var elapsed = time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) ([]string, map[string]Module) {
	var lines = strings.Split(string(input), "\n")
	var modulesMap = make(map[string]Module)
	var broadcaster []string

	for _, line := range lines {
		if line == "" {
			continue
		}

		var fields = strings.Split(line, " -> ")
		var outputs = strings.Split(fields[1], ", ")

		if fields[0] == "broadcaster" {
			broadcaster = outputs
			continue
		}

		var moduleName = fields[0][1:]

		if fields[0][0] == '%' {
			var module = &FlipFlop{ident: moduleName, outputs: outputs}
			modulesMap[moduleName] = module
			continue
		}

		if fields[0][0] == '&' {
			var module = &Conjunction{ident: moduleName, outputs: outputs}
			modulesMap[moduleName] = module
		}
	}

	// Gotta add conjunction inputs
	for _, line := range lines {
		if line == "" {
			continue
		}

		var fields = strings.Split(line, " -> ")
		var outputs = strings.Split(fields[1], ", ")
		var moduleName = fields[0][1:]

		for _, ident := range outputs {
			if modulesMap[ident] == nil {
				continue
			}
			modulesMap[ident].addInput(moduleName)
		}
	}

	return broadcaster, modulesMap
}

type Broadcast struct {
	sender string
	pulse  Pulse
}

func broadcast(broadcaster []string, modulesMap *map[string]Module) (int, int) {
	var queue []Broadcast
	var lowCount = 1
	var highCount = 0

	for _, receiver := range broadcaster {
		lowCount++
		(*modulesMap)[receiver].receive(LOW, "broadcaster")
		var pulse = (*modulesMap)[receiver].pulse()
		queue = append(queue, Broadcast{receiver, pulse})
	}

	for len(queue) > 0 {
		var bc = queue[0]
		queue = queue[1:]

		if bc.pulse == "" {
			continue
		}

		var receivers = (*modulesMap)[bc.sender].receivers()
		for _, receiver := range receivers {
			switch bc.pulse {
			case HIGH:
				highCount++
			case LOW:
				lowCount++
			}

			if (*modulesMap)[receiver] == nil {
				continue
			}
			(*modulesMap)[receiver].receive(bc.pulse, bc.sender)

			var pulse = (*modulesMap)[receiver].pulse()
			queue = append(queue, Broadcast{receiver, pulse})
		}
	}

	return lowCount, highCount
}
