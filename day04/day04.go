package day04

import (
	"aoc2023/tools"
	"fmt"
)

var rawCardList = tools.ReadFileToString("day04/adventofcode_2023_day_4_input.txt")

func Part01(state string) {
	output := 17782

	if state == "run" {
		output = tools.ProcessCard(rawCardList, 1)
	}

	fmt.Println("Day 04, Part 01:", output)
}

func Part02(state string) {
	output := 8477787

	if state == "run" {
		output = tools.ProcessCard(rawCardList, 2)
	}

	fmt.Println("Day 02, Part 02:", output)
}