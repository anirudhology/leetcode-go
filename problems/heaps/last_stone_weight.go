package heaps

import (
	"container/heap"

	"github.com/anirudhology/leetcode-go/problems/util"
)

func LastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	maxHeap := &util.MaxHeap{}
	heap.Init(maxHeap)
	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}
	for maxHeap.Len() > 1 {
		x := heap.Pop(maxHeap).(int)
		y := heap.Pop(maxHeap).(int)
		if x != y {
			heap.Push(maxHeap, x-y)
		}
	}
	if maxHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(maxHeap).(int)
}
