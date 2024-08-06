package heap

import (
	"container/heap"

	"github.com/anirudhology/leetcode-go/problems/util"
)

type KthLargest struct {
	minHeap *util.MinHeap
	k       int
}

func KthLargestConstructor(k int, nums []int) KthLargest {
	minHeap := &util.MinHeap{}
	heap.Init(minHeap)
	kthLargest := KthLargest{minHeap: minHeap, k: k}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if this.minHeap.Len() < this.k {
		heap.Push(this.minHeap, val)
	} else if (*this.minHeap)[0] < val {
		heap.Pop(this.minHeap)
		heap.Push(this.minHeap, val)
	}
	return (*this.minHeap)[0]
}
