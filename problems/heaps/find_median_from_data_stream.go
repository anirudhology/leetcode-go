package heaps

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	minHeap, maxHeap *IntHeap
}

func MedianFinderConstructor() MedianFinder {
	a := &IntHeap{}
	b := &IntHeap{}
	heap.Init(a)
	heap.Init(b)
	return MedianFinder{a, b}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxHeap, num)
	heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
	if this.minHeap.Len() > this.maxHeap.Len() {
		heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64((*this.maxHeap)[0])
	}
	return float64(-(*this.minHeap)[0]+(*this.maxHeap)[0]) / 2
}
