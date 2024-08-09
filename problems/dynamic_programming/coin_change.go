package dynamic_programming

import "math"

func CoinChange(coins []int, amount int) int {
	// Special case
	if len(coins) == 0 || amount < 0 {
		return -1
	}
	// Lookup table to store number of ways to make amount i
	lookup := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		lookup[i] = math.MaxInt64
		for _, coin := range coins {
			if coin <= i {
				difference := lookup[i-coin]
				if difference != math.MaxInt64 {
					lookup[i] = min(lookup[i], 1+difference)
				}
			}
		}
	}
	if lookup[amount] == math.MaxInt64 {
		return -1
	}
	return lookup[amount]
}
