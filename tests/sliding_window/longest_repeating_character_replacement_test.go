package sliding_window_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/sliding_window"
)

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{"", 2, 0},         // Empty string
		{"A", 2, 1},        // Single character string
		{"ABCDE", 10, 5},   // k greater than length of string
		{"AAAA", 2, 4},     // No replacement needed
		{"AABABBA", 1, 4},  // General case
		{"ABAB", 2, 4},     // Mixed characters
		{"ABCDE", 0, 1},    // All different characters
		{"AABBB", 1, 4},    // Mixed characters with k
		{"ABAA", 0, 2},     // k equal to 0
		{"AABBA", 2, 5},    // All characters can be same with k replacements
		{"AAABBC", 1, 4},   // Partial replacement
		{"ABABABAB", 2, 5}, // Alternating characters with k replacements
		{"AABABBA", 2, 5},  // Multiple replacements
	}

	for _, tt := range tests {
		result := sliding_window.CharacterReplacement(tt.s, tt.k)
		if result != tt.expected {
			t.Errorf("CharacterReplacement(%v, %d) = %d; expected %d", tt.s, tt.k, result, tt.expected)
		}
	}
}
