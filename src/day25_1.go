package main

import (
	"fmt"
)

func day25_1() {

	startCode := 20151125

	goalRow := 2978
	goalColumn := 3083

	nextCode := func(code int) int {
		return (code * 252533) % 33554393
	}

	row := 1
	col := 1
	code := startCode
	for true {

		if row == 1 {
			row = col + 1
			col = 1
		} else {
			row--
			col++
		}

		code = nextCode(code)

		if row == goalRow && col == goalColumn {
			fmt.Println(code)
			break
		}
	}

}
