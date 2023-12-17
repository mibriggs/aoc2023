package main

import (
	"Misc/aoc2023/day10/structs"
	"Misc/aoc2023/day10/utils"
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func main() {
	answer1 := puzzle1("input.txt")
	fmt.Printf("The answer to problem 1 is: %d\n", answer1)

	answer2 := puzzle2("test1.txt")
	fmt.Printf("The answer to problem 2 is: %d\n", answer2)
}

func puzzle1(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	bp := [][]string{}
	scanner := bufio.NewScanner(file)

	rowNum := 0
	sPipe := structs.Pipe{}
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, "")
		sCol := slices.Index(lineArr, "S")
		if sCol != -1 {
			sPipe.Col = sCol
			sPipe.Row = rowNum
			sPipe.Symbol = "S"
		}
		bp = append(bp, lineArr)
		rowNum++
	}

	currentPipe := &sPipe
	seenPipes := []*structs.Pipe{}
	maxRow, maxCol := len(bp), len(bp[0])

	for !structs.ContainsPipe(seenPipes, currentPipe.Row, currentPipe.Col) {
		nextPipe := structs.Pipe{}
		if currentPipe.Symbol == "S" {
			currRow, currCol := currentPipe.Row, currentPipe.Col
			if utils.InRange(currRow-1, currCol, maxRow, maxCol) &&
				utils.CanExitFromSouth(bp[currRow-1][currCol]) {
				nextPipe = structs.Pipe{Symbol: bp[currRow-1][currCol], Row: currRow - 1, Col: currCol}
			} else if utils.InRange(currRow+1, currCol, maxRow, maxCol) &&
				utils.CanExitFromNorth(bp[currRow+1][currCol]) {
				nextPipe = structs.Pipe{Symbol: bp[currRow+1][currCol], Row: currRow + 1, Col: currCol}
			} else if utils.InRange(currRow, currCol+1, maxRow, maxCol) &&
				utils.CanExitFromSouth(bp[currRow][currCol+1]) {
				nextPipe = structs.Pipe{Symbol: bp[currRow][currCol+1], Row: currRow, Col: currCol + 1}
			} else {
				nextPipe = structs.Pipe{Symbol: bp[currRow][currCol-1], Row: currRow, Col: currCol - 1}
			}
		} else {
			nextRow, nextCol := structs.GetNextDirection(currentPipe, seenPipes, bp)
			nextPipe = structs.Pipe{
				Symbol: bp[currentPipe.Row+nextRow][currentPipe.Col+nextCol],
				Row:    currentPipe.Row + nextRow,
				Col:    currentPipe.Col + nextCol,
			}
		}
		seenPipes = append(seenPipes, currentPipe)
		currentPipe = &nextPipe
	}
	return len(seenPipes) / 2
}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return -1
}
