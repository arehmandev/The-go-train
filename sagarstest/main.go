// Find the number:
//
// It's less than 100.
//for loop to 100

// It's one more than a multiple of 3.
// n-1 % 3 == 0

// If you multiply it by 5, the answer is greater than 100.
// n * 5 > 100

// The sum of its digits is prime.
// split digits, isPrime(digit1 + digit2) == true

// Exactly one of the digits is prime.
// split digits, isPrime(digits[1]) == true || isPrime(digits[2])

// If you reverse its digits, you get a prime number.
// split digits, reverse slice, join, isPrime(reversednum)

// It has exactly four factors.
// len(findFactors(number))

package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	for i := 0; 100 > i; i++ {

		// It's one more than a multiple of 3.
		if (i-1)%3 != 0 {
			continue
		}

		// If you multiply it by 5, the answer is greater than 100.
		if i*5 < 100 {
			// fmt.Println(i, "- multiply it by 5, the answer is not greater than 100.")
			continue
		}

		digit1string := string(strconv.Itoa(i)[0])
		digit2string := string(strconv.Itoa(i)[1])

		digit1, err := strconv.Atoi(digit1string)
		digit2, err := strconv.Atoi(digit2string)
		if digit2 == 0 {
			continue
		}

		if err != nil {
			panic(err)
		}

		// The sum of its digits is prime.
		if isPrime(digit1+digit2) == false {
			// fmt.Println(i, "- Digits don't total to a prime number")
			continue
		}
		// The reverse of its digits is prime.
		reversed, err := strconv.Atoi(strings.Join([]string{digit2string, digit1string}, ""))
		if err != nil {
			log.Fatal(err)
		}
		if isPrime(reversed) == false {
			// fmt.Println(i, "- Digits don't total to a prime number")
			continue
		}

		// It has exactly four factors.
		if getfactors(i) != 4 {
			continue
		}
		fmt.Print(i, " - is remaining \n")
	}
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func getfactors(value int) (output int) {

	for i := 1; i < value; i++ {
		if value%i == 0 {
			output++
		}
	}

	return output + 1
}
