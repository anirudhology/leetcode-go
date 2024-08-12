package bfs_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bfs"
)

func TestIslandsAndTreasure_NullRooms(t *testing.T) {
	if res := bfs.WallsAndGates(nil); res != nil {
		t.Errorf("expected nil, got %v", res)
	}
}

func TestIslandsAndTreasure_EmptyRooms(t *testing.T) {
	rooms := [][]int{}
	if res := bfs.WallsAndGates(rooms); !reflect.DeepEqual(res, rooms) {
		t.Errorf("expected %v, got %v", rooms, res)
	}
}

func TestIslandsAndTreasure_SingleRoomWithGate(t *testing.T) {
	rooms := [][]int{{0}}
	expected := [][]int{{0}}
	if res := bfs.WallsAndGates(rooms); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestIslandsAndTreasure_SingleRoomWithWall(t *testing.T) {
	rooms := [][]int{{-1}}
	expected := [][]int{{-1}}
	if res := bfs.WallsAndGates(rooms); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestIslandsAndTreasure_SingleRoomWithEmpty(t *testing.T) {
	rooms := [][]int{{math.MaxInt32}}
	expected := [][]int{{math.MaxInt32}}
	if res := bfs.WallsAndGates(rooms); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestIslandsAndTreasure_GridWithMultipleGates(t *testing.T) {
	rooms := [][]int{
		{math.MaxInt32, -1, 0, math.MaxInt32},
		{math.MaxInt32, math.MaxInt32, math.MaxInt32, -1},
		{math.MaxInt32, -1, math.MaxInt32, -1},
		{0, -1, math.MaxInt32, math.MaxInt32},
	}
	expected := [][]int{
		{3, -1, 0, 1},
		{2, 2, 1, -1},
		{1, -1, 2, -1},
		{0, -1, 3, 4},
	}
	if res := bfs.WallsAndGates(rooms); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestIslandsAndTreasure_GridWithNoGates(t *testing.T) {
	rooms := [][]int{
		{math.MaxInt32, -1, math.MaxInt32, math.MaxInt32},
		{math.MaxInt32, math.MaxInt32, math.MaxInt32, -1},
		{math.MaxInt32, -1, math.MaxInt32, -1},
		{math.MaxInt32, -1, math.MaxInt32, math.MaxInt32},
	}
	expected := rooms // No change expected since there are no gates
	if res := bfs.WallsAndGates(rooms); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
