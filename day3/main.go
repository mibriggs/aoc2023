package main

import (
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	answer1 := puzzle1("input.txt")
	fmt.Printf("Solution to problem 1 is: %d\n", answer1)
	answer2 := puzzle2("input.txt")
	fmt.Printf("Solution to problem 2 is: %d\n", answer2)
}

func puzzle1(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	engineSchematic := [][]string{}
	partNumbersSum := 0

	for scanner.Scan() {
		row := scanner.Text()
		rowArray := strings.Split(row, "")
		engineSchematic = append(engineSchematic, rowArray)
	}

	potentialPartNumbers := createLinkedLists(engineSchematic)
	for _, ll := range potentialPartNumbers {
		if ll.hasAdjacentSymbol() {
			partNumbersSum += ll.toInt()
		}
	}

	return partNumbersSum
	// If we're a corner piece places to check = 2 horizontal/vertical neighbor + 1 diagonal
	// If edge piece places to check = 3 horizontal/vertical neighbor + 2 diagonal
	// Else all diagonals + horiztonal/veritcal

}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	engineSchematic := [][]string{}

	for scanner.Scan() {
		row := scanner.Text()
		rowArray := strings.Split(row, "")
		engineSchematic = append(engineSchematic, rowArray)
	}

	potentialPartNumbers := createLinkedLists(engineSchematic)
	potentialGears := []NumberLinkedList{}
	for _, ll := range potentialPartNumbers {
		if ll.hasAdjacentSymbol() && ll.hasAdjacentAsterick() {
			potentialGears = append(potentialGears, ll)
		}
	}

	flattenedPotentialGears := []PartNumberAsterickMapping{}
	for _, ll := range potentialGears {
		partNumber := ll.toInt()
		asterickCoords := ll.getUniqueAstericks()
		for _, coord := range asterickCoords {
			flattenedPotentialGears = append(flattenedPotentialGears, PartNumberAsterickMapping{partNumber: partNumber, asterickCoord: coord})
		}
	}

	gears := filter(flattenedPotentialGears, func(c Coordinate, pnam []PartNumberAsterickMapping) bool {
		return count(c, pnam) == 2
	})

	gearRatios := map[string]int{}
	for _, gear := range gears {
		newKey := gear.asterickCoord.toString()
		value, exists := gearRatios[newKey]
		if exists {
			gearRatios[newKey] = value * gear.partNumber
		} else {
			gearRatios[newKey] = gear.partNumber
		}
	}

	gearRatiosSum := 0
	for _, val := range gearRatios {
		gearRatiosSum += val
	}

	return gearRatiosSum
}
