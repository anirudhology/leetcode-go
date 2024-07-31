package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestValidateBinarySearchTree(t *testing.T) {
	tests := []struct {
		name string
		root *util.TreeNode
		want bool
	}{
		{"empty tree", nil, true},
		{"single node tree", &util.TreeNode{Val: 1}, true},
		{"valid BST", &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 1}, Right: &util.TreeNode{Val: 3}}, true},
		{"invalid BST", &util.TreeNode{Val: 5, Left: &util.TreeNode{Val: 1}, Right: &util.TreeNode{Val: 4, Left: &util.TreeNode{Val: 3}, Right: &util.TreeNode{Val: 6}}}, false},
		{"valid BST with negative values", &util.TreeNode{Val: 0, Left: &util.TreeNode{Val: -1}}, true},
		{"invalid BST with equal values", &util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 1}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.ValidateBinarySearchTree(tt.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
