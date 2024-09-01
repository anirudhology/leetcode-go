package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestNumDistinctEmptyStrings(t *testing.T) {
	result := dynamic_programming.DistinctSubsequences("", "")
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestNumDistinctEmptyTarget(t *testing.T) {
	if result := dynamic_programming.DistinctSubsequences("abc", ""); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestNumDistinctEmptySource(t *testing.T) {
	if result := dynamic_programming.DistinctSubsequences("", "abc"); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestNumDistinctSameStrings(t *testing.T) {
	if result := dynamic_programming.DistinctSubsequences("abc", "abc"); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestNumDistinctGeneralCases(t *testing.T) {
	if result := dynamic_programming.DistinctSubsequences("babgbag", "bag"); result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
	if result := dynamic_programming.DistinctSubsequences("rabbbit", "rabbit"); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
	if result := dynamic_programming.DistinctSubsequences("abc", "def"); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
