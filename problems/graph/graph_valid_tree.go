package graph

import "container/list"

func GraphValidTree(n int, edges [][]int) bool {
	// Base case
	if len(edges) == 0 {
		return n == 1
	}
	// Adjacency list
	adjacencyList := make([][]int, n)
	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], edge[0])
	}
	// Array to keep track of visited nodes
	visited := make([]bool, n)
	// Queue for BFS traversal
	nodes := list.New()
	nodes.PushBack(edges[0][0])
	for nodes.Len() > 0 {
		node := nodes.Remove(nodes.Front()).(int)
		if visited[node] {
			return false
		}
		visited[node] = true
		for _, neighbor := range adjacencyList[node] {
			nodes.PushBack(neighbor)
			index := 0
			for i := 0; i < len(adjacencyList[neighbor]); i++ {
				if adjacencyList[neighbor][i] == node {
					index = i
					break
				}
			}
			adjacencyList[neighbor] = append(adjacencyList[neighbor][:index], adjacencyList[neighbor][:index]...)
		}
	}
	for _, element := range visited {
		if !element {
			return false
		}
	}
	return true
}
