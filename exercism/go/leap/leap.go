package leap

const testVersion = 3

func IsLeapYear(x int) bool {
	// Used switch instead of if
	switch {
	case x%4 == 0 && x%100 == 0 && x%400 == 0:
		return true
	case x%4 == 0 && x%100 != 0:
		return true
	default:
		return false
	}
}
