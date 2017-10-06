package pangram

import (
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

const testVersion = 2

func IsPangram(upperword string) bool {

	lowerword := strings.ToLower(upperword)

	for _, v := range alphabet {
		if strings.ContainsRune(lowerword, v) == false {
			return false
		}
	}

	return true
}
