package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestMaxCoinsSingleBalloon(t *testing.T) {
	nums := []int{5}
	if result := dynamic_programming.BurstBalloons(nums); result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestMaxCoinsTwoBalloons(t *testing.T) {
	nums := []int{3, 1}
	if result := dynamic_programming.BurstBalloons(nums); result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestMaxCoinsMultipleBalloons(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	if result := dynamic_programming.BurstBalloons(nums); result != 167 {
		t.Errorf("Expected 167, got %d", result)
	}
}

func TestMaxCoinsNoBalloons(t *testing.T) {
	nums := []int{}
	if result := dynamic_programming.BurstBalloons(nums); result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestMaxCoinsLargeCase(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	if result := dynamic_programming.BurstBalloons(nums); result != 252 {
		t.Errorf("Expected 252, got %d", result)
	}
}
