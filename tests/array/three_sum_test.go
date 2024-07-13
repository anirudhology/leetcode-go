package array_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Basic test with positive and negative numbers",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Test with duplicates",
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "Array with all zeros",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := array.ThreeSum(test.nums)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For nums = %v, expected %v, but got %v", test.nums, test.expected, result)
			}
		})
	}
}
