package main

import (
	"bufio" // Import bufio
	"fmt"
	"os"
)

func FilePath(filePath string) ([]string, error) {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		// Improved error message for better clarity
		return nil, fmt.Errorf("could not load the file %s: %v", filePath, err)
	}
	defer f.Close()

	var lines []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 { // Only add non-empty lines
			lines = append(lines, line)
		}
	}

	// Check if there was any error during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file %s: %v", filePath, err)
	}

	return lines, nil
}



func main() {
	if len(os.Args) != 2 {
		fmt.Print("Put Valid Argument")
		return
	}

	// Load templates from files
	standardTemplate, err := FilePath("standard.txt")
	if err != nil {
		fmt.Println("Could'nt find the:", err)
		return
	}

	// Map ASCII art to characters
fmt.Print(standardTemplate)
}
