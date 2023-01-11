package main

import (
	"fmt"
	"strconv"
)

func day12_1() {

	for _, line := range getInput("day12_input.txt") {

		sum := 0
		numStart := -1
		for index, ch := range line {
			if ch == '-' || (ch >= '0' && ch <= '9') {
				if numStart == -1 {
					numStart = index
				}
			} else if numStart != -1 {
				num, _ := strconv.Atoi(line[numStart:index])
				sum += num
				numStart = -1
			}
		}

		fmt.Println(sum)
	}

}
