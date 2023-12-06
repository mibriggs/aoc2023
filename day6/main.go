package main

import (
	"Misc/aoc2023/day6/utils"
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

	waysToWin := 1
	times := []float64{}
	distances := []float64{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, ":")
		floatValues := utils.ToFloatArray(strings.Fields(lineArr[1]))
		if lineArr[0] == "Time" {
			times = floatValues
		} else {
			distances = floatValues
		}
	}

	for i := 0; i < len(times); i++ {
		min, max := utils.ComputeWinningRange(times[i], distances[i])
		waysToWin *= (max - min + 1)
	}
	return waysToWin
}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	time := float64(0)
	distance := float64(0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, ":")
		floatVal := utils.ToFloat(strings.ReplaceAll(lineArr[1], " ", ""))
		if lineArr[0] == "Time" {
			time = floatVal
		} else {
			distance = floatVal
		}
	}

	min, max := utils.ComputeWinningRange(time, distance)
	return (max - min + 1)
}
