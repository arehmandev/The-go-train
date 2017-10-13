package letter

// FreqMap - type for the for frequencymap
type FreqMap map[rune]int

// Frequency - Find the frequency letters
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
