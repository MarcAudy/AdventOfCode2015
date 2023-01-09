package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type opType int

const (
	none opType = iota
	not
	and
	or
	lshift
	rshift
)

type expElement interface {
	getValue() (int, bool)
}

type absolute struct {
	value int
}

func (abs absolute) getValue() (int, bool) {
	return abs.value, true
}

type gate struct {
	name string
}

func (g gate) getValue() (int, bool) {
	val, ok := resolvedGates[g.name]
	return val, ok
}

type expression struct {
	left  expElement
	right expElement
	op    opType
}

func (exp expression) resolve() (int, bool) {
	if lval, lok := exp.left.getValue(); lok {
		var rval int
		var rok bool
		if exp.right != nil {
			rval, rok = exp.right.getValue()
		}
		if exp.op == none || exp.op == not || rok {
			var val int
			switch exp.op {
			case none:
				val = lval

			case not:
				val = ^lval

			case and:
				val = lval & rval

			case or:
				val = lval | rval

			case lshift:
				val = lval << rval

			case rshift:
				val = lval >> rval
			}

			return val, true
		}
	}
	return 0, false
}

var resolvedGates = map[string]int{}
var unresolvedGates = map[string]expression{}

func day7() {

	notExp := regexp.MustCompile(`NOT (.*) -> (.*)`)
	opExp := regexp.MustCompile(`(.*) (.*) (.*) -> (.*)`)
	exp := regexp.MustCompile(`(.*) -> (.*)`)

	parseExpression := func(exp string) expElement {
		val, err := strconv.Atoi(exp)
		if err == nil {
			return absolute{val}
		} else {
			return gate{exp}
		}
	}

	for _, line := range getInput("day7_input.txt") {

		var gate string
		var op opType
		var left expElement
		var right expElement

		matches := notExp.FindStringSubmatch(line)
		if len(matches) > 0 {
			gate = matches[2]
			op = not
			left = parseExpression(matches[1])
		} else {
			matches = opExp.FindStringSubmatch(line)
			if len(matches) > 0 {
				gate = matches[4]
				left = parseExpression(matches[1])
				right = parseExpression(matches[3])
				switch matches[2] {
				case "AND":
					op = and
				case "OR":
					op = or
				case "LSHIFT":
					op = lshift
				case "RSHIFT":
					op = rshift
				}
			} else {
				matches := exp.FindStringSubmatch(line)
				gate = matches[2]
				op = none
				left = parseExpression(matches[1])

				abs, ok := left.(absolute)
				if ok {
					resolvedGates[gate] = abs.value
					continue
				}
			}
		}

		unresolvedGates[gate] = expression{left, right, op}
	}

	initialResolved := CopyMap(resolvedGates)
	initialUnresolved := CopyMap(unresolvedGates)

	resolveGates := func() int {
		for true {
			val, ok := resolvedGates["a"]
			if ok {
				return val
			}

			for gate, exp := range unresolvedGates {
				if val, ok := exp.resolve(); ok {
					resolvedGates[gate] = val
					delete(unresolvedGates, gate)
				}
			}
		}
		return -1
	}

	initialA := resolveGates()
	fmt.Printf("PART 1: %d\n", initialA)

	resolvedGates = initialResolved
	unresolvedGates = initialUnresolved

	resolvedGates["b"] = initialA

	fmt.Printf("PART 2: %d\n", resolveGates())
}
