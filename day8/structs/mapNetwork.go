package structs

type MapNetwork struct {
	Node  string
	Left  *MapNetwork
	Right *MapNetwork
}

// Given a MapNetwork, how many steps will it take to get to home
func (network MapNetwork) Traverse(steps []string, currentIndex int) int {
	// Basically, we'll be given a list of steps (traversals) to take
	// If we get to the end of the traversal list and we haven't reached the end the we keep repeating the order
	// we take is a list of steps, our current index, and the length of the steps
	// Base case: we are at the end of the traversal
	//		Return 0
	// Recursive case 1: we are at the end of the list but not at the end of the traversal
	// 		continue the recursion (with the right child) but point index back at 0
	// Recursive case 2: we are NOT at the end of the list and currently at R
	// 		Recurse on Right increment currentIndex
	// Recursive case 3: we are NOT at the end of the list and currently at L
	// 		Recurse on Left
	maxIndex := len(steps) - 1
	if network.Node == "ZZZ" {
		return 0
	} else {
		currentStep := steps[currentIndex]
		totalSteps := 1
		if currentIndex == maxIndex { // End of list
			if currentStep == "R" {
				totalSteps += network.Right.Traverse(steps, 0)
			} else {
				totalSteps += network.Left.Traverse(steps, 0)
			}
		} else if currentStep == "R" {
			totalSteps += network.Right.Traverse(steps, currentIndex+1)
		} else {
			totalSteps += network.Left.Traverse(steps, currentIndex+1)
		}
		return totalSteps
	}
}

// Creates new MapNetwork based on the values, sets Left and Right to nil
func CreateNode(nodeVal string) *MapNetwork {
	newNode := MapNetwork{Node: nodeVal, Left: nil, Right: nil}
	return &newNode
}

// Add the other MapNetwork as a child of the current MapNetwork
func (mapNetwork *MapNetwork) AddNode(other *MapNetwork, isLeft bool) {
	if isLeft {
		mapNetwork.Left = other
	} else {
		mapNetwork.Right = other
	}
}

// Iteratively calculates how many moves to get to the end goal descibed by the predicate
func (mapNetwork *MapNetwork) TraverseNonRecursive(steps []string, endingPred func(currVal string) bool) int {
	currentIndex := 0
	maxIndex := len(steps) - 1
	currNode := mapNetwork
	count := 0
	for !endingPred(currNode.Node) {
		currentStep := steps[currentIndex]
		if currentIndex == maxIndex {
			currentIndex = 0
		} else {
			currentIndex++
		}

		if currentStep == "R" {
			currNode = currNode.Right
		} else {
			currNode = currNode.Left
		}
		count++
	}
	return count
}
