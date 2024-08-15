package dfs

func SurroundedRegions(board [][]byte) [][]byte {
	// Special case
	if len(board) == 0 {
		return board
	}
	// Dimensions of the board
	m, n := len(board), len(board[0])
	// Process boundary cells with 'O'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && (i == 0 || i == m-1 || j == 0 || j == n-1) {
				boundaryDFS(board, i, j, m, n)
			}
		}
	}
	// Post processing
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
	return board
}

func boundaryDFS(board [][]byte, i int, j int, m int, n int) {
	// Check for valid cells
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
		return
	}
	board[i][j] = '#'
	boundaryDFS(board, i-1, j, m, n)
	boundaryDFS(board, i+1, j, m, n)
	boundaryDFS(board, i, j-1, m, n)
	boundaryDFS(board, i, j+1, m, n)
}
