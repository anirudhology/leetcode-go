package union_find

func NumberOfConnectedComponentsInAnUndirectedGraph(n int, edges [][]int) (ans int) {
	// Array to store parents of each node
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	// Connect edges using union operation
	for _, edge := range edges {
		union(edge[0], edge[1], parents)
	}
	// Count connected components by counting nodes that are their own parent
	count := 0
	for i := 0; i < n; i++ {
		if i == find(parents[i], parents) {
			count++
		}
	}
	return count
}

func union(a int, b int, parents []int) {
	rootA := find(a, parents)
	rootB := find(b, parents)
	parents[rootA] = rootB
}

func find(node int, parents []int) int {
	if node != parents[node] {
		parents[node] = parents[parents[node]]
	}
	return parents[node]
}
