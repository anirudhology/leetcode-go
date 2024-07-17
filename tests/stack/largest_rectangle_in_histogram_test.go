package stack_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/stack"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{}, 0},                        // Test case with empty input
		{[]int{2, 1, 5, 6, 2, 3}, 10},       // Test case with mixed heights
		{[]int{2, 4}, 4},                    // Test case with two elements
		{[]int{2, 3, 4, 5, 6, 7}, 16},       // Test case with increasing heights
		{[]int{7, 6, 5, 4, 3, 2}, 16},       // Test case with decreasing heights
		{[]int{5, 4, 1, 2}, 8},              // Test case with random heights
		{[]int{1, 2, 3, 4, 5}, 9},           // Test case with ascending heights
		{[]int{5, 4, 3, 2, 1}, 9},           // Test case with descending heights
		{[]int{6, 7, 5, 2, 4, 5, 9, 3}, 16}, // Test case with random heights
		{[]int{1}, 1},                       // Test case with single element
		{[]int{1, 1, 1, 1, 1, 1, 1}, 7},     // Test case with all equal elements
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := stack.LargestRectangleInHistogram(tt.heights); got != tt.want {
				t.Errorf("largestRectangleArea(%v) = %v, want %v", tt.heights, got, tt.want)
			}
		})
	}
}
