package intervals_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/intervals"
)

func TestEraseOverlapIntervals(t *testing.T) {
	// Test case: Overlapping intervals
	intervals1 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	result1 := 1
	if intervals.NonOverlappingIntervals(intervals1) != result1 {
		t.Errorf("Expected %v but got %v", result1, intervals.NonOverlappingIntervals(intervals1))
	}

	// Test case: Non-overlapping intervals
	intervals2 := [][]int{{1, 2}, {2, 3}, {3, 4}}
	result2 := 0
	if intervals.NonOverlappingIntervals(intervals2) != result2 {
		t.Errorf("Expected %v but got %v", result2, intervals.NonOverlappingIntervals(intervals2))
	}

	// Test case: Single interval
	intervals3 := [][]int{{1, 2}}
	result3 := 0
	if intervals.NonOverlappingIntervals(intervals3) != result3 {
		t.Errorf("Expected %v but got %v", result3, intervals.NonOverlappingIntervals(intervals3))
	}

	// Test case: Empty array
	intervals4 := [][]int{}
	result4 := 0
	if intervals.NonOverlappingIntervals(intervals4) != result4 {
		t.Errorf("Expected %v but got %v", result4, intervals.NonOverlappingIntervals(intervals4))
	}

	// Test case: Null input
	intervals5 := [][]int(nil)
	result5 := 0
	if intervals.NonOverlappingIntervals(intervals5) != result5 {
		t.Errorf("Expected %v but got %v", result5, intervals.NonOverlappingIntervals(intervals5))
	}

	// Test case: Multiple overlapping intervals
	intervals6 := [][]int{{1, 3}, {2, 4}, {3, 5}, {1, 2}}
	result6 := 2
	if intervals.NonOverlappingIntervals(intervals6) != result6 {
		t.Errorf("Expected %v but got %v", result6, intervals.NonOverlappingIntervals(intervals6))
	}
}
