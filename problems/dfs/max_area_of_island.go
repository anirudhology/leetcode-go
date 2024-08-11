package dfs

func MaxAreaOfIsland(grid [][]int) int {
	// Special case
	if len(grid) == 0 {
		return 0
	}
	// Dimensions of grid
	m, n := len(grid), len(grid[0])
	// Array to keep track of visited cells
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = false
		}
	}
	// Maximum area of island
	maxArea := 0
	// Process each cell in the grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				currentArea := dfsForMaxAreaOfIsland(grid, i, j, m, n, visited)
				maxArea = max(maxArea, currentArea)
			}
		}
	}
	return maxArea
}

func dfsForMaxAreaOfIsland(grid [][]int, i int, j int, m int, n int, visited [][]bool) int {
	// Base case
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
		return 0
	}
	// Mark the current cell visited
	visited[i][j] = true
	if grid[i][j] == 0 {
		return 0
	}
	// Perform DFS
	return 1 + dfsForMaxAreaOfIsland(grid, i-1, j, m, n, visited) + dfsForMaxAreaOfIsland(grid, i+1, j, m, n, visited) + dfsForMaxAreaOfIsland(grid, i, j-1, m, n, visited) + dfsForMaxAreaOfIsland(grid, i, j+1, m, n, visited)
}
