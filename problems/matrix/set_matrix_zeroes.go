package matrix

func SetMatrixZeroes(matrix [][]int) [][]int {
	// Special case
	if len(matrix) == 0 {
		return matrix
	}
	// Dimensions of the matrix
	m, n := len(matrix), len(matrix[0])
	// Flags to denote if first row and first column contain zero
	firstRow, firstColumn := false, false
	// Traverse the matrix
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstColumn = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// Traverse the matrix except first row and column
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// Process if firs row and column are true
	if firstRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColumn {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	return matrix
}
