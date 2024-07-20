package sliding_window

func BestTimeToBuyAndSellStocks(prices []int) int {
    // Special case
    if len(prices) == 0 {
        return 0
    }
    // Maximum profit so far
    maxProfit := 0
    // Minimum price so far
    min_price := prices[0]
    // Process all the prices
    for _, price := range prices {
        maxProfit = max(maxProfit, price - min_price)
        min_price = min(min_price, price)
    }
    return maxProfit
}