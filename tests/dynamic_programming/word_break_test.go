package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestWordBreak(t *testing.T) {
	// Test case for null input string
	if result := dynamic_programming.WordBreak("", []string{"leet", "code"}); result != false {
		t.Errorf("Expected false, got %v", result)
	}

	// Test case for empty string
	if result := dynamic_programming.WordBreak("", []string{"leet", "code"}); result != false {
		t.Errorf("Expected false, got %v", result)
	}

	// Test case for null wordDict
	if result := dynamic_programming.WordBreak("leetcode", nil); result != false {
		t.Errorf("Expected false, got %v", result)
	}

	// Test case for empty wordDict
	if result := dynamic_programming.WordBreak("leetcode", []string{}); result != false {
		t.Errorf("Expected false, got %v", result)
	}

	// Test case for string that can be segmented
	if result := dynamic_programming.WordBreak("leetcode", []string{"leet", "code"}); result != true {
		t.Errorf("Expected true, got %v", result)
	}

	// Test case for string that cannot be segmented
	if result := dynamic_programming.WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}); result != false {
		t.Errorf("Expected false, got %v", result)
	}

	// Test case for string with overlapping words
	if result := dynamic_programming.WordBreak("applepenapple", []string{"apple", "pen"}); result != true {
		t.Errorf("Expected true, got %v", result)
	}

	// Test case for string with single character words
	if result := dynamic_programming.WordBreak("abcd", []string{"a", "b", "c", "d"}); result != true {
		t.Errorf("Expected true, got %v", result)
	}
}
