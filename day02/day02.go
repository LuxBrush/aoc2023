package day02

import "fmt"

func Part01(state string) {
	output := 0

	if state == "run" {
		output = 1
	}

	fmt.Println("Day 02, Part 01:", output)
}