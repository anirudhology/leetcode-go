package binary_search_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/binary_search"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		// Target found cases
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{[]int{1, 3}, 1, 0},
		{[]int{3, 1}, 1, 1},

		// Target not found cases
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{}, 1, -1},

		// Edge cases
		{[]int{1}, 1, 0},
		{nil, 1, -1},
	}

	for _, test := range tests {
		result := binary_search.SearchInRotatedSortedArray(test.nums, test.target)
		if result != test.expected {
			t.Errorf("SearchInRotatedSortedArray(%v, %d) = %d; want %d", test.nums, test.target, result, test.expected)
		}
	}
}
