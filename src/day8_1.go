package main

import (
	"fmt"
)

func day8_1() {

	totalCode := 0
	totalChars := 0

	for _, line := range getInput("day8_input.txt") {

		lineLength := len(line)
		totalCode += lineLength

		for i := 1; i < lineLength-1; i++ {
			totalChars++
			if line[i] == '\\' {
				if line[i+1] == '\\' || line[i+1] == '"' {
					i++
				} else {
					i += 3
				}
			}
		}
	}

	fmt.Println(totalCode - totalChars)
}
