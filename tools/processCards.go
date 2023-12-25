package tools

import (
	"fmt"
	"strconv"
	"strings"
)

// ProcessCard takes a rawCardList string and a part int and returns an int
func ProcessCard(rawCardList string, part int) int {
	output := 0

	// Get the card list from the rawCardList string
	cardList := getCardList(rawCardList)

	// Get the card map from the card list
	cardMap := getCardMap(cardList)

	// Get the cardCount from the card map
	cardCount := readCardMap(cardMap, part)

	// Iterate through the card list
	for _, card := range cardList {
		if card != "" {
			// Read the card and add the output
			output += readCard(card, part, cardList)
		}
	}

	if part == 2 {
		output = cardCount
	}

	return output
}

// Get the card list from the rawCardList string
func getCardList(rawCardList string) []string {
	return strings.Split(rawCardList, "\n")
}

// Get the card ID from the header string
func getCardID(header string) int {
	output := 0
	for _, part := range strings.Split(header, " ") {
		if value, err := strconv.Atoi(part); err == nil {
			output = value
		}
	}

	return output
}

// ProcessBody takes a body string and returns two int slices
func processBody(body string) ([]int, []int) {
	var numberList01 []int
	var numberList02 []int

	// Split the body string into two lists
	rawLists := strings.Split(body, "|")

	// Get the number list from the first list
	list01, err := getNumberList(rawLists[0])
	if err != nil {
		fmt.Println("Error getting list01:", err)
	}
	numberList01 = list01

	// Get the number list from the second list
	list02, err := getNumberList(rawLists[1])
	if err != nil {
		fmt.Println("Error getting list02:", err)
	}
	numberList02 = list02

	// Return the two int slices
	return numberList01, numberList02
}

// Get the number list from the rawNumberList string
func getNumberList(rawNumberList string) ([]int, error) {
	var numberList []int
	// Iterate through the rawNumberList string
	for _, numberString := range strings.Split(rawNumberList, " ") {
		if numberString != "" {
			// Convert the number string to an int
			number, err := strconv.Atoi(numberString)
			if err != nil {
				return numberList, err
			}

			// Append the int to the numberList
			numberList = append(numberList, number)
		}
	}

	// Return the numberList and nil
	return numberList, nil
}

// Read the card from the rawCard string and the part int, and return an int
func readCard(rawCard string, part int, cardList []string) int {
	output := 0
	count := 0

	// Split the rawCard string into two parts
	cardData := strings.Split(rawCard, ":")

	// Get the body from the second part
	body := strings.Split(cardData[1], "\r")[0]

	// Process the body
	numberList01, numberList02 := processBody(body)

	// If the part is 1, iterate through the two number lists
	if part == 1 {
		for _, numberL01 := range numberList01 {
			for _, numberL02 := range numberList02 {
				// If the two numbers are equal, increment the output and the count
				if numberL01 == numberL02 {
					if count == 0 {
						output++
						count++
					} else {
						output = output * 2
						count++
					}
				}
			}
		}
	}

	// Return the output
	return output
}

// Get the card map from the card list
func getCardMap(cardList []string) map[int]cardStruct {

	// Create an empty map
	cardMap := map[int]cardStruct{}

	// Iterate through the card list
	for _, card := range cardList {
		if card != "" {
			// Split the card string into two parts
			cardData := strings.Split(card, ":")

			// Get the header from the first part
			header := cardData[0]
			// Get the body from the second part
			body := cardData[1]
			// Initialize the count
			winList := []int{}

			// Get the card ID from the header
			id := getCardID(header)

			// Process the body
			numberList01, numberList02 := processBody(body)

			// Add the card struct to the map
			cardMap[id] = cardStruct{numberList01, numberList02, id, winList}

		}
	}

	// Return the map
	return cardMap
}

// Function to read card map and return the output
func readCardMap(cardMap map[int]cardStruct, part int) int {
	// Initialize output and winList
	output := 0
	var winList []int
	outCount := 0
	storeIndex := 0

	// Iterate through the card map
	for index := 1; index <= len(cardMap); index++ {
		card := cardMap[index]

		// If the card has no winList, get the winList from the card id
		if len(card.winList) == 0 {

			winList = getWinList(card.id, card)
			card.winList = winList
			cardMap[index] = card
		}

		// Iterate through the winList and store the index
		outCount = len(cardMap) + 1
		for _, win := range card.winList {

			cardMap[outCount] = cardMap[win]
			outCount++

		}
		storeIndex = index
	}
	output = storeIndex

	return output
}


func getWinList(index int, card cardStruct) []int {
	// Initialize an empty winList
	winList := []int{}
	// Iterate through nl01
	for _, n01 := range card.nl01 {
		// Iterate through nl02
		for _, n02 := range card.nl02 {
			// If n01 is equal to n02
			if n01 == n02 {
				// Increment the index
				index++
				// Append the index to winList
				winList = append(winList, index)
			}
		}
	}
	// Return the winList of indexes
	return winList
}
