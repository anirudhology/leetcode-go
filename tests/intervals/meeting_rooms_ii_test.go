package intervals_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/intervals"
)

func TestMinMeetingRooms(t *testing.T) {
	// Test case: Overlapping intervals
	intervals1 := [][]int{{0, 30}, {5, 10}, {15, 20}}
	expected1 := 2
	result1 := intervals.MeetingRoomsII(intervals1)
	if result1 != expected1 {
		t.Errorf("Expected %d, but got %d", expected1, result1)
	}

	// Test case: Non-overlapping intervals
	intervals2 := [][]int{{7, 10}, {2, 4}}
	expected2 := 1
	result2 := intervals.MeetingRoomsII(intervals2)
	if result2 != expected2 {
		t.Errorf("Expected %d, but got %d", expected2, result2)
	}

	// Test case: Single interval
	intervals3 := [][]int{{1, 5}}
	expected3 := 1
	result3 := intervals.MeetingRoomsII(intervals3)
	if result3 != expected3 {
		t.Errorf("Expected %d, but got %d", expected3, result3)
	}

	// Test case: Empty intervals array
	intervals4 := [][]int{}
	expected4 := 0
	result4 := intervals.MeetingRoomsII(intervals4)
	if result4 != expected4 {
		t.Errorf("Expected %d, but got %d", expected4, result4)
	}

	// Test case: Null intervals array
	var intervals5 [][]int
	expected5 := 0
	result5 := intervals.MeetingRoomsII(intervals5)
	if result5 != expected5 {
		t.Errorf("Expected %d, but got %d", expected5, result5)
	}

	// Test case: Multiple meetings at the same time
	intervals6 := [][]int{{1, 10}, {1, 10}, {1, 10}}
	expected6 := 3
	result6 := intervals.MeetingRoomsII(intervals6)
	if result6 != expected6 {
		t.Errorf("Expected %d, but got %d", expected6, result6)
	}
}
