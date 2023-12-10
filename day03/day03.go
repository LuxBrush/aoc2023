package day03

import (
	"aoc2023/tools"
	"fmt"
)

var schematic = tools.ReadFileToString("day03/adventofcode_2023_day_3_input.txt")

func Part01(state string) {
	output := 0

	if state == "run" {
		output = tools.ProcessSchematic(schematic, 1)
	}

	fmt.Println("Day 03, Part 01:", output)
}
