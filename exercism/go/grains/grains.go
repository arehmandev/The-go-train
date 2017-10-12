package grains

import "fmt"

const testVersion = 1

// Square - given the number of squares as input, give the output
func Square(input int) (output uint64, err error) {

	if input > 64 {
		err := fmt.Errorf("Out of range (0-64)")
		return 0, err
	}

	if input <= 0 {
		err := fmt.Errorf("Out of range (0-64)")
		return 0, err
	}

	val := 1
	for index := 0; index+1 < input; index++ {
		val = val * 2
	}

	output = uint64(val)
	return output, err
}

// Total - 2 to the power of 64
func Total() (output uint64) {

	a, err := Square(64)
	if err != nil {
		panic(err)
	}

	b := a + (a - 1) // remove 1 or uint overflows

	return (b)
}
