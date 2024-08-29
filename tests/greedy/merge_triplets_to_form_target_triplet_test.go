package greedy_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/greedy"
)

func TestMergeTriplets_AllPresent(t *testing.T) {
	triplets := [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}
	target := []int{2, 7, 5}
	result := greedy.MergeTripletsToFormTargetTriplet(triplets, target)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestMergeTriplets_NotAllPresent(t *testing.T) {
	triplets := [][]int{{2, 3, 4}, {1, 2, 3}, {1, 2, 3}}
	target := []int{3, 2, 5}
	result := greedy.MergeTripletsToFormTargetTriplet(triplets, target)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestMergeTriplets_SingleTripletMatch(t *testing.T) {
	triplets := [][]int{{2, 5, 3}}
	target := []int{2, 5, 3}
	result := greedy.MergeTripletsToFormTargetTriplet(triplets, target)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestMergeTriplets_EmptyTriplets(t *testing.T) {
	triplets := [][]int{}
	target := []int{1, 2, 3}
	result := greedy.MergeTripletsToFormTargetTriplet(triplets, target)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestMergeTriplets_LargeValues(t *testing.T) {
	triplets := [][]int{{100000, 100000, 100000}, {99999, 99999, 99999}}
	target := []int{100000, 100000, 100000}
	result := greedy.MergeTripletsToFormTargetTriplet(triplets, target)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestMergeTriplets_ZeroTarget(t *testing.T) {
	triplets := [][]int{{0, 0, 0}, {1, 1, 1}, {0, 1, 0}}
	target := []int{0, 0, 0}
	result := greedy.MergeTripletsToFormTargetTriplet(triplets, target)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}
