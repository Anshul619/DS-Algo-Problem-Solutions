package main

import (
	"log"
	"math"
)

/*
- LeetCode - https://leetcode.com/problems/coin-change/
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

func main() {

	// coins := []int{1, 2, 5}
	// amount := 11

	coins := []int{2}
	amount := 3

	// coins := []int{2}
	// amount := 3

	// coins := []int{186, 419, 83, 408}
	// amount := 6249

	// coins := []int{1, 2, 5}
	// amount := 5

	// coins := []int{2}
	// amount := 3

	log.Println(coinChange(coins, amount))
}
