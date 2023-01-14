package main

import (
	"fmt"
	"math"
)

func day20_2() {

	house := 1

	for true {

		maxDiv := int(math.Floor(math.Sqrt(float64(house))))

		prezzies := 0
		for i := 1; i <= maxDiv; i++ {
			if house%i == 0 {

				m := house / i
				if m <= 50 {
					prezzies += i * 11
				}

				if i <= 50 && i*i != house {
					prezzies += m * 11
				}
			}
		}

		if prezzies >= 29000000 {
			fmt.Println(house)
			break
		}

		house++
	}

}
