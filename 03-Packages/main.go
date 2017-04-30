package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	min       = 1
	max       = 10
	randomnum = 64
)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func randomnumber() {
	Myrand := random(min, max)
	fmt.Println("Here is a random number between", min, "and", max, ":", Myrand)
	fmt.Println("Here is the value of Pi (15 d.p):", math.Pi)
}

func squareroot() {
	sqrtnumber := math.Sqrt(randomnum)
	fmt.Printf("Here is the square root of the number %v : %g \n", randomnum, sqrtnumber)
}

// Generates random number between 1 (min) and 10 (max) with no duplicate results
//Functions 4/17

func add(a, b int) int {
	return a + b
}

func reverseletters(c, d string) (string, string) {
	return d, c
}

func main() {
	randomnumber()
	squareroot()
	a := 5
	b := 9
	word1, word2 := "Hello", "Test"
	fmt.Println("Here is the total of 2 numbers:", a, "+", b, "=", add(a, b))
	fmt.Println(reverseletters(word1, word2))
}
