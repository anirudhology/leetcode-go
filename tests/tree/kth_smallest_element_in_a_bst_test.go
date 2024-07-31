package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestKthSmallestElementInABST(t *testing.T) {
	tests := []struct {
		name string
		root *util.TreeNode
		k    int
		want int
	}{
		{"single node tree", &util.TreeNode{Val: 1}, 1, 1},
		{"left skewed tree", &util.TreeNode{Val: 3, Left: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 1}}}, 2, 2},
		{"right skewed tree", &util.TreeNode{Val: 1, Right: &util.TreeNode{Val: 2, Right: &util.TreeNode{Val: 3}}}, 3, 3},
		{"balanced tree", &util.TreeNode{Val: 3, Left: &util.TreeNode{Val: 1, Right: &util.TreeNode{Val: 2}}, Right: &util.TreeNode{Val: 4}}, 3, 3},
		{"k greater than number of nodes", &util.TreeNode{Val: 1, Right: &util.TreeNode{Val: 2, Right: &util.TreeNode{Val: 3}}}, 4, 0}, // This assumes we have a condition to handle this gracefully
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.KthSmallestElementInABST(tt.root, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
