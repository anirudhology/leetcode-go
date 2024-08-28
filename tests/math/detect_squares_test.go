package math_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/math"
)

func TestAddAndCount(t *testing.T) {
	ds := math.DetectSquaresConstructor()

	// Adding points
	ds.Add([]int{1, 2})
	ds.Add([]int{2, 1})
	ds.Add([]int{1, 0})
	ds.Add([]int{0, 1})

	// No squares should exist initially
	if ds.Count([]int{1, 1}) != 0 {
		t.Errorf("Expected 0, got %d", ds.Count([]int{1, 1}))
	}

	// Adding points to form a square
	ds.Add([]int{2, 2})
	if ds.Count([]int{1, 1}) != 1 {
		t.Errorf("Expected 1, got %d", ds.Count([]int{1, 1}))
	}

	// Adding a point to form a second square
	ds.Add([]int{0, 0})
	if ds.Count([]int{1, 1}) != 2 {
		t.Errorf("Expected 2, got %d", ds.Count([]int{1, 1}))
	}
}

func TestNoSquares(t *testing.T) {
	ds := math.DetectSquaresConstructor()

	// No points added, no squares should exist
	if ds.Count([]int{0, 0}) != 0 {
		t.Errorf("Expected 0, got %d", ds.Count([]int{0, 0}))
	}
}

func TestMultipleSamePoints(t *testing.T) {
	ds := math.DetectSquaresConstructor()

	// Adding the same point multiple times
	ds.Add([]int{0, 0})
	ds.Add([]int{0, 0})
	ds.Add([]int{1, 1})
	ds.Add([]int{1, 1})
	ds.Add([]int{0, 1})
	ds.Add([]int{1, 0})

	// Two squares are formed
	if ds.Count([]int{0, 0}) != 2 {
		t.Errorf("Expected 2, got %d", ds.Count([]int{0, 0}))
	}
}
