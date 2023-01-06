package main

import (
	"fmt"
)

func day3_2() {

	for _, line := range getInput("day3_input.txt") {

		visited := NewSet[point]()
		var deliverers [2]point

		for index, ch := range line {
			curLoc := &deliverers[index%2]
			switch ch {
			case '<':
				curLoc.x--
				visited.Add(*curLoc)
				break

			case '>':
				curLoc.x++
				visited.Add(*curLoc)
				break

			case '^':
				curLoc.y--
				visited.Add(*curLoc)
				break

			case 'v':
				curLoc.y++
				visited.Add(*curLoc)
				break
			}
		}

		fmt.Println(len(visited))
	}

}
