package main

/*
- Leetcode - https://leetcode.com/problems/minimum-size-subarray-sum/description/
- Time - O(n)
- Space - O(1)
*/

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func minSubArrayLen(target int, nums []int) int {
	prefix := make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}

	out := 0

	for i := 0; i < len(nums); i++ {
		want := target + prefix[i]

		j := searchInsert(prefix, want)
		if j < len(prefix) {
			if out == 0 || j-i < out {
				out = j - i
			}
		}
	}
	return out
}
