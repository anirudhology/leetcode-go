package bit_manipulation_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bit_manipulation"
)

func TestSingleNumberWithSingleElement(t *testing.T) {
	nums := []int{5}
	expected := 5
	if result := bit_manipulation.SingleNumber(nums); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSingleNumberWithMultipleElements(t *testing.T) {
	nums := []int{2, 2, 1}
	expected := 1
	if result := bit_manipulation.SingleNumber(nums); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSingleNumberWithAllNegativeElements(t *testing.T) {
	nums := []int{-3, -1, -1}
	expected := -3
	if result := bit_manipulation.SingleNumber(nums); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSingleNumberWithMixedElements(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	expected := 4
	if result := bit_manipulation.SingleNumber(nums); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSingleNumberWithZeros(t *testing.T) {
	nums := []int{0, 1, 1}
	expected := 0
	if result := bit_manipulation.SingleNumber(nums); result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
