package greedy_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/greedy"
)

func TestIsNStraightHandValid(t *testing.T) {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	groupSize := 3
	result := greedy.HandOfStraights(hand, groupSize)
	if !result {
		t.Errorf("Expected true but got false")
	}
}

func TestIsNStraightHandInvalid(t *testing.T) {
	hand := []int{1, 2, 3, 4, 5}
	groupSize := 4
	result := greedy.HandOfStraights(hand, groupSize)
	if result {
		t.Errorf("Expected false but got true")
	}
}

func TestIsNStraightHandEmptyHand(t *testing.T) {
	hand := []int{}
	groupSize := 1
	result := greedy.HandOfStraights(hand, groupSize)
	if !result {
		t.Errorf("Expected true but got false")
	}
}

func TestIsNStraightHandSingleGroup(t *testing.T) {
	hand := []int{1, 2, 3}
	groupSize := 3
	result := greedy.HandOfStraights(hand, groupSize)
	if !result {
		t.Errorf("Expected true but got false")
	}
}
