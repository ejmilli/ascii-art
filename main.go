package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)!=3{
		fmt.Print("Put Valid Argument")
		return
	}

standardTemplate, err := FilePath("standard.txt")
if err != nil {
   fmt.Println("Could'nt find the:", err)
	 return
} 

shadowTemplate, err := FilePath("shadow.txt")
if err != nil {
   fmt.Println("Could'nt find the:", err)
	 return
} 

thinkertoyTemplate, err := FilePath("thinkertoy.txt")
if err != nil {
   fmt.Println("Could'nt find the:", err)
	 return
} 

fmt.Println("Templates loaded successfully!")
fmt.Println("Standard Template:", standardTemplate)
fmt.Println("Shadow Template:", shadowTemplate)
fmt.Println("thinkertoy Template:", thinkertoyTemplate)

}