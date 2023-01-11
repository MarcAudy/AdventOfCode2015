package main

import (
	"encoding/json"
	"fmt"
)

func day12_2() {

	var countJson func(jsonObj any) float64
	countJson = func(jsonObj any) float64 {
		if num, ok := jsonObj.(float64); ok {
			return num
		}

		if m, ok := jsonObj.(map[string]any); ok {
			for _, v := range m {
				if str, ok := v.(string); ok {
					if str == "red" {
						return 0
					}
				}
			}
			sum := float64(0)
			for _, v := range m {
				sum += countJson(v)
			}
			return sum
		}

		if arr, ok := jsonObj.([]any); ok {
			sum := float64(0)
			for _, v := range arr {
				sum += countJson(v)
			}
			return sum
		}

		if _, ok := jsonObj.(string); ok {
			return 0
		}

		return 0
	}

	for _, line := range getInput("day12_input.txt") {

		var parsedJson map[string]any
		json.Unmarshal([]byte(line), &parsedJson)

		fmt.Println(countJson(parsedJson))
	}

}
