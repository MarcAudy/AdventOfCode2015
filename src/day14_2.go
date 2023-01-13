package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day14_2() {

	type Reindeer struct {
		speed    int
		duration int
		rest     int

		resting             bool
		actionTimeRemaining int
		travelled           int
		points              int
	}
	var reindeers []Reindeer

	exp := regexp.MustCompile(`(.*) can fly (.*) km/s for (.*) seconds, but then must rest for (.*) seconds.`)

	var raceLength int
	for lIndex, line := range getInput("day14_input.txt") {

		if lIndex == 0 {
			raceLength, _ = strconv.Atoi(line)
			continue
		}

		matches := exp.FindStringSubmatch(line)

		speed, _ := strconv.Atoi(matches[2])
		duration, _ := strconv.Atoi(matches[3])
		rest, _ := strconv.Atoi(matches[4])

		reindeers = append(reindeers, Reindeer{speed: speed, duration: duration, rest: rest, resting: false, actionTimeRemaining: duration})
	}

	for i := 0; i < raceLength; i++ {

		maxDistance := 0
		for rIndex := range reindeers {
			rd := &reindeers[rIndex]
			if !rd.resting {
				rd.travelled += rd.speed
			}
			maxDistance = MaxOf(maxDistance, rd.travelled)
			rd.actionTimeRemaining--
			if rd.actionTimeRemaining == 0 {
				if rd.resting {
					rd.resting = false
					rd.actionTimeRemaining = rd.duration
				} else {
					rd.resting = true
					rd.actionTimeRemaining = rd.rest
				}
			}
		}

		for rIndex := range reindeers {
			rd := &reindeers[rIndex]
			if rd.travelled == maxDistance {
				rd.points++
			}
		}
	}

	maxPoints := 0
	for _, rd := range reindeers {
		maxPoints = MaxOf(maxPoints, rd.points)
	}

	fmt.Println(maxPoints)
}
