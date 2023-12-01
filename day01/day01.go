package day01

import (
	"aoc2023/tools"
	"fmt"
)

var calibrationNumberList = tools.ReadFileToString("day01/adventofcode_2023_day_1_input.txt")
// var calibrationNumberList = tools.ReadFileToString("day01/demo.txt")

func Part01(state string) {
	output := 0

	if state == "run" {
		calSum, err := tools.GetCalValue(calibrationNumberList)
		if err != nil {
			fmt.Println("GetCalValue has had an error:", err)
		}
		output = calSum
	}

	fmt.Println("Day 01, Part 01:", output)
}
