package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessCard(rawCardList string, part int) int {
	output := 0

	cardList := getCardList(rawCardList)

	for _, card := range cardList {
		if card != "" {
			output += readCard(card, part)
		}
	}

	return output
}

func getCardList(rawCardList string) []string {
	return strings.Split(rawCardList, "\n")
}

func getCardID(header string) (int, error) {
	output, err := strconv.Atoi(strings.Split(header, " ")[1])
	return output, err
}

func processBody(body string) ([]int, []int) {
	var numberList01 []int
	var numberList02 []int

	rawLists := strings.Split(body, "|")

	list01, err := getNumberList(rawLists[0])
	if err != nil {
		fmt.Println("Error getting list01:", err)
	}
	numberList01 = list01

	list02, err := getNumberList(rawLists[1])
	if err != nil {
		fmt.Println("Error getting list02:", err)
	}
	numberList02 = list02

	return numberList01, numberList02
}

func getNumberList(rawNumberList string) ([]int, error) {
	var numberList []int
	for _, numberString := range strings.Split(rawNumberList, " ") {
		if numberString != "" {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				return numberList, err
			}

			numberList = append(numberList, number)
		}
	}

	return numberList, nil
}

func readCard(rawCard string, part int) int {
	output := 0
	count := 0

	cardData := strings.Split(rawCard, ":")

	header := cardData[0]
	body := strings.Split(cardData[1], "\r")[0]

	id, err := getCardID(header)
	if err != nil {
		fmt.Println("Error getting card ID:", err)
	}

	fmt.Println(id)

	numberList01, numberList02 := processBody(body)

	if part == 1 {
		for _, numberL01 := range numberList01 {
			for _, numberL02 := range numberList02 {
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

	return output
}
