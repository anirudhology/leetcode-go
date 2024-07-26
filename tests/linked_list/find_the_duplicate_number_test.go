package linked_list_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/linked_list"
)

func TestFindTheDuplicateNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{1, 1}, 1},
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{1, 1}, 1},
		{[]int{4, 2, 1, 3, 2}, 2},
	}

	for _, test := range tests {
		got := linked_list.FindTheDuplicateNumber(test.nums)
		if got != test.want {
			t.Errorf("findDuplicate(%v) = %d; want %d", test.nums, got, test.want)
		}
	}
}
