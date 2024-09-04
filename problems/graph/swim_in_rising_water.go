package graph

import "container/heap"

// MinHeap is a min-heap of [3]int arrays.
type WaterMinHeap [][3]int

func (h WaterMinHeap) Len() int           { return len(h) }
func (h WaterMinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h WaterMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *WaterMinHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *WaterMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func SwimInRisingWater(grid [][]int) int {
	// Dimension of the grid
	n := len(grid)
	// Array to store the visited cells
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
		for j := range visited[i] {
			visited[i][j] = false
		}
	}
	// Min heap to store nearest neighbor from current cell
	minHeap := &WaterMinHeap{}
	heap.Push(minHeap, [3]int{grid[0][0], 0, 0})
	visited[0][0] = true
	// Four directions
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// Process elements in the heap
	for minHeap.Len() > 0 {
		currentCell := heap.Pop(minHeap).([3]int)
		time, i, j := currentCell[0], currentCell[1], currentCell[2]
		if i == n-1 && j == n-1 {
			return time
		}
		// Move in four directions
		for _, direction := range directions {
			x := i + direction[0]
			y := j + direction[1]
			if x < 0 || x >= n || y < 0 || y >= n || visited[x][y] {
				continue
			}
			visited[x][y] = true
			heap.Push(minHeap, [3]int{max(grid[x][y], time), x, y})
		}
	}
	return -1
}
