package main

/*
- Leetcode - https://leetcode.com/problems/container-with-most-water/description/
- Space - O(1)
- Time - O(n)
*/

func maxArea(height []int) int {
	out := 0
	left := 0
	right := len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		out = max(out, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return out
}
