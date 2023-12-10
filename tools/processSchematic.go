package tools

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ProcessSchematic(schematic string, part int) int {
	output := 0

	grid := strings.Split(schematic, "\n")
	specialChars := `@#%&*+\-=/$`

	if part == 2 {
		specialChars = `*`
	}

	visited := map[[2]int]bool{}

	coordinates := getNumCooFromGrid(grid, specialChars, visited, part)

	output = getTotalFromCoo(grid, coordinates, part)

	return output
}

func getNumCooFromGrid(grid []string, specialChars string, visited map[[2]int]bool, part int) [][2]int {
	aKeys := [][2]int{}
	for x, row := range grid {
		for y, r := range row {
			if strings.ContainsRune(specialChars, r) {

				neighborInts := checkNeighborsForInt(grid, x, y, visited, true, 0, 0, part)

				aKeys = nil
				for k := range neighborInts {
					aKeys = append(aKeys, k)
				}

				slices.SortFunc(aKeys, func(a, b [2]int) int {
					if a[0] == b[0] {
						return a[1] - b[1]
					}
					return a[0] - b[0]
				})
			}
		}
	}
	return aKeys
}

func getTotalFromCoo(grid []string, coordinates [][2]int, part int) int {
	output := 0
	numString := ""
	numSlice := []int{}
	preY := 0

	for _, coordinate := range coordinates {

		cX := coordinate[0]
		cY := coordinate[1]
		if val, err := strconv.Atoi(string(string(grid[cX][cY]))); err == nil {
			diff := (preY + 1) - cY

			if diff > 0 || diff < 0 {
				if numString != "" {
					newNum, err := strconv.Atoi(numString)
					if err != nil {
						fmt.Println("Error in getting new number:", err)
					}

					numSlice = append(numSlice, newNum)
					numString = ""
					numString = numString + numToString(val)
					preY = cY
				} else {
					numString = numString + numToString(val)
					preY = cY
				}
			} else {
				numString = numString + numToString(val)
				preY = cY
			}

		}
	}

	if val, err := strconv.Atoi(string(string(numString))); err == nil {
		numSlice = append(numSlice, val)
		numString = ""
	}

	for _, num := range numSlice {
		output += num
	}

	return output
}

func checkNeighborsForInt(grid []string, x int, y int, visited map[[2]int]bool, searchAll bool, dx, dy int, part int) map[[2]int]bool {
	visited[[2]int{x, y}] = true

	directions := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range directions {
		if !searchAll && (dir.dx != 0 || dir.dy != dy) {
			continue // Skip directions other than left and right
		}
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
			if _, err := strconv.Atoi(string(grid[nx][ny])); err == nil {
				if !visited[[2]int{nx, ny}] {
					checkNeighborsForInt(grid, nx, ny, visited, false, dir.dx, dir.dy, part)
				}
			}
		}
	}
	return visited
}

func numToString(num int) string {
	return strconv.Itoa(num)
}