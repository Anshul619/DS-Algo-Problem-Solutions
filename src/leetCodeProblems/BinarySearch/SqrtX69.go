package main

/*
- LeetCode - https://leetcode.com/problems/sqrtx/
- Time - O(logn)
- Space - O(1)
*/

func mySqrt(x int) int {

	low := 1
	high := x
	mid := x

	ans := x

	for low < high {

		mid = (low + high) / 2
		currentSquare := mid * mid

		if currentSquare == x {
			return mid
		} else if currentSquare < x {
			ans = mid
			low = mid + 1
		} else {
			high = mid
		}

	}

	return ans
}
