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
	if len(nums) == 0 {
		return []string{}
	}

	out := []string{}

	start, end := 0, 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[end]+1 {
			end = i
			continue
		} else {
			out = append(out, formatOut(nums, start, end))
			start, end = i, i
		}
	}
	out = append(out, formatOut(nums, start, end))
	return out
}
