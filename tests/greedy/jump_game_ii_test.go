package greedy_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/greedy"
)

func TestJump(t *testing.T) {

	if greedy.JumpGameII([]int{0}) != 0 {
		t.Error("Expected 0 for single element array")
	}

	if greedy.JumpGameII([]int{2, 3, 1, 1, 4}) != 2 {
		t.Error("Expected 2 for multiple jumps required")
	}

	if greedy.JumpGameII([]int{5, 4, 3, 2, 1, 0}) != 1 {
		t.Error("Expected 1 for maximum jump at each step")
	}

	if greedy.JumpGameII([]int{2, 0, 2, 0, 1}) != 2 {
		t.Error("Expected 2 for array with zeros not blocking path")
	}
}
