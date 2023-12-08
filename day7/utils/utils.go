package utils

import (
	"strconv"
)

// Creates a map of card hand/suite to rank depending on if playing with jokers or not
func ConstructCardMapping(hasJoker bool) map[rune]int {
	cardLabelMap := map[rune]int{}

	adder := 0
	if !hasJoker {
		adder = -1
		cardLabelMap['J'] = 10

	} else {
		cardLabelMap['J'] = 1
	}

	for i := 2; i <= 9; i++ {
		asRune := []rune(strconv.Itoa(i))[0]
		cardLabelMap[asRune] = i + adder
	}

	cardLabelMap['Q'] = 11
	cardLabelMap['K'] = 12
	cardLabelMap['A'] = 13
	cardLabelMap['T'] = 10 + adder

	return cardLabelMap
}
