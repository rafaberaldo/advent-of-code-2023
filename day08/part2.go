package day08

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Part2() int {
	var start = time.Now()
	input, err := os.ReadFile("day08/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var nodeMap, directions = parseInput(input)
	var result = calculateStepsLCM(nodeMap, directions)

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)

	return result
}

func calculateStepsLCM(nodeMap map[string]Node, directions []string) int {
	var steps []int
	for key, value := range nodeMap {
		if strings.HasSuffix(key, "A") {
			steps = append(steps, calculateSteps2(value, nodeMap, directions))
		}
	}

	return lcm(steps)
}

func calculateSteps2(init Node, nodeMap map[string]Node, directions []string) int {
	var steps = 0
	var index = -1
	var walk func(node Node)
	walk = func(node Node) {
		index++
		if index == len(directions) {
			index = 0
		}

		if strings.HasSuffix(node.ident, "Z") {
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

	walk(init)

	return steps
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

// This is gonna take hours... or days
func calculateStepsBruteForce(nodeMap map[string]Node, directions []string) int {
	var getNodesFromDirection = func(nodes []Node, direction string) []Node {
		var result []Node
		for _, node := range nodes {
			if direction == "L" {
				result = append(result, nodeMap[node.left])
			} else if direction == "R" {
				result = append(result, nodeMap[node.right])
			}
		}
		return result
	}

	var nodes []Node
	for key, value := range nodeMap {
		if strings.HasSuffix(key, "A") {
			nodes = append(nodes, value)
			fmt.Printf("%#v\n", value)
		}
	}

	var steps = 0
	for i := 0; i <= len(directions); i++ {
		if i == len(directions) {
			i = 0
		}

		if hasEnded(nodes) {
			break
		}

		steps++
		if steps%1_000_000 == 0 {
			fmt.Println(steps)
		}

		nodes = getNodesFromDirection(nodes, directions[i])
	}

	return steps
}

func hasEnded(nodes []Node) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node.ident, "Z") {
			return false
		}
	}
	return true
}
