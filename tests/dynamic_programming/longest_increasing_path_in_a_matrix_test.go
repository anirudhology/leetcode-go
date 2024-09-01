package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
)

func TestLongestIncreasingPath(t *testing.T) {
	// Test case 1: Single element matrix
	matrix1 := [][]int{{5}}
	if result := dynamic_programming.LongestIncreasingPathInAMatrix(matrix1); result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}

	// Test case 2: 2x2 matrix with no increasing path
	matrix2 := [][]int{{3, 2}, {1, 0}}
	if result := dynamic_programming.LongestIncreasingPathInAMatrix(matrix2); result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}

	// Test case 3: 2x2 matrix with a simple increasing path
	matrix3 := [][]int{{1, 2}, {3, 4}}
	if result := dynamic_programming.LongestIncreasingPathInAMatrix(matrix3); result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}

	// Test case 4: 3x3 matrix with a more complex path
	matrix4 := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	if result := dynamic_programming.LongestIncreasingPathInAMatrix(matrix4); result != 4 {
		t.Errorf("Expected 4, but got %d", result)
	}

	// Test case 5: 3x3 matrix with all elements the same
	matrix5 := [][]int{
		{7, 7, 7},
		{7, 7, 7},
		{7, 7, 7},
	}
	if result := dynamic_programming.LongestIncreasingPathInAMatrix(matrix5); result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}

	// Test case 6: 4x4 matrix with multiple increasing paths
	matrix6 := [][]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{6, 5, 4, 3},
		{7, 8, 9, 10},
	}
	if result := dynamic_programming.LongestIncreasingPathInAMatrix(matrix6); result != 9 {
		t.Errorf("Expected 9, but got %d", result)
	}
}
