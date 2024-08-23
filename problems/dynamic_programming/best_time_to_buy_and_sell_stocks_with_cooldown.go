package dynamic_programming

func BestTimeToBuyAndSellStocksWithCooldown(prices []int) int {
	// Special case
	if len(prices) == 0 {
		return 0
	}
	// Total number of stocks
	n := len(prices)
	// Arrays to store buy and selling prices
	buy, sell := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		buy[i] = -2147483648
		sell[i] = 0
	}
	buy[0] = -prices[0]
	// Process remaining stocks
	for i := 1; i < n; i++ {
		if i > 1 {
			buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		} else {
			buy[i] = max(buy[i-1], -prices[i])
		}
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[n-1]
}
