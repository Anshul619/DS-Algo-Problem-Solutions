package main

/*
- LeetCode - https://leetcode.com/problems/coin-change-2/
*/

func change(amount int, coins []int) int {

	dp := make([][]int, len(coins)+1)

	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	// first column, set to 1
	for i := 0; i < len(coins)+1; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < len(coins)+1; i++ {

		for j := 1; j < amount+1; j++ {
			if coins[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}

	}
	return dp[len(coins)][amount]
}

// func main() {

// 	// coins := []int{1, 2, 3}
// 	// amount := 4

// 	coins := []int{1, 2, 5}
// 	amount := 5
// 	//amount := 4
// 	// coins := []int{2}
// 	// amount := 3

// 	// coins := []int{186, 419, 83, 408}
// 	// amount := 6249

// 	// coins := []int{1, 2, 5}
// 	// amount := 5

// 	// coins := []int{3, 5, 7, 8, 9, 10, 11}
// 	// amount := 500

// 	// coins := []int{2}
// 	// amount := 3

// 	log.Println(change(amount, coins))
// }
