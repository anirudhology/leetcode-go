package sliding_window_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/sliding_window"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{nums: []int{}, k: 3, expected: []int{}},
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 0, expected: []int{1, 3, -1, -3, 5, 3, 6, 7}},
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 1, expected: []int{1, 3, -1, -3, 5, 3, 6, 7}},
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, expected: []int{3, 3, 5, 5, 6, 7}},
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 4, expected: []int{3, 5, 5, 6, 7}},
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: -1, expected: []int{1, 3, -1, -3, 5, 3, 6, 7}},
		{nums: []int{5}, k: 1, expected: []int{5}},
		{nums: []int{-4, -2, -7, -5, -3, -6, -1}, k: 3, expected: []int{-2, -2, -3, -3, -1}},
		{nums: []int{2, 2, 2, 2, 2, 2, 2}, k: 3, expected: []int{2, 2, 2, 2, 2}},
	}

	for _, tt := range tests {
		result := sliding_window.MaxSlidingWindow(tt.nums, tt.k)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("maxSlidingWindow(%v, %d) = %v; expected %v", tt.nums, tt.k, result, tt.expected)
		}
	}
}
