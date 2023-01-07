package main

import (
	"fmt"
	"strings"
)

func day5_2() {

	nice := 0

	for _, line := range getInput("day5_input.txt") {

		doubleLetter := false
		doublePair := false

		for index, ch := range []byte(line) {

			lineLen := len(line)
			doubleLetter = doubleLetter || (index < lineLen-2 && line[index+2] == ch)

			if !doublePair && index < lineLen-3 {
				doublePair = strings.Index(line[index+2:], line[index:index+2]) != -1
			}

			if doubleLetter && doublePair {
				nice++
				break
			}
		}

	}

	fmt.Printf("Total nice words: %d\n", nice)
}
