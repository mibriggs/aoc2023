package main

import (
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"strings"
)

var validNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	sol1 := puzzle1()
	fmt.Printf("Answer to problem one is: %d \n", sol1)

	sol2 := puzzle2()
	fmt.Printf("Answer to problem two is: %d \n", sol2)
}

func puzzle1() int {
	file := shared.OpenFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumSoFar := 0

	for scanner.Scan() {
		// get current line and convert to arr
		line := scanner.Text()
		lineArr := strings.Split(line, "")

		// get all chars that can be converted to int and get first/last pair
		sumSoFar += getIntVal(lineArr)
	}

	err := scanner.Err()
	shared.PanicIfError(err)
	return sumSoFar
}

func puzzle2() int {
	file := shared.OpenFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumSoFar := 0

	for scanner.Scan() {
		// get current line and convert to arr
		line := scanner.Text()
		occurrenceMap := findAllMapOccurrences(line, validNumbers)
		newLine := replaceIfRelevant(line, occurrenceMap)
		lineArr := strings.Split(newLine, "")

		// get all chars that can be converted to int and get first/last pair
		sumSoFar += getIntVal(lineArr)
	}

	err := scanner.Err()
	shared.PanicIfError(err)

	return sumSoFar
}
