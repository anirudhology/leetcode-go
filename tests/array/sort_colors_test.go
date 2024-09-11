package array_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestSortColors_AllZeros(t *testing.T) {
	nums := []int{0, 0, 0}
	expected := []int{0, 0, 0}
	result := array.SortColors(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSortColors_AllOnes(t *testing.T) {
	nums := []int{1, 1, 1}
	expected := []int{1, 1, 1}
	result := array.SortColors(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSortColors_AllTwos(t *testing.T) {
	nums := []int{2, 2, 2}
	expected := []int{2, 2, 2}
	result := array.SortColors(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSortColors_Mixed(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	expected := []int{0, 0, 1, 1, 2, 2}
	result := array.SortColors(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSortColors_EmptyArray(t *testing.T) {
	nums := []int{}
	expected := []int{}
	result := array.SortColors(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSortColors_SingleElement(t *testing.T) {
	nums := []int{1}
	expected := []int{1}
	result := array.SortColors(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSortColors_AlreadySorted(t *testing.T) {
	nums := []int{0, 1, 2}
	expected := []int{0, 1, 2}
	result := array.SortColors(nums)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
