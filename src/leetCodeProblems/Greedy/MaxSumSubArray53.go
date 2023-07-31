package main

/*
- Leetcode - https://leetcode.com/problems/maximum-subarray/description/
- Time - O(n)
- Space - O(1)
*/
import (
	"math"
)

func maxSubArray(nums []int) int {
	maxSoFar := math.MinInt
	current := 0

	for _, v := range nums {
		current += v

		if current > maxSoFar {
			maxSoFar = current
		}

		if current < 0 {
			current = 0
		}
	}

	return maxSoFar
}

// func main() {
// 	log.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
// 	log.Println(maxSubArray([]int{1}))
// 	log.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
// }
