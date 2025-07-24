package main

/*
- Leetcode - https://leetcode.com/problems/summary-ranges/description
- Time - O(n)
- Space - O(1)
*/
import (
	"strconv"
)

func formatOut(nums []int, start, end int) string {
	if start == end {
		return strconv.Itoa(nums[start])
	}
	return strconv.Itoa(nums[start]) + "->" + strconv.Itoa(nums[end])
}

func summaryRanges(nums []int) []string {
	out := []string{}

	start, end := 0, 0

	for end < len(nums) {

		if end < len(nums)-1 && nums[end]+1 == nums[end+1] {
			end++
			continue
		}

		out = append(out, formatOut(nums, start, end))
		end++
		start = end
	}
	return out
}
