package main

/*
- LeetCode - https://leetcode.com/problems/remove-element/
- Time - O(n)
- Space - O(1)
*/

func removeElement(nums []int, val int) int {
	out := 0

	for _, v := range nums {
		if v != val {
			nums[out] = v
			out++
		}
	}
	return out
}
