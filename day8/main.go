package main

import (
	"Misc/aoc2023/day8/structs"
	"Misc/aoc2023/day8/utils"
	"Misc/aoc2023/shared"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	answer1 := puzzle1("input.txt")
	fmt.Printf("The answer to problem 1 is: %d\n", answer1)

	answer2 := puzzle2("input.txt")
	fmt.Printf("The answer to problem 2 is: %d\n", answer2)
}

func puzzle1(filePath string) int {
	// list of created nodes
	// for every node that is itself, just make it nil
	// if its ZZZ just make left/right nil
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // point to first line
	steps := strings.Split(scanner.Text(), "")

	seenNodes := map[string]*structs.MapNetwork{}

	for scanner.Scan() {
		if scanner.Text() != "" {
			lineArr := utils.GetNodeArray(scanner.Text())
			currNodeVal, leftNodeVal, rightNodeVal := lineArr[0], lineArr[1], lineArr[2]
			currNode, currNodeExists := seenNodes[currNodeVal]

			if !currNodeExists { // we haven't seen the current node, create and add to seen and set to current
				newNode := structs.CreateNode(currNodeVal)
				seenNodes[newNode.Node] = newNode
				currNode = newNode
			}

			leftNode, leftNodeExists := seenNodes[leftNodeVal]
			rightNode, rightNodeExists := seenNodes[rightNodeVal]

			// right node and left node can be same
			if leftNodeVal == rightNodeVal {
				if !leftNodeExists { // same but both dont exist
					newNode := structs.CreateNode(leftNodeVal)
					seenNodes[newNode.Node] = newNode
					leftNode = newNode
					rightNode = newNode
				} else {
					rightNode = leftNode
				}
			} else {
				if !rightNodeExists {
					newNode := structs.CreateNode(rightNodeVal)
					seenNodes[newNode.Node] = newNode
					rightNode = newNode
				}
				if !leftNodeExists {
					newNode := structs.CreateNode(leftNodeVal)
					seenNodes[newNode.Node] = newNode
					leftNode = newNode
				}
			}
			currNode.AddNode(rightNode, false)
			currNode.AddNode(leftNode, true)
		}
	}

	return seenNodes["AAA"].TraverseNonRecursive(steps, utils.IsTripleZ)
}

func puzzle2(filePath string) int {
	file := shared.OpenFile(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // point to first line
	steps := strings.Split(scanner.Text(), "")

	seenNodes := map[string]*structs.MapNetwork{}
	startNodes := []*structs.MapNetwork{}

	for scanner.Scan() {
		if scanner.Text() != "" {
			lineArr := utils.GetNodeArray(scanner.Text())
			currNodeVal, leftNodeVal, rightNodeVal := lineArr[0], lineArr[1], lineArr[2]
			currNode, currNodeExists := seenNodes[currNodeVal]

			if !currNodeExists { // we haven't seen the current node, create and add to seen and set to current
				newNode := structs.CreateNode(currNodeVal)
				seenNodes[newNode.Node] = newNode
				currNode = newNode
			}

			leftNode, leftNodeExists := seenNodes[leftNodeVal]
			rightNode, rightNodeExists := seenNodes[rightNodeVal]

			// right node and left node can be same
			if leftNodeVal == rightNodeVal {
				if !leftNodeExists { // same but both dont exist
					newNode := structs.CreateNode(leftNodeVal)
					seenNodes[newNode.Node] = newNode
					leftNode = newNode
					rightNode = newNode
				} else {
					rightNode = leftNode
				}
			} else {
				if !rightNodeExists {
					newNode := structs.CreateNode(rightNodeVal)
					seenNodes[newNode.Node] = newNode
					rightNode = newNode
				}
				if !leftNodeExists {
					newNode := structs.CreateNode(leftNodeVal)
					seenNodes[newNode.Node] = newNode
					leftNode = newNode
				}
			}
			currNode.AddNode(rightNode, false)
			currNode.AddNode(leftNode, true)

			if utils.EndsInA(currNode.Node) {
				startNodes = append(startNodes, currNode)
			}
		}
	}

	numberOfTraversals := []int{}
	for _, startNode := range startNodes {
		numberOfTraversals = append(numberOfTraversals, startNode.TraverseNonRecursive(steps, utils.EndsInZ))
	}

	return utils.GetLcm(numberOfTraversals)
}
