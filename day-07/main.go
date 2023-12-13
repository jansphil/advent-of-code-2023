package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/jansphil/advent-of-code-2023/utils"
)

// five of a kind  => 6
// four of a kind  => 5
// full house      => 4
// three of a kind => 3
// two pair        => 2
// one pair        => 1
// high card       => 0
func scoreHand(hand []string) int {
	cardCounts := make(map[string]int)
	for _, card := range hand {
		cardCounts[card]++
	}

	switch len(cardCounts) {
	case 1: // five of a kind
		return 6
	case 2: // four of a kind / full house
		for _, v := range cardCounts {
			if v == 4 || v == 1 {
				// four of a kind
				return 5
			} else {
				// full house
				return 4
			}
		}
	case 3: // three of a kind / two pair
		for _, v := range cardCounts {
			if v == 3 {
				// three of a kind
				return 3
			} else if v == 2 {
				// two pair
				return 2
			}
		}
	case 4: // three of a kind / one pair
		for _, v := range cardCounts {
			if v == 3 {
				// three of a kind
				return 3
			} else if v == 2 {
				// one pair
				return 1
			}
		}
	case 5: // high card
		return 0
	}

	panic(fmt.Sprintf("Hand %s can not be scored", hand))
}

// maps the card faces to make them lexicographically comparable
// AKQJT98765432 => ABCDE98765432
func makeLexicographicallyComparable(cards []string) []string {
	mappingTable := map[string]string{"K": "E", "Q": "D", "J": "C", "T": "B"}
	comparableCards := make([]string, len(cards))
	for i, card := range cards {
		mappedCard, exists := mappingTable[card]
		if exists {
			comparableCards[i] = mappedCard
		} else {
			comparableCards[i] = card
		}
	}
	return comparableCards
}

func main() {
	dat, err := os.ReadFile("input.txt")
	// dat, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	type round struct {
		cards []string
		bid   int
	}

	lines := strings.Split(string(dat), "\n")
	rounds := make([]round, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		parts := strings.Split(line, " ")
		rounds[i] = round{cards: strings.Split(parts[0], ""), bid: utils.ParseInt(parts[1])}
	}
	slices.SortFunc(rounds, func(a, b round) int {
		scoreA := scoreHand(a.cards)
		scoreB := scoreHand(b.cards)

		if scoreA == scoreB {
			return strings.Compare(
				strings.Join(makeLexicographicallyComparable(a.cards), ""),
				strings.Join(makeLexicographicallyComparable(b.cards), ""),
			)
		} else {
			return scoreA - scoreB
		}
	})

	sum := 0
	for i, round := range rounds {
		sum += (i + 1) * round.bid
	}

	fmt.Println(rounds)
	fmt.Printf("Part 1: %d\n", sum)
}
