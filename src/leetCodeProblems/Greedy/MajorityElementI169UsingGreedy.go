package main

/*
- LeetCode - https://leetcode.com/problems/majority-element/
- Time - O(n)
- Space - O(1)
*/

func findCandidate(nums []int) int {
	majorIndex, count := 0, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[majorIndex] {
			count++
		} else {
			count--
		}

		if count == 0 {
			majorIndex, count = i, 1
		}
	}
	return nums[majorIndex]
}

func isMajority1(nums []int, candidate int) bool {
	count := 0

	for _, v := range nums {
		if v == candidate {
			count++
		}
	}

	return count > len(nums)/2
}
func majorityElement1(nums []int) int {

	c := findCandidate(nums)

	if isMajority(nums, c) {
		return c
	}
	return -1
}
