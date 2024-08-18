package intervals

import (
	"container/heap"
	"sort"

	"github.com/anirudhology/leetcode-go/problems/util"
)

func MinimumIntervalToIncludeEachQuery(intervals [][]int, queries []int) []int {
	// Total number of queries
	n := len(queries)
	// Create a new slice in order to store pair of query value
	// and its corresponding index
	queryIndex := make([][]int, n)
	for i := 0; i < n; i++ {
		queryIndex[i] = []int{queries[i], i}
	}
	// Sort the array by their query value
	sort.Slice(queryIndex, func(i, j int) bool {
		return queryIndex[i][0] < queryIndex[j][0]
	})
	// Sort the intervals by their left value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// Array to store the final output
	result := make([]int, n)
	// Variable to keep track of intervals
	j := 0
	// Heap to store the pair of size of the intervals and
	// its corresponding right value
	minHeap := util.IntHeap{}
	heap.Init(&minHeap)
	// Process all the queries
	for i := 0; i < n; i++ {
		query := queryIndex[i][0]
		index := queryIndex[i][1]
		// Add pairs to heap for which left value is smaller
		// than the current query value
		for j < len(intervals) && intervals[j][0] <= query {
			right := intervals[j][1]
			size := right - intervals[j][0] + 1
			heap.Push(&minHeap, [2]int{size, right})
			j++
		}
		// Remove pairs from the heap for which the right value
		// is smaller than the current query value
		for minHeap.Len() > 0 && minHeap[0][1] < query {
			heap.Pop(&minHeap)
		}
		if minHeap.Len() > 0 {
			result[index] = minHeap[0][0]
		} else {
			result[index] = -1
		}
	}
	return result
}
