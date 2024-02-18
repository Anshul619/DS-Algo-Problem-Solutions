package main

/*
- Leetcode - https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
- Time - O(log n)
- Space - O(1)
*/
import (
	"math"
)

func findMinUtil(nums []int, start, end int, min *int) {
	if start < 0 || end < 0 || end > len(nums)-1 {
		return
	}

	if start == end {
		if nums[start] < *min {
			*min = nums[start]
		}
		return
	}

	mid := start + (end-start)/2

	// left part is sorted
	if nums[start] <= nums[mid] {
		if nums[start] < *min {
			*min = nums[start]
		}
		start = mid + 1
	} else { // right part is sorted
		if nums[mid] < *min {
			*min = nums[mid]
		}
		end = mid - 1
	}
	findMinUtil(nums, start, end, min)
}

func findMin(nums []int) int {
	min := math.MaxInt
	findMinUtil(nums, 0, len(nums)-1, &min)
	return min
}
