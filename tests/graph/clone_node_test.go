package main

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestCloneGraph_NullNode(t *testing.T) {
	if res := graph.CloneGraph(nil); res != nil {
		t.Errorf("expected nil, got %v", res)
	}
}

func TestCloneGraph_SingleNode(t *testing.T) {
	node := &util.GraphNode{Val: 1}
	clone := graph.CloneGraph(node)
	if clone == node || clone.Val != node.Val || len(clone.Neighbors) != 0 {
		t.Errorf("clone is not correct: %+v", clone)
	}
}

func TestCloneGraph_GraphWithOneEdge(t *testing.T) {
	node1 := &util.GraphNode{Val: 1}
	node2 := &util.GraphNode{Val: 2}
	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)

	clone := graph.CloneGraph(node1)
	if clone == node1 || clone.Val != node1.Val || len(clone.Neighbors) != 1 || clone.Neighbors[0].Val != 2 || clone.Neighbors[0] == node2 {
		t.Errorf("clone is not correct: %+v", clone)
	}
}

func TestCloneGraph_GraphWithCycle(t *testing.T) {
	node1 := &util.GraphNode{Val: 1}
	node2 := &util.GraphNode{Val: 2}
	node3 := &util.GraphNode{Val: 3}
	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node3)
	node3.Neighbors = append(node3.Neighbors, node1)

	clone := graph.CloneGraph(node1)
	if clone == node1 || clone.Val != node1.Val || len(clone.Neighbors) != 1 || clone.Neighbors[0].Val != 2 || clone.Neighbors[0].Neighbors[0].Val != 3 || clone.Neighbors[0].Neighbors[0].Neighbors[0].Val != 1 {
		t.Errorf("clone is not correct: %+v", clone)
	}
}
