package graph_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestFindItinerary(t *testing.T) {
	// Test case 1
	tickets1 := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	expected1 := []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
	result1 := graph.ReconstructItinerary(tickets1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	// Test case 2
	tickets2 := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	expected2 := []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}
	result2 := graph.ReconstructItinerary(tickets2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, result2)
	}
}
