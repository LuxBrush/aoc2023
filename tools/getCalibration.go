package tools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetCalValue(rawString string) (int, error) {
	output := 0
	for _, rawValue := range strings.Split(rawString, "\n") {
		if rawValue != "" {
			lookIn := regexp.MustCompile(`\d`)
			rawNumbers := lookIn.FindAllString(rawValue, -1)

			fmt.Println(rawNumbers)
			switch len(rawNumbers) {
			case 1:
				for _, number := range rawNumbers {
					outNumberText := ""
					outNumberText += number
					outNumberText += number
					num, err := strconv.Atoi(outNumberText)
					if err != nil {
						return output, err
					}
					fmt.Println(num)
					output += num
				}
			default:
				newCalValueText := ""
				newCalValueText += rawNumbers[0]
				newCalValueText += rawNumbers[len(rawNumbers)-1]
				
				calValue, err := strconv.Atoi(newCalValueText)
				if err != nil {
					return output, err
				}
				fmt.Println(calValue)
				output += calValue
			}
			
		}
	}
	return output, nil
}
