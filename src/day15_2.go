package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day15_2() {

	exp := regexp.MustCompile(`(.*): capacity (.*), durability (.*), flavor (.*), texture (.*), calories (.*)`)

	type Ingredient struct {
		name       string
		capacity   int
		durability int
		flavor     int
		texture    int
		calories   int
	}

	var ingredients []Ingredient

	for _, line := range getInput("day15_input.txt") {
		matches := exp.FindStringSubmatch(line)

		capacity, _ := strconv.Atoi(matches[2])
		durability, _ := strconv.Atoi(matches[3])
		flavor, _ := strconv.Atoi(matches[4])
		texture, _ := strconv.Atoi(matches[5])
		calories, _ := strconv.Atoi(matches[6])

		ingredients = append(ingredients, Ingredient{matches[1], capacity, durability, flavor, texture, calories})
	}

	scoreCookie := func(amounts []int) int {
		capacity := 0
		durability := 0
		flavor := 0
		texture := 0
		calories := 0
		for index := range ingredients {
			amt := amounts[index]
			ing := &ingredients[index]
			capacity += amt * ing.capacity
			durability += amt * ing.durability
			flavor += amt * ing.flavor
			texture += amt * ing.texture
			calories += amt * ing.calories
		}
		if calories == 500 {
			return MaxOf(0, capacity) * MaxOf(0, durability) * MaxOf(0, flavor) * MaxOf(0, texture)
		} else {
			return 0
		}

	}

	var testRecipe func(amount int, amounts []int) int
	testRecipe = func(amount int, amounts []int) int {

		aIndex := len(amounts)
		amounts = append(amounts, amount)
		if aIndex+1 == len(ingredients) {
			return scoreCookie(amounts)
		}

		best := 0
		for a := 0; a <= amount; a++ {
			amounts[aIndex] = a
			score := testRecipe(amount-a, amounts)
			best = MaxOf(score, best)
		}
		return best
	}

	fmt.Println(testRecipe(100, []int{}))

}
