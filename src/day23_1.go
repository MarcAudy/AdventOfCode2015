package main

import (
	"fmt"
	"strconv"
)

func day23_1() {

	type operation int
	const (
		hlf = iota
		tpl
		inc
		jmp
		jie
		jio
	)

	type instruction struct {
		op     operation
		regA   bool
		offset int
	}

	var program []instruction

	for _, line := range getInput("day23_input.txt") {
		var instr instruction
		switch line[0:3] {
		case "hlf":
			instr.op = hlf
			instr.regA = line[4] == 'a'
		case "tpl":
			instr.op = tpl
			instr.regA = line[4] == 'a'
		case "inc":
			instr.op = inc
			instr.regA = line[4] == 'a'
		case "jmp":
			instr.op = jmp
			instr.offset, _ = strconv.Atoi(line[4:])
		case "jie":
			instr.op = jie
			instr.regA = line[4] == 'a'
			instr.offset, _ = strconv.Atoi(line[7:])
		case "jio":
			instr.op = jio
			instr.regA = line[4] == 'a'
			instr.offset, _ = strconv.Atoi(line[7:])
		}
		program = append(program, instr)
	}

	registerA := 0
	registerB := 0
	counter := 0
	for counter < len(program) {
		instr := program[counter]
		switch instr.op {
		case hlf:
			if instr.regA {
				registerA /= 2
			} else {
				registerB /= 2
			}
			counter++
		case tpl:
			if instr.regA {
				registerA *= 3
			} else {
				registerB *= 3
			}
			counter++
		case inc:
			if instr.regA {
				registerA++
			} else {
				registerB++
			}
			counter++
		case jmp:
			counter += instr.offset
		case jie:
			if instr.regA && registerA%2 == 0 {
				counter += instr.offset
			} else if !instr.regA && registerB%2 == 0 {
				counter += instr.offset
			} else {
				counter++
			}
		case jio:
			if instr.regA && registerA == 1 {
				counter += instr.offset
			} else if !instr.regA && registerB == 1 {
				counter += instr.offset
			} else {
				counter++
			}
		}
	}

	fmt.Printf("A: %d\n", registerA)
	fmt.Printf("B: %d\n", registerB)
}
