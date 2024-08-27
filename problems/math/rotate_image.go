package math

func RotateImage(matrix [][]int) [][]int {
	// Special case
	if len(matrix) == 0 {
		return matrix
	}
	// Order of the matrix
	n := len(matrix)
	// Left and right pointers
	left, right := 0, n-1
	for left < right {
		matrix[left], matrix[right] = matrix[right], matrix[left]
		left++
		right--
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}
