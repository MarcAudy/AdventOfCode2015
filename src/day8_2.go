package main

import (
	"fmt"
)

func day8_2() {

	totalCode := 0
	totalEncoded := 0

	for _, line := range getInput("day8_input.txt") {

		lineLength := len(line)
		totalCode += lineLength
		totalEncoded += lineLength + 2

		for i := 0; i < lineLength; i++ {
			if line[i] == '\\' || line[i] == '"' {
				totalEncoded++
			}
		}
	}

	fmt.Println(totalEncoded - totalCode)
}
