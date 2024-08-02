/*
	SVG graph generated with Graphviz with cmd `neato -Tsvg input.dot > graph.svg`.

	graph {
		jqt -- {rhn xhk nvd}
		rsh -- {frs pzl lsr}
		xhk -- {hfx}
		cmg -- {qnr nvd lhk bvb}
		rhn -- {xhk bvb hfx}
		bvb -- {xhk hfx}
		pzl -- {lsr hfx nvd}
		qnr -- {nvd}
		ntq -- {jqt hfx bvb xhk}
		nvd -- {lhk}
		lsr -- {lhk}
		rzs -- {qnr cmg lsr rsh}
		frs -- {qnr lhk lsr}
	}
*/

package day25

import (
	"aoc2023/assert"
	_slices "aoc2023/lib"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func Part1() int {
	started := time.Now()
	input, err := os.ReadFile("day25/data.txt")
	assert.Assert(err == nil, "error opening the file: %v", err)

	compMap, _ := parseInput(input)
	keys := mapKeys(compMap)
	tempMap := removeLink(compMap, []string{"tmb", "gpj"}, []string{"rhh", "mtc"}, []string{"njn", "xtx"})
	result, _ := totalLinks(tempMap, keys)

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) (map[string][]string, [][]string) {
	lines := strings.Split(string(input), "\n")
	compMap := make(map[string][]string)
	var pairs [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(strings.Replace(line, ":", "", 1))
		compMap[fields[0]] = fields[1:]
		for _, v := range fields[1:] {
			pairs = append(pairs, []string{fields[0], v})
		}
	}

	for k := range compMap {
		for _, comp := range compMap[k] {
			if slices.Index(compMap[comp], k) == -1 {
				compMap[comp] = append(compMap[comp], k)
			}
		}
	}

	return compMap, pairs
}

func totalLinks(compMap map[string][]string, keys []string) (int, bool) {
	found := 0
	for _, k := range keys {
		count := countLinks(compMap, k)
		if count == 1 || count == len(keys) {
			return 0, false
		} else {
			found = count
			break
		}
	}
	delta := len(keys) - found
	result := delta * found
	return result, true
}

func countLinks(compMap map[string][]string, start string) int {
	seen := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		comp := queue[0]
		queue = queue[1:]

		if seen[comp] {
			continue
		}

		seen[comp] = true

		for _, link := range compMap[comp] {
			if seen[link] {
				continue
			}

			queue = append(queue, link)
		}
	}
	return len(seen)
}

func mapKeys[K comparable, T any](m map[K]T) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func removeLink(compMap map[string][]string, p1, p2, p3 []string) map[string][]string {
	result := make(map[string][]string)
	var verifyAndAppend = func(k string, pair []string) bool {
		if k == pair[0] {
			result[k] = _slices.Filter(compMap[k], func(_ int, str string) bool {
				return str != pair[1]
			})
			return true
		}
		if k == pair[1] {
			result[k] = _slices.Filter(compMap[k], func(_ int, str string) bool {
				return str != pair[0]
			})
			return true
		}
		return false
	}
	for k := range compMap {
		if verifyAndAppend(k, p1) {
			continue
		}
		if verifyAndAppend(k, p2) {
			continue
		}
		if verifyAndAppend(k, p3) {
			continue
		}
		result[k] = compMap[k]
	}
	return result
}
