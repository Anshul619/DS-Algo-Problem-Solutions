package main

/*
- LeetCode - https://leetcode.com/problems/find-peak-element/submissions/
- Time - O(logn)
- Space - O(1)
*/

func binarySearch1(nums []int, start int, end int) int {
	mid := start + (end-start)/2

	if (mid == 0 || nums[mid] >= nums[mid-1]) && (mid == len(nums)-1 || nums[mid] >= nums[mid+1]) {
		return mid
	} else if mid > 0 && nums[mid-1] >= nums[mid] {
		return binarySearch1(nums, start, mid-1)
	} else {
		return binarySearch1(nums, mid+1, end)
	}
}

func findPeakElement(nums []int) int {
	return binarySearch1(nums, 0, len(nums))
}
