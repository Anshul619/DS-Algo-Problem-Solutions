package main

/*
- LeetCode - https://leetcode.com/problems/house-robber/
- Time - O(n)
- Space - O(1)
*/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	value1 := nums[0]
	value2 := max(value1, nums[1])

	out := max(value1, value2)

	for i := 2; i < len(nums); i++ {
		out = max(nums[i]+value1, value2)
		value1 = value2
		value2 = out
	}
	return out
}
