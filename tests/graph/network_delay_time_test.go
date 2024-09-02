package graph_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestNetworkDelayTime(t *testing.T) {
	// Test case 1: Basic case with all nodes reachable
	if res := graph.NetworkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2); res != 2 {
		t.Errorf("Expected 2, got %d", res)
	}

	// Test case 2: Node unreachable
	if res := graph.NetworkDelayTime([][]int{{1, 2, 1}}, 2, 1); res != 1 {
		t.Errorf("Expected 1, got %d", res)
	}

	// Test case 3: Single node, self delay is 0
	if res := graph.NetworkDelayTime([][]int{}, 1, 1); res != 0 {
		t.Errorf("Expected 0, got %d", res)
	}

	// Test case 4: All nodes connected in a line
	if res := graph.NetworkDelayTime([][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 1); res != 3 {
		t.Errorf("Expected 3, got %d", res)
	}

	// Test case 5: Empty times array, unreachable nodes
	if res := graph.NetworkDelayTime([][]int{}, 2, 1); res != -1 {
		t.Errorf("Expected -1, got %d", res)
	}

	// Test case 6: Cycle in the graph
	if res := graph.NetworkDelayTime([][]int{{1, 2, 1}, {2, 3, 2}, {3, 1, 3}}, 3, 1); res != 3 {
		t.Errorf("Expected 3, got %d", res)
	}
}
