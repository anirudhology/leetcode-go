package quick_select_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/quick_select"
)

func TestFindKthLargest(t *testing.T) {
	if res := quick_select.FindKthLargestElementInAnArray([]int{3, 2, 1, 5, 6, 4}, 2); res != 5 {
		t.Errorf("expected 5, got %v", res)
	}
	if res := quick_select.FindKthLargestElementInAnArray([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4); res != 4 {
		t.Errorf("expected 4, got %v", res)
	}
	if res := quick_select.FindKthLargestElementInAnArray([]int{1}, 1); res != 1 {
		t.Errorf("expected 1, got %v", res)
	}
	if res := quick_select.FindKthLargestElementInAnArray([]int{1, 2}, 1); res != 2 {
		t.Errorf("expected 2, got %v", res)
	}
	if res := quick_select.FindKthLargestElementInAnArray([]int{-1, -1}, 2); res != -1 {
		t.Errorf("expected -1, got %v", res)
	}
}
