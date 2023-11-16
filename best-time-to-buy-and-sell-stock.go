package main

func maxProfit(prices []int) int {
	var maxProfit int
	minPrice := prices[0]

	for _, price := range prices {
		if price > minPrice {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			minPrice = price
		}
	}

	return maxProfit
}
