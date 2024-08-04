package backtracking_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected: [][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			candidates: []int{},
			target:     3,
			expected:   [][]int{},
		},
		{
			candidates: []int{1, 2, 3},
			target:     0,
			expected:   [][]int{{}},
		},
		{
			candidates: nil,
			target:     7,
			expected:   [][]int{},
		},
	}

	for _, tt := range tests {
		result := backtracking.CombinationSum2(tt.candidates, tt.target)
		if !compareSlices(result, tt.expected) {
			t.Errorf("CombinationSum2(%v, %d) = %v; expected %v", tt.candidates, tt.target, result, tt.expected)
		}
	}
}
