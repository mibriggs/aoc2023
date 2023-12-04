package main

import (
	"strconv"
)

type PartNumberAsterickMapping struct {
	partNumber    int
	asterickCoord Coordinate
}

type Coordinate struct {
	row int
	col int
}

var topLeft = []Coordinate{{row: 1, col: 0}, {row: 0, col: 1}, {row: 1, col: 1}}
var topRight = []Coordinate{{row: 1, col: 0}, {row: 0, col: -1}, {row: 1, col: -1}}
var bottomRight = []Coordinate{{row: -1, col: 0}, {row: 0, col: -1}, {row: -1, col: -1}}
var bottomLeft = []Coordinate{{row: -1, col: 0}, {row: 0, col: 1}, {row: -1, col: 1}}
var topEdge = []Coordinate{{row: 0, col: 1}, {row: 0, col: -1}, {row: 1, col: 0}, {row: 1, col: -1}, {row: 1, col: 1}}
var bottomEdge = []Coordinate{{row: 0, col: 1}, {row: 0, col: -1}, {row: -1, col: 0}, {row: -1, col: -1}, {row: -1, col: 1}}
var leftEdge = []Coordinate{{row: 1, col: 0}, {row: -1, col: 0}, {row: 0, col: 1}, {row: -1, col: 1}, {row: 1, col: 1}}
var rightEdge = []Coordinate{{row: 1, col: 0}, {row: -1, col: 0}, {row: 0, col: -1}, {row: -1, col: -1}, {row: 1, col: -1}}
var centerPiece = []Coordinate{{row: 0, col: -1}, {row: 0, col: 1}, {row: -1, col: 0}, {row: 1, col: 0}, {row: -1, col: -1}, {row: -1, col: 1}, {row: 1, col: -1}, {row: 1, col: 1}}

// Creates a list of linked lists connected to a symbol
func createLinkedLists(engineSchematic [][]string) []NumberLinkedList {
	numbers := []NumberLinkedList{}
	currentNumber := NumberLinkedList{}
	colLength := len(engineSchematic)
	rowLength := len(engineSchematic[0])

	for i, rows := range engineSchematic {
		for j, item := range rows {
			if isNumber(item) {
				itemCoord := Coordinate{row: i, col: j}
				hasAdjacentSymbol, symbol, symbolCoord := doesNodeHaveAdjacentSymbol(itemCoord, rowLength, colLength, engineSchematic)
				if currentNumber.isEmpty() {
					currentNumber.value = item
					currentNumber.symbol = symbol
					currentNumber.isConnectedToSymbol = hasAdjacentSymbol
					currentNumber.isSymbolAnAsterick = symbol == "*"
					currentNumber.symbolCoord = symbolCoord
				} else {
					currentNumber = currentNumber.addChild(item, symbol, hasAdjacentSymbol, symbol == "*", symbolCoord)
				}
			} else {
				if !currentNumber.isEmpty() {
					numbers = append(numbers, currentNumber)
					currentNumber = NumberLinkedList{}
				}
			}
		}
	}
	return numbers
}

// Affirms if the current string can be converted to a number
func isNumber(maybeNumber string) bool {
	_, err := strconv.Atoi(maybeNumber)
	return err == nil
}

// Does the given coordinates belong to a corner piece?
func isCornerPiece(coord Coordinate, rowLength, columnLength int) bool {
	return (coord.row == 0 && coord.col == 0) ||
		(coord.row == columnLength-1 && coord.col == 0) ||
		(coord.row == columnLength-1 && coord.col == rowLength-1) ||
		(coord.row == 0 && coord.col == rowLength-1)
}

// Does the given coordinates belong to a edge piece?
func isEdgePiece(coord Coordinate, rowLength, columnLength int) bool {
	return (coord.row == 0 || coord.row == columnLength-1 || coord.col == 0 || coord.col == rowLength-1) &&
		!isCornerPiece(coord, rowLength, columnLength)
}

// Is the given string a symbol?
func isSymbol(maybeSymbol string) bool {
	return !(isNumber(maybeSymbol) || maybeSymbol == ".")
}

// Checks if a given node has any adjacent symbols
func doesNodeHaveAdjacentSymbol(node Coordinate, rowLength, colLength int, engineSchematic [][]string) (bool, string, Coordinate) {
	if isCornerPiece(node, rowLength, colLength) {
		if node.col == 0 && node.row == 0 { // top left
			return getSymbol(node, topLeft, engineSchematic)
		} else if node.row == 0 && node.col == rowLength-1 { // top right
			return getSymbol(node, topRight, engineSchematic)
		} else if node.row == colLength-1 && node.col == rowLength-1 { // bottom right
			return getSymbol(node, bottomRight, engineSchematic)
		} else { //bottom left
			return getSymbol(node, bottomLeft, engineSchematic)
		}
	} else if isEdgePiece(node, rowLength, colLength) {
		if node.row == 0 { //top edge
			return getSymbol(node, topEdge, engineSchematic)
		} else if node.row == colLength-1 { // bottom edge
			return getSymbol(node, bottomEdge, engineSchematic)
		} else if node.col == 0 { // left edge
			return getSymbol(node, leftEdge, engineSchematic)
		} else { // right edge
			return getSymbol(node, rightEdge, engineSchematic)
		}
	} else {
		return getSymbol(node, centerPiece, engineSchematic)
	}
}

// returns if its a symbol, type of symbol, and coordinate of symbol
func getSymbol(currentNode Coordinate, coordinatesToCheck []Coordinate, engineSchematic [][]string) (bool, string, Coordinate) {
	for _, coord := range coordinatesToCheck {
		nodeToCheck := Coordinate{row: currentNode.row + coord.row, col: currentNode.col + coord.col}
		maybeSymbol := engineSchematic[nodeToCheck.row][nodeToCheck.col]
		if isSymbol(maybeSymbol) {
			return true, maybeSymbol, nodeToCheck
		}
	}
	return false, "", Coordinate{row: -1, col: -1}
}

// Does the list of coordinates contain the given coordinate
func contains(list []Coordinate, item Coordinate) bool {
	for _, coordinate := range list {
		if coordinate.col == item.col && coordinate.row == item.row {
			return true
		}
	}
	return false
}

// Filters the given list by a predicate of that signature
func filter(listToFilter []PartNumberAsterickMapping, predicate func(Coordinate, []PartNumberAsterickMapping) bool) []PartNumberAsterickMapping {
	newList := []PartNumberAsterickMapping{}
	for _, item := range listToFilter {
		if predicate(item.asterickCoord, listToFilter) {
			newList = append(newList, item)
		}
	}
	return newList
}

// Gets the count of a given coordinate in a list
func count(coordinate Coordinate, list []PartNumberAsterickMapping) int {
	currentCount := 0
	for _, item := range list {
		if item.asterickCoord.col == coordinate.col && item.asterickCoord.row == coordinate.row {
			currentCount++
		}
	}
	return currentCount
}

// Converts a Coordinate to a string
func (coord Coordinate) toString() string {
	return strconv.Itoa(coord.row) + " " + strconv.Itoa(coord.col)
}
