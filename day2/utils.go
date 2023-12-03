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

// uses a map to maintain the min cubes/color needed for a game to be valid
func minCountPerCube(game string) map[string]int {
	minCounts := map[string]int{}
	gameInfo := strings.FieldsFunc(game, customSplitFunc)
	rounds := gameInfo[1:]
	for _, round := range rounds {
		cubes := strings.Split(round, ", ")
		for _, cube := range cubes {
			cubeInfo := strings.Fields(cube)
			count, color := cubeInfo[0], cubeInfo[1]
			countAsInt, _ := strconv.Atoi(count)

			val, exists := minCounts[color]
			if exists {
				minCounts[color] = max(val, countAsInt)
			} else {

				minCounts[color] = countAsInt
			}

		}
	}
	return minCounts
}

// computes the power of that game (aka product of min cubes/color)
func computePower(mapping map[string]int) int {
	power := 1
	for _, count := range mapping {
		power *= count
	}
	return power
}

// gets the max of two ints
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Splits a game into an array where the first elem gives the id and the rest = rounds info
func customSplitFunc(r rune) bool {
	return r == ':' || r == ';'
}
