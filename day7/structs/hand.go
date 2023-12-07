package structs

import (
	"slices"
	"strconv"
)

type Hand struct {
	Type  HandType
	Value string
}

// Is the current Hand stronger than the other Hand?
func (curr Hand) IsStronger(other Hand, mapping map[rune]int) bool {
	if curr.Type == other.Type {
		currValRunes := []rune(curr.Value)
		otherValRunes := []rune(other.Value)

		for i := 0; i < len(currValRunes); i++ {
			currVal := currValRunes[i]
			otherVal := otherValRunes[i]
			if mapping[currVal] != mapping[otherVal] {
				return mapping[currVal] > mapping[otherVal]
			}
		}
		return false
	}
	return curr.Type.IsGreater(other.Type)
}

// Creates a Hand from a cards string
func constructHand(cards string) Hand {

	cardCounts := getCardsCount(cards)
	vals := getVals(cardCounts)
	slices.SortFunc(vals, sortCountAscending)

	countMax := vals[0]
	handType := getHandType(countMax, vals)

	return Hand{Type: handType, Value: cards}
}

// Creates a Hand from a cards string with new Joker Directive
func constructHandWithJoker(cards string) Hand {

	cardCounts := getCardsCount(cards)
	oldVal, jExists := cardCounts[strconv.QuoteRune('J')]
	if jExists {
		cardCounts[strconv.QuoteRune('J')] = 0
	}

	vals := getVals(cardCounts)
	slices.SortFunc(vals, sortCountAscending)

	countMax := vals[0]
	countMax += oldVal

	handType := getHandType(countMax, vals)

	return Hand{Type: handType, Value: cards}
}

// Gets card to count map of all cards in your hand
func getCardsCount(cards string) map[string]int {
	cardCount := map[string]int{}
	for _, currRune := range cards {
		asString := strconv.QuoteRune(currRune)
		count := 1
		val, exists := cardCount[asString]
		if exists {
			count += val
		}
		cardCount[asString] = count
	}

	return cardCount
}

// Gets count of all cards in your hand
func getVals(cardCount map[string]int) []int {
	vals := make([]int, 0, len(cardCount))
	for _, val := range cardCount {
		vals = append(vals, val)
	}
	return vals
}

// Sorts int array in ascending order
func sortCountAscending(a, b int) int {
	if a > b {
		return -1
	} else if a < b {
		return 1
	} else {
		return 0
	}
}

// Returns HandType based on highest number of cards
func getHandType(maxCount int, vals []int) HandType {
	var handType HandType
	if maxCount == 5 {
		handType = FIVE
	} else if maxCount == 4 {
		handType = FOUR
	} else if maxCount == 3 && vals[1] == 2 {
		handType = FULL
	} else if maxCount == 3 {
		handType = THREE
	} else if maxCount == 2 && vals[1] == 2 {
		handType = TWO
	} else if maxCount == 2 {
		handType = ONE
	} else {
		handType = HIGH
	}
	return handType
}
