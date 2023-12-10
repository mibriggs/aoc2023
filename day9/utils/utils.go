package utils

import (
	"slices"
)

// Reduces an int array to sum of all its items
func isAllSame(items []int) bool {
	return slices.Min(items) == slices.Max(items)
}

// Gets the difference of all items in array
func getDiffArray(items []int) []int {
	newArray := make([]int, len(items)-1)
	for i := 1; i < len(items); i++ {
		newArray[i-1] = items[i] - items[i-1]
	}
	return newArray
}

// Gets the next element in the sequence
func GetNextHistory(histories []int) int {
	// Base case:
	//	all same return 0 index
	// Recursive case:
	//	last item + recurse on the difference
	if isAllSame(histories) {
		return histories[0]
	}
	length := len(histories)
	lastItem := histories[length-1]
	return lastItem + GetNextHistory(getDiffArray(histories))
}

// Gets the previous element in the sequence
func GetPrevHistory(histories []int) int {
	// Base case:
	//	all same return 0 index
	// Recursive case:
	//	first item - recurse on the difference
	if isAllSame(histories) {
		return histories[0]
	}
	return histories[0] - GetPrevHistory(getDiffArray(histories))
}
