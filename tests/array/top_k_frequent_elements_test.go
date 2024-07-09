package array

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"empty array", []int{}, 1, []int{}},
		{"k is zero", []int{1, 1, 1, 2, 2, 3}, 0, []int{}},
		{"single element array", []int{1}, 1, []int{1}},
		{"all elements the same", []int{1, 1, 1, 1}, 1, []int{1}},
		{"multiple frequencies", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"k greater than unique elements", []int{1, 1, 2}, 5, []int{1, 2}},
		{"mixed frequencies", []int{4, 1, -1, 2, -1, 2, 3}, 2, []int{-1, 2}},
		{"negative k", []int{1, 1, 2, 2, 3, 3}, -1, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.TopKFrequent(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
