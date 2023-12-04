package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessGames(cubeGames string) int {
	output := 0

	cubeGameBag := CubeBag{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games := strings.Split(cubeGames, "\n")
	for _, rawGame := range games {
		if rawGame == "" {
			continue
		}
		ID, err := getID(rawGame)
		if err != nil {
			fmt.Println("Error in getting ID:", err)
			continue
		}
		hands := strings.Split(strings.Split(rawGame, ":")[1], ";")
		if isValidGame(hands, cubeGameBag) {
			output += ID
		}
	}
	return output
}

func getID(rawGame string) (int, error) {
	idString := strings.Split(strings.Split(rawGame, ":")[0], " ")[1]
	return strconv.Atoi(idString)
}

func isValidGame(hands []string, cubeGameBag CubeBag) bool {
	for _, hand := range hands {
		cubeInHand := strings.Split(hand, ",")
		for _, cube := range cubeInHand {
			numberOfCubes, color := getNumberOfCubesAndColor(cube)
			if numberOfCubes > cubeGameBag[color] {
				return false
			}
		}
	}
	return true
}

func getNumberOfCubesAndColor(cube string) (int, string) {
	parts := strings.Fields(cube)
	fmt.Println(parts)
	num, _ := strconv.Atoi(parts[0])
	return num, parts[1]
}