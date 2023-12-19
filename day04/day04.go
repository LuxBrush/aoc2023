package day04

import (
	"aoc2023/tools"
	"fmt"
)

var rawCardList = tools.ReadFileToString("day04/adventofcode_2023_day_4_input.txt")

func Part01(state string) {
	output := 0

	if state == "run" {
		output = tools.ProcessCard(rawCardList, 1)
	}

	fmt.Println("Day 04, Part 01:", output)
}