package main

/*
- LeetCode - https://leetcode.com/problems/maximum-number-of-points-with-cost/
*/
import "log"

func maxNum(a int, b int) int {

	if a > b {
		return a
	}
	return b
}

func maxPoints(points [][]int) int64 {

	rowsLen := len(points)
	colsLen := len(points[0])

	dp := make([]int, colsLen)
	ans := 0

	for i := 0; i < rowsLen; i++ {

		l2r := make([]int, colsLen)
		r2l := make([]int, colsLen)

		l2r[0] = dp[0]
		r2l[colsLen-1] = dp[colsLen-1]

		for j := 1; j < colsLen; j++ {
			l2r[j] = maxNum(l2r[j-1]-1, dp[j])
			r2l[colsLen-j-1] = maxNum(r2l[colsLen-j]-1, dp[colsLen-j-1])
		}

		log.Println(l2r, r2l, dp)
		for j := 0; j < colsLen; j++ {
			dp[j] = points[i][j] + maxNum(l2r[j], r2l[j])
		}
	}

	for j := 0; j < colsLen; j++ {
		ans = maxNum(ans, dp[j])
	}

	return int64(ans)
}

func main() {

	//input := [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}
	//input := [][]int{{1, 5}, {2, 3}, {4, 2}}

	input := [][]int{{2, 4, 0, 5, 5}, {0, 5, 4, 2, 5}, {2, 0, 2, 3, 1}, {3, 0, 5, 5, 2}}
	//input := [][]int{{1, 5}, {3, 2}, {4, 2}} //11
	log.Println(maxPoints(input))

}
