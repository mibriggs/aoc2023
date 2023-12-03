package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type NumberIndexTuple struct {
	val   string
	index int
}

func getIntVal(lineArr []string) int {
	// get all chars that can be converted to int and get first/last pair
	intArr := stringArrToIntArr(lineArr)
	stringVal := intArr[0] + intArr[len(intArr)-1]

	// convert to int and add to seen
	intVal, _ := strconv.Atoi(stringVal)
	return intVal
}

// Given a string arr will reduce it to its value that can be parsed as an int
func stringArrToIntArr(stringArr []string) []string {
	intArr := []string{}
	for indx, val := range stringArr {
		_, err := strconv.Atoi(val)
		if err == nil {
			intArr = append(intArr, stringArr[indx])
		}
	}
	return intArr
}

// opens a file lol
func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	return file, nil
}

// Finds all occurrences of each number
func findAllMapOccurrences(line string, numbers []string) []NumberIndexTuple {
	numbersMap := []NumberIndexTuple{}
	for _, val := range numbers {
		occurrencesOfCurrentString := findAllOccurrencesOfString(line, val)
		numbersMap = append(numbersMap, occurrencesOfCurrentString...)
	}
	slices.SortFunc(numbersMap, sortByIndex)
	return numbersMap
}

// finds all occurrences of substring in a given strings
func findAllOccurrencesOfString(fullString string, subString string) []NumberIndexTuple {
	startIndex := 0
	occurrenceIndices := []NumberIndexTuple{}
	for {
		wordIndex := strings.Index(fullString[startIndex:], subString)
		if wordIndex == -1 {
			break
		}

		occurrenceIndices = append(occurrenceIndices, NumberIndexTuple{val: subString, index: wordIndex + startIndex})
		startIndex += (wordIndex + len(subString))
	}
	return occurrenceIndices
}

// sorts an array of my struct by its index value increasing
func sortByIndex(a, b NumberIndexTuple) int {
	if a.index < b.index {
		return -1
	} else if a.index > b.index {
		return 1
	} else {
		return 0
	}
}

// will replace with number representation if still needed
func replaceIfRelevant(line string, mapping []NumberIndexTuple) string {
	adjustmentValue := 0
	validMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for _, currTuple := range mapping {
		startIndex := currTuple.index - adjustmentValue
		endIndex := startIndex + len(currTuple.val)
		isStillRelevant := line[startIndex:endIndex] == currTuple.val

		if isStillRelevant {
			line = strings.Replace(line, line[startIndex:endIndex-1], validMap[currTuple.val], 1)
			adjustmentValue += (len(currTuple.val) - 2)
		}
	}
	return line
}
