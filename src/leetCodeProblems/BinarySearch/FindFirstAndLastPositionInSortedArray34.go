package main

/*
- Leetcode - https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
- Time - O(log n)
- Space - O(1)
*/
func findLowerRange(nums []int, target int) int {
	out := -1

	start := 0
	end := len(nums) - 1

	for start <= end {

		mid := start + (end-start)/2

		if nums[mid] == target {
			out = mid
		}

		if nums[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return out
}

func findUpperRange(nums []int, target int) int {
	out := -1

	start := 0
	end := len(nums) - 1

	for start <= end {

		mid := start + (end-start)/2

		if nums[mid] == target {
			out = mid
		}

		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return out
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	out := []int{}
	out = append(out, findLowerRange(nums, target))
	out = append(out, findUpperRange(nums, target))
	return out
}
