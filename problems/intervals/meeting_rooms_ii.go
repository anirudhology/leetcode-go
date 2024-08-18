package intervals

import "sort"

func MeetingRoomsII(intervals [][]int) int {
	// Special case
	if len(intervals) == 0 {
		return 0
	}
	// Total number of intervals
	n := len(intervals)
	// Arrays to store start and end times
	startTimes := make([]int, n)
	endTimes := make([]int, n)
	for i := 0; i < n; i++ {
		startTimes[i] = intervals[i][0]
		endTimes[i] = intervals[i][1]
	}
	// Sort both the arrays
	sort.Slice(startTimes, func(i, j int) bool {
		return startTimes[i] < startTimes[j]
	})
	sort.Slice(endTimes, func(i, j int) bool {
		return endTimes[i] < endTimes[j]
	})
	// Total number of meeting rooms required
	requiredMeetingRooms := 0
	// Meetings in progress
	meetingsInProgress := 0
	// Pointers to traverse startTimes and endTimes
	start, end := 0, 0
	for start < n {
		if startTimes[start] < endTimes[end] {
			meetingsInProgress++
			start++
		} else {
			meetingsInProgress--
			end++
		}
		requiredMeetingRooms = max(requiredMeetingRooms, meetingsInProgress)
	}
	return requiredMeetingRooms
}
