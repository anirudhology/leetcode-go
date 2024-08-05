package backtracking_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/backtracking"
)

func TestPartition(t *testing.T) {
	result1 := backtracking.Partition("aab")
	expected1 := [][]string{{"a", "a", "b"}, {"aa", "b"}}
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v but got %v", expected1, result1)
	}

	result2 := backtracking.Partition("a")
	expected2 := [][]string{{"a"}}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v but got %v", expected2, result2)
	}

	result3 := backtracking.Partition("aba")
	expected3 := [][]string{{"a", "b", "a"}, {"aba"}}
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v but got %v", expected3, result3)
	}
}
