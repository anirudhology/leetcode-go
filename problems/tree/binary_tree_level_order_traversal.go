package tree

import (
	"container/list"

	"github.com/anirudhology/leetcode-go/problems/util"
)

func BinaryTreeLevelOrderTraversal(root *util.TreeNode) [][]int {
	// List to store the final output
	var output [][]int
	if root == nil {
		return output
	}
	// Queue to store the nodes in the BFS fashion
	nodes := list.New()
	// Add root to the list
	nodes.PushBack(root)
	// Process all nodes in the tree
	for nodes.Len() > 0 {
		// Count of nodes at the current level
		size := nodes.Len()
		// List to store the nodes at the current level
		currentLevelNodes := make([]int, 0)
		for i := 0; i < size; i++ {
			// Get front of the queue
			node := nodes.Remove(nodes.Front()).(*util.TreeNode)
			currentLevelNodes = append(currentLevelNodes, node.Val)
			// Process left and right subtrees, if present
			if node.Left != nil {
				nodes.PushBack(node.Left)
			}
			if node.Right != nil {
				nodes.PushBack(node.Right)
			}
		}
		// Add current level nodes list to the final output
		output = append(output, currentLevelNodes)
	}
	return output
}
