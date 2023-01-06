package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day2() {

	exp := regexp.MustCompile(`(.*)x(.*)x(.*)`)

	totalArea := 0
	totalRibbon := 0

	for _, line := range getInput("day2_input.txt") {
		matches := exp.FindStringSubmatch(line)

		l, _ := strconv.Atoi(matches[1])
		w, _ := strconv.Atoi(matches[2])
		h, _ := strconv.Atoi(matches[3])

		area := 2*l*w + 2*l*h + 2*w*h + MinOf(l*w, l*h, w*h)
		totalArea += area

		longEdge := MaxOf(l, w, h)
		ribbon := 2*(l+w+h-longEdge) + l*w*h
		totalRibbon += ribbon
	}

	fmt.Printf("Wrapping paper required: %d\n", totalArea)
	fmt.Printf("Ribbon required: %d\n", totalRibbon)

}
