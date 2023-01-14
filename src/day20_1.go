package main

import (
	"fmt"
	"math"
)

func day20_1() {

	house := 1

	for true {

		maxDiv := int(math.Floor(math.Sqrt(float64(house))))

		prezzies := house + 1
		for i := 2; i <= maxDiv; i++ {
			if house%i == 0 {
				prezzies += i
				if i*i != house {
					prezzies += house / i
				}
			}
		}

		if prezzies >= 2900000 {
			fmt.Println(house)
			break
		}

		house++
	}

}
