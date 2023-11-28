package day01

import "fmt"

func Part01(state string) {
	output := 0

	if state == "run" {
		output = 1
	}

	fmt.Println("Day 01, Part01:", output)
}