package main

import (
	"bufio"
	"fmt"
	"os"
)

func FilePath(filePath string) ([]string ,error) { 
  
	f, err := os.Open(filePath)  
     if err != nil {
        fmt.Print("Could'nt load the file, must be some issue")
			return	nil, err

		 }
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
 return lines, nil
}
