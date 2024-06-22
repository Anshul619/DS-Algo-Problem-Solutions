package main

/*
- LeetCode - https://leetcode.com/problems/house-robber/
- Time - O(n)
- Space - O(1)
*/
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func rob(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	value1 := nums[0]

	if len(nums) == 1 {
		return value1
	}

	value2 := max(value1, nums[1])

	if len(nums) == 2 {
		return value2
	}

	var out int

	for i := 2; i < len(nums); i++ {
		out = max(nums[i]+value1, value2)
		value1 = value2
		value2 = out
	}
	return out
}
