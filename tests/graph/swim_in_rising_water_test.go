package graph_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestSwimInRisingWater(t *testing.T) {
	// Test case 1
	grid1 := [][]int{{0, 2}, {1, 3}}
	expected1 := 3
	if result := graph.SwimInRisingWater(grid1); result != expected1 {
		t.Errorf("expected %d, got %d", expected1, result)
	}

	// Test case 2
	grid2 := [][]int{
		{0, 1, 2, 3, 4},
		{24, 23, 22, 21, 5},
		{12, 13, 14, 15, 16},
		{11, 17, 18, 19, 20},
		{10, 9, 8, 7, 6},
	}
	expected2 := 16
	if result := graph.SwimInRisingWater(grid2); result != expected2 {
		t.Errorf("expected %d, got %d", expected2, result)
	}

	// Edge case: single cell
	grid3 := [][]int{{0}}
	expected3 := 0
	if result := graph.SwimInRisingWater(grid3); result != expected3 {
		t.Errorf("expected %d, got %d", expected3, result)
	}

	// Edge case: larger grid
	grid4 := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	expected4 := 8
	if result := graph.SwimInRisingWater(grid4); result != expected4 {
		t.Errorf("expected %d, got %d", expected4, result)
	}
}
