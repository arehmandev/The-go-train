package queenattack

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	testVersion = 2
)

var alpha = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

// CanQueenAttack -
func CanQueenAttack(queen1, queen2 string) (attack bool, err error) {

	// Are the pieces overlapping
	if queen2 == queen1 {
		return false, fmt.Errorf("Same square")
	}

	queenarray1 := strings.Split(queen1, "")
	// letter
	queen1x := (queenarray1[0])
	// number
	queen1y, err := strconv.Atoi(queenarray1[1])
	if err != nil {
		return false, err
	}

	// Are the pieces on the board
	queenarray2 := strings.Split(queen2, "")
	// letter
	queen2x := (queenarray2[0])
	// number
	queen2y, err := strconv.Atoi(queenarray2[1])
	if err != nil {
		return false, err
	}

	if queen1y > 8 || queen2y > 8 {
		return false, fmt.Errorf("Off of top of board (numbers)")
	}

	if queen1y < 1 || queen2y < 1 {
		return false, fmt.Errorf("Off of bottom of board (numbers)")
	}

	alphabetstring := strings.Join(alpha, "")
	if !strings.Contains(alphabetstring, queen1x) || !strings.Contains(alphabetstring, queen2x) {
		return true, fmt.Errorf("Off of board horizonally (letters)")
	}

	// If on same axis return true
	if queen1x == queen2x || queen1y == queen2y {
		return true, err
	}

	// If on diagonals return true
	verticaldiff := queen1y - queen2y
	fmt.Println(verticaldiff, ": vertical difference")

	horizontalindex1 := SliceIndex(len(alpha), func(i int) bool { return alpha[i] == queen1x })
	horizontalindex2 := SliceIndex(len(alpha), func(i int) bool { return alpha[i] == queen2x })
	horizontaldiff := horizontalindex1 - horizontalindex2
	fmt.Println(horizontaldiff, ": horizonal difference")

	// remove negative by squaring
	if verticaldiff*verticaldiff == horizontaldiff*horizontaldiff {
		return true, err
	}

	return false, err

}

// SliceIndex - Find the index of an element in a slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
