package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestChange_SimpleCase(t *testing.T) {
	result := dynamic_programming.CoinChangeII(5, []int{1, 2, 5})
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestChange_ZeroAmount(t *testing.T) {
	result := dynamic_programming.CoinChangeII(0, []int{1, 2, 5})
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestChange_NoCoins(t *testing.T) {
	result := dynamic_programming.CoinChangeII(5, []int{})
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestChange_NegativeAmount(t *testing.T) {
	result := dynamic_programming.CoinChangeII(-1, []int{1, 2, 5})
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestChange_NoSolution(t *testing.T) {
	result := dynamic_programming.CoinChangeII(3, []int{2})
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
