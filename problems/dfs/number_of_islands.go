package dfs

func NumberOfIslands(grid [][]byte) int {
	// Special case
	if len(grid) == 0 {
		return 0
	}
	// Dimensions of the grid
	m, n := len(grid), len(grid[0])
	// Total number of islands
	count := 0
	// Process every cell in the grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j, m, n)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i int, j int, m int, n int) {
	// Base case
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j, m, n)
	dfs(grid, i+1, j, m, n)
	dfs(grid, i, j-1, m, n)
	dfs(grid, i, j+1, m, n)
}
