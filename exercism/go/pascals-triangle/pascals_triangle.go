package pascal

const testVersion = 1

// Triangle returns Pascal's triangle with the given depth.
func Triangle(rows int) [][]int {
	triangle := [][]int{}

	triangle = append(triangle, []int{1})
	for i := 1; i < rows; i++ {
		triangle = append(triangle, nextRow(triangle[i-1]))
	}
	return triangle
}

func nextRow(prevRow []int) []int {
	current := make([]int, 0, len(prevRow)+1)

	current = append(current, 1)

	for i := 1; i < len(prevRow); i++ {
		current = append(current, prevRow[i-1]+prevRow[i])
	}

	current = append(current, 1)

	return current
}
