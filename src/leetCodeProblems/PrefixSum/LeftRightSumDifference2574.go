package main

/*
- LeetCode - https://leetcode.com/problems/left-and-right-sum-differences/description/
- Time - O(n)
- Extra Space - O(1)
*/

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func leftRightDifference(nums []int) []int {
	totalSum := 0

	for _, v := range nums {
		totalSum += v
	}

	out := make([]int, len(nums))
	leftSum := 0

	for i, v := range nums {
		out[i] = diff(leftSum, totalSum-v-leftSum)
		leftSum += v
	}

	return out
}
