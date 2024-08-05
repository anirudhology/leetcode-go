package backtracking_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestLetterCombinations(t *testing.T) {
	result1 := backtracking.LetterCombinations("23")
	expected1 := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v but got %v", expected1, result1)
	}

	result2 := backtracking.LetterCombinations("2")
	expected2 := []string{"a", "b", "c"}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v but got %v", expected2, result2)
	}

	result3 := backtracking.LetterCombinations("9")
	expected3 := []string{"w", "x", "y", "z"}
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v but got %v", expected3, result3)
	}
}
