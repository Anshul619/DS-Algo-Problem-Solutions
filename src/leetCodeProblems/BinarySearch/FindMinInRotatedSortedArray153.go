package main

/*
- Leetcode - https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
- Time - O(log n)
- Space - O(1)
*/
import (
	"math"
)

// Another option
// - Cleaner & simple version where we keep tracking of sorted array and return nums[start] in the end
func findMin(nums []int) int {

	min := math.MaxInt
	start := 0
	end := len(nums) - 1

	for start <= end {

		// sorted array found
		if nums[start] < nums[end] {
			if min > nums[start] {
				return nums[start]
			}
		}

		mid := start + (end-start)/2

		if min > nums[mid] {
			min = nums[mid]
		}

		// left is sorted
		if nums[start] <= nums[mid] {
			if min > nums[start] {
				min = nums[start]
			}
			start = mid + 1
		} else { // right is sorted
			if min > nums[mid+1] {
				min = nums[mid+1]
			}
			end = mid - 1
		}
	}

	return min
}
