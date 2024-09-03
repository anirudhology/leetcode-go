package graph

import "container/heap"

// Edge represents an edge in the graph
type PointsEdge struct {
	u, v   int
	weight int
}

// PriorityQueue implements heap.Interface and holds Edges
type PriorityQueue []*PointsEdge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*PointsEdge)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func MinCostToConnectAllPoints(points [][]int) int {
	// Special case
	if len(points) == 0 {
		return 0
	}
	// Total number of points
	n := len(points)
	// Array to store visited nodes
	visited := make([]bool, n)
	for i := range visited {
		visited[i] = false
	}
	// We need to visit only n - 1 vertices as we don't want
	// cycle in the graph
	edgesLeftToVisit := n - 1
	// Total cost of connecting points
	cost := 0
	// Since we always want nearest edge to be chosen first,
	// we will use minHeap to store edges by distances
	edges := &PriorityQueue{}
	heap.Init(edges)
	// We will choose first vertex 0 to find cost to reach all
	// points from here
	for i := 1; i < n; i++ {
		distance := Abs(points[i][0]-points[0][0]) + Abs(points[i][1]-points[0][1])
		heap.Push(edges, &PointsEdge{u: 0, v: i, weight: distance})
	}
	// Mark 0 node as visited
	visited[0] = true
	// Process edges until edgesLeftToVisit becomes 0
	for edges.Len() > 0 && edgesLeftToVisit > 0 {
		// Get current edge
		edge := heap.Pop(edges).(*PointsEdge)
		v, weight := edge.v, edge.weight
		if !visited[v] {
			// Add the cost and mark as visited
			cost += weight
			visited[v] = true
			// Now traverse all possible points from v
			for i := 0; i < n; i++ {
				if !visited[i] {
					distance := Abs(points[i][0]-points[v][0]) + Abs(points[i][1]-points[v][1])
					heap.Push(edges, &PointsEdge{u: v, v: i, weight: distance})
				}
			}
			edgesLeftToVisit--
		}
	}
	return cost
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
