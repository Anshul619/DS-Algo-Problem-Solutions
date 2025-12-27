package main

/*
- LeetCode - https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/
- Time - O(n)
- Space - O(1)
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0] // Track the minimum price seen so far
	maxProfit := 0        // Track the maximum profit so far

	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfit
}
