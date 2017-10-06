package triangle

import "math"

const testVersion = 3

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Pick values for the following identifiers used by the test program.

const (
	NaT = iota // not a triangle
	Sca        // scalene
	Iso        // isosceles
	Equ        // equilateral
)

// Organize your code for readability.

func KindFromSides(a, b, c float64) Kind {

	if !isValid(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}
	if (a == b) || (a == c) || (b == c) {
		return Iso
	}
	if a != b && b != c && a != c {
		return Sca
	}

	return NaT
}

func isValid(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}

	if a == math.Inf(1) || b == math.Inf(1) || c == math.Inf(1) {
		return false
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a+b < c || a+c < b || b+c < a {
		return false
	}

	return true
}
