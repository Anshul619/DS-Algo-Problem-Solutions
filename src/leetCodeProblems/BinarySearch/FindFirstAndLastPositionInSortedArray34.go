package main

/*
- Leetcode - https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
- Time - O(log n)
- Space - O(1)
*/
func searchRange(nums []int, target int) []int {
	out := []int{-1, -1}
	start := 0
	end := len(nums) - 1

	// Find the leftmost index
	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			if mid > 0 && nums[mid-1] == target {
				end = mid - 1
			} else {
				out[0] = mid
				break
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// Find the rightmost index
	start = 0
	end = len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			if mid < len(nums)-1 && nums[mid+1] == target {
				start = mid + 1
			} else {
				out[1] = mid
				break
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return out
}
