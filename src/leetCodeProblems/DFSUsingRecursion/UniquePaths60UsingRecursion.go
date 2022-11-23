package main

/*
- LeetCode - https://leetcode.com/problems/unique-paths/
- Time Complexity - Exponential

Note
- Time limit exceeding for large test case
- Optimial solution is through DP
*/
import "log"

func uniquePathsUtil(x int, y int, m int, n int, out *int) {

	if x == m-1 && y == n-1 {
		*out++
		return
	}

	if x >= m || y >= n {
		return
	}

	dx := []int{0, 1}
	dy := []int{1, 0}

	for i := 0; i < 2; i++ {
		uniquePathsUtil(x+dx[i], y+dy[i], m, n, out)
	}
}

func uniquePaths(m int, n int) int {

	out := new(int)
	uniquePathsUtil(0, 0, m, n, out)
	return *out
}

func main() {
	// m := 3
	// n := 2

	// m := 3
	// n := 7

	m := 23
	n := 12

	log.Println(uniquePaths(m, n))
}
