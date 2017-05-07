package hamming

import "fmt"

const testVersion = 6

func Distance(a, b string) (int, error) {
	hammingdistance := 0
	switch {
	case len(a) == len(b):
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammingdistance++
			}
		}
		return hammingdistance, nil
	default:
		return 0, fmt.Errorf("Strings are of different length")
	}
}
