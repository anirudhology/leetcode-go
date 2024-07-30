package tree

import (
	"container/list"

	"github.com/anirudhology/leetcode-go/problems/util"
)

func BinaryTreeRightSideView(root *util.TreeNode) []int {
	// List to store nodes visible from the right side
	var rightNodes []int
	// Special case
	if root == nil {
		return rightNodes
	}
	// Queue to perform BFS
	nodes := list.New()
	// Add root to the list
	nodes.PushBack(root)
	// Process all nodes in the tree
	for nodes.Len() > 0 {
		// Count of nodes at the current level
		size := nodes.Len()
		for i := 0; i < size; i++ {
			// Get node from the front of the queue
			node := nodes.Remove(nodes.Front()).(*util.TreeNode)
			if i == size-1 {
				rightNodes = append(rightNodes, node.Val)
			}
			// Add left and right children, if present
			if node.Left != nil {
				nodes.PushBack(node.Left)
			}
			if node.Right != nil {
				nodes.PushBack(node.Right)
			}
		}
	}
	return rightNodes
}
