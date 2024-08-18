package intervals_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/intervals"
)

func TestMergeIntervals(t *testing.T) {
	// Test case: Overlapping intervals
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result1 := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !reflect.DeepEqual(intervals.MergeIntervals(intervals1), result1) {
		t.Errorf("Expected %v but got %v", result1, intervals.MergeIntervals(intervals1))
	}

	// Test case: Non-overlapping intervals
	intervals2 := [][]int{{1, 4}, {5, 6}}
	result2 := [][]int{{1, 4}, {5, 6}}
	if !reflect.DeepEqual(intervals.MergeIntervals(intervals2), result2) {
		t.Errorf("Expected %v but got %v", result2, intervals.MergeIntervals(intervals2))
	}

	// Test case: Single interval
	intervals3 := [][]int{{1, 4}}
	result3 := [][]int{{1, 4}}
	if !reflect.DeepEqual(intervals.MergeIntervals(intervals3), result3) {
		t.Errorf("Expected %v but got %v", result3, intervals.MergeIntervals(intervals3))
	}

	// Test case: Empty array
	intervals4 := [][]int{}
	result4 := [][]int{}
	if !reflect.DeepEqual(intervals.MergeIntervals(intervals4), result4) {
		t.Errorf("Expected %v but got %v", result4, intervals.MergeIntervals(intervals4))
	}

	// Test case: Null input
	intervals5 := [][]int(nil)
	result5 := [][]int(nil)
	if !reflect.DeepEqual(intervals.MergeIntervals(intervals5), result5) {
		t.Errorf("Expected %v but got %v", result5, intervals.MergeIntervals(intervals5))
	}

	// Test case: Continuous intervals
	intervals6 := [][]int{{1, 4}, {4, 5}}
	result6 := [][]int{{1, 5}}
	if !reflect.DeepEqual(intervals.MergeIntervals(intervals6), result6) {
		t.Errorf("Expected %v but got %v", result6, intervals.MergeIntervals(intervals6))
	}
}
