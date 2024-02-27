package main

/*
- LeetCode - https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
- Time - O(n)
- Space - O(1)
*/

func maxProfit1(prices []int) int {
	currentBuy, currentProfit, totalProfit := 0, 0, 0

	if len(prices) > 0 {
		currentBuy = prices[0]
	}

	for _, v := range prices {

		// Profit is increasing
		if (v - currentBuy) > currentProfit {
			currentProfit = v - currentBuy
		} else {
			// sell and buy since we are getting a lower stock to buy
			totalProfit += currentProfit
			currentBuy = v
			currentProfit = 0
		}
	}

	totalProfit += currentProfit

	return totalProfit
}
