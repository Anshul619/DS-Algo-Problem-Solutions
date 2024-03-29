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

	minPrice := prices[0]
	maxProfit := 0

	for _, v := range prices {

		if v < minPrice {
			minPrice = v
		} else if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}

	return maxProfit
}
