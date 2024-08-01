package tree_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/tree"
	"github.com/anirudhology/leetcode-go/problems/util"
)

func TestConstructBinaryTreeFromPreorderAndInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *util.TreeNode
	}{
		{"single node", []int{1}, []int{1}, &util.TreeNode{Val: 1}},
		{"two nodes", []int{1, 2}, []int{2, 1}, &util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2}}},
		{"balanced tree", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, &util.TreeNode{
			Val: 3, Left: &util.TreeNode{Val: 9}, Right: &util.TreeNode{Val: 20, Left: &util.TreeNode{Val: 15}, Right: &util.TreeNode{Val: 7}},
		}},
		{"empty tree", []int{}, []int{}, nil},
		{"unbalanced tree", []int{3, 2, 1, 4, 5}, []int{1, 2, 3, 4, 5}, &util.TreeNode{
			Val: 3, Left: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 1}}, Right: &util.TreeNode{Val: 4, Right: &util.TreeNode{Val: 5}},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tree.ConstructBinaryTreeFromPreorderAndInorderTraversal(tt.preorder, tt.inorder)
			if !isEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(a, b *util.TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val != b.Val {
		return false
	}
	return isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
}
