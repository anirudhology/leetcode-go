package backtracking_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
	}

	for _, tt := range tests {
		result := backtracking.Subsets(tt.nums)
		sort.Slice(result, func(i, j int) bool {
			if len(result[i]) != len(result[j]) {
				return len(result[i]) < len(result[j])
			}
			for k := range result[i] {
				if result[i][k] != result[j][k] {
					return result[i][k] < result[j][k]
				}
			}
			return false
		})
		sort.Slice(tt.expected, func(i, j int) bool {
			if len(tt.expected[i]) != len(tt.expected[j]) {
				return len(tt.expected[i]) < len(tt.expected[j])
			}
			for k := range tt.expected[i] {
				if tt.expected[i][k] != tt.expected[j][k] {
					return tt.expected[i][k] < tt.expected[j][k]
				}
			}
			return false
		})
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("subsets(%v) = %v; expected %v", tt.nums, result, tt.expected)
		}
	}
}
