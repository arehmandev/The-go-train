package diffsquares

import "fmt"

const testVersion = 1

func SquareOfSums(input int) int {

	var sum int

	for index := 0; index < input+1; index++ {
		sum += index
	}

	square := sum * sum
	return square
}

func SumOfSquares(input int) int {

	var sum int

	for index := 0; index < input+1; index++ {
		square := index * index
		fmt.Println(square)
		sum += square
	}

	return sum
}

func Difference(input int) int {
	return SquareOfSums(input) - SumOfSquares(input)
}
