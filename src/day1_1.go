package main

import (
	"fmt"
)

func day1_1() {

	for _, line := range getInput("day1_input.txt") {
		floor := 0
		for _, ch := range line {
			switch ch {
			case '(':
				floor++
				break

			case ')':
				floor--
				break
			}
		}
		fmt.Println(floor)
	}

}
