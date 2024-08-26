package matrix_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/matrix"
)

func TestSpiralOrder_SingleElement(t *testing.T) {
	newMatrix := [][]int{{1}}
	expected := []int{1}
	result := matrix.SpiralMatrix(newMatrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSpiralOrder_SingleRow(t *testing.T) {
	newMatrix := [][]int{{1, 2, 3, 4}}
	expected := []int{1, 2, 3, 4}
	result := matrix.SpiralMatrix(newMatrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSpiralOrder_SingleColumn(t *testing.T) {
	newMatrix := [][]int{{1}, {2}, {3}, {4}}
	expected := []int{1, 2, 3, 4}
	result := matrix.SpiralMatrix(newMatrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSpiralOrder_RectangularMatrix(t *testing.T) {
	newMatrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	result := matrix.SpiralMatrix(newMatrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSpiralOrder_RectangularMatrix_2x3(t *testing.T) {
	newMatrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := []int{1, 2, 3, 6, 5, 4}
	result := matrix.SpiralMatrix(newMatrix)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
