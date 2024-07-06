package array_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"empty array", []int{}, false},
		{"nil array", nil, false},
		{"one element", []int{1}, false},
		{"no duplicates", []int{1, 2, 3, 4}, false},
		{"has duplicates", []int{1, 2, 3, 1}, true},
		{"multiple duplicates", []int{1, 1, 2, 2, 3, 3}, true},
		{"unique elements with negatives", []int{-1, -2, -3, 0}, false},
		{"duplicate negative numbers", []int{-1, -2, -3, -1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.ContainsDuplicate(tt.nums); got != tt.want {
				t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
