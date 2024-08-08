package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestHouseRobberII(t *testing.T) {
	// Test case for empty slice
	if result := dynamic_programming.HouseRobberII([]int{}); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for single element slice
	if result := dynamic_programming.HouseRobberII([]int{10}); result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}

	// Test case for two elements slice
	if result := dynamic_programming.HouseRobberII([]int{10, 20}); result != 20 {
		t.Errorf("Expected 20, got %d", result)
	}

	// Test case for more elements
	if result := dynamic_programming.HouseRobberII([]int{2, 3, 2}); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
	if result := dynamic_programming.HouseRobberII([]int{1, 2, 3, 1}); result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
	if result := dynamic_programming.HouseRobberII([]int{100, 2, 3, 100}); result != 103 {
		t.Errorf("Expected 103, got %d", result)
	}
}
