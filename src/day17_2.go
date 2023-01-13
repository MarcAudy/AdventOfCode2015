package main

import (
	"fmt"
	"strconv"
)

func day17_2() {

	var containersUsedCounts []int
	var countWaysToFill func(amount int, containersUsed int, containers []int) int

	countWaysToFill = func(amount int, containersUsed int, containers []int) int {
		if len(containers) == 1 {
			if containers[0] == amount {
				containersUsedCounts[containersUsed+1]++
				return 1
			}
			return 0
		}

		waysWithout := countWaysToFill(amount, containersUsed, containers[1:])
		waysWith := 0

		if containers[0] == amount {
			containersUsedCounts[containersUsed+1]++
			waysWith = 1
		} else if containers[0] < amount {
			waysWith = countWaysToFill(amount-containers[0], containersUsed+1, containers[1:])
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
	containersUsedCounts = make([]int, len(containers))

	fmt.Printf("Total ways to fill: %d\n", countWaysToFill(amount, 0, containers))
	for count, ways := range containersUsedCounts {
		if ways > 0 {
			fmt.Printf("Least containers: %d - Ways to use least containers: %d", count, ways)
			break
		}
	}

}
