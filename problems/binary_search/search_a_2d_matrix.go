package binary_search

func SearchA2DMatrix(matrix [][]int, target int) bool {
	// Special case
	if len(matrix) == 0 {
		return false
	}
	// Dimensions of the matrix
	m := len(matrix)
	n := len(matrix[0])
	// Start and end pointers
	start := 0
	end := m*n - 1
	// Process the array between range
	for start <= end {
		middle := start + (end-start)/2
		i := middle / n
		j := middle % n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return false
}
