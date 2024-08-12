package graph

import "github.com/anirudhology/leetcode-go/problems/util"

func CloneGraph(node *util.GraphNode) *util.GraphNode {
	// Special case
	if node == nil {
		return nil
	}
	// Map to store mappings of original and cloned nodes
	mappings := make(map[*util.GraphNode]*util.GraphNode)
	// Perform DFS starting from the current node
	return dfs(node, mappings)
}

func dfs(node *util.GraphNode, mappings map[*util.GraphNode]*util.GraphNode) *util.GraphNode {
	// Create clone for the current node
	cloneNode := new(util.GraphNode)
	cloneNode.Val = node.Val
	mappings[node] = cloneNode
	for _, neighbor := range node.Neighbors {
		if mappings[neighbor] == nil {
			dfs(neighbor, mappings)
		}
		cloneNode.Neighbors = append(cloneNode.Neighbors, mappings[neighbor])
	}
	return cloneNode
}
