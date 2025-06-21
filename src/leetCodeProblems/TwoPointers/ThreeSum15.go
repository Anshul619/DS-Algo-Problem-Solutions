package main

/*
- LeetCode - https://leetcode.com/problems/3sum/
- Time - O(n^2)
- Space - O(1)
*/
import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	out := [][]int{}

	// Outer loop
	for i := 0; i < len(nums)-2; i++ {

		// If nums[i] is greater than 0,
		// there can't be any chance that sum with subsequent elements in sorted array, can be zero - hence break
		if nums[i] > 0 {
			break
		}

		// Duplicate check
		// - Both previous and this number are same, skip
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Inner Loop - Two-Pointers search
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			switch {
			case sum == 0:

				out = append(out, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Duplicate handling
				// - If left & left-1 (previous, which is already checked) are same,
				// - hence skip this left and increment left (to increase sum towards 0)
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Duplicate handling
				// - If right & right+1 (next, which is already checked) are same,
				// - hence skip this right & decrement right (to decrease sum towards 0)
				for left < right && right+1 < len(nums) && nums[right] == nums[right+1] {
					right--
				}

			case sum < 0:
				left++
			case sum > 0:
				right--
			}
		}
	}

	return out
}
