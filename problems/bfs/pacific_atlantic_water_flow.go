package bfs

import "container/list"

func PacificAtlanticWaterFlow(heights [][]int) [][]int {
	// Final output
	var coordinates [][]int
	// Special case
	if len(heights) == 0 {
		return coordinates
	}
	// Dimensions of the island
	m, n := len(heights), len(heights[0])
	// Queues to perform BFS
	pacificCells := list.New()
	atlanticCells := list.New()
	// Arrays to mark visited cells
	pacificVisited := make([][]bool, m)
	atlanticVisited := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacificVisited[i] = make([]bool, n)
		atlanticVisited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			pacificVisited[i][j] = false
			atlanticVisited[i][j] = false
		}
	}
	// Populate queues and arrays
	for i := 0; i < m; i++ {
		pacificCells.PushBack([2]int{i, 0})
		atlanticCells.PushBack([2]int{i, n - 1})
		pacificVisited[i][0] = true
		atlanticVisited[i][n-1] = true
	}
	for j := 0; j < n; j++ {
		pacificCells.PushBack([2]int{0, j})
		atlanticCells.PushBack([2]int{m - 1, j})
		pacificVisited[0][j] = true
		atlanticVisited[m-1][j] = true
	}
	// Perform BFS for both oceans
	bfs(heights, pacificCells, pacificVisited, m, n)
	bfs(heights, atlanticCells, atlanticVisited, m, n)
	// Populate the result
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacificVisited[i][j] && atlanticVisited[i][j] {
				coordinates = append(coordinates, []int{i, j})
			}
		}
	}
	return coordinates
}

func bfs(heights [][]int, cells *list.List, visited [][]bool, m int, n int) {
	// Four directions
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for cells.Len() > 0 {
		cell := cells.Remove(cells.Front()).([2]int)
		for _, direction := range directions {
			x := cell[0] + direction[0]
			y := cell[1] + direction[1]
			if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] || heights[x][y] < heights[cell[0]][cell[1]] {
				continue
			}
			visited[x][y] = true
			cells.PushBack([2]int{x, y})
		}
	}
}
