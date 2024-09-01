package dynamic_programming

func LongestIncreasingPathInAMatrix(matrix [][]int) int {
	// Dimensions of matrix
	m, n := len(matrix), len(matrix[0])
	// Lookup table to store longest increasing path
	lookup := make([][]int, m)
	for i := range lookup {
		lookup[i] = make([]int, n)
		for j := range lookup[i] {
			lookup[i][j] = 0
		}
	}
	// Longest increasing path
	maxPath := 0
	// Process each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxPath = max(maxPath, dfsForLongestPath(matrix, i, j, m, n, lookup))
		}
	}
	return maxPath
}

func dfsForLongestPath(matrix [][]int, i int, j int, m int, n int, lookup [][]int) int {
	// If we have already processed this cell
	if lookup[i][j] != 0 {
		return lookup[i][j]
	}
	// Directions
	directions := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
	longestPath := 1
	// DFS in all directions
	for _, direction := range directions {
		x := i + direction[0]
		y := j + direction[1]
		// Check for boundary conditions
		if x < 0 || x >= m || y < 0 || y >= n {
			continue
		}
		if matrix[x][y] <= matrix[i][j] {
			continue
		}
		longestPath = max(longestPath, 1+dfsForLongestPath(matrix, x, y, m, n, lookup))
	}
	lookup[i][j] = longestPath
	return longestPath
}
