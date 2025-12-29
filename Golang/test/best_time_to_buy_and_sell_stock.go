package test

func BestTimeToBuyAndSellStock() int {
	prices := []int{7,1,5,3,6,4}
	maxProfit := 0
	minPrice := prices[0]


	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice

			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}