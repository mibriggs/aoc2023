package utils

import (
	"slices"
	"strconv"
	"strings"
)

// Splits a line into the array of map values
func GetNodeArray(line string) []string {
	firstNode := strings.ReplaceAll(line, " ", "")
	firstNode = strings.ReplaceAll(firstNode, "=(", ",")
	firstNode = strings.ReplaceAll(firstNode, ")", "")
	return strings.Split(firstNode, ",")
}

// Is the given string 3 capital Zs
func IsTripleZ(currNodeVal string) bool {
	return currNodeVal == "ZZZ"
}

// Does the given string end in Z
func EndsInZ(currNodeVal string) bool {
	asRunes := []rune(currNodeVal)
	return strconv.QuoteRune(asRunes[len(asRunes)-1]) == strconv.QuoteRune('Z')
}

// Does the given string end in A
func EndsInA(currNodeVal string) bool {
	asRunes := []rune(currNodeVal)
	return strconv.QuoteRune(asRunes[len(asRunes)-1]) == strconv.QuoteRune('A')
}

// Computes the LCM of all numbers in array
func GetLcm(numbers []int) int {
	slices.Sort(numbers)
	maxNum := numbers[len(numbers)-1]
	rest := numbers[:len(numbers)-1]
	i := 1
	for i > 0 {
		numerator := maxNum * i
		if canDivideEvenly(numerator, rest) {
			break
		}
		i++

	}
	return i * maxNum
}

// Do all numbers in the array evenly divide the current number
func canDivideEvenly(num int, nums []int) bool {
	for _, currentNum := range nums {
		if num%currentNum != 0 {
			return false
		}
	}
	return true
}
