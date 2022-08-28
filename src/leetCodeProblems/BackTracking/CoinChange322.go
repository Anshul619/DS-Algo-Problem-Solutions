package main

/*
- LeetCode - https://leetcode.com/problems/coin-change-2/
*/

import (
	"log"
)

func coinChangeUtils(coins []int, index int, sum int, count int, targetSum int) int {

	if sum == targetSum {
		log.Println("Sum found", count)
		return 1
	}

	if index >= len(coins) || sum > targetSum || sum < 0 {
		return 0
	}

	count++

	//log.Println(coins[index])

	return coinChangeUtils(coins, index, sum+coins[index], count, targetSum) + coinChangeUtils(coins, index+1, sum, count, targetSum)

}

func coinChange(coins []int, amount int) int {
	return coinChangeUtils(coins, 0, 0, 0, amount)
}

func main() {

	// coins := []int{1, 2, 5}
	// amount := 11

	// coins := []int{2}
	// amount := 3

	coins := []int{186, 419, 83, 408}
	amount := 6249

	// coins := []int{1, 2, 5}
	// amount := 5

	// coins := []int{2}
	// amount := 3

	log.Println(coinChange(coins, amount))
}
