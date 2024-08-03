package backtracking_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums:     []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums:     []int{1},
			expected: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		result := backtracking.Permute(tt.nums)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("permute(%v) = %v; expected %v", tt.nums, result, tt.expected)
		}
	}
}
