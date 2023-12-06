package utils

import (
	"Misc/aoc2023/shared"
	"math"
	"strconv"
)

// Wrapper around quadratic formula lol
func ComputeWinningRange(time, distance float64) (int, int) {
	min, max := quadratic(1, (-1 * time), distance)
	intMin, intMax := int(math.Ceil(min)), int(math.Floor(max))
	if isSame(intMax, max) {
		intMax--
	}
	if isSame(intMin, min) {
		intMin++
	}
	return intMin, intMax
}

// Converts string array to float64 array
func ToFloatArray(stringReps []string) []float64 {
	floats := []float64{}
	for _, rep := range stringReps {
		floats = append(floats, ToFloat(rep))
	}
	return floats
}

// Converts singular string number to float
func ToFloat(stringRep string) float64 {
	asInt, err := strconv.Atoi(stringRep)
	shared.PanicIfError(err)
	return float64(asInt)
}

// helper function that does quadratic and returns both roots
func quadratic(a, b, c float64) (float64, float64) {
	squaredPart := math.Sqrt(math.Pow(b, 2) - (4 * a * c))
	denom := 2 * a

	return ((-1 * b) - squaredPart) / denom, ((-1 * b) + squaredPart) / denom
}

// Checks if the int and float representation of a number are the exact same
func isSame(intRep int, floatRep float64) bool {
	return math.Abs(floatRep-float64(intRep)) == float64(0)
}
