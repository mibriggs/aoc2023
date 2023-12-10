package main

import (
	"Misc/aoc2023/day9/utils"
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"strings"
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
	totalNextHistories := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Fields(line)
		intArr := make([]int, len(lineArr))
		for indx, val := range lineArr {
			intArr[indx] = shared.ToInt(val)
		}
		totalNextHistories += utils.GetNextHistory(intArr)
	}

	return totalNextHistories
}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalNextHistories := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Fields(line)
		intArr := make([]int, len(lineArr))
		for indx, val := range lineArr {
			intArr[indx] = shared.ToInt(val)
		}
		totalNextHistories += utils.GetPrevHistory(intArr)
	}

	return totalNextHistories
}
