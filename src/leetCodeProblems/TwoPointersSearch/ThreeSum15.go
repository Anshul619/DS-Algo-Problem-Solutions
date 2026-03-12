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
		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			switch {
			case sum == 0:

				out = append(out, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				// Duplicate handling
				// - If l & l-1 (which is already checked) are same,
				// - increment l (to increase sum towards 0)
				for l < r && nums[l] == nums[l-1] {
					l++
				}

				// Duplicate handling
				// - If r & r+1 (which is already checked) are same,
				// - decrement r (to decrease sum towards 0)
				for l < r && nums[r] == nums[r+1] {
					r--
				}

			case sum < 0:
				l++
			case sum > 0:
				r--
			}
		}
	}

	return out
}
