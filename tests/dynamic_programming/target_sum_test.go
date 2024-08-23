package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestFindTargetSumWays_SimpleCase(t *testing.T) {
	result := dynamic_programming.TargetSum([]int{1, 1, 1, 1, 1}, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestFindTargetSumWays_EmptyArray(t *testing.T) {
	result := dynamic_programming.TargetSum([]int{}, 3)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestFindTargetSumWays_NullArray(t *testing.T) {
	result := dynamic_programming.TargetSum(nil, 3)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestFindTargetSumWays_NoSolution(t *testing.T) {
	result := dynamic_programming.TargetSum([]int{1, 2, 3}, 7)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestFindTargetSumWays_SingleElementTargetMet(t *testing.T) {
	result := dynamic_programming.TargetSum([]int{5}, 5)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestFindTargetSumWays_SingleElementTargetNotMet(t *testing.T) {
	result := dynamic_programming.TargetSum([]int{5}, 3)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
