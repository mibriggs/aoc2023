package main

import (
	"Misc/aoc2023/day7/structs"
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"slices"
)

func main() {
	answer1 := puzzle1("input.txt")
	fmt.Printf("The answer to problem 1 is: %d\n", answer1)

	answer2 := puzzle2("input.txt")
	fmt.Printf("The answer to problem 2 is: %d\n", answer2)
}

func puzzle1(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bids := structs.GetBids(scanner, false)
	totalWinnings := 0

	slices.SortFunc(bids, structs.GetAppropriateSortFunction(false))
	for indx, bid := range bids {
		totalWinnings += (bid.Amount * (indx + 1))
	}
	return totalWinnings
}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bids := structs.GetBids(scanner, true)
	totalWinnings := 0

	slices.SortFunc(bids, structs.GetAppropriateSortFunction(true))
	for indx, bid := range bids {
		totalWinnings += (bid.Amount * (indx + 1))
	}
	return totalWinnings
}
