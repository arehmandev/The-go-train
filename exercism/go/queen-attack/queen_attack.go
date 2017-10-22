package queenattack

import (
	"errors"
	"math"
)

const testVersion = 2

func CanQueenAttack(w string, b string) (bool, error) {
	if len(w) < 2 || len(b) < 2 ||
		w[0] < 'a' || b[0] < 'a' ||
		w[0] > 'h' || b[0] > 'h' ||
		w[1] < '1' || b[1] < '1' ||
		w[1] > '8' || b[1] > '8' ||
		(w[0] == b[0] && w[1] == b[1]) {
		return false, errors.New("invalid position")
	}

	wx := float64(w[0] - 96)
	wy := float64(w[1] - 48)

	bx := float64(b[0] - 96)
	by := float64(b[1] - 48)

	if wx == bx || wy == by || (math.Abs(wx-bx) == math.Abs(wy-by)) {
		return true, nil
	}

	return false, nil
}
