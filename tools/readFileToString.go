package tools

import (
	"log"
	"os"
)

// ReadString reads the contents of a file and returns the file content as a string.
func ReadFileToString(filePath string) string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln("Error getting file for string:", err, filePath)
	}
	return string(file)
}
