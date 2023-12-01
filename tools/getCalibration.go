package tools

import (
	// "fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func convertTextToNumbers(textNumbers []string) []string {
	output := []string{}
	for _, number := range textNumbers {

		wordToNumber := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}

		textToNumberString, ok := wordToNumber[number]
		if ok {
			output = append(output, textToNumberString)
		} else {
			output = append(output, number)
		}

	}
	return output
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		if m.Groups()[0].String() != "" {
			matches = append(matches, m.Groups()[0].String())
		}
		if m.Groups()[1].String() != "" {
			matches = append(matches, m.Groups()[1].String())
		}
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

func GetCalValue(rawString string, part int) (int, error) {
	output := 0
	for _, rawValue := range strings.Split(rawString, "\n") {
		if rawValue != "" {
			lookForNumbers := regexp.MustCompile(`\d`)
			foundNumbers := lookForNumbers.FindAllString(rawValue, -1)
			if part == 2 {
				test := regexp2.MustCompile(`(?=(one|two|three|four|five|six|seven|eight|nine))|\d`, 0)
				foundNumbers = regexp2FindAllString(test, rawValue)
			}

			foundNumbers = convertTextToNumbers(foundNumbers)

			// fmt.Println(foundNumbers)
			switch len(foundNumbers) {
			case 1:
				for _, number := range foundNumbers {
					outNumberText := ""
					outNumberText += number
					outNumberText += number
					num, err := strconv.Atoi(outNumberText)
					if err != nil {
						return output, err
					}
					// fmt.Println(num)
					output += num
				}
			default:
				newCalValueText := ""
				newCalValueText += foundNumbers[0]
				newCalValueText += foundNumbers[len(foundNumbers)-1]

				calValue, err := strconv.Atoi(newCalValueText)
				if err != nil {
					return output, err
				}
				// fmt.Println(calValue)
				output += calValue
			}

		}
	}
	return output, nil
}
