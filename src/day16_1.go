package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day16_1() {

	sueExp := regexp.MustCompile(`Sue (.*): (.*): (.*), (.*): (.*), (.*): (.*)`)
	factExp := regexp.MustCompile(`(.*): (.*)`)

	facts := make(map[string]int)

	for _, line := range getInput("day16_input.txt") {

		if !strings.HasPrefix(line, "Sue") {
			factMatches := factExp.FindStringSubmatch(line)
			factVal, _ := strconv.Atoi(factMatches[2])
			facts[factMatches[1]] = factVal
			continue
		} else {
			sueMatch := sueExp.FindStringSubmatch(line)

			isThisSue := true
			for fIndex := 2; fIndex < len(sueMatch); fIndex += 2 {
				sueFact := sueMatch[fIndex]
				sueFactVal, _ := strconv.Atoi(sueMatch[fIndex+1])

				if facts[sueFact] != sueFactVal {
					isThisSue = false
					break
				}
			}

			if isThisSue {
				fmt.Println(sueMatch[1])
				break
			}
		}
	}
}
