package etl

import (
	"fmt"
	"strings"
)

const testVersion = 1

// Transform - do the `Transform` step of an Extract-Transform-Load.
func Transform(input map[int][]string) (output map[string]int) {

	newmap := make(map[string]int)

	fmt.Println(newmap)

	for i := range input {

		list := input[i]

		for _, element := range list {
			newmap[strings.ToLower(element)] = i
		}
	}

	return newmap
}
