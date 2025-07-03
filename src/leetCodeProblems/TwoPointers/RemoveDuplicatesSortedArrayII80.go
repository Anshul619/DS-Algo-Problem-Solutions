package main

/*
- LeetCode - https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
- Time - O(n)
- Space - O(1)
*/

func removeDuplicatesII(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	out := 2

	for i := 2; i < len(nums); i++ {

		// nums[:out] contains all cleaned up sorted numbers (max 2 duplicates)
		// hence we need to compare nums[i] with nums[out-2] (2 elements before out)
		// to check if it should be included in the out
		// We can't have nums[i] != nums[i-2], since we want to compare with out array
		// (not original array as it would fail once we started skipping elements)
		if nums[i] != nums[out-2] {
			nums[out] = nums[i]
			out++
		}
	}
	return out
}
