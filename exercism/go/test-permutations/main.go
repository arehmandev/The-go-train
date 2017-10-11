package main

import "fmt"

func main() {

fmt.Println(permutations("test"))

}

func permutations(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}

	runes := []rune(input)
	subPermutations := permutations(string(runes[0:len(input) - 1]))

	result := []string{}
	for _, s := range subPermutations {
		result = append(result, merge([]rune(s), runes[len(input)-1])...)
	}

	return result
}

func merge(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}
