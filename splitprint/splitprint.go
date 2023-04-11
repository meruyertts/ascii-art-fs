package splitprint

import (
	"ascii-art-fs/read"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func SplitWord(myStr, myFile string) {
	re := regexp.MustCompile(`\\n`)
	newStr := re.Split(myStr, -1)
	for i := 0; i < len(newStr); i++ {
		if len(newStr[i]) > 0 {
			PrintWord(newStr[i], myFile)
		}
		if newStr[i] == "" {
			fmt.Println("")
		}
	}
}

func PrintWord(s, fileName string) {
	myLine := ""
	myarray := [8]string{}
	for _, char := range s {
		for line := 2; line <= 9; line++ {
			myrune := int(char)
			for i := ' '; i <= '~'; i++ {
				j := (int(i) - ' ') * 9
				if myrune == int(i) {
					firstline, err := read.ReadExactLine(fileName, line+j)
					if err != nil {
						log.Print(err)
						return
					}
					myLine += firstline
				}
			}
		}
		temp := strings.Split(myLine, "\n")
		for index, s := range temp[:len(temp)-1] {
			myarray[index] += s
		}
		myLine = ""
	}
	for _, i := range myarray {
		fmt.Println(i)
	}
}
