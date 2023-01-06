package main

import (
	"fmt"
)

func day1() {

	for _, line := range getInput("day1_input.txt") {
		floor := 0
		foundBasement := false
		for index, ch := range line {
			switch ch {
			case '(':
				floor++
				break

			case ')':
				floor--
				if !foundBasement && floor < 0 {
					foundBasement = true
					fmt.Printf("First entry to basement: %d\n", index+1)
				}
				break
			}
		}
		fmt.Printf("Final floor: %d\n", floor)
	}

}
