package graph_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestValidTree_SingleNode(t *testing.T) {
	if !graph.GraphValidTree(1, [][]int{}) {
		t.Errorf("Expected true for single node")
	}
}

func TestValidTree_TreeStructure(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	if graph.GraphValidTree(5, edges) {
		t.Errorf("Expected true for valid tree structure")
	}
}

func TestValidTree_CycleExists(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
	if graph.GraphValidTree(5, edges) {
		t.Errorf("Expected false for graph with a cycle")
	}
}

func TestValidTree_DisconnectedGraph(t *testing.T) {
	edges := [][]int{{0, 1}, {2, 3}}
	if graph.GraphValidTree(4, edges) {
		t.Errorf("Expected false for disconnected graph")
	}
}

func TestValidTree_NoEdgesMultipleNodes(t *testing.T) {
	if graph.GraphValidTree(4, [][]int{}) {
		t.Errorf("Expected false for no edges with multiple nodes")
	}
}
