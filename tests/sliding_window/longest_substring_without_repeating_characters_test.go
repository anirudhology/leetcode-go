package sliding_window_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/sliding_window"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdefg", 7},
		{"abba", 2},
		{"a", 1},
		{"au", 2},
		{"dvdf", 3},
		{"anviaj", 5},
		{"abcabcabc", 3},
		{"abbaacb", 3},
		{"tmmzuxt", 5},
	}

	for _, test := range tests {
		result := sliding_window.LongestSubstringWithoutRepeatingCharacters(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
