package main

/*
- Leetcode - https://leetcode.com/problems/climbing-stairs/solutions/3742465/go-dynamic-programming/
- Time - O(n)
- Space - O(n)
*/

func climbStairsRecur(dp []int, n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if dp[n-1] == 0 {
		dp[n-1] = climbStairsRecur(dp, n-1)
	}

	if dp[n-2] == 0 {
		dp[n-2] = climbStairsRecur(dp, n-2)
	}

	return dp[n-1] + dp[n-2]
}
func climbStairs(n int) int {
	dp := make([]int, n)
	return climbStairsRecur(dp, n)
}
