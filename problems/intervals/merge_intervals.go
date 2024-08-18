package intervals

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	// Special case
	if len(intervals) == 0 {
		return intervals
	}
	// Sort the intervals based on the start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Slice to store merged intervals
	var mergedIntervals [][]int
	// Current interval
	currentInterval := intervals[0]
	mergedIntervals = append(mergedIntervals, currentInterval)
	// Process all intervals
	for _, nextInterval := range intervals {
		currentEnd := currentInterval[1]
		nextStart := nextInterval[0]
		nextEnd := nextInterval[1]
		// If intervals are overlapping
		if currentEnd >= nextStart {
			currentInterval[1] = max(currentEnd, nextEnd)
		} else {
			currentInterval = nextInterval
			mergedIntervals = append(mergedIntervals, nextInterval)
		}
	}
	return mergedIntervals
}
