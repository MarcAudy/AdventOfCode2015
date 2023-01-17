package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"golang.org/x/exp/slices"
)

func day24_2() {

	totalWeight := 0
	var allPresents []int
	for _, line := range getInput("day24_input.txt") {
		present, _ := strconv.Atoi(line)
		totalWeight += present
		allPresents = append(allPresents, present)
	}

	targetWeight := totalWeight / 4
	sort.Slice(allPresents, func(i, j int) bool { return allPresents[i] < allPresents[j] })

	smallestWorkable := math.MaxInt
	var workableArrangements [][]int

	var pickPresents func(target int, picked []int, options []int, onFound func(picked []int))
	pickPresents = func(target int, picked []int, options []int, onFound func(picked []int)) {
		if len(picked) < smallestWorkable {
			optionCount := len(options)
			p := options[optionCount-1]
			if p == target {
				onFound(append(picked, p))
			}
			if optionCount > 1 {
				remainingOptions := options[:optionCount-1]
				if p < target {
					pickPresents(target-p, append(picked, p), remainingOptions, onFound)
				}
				pickPresents(target, picked, remainingOptions, onFound)
			}
		}
	}

	foundWorkable := func(thisArrangement []int) {

		curSmallest := smallestWorkable
		smallestWorkable = math.MaxInt
		isWorkable := false

		var checkFinishedPicking func(picked []int)
		checkFinishedPicking = func(picked []int) {
			if SumOf(picked...) == totalWeight-targetWeight {
				isWorkable = true
				smallestWorkable = 0
			} else {

				var unpicked []int
				for _, p := range allPresents {
					if !slices.Contains(picked, p) {
						unpicked = append(unpicked, p)
					}
				}

				pickPresents(targetWeight, picked, unpicked, checkFinishedPicking)
			}
		}
		checkFinishedPicking(thisArrangement)

		if isWorkable {
			if len(thisArrangement) < curSmallest {
				workableArrangements = [][]int{slices.Clone(thisArrangement)}
				smallestWorkable = len(thisArrangement)
			} else {
				workableArrangements = append(workableArrangements, slices.Clone(thisArrangement))
				smallestWorkable = curSmallest
			}
		} else {
			smallestWorkable = curSmallest
		}

	}

	pickPresents(targetWeight, []int{}, allPresents, foundWorkable)

	smallestQE := math.MaxInt
	for _, workableArrangement := range workableArrangements {
		QE := 1
		for _, present := range workableArrangement {
			QE *= present
		}
		smallestQE = MinOf(smallestQE, QE)
	}

	fmt.Println(smallestQE)
}
