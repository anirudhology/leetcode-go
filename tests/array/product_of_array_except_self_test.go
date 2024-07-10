package array_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty array", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"three elements", []int{1, 2, 3}, []int{6, 3, 2}},
		{"multiple elements", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"array with zero", []int{1, 0, 3, 4}, []int{0, 12, 0, 0}},
		{"all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"negative elements", []int{-1, -2, -3, -4}, []int{-24, -12, -8, -6}},
		{"mix of positive and negative elements", []int{-1, 2, -3, 4}, []int{-24, 12, -8, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := array.ProductOfArrayExceptSelf(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
