package array_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"found at start", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"found at end", []int{3, 2, 4}, 6, []int{1, 2}},
		{"no result", []int{1, 2, 3, 4}, 8, nil},
		{"empty array", []int{}, 5, []int{}},
		{"single element", []int{5}, 5, []int{5}},
		{"duplicates", []int{3, 3}, 6, []int{0, 1}},
		{"negative numbers", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"zeroes", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"large numbers", []int{1000000000, 999999999, 1, 2}, 1999999999, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.TwoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
