package main

import (
	"math"
)

/*
- LeetCode - https://leetcode.com/problems/coin-change/
- Time - O(amount * numberOfCoins)
- Space - O(amount)
*/

func coinChangeUtil(count []int, pending int, coins []int) int {

	if pending == 0 {
		return 0
	}

	if pending < 0 {
		return -1
	}

	if count[pending-1] != 0 {
		return count[pending-1]
	}

	minCount := math.MaxInt

	for _, v := range coins {
		res := coinChangeUtil(count, pending-v, coins)

		if res >= 0 && res < minCount {
			minCount = res + 1
		}
	}

	if minCount == math.MaxInt {
		minCount = -1
	}

	count[pending-1] = minCount
	return count[pending-1]
}

func coinChange(coins []int, amount int) int {
	return coinChangeUtil(make([]int, amount), amount, coins)
}
