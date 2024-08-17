package heaps_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/heaps"
)

func TestEmptyTasks(t *testing.T) {
	result := heaps.TaskScheduler([]byte{}, 2)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestNullTasks(t *testing.T) {
	result := heaps.TaskScheduler(nil, 2)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestSingleTask(t *testing.T) {
	result := heaps.TaskScheduler([]byte{'A'}, 2)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestMultipleTasksWithCooldown(t *testing.T) {
	result := heaps.TaskScheduler([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}

func TestMultipleTasksNoCooldown(t *testing.T) {
	result := heaps.TaskScheduler([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestComplexTaskPattern(t *testing.T) {
	result := heaps.TaskScheduler([]byte{'A', 'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D'}, 2)
	if result != 9 {
		t.Errorf("Expected 10, got %d", result)
	}
}
