package structs

type HandType int64

const (
	FIVE HandType = iota
	FOUR
	FULL
	THREE
	TWO
	ONE
	HIGH
)

// String rep of a HandType
func (handType HandType) String() string {
	var stringRep string
	switch handType {
	case FIVE:
		stringRep = "Five of a kind"
	case FOUR:
		stringRep = "Four of a kind"
	case FULL:
		stringRep = "Full house"
	case THREE:
		stringRep = "Three of a kind"
	case TWO:
		stringRep = "Two pair"
	case ONE:
		stringRep = "One pair"
	case HIGH:
		stringRep = "High card"
	default:
		stringRep = ""
	}
	return stringRep
}

// Is the current HandType greater than the other HandType
func (curr HandType) IsGreater(other HandType) bool {
	return curr < other
}
