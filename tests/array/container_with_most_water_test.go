package array_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{"Empty array", []int{}, 0},
		{"Single element", []int{1}, 0},
		{"Two elements", []int{1, 1}, 1},
		{"Typical case", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"Decreasing heights", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 20},
		{"Increasing heights", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 20},
		{"Same heights", []int{4, 4, 4, 4, 4}, 16},
		{"Single high element", []int{1, 2, 4, 3}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := array.ContainerWithMostWater(tt.height)
			if result != tt.expected {
				t.Errorf("For heights %v, expected %d, but got %d", tt.height, tt.expected, result)
			}
		})
	}
}
