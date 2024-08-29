package greedy_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/greedy"
)

func TestCanCompleteCircuitBasic(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	result := greedy.GasStation(gas, cost)
	expected := 3
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestCanCompleteCircuitNoSolution(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 5}
	result := greedy.GasStation(gas, cost)
	expected := -1
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestCanCompleteCircuitSingleStation(t *testing.T) {
	gas := []int{5}
	cost := []int{4}
	result := greedy.GasStation(gas, cost)
	expected := 0
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestCanCompleteCircuitAllStationsEqual(t *testing.T) {
	gas := []int{1, 1, 1, 1}
	cost := []int{1, 1, 1, 1}
	result := greedy.GasStation(gas, cost)
	expected := 0
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestCanCompleteCircuitMultipleSolutions(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{2, 1, 2}
	result := greedy.GasStation(gas, cost)
	expected := 0
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
