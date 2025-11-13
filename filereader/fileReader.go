package main

import (
	"bufio"
	"fmt"
	"os"
)

const filePath = "myFile.txt"

func main() {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Create a buffered reader
	reader := bufio.NewReader(file)

	for {
		// Read line by line
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}

	// Close the file
	file.Close()
}
