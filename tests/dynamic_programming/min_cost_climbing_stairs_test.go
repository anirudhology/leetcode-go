package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestMinCostClimbingStairs(t *testing.T) {
	// Test case for two elements slice
	if result := dynamic_programming.MinCostClimbingStairs([]int{10, 15}); result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}

	// Test case for more elements
	if result := dynamic_programming.MinCostClimbingStairs([]int{10, 15, 20}); result != 15 {
		t.Errorf("Expected 15, got %d", result)
	}
	if result := dynamic_programming.MinCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}); result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
