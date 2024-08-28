package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestMaxSubArraySingleElement(t *testing.T) {
	result := dynamic_programming.MaximumSubArray([]int{5})
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestMaxSubArrayAllNegative(t *testing.T) {
	result := dynamic_programming.MaximumSubArray([]int{-1, -2, -3, -4})
	if result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}
}

func TestMaxSubArrayMixed(t *testing.T) {
	result := dynamic_programming.MaximumSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestMaxSubArrayAllPositive(t *testing.T) {
	result := dynamic_programming.MaximumSubArray([]int{1, 2, 3, 4, 5})
	if result != 15 {
		t.Errorf("Expected 15, got %d", result)
	}
}
