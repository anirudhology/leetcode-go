package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestLongestPalindrome(t *testing.T) {
	// Test case for null string
	if result := dynamic_programming.LongestPalindromicSubstring(""); result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}

	// Test case for single character string
	if result := dynamic_programming.LongestPalindromicSubstring("a"); result != "a" {
		t.Errorf("Expected 'a', got %s", result)
	}

	// Test case for string with all same characters
	if result := dynamic_programming.LongestPalindromicSubstring("aaaa"); result != "aaaa" {
		t.Errorf("Expected 'aaaa', got %s", result)
	}

	// Test case for string with a palindromic substring in the middle
	if result := dynamic_programming.LongestPalindromicSubstring("babad"); result != "aba" {
		t.Errorf("Expected 'aba', got %s", result)
	}

	// Test case for string with a palindromic substring at the end
	if result := dynamic_programming.LongestPalindromicSubstring("cbbd"); result != "bb" {
		t.Errorf("Expected 'bb', got %s", result)
	}

	// Test case for string with no palindromic substring longer than one character
	if result := dynamic_programming.LongestPalindromicSubstring("abc"); result != "c" {
		t.Errorf("Expected 'c', got %s", result)
	}

	// Test case for string with the entire string as a palindrome
	if result := dynamic_programming.LongestPalindromicSubstring("racecar"); result != "racecar" {
		t.Errorf("Expected 'racecar', got %s", result)
	}
}
