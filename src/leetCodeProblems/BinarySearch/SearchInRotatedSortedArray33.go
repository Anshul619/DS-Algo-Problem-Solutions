package main

/*
- LeetCode - https://leetcode.com/problems/search-in-rotated-sorted-array/description/
- Time - O(log n)
- Space - O(1)
*/

func binarySearch(nums []int, start, end, target int) int {
	if start < 0 || end < 0 || end > len(nums)-1 {
		return -1
	}

	mid := start + (end-start)/2

	if nums[mid] == target {
		return mid
	}

	if start == end {
		return -1
	}

	// left part is sorted
	if nums[start] <= nums[mid] {

		if target >= nums[start] && target <= nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	} else if target >= nums[mid] && target <= nums[end] { // Right part is sorted
		start = mid + 1
	} else {
		end = mid - 1
	}

	return binarySearch(nums, start, end, target)
}
func searchSortedArray(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}
