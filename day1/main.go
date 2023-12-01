package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	sol1, err := puzzle1()
	if err == nil {
		fmt.Printf("Answer to solution one is: %d \n", sol1)
	}
}

func puzzle1() (int, error) {
	file, err := openFile("puzzle1.txt")

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
		intArr := stringArrToIntArr(lineArr)
		stringVal := intArr[0] + intArr[len(intArr)-1]

		// convert to int and add to seen
		intVal, _ := strconv.Atoi(stringVal)
		sumSoFar += intVal
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return sumSoFar, nil
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
