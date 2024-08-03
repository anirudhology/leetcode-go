package backtracking

func WordSearch(board [][]byte, word string) bool {
	// Special case
	if len(board) == 0 || len(word) == 0 {
		return false
	}
	// Dimensions of the board
	m, n := len(board), len(board[0])
	// Process cells on the board
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && Search(board, i, j, m, n, word, 0) {
				return true
			}
		}
	}
	return false
}

func Search(board [][]byte, i int, j int, m int, n int, word string, index int) bool {
	// Base case
	if index == len(word) {
		return true
	}
	// Handle out of bound indices
	if i >= m || i < 0 || j >= n || j < 0 || board[i][j] != word[index] {
		return false
	}
	if board[i][j] == '#' {
		return false
	}
	temp := board[i][j]
	board[i][j] = '#'
	found := Search(board, i-1, j, m, n, word, index+1) ||
		Search(board, i+1, j, m, n, word, index+1) ||
		Search(board, i, j-1, m, n, word, index+1) ||
		Search(board, i, j+1, m, n, word, index+1)
	board[i][j] = temp
	return found
}
