package main

import (
	"fmt"
)

func main() {
	// red blue or green cube
	// x amount of each
	// each round pull out x number of cubes

	text := "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	fmt.Println(isValidGame(text))
}
