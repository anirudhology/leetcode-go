package heaps_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/heaps"
)

func TestLastStoneWeight(t *testing.T) {
	if res := heaps.LastStoneWeight([]int{2, 7, 4, 1, 8, 1}); res != 1 {
		t.Errorf("expected 1, got %d", res)
	}
	if res := heaps.LastStoneWeight([]int{2, 2}); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
	if res := heaps.LastStoneWeight([]int{1}); res != 1 {
		t.Errorf("expected 1, got %d", res)
	}
	if res := heaps.LastStoneWeight([]int{}); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
	if res := heaps.LastStoneWeight(nil); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
}
