package bfs

import "container/list"

func RottingOranges(grid [][]int) int {
	// Special case
	if len(grid) == 0 {
		return 0
	}
	// Dimensions of grid
	m, n := len(grid), len(grid[0])
	// Queue to store rotten oranges
	cells := list.New()
	// Total number of fresh oranges
	freshOranges := 0
	// Populate the queue
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				cells.PushBack([2]int{i, j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}
	// If there are no fresh oranges return 0
	if freshOranges == 0 {
		return 0
	}
	// Total number of minutes to rot all oranges
	minutes := 0
	// Four directions
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	// Process all the rotten oranges
	for cells.Len() > 0 {
		minutes++
		size := cells.Len()
		for i := 0; i < size; i++ {
			cell := cells.Remove(cells.Front()).([2]int)
			// Check in all four directions
			for _, direction := range directions {
				x := cell[0] + direction[0]
				y := cell[1] + direction[1]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 1 {
					continue
				}
				cells.PushBack([2]int{x, y})
				grid[x][y] = 2
				freshOranges--
			}
		}
	}
	if freshOranges == 0 {
		return minutes - 1
	}
	return -1
}
