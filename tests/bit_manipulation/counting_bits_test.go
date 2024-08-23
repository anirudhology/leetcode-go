package bit_manipulation_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/bit_manipulation"
)

func TestCountBitsWithZero(t *testing.T) {
	expected := []int{0}
	result := bit_manipulation.CountingBits(0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCountBitsWithOne(t *testing.T) {
	expected := []int{0, 1}
	result := bit_manipulation.CountingBits(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCountBitsWithTwo(t *testing.T) {
	expected := []int{0, 1, 1}
	result := bit_manipulation.CountingBits(2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCountBitsWithFive(t *testing.T) {
	expected := []int{0, 1, 1, 2, 1, 2}
	result := bit_manipulation.CountingBits(5)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCountBitsWithTen(t *testing.T) {
	expected := []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}
	result := bit_manipulation.CountingBits(10)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
