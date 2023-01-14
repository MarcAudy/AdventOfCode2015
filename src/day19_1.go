package main

import (
	"fmt"
	"regexp"
)

func day19_1() {

	exp := regexp.MustCompile(`(.*) => (.*)`)

	transitions := make(map[string][]string)

	var molecule string
	for _, line := range getInput("day19_input.txt") {
		if len(line) > 0 {
			matches := exp.FindStringSubmatch(line)
			if len(matches) > 0 {
				if _, ok := transitions[matches[1]]; !ok {
					transitions[matches[1]] = make([]string, 0)
				}
				transitions[matches[1]] = append(transitions[matches[1]], matches[2])
			} else {
				molecule = line
			}
		}
	}

	molecules := NewSet[string]()

	for index := range molecule {
		ch := molecule[index : index+1]
		if trans, ok := transitions[ch]; ok {
			for _, t := range trans {
				newMolecule := molecule[:index] + t + molecule[index+1:]
				molecules.Add(newMolecule)
			}
		} else if index < len(molecule)-1 {
			ch := molecule[index : index+2]
			if trans, ok := transitions[ch]; ok {
				for _, t := range trans {
					newMolecule := molecule[:index] + t + molecule[index+2:]
					molecules.Add(newMolecule)
				}
			}
		}
	}

	fmt.Println(len(molecules))
}
