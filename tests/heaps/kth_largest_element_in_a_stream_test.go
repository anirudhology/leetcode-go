package heaps_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/heaps"
)

func TestKthLargest(t *testing.T) {
	nums := []int{4, 5, 8, 2}
	kthLargest := heaps.KthLargestConstructor(3, nums)

	if res := kthLargest.Add(3); res != 4 {
		t.Errorf("expected 4, got %d", res)
	}
	if res := kthLargest.Add(5); res != 5 {
		t.Errorf("expected 5, got %d", res)
	}
	if res := kthLargest.Add(10); res != 5 {
		t.Errorf("expected 5, got %d", res)
	}
	if res := kthLargest.Add(9); res != 8 {
		t.Errorf("expected 8, got %d", res)
	}
	if res := kthLargest.Add(4); res != 8 {
		t.Errorf("expected 8, got %d", res)
	}
}
