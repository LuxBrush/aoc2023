package day01

import (
	"aoc2023/tools"
	"fmt"
)

var calibrationNumberList = tools.ReadFileToString("day01/adventofcode_2023_day_1_input.txt")
// var calibrationNumberList = tools.ReadFileToString("day01/demo.txt")
// var calibrationNumberList = tools.ReadFileToString("day01/demoPart02.txt")

func Part01(state string) {
	output := 57346

	if state == "run" {
		calSum, err := tools.GetCalValue(calibrationNumberList, 1)
		if err != nil {
			fmt.Println("GetCalValue has had an error:", err)
		}
		output = calSum
	}

	fmt.Println("Day 01, Part 01:", output)
}

func Part02(state string) {
	output := 57345

	if state == "run" {
		calSum, err := tools.GetCalValue(calibrationNumberList, 2)
		if err != nil {
			fmt.Println("GetCalValue has had an error:", err)
		}
		output = calSum
	}

	fmt.Println("Day 01, Part 02:", output)
}