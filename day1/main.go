package main

import (
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"strings"
)

var validNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	sol1, err := puzzle1()
	if err == nil {
		fmt.Printf("Answer to problem one is: %d \n", sol1)
	}

	sol2, err := puzzle2()
	if err == nil {
		fmt.Printf("Answer to problem two is: %d \n", sol2)
	}
}

func puzzle1() (int, error) {
	file, err := shared.OpenFile("input.txt")

	if err != nil {
		return -1, err
	}
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

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return sumSoFar, nil
}

func puzzle2() (int, error) {
	file, err := shared.OpenFile("input.txt")

	if err != nil {
		return -1, err
	}
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
	return sumSoFar, nil
}
