package graph_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestCanFinish_NoPrerequisites(t *testing.T) {
	if !graph.CourseSchedule(3, [][]int{}) {
		t.Errorf("Expected true, got false")
	}
}

func TestCanFinish_ValidPrerequisites(t *testing.T) {
	if !graph.CourseSchedule(2, [][]int{{1, 0}}) {
		t.Errorf("Expected true, got false")
	}
}

func TestCanFinish_CycleInPrerequisites(t *testing.T) {
	if graph.CourseSchedule(2, [][]int{{1, 0}, {0, 1}}) {
		t.Errorf("Expected false, got true")
	}
}

func TestCanFinish_ComplexGraph(t *testing.T) {
	if !graph.CourseSchedule(4, [][]int{{1, 0}, {2, 1}, {3, 2}}) {
		t.Errorf("Expected true, got false")
	}
}

func TestCanFinish_IsolatedCourses(t *testing.T) {
	if !graph.CourseSchedule(5, [][]int{{1, 0}, {3, 2}}) {
		t.Errorf("Expected true, got false")
	}
}
