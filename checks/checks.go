package checks

import (
	"bufio"
	"errors"
	"os"
	"unicode"
)

func FileNameCheck(fName string) string {
	myFile := ""
	switch fName {
	case "standard":
		myFile = "standard.txt"
	case "shadow":
		myFile = "shadow.txt"
	case "thinkertoy":
		myFile = "thinkertoy.txt"
	default:
		myFile = ""
	}

	return myFile
}

func LineCounter(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	if lineCount < 855 {
		return errors.New("file does not contain all characters")
	}
	return nil
}

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
