package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestClimbStairs(t *testing.T) {
	// Test case for n = 0
	if result := dynamic_programming.ClimbingStairs(0); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}

	// Test case for n = 1
	if result := dynamic_programming.ClimbingStairs(1); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	// Test case for n = 2
	if result := dynamic_programming.ClimbingStairs(2); result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}

	// Test case for n = 3
	if result := dynamic_programming.ClimbingStairs(3); result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}

	// Test case for n = 4
	if result := dynamic_programming.ClimbingStairs(4); result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}

	// Test case for n = 5
	if result := dynamic_programming.ClimbingStairs(5); result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}
