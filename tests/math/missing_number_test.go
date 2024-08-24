package math_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/math"
)

func TestMissingNumberWithSingleElementZero(t *testing.T) {
	nums := []int{0}
	expected := 1
	result := math.MissingNumber(nums)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMissingNumberWithSingleElementOne(t *testing.T) {
	nums := []int{1}
	expected := 0
	result := math.MissingNumber(nums)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMissingNumberWithNoMissing(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	expected := 5
	result := math.MissingNumber(nums)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMissingNumberWithMissingInMiddle(t *testing.T) {
	nums := []int{3, 0, 1}
	expected := 2
	result := math.MissingNumber(nums)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMissingNumberWithMissingAtEnd(t *testing.T) {
	nums := []int{0, 1, 3}
	expected := 2
	result := math.MissingNumber(nums)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMissingNumberWithLargeArray(t *testing.T) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expected := 8
	result := math.MissingNumber(nums)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMissingNumberEmptyArray(t *testing.T) {
	nums := []int{}
	expected := 0
	result := math.MissingNumber(nums)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
