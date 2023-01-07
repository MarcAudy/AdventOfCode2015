package main

import (
	"fmt"
)

func day5_1() {

	nice := 0

	for _, line := range getInput("day5_input.txt") {

		naughty := false
		vowels := 0
		doubleLetter := false

	line:
		for index, ch := range []byte(line) {

			checkNext := func(ch byte) bool { return index < len(line)-1 && line[index+1] == ch }

			switch ch {
			case 'a':
				vowels++
				if checkNext('b') {
					naughty = true
					break line
				}

			case 'c':
				if checkNext('d') {
					naughty = true
					break line
				}

			case 'e':
				vowels++
			case 'i':
				vowels++
			case 'o':
				vowels++
			case 'u':
				vowels++

			case 'p':
				if checkNext('q') {
					naughty = true
					break line
				}

			case 'x':
				if checkNext('y') {
					naughty = true
					break line
				}
			}

			doubleLetter = doubleLetter || checkNext(ch)
		}

		if !naughty && vowels >= 3 && doubleLetter {
			nice++
		}
	}

	fmt.Printf("Total nice words: %d\n", nice)
}
