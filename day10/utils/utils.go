package utils

import "strings"

func InRange(row, col, maxRow, maxCol int) bool {
	return row >= 0 && row < maxRow && col >= 0 && col <= maxCol
}

func CanExitFromSouth(symbol string) bool {
	validPipes := "|7F"
	return strings.Contains(validPipes, symbol)
}

func CanExitFromNorth(symbol string) bool {
	validPipes := "|LJ"
	return strings.Contains(validPipes, symbol)
}

func CanExitFromWest(symbol string) bool {
	validPipes := "-LF"
	return strings.Contains(validPipes, symbol)
}

func CanExitFromEast(symbol string) bool {
	validPipes := "-J7"
	return strings.Contains(validPipes, symbol)
}
