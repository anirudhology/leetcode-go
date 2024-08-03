package backtracking_test

import (
	"sort"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{1, 2, 2},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			nums:     []int{4, 4, 4, 1, 4},
			expected: [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}},
		},
		{
			nums:     []int{},
			expected: [][]int{},
		},
		{
			nums:     nil,
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		result := backtracking.SubsetsWithDup(tt.nums)
		if !compareSlices(result, tt.expected) {
			t.Errorf("SubsetsWithDup(%v) = %v; expected %v", tt.nums, result, tt.expected)
		}
	}
}

func compareSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, slice := range a {
		sort.Ints(slice)
	}
	for _, slice := range b {
		sort.Ints(slice)
	}

	sort.Slice(a, func(i, j int) bool {
		return less(a[i], a[j])
	})
	sort.Slice(b, func(i, j int) bool {
		return less(b[i], b[j])
	})

	for i := range a {
		if !equal(a[i], b[i]) {
			return false
		}
	}

	return true
}

func less(a, b []int) bool {
	for i := range a {
		if i >= len(b) {
			return false
		}
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
