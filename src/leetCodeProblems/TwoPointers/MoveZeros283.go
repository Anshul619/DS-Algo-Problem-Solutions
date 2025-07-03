package main

/*
- LeetCode - https://leetcode.com/problems/move-zeroes/
- Time - O(n)
- Space - O(1)
*/

func moveZeroes(nums []int) {
	out := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[out] = nums[i]
			if i != out {
				nums[i] = 0
			}

			out++
		}
	}
}
