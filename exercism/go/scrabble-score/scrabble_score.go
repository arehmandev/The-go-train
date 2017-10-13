package scrabble

import "strings"

const testVersion = 5

var mapboard = map[int]string{
	1:  "AEIOULNRST",
	2:  "DG",
	3:  "BCMP",
	4:  "FHVWY",
	5:  "K",
	8:  "JX",
	10: "QZ",
}

// Score - returns the scrabble score of a string
func Score(input string) (output int) {

	for _, element := range input {

		letter := strings.ToUpper(string(element))

		for key, value := range mapboard {
			if strings.Contains(value, letter) {
				output += key
			}
		}
	}

	return output
}
