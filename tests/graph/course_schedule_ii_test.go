package graph_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/graph"
)

func TestFindOrder_NoPrerequisites(t *testing.T) {
	result := graph.CourseScheduleII(3, [][]int{})
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFindOrder_ValidPrerequisites(t *testing.T) {
	result := graph.CourseScheduleII(3, [][]int{})
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
