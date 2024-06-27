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
	max := math.MinInt
	cur := 0

	for _, v := range nums {
		cur += v

		if cur > max {
			max = cur
		}

		if cur < 0 {
			cur = 0
		}
	}
	return max
}
