package main

import (
	"testing"
)

func TestCountCharOccurrences(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{
			input:    "hello world",
			expected: map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, ' ': 1, 'w': 1, 'r': 1, 'd': 1},
		},
		{
			input:    "go programming",
			expected: map[rune]int{'g': 3, 'o': 2, ' ': 1, 'p': 1, 'r': 2, 'a': 1, 'm': 2, 'i': 1, 'n': 1},
		},
	}

	for _, test := range tests {
		result := CountCharOccurrences(test.input)
		for char, expectedCount := range test.expected {
			if result[char] != expectedCount {
				t.Errorf("For input %q, expected count of %c to be %d, but got %d", test.input, char, expectedCount, result[char])
			}
		}
	}
}
