package graph

import (
	"container/heap"
)

// Edge represents a directed edge with a weight.
type Edge struct {
	node, weight int
}

// MinHeap is a priority queue to store nodes and their cumulative distances from the start.
type MinHeap [][]int

// Implement heap.Interface for MinHeap

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// networkDelayTime calculates the time it will take for all nodes to receive a signal sent from a given node.
func NetworkDelayTime(times [][]int, n int, k int) int {
	// Create the graph using an adjacency list representation.
	graph := make(map[int][]Edge)
	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		graph[u] = append(graph[u], Edge{v, w})
	}

	// Min-heap to get the node with the shortest cumulative distance first.
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, []int{0, k}) // Start with the source node `k` with a distance of 0.

	// Visited set to store nodes that have been processed.
	visited := make(map[int]bool)
	maxDelay := 0

	// Process nodes in the heap.
	for minHeap.Len() > 0 {
		current := heap.Pop(minHeap).([]int)
		currentDistance := current[0]
		currentNode := current[1]

		// Skip processing if the node has already been visited.
		if visited[currentNode] {
			continue
		}
		visited[currentNode] = true
		maxDelay = currentDistance

		// Iterate over all neighbors of the current node.
		for _, edge := range graph[currentNode] {
			if !visited[edge.node] {
				heap.Push(minHeap, []int{currentDistance + edge.weight, edge.node})
			}
		}
	}

	// If all nodes are visited, return the maximum delay time; otherwise, return -1.
	if len(visited) == n {
		return maxDelay
	}
	return -1
}
