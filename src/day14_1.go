package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day14_1() {

	const raceLength = 2503

	exp := regexp.MustCompile(`(.*) can fly (.*) km/s for (.*) seconds, but then must rest for (.*) seconds.`)

	maxDistance := 0
	for _, line := range getInput("day14_input.txt") {

		matches := exp.FindStringSubmatch(line)

		speed, _ := strconv.Atoi(matches[2])
		duration, _ := strconv.Atoi(matches[3])
		rest, _ := strconv.Atoi(matches[4])

		cycles := raceLength / (duration + rest)
		remaining := raceLength - (cycles * (duration + rest))
		distance := 0
		if remaining >= duration {
			cycles++
		} else {
			distance = speed * remaining
		}
		distance += speed * duration * cycles

		maxDistance = MaxOf(maxDistance, distance)
	}

	fmt.Println(maxDistance)
}
