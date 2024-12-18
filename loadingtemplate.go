package main

import (
	"bufio"
	"os"
)

// LoadTemplate reads an ASCII art template file and stores each character's ASCII art
// representation in a map. The map's keys are runes (characters) and values are slices of strings,
// where each string is a line in the ASCII representation of that character.
func LoadTemplate(filePath string) (map[rune][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize a map to store ASCII art representations, with runes as keys and
	// slices of strings as values, where each string is a line in the ASCII art.
	asciiMap := make(map[rune][]string)

	// Prepares the file to be read, we dont actually start reading it yet. . it just sets up the scanner to the input source (file)
	scanner := bufio.NewScanner(file)

	// Initialize `character` to ' ', which will correspond to the ASCII art for space.
	var character rune = ' '
	// Temporary slice to store lines for each character's ASCII representation.
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(lines) > 0 {
				asciiMap[character] = lines // Map the character to its ASCII representation.
				character++                 // Move to the next character in the ASCII sequence.
				lines = []string{}          // Reset the lines slice for the next character.
			}
		} else {
			// If the line is not empty, add it to the current character's lines.
			lines = append(lines, line)
		}
	}
	// After the loop, check if there are remaining lines to add to the map for the final character.
	if len(lines) > 0 {
		asciiMap[character] = lines
	}
	return asciiMap, nil
}
