package graph_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestValidOrder(t *testing.T) {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	result := graph.AlienDictionary(words)
	expected := "wertf"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestInvalidOrderDueToCycle(t *testing.T) {
	words := []string{"z", "x", "z"}
	result := graph.AlienDictionary(words)
	expected := ""
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestInvalidOrderDueToPrefix(t *testing.T) {
	words := []string{"abc", "ab"}
	result := graph.AlienDictionary(words)
	expected := ""
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestSingleWord(t *testing.T) {
	words := []string{"abc"}
	result := graph.AlienDictionary(words)
	expected := "abc"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestEmptyWords(t *testing.T) {
	words := []string{}
	result := graph.AlienDictionary(words)
	expected := ""
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestNoValidOrder(t *testing.T) {
	words := []string{"a", "b", "c"}
	result := graph.AlienDictionary(words)
	expected := "abc"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
