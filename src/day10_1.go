package main

import (
	"fmt"
)

func day10_1() {

	nullRune := '\x00'

	for _, line := range getInput("day10_input.txt") {

		sequence := line

		for i := 0; i < 40; i++ {
			lastChar := nullRune
			charRepeat := 0
			newSequence := ""
			for _, ch := range sequence {
				if lastChar == nullRune {
					lastChar = ch
					charRepeat++
				} else if lastChar == ch {
					charRepeat++
				} else {
					newSequence += fmt.Sprintf("%d%c", charRepeat, lastChar)
					lastChar = ch
					charRepeat = 1
				}
			}
			sequence = newSequence + fmt.Sprintf("%d%c", charRepeat, lastChar)
		}

		fmt.Println(len(sequence))
	}
}
