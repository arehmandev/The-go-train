package main

import "fmt"

func main() {

	initialnum := 3

	fmt.Println("U(n) = 5U(n-1)+1\nIf U(0) =", initialnum)

	array := []int{initialnum}

	for i := 0; i < 10; i++ {
		initialu := array[i]
		number := (5 * initialu) + 1
		array = append(array, number)
		fmt.Printf("%s(%v) = %v\n", "U", i+1, number)
	}

}
