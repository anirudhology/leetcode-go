package intervals_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/intervals"
)

func TestCanAttendMeetings(t *testing.T) {
	// Test case: No overlapping intervals
	intervals1 := [][]int{{0, 30}, {35, 50}, {60, 75}}
	if !intervals.MeetingRooms(intervals1) {
		t.Errorf("Expected true but got false for intervals1")
	}

	// Test case: Overlapping intervals
	intervals2 := [][]int{{0, 30}, {20, 50}, {60, 75}}
	if intervals.MeetingRooms(intervals2) {
		t.Errorf("Expected false but got true for intervals2")
	}

	// Test case: Empty array
	intervals3 := [][]int{}
	if !intervals.MeetingRooms(intervals3) {
		t.Errorf("Expected true but got false for intervals3")
	}

	// Test case: Single interval
	intervals4 := [][]int{{10, 20}}
	if !intervals.MeetingRooms(intervals4) {
		t.Errorf("Expected true but got false for intervals4")
	}

	// Test case: Continuous but non-overlapping intervals
	intervals5 := [][]int{{0, 10}, {10, 20}, {20, 30}}
	if !intervals.MeetingRooms(intervals5) {
		t.Errorf("Expected true but got false for intervals5")
	}
}
