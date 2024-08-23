package dynamic_programming

func CoinChangeII(amount int, coins []int) int {
	// Special case
	if len(coins) == 0 || amount < 0 {
		return 0
	}
	// Lookup table to store number of ways to make amount
	lookup := make([]int, amount+1)
	// There is only one way to make amount = 0;
	// don't take anything
	lookup[0] = 1
	// Populate for remaining amount values
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if coin <= i {
				lookup[i] += lookup[i-coin]
			}
		}
	}
	return lookup[amount]
}
