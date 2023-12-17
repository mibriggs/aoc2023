package structs

import "slices"

type Pipe struct {
	Symbol string
	Row    int
	Col    int
}

func ContainsPipe(seen []*Pipe, row, col int) bool {
	return slices.IndexFunc(seen, func(p *Pipe) bool { return p.Row == row && p.Col == col }) != -1
}

func GetNextDirection(currentPipe *Pipe, seenPipes []*Pipe, pipeMap [][]string) (int, int) {
	switch currentPipe.Symbol {
	case "|":
		// check if currentRow - 1, col has been seen else return row+1, col
		// Don't really need to worry about out of bounds cause we came from one and the other is the next dir?
		if ContainsPipe(seenPipes, currentPipe.Row-1, currentPipe.Col) {
			return 1, 0
		}
		return -1, 0
	case "-":
		if ContainsPipe(seenPipes, currentPipe.Row, currentPipe.Col-1) {
			return 0, 1
		}
		return 0, -1
	case "L":
		if ContainsPipe(seenPipes, currentPipe.Row-1, currentPipe.Col) {
			return 0, 1
		}
		return -1, 0
	case "J":
		if ContainsPipe(seenPipes, currentPipe.Row-1, currentPipe.Col) {
			return 0, -1
		}
		return -1, 0
	case "7":
		if ContainsPipe(seenPipes, currentPipe.Row, currentPipe.Col-1) {
			return 1, 0
		}
		return 0, -1
	case "F":
		if ContainsPipe(seenPipes, currentPipe.Row+1, currentPipe.Col) {
			return 0, 1
		}
		return 1, 0
	default:
		return 0, 0
	}

}
