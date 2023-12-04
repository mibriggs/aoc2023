package main

import (
	"Misc/aoc2023/shared"
	"strconv"
	"strings"
)

type WinningNumbers struct {
	cardId              int
	totalWinningNumbers int
}

// Gets card id and winning count for each scratch card
func getCardWinners(scratchCard string) WinningNumbers {
	cardArr := strings.FieldsFunc(scratchCard, splitByBarOrColon)

	cardId, err := strconv.Atoi(strings.Fields(cardArr[0])[1])
	shared.PanicIfError(err)

	winningNumbers := strings.Fields(cardArr[1])
	myPicks := strings.Fields(cardArr[2])
	seen := []string{}

	for _, pick := range myPicks {
		if indexOf(pick, winningNumbers) != -1 && indexOf(pick, seen) == -1 {
			seen = append(seen, pick)
		}
	}
	return WinningNumbers{cardId: cardId, totalWinningNumbers: len(seen)}
}

// Gets the index of the item we are searching for in given list or -1 if not found
func indexOf(itemToFind string, list []string) int {
	for indx, val := range list {
		if val == itemToFind {
			return indx
		}
	}
	return -1
}

// split string by this predicate
func splitByBarOrColon(r rune) bool {
	return r == '|' || r == ':'
}
