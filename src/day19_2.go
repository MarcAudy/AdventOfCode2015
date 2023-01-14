package main

import (
	"fmt"
	"regexp"
	"strings"
)

func day19_2() {

	exp := regexp.MustCompile(`(.*) => (.*)`)

	transitions := make(map[string][]string)

	var targetMolecule string
	for _, line := range getInput("day19_input.txt") {
		if len(line) > 0 {
			matches := exp.FindStringSubmatch(line)
			if len(matches) > 0 {
				if _, ok := transitions[matches[1]]; !ok {
					transitions[matches[2]] = make([]string, 0)
				}
				transitions[matches[2]] = append(transitions[matches[2]], matches[1])
			} else {
				targetMolecule = line
			}
		}
	}

	seenMolecules := NewSet[string]()

	type MoleculeState struct {
		molecule    string
		transitions int
		score       int
	}

	moleculesToConsider := []MoleculeState{{targetMolecule, 0, len(targetMolecule)}}

	for true {
		mtc := Pop(&moleculesToConsider)

		if mtc.molecule == "e" {
			fmt.Println(mtc.transitions)
			break
		}

		molecule := mtc.molecule
		newTransitionCount := mtc.transitions + 1

		for index := range molecule {
			curSlice := molecule[index:]
			for strand := range transitions {
				if strings.HasPrefix(curSlice, strand) {
					strandLen := len(strand)
					for _, t := range transitions[strand] {
						newMolecule := molecule[:index] + t + molecule[index+strandLen:]
						if seenMolecules.Add(newMolecule) {
							InsertSorted(
								MoleculeState{newMolecule, newTransitionCount, len(newMolecule)},
								&moleculesToConsider,
								func(a *MoleculeState, b *MoleculeState) bool { return a.score > b.score })
						}
					}
				}
			}
		}
	}
}
