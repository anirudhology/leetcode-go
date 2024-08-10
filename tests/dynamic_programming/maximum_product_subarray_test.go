package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestMaxProduct(t *testing.T) {
	// Test case for null input
	if result := dynamic_programming.MaximumProductSubarray(nil); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for empty array
	if result := dynamic_programming.MaximumProductSubarray([]int{}); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for single positive element
	if result := dynamic_programming.MaximumProductSubarray([]int{2}); result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}

	// Test case for single negative element
	if result := dynamic_programming.MaximumProductSubarray([]int{-2}); result != -2 {
		t.Errorf("Expected -2, got %d", result)
	}

	// Test case for array with both positive and negative elements
	if result := dynamic_programming.MaximumProductSubarray([]int{2, 3, -2, 4}); result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}

	// Test case for array with a zero
	if result := dynamic_programming.MaximumProductSubarray([]int{-2, 0, -1, 6, -4}); result != 24 {
		t.Errorf("Expected 24, got %d", result)
	}

	// Test case for all negative elements
	if result := dynamic_programming.MaximumProductSubarray([]int{-1, -3, -4}); result != 12 {
		t.Errorf("Expected 12, got %d", result)
	}

	// Test case for all positive elements
	if result := dynamic_programming.MaximumProductSubarray([]int{1, 2, 3, 4, 5, 6}); result != 720 {
		t.Errorf("Expected 720, got %d", result)
	}

	// Test case for array with a mix of positive, negative, and zero elements
	if result := dynamic_programming.MaximumProductSubarray([]int{-2, 0, -1, -3, 10, 6}); result != 180 {
		t.Errorf("Expected 180, got %d", result)
	}
}
