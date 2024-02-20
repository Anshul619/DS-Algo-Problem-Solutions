package main

/*
- LeetCode - https://leetcode.com/problems/jump-game-ii
- Time - O(n)
- Space - O(1)
*/

func jump(nums []int) int {
	max, out, last := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		if max < i+nums[i] {
			max = i + nums[i]
		}

		if last == i {
			last = max
			max = 0
			out++
		}
	}
	return out
}
