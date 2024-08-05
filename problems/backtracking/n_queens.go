package backtracking

func SolveNQueens(n int) [][]string {
	// List to store combinations
	var combinations [][]string
	// Special case
	if n == 0 {
		return combinations
	}
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	// Perform backtracking
	backtrackForNQueens(board, 0, &combinations)
	return combinations
}

func backtrackForNQueens(board [][]rune, index int, combinations *[][]string) {
	if index == len(board) {
		*combinations = append(*combinations, build(board))
		return
	}
	for i := 0; i < len(board); i++ {
		if check(board, i, index) {
			board[i][index] = 'Q'
			backtrackForNQueens(board, index+1, combinations)
			board[i][index] = '.'
		}
	}
}

func check(board [][]rune, row, column int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < column; j++ {
			if board[i][j] == 'Q' && (row+j == column+i || row+column == i+j || row == i) {
				return false
			}
		}
	}
	return true
}

func build(board [][]rune) []string {
	var combination []string
	for _, row := range board {
		combination = append(combination, string(row))
	}
	return combination
}
