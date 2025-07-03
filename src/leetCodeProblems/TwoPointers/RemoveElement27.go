package main

/*
- LeetCode - https://leetcode.com/problems/remove-element/
- Time - O(n)
- Space - O(1)
*/

func removeElement(nums []int, val int) int {
	out := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[out] = nums[i]
			out++
		}
	}
	return out
}
