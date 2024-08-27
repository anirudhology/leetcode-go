package math_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/math"
)

func TestRotate(t *testing.T) {
	// Test case 1: 3x3 matrix
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected1 := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	result1 := math.RotateImage(matrix1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, matrix1)
	}

	// Test case 2: 4x4 matrix
	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	expected2 := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	result2 := math.RotateImage(matrix2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, matrix2)
	}

	// Test case 3: 1x1 matrix
	matrix3 := [][]int{{1}}
	expected3 := [][]int{{1}}
	result3 := math.RotateImage(matrix3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, but got %v", expected3, matrix3)
	}

	// Test case 4: Empty matrix
	matrix4 := [][]int{}
	expected4 := [][]int{}
	result4 := math.RotateImage(matrix4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Expected %v, but got %v", expected4, matrix4)
	}

	// Test case 5: Nil matrix
	var matrix5 [][]int = nil
	result5 := math.RotateImage(matrix5)
	if result5 != nil {
		t.Errorf("Expected nil, but got %v", matrix5)
	}
}
