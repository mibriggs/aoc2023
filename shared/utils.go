package shared

import (
	"os"
)

// opens a file lol
func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	PanicIfError(err)
	return file
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
