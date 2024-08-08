package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestPalindromicSubstrings(t *testing.T) {
	// Test case for null string
	if result := dynamic_programming.PalindromicSubstrings(""); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for single character string
	if result := dynamic_programming.PalindromicSubstrings("a"); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	// Test case for string with all same characters
	if result := dynamic_programming.PalindromicSubstrings("aaaa"); result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}

	// Test case for string with mixed characters
	if result := dynamic_programming.PalindromicSubstrings("abc"); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}

	// Test case for string with palindromic substrings
	if result := dynamic_programming.PalindromicSubstrings("aaa"); result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
