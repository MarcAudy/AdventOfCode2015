package main

import (
	"fmt"
	"strings"
)

func day4_2() {

	for _, line := range getInput("day4_input.txt") {

		num := 1

		for true {
			hash := GetMD5Hash(fmt.Sprintf("%s%d", line, num))
			if strings.HasPrefix(hash, "000000") {
				break
			}
			num++
		}

		fmt.Println(num)
	}
}
