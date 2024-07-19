package binary_search_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/binary_search"
)

func TestFindMinimumInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{2, 3, 4, 5, 1}, 1},
		{[]int{1}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{3, 1, 2}, 1},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 1},
	}

	for _, test := range tests {
		result := binary_search.FindMinimumInRotatedSortedArray(test.nums)
		if result != test.expected {
			t.Errorf("For nums %v, expected %d, but got %d", test.nums, test.expected, result)
		}
	}
}
