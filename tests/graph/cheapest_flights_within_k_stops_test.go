package graph_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestBasicScenario(t *testing.T) {
	n := 3
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src := 0
	dst := 2
	k := 1
	result := graph.CheapestFlightsWithinKStops(n, flights, src, dst, k)
	if result != 200 {
		t.Errorf("expected 200, got %d", result)
	}
}

func TestNoAvailableFlight(t *testing.T) {
	n := 3
	flights := [][]int{{0, 1, 100}}
	src := 0
	dst := 2
	k := 1
	result := graph.CheapestFlightsWithinKStops(n, flights, src, dst, k)
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
}

func TestDirectFlight(t *testing.T) {
	n := 3
	flights := [][]int{{0, 2, 300}}
	src := 0
	dst := 2
	k := 1
	result := graph.CheapestFlightsWithinKStops(n, flights, src, dst, k)
	if result != 300 {
		t.Errorf("expected 300, got %d", result)
	}
}

func TestNoStopAllowed(t *testing.T) {
	n := 3
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src := 0
	dst := 2
	k := 0
	result := graph.CheapestFlightsWithinKStops(n, flights, src, dst, k)
	if result != 500 {
		t.Errorf("expected 500, got %d", result)
	}
}

func TestLargeKValue(t *testing.T) {
	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 3, 100}, {0, 3, 500}}
	src := 0
	dst := 3
	k := 2
	result := graph.CheapestFlightsWithinKStops(n, flights, src, dst, k)
	if result != 300 {
		t.Errorf("expected 300, got %d", result)
	}
}
