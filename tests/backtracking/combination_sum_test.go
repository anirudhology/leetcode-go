package backtracking_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, tt := range tests {
		result := backtracking.CombinationSum(tt.candidates, tt.target)
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
			t.Errorf("combinationSum(%v, %v) = %v; expected %v", tt.candidates, tt.target, result, tt.expected)
		}
	}
}
