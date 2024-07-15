package day07

import (
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HandType string

var handTypeMap = map[HandType]int{
	"FiveKind":  6,
	"FourKind":  5,
	"FullHouse": 4,
	"ThreeKind": 3,
	"TwoPair":   2,
	"OnePair":   1,
	"None":      0,
}

type Hand struct {
	cards       string
	bid         int
	cardsValues []int
	handType    HandType
	value       int
}

func Part1() int {
	input, err := os.ReadFile("day07/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var hands = parseInput(input)

	var result = 0
	for i, hand := range hands {
		// fmt.Printf("%#v\n", hand)
		var rank = i + 1
		result += hand.bid * rank
	}

	return result
}

func parseInput(input []byte) []Hand {
	var lines = strings.Split(string(input), "\n")
	var hands []Hand
	for _, line := range lines {
		if line == "" {
			continue
		}
		var fields = strings.Fields(line)
		var cards = fields[0]
		var bid, _ = strconv.Atoi(fields[1])

		var cardsValues = calculateCardsValue(cards)
		var handType = calculateHandType(cards)
		hands = append(hands, Hand{cards, bid, cardsValues, handType, handTypeMap[handType]})
	}
	// Sort by hand type then cards
	slices.SortFunc(hands, func(a Hand, b Hand) int {
		if a.value < b.value {
			return -1
		}
		if a.value > b.value {
			return +1
		}

		for i := 0; i < 5; i++ {
			if a.cardsValues[i] < b.cardsValues[i] {
				return -1
			}
			if a.cardsValues[i] > b.cardsValues[i] {
				return +1
			}
		}

		return 0
	})

	return hands
}

func calculateCardsValue(hand string) []int {
	var cardValueMap = map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}
	var result []int
	for i := range len(hand) {
		var card = string(hand[i])
		if v, ok := cardValueMap[card]; ok {
			result = append(result, v)
		} else {
			log.Fatal("Card should be in map!")
		}
	}
	return result
}

func calculateHandType(hand string) HandType {
	var cards []string
	for i := range len(hand) {
		cards = append(cards, string(hand[i]))
	}

	var cardQtyMap = make(map[string]int)
	for i := range len(cards) {
		cardQtyMap[cards[i]] += 1
	}

	// Organize all cards quantities sorted.
	// e.g. [2, 2, 1] is a two-pair; [3, 2] is a full-house
	var cardQty []int
	for _, v := range cardQtyMap {
		cardQty = append(cardQty, v)
	}
	slices.SortFunc(cardQty, func(a int, b int) int {
		return cmp.Compare(b, a)
	})

	var handType HandType
	switch cardQty[0] {
	case 5:
		handType = "FiveKind" // [5]
	case 4:
		handType = "FourKind" // [4, 1]
	case 3:
		if cardQty[1] == 2 {
			handType = "FullHouse" // [3, 2]
		} else {
			handType = "ThreeKind" // [3, 1, 1]
		}
	case 2:
		if cardQty[1] == 2 {
			handType = "TwoPair" // [2, 2, 1]
		} else {
			handType = "OnePair" // [2, 1, 1, 1]
		}
	default:
		handType = "None" // [1, 1, 1, 1, 1]
	}

	return handType
}
