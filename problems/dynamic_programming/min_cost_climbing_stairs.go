package dynamic_programming

func MinCostClimbingStairs(cost []int) int {
	// Total number of stairs
	n := len(cost)
	// Ways to climb first two stairs
	a, b := cost[0], cost[1]
	if n == 2 {
		return min(a, b)
	}
	for i := 2; i < n; i++ {
		c := cost[i] + min(a, b)
		a = b
		b = c
	}
	return min(a, b)
}
