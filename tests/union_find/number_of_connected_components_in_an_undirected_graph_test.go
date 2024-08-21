package union_find_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/union_find"
)

func TestCountComponentsNoEdges(t *testing.T) {
	edges := [][]int{}
	result := union_find.NumberOfConnectedComponentsInAnUndirectedGraph(4, edges)
	expected := 4
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountComponentsFullyConnectedGraph(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	result := union_find.NumberOfConnectedComponentsInAnUndirectedGraph(4, edges)
	expected := 1
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountComponentsMultipleComponents(t *testing.T) {
	edges := [][]int{{0, 1}, {2, 3}}
	result := union_find.NumberOfConnectedComponentsInAnUndirectedGraph(4, edges)
	expected := 2
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountComponentsSingleNode(t *testing.T) {
	edges := [][]int{}
	result := union_find.NumberOfConnectedComponentsInAnUndirectedGraph(1, edges)
	expected := 1
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountComponentsEmptyGraph(t *testing.T) {
	edges := [][]int{}
	result := union_find.NumberOfConnectedComponentsInAnUndirectedGraph(0, edges)
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
