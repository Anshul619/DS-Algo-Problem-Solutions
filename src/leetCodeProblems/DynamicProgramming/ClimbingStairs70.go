package main

/*
- Leetcode - https://leetcode.com/problems/climbing-stairs/solutions/3742465/go-dynamic-programming/
- Time - O(n)
- Space - O(n)
*/

func climbStairsUtils(dp []int, n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if dp[n-1] == 0 {
		dp[n-1] = climbStairsUtils(dp, n-1)
	}

	if dp[n-2] == 0 {
		dp[n-2] = climbStairsUtils(dp, n-2)
	}

	return dp[n-1] + dp[n-2]
}
func climbStairs(n int) int {
	dp := make([]int, n+1)
	return climbStairsUtils(dp, n)
}

// func main() {
// 	log.Println(climbStairs(1))
// 	log.Println(climbStairs(2))
// 	log.Println(climbStairs(3))
// 	log.Println(climbStairs(45))
// }
