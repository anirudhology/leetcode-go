package graph

import "container/heap"

// PriorityQueue implements a min-heap for strings to maintain lexical order
type ItineraryPriorityQueue []string

func (pq ItineraryPriorityQueue) Len() int           { return len(pq) }
func (pq ItineraryPriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq ItineraryPriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *ItineraryPriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(string))
}

func (pq *ItineraryPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func ReconstructItinerary(tickets [][]string) []string {
	flights := make(map[string]*ItineraryPriorityQueue)
	for _, ticket := range tickets {
		departure, arrival := ticket[0], ticket[1]
		if flights[departure] == nil {
			flights[departure] = &ItineraryPriorityQueue{}
			heap.Init(flights[departure])
		}
		heap.Push(flights[departure], arrival)
	}

	var path []string
	var dfs func(string)
	dfs = func(departure string) {
		for flights[departure] != nil && flights[departure].Len() > 0 {
			next := heap.Pop(flights[departure]).(string)
			dfs(next)
		}
		path = append([]string{departure}, path...)
	}

	dfs("JFK")
	return path
}
