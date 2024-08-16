package intervals

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	// List to store the final output
	var result [][]int
	// Special case
	if len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}
	// Total number of intervals
	n := len(intervals)
	// Add intervals smaller than newInterval to the list
	index := 0
	for index < n && intervals[index][1] < newInterval[0] {
		result = append(result, intervals[index])
		index++
	}
	// Add newInterval and merge, if required
	for index < n && intervals[index][0] <= newInterval[1] {
		newInterval[0] = min(intervals[index][0], newInterval[0])
		newInterval[1] = max(intervals[index][1], newInterval[1])
		index++
	}
	result = append(result, newInterval)
	// Add remaining intervals
	for index < n {
		result = append(result, intervals[index])
		index++
	}
	return result
}
