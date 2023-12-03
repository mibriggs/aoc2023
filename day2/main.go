package main

import (
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
)

func main() {
	validGameIdSum := puzzle1("input.txt")
	fmt.Printf("Valid games sum up to: %d\n", validGameIdSum)

	powers := puzzle2("input.txt")
	fmt.Printf("Powers sum up to: %d\n", powers)
}

func puzzle1(filePath string) int {
	file, err := shared.OpenFile(filePath)

	if err != nil {
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validGameSum := 0

	for scanner.Scan() {
		game := scanner.Text()
		result := isValidGame(game)
		if result.isValid {
			validGameSum += result.gameId
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	return validGameSum
}

func puzzle2(filePath string) int {
	file, err := shared.OpenFile(filePath)

	if err != nil {
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPowers := 0

	for scanner.Scan() {
		game := scanner.Text()
		minCounts := minCountPerCube(game)
		power := computePower(minCounts)
		totalPowers += power
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	return totalPowers
}
