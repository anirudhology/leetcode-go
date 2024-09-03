package graph_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestMinCostConnectPoints_EmptyInput(t *testing.T) {
	points := [][]int{}
	result := graph.MinCostToConnectAllPoints(points)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestMinCostConnectPoints_SinglePoint(t *testing.T) {
	points := [][]int{{0, 0}}
	result := graph.MinCostToConnectAllPoints(points)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}
}

func TestMinCostConnectPoints_TwoPoints(t *testing.T) {
	points := [][]int{{0, 0}, {1, 1}}
	result := graph.MinCostToConnectAllPoints(points)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestMinCostConnectPoints_MultiplePoints(t *testing.T) {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	result := graph.MinCostToConnectAllPoints(points)
	if result != 20 {
		t.Errorf("Expected 20, but got %d", result)
	}
}

func TestMinCostConnectPoints_NegativePoints(t *testing.T) {
	points := [][]int{{-1, -2}, {1, 3}, {4, 5}}
	result := graph.MinCostToConnectAllPoints(points)
	if result != 12 {
		t.Errorf("Expected 12, but got %d", result)
	}
}
