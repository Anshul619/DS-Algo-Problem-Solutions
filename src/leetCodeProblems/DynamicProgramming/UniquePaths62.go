package main

/*
- LeetCode - https://leetcode.com/problems/unique-paths/
- TimeComplexity - O(mn)
- Space - O(n)
*/
import "log"

func uniquePaths(m int, n int) int {

	dp := make([]int, n)

	dp[0] = 1

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
		log.Println(dp)
	}
	return dp[n-1]
}

func main() {
	// m := 3
	// n := 2

	m := 3
	n := 7

	// m := 23
	// n := 12

	log.Println(uniquePaths(m, n))
}
