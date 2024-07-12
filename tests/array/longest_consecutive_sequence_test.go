package array

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Empty array", []int{}, 0},
		{"Single element", []int{1}, 1},
		{"No consecutive sequence", []int{10, 5, 100}, 1},
		{"Consecutive sequence", []int{100, 4, 200, 1, 3, 2}, 4},
		{"Mixed numbers with repeated elements", []int{100, 4, 200, 1, 3, 2, 2, 2}, 4},
		{"Duplicates", []int{1, 2, 0, 1}, 3},
		{"Large range", []int{10, 1, 3, 5, 2, 4, 11}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.LongestConsecutiveSequence(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
