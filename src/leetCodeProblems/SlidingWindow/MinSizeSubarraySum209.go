package main

/*
- Leetcode - https://leetcode.com/problems/minimum-size-subarray-sum/description/
- Time - O(n)
- Space - O(1)
*/

func minSubArrayLen(target int, nums []int) int {
	out := 0

	start, sum := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= target {
			if out == 0 || (i-start+1) < out {
				out = i - start + 1
			}

			sum -= nums[start]
			start++
		}
	}
	return out
}
