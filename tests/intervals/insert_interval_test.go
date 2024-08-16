package intervals_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/intervals"
)

func TestInsertEmptyIntervals(t *testing.T) {
	inputIntervals := [][]int{}
	newInterval := []int{2, 5}
	expected := [][]int{{2, 5}}
	result := intervals.InsertInterval(inputIntervals, newInterval)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInsertNonOverlappingBefore(t *testing.T) {
	inputIntervals := [][]int{{1, 2}, {3, 5}}
	newInterval := []int{6, 8}
	expected := [][]int{{1, 2}, {3, 5}, {6, 8}}
	result := intervals.InsertInterval(inputIntervals, newInterval)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInsertOverlappingMerge(t *testing.T) {
	inputIntervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	expected := [][]int{{1, 5}, {6, 9}}
	result := intervals.InsertInterval(inputIntervals, newInterval)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestInsertOverlappingAll(t *testing.T) {
	inputIntervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	expected := [][]int{{1, 2}, {3, 10}, {12, 16}}
	result := intervals.InsertInterval(inputIntervals, newInterval)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
