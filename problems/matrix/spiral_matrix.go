package matrix

func SpiralMatrix(matrix [][]int) []int {
	// List to store final spiral order
	var spiral []int
	// Special case
	if len(matrix) == 0 {
		return spiral
	}
	// Dimensions of the matrix
	m, n := len(matrix), len(matrix[0])
	// Variables to keep track of boundaries
	top, bottom, left, right := 0, m-1, 0, n-1
	// Process all elements in the matrix
	for top <= bottom && left <= right {
		// Move from left to right
		for i := left; i <= right; i++ {
			spiral = append(spiral, matrix[top][i])
		}
		top++
		// Move from top to bottom
		for i := top; i <= bottom; i++ {
			spiral = append(spiral, matrix[i][right])
		}
		right--
		// Check if we are out of bounds
		if left > right || top > bottom {
			break
		}
		// Move from right to left
		for i := right; i >= left; i-- {
			spiral = append(spiral, matrix[bottom][i])
		}
		bottom--
		// Move from bottom to top
		for i := bottom; i >= top; i-- {
			spiral = append(spiral, matrix[i][left])
		}
		left++
	}
	return spiral
}
