package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestCoinChange(t *testing.T) {
	// Test case for null coins array
	if result := dynamic_programming.CoinChange(nil, 5); result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}

	// Test case for empty coins array
	if result := dynamic_programming.CoinChange([]int{}, 5); result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}

	// Test case for negative amount
	if result := dynamic_programming.CoinChange([]int{1, 2, 5}, -5); result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}

	// Test case for amount 0
	if result := dynamic_programming.CoinChange([]int{1, 2, 5}, 0); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for amount that can't be reached
	if result := dynamic_programming.CoinChange([]int{2}, 3); result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}

	// Test case for amount with exact coin matches
	if result := dynamic_programming.CoinChange([]int{1, 2, 5}, 5); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	// Test case for amount with no exact match but can be reached
	if result := dynamic_programming.CoinChange([]int{1, 2, 5}, 11); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}

	// Test case for large amount
	if result := dynamic_programming.CoinChange([]int{1, 2, 5}, 100); result != 20 {
		t.Errorf("Expected 20, got %d", result)
	}
}
