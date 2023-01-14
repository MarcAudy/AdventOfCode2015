package main

import (
	"fmt"
)

func day18_1() {

	var lightGrid [][]bool

	onCount := 0
	for y, line := range getInput("day18_input.txt") {
		lightGrid = append(lightGrid, make([]bool, len(line)))
		for x, ch := range line {
			if ch == '#' {
				lightGrid[y][x] = true
				onCount++
			}
		}

	}

	countNeighbors := func(x int, y int) int {
		xStart := MaxOf(0, x-1)
		xEnd := MinOf(x+1, len(lightGrid[0])-1)
		yStart := MaxOf(0, y-1)
		yEnd := MinOf(y+1, len(lightGrid)-1)

		neighbors := 0
		for nx := xStart; nx <= xEnd; nx++ {
			for ny := yStart; ny <= yEnd; ny++ {
				if nx != x || ny != y {
					if lightGrid[ny][nx] {
						neighbors++
					}
				}
			}
		}
		return neighbors
	}

	for i := 0; i < 100; i++ {

		newGrid := DeepCopy(lightGrid)

		for y := range lightGrid {
			for x, on := range lightGrid[y] {
				neighbors := countNeighbors(x, y)
				if on {
					if neighbors != 2 && neighbors != 3 {
						newGrid[y][x] = false
						onCount--
					}
				} else {
					if neighbors == 3 {
						newGrid[y][x] = true
						onCount++
					}
				}
			}
		}
		lightGrid = newGrid
	}

	fmt.Println(onCount)
}
