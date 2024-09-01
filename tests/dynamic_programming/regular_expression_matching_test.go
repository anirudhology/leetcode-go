package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s, p   string
		result bool
	}{
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"mississippi", "mis*is*p*.", false},
		{"aab", "c*a*b", true},
		{"ab", "a", false},
		{"", ".*", true},
		{"", "a*", true},
		{"abcd", "d*", false},
		{"aaa", "a*a", true},
		{"aaa", "ab*a", false},
	}

	for _, test := range tests {
		if got := dynamic_programming.RegularExpressionMatching(test.s, test.p); got != test.result {
			t.Errorf("isMatch(%v, %v) = %v; want %v", test.s, test.p, got, test.result)
		}
	}
}
