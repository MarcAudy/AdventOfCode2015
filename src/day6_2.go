package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day6_2() {

	exp := regexp.MustCompile(`(.*) (.*),(.*) through (.*),(.*)`)

	var grid [1000][1000]int
	brightness := 0

	for _, line := range getInput("day6_input.txt") {

		matches := exp.FindStringSubmatch(line)
		action := matches[1]
		x1, _ := strconv.Atoi(matches[2])
		y1, _ := strconv.Atoi(matches[3])
		x2, _ := strconv.Atoi(matches[4])
		y2, _ := strconv.Atoi(matches[5])

		if action == "turn on" {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					grid[x][y]++
					brightness++
				}
			}
		} else if action == "turn off" {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					if grid[x][y] > 0 {
						brightness--
						grid[x][y]--
					}
				}
			}
		} else {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					grid[x][y] += 2
					brightness += 2
				}
			}
		}
	}

	fmt.Println(brightness)
}
