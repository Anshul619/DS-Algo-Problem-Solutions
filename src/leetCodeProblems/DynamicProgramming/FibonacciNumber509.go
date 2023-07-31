package main

/*
- LeetCode - https://leetcode.com/problems/fibonacci-number/description/
- Space - O(n)
- Time - O(n)
*/

func fibUtil(dp []int, n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if dp[n-1] == 0 {
		dp[n-1] = fibUtil(dp, n-1)
	}

	if dp[n-2] == 0 {
		dp[n-2] = fibUtil(dp, n-2)
	}
	return dp[n-1] + dp[n-2]
}
func fib(n int) int {
	dp := make([]int, n+1)
	return fibUtil(dp, n)
}

// func main() {
// 	log.Println(fib(2))
// 	log.Println(fib(3))
// 	log.Println(fib(4))
// }
