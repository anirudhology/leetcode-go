package util

type Point struct {
	Distance int
	Index    int
}

type MinHeapPoint []Point

func (h MinHeapPoint) Len() int           { return len(h) }
func (h MinHeapPoint) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h MinHeapPoint) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapPoint) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MinHeapPoint) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
