package binary_search_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/binary_search"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		// Test with an empty array
		{[]int{}, 5, -1},
		// Test with one element array where element is the target
		{[]int{5}, 5, 0},
		{[]int{5}, 3, -1},
		// Test with an array where target is present
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		// Test with an array where target is not present
		{[]int{1, 2, 3, 4, 5}, 0, -1},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		// Test with an array with duplicate elements
		{[]int{1, 2, 2, 2, 3, 4, 5}, 2, 3}, // The function will return the first occurrence of 2
	}

	for _, test := range tests {
		result := binary_search.BinarySearch(test.nums, test.target)
		if result != test.expected {
			t.Errorf("For input nums=%v and target=%d, expected %d but got %d", test.nums, test.target, test.expected, result)
		}
	}
}
