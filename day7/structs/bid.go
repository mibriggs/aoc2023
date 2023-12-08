package structs

import (
	"Misc/aoc2023/day7/utils"
	"Misc/aoc2023/shared"
	"bufio"
	"strings"
)

type Bid struct {
	Amount int
	Hand   Hand
}

// Creates a Bid from the give row array using different logic depending on if jokers exist
func ConstructBid(row []string, withJoker bool) Bid {
	if !withJoker {
		return Bid{Amount: shared.ToInt(row[1]), Hand: constructHand(row[0])}
	}
	return Bid{Amount: shared.ToInt(row[1]), Hand: constructHandWithJoker(row[0])}
}

// Returns correct sorting function based on whether this is a joker game or not
func GetAppropriateSortFunction(hasJoker bool) func(a, b Bid) int {
	cardLabelMap := utils.ConstructCardMapping(hasJoker)
	return func(a, b Bid) int {
		if b.Hand.IsStronger(a.Hand, cardLabelMap) {
			return -1
		} else if a.Hand.IsStronger(b.Hand, cardLabelMap) {
			return 1
		} else {
			return 0
		}
	}
}

// Gets bids slice based on input and if we playing with joker rules
func GetBids(scanner *bufio.Scanner, hasJoker bool) []Bid {
	bids := []Bid{}

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Fields(line)
		bids = append(bids, ConstructBid(lineArr, hasJoker))
	}

	return bids
}
