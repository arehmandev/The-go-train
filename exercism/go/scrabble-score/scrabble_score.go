package scrabble

import "strings"

const testVersion = 5

var onepoint = []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
var twopoint = []string{"D", "G"}
var threepoint = []string{"B", "C", "M", "P"}
var fourpoint = []string{"F", "H", "V", "W", "Y"}
var fivepoint = []string{"K"}
var eightpoint = []string{"J", "X"}
var tenpoint = []string{"Q", "Z"}

// Score - returns the scrabble score of a string
func Score(input string) (output int) {

	scoreboard := createscoreboard()
	stringslice := strings.Split(input, "")

	for _, element := range stringslice {
		output += scoreboard[element]
	}

	return output
}

func createscoreboard() map[string]int {

	outputmap := make(map[string]int)

	for _, letter := range onepoint {
		outputmap[letter] = 1
		outputmap[strings.ToLower(letter)] = 1
	}

	for _, letter := range twopoint {
		outputmap[letter] = 2
		outputmap[strings.ToLower(letter)] = 2
	}

	for _, letter := range threepoint {
		outputmap[letter] = 3
		outputmap[strings.ToLower(letter)] = 3
	}

	for _, letter := range fourpoint {
		outputmap[letter] = 4
		outputmap[strings.ToLower(letter)] = 4
	}

	for _, letter := range fivepoint {
		outputmap[letter] = 5
		outputmap[strings.ToLower(letter)] = 5
	}

	for _, letter := range eightpoint {
		outputmap[letter] = 8
		outputmap[strings.ToLower(letter)] = 8
	}

	for _, letter := range tenpoint {
		outputmap[letter] = 10
		outputmap[strings.ToLower(letter)] = 10
	}

	return outputmap
}
