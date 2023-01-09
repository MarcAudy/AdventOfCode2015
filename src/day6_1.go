package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day6_1() {

	exp := regexp.MustCompile(`(.*) (.*),(.*) through (.*),(.*)`)

	var grid [1000][1000]bool
	on := 0

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
					if !grid[x][y] {
						on++
					}
					grid[x][y] = true
				}
			}
		} else if action == "turn off" {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					if grid[x][y] {
						on--
						grid[x][y] = false
					}
				}
			}
		} else {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					grid[x][y] = !grid[x][y]
					if grid[x][y] {
						on++
					} else {
						on--
					}

				}
			}
		}
	}

	fmt.Println(on)
}
