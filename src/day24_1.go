package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"golang.org/x/exp/slices"
)

func day24_1() {

	totalWeight := 0
	var allPresents []int
	for _, line := range getInput("day24_input.txt") {
		present, _ := strconv.Atoi(line)
		totalWeight += present
		allPresents = append(allPresents, present)
	}

	targetWeight := totalWeight / 3
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
		var unpicked []int
		for _, p := range allPresents {
			if !slices.Contains(thisArrangement, p) {
				unpicked = append(unpicked, p)
			}
		}
		curSmallest := smallestWorkable
		smallestWorkable = math.MaxInt
		unpickedCount := len(unpicked)
		largest := unpicked[unpickedCount-1]
		isWorkable := false
		pickPresents(
			targetWeight-largest,
			[]int{largest}, unpicked[:unpickedCount-1],
			func(_ []int) {
				isWorkable = true
				smallestWorkable = 0
			})
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
