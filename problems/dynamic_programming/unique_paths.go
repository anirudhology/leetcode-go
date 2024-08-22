package dynamic_programming

func UniquePaths(m int, n int) int {
	// Lookup table to store unique paths for a cell {i, j}
	lookup := make([][]int, m)
	for i := range lookup {
		lookup[i] = make([]int, n)
	}
	// For the first row, we can only move horizontally
	for j := 0; j < n; j++ {
		lookup[0][j] = 1
	}
	// For the first column, we can only move vertically
	for i := 0; i < m; i++ {
		lookup[i][0] = 1
	}
	// Populate remaining table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			lookup[i][j] = lookup[i-1][j] + lookup[i][j-1]
		}
	}
	return lookup[m-1][n-1]
}
