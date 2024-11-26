package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid text argument. Usage: go run. <text> <template>")
		return
	}

	// Parse escape sequences (like "\n") in the provided text argument.
	// This function replaces any occurrences of `\n` with actual newline characters.
	text := ParseEscapeSequences(os.Args[1])
	template := os.Args[2]

	templates := map[string]string{
		"standard":   "standard.txt",
		"shadow":     "shadow.txt",
		"thinkertoy": "thinkertoy.txt",
	}
	templatePath, exists := templates[template]
	if !exists {
		fmt.Println("Template not found.")
		return
	}
	// Load the ASCII art template from a file called "standard.txt".
	// This function reads the file, stores ASCII representations of characters in a map, and returns the map.
	asciiMap, err := LoadTemplate(templatePath)
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}
	PrintASCII(asciiMap, text)
}
