package dfs_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dfs"
)

func TestNumIslands_EmptyGrid(t *testing.T) {
	grid := [][]byte{}
	if res := dfs.NumberOfIslands(grid); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
}

func TestNumIslands_NullGrid(t *testing.T) {
	var grid [][]byte
	if res := dfs.NumberOfIslands(grid); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
}

func TestNumIslands_SingleIsland(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0'},
		{'1', '1', '0', '0'},
		{'0', '0', '0', '0'},
		{'0', '0', '0', '0'},
	}
	if res := dfs.NumberOfIslands(grid); res != 1 {
		t.Errorf("expected 1, got %d", res)
	}
}

func TestNumIslands_MultipleIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '1'},
		{'1', '0', '0', '1', '1'},
		{'0', '0', '1', '0', '0'},
		{'1', '0', '0', '1', '1'},
	}
	if res := dfs.NumberOfIslands(grid); res != 5 {
		t.Errorf("expected 5, got %d", res)
	}
}

func TestNumIslands_NoIslands(t *testing.T) {
	grid := [][]byte{
		{'0', '0', '0', '0'},
		{'0', '0', '0', '0'},
	}
	if res := dfs.NumberOfIslands(grid); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
}

func TestNumIslands_MixedGrid(t *testing.T) {
	grid := [][]byte{
		{'1', '0', '1', '0'},
		{'0', '1', '0', '1'},
		{'1', '0', '1', '0'},
	}
	if res := dfs.NumberOfIslands(grid); res != 6 {
		t.Errorf("expected 6, got %d", res)
	}
}
