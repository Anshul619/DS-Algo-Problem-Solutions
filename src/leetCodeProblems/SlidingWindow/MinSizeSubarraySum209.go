package main

/*
- Leetcode - https://leetcode.com/problems/minimum-size-subarray-sum/description/
- Time - O(n)
- Space - O(1)
*/
import "math"

func minSubArrayLen(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	out := math.MaxInt

	start, sum := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= target {
			if (i - start + 1) < out {
				out = i - start + 1
			}

			sum -= nums[start]
			start++
		}
	}

	if out == math.MaxInt {
		return 0
	}
	return out
}
