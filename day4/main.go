package main

import (
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"math"
)

func main() {
	winnings := puzzle1("input.txt")
	fmt.Printf("Answer to problem 1 is: %d\n", winnings)

	scratchCards := puzzle2("input.txt")
	fmt.Printf("Answer to problem 2 is: %d\n", scratchCards)
}

func puzzle1(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	winningSum := 0

	for scanner.Scan() {
		game := scanner.Text()
		result := getCardWinners(game)
		points := math.Pow(2, float64(result.totalWinningNumbers-1))
		winningSum += int(points)
	}

	err := scanner.Err()
	shared.PanicIfError(err)

	return winningSum
}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	winnings := map[int]int{}
	maxCardNumber := 0

	for scanner.Scan() {
		game := scanner.Text()
		result := getCardWinners(game)

		maxCardNumber = result.cardId

		cardOccurrences, exists := winnings[result.cardId]
		newOccurrences := 1
		if exists {
			newOccurrences += cardOccurrences
		}
		winnings[result.cardId] = newOccurrences

		// compute new winnings
		for i := 1; i <= result.totalWinningNumbers; i++ {
			newCardId := result.cardId + i
			cardOccurrences, exists = winnings[newCardId]
			newNewCardOccurrences := newOccurrences
			if exists {
				newNewCardOccurrences += cardOccurrences
			}
			winnings[newCardId] = newNewCardOccurrences
		}
	}

	totalScratchcards := 0
	for key, val := range winnings {
		if key <= maxCardNumber && key > 0 {
			totalScratchcards += val
		}
	}
	return totalScratchcards
}
