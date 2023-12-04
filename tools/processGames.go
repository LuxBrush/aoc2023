package tools

import (
	"fmt"
	"strconv"
	"strings"
)

// ProcessGames processes the cube games based on the specified part
// and returns the output.
func ProcessGames(cubeGames string, part int) int {
	output := 0

	// Define the cube game bag
	cubeGameBag := CubeBag{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// Split the raw games
	rawGames := strings.Split(cubeGames, "\n")
	for _, rawGame := range rawGames {
		if rawGame == "" {
			continue
		}
		// Get the ID from the raw game
		ID, err := getID(rawGame)
		if err != nil {
			fmt.Println("Error in getting ID:", err)
			continue
		}
		// Split the hands
		hands := strings.Split(strings.Split(rawGame, ":")[1], ";")
		if part == 1 {
			// Check if the game is valid and update the output
			if isValidGame(hands, cubeGameBag) {
				output += ID
			}
		} else if part == 2 {
			// Update the output based on the sum of the hands
			output += getSum(hands)
		}
	}
	return output
}

// getID extracts the ID from the raw game
func getID(rawGame string) (int, error) {
	idString := strings.Split(strings.Split(rawGame, ":")[0], " ")[1]
	return strconv.Atoi(idString)
}

// isValidGame checks if the game is valid based on the cube game bag
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

// getPower calculates the power of the possible cubes
func getPower(possibleCubes CubeBag) int {
	output := 1
	for _, number := range possibleCubes {
		output *= number
	}
	return output
}

// getSum calculates the sum of the hands
func getSum(hands []string) int {
	output := 0
	possibleCubes := CubeBag{}
	for _, hand := range hands {
		cubeInHand := strings.Split(hand, ",")
		for _, cube := range cubeInHand {
			numberOfCubes, color := getNumberOfCubesAndColor(cube)
			if numberOfCubes > possibleCubes[color] {
				possibleCubes[color] = numberOfCubes
			}
		}
	}
	output = getPower(possibleCubes)
	return output
}

// getNumberOfCubesAndColor extracts the number of cubes and the color from the cube
func getNumberOfCubesAndColor(cube string) (int, string) {
	parts := strings.Fields(cube)
	num, _ := strconv.Atoi(parts[0])
	return num, parts[1]
}
