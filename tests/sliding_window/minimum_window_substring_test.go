package sliding_window_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/sliding_window"
)

func TestMinimumWindowSubstring(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"abc", "xyz", ""},
		{"aA", "AA", ""},
		{"ADOBECODEBANC", "AABC", "ADOBECODEBA"},
		{"abcdefgh", "h", "h"},
		{"aabbcc", "abc", "abbc"},
		{"abc", "cba", "abc"},
	}

	for _, test := range tests {
		result := sliding_window.MinimumWindowSubstring(test.s, test.t)
		if result != test.expected {
			t.Errorf("For s='%s' and t='%s', expected '%s' but got '%s'", test.s, test.t, test.expected, result)
		}
	}
}
