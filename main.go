package main

import (
	"ascii-art-fs/checks"
	"ascii-art-fs/splitprint"
	"log"
	"os"
)

func main() {
	input := os.Args[1:]
	if len(input) != 2 {
		log.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}
	myStr := input[0]
	fileName := checks.FileNameCheck(input[1])

	if !checks.IsASCII(myStr) {
		log.Println("non-ASCII character was entered")
		return
	}
	if len(myStr) == 0 {
		return
	}
	if fileName == "" {
		log.Println("only following templates names (standard, shadow, thinkertoy) are available")
		return
	}
	if checks.LineCounter(fileName) != nil {
		return
	}
	splitprint.SplitWord(myStr, fileName)
}
