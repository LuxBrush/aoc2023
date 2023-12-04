package day02

import (
	"aoc2023/tools"
	"fmt"
)

var cubeGameRecords = tools.ReadFileToString("day02/adventofcode_2023_day_2_input.txt")
// var cubeGameRecords = tools.ReadFileToString("day02/demoPart01.txt")

func Part01(state string) {
	output := 2551

	if state == "run" {
		output = tools.ProcessGames(cubeGameRecords, 1)
	}

	fmt.Println("Day 02, Part 01:", output)
}

func Part02(state string) {
	output := 62811

	if state == "run" {
		output = tools.ProcessGames(cubeGameRecords, 2)	
	}

	fmt.Println("Day 02, Part 02:", output)
}