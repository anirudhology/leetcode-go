package backtracking_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestSolveNQueens(t *testing.T) {
	normalize := func(input [][]string) [][]string {
		if input == nil {
			return [][]string{}
		}
		return input
	}

	result1 := backtracking.SolveNQueens(4)
	expected1 := [][]string{
		{"..Q.", "Q...", "...Q", ".Q.."},
		{".Q..", "...Q", "Q...", "..Q."},
	}
	if !reflect.DeepEqual(normalize(result1), normalize(expected1)) {
		t.Errorf("Expected %v but got %v", expected1, result1)
	}

	result2 := backtracking.SolveNQueens(1)
	expected2 := [][]string{
		{"Q"},
	}
	if !reflect.DeepEqual(normalize(result2), normalize(expected2)) {
		t.Errorf("Expected %v but got %v", expected2, result2)
	}

	result3 := backtracking.SolveNQueens(0)
	expected3 := [][]string{}
	if !reflect.DeepEqual(normalize(result3), normalize(expected3)) {
		t.Errorf("Expected %v but got %v", expected3, result3)
	}

	result4 := backtracking.SolveNQueens(2)
	expected4 := [][]string{}
	if !reflect.DeepEqual(normalize(result4), normalize(expected4)) {
		t.Errorf("Expected %v but got %v", expected4, result4)
	}

	result5 := backtracking.SolveNQueens(3)
	expected5 := [][]string{}
	if !reflect.DeepEqual(normalize(result5), normalize(expected5)) {
		t.Errorf("Expected %v but got %v", expected5, result5)
	}
}
