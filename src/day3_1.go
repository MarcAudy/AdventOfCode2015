package main

import (
	"fmt"
)

func day3_1() {

	for _, line := range getInput("day3_input.txt") {

		visited := NewSet[point]()
		curLoc := point{0, 0}

		for _, ch := range line {
			switch ch {
			case '<':
				curLoc.x--
				visited.Add(curLoc)
				break

			case '>':
				curLoc.x++
				visited.Add(curLoc)
				break

			case '^':
				curLoc.y--
				visited.Add(curLoc)
				break

			case 'v':
				curLoc.y++
				visited.Add(curLoc)
				break
			}
		}

		fmt.Println(len(visited))
	}

}
