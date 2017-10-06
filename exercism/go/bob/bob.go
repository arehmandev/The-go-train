package bob

import (
	"strings"
)

const testVersion = 3

// Hey responds like lackadaisical teenager.
func Hey(input string) (output string) {
	input = strings.TrimSpace(input)
	if input == strings.ToUpper(input) && strings.ContainsAny(input, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(input, "?") {
		return "Sure."
	}
	if input == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
