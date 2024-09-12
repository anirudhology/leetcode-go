package array_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/array"
)

func TestMerge_BothNonEmpty(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	expected := []int{1, 2, 2, 3, 5, 6}
	result := array.MergeSortedArray(nums1, m, nums2, n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMerge_EmptyNums2(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{}
	m := 3
	n := 0
	expected := []int{1, 2, 3}
	result := array.MergeSortedArray(nums1, m, nums2, n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMerge_EmptyNums1(t *testing.T) {
	nums1 := []int{0, 0, 0}
	nums2 := []int{1, 2, 3}
	m := 0
	n := 3
	expected := []int{1, 2, 3}
	result := array.MergeSortedArray(nums1, m, nums2, n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestMerge_EmptyBoth(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{}
	m := 0
	n := 0
	expected := []int{}
	result := array.MergeSortedArray(nums1, m, nums2, n)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
