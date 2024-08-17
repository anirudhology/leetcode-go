package intervals

import "sort"

func MeetingRooms(intervals [][]int) bool {
	// Sort the intervals based on the start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Process all intervals
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			return false
		}
	}
	return true
}
