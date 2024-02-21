package main

/*
- LeetCode - https://leetcode.com/problems/remove-element/
- Time - O(n)
- Space - O(1)
*/

func removeElement(nums []int, val int) int {
	start, end := 0, len(nums)-1

	for start <= end {

		if nums[end] == val {
			end--
			continue
		}

		if nums[start] == val {
			nums[start] = nums[end]
			end--
		}
		start++
	}
	return start
}
