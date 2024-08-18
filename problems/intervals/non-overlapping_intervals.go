package intervals

import "sort"

func NonOverlappingIntervals(intervals [][]int) int {
	// Special case
	if len(intervals) == 0 {
		return 0
	}
	// Sort the intervals by their end times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// Current end time
	currentEnd := intervals[0][1]
	// Non-overlapping intervals
	nonOverlappingIntervals := 1
	// Process remaining intervals
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= currentEnd {
			currentEnd = intervals[i][1]
			nonOverlappingIntervals++
		}
	}
	return len(intervals) - nonOverlappingIntervals
}
