package series

import (
	"strings"
)

const testVersion = 2

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	slice := strings.Split(s, "")

	var substring []string

	for index := 0; index < len(slice)-n+1; index++ {
		test := slice[index : index+n]
		string := strings.Join(test, "")
		substring = append(substring, string)
	}

	return substring
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	slice := All(n, s)
	return slice[0]
}
