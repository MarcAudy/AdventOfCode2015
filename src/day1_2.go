package main

import (
	"fmt"
)

func day1_2() {

	for _, line := range getInput("day1_input.txt") {
		floor := 0
	line:
		for index, ch := range line {
			switch ch {
			case '(':
				floor++
				break

			case ')':
				floor--
				if floor < 0 {
					fmt.Printf("Basement: %d\n", index+1)
					break line
				}
				break
			}
		}
	}

}
