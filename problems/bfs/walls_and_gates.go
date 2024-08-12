package bfs

import "container/list"

func WallsAndGates(rooms [][]int) [][]int {
	// Special case
	if len(rooms) == 0 {
		return rooms
	}
	// Dimensions of the rooms
	m, n := len(rooms), len(rooms[0])
	// Queue to store all rooms with gates
	cells := list.New()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rooms[i][j] == 0 {
				cells.PushBack([2]int{i, j})
			}
		}
	}
	// Distance from the gate
	distance := 0
	// Four directions
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for cells.Len() > 0 {
		distance++
		for i := cells.Len(); i > 0; i-- {
			cell := cells.Remove(cells.Front()).([2]int)
			// Check in four directions
			for _, direction := range directions {
				x := direction[0] + cell[0]
				y := direction[1] + cell[1]
				if x >= 0 && x < m && y >= 0 && y < m && rooms[x][y] == 2147483647 {
					rooms[x][y] = distance
					cells.PushBack([2]int{x, y})
				}
			}
		}
	}
	return rooms
}
