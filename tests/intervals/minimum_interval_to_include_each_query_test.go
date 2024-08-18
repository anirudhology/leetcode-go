package intervals_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/intervals"
)

func TestMinInterval_basicCase(t *testing.T) {
	newIntervals := [][]int{{1, 4}, {2, 4}, {3, 6}}
	queries := []int{2, 3, 4}
	expected := []int{3, 3, 3}
	result := intervals.MinimumIntervalToIncludeEachQuery(newIntervals, queries)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMinInterval_noOverlap(t *testing.T) {
	newIntervals := [][]int{{1, 2}, {3, 4}, {5, 6}}
	queries := []int{7, 8}
	expected := []int{-1, -1}
	result := intervals.MinimumIntervalToIncludeEachQuery(newIntervals, queries)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMinInterval_multipleQueries(t *testing.T) {
	newIntervals := [][]int{{2, 5}, {1, 10}}
	queries := []int{3, 4, 5, 6}
	expected := []int{4, 4, 4, 10}
	result := intervals.MinimumIntervalToIncludeEachQuery(newIntervals, queries)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMinInterval_emptyIntervals(t *testing.T) {
	newIntervals := [][]int{}
	queries := []int{1, 2, 3}
	expected := []int{-1, -1, -1}
	result := intervals.MinimumIntervalToIncludeEachQuery(newIntervals, queries)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMinInterval_emptyQueries(t *testing.T) {
	newIntervals := [][]int{{1, 5}, {3, 7}}
	queries := []int{}
	expected := []int{}
	result := intervals.MinimumIntervalToIncludeEachQuery(newIntervals, queries)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
