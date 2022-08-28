package main

/*
- LeetCode - https://leetcode.com/problems/coin-change-2/
*/

// import (
// 	"log"
// )

func changeUtils(coins []int, index int, sum int, targetSum int) int {

	if sum == targetSum {
		return 1
	}

	if index >= len(coins) || sum > targetSum || sum < 0 {
		return 0
	}

	return changeUtils(coins, index, sum+coins[index], targetSum) + changeUtils(coins, index+1, sum, targetSum)

}

func change(amount int, coins []int) int {
	return changeUtils(coins, 0, 0, amount)
}

// func main() {

// 	// coins := []int{1, 2, 5}
// 	// amount := 11

// 	// coins := []int{2}
// 	// amount := 3

// 	// coins := []int{186, 419, 83, 408}
// 	// amount := 6249

// 	// coins := []int{1, 2, 5}
// 	// amount := 5

// 	coins := []int{2}
// 	amount := 3

// 	log.Println(coinChange(coins, amount))
// }
