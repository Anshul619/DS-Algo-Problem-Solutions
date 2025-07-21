package main

/*
- LeetCode - https://leetcode.com/problems/sort-an-array/description/
*/
func sortArray(nums []int) []int {
	max := 0

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	countArray := make([]int, max+1)

	for _, v := range nums {
		countArray[v]++
	}

	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i-1]
	}

	out := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		out[countArray[nums[i]]-1] = nums[i]
		countArray[nums[i]]--
	}
	return out
}
