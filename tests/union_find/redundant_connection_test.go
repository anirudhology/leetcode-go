package union_find_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/union_find"
)

func TestFindRedundantConnection(t *testing.T) {
	edges1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	expected1 := []int{2, 3}
	if result := union_find.RedundantConnection(edges1); !reflect.DeepEqual(result, expected1) {
		t.Errorf("Expected %v, got %v", expected1, result)
	}

	edges2 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	expected2 := []int{1, 5}
	if result := union_find.RedundantConnection(edges2); !reflect.DeepEqual(result, expected2) {
		t.Errorf("Expected %v, got %v", expected2, result)
	}
}
