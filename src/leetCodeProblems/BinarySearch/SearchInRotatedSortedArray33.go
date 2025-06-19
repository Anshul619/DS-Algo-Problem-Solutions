package main

/*
- LeetCode - https://leetcode.com/problems/search-in-rotated-sorted-array/description/
- Time - O(log n)
- Space - O(1)
*/

func searchSortedArray(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}

		// left array is sorted
		if nums[start] <= nums[mid] {

			// target within range, hence binary search in left part
			if target >= nums[start] && target <= nums[mid] {
				end = mid - 1
			} else { // otherwise check right part
				start = mid + 1
			}
		} else { // right array is sorted

			// target within range, hence binary search in right part
			if target >= nums[mid] && target <= nums[end] {
				start = mid + 1
			} else { // other wise check left part
				end = mid - 1
			}
		}
	}
	return -1
}
