package main

import (
	"fmt"
)

func day10_2() {

	for _, line := range getInput("day11_input.txt") {

		var sequence []int
		for _, ch := range line {
			val := int(ch - '0')
			sequence = append(sequence, val)
		}

		for iters := 0; iters < 50; iters++ {
			charRepeat := 1
			lastChar := sequence[0]
			var newSequence []int
			for i := 1; i < len(sequence); i++ {
				if lastChar == sequence[i] {
					charRepeat++
				} else {
					newSequence = append(newSequence, charRepeat, lastChar)
					lastChar = sequence[i]
					charRepeat = 1
				}
			}

			sequence = append(newSequence, charRepeat, lastChar)
		}

		fmt.Println(len(sequence))
	}
}
