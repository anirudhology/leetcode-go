package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestBinaryTreeMaximumPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *util.TreeNode
		want int
	}{
		{"single node", &util.TreeNode{Val: 1}, 1},
		{"two nodes", &util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2}}, 3},
		{"negative nodes", &util.TreeNode{
			Val:  -10,
			Left: &util.TreeNode{Val: 9},
			Right: &util.TreeNode{
				Val:   20,
				Left:  &util.TreeNode{Val: 15},
				Right: &util.TreeNode{Val: 7},
			},
		}, 42},
		{"mixed values", &util.TreeNode{
			Val:   2,
			Left:  &util.TreeNode{Val: -1},
			Right: &util.TreeNode{Val: -2},
		}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.BinaryTreeMaximumPathSum(tt.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
