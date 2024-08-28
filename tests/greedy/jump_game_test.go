package greedy_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/greedy"
)

func TestCanJump(t *testing.T) {
	if !greedy.JumpGame([]int{0}) {
		t.Error("Expected true for single element array")
	}

	if greedy.JumpGame([]int{1, 0, 0, 0}) {
		t.Error("Expected false for all zeros except first")
	}

	if !greedy.JumpGame([]int{2, 3, 1, 1, 4}) {
		t.Error("Expected true for sufficient jumps")
	}

	if greedy.JumpGame([]int{3, 2, 1, 0, 4}) {
		t.Error("Expected false for insufficient jumps")
	}
}
