package secret

import (
	"fmt"
	"strconv"
	"strings"
)

const testVersion = 2

// Handshake -
func Handshake(input uint) (output []string) {

	// convert to int64 for next step
	u := int64(input)

	// Convert int to binary (string)
	binarystring := strconv.FormatInt(u, 2)
	fmt.Println(binarystring)

	binaryslice := strings.Split(binarystring, "")

	for i, element := range binaryslice {

		if i == len(binaryslice)-1-0 && element == "1" {
			output = append(output, "wink")
		}

		if i == len(binaryslice)-1-1 && element == "1" {
			output = append(output, "double blink")
		}
		if i == len(binaryslice)-1-2 && element == "1" {
			output = append(output, "close your eyes")
		}
		if i == len(binaryslice)-1-3 && element == "1" {
			output = append(output, "jump")
		}
	}

	output = reverse(output)

	// reverse again if 10000+
	if len(binaryslice) == 5 && binaryslice[4] == "1" {
		return reverse(output)
	}

	return output
}

func reverse(ss []string) []string {
	// fmt.Println("before:", ss)
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}

	fmt.Println("after:", ss)

	return ss
}
