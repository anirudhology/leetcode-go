package math_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/math"
)

func TestPlusOneNoCarryOver(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 4}
	result := math.PlusOne(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPlusOneWithCarryOverInMiddle(t *testing.T) {
	input := []int{1, 2, 9}
	expected := []int{1, 3, 0}
	result := math.PlusOne(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPlusOneWithAllNines(t *testing.T) {
	input := []int{9, 9, 9}
	expected := []int{1, 0, 0, 0}
	result := math.PlusOne(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPlusOneEmptyArray(t *testing.T) {
	input := []int{}
	expected := []int{1}
	result := math.PlusOne(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
