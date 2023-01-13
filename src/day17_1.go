package main

import (
	"fmt"
	"strconv"
)

func day17_1() {

	var countWaysToFill func(amount int, containers []int) int

	countWaysToFill = func(amount int, containers []int) int {
		if len(containers) == 1 {
			if containers[0] == amount {
				return 1
			}
			return 0
		}

		waysWithout := countWaysToFill(amount, containers[1:])
		waysWith := 0

		if containers[0] == amount {
			waysWith = 1
		} else if containers[0] < amount {
			waysWith = countWaysToFill(amount-containers[0], containers[1:])
		}

		return waysWith + waysWithout
	}

	var containers []int
	var amount int

	for lIndex, line := range getInput("day17_input.txt") {

		if lIndex == 0 {
			amount, _ = strconv.Atoi(line[8:])
		} else {
			container, _ := strconv.Atoi(line)
			containers = append(containers, container)
		}
	}

	fmt.Println(countWaysToFill(amount, containers))

}
