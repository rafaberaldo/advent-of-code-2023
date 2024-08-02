/*
	SVG graph generated with Graphviz with cmd `neato -Tsvg input.dot > graph.svg`.

	Since the description says we have to disconnect at least half the components,
	we know there are two clusters with three links:

	  * * *       * * *
  * * * * * - * * * * *
  * * * * * - * * * * *
  * * * * * - * * * * *
    * * *       * * *

	Theoretically, running BFS on every pair (as start/end) and counting all seen links,
	the ones we're looking for are the ones most seen.
*/

package day25

import (
	"aoc2023/assert"
	_slices "aoc2023/lib"
	"cmp"
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
	visitedCount := make(map[string]int)
	// we don't actually need to run on all of them since the input is huge
	// use 15 for the test data!
	const MAX_KEYS = 200
	for i, k1 := range keys[:MAX_KEYS] {
		for _, k2 := range keys[:i] {
			search(compMap, k1, k2, visitedCount)
		}
	}

	// Brute force ordererd by most visited links
	// Time may vary between 6-20 secs
	result := 0
	mostVisited := getMostVisited(visitedCount)
outer:
	for i, l1 := range mostVisited {
		for j, l2 := range mostVisited[:i] {
			for _, l3 := range mostVisited[:j] {
				tempMap := removeLink(compMap, l1.link, l2.link, l3.link)
				if totals, ok := totalLinks(tempMap, keys); ok {
					result = totals
					break outer
				}
			}
		}
	}

	elapsed := time.Since(started)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func parseInput(input []byte) (map[string][]string, [][]string) {
	lines := strings.Split(string(input), "\n")
	compMap := make(map[string][]string)
	var links [][]string
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(strings.Replace(line, ":", "", 1))
		compMap[fields[0]] = fields[1:]
		for _, v := range fields[1:] {
			links = append(links, []string{fields[0], v})
		}
	}

	for k := range compMap {
		for _, comp := range compMap[k] {
			if slices.Index(compMap[comp], k) == -1 {
				compMap[comp] = append(compMap[comp], k)
			}
		}
	}

	return compMap, links
}

type LinkCount struct {
	link  []string
	count int
}

func getMostVisited(visited map[string]int) []LinkCount {
	result := make([]LinkCount, 0, len(visited))
	for k, v := range visited {
		result = append(result, LinkCount{strings.Split(k, "-"), v})
	}
	slices.SortFunc(result, func(a, b LinkCount) int {
		return cmp.Compare(b.count, a.count)
	})
	return result
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

func search(compMap map[string][]string, start, end string, visited map[string]int) {
	seen := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		comp := queue[0]
		queue = queue[1:]

		if seen[comp] {
			continue
		}

		seen[comp] = true

		if comp == end {
			return
		}

		for _, edge := range compMap[comp] {
			if seen[edge] {
				continue
			}

			if _, ok := visited[comp+"-"+edge]; ok {
				visited[comp+"-"+edge]++
			} else {
				visited[edge+"-"+comp]++
			}

			queue = append(queue, edge)
		}
	}
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
