package heaps_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/heaps"
)

func TestKClosest(t *testing.T) {
	if res := heaps.KClosestPointsToOrigin([][]int{{1, 3}, {-2, 2}}, 1); len(res) != 1 || res[0][0] != -2 || res[0][1] != 2 {
		t.Errorf("expected [[-2, 2]], got %v", res)
	}
	if res := heaps.KClosestPointsToOrigin([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2); len(res) != 2 || res[0][0] != 3 || res[0][1] != 3 || res[1][0] != -2 || res[1][1] != 4 {
		t.Errorf("expected [[3, 3], [-2, 4]], got %v", res)
	}
	if res := heaps.KClosestPointsToOrigin([][]int{{0, 1}, {1, 0}}, 1); len(res) != 1 || res[0][0] != 0 || res[0][1] != 1 {
		t.Errorf("expected [[0, 1]], got %v", res)
	}
	if res := heaps.KClosestPointsToOrigin([][]int{{1, 1}}, 1); len(res) != 1 || res[0][0] != 1 || res[0][1] != 1 {
		t.Errorf("expected [[1, 1]], got %v", res)
	}
	if res := heaps.KClosestPointsToOrigin([][]int{}, 0); len(res) != 0 {
		t.Errorf("expected [], got %v", res)
	}
}
