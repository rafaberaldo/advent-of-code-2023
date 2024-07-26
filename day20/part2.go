package day20

import (
	"aoc2023/assert"
	"fmt"
	"os"
	"slices"
	"time"
)

func Part2() int {
	var start = time.Now()
	var input, err = os.ReadFile("day20/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	var broadcaster, modulesMap = parseInput(input)
	var rxConn, rxConnCount = findRxConnector(modulesMap)
	assert.Assert(rxConnCount > 0, "could not find rx connectors!")
	var rxConnsClicks []int

	var clicks = 0
	for rxConnCount > len(rxConnsClicks) {
		clicks++
		if broadcast2(broadcaster, &modulesMap, rxConn) {
			rxConnsClicks = append(rxConnsClicks, clicks)
		}
	}

	var result = lcm(rxConnsClicks)

	var elapsed = time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func broadcast2(broadcaster []string, modulesMap *map[string]Module, rxConn string) bool {
	var queue []Broadcast
	for _, receiver := range broadcaster {
		(*modulesMap)[receiver].receive(LOW, "broadcaster")
		var pulse = (*modulesMap)[receiver].pulse()
		queue = append(queue, Broadcast{receiver, pulse})
	}

	var found = false
	for len(queue) > 0 {
		var bc = queue[0]
		queue = queue[1:]

		if bc.pulse == "" {
			continue
		}

		var receivers = (*modulesMap)[bc.sender].receivers()
		for _, receiver := range receivers {

			if (*modulesMap)[receiver] == nil {
				continue
			}
			(*modulesMap)[receiver].receive(bc.pulse, bc.sender)

			if receiver == rxConn && bc.pulse == HIGH {
				found = true
			}

			var pulse = (*modulesMap)[receiver].pulse()
			queue = append(queue, Broadcast{receiver, pulse})
		}
	}

	return found
}

// return the Conjunction ident/name and its input count
func findRxConnector(modulesMap map[string]Module) (string, int) {
	for k, v := range modulesMap {
		if slices.Contains(v.receivers(), "rx") {
			return k, modulesMap[k].inputCount()
		}
	}
	return "", 0
}

// greatest commom divisor
func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// least common multiple
func lcm(n []int) int {
	var result = n[0]
	for i := 1; i < len(n); i++ {
		result = (n[i] * result) / gcd(n[i], result)
	}
	return result
}
