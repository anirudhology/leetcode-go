package dfs_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dfs"
)

func TestSolveEmptyBoard(t *testing.T) {
	board := [][]byte{}
	result := dfs.SurroundedRegions(board)
	if !reflect.DeepEqual(result, board) {
		t.Errorf("Expected %v, but got %v", board, result)
	}
}

func TestSolveNilBoard(t *testing.T) {
	var board [][]byte = nil
	result := dfs.SurroundedRegions(board)
	if !reflect.DeepEqual(result, board) {
		t.Errorf("Expected %v, but got %v", board, result)
	}
}

func TestSolveSingleCellO(t *testing.T) {
	board := [][]byte{{'O'}}
	expected := [][]byte{{'O'}}
	result := dfs.SurroundedRegions(board)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSolveSingleCellX(t *testing.T) {
	board := [][]byte{{'X'}}
	expected := [][]byte{{'X'}}
	result := dfs.SurroundedRegions(board)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSolveSmallBoard(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	expected := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	result := dfs.SurroundedRegions(board)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
