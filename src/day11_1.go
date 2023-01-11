package main

import (
	"fmt"
)

func day11_1() {

	const iVal = int('i' - 'a')
	const lVal = int('l' - 'a')
	const oVal = int('o' - 'a')

	for _, line := range getInput("day11_input.txt") {

		var password []int
		for _, ch := range line {
			val := int(ch - 'a')
			password = append(password, val)
		}

		passwordLen := len(password)

		for true {
			// increment password
			for i := passwordLen - 1; i >= 0; i-- {
				password[i]++
				if password[i] == iVal || password[i] == lVal || password[i] == oVal {
					password[i]++
					break
				} else if password[i] < 26 {
					break
				} else {
					password[i] = 0
				}
			}

			// validate password
			threeStraight := false
			doubleDouble := false
			firstDouble := -1

			for i := 0; i < passwordLen-1; i++ {
				if !threeStraight && i < passwordLen-2 {
					threeStraight = password[i]+1 == password[i+1] && password[i]+2 == password[i+2]
				}
				if !doubleDouble {
					if password[i] == password[i+1] {
						if firstDouble == -1 {
							firstDouble = i
						} else if firstDouble < i-1 {
							doubleDouble = true
						}
					}
				}
			}

			if threeStraight && doubleDouble {
				finalPassword := ""
				for _, v := range password {
					finalPassword += string(rune(v + 'a'))
				}
				fmt.Println(finalPassword)
				break
			}
		}
	}
}
