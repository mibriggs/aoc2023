package shared

import (
	"os"
	"strconv"
)

// opens a file lol
func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	PanicIfError(err)
	return file
}

// Crashes program if an error occurs
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// Converts given string to an int
func ToInt(intString string) int {
	asInt, err := strconv.Atoi(intString)
	PanicIfError(err)
	return asInt
}
