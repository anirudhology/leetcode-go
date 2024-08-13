package bfs_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bfs"
	"github.com/stretchr/testify/assert"
)

func TestPacificAtlantic_EmptyGrid(t *testing.T) {
	heights := [][]int{}
	result := bfs.PacificAtlanticWaterFlow(heights)
	assert.Equal(t, 0, len(result))
}

func TestPacificAtlantic_SingleCellGrid(t *testing.T) {
	heights := [][]int{{1}}
	result := bfs.PacificAtlanticWaterFlow(heights)
	assert.Equal(t, [][]int{{0, 0}}, result)
}

func TestPacificAtlantic_FlatGrid(t *testing.T) {
	heights := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	result := bfs.PacificAtlanticWaterFlow(heights)
	assert.Equal(t, 9, len(result)) // All cells should be reachable by both oceans
}

func TestPacificAtlantic_MountainGrid(t *testing.T) {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	result := bfs.PacificAtlanticWaterFlow(heights)
	assert.Equal(t, 7, len(result)) // Specific cells should be reachable by both oceans
}
