package array_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestTrappingRainWater(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Empty array",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			height:   []int{1},
			expected: 0,
		},
		{
			name:     "Two elements",
			height:   []int{1, 2},
			expected: 0,
		},
		{
			name:     "No trapping",
			height:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "Typical case",
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "All same height",
			height:   []int{3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "Decreasing heights",
			height:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "V-shaped array",
			height:   []int{2, 0, 2},
			expected: 2,
		},
		{
			name:     "Complex case",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
		{
			name:     "Undefined input",
			height:   nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := array.TrappingRainWater(tt.height)
			if result != tt.expected {
				t.Errorf("for height %v, expected %d, but got %d", tt.height, tt.expected, result)
			}
		})
	}
}
