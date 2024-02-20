package main

/*
- LeetCode - https://leetcode.com/problems/jump-game/
- Time - O(n)
- Space - O(1)
*/

func canJump1(nums []int) bool {
	max := 0

	for i, v := range nums {
		if i > max {
			return false
		}

		if max < i+v {
			max = i + v
		}
	}
	return true
}
