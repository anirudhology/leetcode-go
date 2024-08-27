package matrix_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/matrix"
)

func TestSetZeroes_EmptyMatrix(t *testing.T) {
	newMatrix := [][]int{}
	expected := matrix.SetMatrixZeroes(newMatrix)
	if !reflect.DeepEqual(newMatrix, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMatrix)
	}
}

func TestSetZeroes_SingleElementZero(t *testing.T) {
	newMatrix := [][]int{{0}}
	expected := matrix.SetMatrixZeroes(newMatrix)
	if !reflect.DeepEqual(newMatrix, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMatrix)
	}
}

func TestSetZeroes_SingleElementNonZero(t *testing.T) {
	newMatrix := [][]int{{1}}
	expected := matrix.SetMatrixZeroes(newMatrix)
	if !reflect.DeepEqual(newMatrix, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMatrix)
	}
}

func TestSetZeroes_MatrixWithNoZeroes(t *testing.T) {
	newMatrix := [][]int{{1, 2}, {3, 4}}
	expected := matrix.SetMatrixZeroes(newMatrix)
	if !reflect.DeepEqual(newMatrix, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMatrix)
	}
}

func TestSetZeroes_MatrixWithZeroes(t *testing.T) {
	newMatrix := [][]int{{1, 0, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := matrix.SetMatrixZeroes(newMatrix)
	if !reflect.DeepEqual(newMatrix, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMatrix)
	}
}

func TestSetZeroes_MatrixWithMultipleZeroes(t *testing.T) {
	newMatrix := [][]int{{0, 2, 3}, {4, 5, 0}, {7, 8, 9}}
	expected := matrix.SetMatrixZeroes(newMatrix)
	if !reflect.DeepEqual(newMatrix, expected) {
		t.Errorf("Expected %v, but got %v", expected, newMatrix)
	}
}
