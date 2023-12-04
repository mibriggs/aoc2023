package main

import (
	"Misc/aoc2023/shared"
	"strconv"
)

type NumberLinkedList struct {
	value               string
	symbol              string
	isConnectedToSymbol bool
	isSymbolAnAsterick  bool
	symbolCoord         Coordinate
	parent              *NumberLinkedList
}

// Checks if this linked list is empty
func (ll NumberLinkedList) isEmpty() bool {
	return !ll.isConnectedToSymbol && ll.value == "" && ll.parent == nil
}

// Takes converts this linkedList to its part number
func (ll NumberLinkedList) toInt() int {
	stringRep := linkedListToIntHelper(ll)
	intRep, err := strconv.Atoi(stringRep)
	shared.PanicIfError(err)
	return intRep
}

// Recursively creates the full number string
func linkedListToIntHelper(ll NumberLinkedList) string {
	if ll.parent == nil {
		return ll.value
	}
	return linkedListToIntHelper(*ll.parent) + ll.value
}

// Checks if any node in this linked list is connected to a symbol
func (ll NumberLinkedList) hasAdjacentSymbol() bool {
	if ll.parent == nil {
		return ll.isConnectedToSymbol
	} else if ll.isConnectedToSymbol {
		return true // early return so we don't have to go deeper if we already know the answer
	} else {
		return ll.isConnectedToSymbol || ll.parent.hasAdjacentSymbol()
	}
}

// Checks if any node in this linked list is connected to an asterick
func (ll NumberLinkedList) hasAdjacentAsterick() bool {
	if ll.parent == nil {
		return ll.isSymbolAnAsterick
	} else if ll.isSymbolAnAsterick {
		return true // early return so we don't have to go deeper if we already know the answer
	} else {
		return ll.isSymbolAnAsterick || ll.parent.hasAdjacentAsterick()
	}
}

// Adds a new value to this linked list
func (parent NumberLinkedList) addChild(val string, symbol string, isConnectedToSymbol bool, isSymbolAnAsterick bool, symbolCoord Coordinate) NumberLinkedList {
	newList := NumberLinkedList{
		value:               val,
		symbol:              symbol,
		isConnectedToSymbol: isConnectedToSymbol,
		isSymbolAnAsterick:  isSymbolAnAsterick,
		symbolCoord:         symbolCoord,
		parent:              &parent,
	}
	return newList
}

// Gets all UNIQUE astericks connected to this linked list
func (ll NumberLinkedList) getUniqueAstericks() []Coordinate {
	allAstericks := getAllAsterickCoords(ll)
	uniqueList := []Coordinate{}
	for _, asterickCoord := range allAstericks {
		if !contains(uniqueList, asterickCoord) {
			uniqueList = append(uniqueList, asterickCoord)
		}
	}
	return uniqueList
}

// Gets all astericks that are connected to this linked list
func getAllAsterickCoords(ll NumberLinkedList) []Coordinate {
	if ll.parent == nil { // end if list
		if ll.isSymbolAnAsterick {
			return []Coordinate{ll.symbolCoord}
		}
		return []Coordinate{}
	} else { // not at end
		toReturn := []Coordinate{}
		if ll.isSymbolAnAsterick {
			toReturn = append(toReturn, ll.symbolCoord)
		}
		return append(getAllAsterickCoords(*ll.parent), toReturn...)
	}
}
