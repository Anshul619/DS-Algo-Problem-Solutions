package main

/*
- LeetCode - https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
- Time - O(n)
- Space - O(1)
*/

func removeDuplicatesII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	unique := 0
	dup := false

	for i := 1; i < len(nums); i++ {

		// if numbers are same, increase i counter
		// otherwise, swap elements
		if nums[unique] != nums[i] || !dup {
			if nums[unique] != nums[i] {
				dup = false
			} else {
				dup = true
			}

			unique++

			// swap elements
			if i != unique {
				t := nums[i]
				nums[i] = nums[unique]
				nums[unique] = t
			}

		}
	}
	return unique + 1
}
