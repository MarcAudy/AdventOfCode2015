package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/constraints"
)

type point struct {
	x int
	y int
}

func getInput(fileName string) []string {
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	return fileLines
}

func MinOf[T constraints.Ordered](vars ...T) T {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf[T constraints.Ordered](vars ...T) T {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
