package shared

import (
	"fmt"
	"os"
)

// opens a file lol
func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	return file, nil
}
