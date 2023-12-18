package tools

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// directions represents the possible movements in a 2D grid.
var directions = []struct{ dx, dy int }{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

// ProcessSchematic processes a schematic string based on the specified part of the day.
// It returns the result according to the part parameter.
func ProcessSchematic(schematic string, part int) int {
	grid := strings.Split(schematic, "\n")
	specialChars := `@#%&*+\-=/$`
	visited, visited2 := map[[2]int]bool{}, map[[2]int]bool{}

	if part == 2 {
		specialChars = `*`
	}

	// Get the coordinates of special characters in the grid.
	coordinates := getNumCooFromGrid(grid, specialChars, visited, part)

	// Process the coordinates and return the total result.
	newList := processCoordinates(grid, coordinates, specialChars, visited2, part)
	return getTotal(newList, part)
}

// getNumCooFromGrid retrieves the coordinates of special characters in the grid.
func getNumCooFromGrid(grid []string, specialChars string, visited map[[2]int]bool, part int) [][2]int {
	coordinates := [][2]int{}
	for x, row := range grid {
		for y, r := range row {
			if strings.ContainsRune(specialChars, r) {

				// Get neighbors' coordinates and sort them.
				coordinates = [][2]int{}
				for c := range checkNeighborsForInt(grid, x, y, visited, true, 0, 0) {
					coordinates = append(coordinates, c)
				}
				sortCoordinates(coordinates)
			}
		}
	}
	return coordinates
}

// checkValidNeighbors checks if a special character has at least two valid neighbors.
func checkValidNeighbors(grid []string, x int, y int, visited map[[2]int]bool) bool {
	visited[[2]int{x, y}] = true
	numLocations, count, preX, preY := [][2]int{}, 0, 0, 0

	// Check neighbors in all directions.
	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
			if _, err := strconv.Atoi(string(grid[nx][ny])); err == nil {
				if !visited[[2]int{nx, ny}] {
					coo := [2]int{nx, ny}
					numLocations = append(numLocations, coo)
				}
			}
		}
	}

	// Count valid neighbors.
	for _, loc := range numLocations {
		lX := loc[0]
		lY := loc[1]
		if lY != preY+1 || lX != preX {
			count++
		}
		preX = lX
		preY = lY
	}

	return count >= 2
}

// checkNeighborsForInt explores neighbors to find integer values.
func checkNeighborsForInt(grid []string, x int, y int, visited map[[2]int]bool, searchAll bool, dx, dy int) map[[2]int]bool {
	visited[[2]int{x, y}] = true

	// Explore neighbors in all or specific directions.
	for _, dir := range directions {
		if !searchAll && (dir.dx != 0 || dir.dy != dy) {
			continue // Skip directions other than left and right
		}
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
			if _, err := strconv.Atoi(string(grid[nx][ny])); err == nil {
				if !visited[[2]int{nx, ny}] {
					checkNeighborsForInt(grid, nx, ny, visited, false, dir.dx, dir.dy)
				}
			}
		}
	}
	return visited
}

// numToString converts an integer to a string.
func numToString(num int) string {
	return strconv.Itoa(num)
}

// processCoordinates processes the coordinates and builds numbers based on neighboring integers.
func processCoordinates(grid []string, coordinates [][2]int, specialChars string, visited map[[2]int]bool, part int) [][]int {
	var numberList [][]int
	inVisited := visited
	for _, coordinate := range coordinates {
		cX, cY := coordinate[0], coordinate[1]

		if isSpecialChar(specialChars, rune(grid[cX][cY])) {
			if part != 1 {
				if checkValidNeighbors(grid, cX, cY, visited) {
					numberList = makeNumberList(grid, cX, cY, inVisited, numberList)
					inVisited = map[[2]int]bool{}
				}
			} else {
				numberList = makeNumberList(grid, cX, cY, inVisited, numberList)
				inVisited = map[[2]int]bool{}
			}
		}
	}

	return numberList
}

func makeNumberList(grid []string, cX, cY int, inVisited map[[2]int]bool, numberList [][]int) [][]int {
	var newCoordinates [][2]int

	// Get neighbors' coordinates and sort them.
	for value := range checkNeighborsForInt(grid, cX, cY, inVisited, true, 0, 0) {
		newCoordinates = append(newCoordinates, value)
	}
	sortCoordinates(newCoordinates)

	// Build a number from the coordinates.
	newNumber := buildNumber(grid, newCoordinates)

	numberList = append(numberList, newNumber)

	return numberList
}

// isSpecialChar checks if a character is a special character.
func isSpecialChar(specialChars string, ch rune) bool {
	return strings.ContainsRune(specialChars, ch)
}

// sortCoordinates sorts coordinates first by X and then by Y.
func sortCoordinates(coo [][2]int) {
	slices.SortFunc(coo, func(a, b [2]int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})
}

// buildNumber constructs a number from the grid based on the given coordinates.
func buildNumber(grid []string, coordinates [][2]int) []int {
	numString := ""
	numSlice := []int{}
	preY := 0
	preX := 0

	for _, coordinate := range coordinates {
		cX := coordinate[0]
		cY := coordinate[1]

		if val, err := strconv.Atoi(string(grid[cX][cY])); err == nil {
			diffY := (preY + 1) - cY
			diffX := preX - cX

			// Check if there is a gap between the coordinates.
			if diffY > 0 || diffY < 0 || diffX < 0 {
				if numString != "" {
					// Convert the accumulated string to a number.
					newNum, err := strconv.Atoi(numString)
					if err != nil {
						fmt.Println("Error in getting new number:", err)
					}

					numSlice = append(numSlice, newNum)
					numString = ""
					numString = numString + numToString(val)
					preY = cY
					preX = cX
				} else {
					numString = numString + numToString(val)
					preY = cY
					preX = cX
				}
			} else {
				numString = numString + numToString(val)
				preY = cY
				preX = cX
			}
		}
	}

	// Convert the remaining string to a number.
	if val, err := strconv.Atoi(string(numString)); err == nil {
		numSlice = append(numSlice, val)
		numString = ""
	}
	return numSlice
}

// getTotal calculates the total based on the numbers array and the specified part.
func getTotal(numbersArray [][]int, part int) int {
	var output int

	for _, value := range numbersArray {
		switch part {
		case 1:
			// Sum all numbers for part 1.
			for _, num := range value {
				output += num
			}
		case 2:
			// Multiply the first two numbers for part 2.
			num01 := value[0]
			num02 := value[1]
			output += num01 * num02
		}
	}

	return output
}
