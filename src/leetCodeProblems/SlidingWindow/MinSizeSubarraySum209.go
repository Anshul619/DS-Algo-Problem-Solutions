package main

/*
- Leetcode - https://leetcode.com/problems/minimum-size-subarray-sum/description/
- Time - O(n)
- Space - O(1)
*/

func minSubArrayLen(target int, nums []int) int {
	start := 0
	out := 0
	cur := 0

	for i, v := range nums {
		cur += v

		for cur >= target {
			if out == 0 || i-start+1 < out {
				out = i - start + 1
			}
			cur -= nums[start]
			start++
		}
	}

	return out
}
