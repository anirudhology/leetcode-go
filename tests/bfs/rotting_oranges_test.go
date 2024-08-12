package bfs_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bfs"
	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting_EmptyGrid(t *testing.T) {
	grid := [][]int{}
	result := bfs.RottingOranges(grid)
	assert.Equal(t, 0, result)
}

func TestOrangesRotting_NoFreshOranges(t *testing.T) {
	grid := [][]int{
		{2, 2, 0},
		{2, 0, 2},
		{0, 2, 2},
	}
	result := bfs.RottingOranges(grid)
	assert.Equal(t, 0, result)
}

func TestOrangesRotting_AllFreshOranges(t *testing.T) {
	grid := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	result := bfs.RottingOranges(grid)
	assert.Equal(t, -1, result)
}

func TestOrangesRotting_RottenSpread(t *testing.T) {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	result := bfs.RottingOranges(grid)
	assert.Equal(t, 4, result)
}

func TestOrangesRotting_RottenSpreadWithUnreachable(t *testing.T) {
	grid := [][]int{
		{2, 1, 1},
		{0, 1, 0},
		{1, 0, 1},
	}
	result := bfs.RottingOranges(grid)
	assert.Equal(t, -1, result)
}
