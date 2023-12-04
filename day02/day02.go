package day02

import (
	"aoc2023/tools"
	"fmt"
)

var cubeGameRecords = tools.ReadFileToString("day02/adventofcode_2023_day_2_input.txt")
// var cubeGameRecords = tools.ReadFileToString("day02/demoPart01.txt")

func Part01(state string) {
	output := 0

	if state == "run" {
		output = tools.ProcessGames(cubeGameRecords)
	}

	fmt.Println("Day 02, Part 01:", output)
}