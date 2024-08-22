package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestLCS_SimpleCase(t *testing.T) {
	if result := dynamic_programming.LongestCommonSubsequence("abcde", "ace"); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestLCS_NoCommonSubsequence(t *testing.T) {
	if result := dynamic_programming.LongestCommonSubsequence("abc", "def"); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestLCS_EmptyStrings(t *testing.T) {
	if result := dynamic_programming.LongestCommonSubsequence("", ""); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestLCS_OneEmptyString(t *testing.T) {
	if result := dynamic_programming.LongestCommonSubsequence("abcde", ""); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestLCS_FullMatch(t *testing.T) {
	if result := dynamic_programming.LongestCommonSubsequence("abcde", "abcde"); result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestLCS_ReversedStrings(t *testing.T) {
	if result := dynamic_programming.LongestCommonSubsequence("abcde", "edcba"); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
