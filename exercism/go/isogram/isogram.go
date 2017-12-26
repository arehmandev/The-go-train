package isogram

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// IsIsogram - Returns boolean depending on if string is an isogram
func IsIsogram(input string) bool {
	alphabet := make(map[string]bool)
	capital := strings.ToUpper(input)
	alphanum := RemoveNonAlphaNumeric(capital)

	word := strings.Split(alphanum, "")

	for _, value := range word {
		_, ok := alphabet[value]
		if ok == false {
			alphabet[value] = true
		} else {
			return false
		}
	}

	fmt.Println(alphabet)

	return true
}

func RemoveNonAlphaNumeric(input string) string {

	// Make a Regex to say we only want
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(input, "")
	return processedString
}
