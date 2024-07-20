package binary_search_test

import (
	"fmt"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/binary_search"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		median float64
	}{
		{[]int{}, []int{1}, 1},
		{[]int{2}, []int{}, 2},
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 3}, []int{2, 4, 5}, 3},
		{[]int{-3, -2, -1}, []int{1, 2, 3}, 0},
		{[]int{-5, -3, -1}, []int{-4, -2, 0, 1}, -2},
		{[]int{1, 2, 2}, []int{2, 2, 3}, 2},
		{[]int{1, 1, 2, 2}, []int{2, 3, 4, 4}, 2},
		{[]int{1000000}, []int{1000001}, 1000000.5},
		{[]int{1000000, 1000002}, []int{1000001, 1000003}, 1000001.5},
		{[]int{1, 2}, []int{1000000, 1000001}, 500001},
		{[]int{1, 1000001}, []int{2, 1000000}, 500001},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("nums1: %v, nums2: %v", test.nums1, test.nums2), func(t *testing.T) {
			result := binary_search.FindMedianSortedArrays(test.nums1, test.nums2)
			if result != test.median {
				t.Errorf("expected %f, got %f", test.median, result)
			}
		})
	}
}
