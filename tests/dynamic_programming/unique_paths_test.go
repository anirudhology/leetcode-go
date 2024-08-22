package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestUniquePaths_3x7(t *testing.T) {
	if result := dynamic_programming.UniquePaths(3, 7); result != 28 {
		t.Errorf("Expected 28, got %d", result)
	}
}

func TestUniquePaths_1x1(t *testing.T) {
	if result := dynamic_programming.UniquePaths(1, 1); result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestUniquePaths_2x2(t *testing.T) {
	if result := dynamic_programming.UniquePaths(2, 2); result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestUniquePaths_3x3(t *testing.T) {
	if result := dynamic_programming.UniquePaths(3, 3); result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestUniquePaths_7x3(t *testing.T) {
	if result := dynamic_programming.UniquePaths(7, 3); result != 28 {
		t.Errorf("Expected 28, got %d", result)
	}
}
