package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestRob(t *testing.T) {
	// Test case for single element slice
	if result := dynamic_programming.HouseRobber([]int{10}); result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}

	// Test case for two elements slice
	if result := dynamic_programming.HouseRobber([]int{10, 20}); result != 20 {
		t.Errorf("Expected 20, got %d", result)
	}

	// Test case for more elements
	if result := dynamic_programming.HouseRobber([]int{1, 2, 3, 1}); result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
	if result := dynamic_programming.HouseRobber([]int{2, 7, 9, 3, 1}); result != 12 {
		t.Errorf("Expected 12, got %d", result)
	}
	if result := dynamic_programming.HouseRobber([]int{1, 2, 9, 4, 5, 10}); result != 20 {
		t.Errorf("Expected 15, got %d", result)
	}
}
