package heaps

import (
	"container/heap"

	"github.com/anirudhology/leetcode-go/problems/util"
)

func KClosestPointsToOrigin(points [][]int, k int) [][]int {
	// List to store k closest points
	var kClosestPoints [][]int
	// Special case
	if len(points) == 0 {
		return kClosestPoints
	}
	minHeap := &util.MinHeapPoint{}
	heap.Init(minHeap)
	for i, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		heap.Push(minHeap, util.Point{Distance: distance, Index: i})
	}
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		point := heap.Pop(minHeap).(util.Point)
		result[i] = points[point.Index]
	}
	return result
}
