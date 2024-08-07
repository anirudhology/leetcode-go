package dynamic_programming

func ClimbingStairs(n int) int {
	// Special case
	if n <= 2 {
		return n
	}
	// Ways to climb previous two stairs
	a, b := 1, 2
	c := a + b
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
