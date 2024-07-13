package array_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{1, 2, 3, 4, 4, 9, 56, 90}, 8, []int{4, 5}},
		{[]int{1, 3, 5, 7, 9}, 10, []int{1, 5}},
		{[]int{1, 2}, 3, []int{1, 2}},
		{[]int{1, 2, 3}, 5, []int{2, 3}},
		{[]int{2, 3, 3, 7}, 6, []int{2, 3}},
		{[]int{}, 5, nil},
		{[]int{5}, 5, nil},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := array.TwoSumII(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}
