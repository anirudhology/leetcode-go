package union_find

func RedundantConnection(edges [][]int) []int {
	// Total number of edges
	n := len(edges)
	// Arrays to store parents and ranks
	parents, ranks := make([]int, n+1), make([]int, n+1)
	for i := 0; i <= n; i++ {
		parents[i] = i
		ranks[i] = 1
	}
	// Process for every edge
	for _, edge := range edges {
		if !RedundantConnectionUnion(edge, parents, ranks) {
			return edge
		}
	}
	return nil
}

func RedundantConnectionUnion(edge []int, parents []int, ranks []int) bool {
	// Find parents of both nodes
	p1, p2 := RedundantConnectionFind(edge[0], parents), RedundantConnectionFind(edge[1], parents)
	if p1 == p2 {
		return false
	}
	if ranks[p1] > ranks[p2] {
		parents[p2] = p1
		ranks[p1] += ranks[p2]
	} else {
		parents[p1] = p2
		ranks[p2] += ranks[p1]
	}
	return true
}

func RedundantConnectionFind(node int, parents []int) int {
	p := parents[node]
	for p != parents[p] {
		// Path compression
		parents[p] = parents[parents[p]]
		p = parents[p]
	}
	return p
}
