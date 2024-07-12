package strings_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/strings"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"A", true},
		{"1", true},
		{"racecar", true},
		{"A man, a plan, a canal: Panama", true},
		{"No 'x' in Nixon", true},
		{"hello", false},
		{"race a car", false},
		{"abc123", false},
		{"A1b2B1a", true},
		{"A1b2B3a", false},
		{"!@#$%^&*()_+", true},
		{"A man, a plan, a canal, Panama!", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := strings.ValidPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%s) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
