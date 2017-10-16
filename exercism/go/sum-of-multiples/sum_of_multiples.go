package summultiples

const testVersion = 2

// SumMultiples -
func SumMultiples(limit int, divisors ...int) (total int) {

	var multipletotal []int

	// for each divisor
	for _, element := range divisors {

		// count up to limit
		for index := 0; index < limit; index++ {

			// if number is divisible by divisor, append to slice
			if index%element == 0 {
				multipletotal = append(multipletotal, index)
			}

		}
	}

	// filter duplicates from array and total
	multipletotal = removeDuplicates(multipletotal)

	for _, element := range multipletotal {
		total += element

	}

	return total

}

func removeDuplicates(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}
