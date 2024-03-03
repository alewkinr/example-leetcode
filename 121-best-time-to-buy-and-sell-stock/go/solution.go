package solution

func maxProfit(prices []int) int {
	var minPrice = prices[0]
	var maxProfit int

	for i := 0; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		maxProfit = max(prices[i]-minPrice, maxProfit)
	}
	return maxProfit
}
