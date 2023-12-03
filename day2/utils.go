package main

import (
	"strconv"
	"strings"
)

type GameResult struct {
	gameId  int
	isValid bool
}

var maxCountPerCubeColor = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// For a given game, checks if each round passes the condition
// @Returns an object with the game id and whether its valid or not
func isValidGame(game string) GameResult {
	gameInfo := strings.FieldsFunc(game, customSplitFunc)
	gameNumber, _ := strconv.Atoi(strings.Fields(gameInfo[0])[1])

	for _, round := range gameInfo[1:] {
		picks := strings.Split(round, ", ")
		for _, pick := range picks {
			cubeInfo := strings.Fields(pick)
			count, _ := strconv.Atoi(cubeInfo[0])
			cubeColor := cubeInfo[1]
			if maxCountPerCubeColor[cubeColor] < count {
				return GameResult{gameId: gameNumber, isValid: false}
			}
		}
	}
	return GameResult{gameId: gameNumber, isValid: true}
}

// Splits a game into an array where the first elem gives the id and the rest = rounds info
func customSplitFunc(r rune) bool {
	return r == ':' || r == ';'
}
