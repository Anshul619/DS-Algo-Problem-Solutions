package main

import (
	"math"
)

/*
- LeetCode - https://leetcode.com/problems/coin-change/
*/

func coinChangeUtil(coinIndex int, coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	minCount := math.MaxInt

	if coinIndex < len(coins) && amount > 0 {

		maxCount := amount / coins[coinIndex]

		//log.Println("Coin Val", coins[coinIndex])
		// log.Println("Current Coin Val", coins[coinIndex], ", Max count", maxCount)
		for i := 0; i <= maxCount; i++ {

			if amount >= i*coins[coinIndex] {

				// log.Println("In loop---")
				// log.Println("In loop - Coin Val", coins[coinIndex])
				// log.Println("i", i)
				// log.Println("Pending amount", amount-i*coins[coinIndex])

				//log.Println(coinIndex, i, amount-i*coins[coinIndex], coins)
				res := coinChangeUtil(coinIndex+1, coins, amount-i*coins[coinIndex])

				// log.Println("Res", res)

				if res != -1 {
					if minCount > res+i {
						// log.Println(minCount)
						minCount = res + i
					}
				}
			}
		}
	}

	if minCount == math.MaxInt {
		return -1
	}

	return minCount
}

func coinChange(coins []int, amount int) int {
	return coinChangeUtil(0, coins, amount)
}

// func main() {

// 	coins := []int{1, 2, 5}
// 	amount := 11

// 	// coins := []int{2}
// 	// amount := 3

// 	// coins := []int{186, 419, 83, 408}
// 	// amount := 6249

// 	// coins := []int{1, 2, 5}
// 	// amount := 5

// 	// coins := []int{2}
// 	// amount := 3

// 	log.Println(coinChange(coins, amount))
// }
