package main

import (
	"fmt"
)

func day18_2() {

	var lightGrid [][]bool

	onCount := 0
	for y, line := range getInput("day18_input.txt") {
		lightGrid = append(lightGrid, make([]bool, len(line)))
		for x, ch := range line {
			if ch == '#' {
				lightGrid[y][x] = true
				onCount++
			} else if y == 0 {
				if x == 0 || x == len(line)-1 {
					lightGrid[y][x] = true
					onCount++
				}
			}
		}
	}
	gridWidth := len(lightGrid[0])
	gridHeight := len(lightGrid)
	if !lightGrid[gridHeight-1][0] {
		lightGrid[gridHeight-1][0] = true
		onCount++
	}
	if !lightGrid[gridHeight-1][gridWidth-1] {
		lightGrid[gridHeight-1][gridWidth-1] = true
		onCount++
	}

	isCorner := func(x int, y int) bool {
		if x == 0 || x == gridWidth-1 {
			return y == 0 || y == gridHeight-1
		}
		return false
	}

	countNeighbors := func(x int, y int) int {
		xStart := MaxOf(0, x-1)
		xEnd := MinOf(x+1, gridWidth-1)
		yStart := MaxOf(0, y-1)
		yEnd := MinOf(y+1, gridHeight-1)

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
				if !isCorner(x, y) {
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
		}
		lightGrid = newGrid
	}

	fmt.Println(onCount)
}
