package main

import (
	"Misc/aoc2023/day8/structs"
	"Misc/aoc2023/shared"
	"fmt"
)

func main() {
	answer1 := puzzle1("test.txt")
	fmt.Printf("The answer to problem 1 is: %d\n", answer1)

	answer2 := puzzle2("test.txt")
	fmt.Printf("The answer to problem 2 is: %d\n", answer2)

	steps := []string{"L", "L", "R"}
	stepLength := len(steps)
	homeNode := structs.MapNetwork{Node: "ZZZ", Left: nil, Right: nil}
	bNode := structs.MapNetwork{Node: "BBB", Left: nil, Right: &homeNode}
	aNode := structs.MapNetwork{Node: "AAA", Left: &bNode, Right: &bNode}
	bNode.AddNode(&aNode, true)

	fmt.Println(aNode.Traverse(steps, 0, stepLength-1))
}

func puzzle1(filePath string) int {
	// list of created nodes
	// for every node that is itself, just make it nil
	// if its ZZZ just make left/right nil
	file := shared.OpenFile(filePath)
	defer file.Close()

	return -1
}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	return -1
}
