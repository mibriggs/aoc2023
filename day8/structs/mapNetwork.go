package structs

type MapNetwork struct {
	Node  string
	Left  *MapNetwork
	Right *MapNetwork
}

// Given a MapNetwork, how many steps will it take to get to home
func (network MapNetwork) Traverse(steps []string, currentIndex int, maxIndex int) int {
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
	if network.Left == nil && network.Right == nil {
		return 0
	} else {
		currentStep := steps[currentIndex]
		totalSteps := 1
		if currentIndex == maxIndex { // End of list
			if currentStep == "R" {
				totalSteps += network.Right.Traverse(steps, 0, maxIndex)
			} else {
				totalSteps += network.Left.Traverse(steps, 0, maxIndex)
			}
		} else if currentStep == "R" {
			totalSteps += network.Right.Traverse(steps, currentIndex+1, maxIndex)
		} else {
			totalSteps += network.Left.Traverse(steps, currentIndex+1, maxIndex)
		}
		return totalSteps
	}
}

func (mapNetwork *MapNetwork) AddNode(other *MapNetwork, isLeft bool) {
	if isLeft {
		mapNetwork.Left = other
	} else {
		mapNetwork.Right = other
	}
}
