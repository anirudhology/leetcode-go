package graph

func CheapestFlightsWithinKStops(n int, flights [][]int, src int, dst int, k int) int {
	// Adjacency list
	// Adjacency list
	graph := make([][][2]int, n)
	for _, flight := range flights {
		u, v, w := flight[0], flight[1], flight[2]
		graph[u] = append(graph[u], [2]int{v, w})
	}
	// Queue to perform BFS
	var nodes [][2]int
	nodes = append(nodes, [2]int{src, 0})
	// Array to minimum cost for each node
	minCost := make([]int, n)
	for i := 0; i < n; i++ {
		minCost[i] = 0x7ffffffff
	}
	// Stops taken
	stops := 0
	// Process until the queue is empty or stops <= k
	for len(nodes) > 0 && stops <= k {
		size := len(nodes)
		tempCost := make([]int, len(minCost))
		copy(tempCost, minCost)
		for i := 0; i < size; i++ {
			current := nodes[0]
			nodes = nodes[1:]
			// Traverse through the neighbors of this node
			for _, neighbor := range graph[current[0]] {
				node, price := neighbor[0], neighbor[1]
				if price+current[1] >= tempCost[node] {
					continue
				}
				tempCost[node] = price + current[1]
				nodes = append(nodes, [2]int{node, tempCost[node]})
			}
		}
		minCost = tempCost
		stops++
	}
	if minCost[dst] == 0x7ffffffff {
		return -1
	}
	return minCost[dst]
}
