package main

import (
	"bufio"
	"fmt"
	"os"
)

// LoadTemplate loads the ASCII template into a map
func LoadTemplate(filePath string) (map[rune][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	asciiMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var character rune = ' '
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Move to next character in ASCII map
			asciiMap[character] = lines
			character++
			lines = []string{}
		} else {
			lines = append(lines, line)
		}
	}
	if len(lines) > 0 {
		asciiMap[character] = lines
	}
	return asciiMap, nil
}

// PrintASCII prints the ASCII representation for the given text
func PrintASCII(asciiMap map[rune][]string, text string) {
	for i := 0; i < 8; i++ { // assuming 8 lines per character
		for _, char := range text {
			fmt.Print(asciiMap[char][i])
		}
		fmt.Println()
	}
}

func main() {
	asciiMap, err := LoadTemplate("standard.txt")
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}

	// Example input
text := "mno"
	PrintASCII(asciiMap, text )
}
