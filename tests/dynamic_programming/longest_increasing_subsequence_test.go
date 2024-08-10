package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestLengthOfLIS(t *testing.T) {
	// Test case for null input
	if result := dynamic_programming.LongestIncreasingSubsequence(nil); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for empty array
	if result := dynamic_programming.LongestIncreasingSubsequence([]int{}); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for single element array
	if result := dynamic_programming.LongestIncreasingSubsequence([]int{10}); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	// Test case for array with all elements the same
	if result := dynamic_programming.LongestIncreasingSubsequence([]int{5, 5, 5, 5}); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	// Test case for array with strictly increasing elements
	if result := dynamic_programming.LongestIncreasingSubsequence([]int{1, 2, 3, 4}); result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}

	// Test case for array with strictly decreasing elements
	if result := dynamic_programming.LongestIncreasingSubsequence([]int{5, 4, 3, 2, 1}); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	// Test case for array with mixed increasing and decreasing elements
	if result := dynamic_programming.LongestIncreasingSubsequence([]int{10, 9, 2, 5, 3, 7, 101, 18}); result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}

	// Test case for array with random elements
	if result := dynamic_programming.LongestIncreasingSubsequence([]int{3, 10, 2, 1, 20}); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}
