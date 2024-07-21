package sliding_window_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/sliding_window"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		s1     string
		s2     string
		result bool
	}{
		{"", "anything", true},         // s1 is empty string
		{"anything", "", false},        // s2 is empty string
		{"longstring", "short", false}, // s1 is longer than s2
		{"abc", "cba", true},           // exact match
		{"ab", "eidbaooo", true},       // permutation in the middle of s2
		{"ab", "abdcba", true},         // permutation at the start of s2
		{"ab", "eidboaooab", true},     // permutation at the end of s2
		{"ab", "eidboaoo", false},      // no permutation
		{"aabb", "bbbaaabbcc", true},   // repeated characters in s1
		{"abc", "aabbcc", false},       // repeated characters in s2
		{"abc", "defghijkl", false},    // different characters
	}

	for _, test := range tests {
		result := sliding_window.CheckInclusion(test.s1, test.s2)
		if result != test.result {
			t.Errorf("checkInclusion(%q, %q) = %v; want %v", test.s1, test.s2, result, test.result)
		}
	}
}
