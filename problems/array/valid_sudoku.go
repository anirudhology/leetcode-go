package array

import "strconv"

func ValidSudoku(board [][]byte) bool {
	// We need a map to store the encodings of elements on the board
	encodings := make(map[string]bool)
	// Process the board one cell at a time
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Current number
			c := board[i][j]
			// Process only filled cells
			if c != '.' {
				rowString := string(c) + " is present in row: " + strconv.Itoa(i)
				columnString := string(c) + " is present in column: " + strconv.Itoa(j)
				blockString := string(c) + " is present in block: " + strconv.Itoa(i/3) + "-" + strconv.Itoa(j/3)

				// Check if any encoding is already present
				if encodings[rowString] || encodings[columnString] || encodings[blockString] {
					return false
				}

				// Add encodings to the map
				encodings[rowString] = true
				encodings[columnString] = true
				encodings[blockString] = true
			}
		}
	}
	return true
}
