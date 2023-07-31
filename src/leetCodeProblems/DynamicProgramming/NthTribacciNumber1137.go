package main

/*
- LeetCode - https://leetcode.com/problems/n-th-tribonacci-number/description/
- Time - O(n)
- Space - O(n)
*/

func tribonacciUtil(dp []int, n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	if dp[n-3] == 0 {
		dp[n-3] = tribonacciUtil(dp, n-3)
	}

	if dp[n-2] == 0 {
		dp[n-2] = tribonacciUtil(dp, n-2)
	}

	if dp[n-1] == 0 {
		dp[n-1] = tribonacciUtil(dp, n-1)
	}

	return dp[n-3] + dp[n-2] + dp[n-1]
}
func tribonacci(n int) int {
	dp := make([]int, n+1)
	return tribonacciUtil(dp, n)
}

// func main() {
// 	log.Println(tribonacci(4))
// 	log.Println(tribonacci(25))
// }
