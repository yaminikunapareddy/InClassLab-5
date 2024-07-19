package main

// CountCharOccurrences returns a map with the count of each character in the input string
func CountCharOccurrences(s string) map[rune]int {
	// Create a map to store character counts
	counts := make(map[rune]int)

	// Iterate over each character in the string
	for _, char := range s {
		counts[char]++
	}

	return counts
}
