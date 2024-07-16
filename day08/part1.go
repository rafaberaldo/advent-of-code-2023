package day08

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	ident string
	left  string
	right string
}

const (
	START = "AAA"
	END   = "ZZZ"
)

func Part1() int {
	input, err := os.ReadFile("day08/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var nodeMap, directions = parseInput(input)

	return calculateSteps(nodeMap, directions)
}

func parseInput(input []byte) (map[string]Node, []string) {
	var lines = strings.Split(string(input), "\n")
	var nodeMap = make(map[string]Node)
	var directions []string
	var regex = regexp.MustCompile(`\w{3}`)

	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			directions = strToSlice(line)
			continue
		}

		var matches = regex.FindAllString(line, -1)
		nodeMap[matches[0]] = Node{matches[0], matches[1], matches[2]}
	}

	return nodeMap, directions
}

func strToSlice(str string) []string {
	var result []string
	for i := range len(str) {
		result = append(result, string(str[i]))
	}
	return result
}

func calculateSteps(nodeMap map[string]Node, directions []string) int {
	var steps = 0
	var index = -1
	var walk func(node Node)
	walk = func(node Node) {
		index++
		if index == len(directions) {
			index = 0
		}

		if node.ident == END {
			return
		}

		steps++

		if directions[index] == "L" {
			walk(nodeMap[node.left])
			return
		}

		if directions[index] == "R" {
			walk(nodeMap[node.right])
			return
		}
	}

	var init = nodeMap[START]
	walk(init)

	return steps
}
