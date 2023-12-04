package main

import "strconv"

type Coordinate struct {
	row int
	col int
}

// Gets the count of the current Coordinate in a given list
func (coordinate Coordinate) count(list []PartNumberAsterickMapping) int {
	currentCount := 0
	for _, item := range list {
		if item.asterickCoord.col == coordinate.col && item.asterickCoord.row == coordinate.row {
			currentCount++
		}
	}
	return currentCount
}

// Converts the current Coordinate to a string
func (coord Coordinate) toString() string {
	return strconv.Itoa(coord.row) + " " + strconv.Itoa(coord.col)
}
