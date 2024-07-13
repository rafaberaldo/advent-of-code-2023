package day04

import (
	"strconv"
	"strings"
)

func Part2() int {
	var input = GetInput()
	var count = 0

	var totalPerCard = make([]int, len(input))

	for row := 0; row < len(input); row++ {
		var numbers = strings.Split(strings.Split(input[row], ":")[1], "|")
		var realNumbers = strings.Split(numbers[0], " ")
		var myNumbers = strings.Split(numbers[1], " ")

		var winningNumbers = []int{}
		for i := 0; i < len(realNumbers); i++ {
			for j := 0; j < len(myNumbers); j++ {
				if realNumbers[i] == myNumbers[j] {
					if value, err := strconv.Atoi(myNumbers[j]); err == nil {
						winningNumbers = append(winningNumbers, value)
					}
				}
			}
		}

		for range totalPerCard[row] + 1 {
			for i := 1; i <= len(winningNumbers); i++ {
				// fmt.Println("copying card ", row+i+1)
				totalPerCard[row+i]++
			}
		}

		count += totalPerCard[row] + 1
	}

	return count
}
