package main

/*
- LeetCode - https://leetcode.com/problems/sqrtx/
- Time - O(logn)
- Space - O(1)
*/
import "log"

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

func main() {
	log.Println(mySqrt(4))
	log.Println(mySqrt(8))
	log.Println(mySqrt(3))

	log.Println(mySqrt(5))
}