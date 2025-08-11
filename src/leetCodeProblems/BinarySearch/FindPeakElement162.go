package main

/*
- LeetCode - https://leetcode.com/problems/find-peak-element/submissions/
- Time - O(logn)
- Space - O(1)
*/
func findPeakElement(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := start + (end-start)/2

		if (mid == 0 || nums[mid-1] < nums[mid]) &&
			(mid == len(nums)-1 || nums[mid] > nums[mid+1]) {
			return mid
		}

		if mid > 0 && nums[mid-1] > nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}
