package dfs_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dfs"
)

func TestMaxAreaOfIsland_EmptyGrid(t *testing.T) {
	grid := [][]int{}
	if res := dfs.MaxAreaOfIsland(grid); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
}

func TestMaxAreaOfIsland_NullGrid(t *testing.T) {
	var grid [][]int
	if res := dfs.MaxAreaOfIsland(grid); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
}

func TestMaxAreaOfIsland_SingleIsland(t *testing.T) {
	grid := [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	if res := dfs.MaxAreaOfIsland(grid); res != 4 {
		t.Errorf("expected 4, got %d", res)
	}
}

func TestMaxAreaOfIsland_MultipleIslands(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 1, 1},
		{0, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{1, 1, 0, 1, 1},
	}
	if res := dfs.MaxAreaOfIsland(grid); res != 4 {
		t.Errorf("expected 4, got %d", res)
	}
}

func TestMaxAreaOfIsland_NoIslands(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	if res := dfs.MaxAreaOfIsland(grid); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
}

func TestMaxAreaOfIsland_MixedGrid(t *testing.T) {
	grid := [][]int{
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{1, 0, 1, 0},
	}
	if res := dfs.MaxAreaOfIsland(grid); res != 1 {
		t.Errorf("expected 1, got %d", res)
	}
}
