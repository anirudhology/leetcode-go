package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestGoodNodesInBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *util.TreeNode
		want int
	}{
		{"single node tree", &util.TreeNode{Val: 1}, 1},
		{"right leaning tree", &util.TreeNode{Val: 1, Right: &util.TreeNode{Val: 2, Right: &util.TreeNode{Val: 3}}}, 3},
		{"left leaning tree", &util.TreeNode{Val: 3, Left: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 1}}}, 1},
		{"balanced tree", &util.TreeNode{Val: 3, Left: &util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 3}}, Right: &util.TreeNode{Val: 4, Left: &util.TreeNode{Val: 1}, Right: &util.TreeNode{Val: 5}}}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.CountGoodNodesInBinaryTree(tt.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
