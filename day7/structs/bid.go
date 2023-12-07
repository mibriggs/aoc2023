package structs

import (
	"Misc/aoc2023/day7/utils"
	"Misc/aoc2023/shared"
)

type Bid struct {
	Amount int
	Hand   Hand
}

// Creates a Bid from the give row array
func ConstructBid(row []string, withJoker bool) Bid {
	if !withJoker {
		return Bid{Amount: shared.ToInt(row[1]), Hand: constructHand(row[0])}
	}
	return Bid{Amount: shared.ToInt(row[1]), Hand: constructHandWithJoker(row[0])}
}

// Sort by Hand Type using part 1 mapping
type SortByHandType []Bid

func (a SortByHandType) Len() int      { return len(a) }
func (a SortByHandType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByHandType) Less(i, j int) bool {
	return a[j].Hand.IsStronger(a[i].Hand, utils.CardLabelMap)
}

// Sort by Hand Type using part 2 mapping
type SortByHandTypeV2 []Bid

func (a SortByHandTypeV2) Len() int      { return len(a) }
func (a SortByHandTypeV2) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByHandTypeV2) Less(i, j int) bool {
	return a[j].Hand.IsStronger(a[i].Hand, utils.NewCardLabelMap)
}
