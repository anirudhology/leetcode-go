package binary_search_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/binary_search"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		// Test case: empty matrix
		{
			matrix: [][]int{},
			target: 5,
			want:   false,
		},
		// Test case: single element matrix, target found
		{
			matrix: [][]int{{5}},
			target: 5,
			want:   true,
		},
		// Test case: single element matrix, target not found
		{
			matrix: [][]int{{5}},
			target: 3,
			want:   false,
		},
		// Test case: matrix with multiple elements, target found
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 11,
			want:   true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 60,
			want:   true,
		},
		// Test case: matrix with multiple elements, target not found
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 0,
			want:   false,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 6,
			want:   false,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 61,
			want:   false,
		},
		// Test case: matrix with empty rows
		{
			matrix: [][]int{{}, {}, {}},
			target: 5,
			want:   false,
		},
		// Test case: matrix with one row
		{
			matrix: [][]int{{1, 2, 3, 4, 5}},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{{1, 2, 3, 4, 5}},
			target: 6,
			want:   false,
		},
		// Test case: matrix with one column
		{
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			target: 6,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := binary_search.SearchA2DMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("SearchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
